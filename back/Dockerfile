FROM golang:1.23-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /
COPY . .
RUN ls

RUN go mod download
ENV GO_ENV production
RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.8.4 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build_base /bootstrap /bootstrap
COPY --from=build_base /env.production /env.production
ENV PORT=8080 AWS_LWA_ASYNC_INIT=true
RUN ls

EXPOSE 8080

CMD ["/bootstrap"]