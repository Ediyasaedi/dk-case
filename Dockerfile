FROM golang

LABEL maintainer="ediyasaedi@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

ENV JWT_SECRET "secretforjsonwebtoken"

RUN go build

CMD ["./dk-case"]