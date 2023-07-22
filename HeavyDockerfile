FROM golang:latest

WORKDIR /Users/jim/workspace/go-blog-step-by-step
COPY . /Users/jim/workspace/go-blog-step-by-step

RUN export GO111MODULE=on && export GOPROXY=https://goproxy.io && go build .

EXPOSE 8848
ENTRYPOINT ["./go-blog-step-by-step"]