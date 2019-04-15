/**
 * Autogenerated by Thrift Compiler (0.12.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef ThriftTest_TYPES_H
#define ThriftTest_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>


namespace thrift { namespace test {

struct Numberz {
  enum type {
    ONE = 1,
    TWO = 2,
    THREE = 3,
    FIVE = 5,
    SIX = 6,
    EIGHT = 8
  };
};

extern const std::map<int, const char*> _Numberz_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const Numberz::type& val);

typedef int64_t UserId;

typedef std::map<std::string, class Bonk>  MapType;

class Bonk;

class Bools;

class Xtruct;

class Xtruct2;

class Xtruct3;

class Insanity;

class CrazyNesting;

class SomeUnion;

class Xception;

class Xception2;

class EmptyStruct;

class OneField;

class VersioningTestV1;

class VersioningTestV2;

class ListTypeVersioningV1;

class ListTypeVersioningV2;

class GuessProtocolStruct;

class LargeDeltas;

class NestedListsI32x2;

class NestedListsI32x3;

class NestedMixedx2;

class ListBonks;

class NestedListsBonk;

class BoolTest;

class StructA;

class StructB;

typedef struct _Bonk__isset {
  _Bonk__isset() : message(false), type(false) {}
  bool message :1;
  bool type :1;
} _Bonk__isset;

class Bonk {
 public:

  Bonk(const Bonk&);
  Bonk& operator=(const Bonk&);
  Bonk() : message(), type(0) {
  }

  virtual ~Bonk() throw();
  std::string message;
  int32_t type;

  _Bonk__isset __isset;

  void __set_message(const std::string& val);

  void __set_type(const int32_t val);

  bool operator == (const Bonk & rhs) const
  {
    if (!(message == rhs.message))
      return false;
    if (!(type == rhs.type))
      return false;
    return true;
  }
  bool operator != (const Bonk &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Bonk & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Bonk &a, Bonk &b);

std::ostream& operator<<(std::ostream& out, const Bonk& obj);

typedef struct _Bools__isset {
  _Bools__isset() : im_true(false), im_false(false) {}
  bool im_true :1;
  bool im_false :1;
} _Bools__isset;

class Bools {
 public:

  Bools(const Bools&);
  Bools& operator=(const Bools&);
  Bools() : im_true(0), im_false(0) {
  }

  virtual ~Bools() throw();
  bool im_true;
  bool im_false;

  _Bools__isset __isset;

  void __set_im_true(const bool val);

  void __set_im_false(const bool val);

  bool operator == (const Bools & rhs) const
  {
    if (!(im_true == rhs.im_true))
      return false;
    if (!(im_false == rhs.im_false))
      return false;
    return true;
  }
  bool operator != (const Bools &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Bools & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Bools &a, Bools &b);

std::ostream& operator<<(std::ostream& out, const Bools& obj);

typedef struct _Xtruct__isset {
  _Xtruct__isset() : string_thing(false), byte_thing(false), i32_thing(false), i64_thing(false) {}
  bool string_thing :1;
  bool byte_thing :1;
  bool i32_thing :1;
  bool i64_thing :1;
} _Xtruct__isset;

class Xtruct {
 public:

  Xtruct(const Xtruct&);
  Xtruct& operator=(const Xtruct&);
  Xtruct() : string_thing(), byte_thing(0), i32_thing(0), i64_thing(0) {
  }

  virtual ~Xtruct() throw();
  std::string string_thing;
  int8_t byte_thing;
  int32_t i32_thing;
  int64_t i64_thing;

  _Xtruct__isset __isset;

  void __set_string_thing(const std::string& val);

  void __set_byte_thing(const int8_t val);

  void __set_i32_thing(const int32_t val);

  void __set_i64_thing(const int64_t val);

  bool operator == (const Xtruct & rhs) const
  {
    if (!(string_thing == rhs.string_thing))
      return false;
    if (!(byte_thing == rhs.byte_thing))
      return false;
    if (!(i32_thing == rhs.i32_thing))
      return false;
    if (!(i64_thing == rhs.i64_thing))
      return false;
    return true;
  }
  bool operator != (const Xtruct &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Xtruct & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Xtruct &a, Xtruct &b);

std::ostream& operator<<(std::ostream& out, const Xtruct& obj);

typedef struct _Xtruct2__isset {
  _Xtruct2__isset() : byte_thing(false), struct_thing(false), i32_thing(false) {}
  bool byte_thing :1;
  bool struct_thing :1;
  bool i32_thing :1;
} _Xtruct2__isset;

class Xtruct2 {
 public:

  Xtruct2(const Xtruct2&);
  Xtruct2& operator=(const Xtruct2&);
  Xtruct2() : byte_thing(0), i32_thing(0) {
  }

  virtual ~Xtruct2() throw();
  int8_t byte_thing;
  Xtruct struct_thing;
  int32_t i32_thing;

  _Xtruct2__isset __isset;

  void __set_byte_thing(const int8_t val);

  void __set_struct_thing(const Xtruct& val);

  void __set_i32_thing(const int32_t val);

  bool operator == (const Xtruct2 & rhs) const
  {
    if (!(byte_thing == rhs.byte_thing))
      return false;
    if (!(struct_thing == rhs.struct_thing))
      return false;
    if (!(i32_thing == rhs.i32_thing))
      return false;
    return true;
  }
  bool operator != (const Xtruct2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Xtruct2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Xtruct2 &a, Xtruct2 &b);

std::ostream& operator<<(std::ostream& out, const Xtruct2& obj);

typedef struct _Xtruct3__isset {
  _Xtruct3__isset() : string_thing(false), changed(false), i32_thing(false), i64_thing(false) {}
  bool string_thing :1;
  bool changed :1;
  bool i32_thing :1;
  bool i64_thing :1;
} _Xtruct3__isset;

class Xtruct3 {
 public:

  Xtruct3(const Xtruct3&);
  Xtruct3& operator=(const Xtruct3&);
  Xtruct3() : string_thing(), changed(0), i32_thing(0), i64_thing(0) {
  }

  virtual ~Xtruct3() throw();
  std::string string_thing;
  int32_t changed;
  int32_t i32_thing;
  int64_t i64_thing;

  _Xtruct3__isset __isset;

  void __set_string_thing(const std::string& val);

  void __set_changed(const int32_t val);

  void __set_i32_thing(const int32_t val);

  void __set_i64_thing(const int64_t val);

  bool operator == (const Xtruct3 & rhs) const
  {
    if (!(string_thing == rhs.string_thing))
      return false;
    if (!(changed == rhs.changed))
      return false;
    if (!(i32_thing == rhs.i32_thing))
      return false;
    if (!(i64_thing == rhs.i64_thing))
      return false;
    return true;
  }
  bool operator != (const Xtruct3 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Xtruct3 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Xtruct3 &a, Xtruct3 &b);

std::ostream& operator<<(std::ostream& out, const Xtruct3& obj);

typedef struct _Insanity__isset {
  _Insanity__isset() : userMap(false), xtructs(false) {}
  bool userMap :1;
  bool xtructs :1;
} _Insanity__isset;

class Insanity {
 public:

  Insanity(const Insanity&);
  Insanity& operator=(const Insanity&);
  Insanity() {
  }

  virtual ~Insanity() throw();
  std::map<Numberz::type, UserId>  userMap;
  std::vector<Xtruct>  xtructs;

  _Insanity__isset __isset;

  void __set_userMap(const std::map<Numberz::type, UserId> & val);

  void __set_xtructs(const std::vector<Xtruct> & val);

  bool operator == (const Insanity & rhs) const
  {
    if (!(userMap == rhs.userMap))
      return false;
    if (!(xtructs == rhs.xtructs))
      return false;
    return true;
  }
  bool operator != (const Insanity &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Insanity & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Insanity &a, Insanity &b);

std::ostream& operator<<(std::ostream& out, const Insanity& obj);

typedef struct _CrazyNesting__isset {
  _CrazyNesting__isset() : string_field(false), set_field(false), binary_field(false) {}
  bool string_field :1;
  bool set_field :1;
  bool binary_field :1;
} _CrazyNesting__isset;

class CrazyNesting {
 public:

  CrazyNesting(const CrazyNesting&);
  CrazyNesting& operator=(const CrazyNesting&);
  CrazyNesting() : string_field(), binary_field() {
  }

  virtual ~CrazyNesting() throw();
  std::string string_field;
  std::set<Insanity>  set_field;
  std::vector<std::map<std::set<int32_t> , std::map<int32_t, std::set<std::vector<std::map<Insanity, std::string> > > > > >  list_field;
  std::string binary_field;

  _CrazyNesting__isset __isset;

  void __set_string_field(const std::string& val);

  void __set_set_field(const std::set<Insanity> & val);

  void __set_list_field(const std::vector<std::map<std::set<int32_t> , std::map<int32_t, std::set<std::vector<std::map<Insanity, std::string> > > > > > & val);

  void __set_binary_field(const std::string& val);

  bool operator == (const CrazyNesting & rhs) const
  {
    if (!(string_field == rhs.string_field))
      return false;
    if (__isset.set_field != rhs.__isset.set_field)
      return false;
    else if (__isset.set_field && !(set_field == rhs.set_field))
      return false;
    if (!(list_field == rhs.list_field))
      return false;
    if (!(binary_field == rhs.binary_field))
      return false;
    return true;
  }
  bool operator != (const CrazyNesting &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const CrazyNesting & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(CrazyNesting &a, CrazyNesting &b);

std::ostream& operator<<(std::ostream& out, const CrazyNesting& obj);

typedef struct _SomeUnion__isset {
  _SomeUnion__isset() : map_thing(false), string_thing(false), i32_thing(false), xtruct_thing(false), insanity_thing(false) {}
  bool map_thing :1;
  bool string_thing :1;
  bool i32_thing :1;
  bool xtruct_thing :1;
  bool insanity_thing :1;
} _SomeUnion__isset;

class SomeUnion {
 public:

  SomeUnion(const SomeUnion&);
  SomeUnion& operator=(const SomeUnion&);
  SomeUnion() : string_thing(), i32_thing(0) {
  }

  virtual ~SomeUnion() throw();
  std::map<Numberz::type, UserId>  map_thing;
  std::string string_thing;
  int32_t i32_thing;
  Xtruct3 xtruct_thing;
  Insanity insanity_thing;

  _SomeUnion__isset __isset;

  void __set_map_thing(const std::map<Numberz::type, UserId> & val);

  void __set_string_thing(const std::string& val);

  void __set_i32_thing(const int32_t val);

  void __set_xtruct_thing(const Xtruct3& val);

  void __set_insanity_thing(const Insanity& val);

  bool operator == (const SomeUnion & rhs) const
  {
    if (__isset.map_thing != rhs.__isset.map_thing)
      return false;
    else if (__isset.map_thing && !(map_thing == rhs.map_thing))
      return false;
    if (__isset.string_thing != rhs.__isset.string_thing)
      return false;
    else if (__isset.string_thing && !(string_thing == rhs.string_thing))
      return false;
    if (__isset.i32_thing != rhs.__isset.i32_thing)
      return false;
    else if (__isset.i32_thing && !(i32_thing == rhs.i32_thing))
      return false;
    if (__isset.xtruct_thing != rhs.__isset.xtruct_thing)
      return false;
    else if (__isset.xtruct_thing && !(xtruct_thing == rhs.xtruct_thing))
      return false;
    if (__isset.insanity_thing != rhs.__isset.insanity_thing)
      return false;
    else if (__isset.insanity_thing && !(insanity_thing == rhs.insanity_thing))
      return false;
    return true;
  }
  bool operator != (const SomeUnion &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const SomeUnion & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(SomeUnion &a, SomeUnion &b);

std::ostream& operator<<(std::ostream& out, const SomeUnion& obj);

typedef struct _Xception__isset {
  _Xception__isset() : errorCode(false), message(false) {}
  bool errorCode :1;
  bool message :1;
} _Xception__isset;

class Xception : public ::apache::thrift::TException {
 public:

  Xception(const Xception&);
  Xception& operator=(const Xception&);
  Xception() : errorCode(0), message() {
  }

  virtual ~Xception() throw();
  int32_t errorCode;
  std::string message;

  _Xception__isset __isset;

  void __set_errorCode(const int32_t val);

  void __set_message(const std::string& val);

  bool operator == (const Xception & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (!(message == rhs.message))
      return false;
    return true;
  }
  bool operator != (const Xception &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Xception & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
  mutable std::string thriftTExceptionMessageHolder_;
  const char* what() const throw();
};

void swap(Xception &a, Xception &b);

std::ostream& operator<<(std::ostream& out, const Xception& obj);

typedef struct _Xception2__isset {
  _Xception2__isset() : errorCode(false), struct_thing(false) {}
  bool errorCode :1;
  bool struct_thing :1;
} _Xception2__isset;

class Xception2 : public ::apache::thrift::TException {
 public:

  Xception2(const Xception2&);
  Xception2& operator=(const Xception2&);
  Xception2() : errorCode(0) {
  }

  virtual ~Xception2() throw();
  int32_t errorCode;
  Xtruct struct_thing;

  _Xception2__isset __isset;

  void __set_errorCode(const int32_t val);

  void __set_struct_thing(const Xtruct& val);

  bool operator == (const Xception2 & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (!(struct_thing == rhs.struct_thing))
      return false;
    return true;
  }
  bool operator != (const Xception2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Xception2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
  mutable std::string thriftTExceptionMessageHolder_;
  const char* what() const throw();
};

void swap(Xception2 &a, Xception2 &b);

std::ostream& operator<<(std::ostream& out, const Xception2& obj);


class EmptyStruct {
 public:

  EmptyStruct(const EmptyStruct&);
  EmptyStruct& operator=(const EmptyStruct&);
  EmptyStruct() {
  }

  virtual ~EmptyStruct() throw();

  bool operator == (const EmptyStruct & /* rhs */) const
  {
    return true;
  }
  bool operator != (const EmptyStruct &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const EmptyStruct & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(EmptyStruct &a, EmptyStruct &b);

std::ostream& operator<<(std::ostream& out, const EmptyStruct& obj);

typedef struct _OneField__isset {
  _OneField__isset() : field(false) {}
  bool field :1;
} _OneField__isset;

class OneField {
 public:

  OneField(const OneField&);
  OneField& operator=(const OneField&);
  OneField() {
  }

  virtual ~OneField() throw();
  EmptyStruct field;

  _OneField__isset __isset;

  void __set_field(const EmptyStruct& val);

  bool operator == (const OneField & rhs) const
  {
    if (!(field == rhs.field))
      return false;
    return true;
  }
  bool operator != (const OneField &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const OneField & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(OneField &a, OneField &b);

std::ostream& operator<<(std::ostream& out, const OneField& obj);

typedef struct _VersioningTestV1__isset {
  _VersioningTestV1__isset() : begin_in_both(false), old_string(false), end_in_both(false) {}
  bool begin_in_both :1;
  bool old_string :1;
  bool end_in_both :1;
} _VersioningTestV1__isset;

class VersioningTestV1 {
 public:

  VersioningTestV1(const VersioningTestV1&);
  VersioningTestV1& operator=(const VersioningTestV1&);
  VersioningTestV1() : begin_in_both(0), old_string(), end_in_both(0) {
  }

  virtual ~VersioningTestV1() throw();
  int32_t begin_in_both;
  std::string old_string;
  int32_t end_in_both;

  _VersioningTestV1__isset __isset;

  void __set_begin_in_both(const int32_t val);

  void __set_old_string(const std::string& val);

  void __set_end_in_both(const int32_t val);

  bool operator == (const VersioningTestV1 & rhs) const
  {
    if (!(begin_in_both == rhs.begin_in_both))
      return false;
    if (!(old_string == rhs.old_string))
      return false;
    if (!(end_in_both == rhs.end_in_both))
      return false;
    return true;
  }
  bool operator != (const VersioningTestV1 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VersioningTestV1 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VersioningTestV1 &a, VersioningTestV1 &b);

std::ostream& operator<<(std::ostream& out, const VersioningTestV1& obj);

typedef struct _VersioningTestV2__isset {
  _VersioningTestV2__isset() : begin_in_both(false), newint(false), newbyte(false), newshort(false), newlong(false), newdouble(false), newstruct(false), newlist(false), newset(false), newmap(false), newstring(false), end_in_both(false) {}
  bool begin_in_both :1;
  bool newint :1;
  bool newbyte :1;
  bool newshort :1;
  bool newlong :1;
  bool newdouble :1;
  bool newstruct :1;
  bool newlist :1;
  bool newset :1;
  bool newmap :1;
  bool newstring :1;
  bool end_in_both :1;
} _VersioningTestV2__isset;

class VersioningTestV2 {
 public:

  VersioningTestV2(const VersioningTestV2&);
  VersioningTestV2& operator=(const VersioningTestV2&);
  VersioningTestV2() : begin_in_both(0), newint(0), newbyte(0), newshort(0), newlong(0), newdouble(0), newstring(), end_in_both(0) {
  }

  virtual ~VersioningTestV2() throw();
  int32_t begin_in_both;
  int32_t newint;
  int8_t newbyte;
  int16_t newshort;
  int64_t newlong;
  double newdouble;
  Bonk newstruct;
  std::vector<int32_t>  newlist;
  std::set<int32_t>  newset;
  std::map<int32_t, int32_t>  newmap;
  std::string newstring;
  int32_t end_in_both;

  _VersioningTestV2__isset __isset;

  void __set_begin_in_both(const int32_t val);

  void __set_newint(const int32_t val);

  void __set_newbyte(const int8_t val);

  void __set_newshort(const int16_t val);

  void __set_newlong(const int64_t val);

  void __set_newdouble(const double val);

  void __set_newstruct(const Bonk& val);

  void __set_newlist(const std::vector<int32_t> & val);

  void __set_newset(const std::set<int32_t> & val);

  void __set_newmap(const std::map<int32_t, int32_t> & val);

  void __set_newstring(const std::string& val);

  void __set_end_in_both(const int32_t val);

  bool operator == (const VersioningTestV2 & rhs) const
  {
    if (!(begin_in_both == rhs.begin_in_both))
      return false;
    if (!(newint == rhs.newint))
      return false;
    if (!(newbyte == rhs.newbyte))
      return false;
    if (!(newshort == rhs.newshort))
      return false;
    if (!(newlong == rhs.newlong))
      return false;
    if (!(newdouble == rhs.newdouble))
      return false;
    if (!(newstruct == rhs.newstruct))
      return false;
    if (!(newlist == rhs.newlist))
      return false;
    if (!(newset == rhs.newset))
      return false;
    if (!(newmap == rhs.newmap))
      return false;
    if (!(newstring == rhs.newstring))
      return false;
    if (!(end_in_both == rhs.end_in_both))
      return false;
    return true;
  }
  bool operator != (const VersioningTestV2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VersioningTestV2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VersioningTestV2 &a, VersioningTestV2 &b);

std::ostream& operator<<(std::ostream& out, const VersioningTestV2& obj);

typedef struct _ListTypeVersioningV1__isset {
  _ListTypeVersioningV1__isset() : myints(false), hello(false) {}
  bool myints :1;
  bool hello :1;
} _ListTypeVersioningV1__isset;

class ListTypeVersioningV1 {
 public:

  ListTypeVersioningV1(const ListTypeVersioningV1&);
  ListTypeVersioningV1& operator=(const ListTypeVersioningV1&);
  ListTypeVersioningV1() : hello() {
  }

  virtual ~ListTypeVersioningV1() throw();
  std::vector<int32_t>  myints;
  std::string hello;

  _ListTypeVersioningV1__isset __isset;

  void __set_myints(const std::vector<int32_t> & val);

  void __set_hello(const std::string& val);

  bool operator == (const ListTypeVersioningV1 & rhs) const
  {
    if (!(myints == rhs.myints))
      return false;
    if (!(hello == rhs.hello))
      return false;
    return true;
  }
  bool operator != (const ListTypeVersioningV1 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const ListTypeVersioningV1 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(ListTypeVersioningV1 &a, ListTypeVersioningV1 &b);

std::ostream& operator<<(std::ostream& out, const ListTypeVersioningV1& obj);

typedef struct _ListTypeVersioningV2__isset {
  _ListTypeVersioningV2__isset() : strings(false), hello(false) {}
  bool strings :1;
  bool hello :1;
} _ListTypeVersioningV2__isset;

class ListTypeVersioningV2 {
 public:

  ListTypeVersioningV2(const ListTypeVersioningV2&);
  ListTypeVersioningV2& operator=(const ListTypeVersioningV2&);
  ListTypeVersioningV2() : hello() {
  }

  virtual ~ListTypeVersioningV2() throw();
  std::vector<std::string>  strings;
  std::string hello;

  _ListTypeVersioningV2__isset __isset;

  void __set_strings(const std::vector<std::string> & val);

  void __set_hello(const std::string& val);

  bool operator == (const ListTypeVersioningV2 & rhs) const
  {
    if (!(strings == rhs.strings))
      return false;
    if (!(hello == rhs.hello))
      return false;
    return true;
  }
  bool operator != (const ListTypeVersioningV2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const ListTypeVersioningV2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(ListTypeVersioningV2 &a, ListTypeVersioningV2 &b);

std::ostream& operator<<(std::ostream& out, const ListTypeVersioningV2& obj);

typedef struct _GuessProtocolStruct__isset {
  _GuessProtocolStruct__isset() : map_field(false) {}
  bool map_field :1;
} _GuessProtocolStruct__isset;

class GuessProtocolStruct {
 public:

  GuessProtocolStruct(const GuessProtocolStruct&);
  GuessProtocolStruct& operator=(const GuessProtocolStruct&);
  GuessProtocolStruct() {
  }

  virtual ~GuessProtocolStruct() throw();
  std::map<std::string, std::string>  map_field;

  _GuessProtocolStruct__isset __isset;

  void __set_map_field(const std::map<std::string, std::string> & val);

  bool operator == (const GuessProtocolStruct & rhs) const
  {
    if (!(map_field == rhs.map_field))
      return false;
    return true;
  }
  bool operator != (const GuessProtocolStruct &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const GuessProtocolStruct & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(GuessProtocolStruct &a, GuessProtocolStruct &b);

std::ostream& operator<<(std::ostream& out, const GuessProtocolStruct& obj);

typedef struct _LargeDeltas__isset {
  _LargeDeltas__isset() : b1(false), b10(false), b100(false), check_true(false), b1000(false), check_false(false), vertwo2000(false), a_set2500(false), vertwo3000(false), big_numbers(false) {}
  bool b1 :1;
  bool b10 :1;
  bool b100 :1;
  bool check_true :1;
  bool b1000 :1;
  bool check_false :1;
  bool vertwo2000 :1;
  bool a_set2500 :1;
  bool vertwo3000 :1;
  bool big_numbers :1;
} _LargeDeltas__isset;

class LargeDeltas {
 public:

  LargeDeltas(const LargeDeltas&);
  LargeDeltas& operator=(const LargeDeltas&);
  LargeDeltas() : check_true(0), check_false(0) {
  }

  virtual ~LargeDeltas() throw();
  Bools b1;
  Bools b10;
  Bools b100;
  bool check_true;
  Bools b1000;
  bool check_false;
  VersioningTestV2 vertwo2000;
  std::set<std::string>  a_set2500;
  VersioningTestV2 vertwo3000;
  std::vector<int32_t>  big_numbers;

  _LargeDeltas__isset __isset;

  void __set_b1(const Bools& val);

  void __set_b10(const Bools& val);

  void __set_b100(const Bools& val);

  void __set_check_true(const bool val);

  void __set_b1000(const Bools& val);

  void __set_check_false(const bool val);

  void __set_vertwo2000(const VersioningTestV2& val);

  void __set_a_set2500(const std::set<std::string> & val);

  void __set_vertwo3000(const VersioningTestV2& val);

  void __set_big_numbers(const std::vector<int32_t> & val);

  bool operator == (const LargeDeltas & rhs) const
  {
    if (!(b1 == rhs.b1))
      return false;
    if (!(b10 == rhs.b10))
      return false;
    if (!(b100 == rhs.b100))
      return false;
    if (!(check_true == rhs.check_true))
      return false;
    if (!(b1000 == rhs.b1000))
      return false;
    if (!(check_false == rhs.check_false))
      return false;
    if (!(vertwo2000 == rhs.vertwo2000))
      return false;
    if (!(a_set2500 == rhs.a_set2500))
      return false;
    if (!(vertwo3000 == rhs.vertwo3000))
      return false;
    if (!(big_numbers == rhs.big_numbers))
      return false;
    return true;
  }
  bool operator != (const LargeDeltas &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const LargeDeltas & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(LargeDeltas &a, LargeDeltas &b);

std::ostream& operator<<(std::ostream& out, const LargeDeltas& obj);

typedef struct _NestedListsI32x2__isset {
  _NestedListsI32x2__isset() : integerlist(false) {}
  bool integerlist :1;
} _NestedListsI32x2__isset;

class NestedListsI32x2 {
 public:

  NestedListsI32x2(const NestedListsI32x2&);
  NestedListsI32x2& operator=(const NestedListsI32x2&);
  NestedListsI32x2() {
  }

  virtual ~NestedListsI32x2() throw();
  std::vector<std::vector<int32_t> >  integerlist;

  _NestedListsI32x2__isset __isset;

  void __set_integerlist(const std::vector<std::vector<int32_t> > & val);

  bool operator == (const NestedListsI32x2 & rhs) const
  {
    if (!(integerlist == rhs.integerlist))
      return false;
    return true;
  }
  bool operator != (const NestedListsI32x2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const NestedListsI32x2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(NestedListsI32x2 &a, NestedListsI32x2 &b);

std::ostream& operator<<(std::ostream& out, const NestedListsI32x2& obj);

typedef struct _NestedListsI32x3__isset {
  _NestedListsI32x3__isset() : integerlist(false) {}
  bool integerlist :1;
} _NestedListsI32x3__isset;

class NestedListsI32x3 {
 public:

  NestedListsI32x3(const NestedListsI32x3&);
  NestedListsI32x3& operator=(const NestedListsI32x3&);
  NestedListsI32x3() {
  }

  virtual ~NestedListsI32x3() throw();
  std::vector<std::vector<std::vector<int32_t> > >  integerlist;

  _NestedListsI32x3__isset __isset;

  void __set_integerlist(const std::vector<std::vector<std::vector<int32_t> > > & val);

  bool operator == (const NestedListsI32x3 & rhs) const
  {
    if (!(integerlist == rhs.integerlist))
      return false;
    return true;
  }
  bool operator != (const NestedListsI32x3 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const NestedListsI32x3 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(NestedListsI32x3 &a, NestedListsI32x3 &b);

std::ostream& operator<<(std::ostream& out, const NestedListsI32x3& obj);

typedef struct _NestedMixedx2__isset {
  _NestedMixedx2__isset() : int_set_list(false), map_int_strset(false), map_int_strset_list(false) {}
  bool int_set_list :1;
  bool map_int_strset :1;
  bool map_int_strset_list :1;
} _NestedMixedx2__isset;

class NestedMixedx2 {
 public:

  NestedMixedx2(const NestedMixedx2&);
  NestedMixedx2& operator=(const NestedMixedx2&);
  NestedMixedx2() {
  }

  virtual ~NestedMixedx2() throw();
  std::vector<std::set<int32_t> >  int_set_list;
  std::map<int32_t, std::set<std::string> >  map_int_strset;
  std::vector<std::map<int32_t, std::set<std::string> > >  map_int_strset_list;

  _NestedMixedx2__isset __isset;

  void __set_int_set_list(const std::vector<std::set<int32_t> > & val);

  void __set_map_int_strset(const std::map<int32_t, std::set<std::string> > & val);

  void __set_map_int_strset_list(const std::vector<std::map<int32_t, std::set<std::string> > > & val);

  bool operator == (const NestedMixedx2 & rhs) const
  {
    if (!(int_set_list == rhs.int_set_list))
      return false;
    if (!(map_int_strset == rhs.map_int_strset))
      return false;
    if (!(map_int_strset_list == rhs.map_int_strset_list))
      return false;
    return true;
  }
  bool operator != (const NestedMixedx2 &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const NestedMixedx2 & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(NestedMixedx2 &a, NestedMixedx2 &b);

std::ostream& operator<<(std::ostream& out, const NestedMixedx2& obj);

typedef struct _ListBonks__isset {
  _ListBonks__isset() : bonk(false) {}
  bool bonk :1;
} _ListBonks__isset;

class ListBonks {
 public:

  ListBonks(const ListBonks&);
  ListBonks& operator=(const ListBonks&);
  ListBonks() {
  }

  virtual ~ListBonks() throw();
  std::vector<Bonk>  bonk;

  _ListBonks__isset __isset;

  void __set_bonk(const std::vector<Bonk> & val);

  bool operator == (const ListBonks & rhs) const
  {
    if (!(bonk == rhs.bonk))
      return false;
    return true;
  }
  bool operator != (const ListBonks &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const ListBonks & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(ListBonks &a, ListBonks &b);

std::ostream& operator<<(std::ostream& out, const ListBonks& obj);

typedef struct _NestedListsBonk__isset {
  _NestedListsBonk__isset() : bonk(false) {}
  bool bonk :1;
} _NestedListsBonk__isset;

class NestedListsBonk {
 public:

  NestedListsBonk(const NestedListsBonk&);
  NestedListsBonk& operator=(const NestedListsBonk&);
  NestedListsBonk() {
  }

  virtual ~NestedListsBonk() throw();
  std::vector<std::vector<std::vector<Bonk> > >  bonk;

  _NestedListsBonk__isset __isset;

  void __set_bonk(const std::vector<std::vector<std::vector<Bonk> > > & val);

  bool operator == (const NestedListsBonk & rhs) const
  {
    if (!(bonk == rhs.bonk))
      return false;
    return true;
  }
  bool operator != (const NestedListsBonk &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const NestedListsBonk & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(NestedListsBonk &a, NestedListsBonk &b);

std::ostream& operator<<(std::ostream& out, const NestedListsBonk& obj);

typedef struct _BoolTest__isset {
  _BoolTest__isset() : b(true), s(true) {}
  bool b :1;
  bool s :1;
} _BoolTest__isset;

class BoolTest {
 public:

  BoolTest(const BoolTest&);
  BoolTest& operator=(const BoolTest&);
  BoolTest() : b(true), s("true") {
  }

  virtual ~BoolTest() throw();
  bool b;
  std::string s;

  _BoolTest__isset __isset;

  void __set_b(const bool val);

  void __set_s(const std::string& val);

  bool operator == (const BoolTest & rhs) const
  {
    if (__isset.b != rhs.__isset.b)
      return false;
    else if (__isset.b && !(b == rhs.b))
      return false;
    if (__isset.s != rhs.__isset.s)
      return false;
    else if (__isset.s && !(s == rhs.s))
      return false;
    return true;
  }
  bool operator != (const BoolTest &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const BoolTest & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(BoolTest &a, BoolTest &b);

std::ostream& operator<<(std::ostream& out, const BoolTest& obj);


class StructA {
 public:

  StructA(const StructA&);
  StructA& operator=(const StructA&);
  StructA() : s() {
  }

  virtual ~StructA() throw();
  std::string s;

  void __set_s(const std::string& val);

  bool operator == (const StructA & rhs) const
  {
    if (!(s == rhs.s))
      return false;
    return true;
  }
  bool operator != (const StructA &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StructA & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(StructA &a, StructA &b);

std::ostream& operator<<(std::ostream& out, const StructA& obj);

typedef struct _StructB__isset {
  _StructB__isset() : aa(false) {}
  bool aa :1;
} _StructB__isset;

class StructB {
 public:

  StructB(const StructB&);
  StructB& operator=(const StructB&);
  StructB() {
  }

  virtual ~StructB() throw();
  StructA aa;
  StructA ab;

  _StructB__isset __isset;

  void __set_aa(const StructA& val);

  void __set_ab(const StructA& val);

  bool operator == (const StructB & rhs) const
  {
    if (__isset.aa != rhs.__isset.aa)
      return false;
    else if (__isset.aa && !(aa == rhs.aa))
      return false;
    if (!(ab == rhs.ab))
      return false;
    return true;
  }
  bool operator != (const StructB &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StructB & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(StructB &a, StructB &b);

std::ostream& operator<<(std::ostream& out, const StructB& obj);

}} // namespace

#include "ThriftTest_types.tcc"

#endif
