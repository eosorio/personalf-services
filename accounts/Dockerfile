# For debugging
#FROM alpine:3.10.3
#WORKDIR /
#COPY --from=build-env /server /
#CMD ["/server"]

# Go part
FROM golang:1.13
LABEL maintainer="eduardo.osorio.it@gmail.com"
EXPOSE 11111
WORKDIR /go/src/git.osmon.local/personalf-services
COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...

RUN go build -o /accountsService -i /go/src/git.osmon.local/personalf-services/accountsService.go
CMD ["/accountsService"]

