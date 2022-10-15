# go-echo-lambda
APIGateway + Lambda(Go Echo) + ServerlessFramework

- デプロイコマンド
```sh
$ GOARCH=amd64 GOOS=linux go build "-ldflags=-s -w" ./main.go
$ sls deploy --param="account_id=${AWS_ACCOUNT_ID}" --param="allow_id=X.X.X.X/32"
```
