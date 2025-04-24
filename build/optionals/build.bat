@echo off
echo Building DiViNE Loader as Windows GUI application...
go build -ldflags="-H windowsgui" main.go
echo [BUILD COMPLETE]: DiViNE Has Awoken
echo Use main.exe --no-error to run without showing the error message.
echo Use go run build/build.go --analysis to build for analysis mode.
