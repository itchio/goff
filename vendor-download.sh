#!/bin/bash -xe
source vendor-env.sh

mkdir -p $GOFF_PREFIX/src
pushd $GOFF_PREFIX/src
# curl -L https://www.nasm.us/pub/nasm/releasebuilds/2.14.02/nasm-2.14.02.tar.bz2 | tar -xj
curl -L https://repo.or.cz/nasm.git/snapshot/74246c499ea4313fb8837977dc0c135fc50567c0.tar.gz | tar -xz
curl -L http://download.videolan.org/pub/videolan/x264/snapshots/x264-snapshot-20190307-2245.tar.bz2 | tar -xj
curl -L https://ffmpeg.org/releases/ffmpeg-4.1.1.tar.bz2 | tar -xj
popd
