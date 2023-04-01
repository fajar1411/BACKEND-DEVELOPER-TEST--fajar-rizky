FROM golang:1.20-alpine as build

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

EXPOSE 8080

# create executable
RUN go build -o toko

# menjalankan file executablenya
CMD ["./toko"]
