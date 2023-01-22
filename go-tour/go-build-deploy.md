## Go Module
github.com/vjammi/go-learn/go-tour/mr> go mod tidy
    go: finding module for package github.com/vjammi/go-learn/go-tour/mr
    go: downloading github.com/vjammi/go-learn/go-tour/mr v0.0.0-20221230025136-11275fa706d8
    go: downloading github.com/vjammi/go-learn v0.0.0-20221230025136-11275fa706d8
    go: found github.com/vjammi/go-learn/go-tour/mr in github.com/vjammi/go-learn/go-tour/mr v0.0.0-20221230025136-11275fa706d8

## Use the module in another module
github.com\vjammi\go-learn\go-tour\main> go install  github.com/vjammi/go-learn/go-tour/mr
github.com\vjammi\go-learn\go-tour\main> go install  github.com/vjammi/go-learn/go-tour/mr
        no required module provides package github.com/vjammi/go-learn/go-tour/mr; to add it:
        go get github.com/vjammi/go-learn/go-tour/mr
github.com\vjammi\go-learn\go-tour\main> go get github.com/vjammi/go-learn/go-tour/mr
        go: downloading github.com/vjammi/go-learn/go-tour/mr v0.0.0-20230122043044-a2e662f89da0
        go: added github.com/vjammi/go-learn/go-tour/mr v0.0.0-20230122043044-a2e662f89da0

---
github.com/vjammi/go-learn/go-tour/main> go mod init
    go: creating new go.mod: module github.com/vjammi/go-learn/go-tour/main
    go: to add module requirements and sums:
    go mod tidy
github.com/vjammi/go-learn/go-tour/main> go mod tidy
github.com/vjammi/go-learn/go-tour/main> go run main.go

## ??? Build and Deploy
touch main.go
touch main_test.go

go build main.go
go build -o main -ldflags=-X=main.version=${VERSION} main.go
go test
go run main.go
Commit to git

git init
git add .
git commit -m "initial commit"
git remote add origin git@github.com:USERNAME/golang-pipeline
git branch -M main
git push -u origin main

git tag v1.0.0
git push --tags