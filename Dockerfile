FROM golang:1.13.4
WORKDIR /go/src/github.com/chrisevett/helmsync

RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega/...
COPY . .
RUN ginkgo -r -cover
RUN GOOS=linux go build 

FROM alpine:3.10

RUN apk add curl git libc6-compat
RUN curl https://storage.googleapis.com/kubernetes-helm/helm-v2.15.2-linux-amd64.tar.gz -o helm-v2.15.2-linux-amd64.tar.gz
RUN tar -xzvf helm-v2.15.2-linux-amd64.tar.gz
RUN mv linux-amd64/helm /usr/local/bin/helm
RUN helm init --client-only
WORKDIR /app
COPY --from=0 /go/src/github.com/chrisevett/helmsync/helmsync /app/helmsync
CMD ["/app/helmsync"]
