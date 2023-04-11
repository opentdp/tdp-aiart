#!/bin/sh
#

set -e
set -o noglob

###########################################

if [ -n "$GITHUB_WORKSPACE" ]; then
    WKDIR=$GITHUB_WORKSPACE
else
    WKDIR=`dirname $0`
fi

###########################################

echo ">> Prepare workspace"

cd $WKDIR
chmod +x $WKDIR/*/build.sh

###########################################

echo ">> Compile frontend components"

cd $WKDIR/frontend
npm i && ./build.sh

###########################################

echo ">> Compile backend components"

cd $WKDIR/factory
go mod tidy && ./build.sh

###########################################

echo ">> Gzip binary files"

cd $WKDIR

if [ -n "$WITH_UPX" ] && type upx >/dev/null 2>&1; then
    upx `find build/ -type f`
fi

for app in `ls build`; do
    gzip build/$app
done
