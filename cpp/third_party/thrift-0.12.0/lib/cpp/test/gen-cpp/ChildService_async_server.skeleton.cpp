// This autogenerated skeleton file illustrates one way to adapt a synchronous
// interface into an asynchronous interface. You should copy it to another
// filename to avoid overwriting it and rewrite as asynchronous any functions
// that would otherwise introduce unwanted latency.

#include "ChildService.h"
#include <thrift/protocol/TBinaryProtocol.h>

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::async;

using namespace  ::apache::thrift::test;

class ChildServiceAsyncHandler : public ChildServiceCobSvIf {
 public:
  ChildServiceAsyncHandler() {
    syncHandler_ = std::auto_ptr<ChildServiceHandler>(new ChildServiceHandler);
    // Your initialization goes here
  }
  virtual ~ChildServiceAsyncHandler();

  void setValue(::apache::thrift::stdcxx::function<void(int32_t const& _return)> cob, const int32_t value) {
    int32_t _return = 0;
    _return = syncHandler_->setValue(value);
    return cob(_return);
  }

  void getValue(::apache::thrift::stdcxx::function<void(int32_t const& _return)> cob) {
    int32_t _return = 0;
    _return = syncHandler_->getValue();
    return cob(_return);
  }

 protected:
  std::auto_ptr<ChildServiceHandler> syncHandler_;
};

