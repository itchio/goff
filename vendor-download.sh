#!/bin/bash -xe
source vendor-env.sh

mkdir -p $GOFF_PREFIX/src
pushd $GOFF_PREFIX/src
curl -L https://www.nasm.us/pub/nasm/releasebuilds/2.14.02/nasm-2.14.02.tar.bz2 | tar -xj
curl -L https://code.videolan.org/videolan/x264/-/archive/1771b556ee45207f8711744ccbd5d42a3949b14c/x264-1771b556ee45207f8711744ccbd5d42a3949b14c.tar.bz2 | tar -xj
curl -L https://ffmpeg.org/releases/ffmpeg-4.2.2.tar.bz2 | tar -xj
popd
