FROM golang:1.25-alpine3.22 AS build

WORKDIR /wtf
COPY . .
RUN go build -v -mod=vendor -trimpath -ldflags '-d -w' -o /usr/local/bin/uscan-helper

FROM alpine:3.22

RUN apk add --no-cache git

COPY --from=build /usr/local/bin/uscan-helper /usr/local/bin/uscan-helper

EXPOSE 8181
CMD ["uscan-helper"]
