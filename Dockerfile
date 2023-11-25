FROM golang:1.20 as builder
LABEL stage=builder

WORKDIR /data
# COPY project into build image.
COPY . .

# Download depedencies.
RUN go mod download

RUN git clone https://github.com/liuzl/gocc.git

# Build project
RUN CGO_ENABLED=0 go build -o ./app/ ./cmd/...

# Build finish, Copy to runtime
# FROM alpine as runtime
FROM gcr.io/distroless/static-debian12 as runtime
LABEL stage=runtime


WORKDIR /app

COPY --from=builder /data/app ./
COPY --from=builder /data/config.yml ./
COPY --from=builder /data/lang.yml ./
COPY --from=builder /data/gocc/config ./config
COPY --from=builder /data/gocc/dictionary ./dictionary
