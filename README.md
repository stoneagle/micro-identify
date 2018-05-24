# 项目介绍
## 基本功能
### backend
1. 基于图像识别的服务，进行业务层的概念管理
### thrift
1. 提供图像识别的rpc服务

## 环境搭建
### 操作流程
1. 安装docker与docker-compose
2. 拉取cpp的build资源
3. 执行make thrift-build生成图像处理的服务
4. 执行make run-web启动服务
5. 映射backend/config目录的配置为.config，并根据需要调整配置
6. 如果使用本地mysql，可以执行makefile的init-db初始化db 

## 相关工具 
### hack  
1. 使用dockerfile打造基础和发布镜像
2. 使用docker-compose搭建开发环境
3. 使用glide完成第三方包管理
### backend
1. 框架gin
2. mysql数据库xorm
3. redis缓存go-redis
4. log日志zap
### thrift
1. 基于hack/image.thrift管理rpc通信接口
2. 通过Makefile命令进行开发和发布维护
