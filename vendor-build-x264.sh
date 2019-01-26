#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/x264-snapshot*
./configure \
    --disable-cli --enable-shared \
    --enable-pic --enable-lto \
    --prefix=$PREFIX
make -j
make install
popd
