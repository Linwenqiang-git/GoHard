FROM golang:alpine
WORKDIR  /home/MyProject/Go/GoHard/GinStudy
ADD . ./
ENV GOPROXY="https://goproxy.io"
RUN  go  build  -o  main .
EXPOSE 8089
CMD  ["./main"]
