FROM alpine
RUN mkdir n3-csv2json
COPY ./n3-csv2json /n3-csv2json
WORKDIR /n3-csv2json/Server/build/linux64
CMD [ "./server" ]

### docker build --tag=n3-csv2json .

### docker tag IMAGE_ID cdutwhu/n3-csv2json:latest
### docker login
### docker push cdutwhu/n3-csv2json