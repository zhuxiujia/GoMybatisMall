::把create_statik.bat拖到cmd窗口执行,将在D:\GOPATH\src\go-p2p\microservice\everything\src\resource下生成/statik/statik.go
cd ../core/com/example/dao
del /S /Q statik
go run ../../../../statik/main/statik.go -src=./