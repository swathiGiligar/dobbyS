#!/bin/bash

src=$GOPATH/src
vp=github.com/varunamachi/vaali
sp=${src}/${vp}/
dp=${src}/github.com/swathiGiligar/dobbyS/vendor/${vp}/

rsync -av --exclude='vendor/' --exclude='.git/' ${sp} ${dp} && \
cd cmd/dobby && \
go install && \
dobby serve

