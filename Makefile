export GOPROXY=https://goproxy.cn,direct

VERSION="v1.0.0"
COMMIT_ID=`git rev-parse --short HEAD`
BUILD_TIME=`date "+%Y-%m-%d %H:%M:%S"`

LDFLAGS="-w -s \
-X 'main.Version=${VERSION}' \
-X 'main.CommitID=${COMMIT_ID}' \
-X 'main.BuildTime=${BUILD_TIME}'"

all: clear build

clear:
	rm -rf bin/

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/linux/htpasswd main.go
#	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/darwin/htpasswd main.go
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/darwin/htpasswd main.go