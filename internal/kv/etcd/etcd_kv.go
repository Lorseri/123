// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcdkv

import (
	"context"
	"encoding/binary"
	"fmt"
	"path"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/samber/lo"

	"github.com/milvus-io/milvus/internal/common"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/metrics"
	"github.com/milvus-io/milvus/internal/util/timerecord"
)

const (
	// RequestTimeout is default timeout for etcd request.
	RequestTimeout = 30 * time.Second
)

// EtcdKV implements TxnKV interface, it supports to process multiple kvs in a transaction.
type EtcdKV struct {
	client   *clientv3.Client
	rootPath string
}

// NewEtcdKV creates a new etcd kv.
func NewEtcdKV(client *clientv3.Client, rootPath string) *EtcdKV {
	kv := &EtcdKV{
		client:   client,
		rootPath: rootPath,
	}
	return kv
}

// Close closes the connection to etcd.
func (kv *EtcdKV) Close() {
	log.Debug("etcd kv closed", zap.String("path", kv.rootPath))
}

// GetPath returns the path of the key.
func (kv *EtcdKV) GetPath(key string) string {
	return path.Join(kv.rootPath, key)
}

func (kv *EtcdKV) WalkWithPrefix(prefix string, paginationSize int, fn func([]byte, []byte) error) error {
	start := time.Now()
	prefix = path.Join(kv.rootPath, prefix)

	batch := int64(paginationSize)
	opts := []clientv3.OpOption{
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(batch),
		clientv3.WithRange(clientv3.GetPrefixRangeEnd(prefix)),
	}

	key := prefix
	for {
		resp, err := kv.getEtcdMeta(context.TODO(), key, opts...)
		if err != nil {
			return err
		}

		for _, kv := range resp.Kvs {
			if err = fn(kv.Key, kv.Value); err != nil {
				return err
			}
		}

		if !resp.More {
			break
		}
		// move to next key
		key = string(append(resp.Kvs[len(resp.Kvs)-1].Key, 0))
	}

	CheckElapseAndWarn(start, "Slow etcd operation(WalkWithPagination)", zap.String("prefix", prefix))
	return nil
}

// LoadWithPrefix returns all the keys and values with the given key prefix.
func (kv *EtcdKV) LoadWithPrefix(key string) ([]string, []string, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([]string, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, string(kv.Value))
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with prefix", zap.Strings("keys", keys))
	return keys, values, nil
}

// LoadBytesWithPrefix returns all the keys and values with the given key prefix.
func (kv *EtcdKV) LoadBytesWithPrefix(key string) ([]string, [][]byte, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([][]byte, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, kv.Value)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with prefix", zap.Strings("keys", keys))
	return keys, values, nil
}

// LoadWithPrefix2 returns all the the keys,values and key versions with the given key prefix.
func (kv *EtcdKV) LoadWithPrefix2(key string) ([]string, []string, []int64, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, nil, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([]string, 0, resp.Count)
	versions := make([]int64, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, string(kv.Value))
		versions = append(versions, kv.Version)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with prefix2", zap.Strings("keys", keys))
	return keys, values, versions, nil
}

func (kv *EtcdKV) LoadWithRevisionAndVersions(key string) ([]string, []string, []int64, int64, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, nil, 0, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([]string, 0, resp.Count)
	versions := make([]int64, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, string(kv.Value))
		versions = append(versions, kv.Version)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with prefix2", zap.Strings("keys", keys))
	return keys, values, versions, resp.Header.Revision, nil
}

// LoadBytesWithPrefix2 returns all the the keys,values and key versions with the given key prefix.
func (kv *EtcdKV) LoadBytesWithPrefix2(key string) ([]string, [][]byte, []int64, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, nil, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([][]byte, 0, resp.Count)
	versions := make([]int64, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, kv.Value)
		versions = append(versions, kv.Version)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with prefix2", zap.Strings("keys", keys))
	return keys, values, versions, nil
}

// Load returns value of the key.
func (kv *EtcdKV) Load(key string) (string, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key)
	if err != nil {
		return "", err
	}
	if resp.Count <= 0 {
		return "", common.NewKeyNotExistError(key)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load", zap.String("key", key))
	return string(resp.Kvs[0].Value), nil
}

// LoadBytes returns value of the key.
func (kv *EtcdKV) LoadBytes(key string) ([]byte, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key)
	if err != nil {
		return []byte{}, err
	}
	if resp.Count <= 0 {
		return []byte{}, common.NewKeyNotExistError(key)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load", zap.String("key", key))
	return resp.Kvs[0].Value, nil
}

// MultiLoad gets the values of the keys in a transaction.
func (kv *EtcdKV) MultiLoad(keys []string) ([]string, error) {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(keys))
	for _, keyLoad := range keys {
		ops = append(ops, clientv3.OpGet(path.Join(kv.rootPath, keyLoad)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		return []string{}, err
	}

	result := make([]string, 0, len(keys))
	invalid := make([]string, 0, len(keys))
	for index, rp := range resp.Responses {
		if rp.GetResponseRange().Kvs == nil || len(rp.GetResponseRange().Kvs) == 0 {
			invalid = append(invalid, keys[index])
			result = append(result, "")
		}
		for _, ev := range rp.GetResponseRange().Kvs {
			result = append(result, string(ev.Value))
		}
	}
	if len(invalid) != 0 {
		log.Warn("MultiLoad: there are invalid keys", zap.Strings("keys", invalid))
		err = fmt.Errorf("there are invalid keys: %s", invalid)
		return result, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi load", zap.Any("keys", keys))
	return result, nil
}

// MultiLoadBytes gets the values of the keys in a transaction.
func (kv *EtcdKV) MultiLoadBytes(keys []string) ([][]byte, error) {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(keys))
	for _, keyLoad := range keys {
		ops = append(ops, clientv3.OpGet(path.Join(kv.rootPath, keyLoad)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		return [][]byte{}, err
	}

	result := make([][]byte, 0, len(keys))
	invalid := make([]string, 0, len(keys))
	for index, rp := range resp.Responses {
		if rp.GetResponseRange().Kvs == nil || len(rp.GetResponseRange().Kvs) == 0 {
			invalid = append(invalid, keys[index])
			result = append(result, []byte{})
		}
		for _, ev := range rp.GetResponseRange().Kvs {
			result = append(result, ev.Value)
		}
	}
	if len(invalid) != 0 {
		log.Warn("MultiLoad: there are invalid keys", zap.Strings("keys", invalid))
		err = fmt.Errorf("there are invalid keys: %s", invalid)
		return result, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi load", zap.Strings("keys", keys))
	return result, nil
}

// LoadWithRevision returns keys, values and revision with given key prefix.
func (kv *EtcdKV) LoadWithRevision(key string) ([]string, []string, int64, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	if err != nil {
		return nil, nil, 0, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([]string, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, string(kv.Value))
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with revision", zap.Strings("keys", keys))
	return keys, values, resp.Header.Revision, nil
}

// LoadBytesWithRevision returns keys, values and revision with given key prefix.
func (kv *EtcdKV) LoadBytesWithRevision(key string) ([]string, [][]byte, int64, error) {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	resp, err := kv.getEtcdMeta(context.TODO(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, nil, 0, err
	}
	keys := make([]string, 0, resp.Count)
	values := make([][]byte, 0, resp.Count)
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
		values = append(values, kv.Value)
	}
	CheckElapseAndWarn(start, "Slow etcd operation load with revision", zap.Strings("keys", keys))
	return keys, values, resp.Header.Revision, nil
}

// Save saves the key-value pair.
func (kv *EtcdKV) Save(key, value string) error {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	CheckValueSizeAndWarn(key, value)
	_, err := kv.putEtcdMeta(context.TODO(), key, value)
	CheckElapseAndWarn(start, "Slow etcd operation save", zap.String("key", key))
	return err
}

// SaveBytes saves the key-value pair.
func (kv *EtcdKV) SaveBytes(key string, value []byte) error {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	CheckValueSizeAndWarn(key, value)
	_, err := kv.putEtcdMeta(context.TODO(), key, string(value))
	CheckElapseAndWarn(start, "Slow etcd operation save", zap.String("key", key))
	return err
}

// SaveWithLease is a function to put value in etcd with etcd lease options.
func (kv *EtcdKV) SaveWithLease(key, value string, id clientv3.LeaseID) error {
	log.Debug("Etcd saving with lease", zap.String("etcd_key", key))
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	CheckValueSizeAndWarn(key, value)
	_, err := kv.putEtcdMeta(context.TODO(), key, value, clientv3.WithLease(id))
	CheckElapseAndWarn(start, "Slow etcd operation save with lease", zap.String("key", key))
	return err
}

// SaveWithIgnoreLease updates the key without changing its current lease. Must be used when key already exists.
func (kv *EtcdKV) SaveWithIgnoreLease(key, value string) error {
	log.Debug("Etcd saving with ignore lease", zap.String("etcd_key", key))
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	CheckValueSizeAndWarn(key, value)
	_, err := kv.putEtcdMeta(context.TODO(), key, value, clientv3.WithIgnoreLease())
	CheckElapseAndWarn(start, "Slow etcd operation save with lease", zap.String("key", key))
	return err
}

// SaveBytesWithLease is a function to put value in etcd with etcd lease options.
func (kv *EtcdKV) SaveBytesWithLease(key string, value []byte, id clientv3.LeaseID) error {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	CheckValueSizeAndWarn(key, value)
	_, err := kv.putEtcdMeta(context.TODO(), key, string(value), clientv3.WithLease(id))
	CheckElapseAndWarn(start, "Slow etcd operation save with lease", zap.String("key", key))
	return err
}

// MultiSave saves the key-value pairs in a transaction.
func (kv *EtcdKV) MultiSave(kvs map[string]string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(kvs))
	var keys []string
	for key, value := range kvs {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), value))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	CheckTxnStringValueSizeAndWarn(kvs)
	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSave error", zap.Strings("keys", lo.Keys(kvs)), zap.Int("len", len(kvs)), zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save", zap.Strings("keys", keys))
	return err
}

// MultiSaveBytes saves the key-value pairs in a transaction.
func (kv *EtcdKV) MultiSaveBytes(kvs map[string][]byte) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(kvs))
	var keys []string
	for key, value := range kvs {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), string(value)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	CheckTxnBytesValueSizeAndWarn(kvs)
	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSaveBytes err", zap.Any("kvs", kvs), zap.Int("len", len(kvs)), zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save", zap.Strings("keys", keys))
	return err
}

// RemoveWithPrefix removes the keys with given prefix.
func (kv *EtcdKV) RemoveWithPrefix(prefix string) error {
	start := time.Now()
	key := path.Join(kv.rootPath, prefix)
	_, err := kv.removeEtcdMeta(context.TODO(), key, clientv3.WithPrefix())
	CheckElapseAndWarn(start, "Slow etcd operation remove with prefix", zap.String("prefix", prefix))
	return err
}

// Remove removes the key.
func (kv *EtcdKV) Remove(key string) error {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	_, err := kv.removeEtcdMeta(context.TODO(), key)
	CheckElapseAndWarn(start, "Slow etcd operation remove", zap.String("key", key))
	return err
}

// MultiRemove removes the keys in a transaction.
func (kv *EtcdKV) MultiRemove(keys []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(keys))
	for _, key := range keys {
		ops = append(ops, clientv3.OpDelete(path.Join(kv.rootPath, key)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiRemove error", zap.Strings("keys", keys), zap.Int("len", len(keys)), zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi remove", zap.Strings("keys", keys))
	return err
}

// MultiSaveAndRemove saves the key-value pairs and removes the keys in a transaction.
func (kv *EtcdKV) MultiSaveAndRemove(saves map[string]string, removals []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(saves)+len(removals))
	var keys []string
	for key, value := range saves {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), value))
	}

	for _, keyDelete := range removals {
		ops = append(ops, clientv3.OpDelete(path.Join(kv.rootPath, keyDelete)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSaveAndRemove error",
			zap.Any("saves", saves),
			zap.Strings("removes", removals),
			zap.Int("saveLength", len(saves)),
			zap.Int("removeLength", len(removals)),
			zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save and remove", zap.Strings("keys", keys))
	return err
}

// MultiSaveBytesAndRemove saves the key-value pairs and removes the keys in a transaction.
func (kv *EtcdKV) MultiSaveBytesAndRemove(saves map[string][]byte, removals []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(saves)+len(removals))
	var keys []string
	for key, value := range saves {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), string(value)))
	}

	for _, keyDelete := range removals {
		ops = append(ops, clientv3.OpDelete(path.Join(kv.rootPath, keyDelete)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSaveBytesAndRemove error",
			zap.Any("saves", saves),
			zap.Strings("removes", removals),
			zap.Int("saveLength", len(saves)),
			zap.Int("removeLength", len(removals)),
			zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save and remove", zap.Strings("keys", keys))
	return err
}

// Watch starts watching a key, returns a watch channel.
// Watch related can not set timeout
func (kv *EtcdKV) Watch(key string) clientv3.WatchChan {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	rch := kv.client.Watch(context.Background(), key, clientv3.WithCreatedNotify())
	CheckElapseAndWarn(start, "Slow etcd operation watch", zap.String("key", key))
	return rch
}

// WatchWithPrefix starts watching a key with prefix, returns a watch channel.
func (kv *EtcdKV) WatchWithPrefix(key string) clientv3.WatchChan {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	rch := kv.client.Watch(context.Background(), key, clientv3.WithPrefix(), clientv3.WithCreatedNotify())
	CheckElapseAndWarn(start, "Slow etcd operation watch with prefix", zap.String("key", key))
	return rch
}

// WatchWithRevision starts watching a key with revision, returns a watch channel.
func (kv *EtcdKV) WatchWithRevision(key string, revision int64) clientv3.WatchChan {
	start := time.Now()
	key = path.Join(kv.rootPath, key)
	rch := kv.client.Watch(context.Background(), key, clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(revision))
	CheckElapseAndWarn(start, "Slow etcd operation watch with revision", zap.String("key", key))
	return rch
}

// MultiRemoveWithPrefix removes the keys with given prefix.
func (kv *EtcdKV) MultiRemoveWithPrefix(keys []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(keys))
	for _, k := range keys {
		op := clientv3.OpDelete(path.Join(kv.rootPath, k), clientv3.WithPrefix())
		ops = append(ops, op)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiRemoveWithPrefix error", zap.Strings("keys", keys), zap.Int("len", len(keys)), zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi remove with prefix", zap.Strings("keys", keys))
	return err
}

// MultiSaveAndRemoveWithPrefix saves kv in @saves and removes the keys with given prefix in @removals.
func (kv *EtcdKV) MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(saves))
	var keys []string
	for key, value := range saves {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), value))
	}

	for _, keyDelete := range removals {
		ops = append(ops, clientv3.OpDelete(path.Join(kv.rootPath, keyDelete), clientv3.WithPrefix()))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSaveAndRemoveWithPrefix error",
			zap.Any("saves", saves),
			zap.Strings("removes", removals),
			zap.Int("saveLength", len(saves)),
			zap.Int("removeLength", len(removals)),
			zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save and move with prefix", zap.Strings("keys", keys))
	return err
}

// MultiSaveBytesAndRemoveWithPrefix saves kv in @saves and removes the keys with given prefix in @removals.
func (kv *EtcdKV) MultiSaveBytesAndRemoveWithPrefix(saves map[string][]byte, removals []string) error {
	start := time.Now()
	ops := make([]clientv3.Op, 0, len(saves))
	var keys []string
	for key, value := range saves {
		keys = append(keys, key)
		ops = append(ops, clientv3.OpPut(path.Join(kv.rootPath, key), string(value)))
	}

	for _, keyDelete := range removals {
		ops = append(ops, clientv3.OpDelete(path.Join(kv.rootPath, keyDelete), clientv3.WithPrefix()))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	_, err := kv.executeTxn(kv.getTxnWithCmp(ctx), ops...)
	if err != nil {
		log.Warn("Etcd MultiSaveBytesAndRemoveWithPrefix error",
			zap.Any("saves", saves),
			zap.Strings("removes", removals),
			zap.Int("saveLength", len(saves)),
			zap.Int("removeLength", len(removals)),
			zap.Error(err))
	}
	CheckElapseAndWarn(start, "Slow etcd operation multi save and move with prefix", zap.Strings("keys", keys))
	return err
}

// Grant creates a new lease implemented in etcd grant interface.
func (kv *EtcdKV) Grant(ttl int64) (id clientv3.LeaseID, err error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()

	resp, err := kv.client.Grant(ctx, ttl)
	CheckElapseAndWarn(start, "Slow etcd operation grant")
	return resp.ID, err
}

// KeepAlive keeps the lease alive forever with leaseID.
// Implemented in etcd interface.
func (kv *EtcdKV) KeepAlive(id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	start := time.Now()
	ch, err := kv.client.KeepAlive(context.Background(), id)
	if err != nil {
		return nil, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation keepAlive")
	return ch, nil
}

// CompareValueAndSwap compares the existing value with compare, and if they are
// equal, the target is stored in etcd.
func (kv *EtcdKV) CompareValueAndSwap(key, value, target string, opts ...clientv3.OpOption) (bool, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx,
		clientv3.Compare(clientv3.Value(path.Join(kv.rootPath, key)), "=", value)),
		clientv3.OpPut(path.Join(kv.rootPath, key), target, opts...))
	if err != nil {
		return false, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation compare value and swap", zap.String("key", key))
	return resp.Succeeded, nil
}

// CompareValueAndSwapBytes compares the existing value with compare, and if they are
// equal, the target is stored in etcd.
func (kv *EtcdKV) CompareValueAndSwapBytes(key string, value, target []byte, opts ...clientv3.OpOption) (bool, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx,
		clientv3.Compare(clientv3.Value(path.Join(kv.rootPath, key)), "=", string(value))),
		clientv3.OpPut(path.Join(kv.rootPath, key), string(target), opts...))
	if err != nil {
		return false, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation compare value and swap", zap.String("key", key))
	return resp.Succeeded, nil
}

// CompareVersionAndSwap compares the existing key-value's version with version, and if
// they are equal, the target is stored in etcd.
func (kv *EtcdKV) CompareVersionAndSwap(key string, source int64, target string, opts ...clientv3.OpOption) (bool, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx,
		clientv3.Compare(clientv3.Version(path.Join(kv.rootPath, key)), "=", source)),
		clientv3.OpPut(path.Join(kv.rootPath, key), target, opts...))
	if err != nil {
		return false, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation compare version and swap", zap.String("key", key))
	return resp.Succeeded, nil
}

// CompareVersionAndSwapBytes compares the existing key-value's version with version, and if
// they are equal, the target is stored in etcd.
func (kv *EtcdKV) CompareVersionAndSwapBytes(key string, source int64, target []byte, opts ...clientv3.OpOption) (bool, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), RequestTimeout)
	defer cancel()
	resp, err := kv.executeTxn(kv.getTxnWithCmp(ctx,
		clientv3.Compare(clientv3.Version(path.Join(kv.rootPath, key)), "=", source)),
		clientv3.OpPut(path.Join(kv.rootPath, key), string(target), opts...))
	if err != nil {
		return false, err
	}
	CheckElapseAndWarn(start, "Slow etcd operation compare version and swap", zap.String("key", key))
	return resp.Succeeded, nil
}

// CheckElapseAndWarn checks the elapsed time and warns if it is too long.
func CheckElapseAndWarn(start time.Time, message string, fields ...zap.Field) bool {
	elapsed := time.Since(start)
	if elapsed.Milliseconds() > 2000 {
		log.Warn(message, append([]zap.Field{zap.String("time spent", elapsed.String())}, fields...)...)
		return true
	}
	return false
}

func CheckValueSizeAndWarn(key string, value interface{}) bool {
	size := binary.Size(value)
	if size > 102400 {
		log.Warn("value size large than 100kb", zap.String("key", key), zap.Int("value_size(kb)", size/1024))
		return true
	}
	return false
}

func CheckTxnBytesValueSizeAndWarn(kvs map[string][]byte) bool {
	var hasWarn bool
	for key, value := range kvs {
		if CheckValueSizeAndWarn(key, value) {
			hasWarn = true
		}
	}
	return hasWarn
}

func CheckTxnStringValueSizeAndWarn(kvs map[string]string) bool {
	newKvs := make(map[string][]byte, len(kvs))
	for key, value := range kvs {
		newKvs[key] = []byte(value)
	}

	return CheckTxnBytesValueSizeAndWarn(newKvs)
}

func (kv *EtcdKV) getEtcdMeta(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx1, cancel := context.WithTimeout(ctx, RequestTimeout)
	defer cancel()

	start := timerecord.NewTimeRecorder("getEtcdMeta")
	resp, err := kv.client.Get(ctx1, key, opts...)
	elapsed := start.ElapseSpan()
	metrics.MetaOpCounter.WithLabelValues(metrics.MetaGetLabel, metrics.TotalLabel).Inc()

	// cal meta kv size
	if err == nil && resp != nil {
		totalSize := 0
		for _, v := range resp.Kvs {
			totalSize += binary.Size(v)
		}
		metrics.MetaKvSize.WithLabelValues(metrics.MetaGetLabel).Observe(float64(totalSize))
		metrics.MetaRequestLatency.WithLabelValues(metrics.MetaGetLabel).Observe(float64(elapsed.Milliseconds()))
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaGetLabel, metrics.SuccessLabel).Inc()
	} else {
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaGetLabel, metrics.FailLabel).Inc()
	}
	return resp, err
}

func (kv *EtcdKV) putEtcdMeta(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx1, cancel := context.WithTimeout(ctx, RequestTimeout)
	defer cancel()

	start := timerecord.NewTimeRecorder("putEtcdMeta")
	resp, err := kv.client.Put(ctx1, key, val, opts...)
	elapsed := start.ElapseSpan()
	metrics.MetaOpCounter.WithLabelValues(metrics.MetaPutLabel, metrics.TotalLabel).Inc()
	if err == nil {
		metrics.MetaKvSize.WithLabelValues(metrics.MetaPutLabel).Observe(float64(len(val)))
		metrics.MetaRequestLatency.WithLabelValues(metrics.MetaPutLabel).Observe(float64(elapsed.Milliseconds()))
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaPutLabel, metrics.SuccessLabel).Inc()
	} else {
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaPutLabel, metrics.FailLabel).Inc()
	}

	return resp, err
}

func (kv *EtcdKV) removeEtcdMeta(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx1, cancel := context.WithTimeout(ctx, RequestTimeout)
	defer cancel()

	start := timerecord.NewTimeRecorder("removeEtcdMeta")
	resp, err := kv.client.Delete(ctx1, key, opts...)
	elapsed := start.ElapseSpan()
	metrics.MetaOpCounter.WithLabelValues(metrics.MetaRemoveLabel, metrics.TotalLabel).Inc()

	if err == nil {
		metrics.MetaRequestLatency.WithLabelValues(metrics.MetaRemoveLabel).Observe(float64(elapsed.Milliseconds()))
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaRemoveLabel, metrics.SuccessLabel).Inc()
	} else {
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaRemoveLabel, metrics.FailLabel).Inc()
	}

	return resp, err
}

func (kv *EtcdKV) getTxnWithCmp(ctx context.Context, cmp ...clientv3.Cmp) clientv3.Txn {
	return kv.client.Txn(ctx).If(cmp...)
}

func (kv *EtcdKV) executeTxn(txn clientv3.Txn, ops ...clientv3.Op) (*clientv3.TxnResponse, error) {
	start := timerecord.NewTimeRecorder("executeTxn")

	resp, err := txn.Then(ops...).Commit()
	elapsed := start.ElapseSpan()
	metrics.MetaOpCounter.WithLabelValues(metrics.MetaTxnLabel, metrics.TotalLabel).Inc()

	if err == nil && resp.Succeeded {
		// cal put meta kv size
		totalPutSize := 0
		for _, op := range ops {
			if op.IsPut() {
				totalPutSize += binary.Size(op.ValueBytes())
			}
		}
		metrics.MetaKvSize.WithLabelValues(metrics.MetaPutLabel).Observe(float64(totalPutSize))

		// cal get meta kv size
		totalGetSize := 0
		for _, rp := range resp.Responses {
			if rp.GetResponseRange() != nil {
				for _, v := range rp.GetResponseRange().Kvs {
					totalGetSize += binary.Size(v)
				}
			}
		}
		metrics.MetaKvSize.WithLabelValues(metrics.MetaGetLabel).Observe(float64(totalGetSize))
		metrics.MetaRequestLatency.WithLabelValues(metrics.MetaTxnLabel).Observe(float64(elapsed.Milliseconds()))
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaTxnLabel, metrics.SuccessLabel).Inc()
	} else {
		metrics.MetaOpCounter.WithLabelValues(metrics.MetaTxnLabel, metrics.FailLabel).Inc()
	}

	return resp, err
}
