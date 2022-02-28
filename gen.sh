#!/bin/bash

protoc grpc_unary/unary.proto --go_out=plugins=grpc:.
protoc grpc_streaming/streaming.proto --go_out=plugins=grpc:.
