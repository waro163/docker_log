## 概述

采集docker标准输出信息（stdout）和标准出错信息（stderr)；
通过/var/run/docker.sock访问Docker；
监听到容器die的事件后，获取对应的容器ID，再通过docker logs命令获取指定容器全部日志；
可以设置$CONTAINER_NAME来过滤只监听该容器的die事件；

## build

docker build -t docker_log -f ./Dockerfile .

## run

docker run --name docker_log_ins -v /var/run/docker.sock:/var/run/docker.sock -d docker_log