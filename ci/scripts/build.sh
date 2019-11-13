#!/bin/bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
SCRIPTS_DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

MILVUS_CORE_DIR="${SCRIPTS_DIR}/../../core"
CORE_BUILD_DIR="${MILVUS_CORE_DIR}/cmake_build"
BUILD_TYPE="Debug"
BUILD_UNITTEST="OFF"
INSTALL_PREFIX="/opt/milvus"
BUILD_COVERAGE="OFF"
USE_JFROG_CACHE="OFF"
RUN_CPPLINT="OFF"
CPU_VERSION="ON"
WITH_MKL="OFF"
CUDA_COMPILER=/usr/local/cuda/bin/nvcc

while getopts "o:t:b:gulcjmh" arg
do
        case $arg in
             o)
                INSTALL_PREFIX=$OPTARG
                ;;
             t)
                BUILD_TYPE=$OPTARG # BUILD_TYPE
                ;;
             b)
                CORE_BUILD_DIR=$OPTARG # CORE_BUILD_DIR
                ;;
             g)
                CPU_VERSION="OFF";
                ;;
             u)
                echo "Build and run unittest cases" ;
                BUILD_UNITTEST="ON";
                ;;
             l)
                RUN_CPPLINT="ON"
                ;;
             c)
                BUILD_COVERAGE="ON"
                ;;
             j)
                USE_JFROG_CACHE="ON"
                ;;
             m)
                WITH_MKL="ON"
                ;;
             h) # help
                echo "

parameter:
-o: install prefix(default: /opt/milvus)
-t: build type(default: Debug)
-b: core code build directory
-g: gpu version
-u: building unit test options(default: OFF)
-l: run cpplint, clang-format and clang-tidy(default: OFF)
-c: code coverage(default: OFF)
-j: use jfrog cache build directory(default: OFF)
-m: build with MKL(default: OFF)
-h: help

usage:
./build.sh -o \${INSTALL_PREFIX} -t \${BUILD_TYPE} -b \${CORE_BUILD_DIR} [-u] [-l] [-c] [-j] [-m] [-h]
                "
                exit 0
                ;;
             ?)
                echo "ERROR! unknown argument"
        exit 1
        ;;
        esac
done

if [[ ! -d ${CORE_BUILD_DIR} ]]; then
    mkdir ${CORE_BUILD_DIR}
fi

pushd ${CORE_BUILD_DIR}

CMAKE_CMD="cmake \
-DCMAKE_INSTALL_PREFIX=${INSTALL_PREFIX}
-DCMAKE_BUILD_TYPE=${BUILD_TYPE} \
-DCMAKE_CUDA_COMPILER=${CUDA_COMPILER} \
-DMILVUS_CPU_VERSION=${CPU_VERSION} \
-DBUILD_UNIT_TEST=${BUILD_UNITTEST} \
-DBUILD_COVERAGE=${BUILD_COVERAGE} \
-DUSE_JFROG_CACHE=${USE_JFROG_CACHE} \
-DBUILD_FAISS_WITH_MKL=${WITH_MKL} \
-DArrow_SOURCE=AUTO \
${MILVUS_CORE_DIR}"
echo ${CMAKE_CMD}
${CMAKE_CMD}

# compile and build
make -j8 || exit 1
make install || exit 1
