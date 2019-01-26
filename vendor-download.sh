#!/bin/bash -xe
source vendor-env.sh

mkdir -p $PREFIX/src
pushd $PREFIX/src
curl -L https://www.nasm.us/pub/nasm/releasebuilds/2.14.02/nasm-2.14.02.tar.bz2 | tar -xj
curl -L https://download.videolan.org/pub/videolan/x264/snapshots/x264-snapshot-20190125-2245.tar.bz2 | tar -xj
curl -L https://ffmpeg.org/releases/ffmpeg-4.1.tar.bz2 | tar -xj
popd
