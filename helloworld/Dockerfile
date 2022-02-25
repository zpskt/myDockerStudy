#原镜像
FROM golang:1.17
#设置作者
MAINTAINER zp 2233151897@qq.com

#新建文件夹
RUN mkdir /app
#复制代码
COPY . /app
#设置工作目录
WORKDIR /app

#设置代理
ENV GOPROXY https://goproxy.cn,direct

#go构建可执行文件
RUN go build .

#最终运行docker的命令
ENTRYPOINT  ["./helloworld"]