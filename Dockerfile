# https://github.com/GoogleContainerTools/distroless
FROM golang:1.19 AS builder
WORKDIR /go/src/app
COPY . .
RUN make setup
RUN make docker

FROM gcr.io/distroless/static-debian11
WORKDIR /root/
COPY --from=builder /go/src/app/bin/trokhos .
CMD [ "./trokhos" ]