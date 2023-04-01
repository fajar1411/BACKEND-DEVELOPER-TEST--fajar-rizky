FROM golang:1.20-alpine  as build

# membuat direktori app
RUN mkdir /toko

# set working directory /app
WORKDIR /toko

COPY ./ /toko

RUN go mod tidy

RUN go build -o toko

CMD ["./toko"]
