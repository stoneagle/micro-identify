FROM identify:thrift-0.11 

COPY ./release/server /usr/src/server
COPY ./build/model /usr/src/model

WORKDIR /usr/src

CMD ["./server"]
