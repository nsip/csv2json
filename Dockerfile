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
FROM golang:1.15.2-alpine3.12 as builder
RUN apk add --no-cache ca-certificates
RUN apk update && apk add --no-cache git bash
RUN mkdir -p /n3-csv2json
COPY . / /n3-csv2json/
WORKDIR /n3-csv2json/
RUN ["/bin/bash", "-c", "./build_d.sh"]
RUN ["/bin/bash", "-c", "./release_d.sh"] 

############################
# STEP 2 build a small image
############################
#FROM debian:stretch
FROM alpine
COPY --from=builder /n3-csv2json/app/ /
# NOTE - make sure it is the last build that still copies the files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /
CMD ["./server"]
