FROM golang:1.19

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN go build

CMD [ "./fridgesim" ]
