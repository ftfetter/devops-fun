#!/bin/bash

################################################
# Script to prepare and run GoLang application #
################################################

#download dependencies
go get github.com/gorilla/mux

#compile go application
go install calc-service

#run application
$GOPATH/bin/calc-service