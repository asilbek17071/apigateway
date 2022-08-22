#!/bin/bash

CURRENT_DIR=$(pwd)

mkdir $CURRENT_DIR/genproto/user_service
mkdir $CURRENT_DIR/genproto/teacher_service
mkdir $CURRENT_DIR/genproto/student_service
mkdir $CURRENT_DIR/genproto/system_service
mkdir $CURRENT_DIR/genproto/finance_service
mkdir $CURRENT_DIR/genproto/course_service

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/user_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/user_service/ \
        $CURRENT_DIR/protos/user_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/teacher_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/teacher_service/ \
        $CURRENT_DIR/protos/teacher_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/student_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/student_service/ \
        $CURRENT_DIR/protos/student_service/*.proto;


protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/system_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/system_service/ \
        $CURRENT_DIR/protos/system_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/finance_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/finance_service/ \
        $CURRENT_DIR/protos/finance_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/course_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/course_service/ \
        $CURRENT_DIR/protos/course_service/*.proto;


if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i "" -e "s/,omitempty//g" $CURRENT_DIR/genproto/*/*.go
  else
    sed -i -e "s/,omitempty//g" $CURRENT_DIR/genproto/*/*.go
fi