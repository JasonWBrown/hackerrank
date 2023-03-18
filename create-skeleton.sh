#!/bin/bash

# create dir
echo "will create directory $1"
mkdir $1 || echo "assuming $1 exists"
# create files
cd $1 & \
    touch main.go main_test.go Makfile & \
    echo "package main" > main.go & \
    echo "package main" > main_test.go

