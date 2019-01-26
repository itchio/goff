#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/nasm*
./configure \
    --prefix=$PREFIX
make -j
make install
popd

