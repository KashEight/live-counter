FROM golang:1.16 AS build

WORKDIR /root/app/
COPY . /root/app/

RUN go build -o dist/

FROM gcr.io/distroless/base

WORKDIR /root/
COPY --from=build /root/app/dist .

EXPOSE 8080

CMD ["./live-counter"]
