FROM golang:1.20 as builder

WORKDIR /go/src/pinterest

COPY . .


RUN go mod tidy && \
	go mod vendor && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/pinterest .

# 빌드 시 _USER_API_TOKEN이 layer에 포함되므로
# 반드시 빌드 결과물을 추출해서 새로운 이미지로 빌드
# FROM gcr.io/distroless/static-debian11
FROM scratch

WORKDIR /go/src/pinterest

COPY --from=builder /go/src/pinterest/bin/pinterest ./pinterest

ENTRYPOINT ["./pinterest"]
