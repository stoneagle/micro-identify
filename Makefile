.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
USERNAME := $(shell id -u -n)
GROUP := $(shell id -g)
PROJECT := card

run-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" up
restart-web: 
	cd hack && docker-compose restart $(PROJECT)-$(USERNAME)-golang
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" rm 

build-base:
	cd hack/dockerfile && \
	docker build -f ./Dockerfile -t golang-centos:1.8 . && \
	docker build -f ./Dockerfile.beego -t golang-bee:1.8 .

init-db:
	go run ./initial/db.go 

# image identification build
PATH_PREFIX = ./cplus

CXX = /usr/bin/g++

INCLUDE=-I$(PATH_PREFIX)/include \
		-I$(PATH_PREFIX)/include/opencv \
		-I$(PATH_PREFIX)/include/opencv2

OPENCV_LIBS = -lopencv_highgui -lopencv_imgcodecs -llibjasper -llibjpeg -llibtiff -llibpng -llibwebp -lopencv_imgproc -lopencv_core  -lpthread -lstdc++ -fopenmp -ldl -lz

LIBS = -L$(PATH_PREFIX)/lib \
        -lIPCReading-S-r0.0.1 -lDNN -lopenblas\
		$(OPENCV_LIBS) 

DEFINE =-O3 -fPIC -msse2 -msse3 -msse4 -mfpmath=sse -ffast-math -funroll-loops -std=c++11

build-ipc: 
	cd backend && \
	$(CXX) -c -o $(PATH_PREFIX)/ipc.o $(PATH_PREFIX)/ipc.cpp $(DEFINE) $(INCLUDE) $(LIBS) && ar -rs $(PATH_PREFIX)/lib/libipc.a $(PATH_PREFIX)/ipc.o
