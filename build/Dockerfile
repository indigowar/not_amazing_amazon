FROM golang:1.23.2 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o not_amazing_amazon ./cmd/not_amazing_amazon

FROM scratch
WORKDIR /
COPY --from=build /app/not_amazing_amazon /bin/not_amazing_amazon
EXPOSE 80
CMD [ "/bin/not_amazing_amazon" ]
