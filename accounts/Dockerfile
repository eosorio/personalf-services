# For debugging
FROM alpine:3.10.3
WORKDIR /
#COPY --from=build-env /server /
#CMD ["/server"]

# Go part
FROM golang:1.13
LABEL maintainer="eduardo.osorio.it@gmail.com"
EXPOSE 11111
#WORKDIR /go/src/app
ADD / /go/src/

RUN go build -o /accountsService accountsService.go

CMD ["/accountsService"]

