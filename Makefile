.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
USERNAME := $(shell id -u -n)
GROUP := $(shell id -g)
USERNAME := $(shell id -u -n)
PROJECT := identify
GOVERSION = 1.10
DEVELOP_PREFIX =
IDENTIFY_GIT_TAG =

run-web: rm-web-ol 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USERNAME)" up
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USERNAME)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USERNAME)" rm 

run-web-ol: rm-web
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	cd hack && docker-compose -f docker-compose-online.yml -p "$(PROJECT)-web-$(USERNAME)-online" up
stop-web-ol: 
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	cd hack && docker-compose -f docker-compose-online.yml -p "$(PROJECT)-web-$(USERNAME)-online" stop 
rm-web-ol: 
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	cd hack && docker-compose -f docker-compose-online.yml -p "$(PROJECT)-web-$(USERNAME)-online" rm 

build-base:
	cd hack/dockerfile && \
	docker build -f ./Dockerfile.Debian.Thrift -t $(DEVELOP_PREFIX)identify:thrift-0.11 . && \
	docker build -f ./Dockerfile.Debian.Golang -t $(DEVELOP_PREFIX)identify:golang-$(GOVERSION) --build-arg GOV=$(GOVERSION) .

init-db:
	docker exec -w /go/src/toolkit/backend/initial -it card-$(USERNAME)-golang go run ./init.go 

backend-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:/tmp/release \
		-v $(PWD)/backend:/go/src/identify/backend \
		-w /go/src/identify/backend \
		golang:$(GOVERSION) \
		go build -o /tmp/release/backend

tool-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:/tmp/release \
		-v $(PWD)/backend:/go/src/identify/backend \
		-w /go/src/identify/backend/initial \
		golang:$(GOVERSION) \
		go build -o /tmp/release/tool

release: release-backend release-cpp
	docker push $(DEVELOP_PREFIX)identify-backend:$(IDENTIFY_GIT_TAG) && \
	docker push $(DEVELOP_PREFIX)identify-cpp:$(IDENTIFY_GIT_TAG)

release-backend: tool-build backend-build
	docker build -f ./hack/release/Dockerfile.golang -t $(DEVELOP_PREFIX)identify-backend:$(IDENTIFY_GIT_TAG) .

release-cpp: thrift-build
	docker build -f ./hack/release/Dockerfile.cpp -t $(DEVELOP_PREFIX)identify-cpp:$(IDENTIFY_GIT_TAG) .

# rpc-thrfit-cpp-identify
CXX = /usr/bin/g++

OPENCV_LIBS = -lopencv_highgui -lopencv_imgcodecs -llibjasper -llibjpeg -llibtiff -llibpng -llibwebp -lopencv_imgproc -lopencv_core  -lpthread -lstdc++ -fopenmp -ldl -lz

DEFINE =-O3 -fPIC -msse2 -msse3 -msse4 -mfpmath=sse -ffast-math -funroll-loops -std=c++11

THRIFT_PREFIX = /tmp/thrift

THRIFT_INCLUDE=-I $(THRIFT_PREFIX)/include \
		-I $(THRIFT_PREFIX)/include/opencv \
		-I $(THRIFT_PREFIX)/include/opencv2 \
		-I $(THRIFT_PREFIX)/service \
		-I /usr/local/include

THRIFT_LIBS = -L $(THRIFT_PREFIX)/lib \
        -lIPCReading-S-r0.0.1 -lDNN -lopenblas \
		-L /usr/local/lib -lthrift \
		$(OPENCV_LIBS)

THRIFT_FILES = $(THRIFT_PREFIX)/service/ImageService_server.skeleton.cpp \
			   $(THRIFT_PREFIX)/service/ImageService.cpp \
			   $(THRIFT_PREFIX)/service/image_constants.cpp \
			   $(THRIFT_PREFIX)/service/image_types.cpp

thrift-init-golang:
	rm -rf ./backend/rpc/image && mkdir ./backend/rpc/image && \
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/backend:$(THRIFT_PREFIX)/backend \
		-v $(PWD)/hack:$(THRIFT_PREFIX)/hack \
		thrift:0.11.0-cpp \
		thrift --gen go -out $(THRIFT_PREFIX)/backend/rpc $(THRIFT_PREFIX)/hack/image.thrift

thrift-init-cpp:
	rm -rf ./thrift/service && mkdir ./thrift/service && \
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/thrift:$(THRIFT_PREFIX)/thrift \
		-v $(PWD)/hack:$(THRIFT_PREFIX)/hack \
		thrift:0.11.0-cpp \
		thrift --gen cpp -out $(THRIFT_PREFIX)/thrift/service $(THRIFT_PREFIX)/hack/image.thrift

thrift-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:$(THRIFT_PREFIX)/release \
		-v $(PWD)/thrift/service:$(THRIFT_PREFIX)/service \
		-v $(PWD)/build/lib:$(THRIFT_PREFIX)/lib \
		-v $(PWD)/build/include:$(THRIFT_PREFIX)/include \
		-v $(PWD)/build/model:$(THRIFT_PREFIX)/model \
		$(DEVELOP_PREFIX)identify:thrift-0.11 \
		$(CXX) -o $(THRIFT_PREFIX)/release/server $(THRIFT_FILES) $(DEFINE) $(THRIFT_INCLUDE) $(THRIFT_LIBS)
