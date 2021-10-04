#!/bin/bash
set -e

echo "# ================"
echo "# =  GO_BUILDER  ="
echo "# ================"
printf "\n"

echo "> Clean old files"
rm -rvf main main.zip

echo "> Build Linux package..."
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "> Zip package for Serverless environment"
zip main.zip main

echo "> Done!"