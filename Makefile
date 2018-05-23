.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)
USERNAME := $(shell id -u -n)
PROJECT := card

CXX = /usr/bin/g++

OPENCV_LIBS = -lopencv_highgui -lopencv_imgcodecs -llibjasper -llibjpeg -llibtiff -llibpng -llibwebp -lopencv_imgproc -lopencv_core  -lpthread -lstdc++ -fopenmp -ldl -lz

DEFINE =-O3 -fPIC -msse2 -msse3 -msse4 -mfpmath=sse -ffast-math -funroll-loops -std=c++11

GOVERSION = 1.10

run-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" up
restart-web: 
	cd hack && docker-compose restart $(PROJECT)-$(USERNAME)-golang
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" rm 

build-thrift:
	cd hack/dockerfile && \
	docker build -f ./Dockerfile.Debian.Thrift -t debian-thrift:0.11.0-cpp .

build-centos:
	cd hack/dockerfile && \
	docker build -f ./Dockerfile.Centos.Thrift -t centos:thrift . && \
	docker build -f ./Dockerfile.Centos.Golang -t centos:golang-$(GOVERSION) --build-arg GOV=$(GOVERSION) .


init-db:
	go run ./initial/db.go 

# rpc-thrfit模式调用cpp图像识别服务
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
	rm -rf $(THRIFT_PREFIX)/service && mkdir $(THRIFT_PREFIX)/service && \
	thrift --gen cpp -out $(THRIFT_PREFIX)/service ./hack/image.thrift

thrift-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:$(THRIFT_PREFIX)/release \
		-v $(PWD)/thrift/service:$(THRIFT_PREFIX)/service \
		-v $(PWD)/build/lib:$(THRIFT_PREFIX)/lib \
		-v $(PWD)/build/include:$(THRIFT_PREFIX)/include \
		-v $(PWD)/build/model:$(THRIFT_PREFIX)/model \
		thrift:0.11.0-cpp \
		$(CXX) -o $(THRIFT_PREFIX)/release/server $(THRIFT_FILES) $(DEFINE) $(THRIFT_INCLUDE) $(THRIFT_LIBS)

# cgo模式调用cpp图像识别服务
CGO_PREFIX = ./cplus

CGO_INCLUDE=-I$(CGO_PREFIX)/include \
		-I$(CGO_PREFIX)/include/opencv \
		-I$(CGO_PREFIX)/include/opencv2

CGO_LIBS = -L$(CGO_PREFIX)/lib \
        -lIPCReading-S-r0.0.1 -lDNN -lopenblas\
		$(OPENCV_LIBS) 

cgo-soft:
	ln -s $(PWD)/build/lib ./backend/cplus/lib && \
	ln -s $(PWD)/build/include ./backend/cplus/include && \
	ln -s $(PWD)/build/model ./backend/cplus/model

cgo-clean:
	rm -rf ./backend/cplus/lib && \
	rm -rf ./backend/cplus/include && \
	rm -rf ./backend/cplus/model

cgo-build: 
	cd backend && \
	$(CXX) -c -o $(CGO_PREFIX)/ipc.o $(CGO_PREFIX)/ipc.cpp $(DEFINE) $(CGO_INCLUDE) $(CGO_LIBS) && ar -rs $(CGO_PREFIX)/lib/libipc.a $(CGO_PREFIX)/ipc.o
