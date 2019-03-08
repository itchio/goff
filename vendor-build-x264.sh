#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/x264-snapshot*
./configure \
    --disable-cli --enable-static \
    --bit-depth=8 --chroma-format=420 \
    --disable-interlaced \
    --enable-pic --enable-strip \
    --prefix=$PREFIX
make -j
make install
popd
