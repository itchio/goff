#!/bin/bash -xe
source vendor-env.sh

pushd $GOFF_PREFIX/src/ffmpeg*
./configure \
    --disable-all --disable-network --enable-pthreads \
    --enable-gpl --enable-libx264 \
    --enable-avformat --enable-avcodec --enable-swscale \
    --enable-muxer=mp4 --enable-demuxer=mov \
    --enable-decoder=h264 --enable-encoder=libx264 \
    --enable-decoder=aac --enable-encoder=aac \
    --enable-protocol=file \
    --disable-shared --enable-static \
    --enable-pic \
    --prefix=$GOFF_PREFIX
make -j
make install
popd
