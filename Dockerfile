FROM golang:latest
WORKDIR /Code
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

# src code
COPY . .
#RUN go env
RUN go mod tidy
RUN go build ./src/Web
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build ./src/Web
EXPOSE 8080
ENTRYPOINT [ "./Web" ]
