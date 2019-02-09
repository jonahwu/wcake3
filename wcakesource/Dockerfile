#FROM golang:latest AS build
FROM golang:1.11 AS build

ENV GOARCH_SRC=$GOPATH/src/github.com/wcake
#ENV CGO_ENABLED=1
#ENV GOOS=linux
#ENV NOMS_VERSION_NEXT=1
#ENV DOCKER=1

RUN mkdir -pv $GOARCH_SRC
COPY . ${GOARCH_SRC}
RUN go test github.com/wcake/...
RUN ls $GOPATH/src/github.com/wcake/cmd/wcake -alh
RUN cd $GOPATH/bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build github.com/wcake/cmd/wcake
#RUN go install -v 
RUN cp $GOPATH/bin/wcake /bin/wcake
RUN ls $GOPATH/bin/ -alh
RUN ls /bin/ -alh

FROM alpine:latest

COPY --from=build /bin/wcake /wcake
RUN ls / -alh
#VOLUME /data
EXPOSE 8000

ENV NOMS_VERSION_NEXT=1
RUN chmod +x ./wcake
ENTRYPOINT [ "./wcake" ]

#CMD ["serve", "/data"] ]
