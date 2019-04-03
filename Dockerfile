FROM golang
WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM scratch

COPY --from=0 /build/server /opt/bin/server
ENTRYPOINT ["opt/bin/server"]