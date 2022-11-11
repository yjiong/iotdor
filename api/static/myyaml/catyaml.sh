#!/bin/bash
#########################################################################
# File Name: catyaml.sh
# Author: yjiong
# mail: 4418229@qq.com
# Created Time: 2022-08-23 19:05:52
#########################################################################
set -e
PWD=$(cd $(dirname $0);pwd)
pushd $PWD
cp  -f ./iotdor.yaml ../iotdor.yaml
cat {device,user,organizt,organiz,components}.yaml >> ../iotdor.yaml
popd

