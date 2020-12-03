# go-download-manager
golang实现简单的文件下载管理   
使用**apollo**进行文件版本管理，客户端监听配置的变化触发文件下载。同时能在apollo管理端看到已连接的客户端和是否已更新

# Usage

## 配置管理系统apollo部署
[参考](https://github.com/ctripcorp/apollo/wiki/Apollo-Quick-Start-Docker%E9%83%A8%E7%BD%B2)

### 修改docker-compose.yml
#### 修改启动的镜像源
```
apollo-quick-start:
    image: nobodyiam/apollo-quick-start
```
#### 修改端口避免冲突
```
ports:
      - "39001:8080"
      - "39002:8070"
```
#### 安装docker-compose
```
curl -L https://github.com/docker/compose/releases/download/1.25.0/docker-compose-Linux-x86_64 -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose --version
```
### 控制台访问
```
目录下docker-compose up命令启动
转发到8070端口，http://xxx:39002
账号/密码：apollo/admin
```
## 客户端配置
### apollo连接配置
配置文件app.properties
```
{
    "appId":"upgrade", ## apollo命名空间配置的允许的客户端访问id
    "cluster":"default", ## apollo集群名称默认
    "namespaceNames":["application"], ## apollo命名空间
    "ip":"172.16.21.11:39001" ## apollo配置中心连接，转发到8080端口
}
```
### 文件下载管理客户端配置
配置文件config.properties
```
{
    "WatchNs":"application", ## 监听的apollo命名空间
    "WatchKey":"fileVersion", ## 监听的配置key
    "DownLoadDst":".", ## 下载到指定目录
    "DownLoadUrl":"", ## 下载连接，http请求
    "DownLoadFileName":"" ## 下载的文件名
}
```
