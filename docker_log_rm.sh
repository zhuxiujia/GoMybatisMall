#!/usr/bin/env bash
path=/var/lib/docker/containers/

cd $path
for file in $(ls)
do
    #[ -d $file ] && echo $file
    if [ -d $file ];then
        echo $file
        cat /dev/null > $file/$file-json.log
      else
        echo 0
    fi
done
