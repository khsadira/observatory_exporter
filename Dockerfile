#builder container
FROM golang:latest as builder
LABEL maintainer="Khan Sadirac <khan.sadirac42@gmail.com"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o observatory_exporter .

# main container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 9230
CMD ["./observatory_exporter"]