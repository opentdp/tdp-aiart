#!/bin/sh
#

####################################################################

dateline=`date +%s`

if [ -f node_modules/update.txt ]; then
    lastdate=`cat node_modules/update.txt`
    lastdate=`expr $lastdate + 3600 * 48`
    if [ $dateline -gt $lastdate ]; then
        npm update
    fi
else
    npm install
fi

echo $dateline > node_modules/update.txt

####################################################################

npm run dev
