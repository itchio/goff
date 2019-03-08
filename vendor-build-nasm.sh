#!/bin/bash -xe
source vendor-env.sh

pushd $GOFF_PREFIX/src/nasm*
./autogen.sh
./configure \
    --prefix=$GOFF_PREFIX
make -j
touch nasm.1 ndisasm.1 # don't ask
make install
popd

