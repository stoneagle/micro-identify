FROM debian:jessie

ENV THRIFT_VERSION=0.11.0

RUN buildDeps="automake bison curl flex g++ libboost-dev libboost-filesystem-dev libboost-program-options-dev libboost-system-dev libboost-test-dev libevent-dev libssl-dev libtool make pkg-config"; \
    apt-get update \
    && apt-get install -y --no-install-recommends $buildDeps \
    && rm -rf /var/lib/apt/lists/* \
	&& curl -sSL "http://apache.mirrors.spacedump.net/thrift/$THRIFT_VERSION/thrift-$THRIFT_VERSION.tar.gz" -o thrift.tar.gz \
	&& mkdir -p /usr/src/thrift \
    && tar zxf thrift.tar.gz -C /usr/src/thrift --strip-components=1 \
    && rm thrift.tar.gz \
	&& cd /usr/src/thrift \
    && ./configure --libdir=/usr/lib --without-python --without-java \
	&& make  && make install \
    && cd /  && rm -rf /usr/src/thrift

CMD ["bash"]
