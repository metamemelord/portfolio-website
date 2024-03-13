FROM docker.io/library/golang:1-alpine AS server-builder
RUN apk add upx
WORKDIR /build
ENV GO111MODULE=on
ENV PATH=$PATH:$GOPATH/bin
RUN go install github.com/dmarkham/enumer@latest
COPY . .
RUN go generate ./...
RUN go build -v -ldflags="-s -w" -o portfolio
RUN upx -9 -k portfolio
 
FROM docker.io/library/node:20-alpine AS ui-builder
ENV NODE_OPTIONS="--openssl-legacy-provider"
RUN npm install --location=global npm
WORKDIR /build
COPY . .
RUN npx browserslist@latest --update-db
RUN npm install
RUN npm run build
 
FROM docker.io/library/alpine AS pre-prod
COPY --from=server-builder /build/portfolio /bin
COPY --from=ui-builder /build/dist /srv/portfolio/dist
 
FROM scratch
COPY --from=pre-prod / /
CMD ["/bin/portfolio"]
