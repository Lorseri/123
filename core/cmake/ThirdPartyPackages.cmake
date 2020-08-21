# Copyright (C) 2019-2020 Zilliz. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under the License.

set(MILVUS_THIRDPARTY_DEPENDENCIES

        SQLite
        fiu
        AWS
        )

message(STATUS "Using ${MILVUS_DEPENDENCY_SOURCE} approach to find dependencies")

# For each dependency, set dependency source to global default, if unset
foreach (DEPENDENCY ${MILVUS_THIRDPARTY_DEPENDENCIES})
    if ("${${DEPENDENCY}_SOURCE}" STREQUAL "")
        set(${DEPENDENCY}_SOURCE ${MILVUS_DEPENDENCY_SOURCE})
    endif ()
endforeach ()

macro(build_dependency DEPENDENCY_NAME)
    if ("${DEPENDENCY_NAME}" STREQUAL "SQLite")
        build_sqlite()
    elseif ("${DEPENDENCY_NAME}" STREQUAL "fiu")
        build_fiu()
    elseif("${DEPENDENCY_NAME}" STREQUAL "AWS")
        build_aws()
    else ()
        message(FATAL_ERROR "Unknown thirdparty dependency to build: ${DEPENDENCY_NAME}")
    endif ()
endmacro()

# ----------------------------------------------------------------------
# Identify OS
if (UNIX)
    if (APPLE)
        set(CMAKE_OS_NAME "osx" CACHE STRING "Operating system name" FORCE)
    else (APPLE)
        ## Check for Debian GNU/Linux ________________
        find_file(DEBIAN_FOUND debian_version debconf.conf
                PATHS /etc
                )
        if (DEBIAN_FOUND)
            set(CMAKE_OS_NAME "debian" CACHE STRING "Operating system name" FORCE)
        endif (DEBIAN_FOUND)
        ##  Check for Fedora _________________________
        find_file(FEDORA_FOUND fedora-release
                PATHS /etc
                )
        if (FEDORA_FOUND)
            set(CMAKE_OS_NAME "fedora" CACHE STRING "Operating system name" FORCE)
        endif (FEDORA_FOUND)
        ##  Check for RedHat _________________________
        find_file(REDHAT_FOUND redhat-release inittab.RH
                PATHS /etc
                )
        if (REDHAT_FOUND)
            set(CMAKE_OS_NAME "redhat" CACHE STRING "Operating system name" FORCE)
        endif (REDHAT_FOUND)
        ## Extra check for Ubuntu ____________________
        if (DEBIAN_FOUND)
            ## At its core Ubuntu is a Debian system, with
            ## a slightly altered configuration; hence from
            ## a first superficial inspection a system will
            ## be considered as Debian, which signifies an
            ## extra check is required.
            find_file(UBUNTU_EXTRA legal issue
                    PATHS /etc
                    )
            if (UBUNTU_EXTRA)
                ## Scan contents of file
                file(STRINGS ${UBUNTU_EXTRA} UBUNTU_FOUND
                        REGEX Ubuntu
                        )
                ## Check result of string search
                if (UBUNTU_FOUND)
                    set(CMAKE_OS_NAME "ubuntu" CACHE STRING "Operating system name" FORCE)
                    set(DEBIAN_FOUND FALSE)

                    find_program(LSB_RELEASE_EXEC lsb_release)
                    execute_process(COMMAND ${LSB_RELEASE_EXEC} -rs
                            OUTPUT_VARIABLE LSB_RELEASE_ID_SHORT
                            OUTPUT_STRIP_TRAILING_WHITESPACE
                            )
                    STRING(REGEX REPLACE "\\." "_" UBUNTU_VERSION "${LSB_RELEASE_ID_SHORT}")
                endif (UBUNTU_FOUND)
            endif (UBUNTU_EXTRA)
        endif (DEBIAN_FOUND)
    endif (APPLE)
endif (UNIX)

# ----------------------------------------------------------------------
# thirdparty directory
set(THIRDPARTY_DIR "${MILVUS_SOURCE_DIR}/thirdparty")

macro(resolve_dependency DEPENDENCY_NAME)
    if (${DEPENDENCY_NAME}_SOURCE STREQUAL "AUTO")
        find_package(${DEPENDENCY_NAME} MODULE)
        if (NOT ${${DEPENDENCY_NAME}_FOUND})
            build_dependency(${DEPENDENCY_NAME})
        endif ()
    elseif (${DEPENDENCY_NAME}_SOURCE STREQUAL "BUNDLED")
        build_dependency(${DEPENDENCY_NAME})
    elseif (${DEPENDENCY_NAME}_SOURCE STREQUAL "SYSTEM")
        find_package(${DEPENDENCY_NAME} REQUIRED)
    endif ()
endmacro()

# ----------------------------------------------------------------------
# ExternalProject options

string(TOUPPER ${CMAKE_BUILD_TYPE} UPPERCASE_BUILD_TYPE)

set(EP_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${CMAKE_CXX_FLAGS_${UPPERCASE_BUILD_TYPE}}")
set(EP_C_FLAGS "${CMAKE_C_FLAGS} ${CMAKE_C_FLAGS_${UPPERCASE_BUILD_TYPE}}")

# Set -fPIC on all external projects
set(EP_CXX_FLAGS "${EP_CXX_FLAGS} -fPIC")
set(EP_C_FLAGS "${EP_C_FLAGS} -fPIC")

# CC/CXX environment variables are captured on the first invocation of the
# builder (e.g make or ninja) instead of when CMake is invoked into to build
# directory. This leads to issues if the variables are exported in a subshell
# and the invocation of make/ninja is in distinct subshell without the same
# environment (CC/CXX).
set(EP_COMMON_TOOLCHAIN -DCMAKE_C_COMPILER=${CMAKE_C_COMPILER}
        -DCMAKE_CXX_COMPILER=${CMAKE_CXX_COMPILER})

if (CMAKE_AR)
    set(EP_COMMON_TOOLCHAIN ${EP_COMMON_TOOLCHAIN} -DCMAKE_AR=${CMAKE_AR})
endif ()

if (CMAKE_RANLIB)
    set(EP_COMMON_TOOLCHAIN ${EP_COMMON_TOOLCHAIN} -DCMAKE_RANLIB=${CMAKE_RANLIB})
endif ()

# External projects are still able to override the following declarations.
# cmake command line will favor the last defined variable when a duplicate is
# encountered. This requires that `EP_COMMON_CMAKE_ARGS` is always the first
# argument.
set(EP_COMMON_CMAKE_ARGS
        ${EP_COMMON_TOOLCHAIN}
        -DCMAKE_BUILD_TYPE=${CMAKE_BUILD_TYPE}
        -DCMAKE_C_FLAGS=${EP_C_FLAGS}
        -DCMAKE_C_FLAGS_${UPPERCASE_BUILD_TYPE}=${EP_C_FLAGS}
        -DCMAKE_CXX_FLAGS=${EP_CXX_FLAGS}
        -DCMAKE_CXX_FLAGS_${UPPERCASE_BUILD_TYPE}=${EP_CXX_FLAGS})

if (NOT MILVUS_VERBOSE_THIRDPARTY_BUILD)
    set(EP_LOG_OPTIONS LOG_CONFIGURE 1 LOG_BUILD 1 LOG_INSTALL 1 LOG_DOWNLOAD 1)
else ()
    set(EP_LOG_OPTIONS)
endif ()

# Ensure that a default make is set
if ("${MAKE}" STREQUAL "")
    find_program(MAKE make)
endif ()

if (NOT DEFINED MAKE_BUILD_ARGS)
    set(MAKE_BUILD_ARGS "-j8")
endif ()
message(STATUS "Third Party MAKE_BUILD_ARGS = ${MAKE_BUILD_ARGS}")

# ----------------------------------------------------------------------
# Find pthreads

set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

# ----------------------------------------------------------------------
# Versions and URLs for toolchain builds, which also can be used to configure
# offline builds

# Read toolchain versions from cpp/thirdparty/versions.txt
file(STRINGS "${THIRDPARTY_DIR}/versions.txt" TOOLCHAIN_VERSIONS_TXT)
foreach (_VERSION_ENTRY ${TOOLCHAIN_VERSIONS_TXT})
    # Exclude comments
    if (NOT _VERSION_ENTRY MATCHES "^[^#][A-Za-z0-9-_]+_VERSION=")
        continue()
    endif ()

    string(REGEX MATCH "^[^=]*" _LIB_NAME ${_VERSION_ENTRY})
    string(REPLACE "${_LIB_NAME}=" "" _LIB_VERSION ${_VERSION_ENTRY})

    # Skip blank or malformed lines
    if (${_LIB_VERSION} STREQUAL "")
        continue()
    endif ()

    # For debugging
    #message(STATUS "${_LIB_NAME}: ${_LIB_VERSION}")

    set(${_LIB_NAME} "${_LIB_VERSION}")
endforeach ()


if (DEFINED ENV{MILVUS_SQLITE_URL})
    set(SQLITE_SOURCE_URL "$ENV{MILVUS_SQLITE_URL}")
else ()
    set(SQLITE_SOURCE_URL
            "https://www.sqlite.org/2019/sqlite-autoconf-${SQLITE_VERSION}.tar.gz")
endif ()

if (DEFINED ENV{MILVUS_FIU_URL})
    set(FIU_SOURCE_URL "$ENV{MILVUS_FIU_URL}")
else ()
    set(FIU_SOURCE_URL "https://github.com/albertito/libfiu/archive/${FIU_VERSION}.tar.gz"
                       "https://gitee.com/quicksilver/libfiu/repository/archive/${FIU_VERSION}.zip")
endif ()

if (DEFINED ENV{MILVUS_AWS_URL})
    set(AWS_SOURCE_URL "$ENV{MILVUS_AWS_URL}")
else ()
    set(AWS_SOURCE_URL "https://github.com/aws/aws-sdk-cpp/archive/${AWS_VERSION}.tar.gz")
endif ()


# ----------------------------------------------------------------------
# SQLite

macro(build_sqlite)
    message(STATUS "Building SQLITE-${SQLITE_VERSION} from source")
    set(SQLITE_PREFIX "${CMAKE_CURRENT_BINARY_DIR}/sqlite_ep-prefix/src/sqlite_ep")
    set(SQLITE_INCLUDE_DIR "${SQLITE_PREFIX}/include")
    set(SQLITE_STATIC_LIB
            "${SQLITE_PREFIX}/lib/${CMAKE_STATIC_LIBRARY_PREFIX}sqlite3${CMAKE_STATIC_LIBRARY_SUFFIX}")

    set(SQLITE_CONFIGURE_ARGS
            "--prefix=${SQLITE_PREFIX}"
            "CC=${CMAKE_C_COMPILER}"
            "CXX=${CMAKE_CXX_COMPILER}"
            "CFLAGS=${EP_C_FLAGS}"
            "CXXFLAGS=${EP_CXX_FLAGS}")

    ExternalProject_Add(sqlite_ep
            URL
            ${SQLITE_SOURCE_URL}
            ${EP_LOG_OPTIONS}
            URL_MD5
            "3c68eb400f8354605736cd55400e1572"
            CONFIGURE_COMMAND
            "./configure"
            ${SQLITE_CONFIGURE_ARGS}
            BUILD_COMMAND
            ${MAKE}
            ${MAKE_BUILD_ARGS}
            BUILD_IN_SOURCE
            1
            BUILD_BYPRODUCTS
            "${SQLITE_STATIC_LIB}")

    file(MAKE_DIRECTORY "${SQLITE_INCLUDE_DIR}")
    add_library(sqlite STATIC IMPORTED)
    set_target_properties(
            sqlite
            PROPERTIES IMPORTED_LOCATION "${SQLITE_STATIC_LIB}"
            INTERFACE_INCLUDE_DIRECTORIES "${SQLITE_INCLUDE_DIR}")

    add_dependencies(sqlite sqlite_ep)
endmacro()

if (MILVUS_WITH_SQLITE)
    resolve_dependency(SQLite)
    include_directories(SYSTEM "${SQLITE_INCLUDE_DIR}")
    link_directories(SYSTEM ${SQLITE_PREFIX}/lib/)
endif ()

# ----------------------------------------------------------------------
# fiu
macro(build_fiu)
    message(STATUS "Building FIU-${FIU_VERSION} from source")
    set(FIU_PREFIX "${CMAKE_CURRENT_BINARY_DIR}/fiu_ep-prefix/src/fiu_ep")
    set(FIU_SHARED_LIB "${FIU_PREFIX}/lib/${CMAKE_SHARED_LIBRARY_PREFIX}fiu${CMAKE_SHARED_LIBRARY_SUFFIX}")
    set(FIU_INCLUDE_DIR "${FIU_PREFIX}/include")

    ExternalProject_Add(fiu_ep
            URL
            ${FIU_SOURCE_URL}
            ${EP_LOG_OPTIONS}
            URL_MD5
            "75f9d076daf964c9410611701f07c61b"
            CONFIGURE_COMMAND
            ""
            BUILD_IN_SOURCE
            1
            BUILD_COMMAND
            ${MAKE}
            ${MAKE_BUILD_ARGS}
            INSTALL_COMMAND
            ${MAKE}
            "PREFIX=${FIU_PREFIX}"
            install
            BUILD_BYPRODUCTS
            ${FIU_SHARED_LIB}
            )

        file(MAKE_DIRECTORY "${FIU_INCLUDE_DIR}")
        add_library(fiu SHARED IMPORTED)
    set_target_properties(fiu
        PROPERTIES IMPORTED_LOCATION "${FIU_SHARED_LIB}"
        INTERFACE_INCLUDE_DIRECTORIES "${FIU_INCLUDE_DIR}")

    add_dependencies(fiu fiu_ep)
endmacro()

resolve_dependency(fiu)

get_target_property(FIU_INCLUDE_DIR fiu INTERFACE_INCLUDE_DIRECTORIES)
include_directories(SYSTEM ${FIU_INCLUDE_DIR})


# ----------------------------------------------------------------------
# aws
macro(build_aws)
    message(STATUS "Building aws-${AWS_VERSION} from source")
    set(AWS_PREFIX "${CMAKE_CURRENT_BINARY_DIR}/aws_ep-prefix/src/aws_ep")

    set(AWS_CMAKE_ARGS
            ${EP_COMMON_TOOLCHAIN}
            "-DCMAKE_INSTALL_PREFIX=${AWS_PREFIX}"
            -DCMAKE_BUILD_TYPE=Release
            -DCMAKE_INSTALL_LIBDIR=lib
            -DBUILD_ONLY=s3
            -DBUILD_SHARED_LIBS=off
            -DENABLE_TESTING=off
            -DENABLE_UNITY_BUILD=on
            -DNO_ENCRYPTION=off)

    set(AWS_CPP_SDK_CORE_STATIC_LIB
            "${AWS_PREFIX}/lib/${CMAKE_STATIC_LIBRARY_PREFIX}aws-cpp-sdk-core${CMAKE_STATIC_LIBRARY_SUFFIX}")
    set(AWS_CPP_SDK_S3_STATIC_LIB
            "${AWS_PREFIX}/lib/${CMAKE_STATIC_LIBRARY_PREFIX}aws-cpp-sdk-s3${CMAKE_STATIC_LIBRARY_SUFFIX}")
    set(AWS_INCLUDE_DIR "${AWS_PREFIX}/include")
    set(AWS_CMAKE_ARGS
            ${AWS_CMAKE_ARGS}
            -DCMAKE_C_COMPILER=${CMAKE_C_COMPILER}
            -DCMAKE_CXX_COMPILER=${CMAKE_CXX_COMPILER}
            -DCMAKE_C_FLAGS=${EP_C_FLAGS}
            -DCMAKE_CXX_FLAGS=${EP_CXX_FLAGS})

    ExternalProject_Add(aws_ep
            ${EP_LOG_OPTIONS}
            CMAKE_ARGS
            ${AWS_CMAKE_ARGS}
            BUILD_COMMAND
            ${MAKE}
            ${MAKE_BUILD_ARGS}
            INSTALL_DIR
            ${AWS_PREFIX}
            URL
            ${AWS_SOURCE_URL}
            BUILD_BYPRODUCTS
            "${AWS_CPP_SDK_S3_STATIC_LIB}"
            "${AWS_CPP_SDK_CORE_STATIC_LIB}")

    file(MAKE_DIRECTORY "${AWS_INCLUDE_DIR}")
    add_library(aws-cpp-sdk-s3 STATIC IMPORTED)
    add_library(aws-cpp-sdk-core STATIC IMPORTED)

    set_target_properties(aws-cpp-sdk-s3
            PROPERTIES
            IMPORTED_LOCATION "${AWS_CPP_SDK_S3_STATIC_LIB}"
            INTERFACE_INCLUDE_DIRECTORIES "${AWS_INCLUDE_DIR}"
            )

    set_target_properties(aws-cpp-sdk-core
            PROPERTIES
            IMPORTED_LOCATION "${AWS_CPP_SDK_CORE_STATIC_LIB}"
            INTERFACE_INCLUDE_DIRECTORIES "${AWS_INCLUDE_DIR}"
            )

    if(REDHAT_FOUND)
        set_target_properties(aws-cpp-sdk-s3
                PROPERTIES
                INTERFACE_LINK_LIBRARIES
                "${AWS_PREFIX}/lib64/libaws-c-event-stream.a;${AWS_PREFIX}/lib64/libaws-checksums.a;${AWS_PREFIX}/lib64/libaws-c-common.a")
        set_target_properties(aws-cpp-sdk-core
                PROPERTIES
                INTERFACE_LINK_LIBRARIES
                "${AWS_PREFIX}/lib64/libaws-c-event-stream.a;${AWS_PREFIX}/lib64/libaws-checksums.a;${AWS_PREFIX}/lib64/libaws-c-common.a")
    else()
        set_target_properties(aws-cpp-sdk-s3
                PROPERTIES
                INTERFACE_LINK_LIBRARIES
                "${AWS_PREFIX}/lib/libaws-c-event-stream.a;${AWS_PREFIX}/lib/libaws-checksums.a;${AWS_PREFIX}/lib/libaws-c-common.a")
        set_target_properties(aws-cpp-sdk-core
                PROPERTIES
                INTERFACE_LINK_LIBRARIES
                "${AWS_PREFIX}/lib/libaws-c-event-stream.a;${AWS_PREFIX}/lib/libaws-checksums.a;${AWS_PREFIX}/lib/libaws-c-common.a")
    endif()

    add_dependencies(aws-cpp-sdk-s3 aws_ep)
    add_dependencies(aws-cpp-sdk-core aws_ep)

endmacro()

if(MILVUS_WITH_AWS)
    resolve_dependency(AWS)

    link_directories(SYSTEM ${AWS_PREFIX}/lib)

    get_target_property(AWS_CPP_SDK_S3_INCLUDE_DIR aws-cpp-sdk-s3 INTERFACE_INCLUDE_DIRECTORIES)
    include_directories(SYSTEM ${AWS_CPP_SDK_S3_INCLUDE_DIR})

    get_target_property(AWS_CPP_SDK_CORE_INCLUDE_DIR aws-cpp-sdk-core INTERFACE_INCLUDE_DIRECTORIES)
    include_directories(SYSTEM ${AWS_CPP_SDK_CORE_INCLUDE_DIR})

endif()
