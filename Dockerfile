FROM golang:alpine
# Install git and build in the docker image
RUN apk update && apk add git
COPY . $GOPATH/src/github.com/ZJU-DistributedAI/Common
RUN ls -la $GOPATH/src/github.com/ZJU-DistributedAI/Common/public
WORKDIR $GOPATH/src/github.com/ZJU-DistributedAI/Common
RUN go get -d -v
RUN go build -o /go/bin/Common
# Define entrypoint
ENTRYPOINT ["/go/bin/Common"]