FROM golang:alpine AS builderimage
WORKDIR /go/src/tftpd-go
COPY . .
RUN go build -o tftpd-go main.go


###################################################################

FROM alpine
COPY --from=builderimage /go/src/tftpd-go/tftpd-go /app/
WORKDIR /app
EXPOSE 69
CMD ["./tftpd-go"]