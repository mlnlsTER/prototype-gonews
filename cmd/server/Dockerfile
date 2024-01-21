FROM golang:latest AS compiling_stage
RUN mkdir -p /go/src/gonews
WORKDIR /go/src/gonews
ADD server.go .
ADD go.mod .
RUN go install .
 
FROM alpine:latest
LABEL version="1.0.0"
LABEL maintainer="Zhdan Baliuk<balyuk603@gmail.com>"
WORKDIR /root/
COPY --from=compiling_stage /go/bin/gonews .
ENTRYPOINT ./gonews