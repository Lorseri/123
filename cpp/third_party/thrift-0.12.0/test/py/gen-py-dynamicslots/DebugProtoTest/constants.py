#
# Autogenerated by Thrift Compiler (0.12.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py:dynamic,slots
#

from thrift.Thrift import TType, TMessageType, TFrozenDict, TException, TApplicationException
from thrift.protocol.TProtocol import TProtocolException
from thrift.TRecursive import fix_spec

import sys
from .ttypes import *
COMPACT_TEST = CompactProtoTestStruct(**{
    "a_byte": 127,
    "a_double": 5.6788999999999996,
    "a_i16": 32000,
    "a_i32": 1000000000,
    "a_i64": 1099511627775,
    "a_string": "my string",
    "boolean_byte_map": {
        False: 0,
        True: 1,
    },
    "boolean_list": [
        True,
        True,
        True,
        False,
        False,
        False,
    ],
    "boolean_set": set((
        True,
        False,
    )),
    "byte_boolean_map": {
        1: True,
        2: False,
    },
    "byte_byte_map": {
        1: 2,
    },
    "byte_double_map": {
        1: 0.1000000000000000,
        2: -0.1000000000000000,
        3: 1000000.0999999999767169,
    },
    "byte_i16_map": {
        1: 1,
        2: -1,
        3: 32767,
    },
    "byte_i32_map": {
        1: 1,
        2: -1,
        3: 2147483647,
    },
    "byte_i64_map": {
        1: 1,
        2: -1,
        3: 9223372036854775807,
    },
    "byte_list": [
        -127,
        -1,
        0,
        1,
        127,
    ],
    "byte_list_map": {
        0: [
        ],
        1: [
            1,
        ],
        2: [
            1,
            2,
        ],
    },
    "byte_map_map": {
        0: {
        },
        1: {
            1: 1,
        },
        2: {
            1: 1,
            2: 2,
        },
    },
    "byte_set": set((
        -127,
        -1,
        0,
        1,
        127,
    )),
    "byte_set_map": {
        0: set((
        )),
        1: set((
            1,
        )),
        2: set((
            1,
            2,
        )),
    },
    "byte_string_map": {
        1: "",
        2: "blah",
        3: "loooooooooooooong string",
    },
    "double_byte_map": {
        -1.1000000000000001: 1,
        1.1000000000000001: 1,
    },
    "double_list": [
        0.1000000000000000,
        0.2000000000000000,
        0.3000000000000000,
    ],
    "double_set": set((
        0.1000000000000000,
        0.2000000000000000,
        0.3000000000000000,
    )),
    "empty_struct_field": Empty(**{
    }),
    "false_field": False,
    "i16_byte_map": {
        -1: 1,
        1: 1,
        32767: 1,
    },
    "i16_list": [
        -1,
        0,
        1,
        32767,
    ],
    "i16_set": set((
        -1,
        0,
        1,
        32767,
    )),
    "i32_byte_map": {
        -1: 1,
        1: 1,
        2147483647: 1,
    },
    "i32_list": [
        -1,
        0,
        255,
        65535,
        16777215,
        2147483647,
    ],
    "i32_set": set((
        1,
        2,
        3,
    )),
    "i64_byte_map": {
        -1: 1,
        0: 1,
        1: 1,
        9223372036854775807: 1,
    },
    "i64_list": [
        -1,
        0,
        255,
        65535,
        16777215,
        4294967295,
        1099511627775,
        281474976710655,
        72057594037927935,
        9223372036854775807,
    ],
    "i64_set": set((
        -1,
        0,
        255,
        65535,
        16777215,
        4294967295,
        1099511627775,
        281474976710655,
        72057594037927935,
        9223372036854775807,
    )),
    "list_byte_map": {
        (
        ): 0,
        (
            0,
            1,
        ): 2,
        (
            1,
            2,
            3,
        ): 1,
    },
    "map_byte_map": {
        TFrozenDict({
        }): 0,
        TFrozenDict({
            1: 1,
        }): 1,
        TFrozenDict({
            2: 2,
        }): 2,
    },
    "set_byte_map": {
        frozenset((
        )): 0,
        frozenset((
            0,
            1,
        )): 2,
        frozenset((
            1,
            2,
            3,
        )): 1,
    },
    "string_byte_map": {
        "": 0,
        "first": 1,
        "second": 2,
        "third": 3,
    },
    "string_list": [
        "first",
        "second",
        "third",
    ],
    "string_set": set((
        "first",
        "second",
        "third",
    )),
    "struct_list": [
        Empty(**{
        }),
        Empty(**{
        }),
    ],
    "struct_set": set((
        Empty(**{
        }),
    )),
    "true_field": True,
})
MYCONST = 2
MY_SOME_ENUM = 1
MY_SOME_ENUM_1 = 1
MY_ENUM_MAP = {
    1: 2,
}
EXTRA_CRAZY_MAP = {
    1: StructWithSomeEnum(**{
        "blah": 2,
    }),
}
