FROM golang:latest as build

WORKDIR /go/src/app
ENV GOBIN=/go/bin
ENV GO111MODULE=on
COPY . .

RUN make build
RUN make install

FROM gcr.io/distroless/base
EXPOSE 3000
COPY --from=build /go/bin /
CMD ["/wasataro"]
