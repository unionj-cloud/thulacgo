#!/bin/bash

echo '======retrieve models======'

if [ ! -d "models" ]; then
  echo '======models does not exists, download it======'
  wget "http://thulac.thunlp.org/source/Models_v1_v2(v1_2).zip"
  unzip -o "Models_v1_v2(v1_2).zip"
fi

ls -l -a -h

echo '======models content======'
ls -l -a -h models
