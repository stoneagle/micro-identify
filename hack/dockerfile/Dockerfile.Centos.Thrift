FROM centos:7

# 安装c++基础组件和c++组件
RUN yum install -y gcc libstdc++-devel gcc-c++ libtool flex bison pkgconfig boost-devel libevent-devel zlib-devel python-devel ruby-devel boost-devel-static

RUN yum install -y wget tar autoconf make openssl-devel automake git

# 安装thrift
RUN cd /tmp && \
    wget http://mirrors.shu.edu.cn/apache/thrift/0.11.0/thrift-0.11.0.tar.gz && \
    tar -zxvf ./thrift-0.11.0.tar.gz && \
    cd thrift-0.11.0 && ./configure && make && make install && \
    cd .. && rm -rf ./thrift-0.11.0.tar.gz

RUN echo "/usr/local/lib" >> /etc/ld.so.conf && \
    /sbin/ldconfig -v
