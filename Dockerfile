FROM golang:1.17-stretch

# dlv for debugging
RUN go get github.com/go-delve/delve/cmd/dlv
# reflex for watching source files and re-running the app
RUN go install github.com/cespare/reflex@latest

# rsync makes life easier
RUN apt-get install rsync grsync

WORKDIR /work
COPY ./go.* /work/
RUN go mod download

COPY ./dev.sh /work/

CMD [ "./dev.sh" ]

