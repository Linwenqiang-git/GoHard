FROM golang:latest as build
WORKDIR  /home/MyProject/Go/GoHard/GinStudy
COPY  .  .
RUN  go  build  -o  main
#产物运行镜像
FROM  alpine:latest
WORKDIR  /home/MyProject/Go/GoHard/GinStudy
COPY  --from=builder /home/MyProject/Go/GoHard/GinStudy  .
EXPOSE 8088
CMD  [“./main]