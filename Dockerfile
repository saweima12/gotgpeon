FROM golang:1.20 as build

WORKDIR /data
# COPY project into build image.
COPY . .

# Download depedencies.
RUN go mod download

RUN git clone https://github.com/liuzl/gocc.git

# Build project
RUN CGO_ENABLED=0 go build -o ./app/ ./cmd/...

# Build finish, Copy to runtime
from alpine as runtime
# FROM gcr.io/distroless/static-debian12 as runtime
WORKDIR /app

COPY --from=build /data/app ./
COPY --from=build /data/config.yml ./
COPY --from=build /data/lang.yml ./
COPY --from=build /data/gocc/config ./
COPY --from=build /data/gocc/dictionary ./
