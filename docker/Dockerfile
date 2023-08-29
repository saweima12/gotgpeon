FROM golang:1.20 as build

WORKDIR /data
# COPY project into build image.
COPY . .

# Download depedencies.
RUN go mod download

# Build project
RUN CGO_ENABLE=0 go build -o ./app/ ./cmd/...

# Build finish, Copy to runtime
FROM gcr.io/distroless/static-debian12 as runtime

COPY --from=build /data/app /
