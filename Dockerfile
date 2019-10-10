#builder container
FROM golang:latest as builder
LABEL maintainer="Khan Sadirac <khan.sadirac42@gmail.com"
WORKDIR /app
COPY . .
RUN go build -o observatory_exporter

# main container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 9230
CMD ["./observatory_exporter"]
