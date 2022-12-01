#!/bin/bash

day_directories=$(ls | grep -E '^\d\d$');

for d in $day_directories; do
  (cd $d && go test -v ./...)
done
