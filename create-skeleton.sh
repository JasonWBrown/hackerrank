#!/bin/bash

# create dir
echo "will create directory $1"
mkdir -p $1 || echo "assuming $1 exists"
# create files
touch $1/main.go $1/main_test.go $1/Makefile & \
    echo "package main" > $1/main.go & \
    echo "func main() {}" >> $1/main.go & \
    echo "package main" > $1/main_test.go


