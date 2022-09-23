FROM golang:1.19

WORKDIR /work

# download deps
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
COPY . .
RUN go build

CMD [ "./fridgesim" ]
