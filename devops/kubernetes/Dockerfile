FROM golang:1.17.1-alpine AS builder

ADD . /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

RUN go build -o server .

FROM nginx:alpine
COPY --from=builder ./app .
COPY nginx.conf /etc/nginx/nginx.conf
CMD ["./server"]