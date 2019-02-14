#!/bin/bash -xe
source vendor-env.sh

pushd $PREFIX/src/nasm*
./autogen.sh
./configure \
    --prefix=$PREFIX
make -j
touch nasm.1 ndisasm.1 # don't ask
make install
popd

