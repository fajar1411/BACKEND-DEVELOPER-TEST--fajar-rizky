FROM golang:1.20-alpine 

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o toko

CMD ["./toko"]
