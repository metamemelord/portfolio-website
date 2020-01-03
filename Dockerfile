FROM golang:alpine as builder
RUN apk add --update nodejs nodejs-npm build-base git
WORKDIR /build
COPY . .
RUN npm install
RUN npm run build
RUN go build -o portfolio

FROM alpine
WORKDIR /portfolio
COPY --from=builder /build/dist ./dist
COPY --from=builder /build/portfolio .
CMD ["./portfolio"]