FROM alpine
RUN mkdir /n3-csv2json
COPY . / /n3-csv2json/
WORKDIR /n3-csv2json/
CMD [ "./server" ]

### docker build --tag=n3-csv2json .

### ! run this docker image
### docker run --name csv2json --net host n3-csv2json:latest

### docker tag IMAGE_ID dockerhub-user/n3-csv2json:latest
### docker login
### docker push dockerhub-user/n3-csv2json