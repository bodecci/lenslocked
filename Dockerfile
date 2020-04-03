FROM golang:1.13.1-alpine3.10

WORKDIR /bbfalade/go/src/github.com/bodecci/use_golang/lenslocked

EXPOSE 8080

COPY ./lenslocked .
CMD ["./lenslocked/main"]