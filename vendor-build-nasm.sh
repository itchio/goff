#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/nasm*
./autogen.sh
./configure \
    --prefix=$PREFIX
make -j
make install
popd

