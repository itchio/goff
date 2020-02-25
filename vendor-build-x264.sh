#!/bin/bash -xe
source vendor-env.sh

pushd $GOFF_PREFIX/src/x264-*
./configure \
    --disable-cli --enable-static \
    --bit-depth=8 --chroma-format=420 \
    --disable-interlaced \
    --enable-pic --enable-strip \
    --prefix=$GOFF_PREFIX
make -j
make install
popd
