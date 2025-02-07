FROM golang:1.23.2

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
COPY ./handler ./handler
COPY ./model ./model
COPY ./service ./service
RUN go build -v -o /usr/local/bin/app ./

EXPOSE 8090

CMD ["app"]