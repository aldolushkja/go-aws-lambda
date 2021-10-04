# Go AWS Lambda

Sample HTTP Endpoint build with Go and deployed to AWS Lambda

## Uses

1. Go 1.x
2. Go AWS Lambda
2. Go AWS Events

## Build

1. Package in ZIP format

```shell
sh 0_buildLinux.sh
```

2. Package using Docker Image

```shell
sh 0_buildLinux.sh
```

## Test Application

1. Run unit test

```shell
go test
```

2. Print unit test coverage with html output

```shell
go test -coverage
```

## Deploy

1. Use the AWS CLI :: TODO
2. Copy **main.zip** into your source for your AWS Lambda

## Test Deployment

1. Goto: API Gateway Service from your AWS Console
2. Create new HTTP Api
3. Integrate with previously create Lambda Function
4. Create routes for HTTP Verbs (GET,POST, ecc...)
5. Test against endpoint provided by AWS