FROM alpine:edge AS builder

COPY . /go/src/github.com/42wim/matterbridge
RUN apk update && apk add go git gcc musl-dev \
        && cd /go/src/github.com/42wim/matterbridge \
        && export GOPATH=/go \
        && go get \
        && go build -x -ldflags "-X main.githash=$(git log --pretty=format:'%h' -n 1)" -o /bin/matterbridge

FROM alpine:edge
RUN apk --no-cache add ca-certificates mailcap
COPY --from=builder /bin/matterbridge /bin/matterbridge
ENTRYPOINT ["/bin/matterbridge"]
