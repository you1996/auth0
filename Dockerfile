FROM golang:1.12

RUN mkdir app
ADD . app
WORKDIR app

RUN go build -o main .
CMD ["./main"]
EXPOSE 3000