#/bin/bash
rm -rf build
gox -os="windows linux darwin" -arch="amd64" -output="build/{{.OS}}/{{.Arch}}/{{.Dir}}"