#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/ffmpeg*
./configure \
    --disable-all --disable-network \
    --enable-gpl --enable-libx264 \
    --enable-avformat --enable-avcodec --enable-swscale \
    --enable-muxer=mp4 --enable-demuxer=mov \
    --enable-decoder=h264 --enable-encoder=libx264 \
    --enable-decoder=aac --enable-encoder=aac \
    --enable-protocol=file \
    --enable-shared --enable-pic \
    --enable-lto \
    --prefix=$PREFIX
make -j
make install
popd
