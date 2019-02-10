#!/bin/bash
[ -f lendr.env ] || cp app/lendr.env.ex app/lendr.env
set -a
source lendr.env
set +a

# lendr-deploy () {
  # --handler is the path to the executable inside the .zip
  # aws lambda create-function \
  #   --region region \
  #   --function-name lambda-handler \
  #   --memory 128 \
  #   --role arn:aws:iam::account-id:role/execution_role \
  #   --runtime go1.x \
  #   --zip-file fileb://path-to-your-zip-file/handler.zip \
  #   --handler lambda-handler{}
# }

# AWS lambda requires your project as a zip file
lendr-package () {
  zip -r lendr.zip ./build
}
# Bundle, use this command before each lendr-run-test
lendr-bundle () {
  #mkdir -p build/internal/views
  #rsync -avhP --stats --progress internal/views/ build/internal/views/
  lendr-compile
}

lendr-clear-cache () {
  sudo rm -rf build
}

# run this to compile it to static binaries
lendr-compile () {
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/main .
}

# Run the lambda
lendr-run-test () {
    sam local start-api
}

# Publish it for non-local testing
lendr-pub-test () {
    ./ngrok http 127.0.0.1:8080 -subdomain=lendr
}
