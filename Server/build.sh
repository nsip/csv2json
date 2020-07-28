 #!/bin/bash

set -e

rm -rf ./build/

PROJECTPATH="github.com/nsip/n3-csv2json/Server"
go test -v -timeout 2s $PROJECTPATH/preprocess -run TestGenSvrCfgStruct
go test -v -timeout 2s $PROJECTPATH/config -run TestGenCltCfg -args "Version" "Log" "WebService" "NATS" "Loggly"
go test -v -timeout 2s $PROJECTPATH/preprocess -run TestGenCltCfgStruct

go get

GOARCH=amd64
LDFLAGS="-s -w"
OUT=server

OUTPATH=./build/linux64/
mkdir -p $OUTPATH
CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config/config.toml $OUTPATH

# OUTPATH=./build/win64/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
# mv $OUT.exe $OUTPATH
# cp ./config/config.toml $OUTPATH

# OUTPATH=./build/mac/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./config/config.toml $OUTPATH

# GOARCH=arm
# OUTPATH=./build/linuxarm/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./config/config.toml $OUTPATH
