import logging
import pytest
import mock
import threading

from milvus import Milvus
from mishards.connections import (ConnectionMgr, Connection,
        ConnectionPool, ConnectionTopology, ConnectionGroup)
from mishards.topology import StatusType
from mishards import exceptions

logger = logging.getLogger(__name__)


@pytest.mark.usefixtures('app')
class TestConnection:
    def test_manager(self):
        mgr = ConnectionMgr()

        mgr.register('pod1', '111')
        mgr.register('pod2', '222')
        mgr.register('pod2', '222')
        mgr.register('pod2', '2222')
        assert len(mgr.conn_names) == 2

        mgr.unregister('pod1')
        assert len(mgr.conn_names) == 1

        mgr.unregister('pod2')
        assert len(mgr.conn_names) == 0

        mgr.register('WOSERVER', 'xxxx')
        assert len(mgr.conn_names) == 0

        assert not mgr.conn('XXXX', None)
        with pytest.raises(exceptions.ConnectionNotFoundError):
            mgr.conn('XXXX', None, True)

        mgr.conn('WOSERVER', None)

    def test_connection(self):
        class Conn:
            def __init__(self, state):
                self.state = state

            def connect(self, uri):
                return self.state

            def connected(self):
                return self.state

        FAIL_CONN = Conn(False)
        PASS_CONN = Conn(True)

        class Retry:
            def __init__(self):
                self.times = 0

            def __call__(self, conn):
                self.times += 1
                logger.info('Retrying {}'.format(self.times))

        class Func():
            def __init__(self):
                self.executed = False

            def __call__(self):
                self.executed = True

        max_retry = 3

        RetryObj = Retry()

        c = Connection('client',
                       uri='xx',
                       max_retry=max_retry,
                       on_retry_func=RetryObj)
        c.conn = FAIL_CONN
        ff = Func()
        this_connect = c.connect(func=ff)
        with pytest.raises(exceptions.ConnectionConnectError):
            this_connect()
        assert RetryObj.times == max_retry
        assert not ff.executed
        RetryObj = Retry()

        c.conn = PASS_CONN
        this_connect = c.connect(func=ff)
        this_connect()
        assert ff.executed
        assert RetryObj.times == 0

        this_connect = c.connect(func=None)
        with pytest.raises(TypeError):
            this_connect()

        errors = []

        def error_handler(err):
            errors.append(err)

        this_connect = c.connect(func=None, exception_handler=error_handler)
        this_connect()
        assert len(errors) == 1

    def test_topology(self):
        ConnectionGroup.on_added = mock.MagicMock(return_value=(True,))
        w_topo = ConnectionTopology()
        status, wg1 = w_topo.create(name='wg1')
        assert w_topo.has_group(wg1)
        assert status == StatusType.OK

        status, wg1_dup = w_topo.create(name='wg1')
        assert wg1_dup is None
        assert status == StatusType.DUPLICATED

        fetched_group = w_topo.get_group('wg1')
        assert id(fetched_group) == id(wg1)

        with pytest.raises(RuntimeError):
            wg1.create(name='wg1_p1')

        status, wg1_p1 = wg1.create(name='wg1_p1', uri='127.0.0.1:19530')
        assert status == StatusType.OK
        assert wg1_p1 is not None
        assert len(wg1) == 1

        status, wg1_p1_dup = wg1.create(name='wg1_p1', uri='127.0.0.1:19530')
        assert status == StatusType.DUPLICATED
        assert wg1_p1_dup is None
        assert len(wg1) == 1

        status, wg1_p2 = wg1.create('wg1_p2', uri='127.0.0.1:19530')
        assert status == StatusType.OK
        assert wg1_p2 is not None
        assert len(wg1) == 2

        poped = wg1.remove('wg1_p3')
        assert poped is None
        assert len(wg1) == 2

        poped = wg1.remove('wg1_p2')
        assert poped.name == 'wg1_p2'
        assert len(wg1) == 1

        fetched_p1 = wg1.get(wg1_p1.name)
        assert fetched_p1 == wg1_p1

        fetched_p1 = w_topo.get_group('wg1').get('wg1_p1')

        conn1 = fetched_p1.fetch()
        assert len(fetched_p1) == 1
        assert fetched_p1.active_num == 1

        conn2 = fetched_p1.fetch()
        assert len(fetched_p1) == 2
        assert fetched_p1.active_num == 2

        conn2.release()
        assert len(fetched_p1) == 2
        assert fetched_p1.active_num == 1

        assert len(w_topo.group_names) == 1

    def test_connection_pool(self):

        def check_mp_fetch(capacity=-1):
            w2 = ConnectionPool(name='w2', uri='127.0.0.1:19530', max_retry=2, capacity=capacity)
            connections = []
            def GetConnection(pool):
                conn = pool.fetch(timeout=0.1)
                if conn:
                    connections.append(conn)

            threads = []
            threads_num = 10 if capacity < 0 else 2*capacity
            for _ in range(threads_num):
                t = threading.Thread(target=GetConnection, args=(w2,))
                threads.append(t)
                t.start()

            for t in threads:
                t.join()

            expected_size = threads_num if capacity < 0 else capacity

            assert len(connections) == expected_size

        check_mp_fetch(5)
        check_mp_fetch()

        w1 = ConnectionPool(name='w1', uri='127.0.0.1:19530', max_retry=2, capacity=2)
        w1_1 = w1.fetch()
        assert len(w1) == 1
        assert w1.active_num == 1
        w1_2 = w1.fetch()
        assert len(w1) == 2
        assert w1.active_num == 2
        w1_3 = w1.fetch()
        assert w1_3 is None
        assert len(w1) == 2
        assert w1.active_num == 2

        w1_1.release()
        assert len(w1) == 2
        assert w1.active_num == 1

        def check(pool, expected_size, expected_active_num):
            w = pool.fetch()
            assert len(pool) == expected_size
            assert pool.active_num == expected_active_num

        check(w1, 2, 2)

        assert len(w1) == 2
        assert w1.active_num == 1

        wild_w = w1.create()
        with pytest.raises(RuntimeError):
            w1.release(wild_w)

        ret = w1_2.can_retry
        assert ret == w1_2.connection.can_retry
