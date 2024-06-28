FROM centos:7
ADD qq-bot /root/qq-bot
EXPOSE 19997
## operate
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone
USER root
CMD ["./root/qq-bot", "/root/config.yaml"]