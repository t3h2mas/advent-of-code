#!/bin/bash

title=$1

mkdir -p $1/spec $1/lib

cp ./day-nine/Gemfile $1/Gemfile

cd $1

bundle install

echo "Get to work."
