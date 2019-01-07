# BUILD STAGE
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/github.com/melodiez14/meiko
COPY . .
COPY ./files/var/www/meiko /var/www/meiko
COPY ./files/etc/meiko /etc/meiko
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app

# FINAL STAGE
FROM alpine
COPY --from=builder /var/www/meiko /var/www/meiko
COPY --from=builder /etc/meiko /etc/meiko
COPY --from=builder /go/bin/app /go/bin/app
RUN chmod 744 /go/bin/app
ENV LCENV=production
EXPOSE 9000
ENTRYPOINT [ "/go/bin/app" ]
