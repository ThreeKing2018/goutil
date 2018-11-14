#!/usr/bin/env sh
file=$1 # 文件名
func=$2 # 方法名

if [[ $func && $file ]];then
    go test -v -count=1 $file -test.run $func
elif [ $file ];then
    go test -v -count=1 $file
else
    go test -count=1 
fi
