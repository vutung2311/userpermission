# Exercise solution for user permission problem

## How to run
* Install Golang as per: https://golang.org/doc/install
* Clone source of this project into $GOROOT
```
git clone https://github.com/vutung2311/userpermission $GOROOT/src/
```
* Run example input and compare with output
```
cat input1.txt | go run $GOROOT/src/userpermission/cmd/main.go | diff -u output1.txt -
cat input2.txt | go run $GOROOT/src/userpermission/cmd/main.go | diff -u output2.txt -
```
* Change directory to $GOROOT/src/userpermission and then run test
```
go test ./...
```
