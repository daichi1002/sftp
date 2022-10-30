lint:
	golangci-lint run
	protolint lint -fix .

run:
	go run main.go

test:
	# profile.outというファイルテストのカバー内容を吐き出す
	go test ./... -v -coverprofile=profile.out -p 1
	# go toolを用いてprofile.htmlを作成する
	go tool cover -func=profile.out

task:
	golangci-lint run
	go test ./... -v -coverprofile=profile.out -p 1
	go tool cover -func=profile.out