FROM golang:1.20.5-alpine AS build
RUN mkdir /src
ADD ./*.go /src
ADD ./go.mod /src
ADD ./go.sum /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux GOARCH=amd64 go build -v -o pinger 
RUN chmod +x pinger

FROM scratch
ENV DB_PORT=5432 DB_USER=postgres DB_NAME=pinger
COPY --from=build /src/pinger /usr/local/bin/pinger
EXPOSE 8080
CMD ["pinger"]
