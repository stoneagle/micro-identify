FROM centos:7

# 安装基础组件
RUN yum install -y wget tar

# 安装go组件
RUN cd /tmp && \
    wget https://mirrors.nju.edu.cn/golang/go1.8.linux-amd64.tar.gz && \
    tar -zxvf ./go1.8.linux-amd64.tar.gz && \
    mkdir /home/roobo && cp -r ./go /home/roobo/go && \
    mkdir /home/roobo/go/gopkg && \
    rm -rf ./go && \
    rm ./go1.8.linux-amd64.tar.gz

# 安装c++组件
RUN yum install -y gcc libstdc++-devel gcc-c++

# 配置golang环境变量
ENV GOARCH amd64
ENV GOOS linux
ENV GOROOT /home/roobo/go
ENV GOPATH $GOROOT/gopkg:/home/roobo/go_project
ENV PATH $PATH:$GOROOT/bin:$GOROOT/gopkg/bin:/home/wuzhongyang/go_project/bin
