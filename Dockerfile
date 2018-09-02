FROM golang:1.9 as base
WORKDIR /usr/src
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM scratch
COPY --from=base /usr/src/main /geekbread-api
CMD ["/geekbread-api"]
