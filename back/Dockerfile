FROM golang:1.23-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /

COPY . .
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.8.4 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build_base /bootstrap /bootstrap

ENV PORT=8000 AWS_LWA_ASYNC_INIT=true
EXPOSE 8000

CMD ["/bootstrap"]