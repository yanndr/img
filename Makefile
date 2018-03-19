VERSION=`cat version`
BUILD=`git rev-parse HEAD`
DATE=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/yanndr/imgresize/cmd.Version=${VERSION} -X github.com/yanndr/imgresize/cmd.CommitHash=${BUILD} -X github.com/yanndr/imgresize/cmd.BuildDate=${DATE}"

build:
	go build ${LDFLAGS} 

linux:
	GOOS=linux go build ${LDFLAGS} 

windows:
	GOOS=windows go build ${LDFLAGS} 

install:
	go install ${LDFLAGS} 