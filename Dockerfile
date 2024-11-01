FROM golang:1.22.5-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add git gcc g++ make 


#dependencies
COPY ["go.mod", "go.sum", "./"]

RUN go mod download

#copy

COPY cmd ./cmd
COPY configs ./configs
COPY pkg ./pkg
COPY app-models ./app-models
COPY schema ./schema
COPY .env ./

#build 

RUN go build -o ./bin/app ./cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY --from=builder /usr/local/src/configs/config.yaml configs/config.yaml
COPY --from=builder /usr/local/src/schema /schema
COPY --from=builder /usr/local/src/.env /

CMD ["/app"]