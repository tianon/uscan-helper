FROM golang:1.12-alpine3.10

RUN apk add --no-cache git

WORKDIR /usr/src/uscan-helper
COPY . .
RUN go install -v -mod=vendor

EXPOSE 8181
CMD ["uscan-helper"]
