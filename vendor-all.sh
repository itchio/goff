#!/bin/bash -xe
./vendor-download.sh
./vendor-build-nasm.sh
./vendor-build-x264.sh
./vendor-build-ffmpeg.sh
