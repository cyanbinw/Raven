FROM golang:latest
WORKDIR /Code

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

# src code
COPY ./bin/linux/Web_linux .
COPY ./config.yml .
#RUN go env

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build ./src/Web
EXPOSE 8080
ENTRYPOINT [ "./Web_linux" ]

