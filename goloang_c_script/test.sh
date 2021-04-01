#!/bin/bash


if [ $# == 0 ]; then
    echo "hello"
    exit 0
fi

for i in $* ; do
    echo "${i}"
done
