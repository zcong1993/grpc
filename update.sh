#!/bin/bash
protoc -I echo/ echo/echo.proto --go_out=plugins=grpc:echo
