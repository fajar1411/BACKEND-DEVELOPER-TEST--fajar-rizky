FROM golang:1.20-alpine  as build

# membuat direktori app
RUN mkdir /olshop

# set working directory /app
WORKDIR /olshop

COPY ./ /olshop

RUN go mod tidy

RUN go build -o olshop

CMD ["./olshop"]
