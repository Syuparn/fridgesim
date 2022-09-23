FROM golang:1.19

WORKDIR /work

COPY go.* .
RUN go mod tidy

COPY . .
RUN go build

CMD [ "./fridgesim" ]
