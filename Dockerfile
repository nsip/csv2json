###########################
# INSTRUCTIONS
############################
# BUILD
#	docker build -t nsip/n3-csv2json:latest -t nsip/n3-csv2json:v0.1.0 .
# TEST: docker run -it -v $PWD/test/data:/data -v $PWD/test/config.json:/config.json nsip/n3-csv2json:develop .
# RUN: docker run -d nsip/n3-csv2json:develop
#
# PUSH
#	Public:
#		docker push nsip/n3-csv2json:v0.1.0
#		docker push nsip/n3-csv2json:latest
#
#	Private:
#		docker tag nsip/n3-csv2json:v0.1.0 the.hub.nsip.edu.au:3500/nsip/n3-csv2json:v0.1.0
#		docker tag nsip/n3-csv2json:latest the.hub.nsip.edu.au:3500/nsip/n3-csv2json:latest
#		docker push the.hub.nsip.edu.au:3500/nsip/n3-csv2json:v0.1.0
#		docker push the.hub.nsip.edu.au:3500/nsip/n3-csv2json:latest
#
###########################
# DOCUMENTATION
############################

###########################
# STEP 0 Get them certificates
############################
# (note, step 2 is using alpine now) 
# FROM alpine:latest as certs

############################
# STEP 1 build executable binary (go.mod version)
############################
FROM golang:1.15.0-alpine3.12 as builder
RUN apk --no-cache add ca-certificates
RUN apk update && apk add git
RUN apk add gcc g++
RUN mkdir -p /build
WORKDIR /build
COPY . .
WORKDIR Server
RUN go get github.com/cdutwhu/n3-util/n3cfg
RUN CGO_ENABLED=0 go build -o /build/app

############################
# STEP 2 build a small image
############################
#FROM debian:stretch
FROM alpine
COPY --from=builder /build/app /app
# NOTE - make sure it is the last build that still copies the files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/Server/config/config.toml /
CMD ["./app"]
