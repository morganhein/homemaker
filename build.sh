#/bin/bash
rm -rf build
GOOS=darwin GOARCH=amd64 && go build -o build/$GOOS/homemaker
GOOS=linux GOARCH=amd64 && go build -o build/$GOOS/homemaker
GOOS=linux GOARCH=arm && go build -o build/$GOOS/homemaker-$GOARCH
GOOS=windows GOARCH=amd64 && go build -o build/$GOOS/homemaker.exe
