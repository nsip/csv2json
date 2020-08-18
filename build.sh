 #!/bin/bash

set -e

red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
reset=`tput sgr0`

ORIGINALPATH=`pwd`

####

WORKPATH="./Preprocess"

# sudo password
sudopwd="password"

# generate config.go for [Server]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "server"

# Trim Server config.toml for [goclient]
go test -v -timeout 1s -count=1 $WORKPATH/CfgGen -run TestMkCltCfg -args "Path" "Service" "Route" "Server" "Access"
echo "${green}goclient Config.toml Generated${reset}"

# generate config.go fo [goclient]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "goclient"

####

cd ./Server && ./build.sh && cd $ORIGINALPATH && echo "${green}Server Built${reset}"
