FROM golang:1.11
LABEL MAINTAINER Gus Luxton <gus@gravitational.com>

ENV GOPATH /go
RUN mkdir -p /go && \
    go get -d -v github.com/webvictim/eventing-sources/cmd/event_display && \
    go install -v github.com/webvictim/eventing-sources/cmd/event_display

RUN command -v event_display

CMD ["event_display"]