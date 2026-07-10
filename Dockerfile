# builder image
FROM golang:1.18 as builder

COPY /build/src /build
WORKDIR /build

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o maximilian .


FROM ghcr.io/unb-libraries/base:2.x

COPY --from=builder /build/maximilian /app/maximilian
COPY ./build /build

ENTRYPOINT /app/maximilian

# Container metadata.
LABEL ca.unb.lib.generator="go" \
  org.opencontainers.image.title="maximilian.lib.unb.ca" \
  org.opencontainers.image.description="maximilian.lib.unb.ca is a Slack ChatOps app to interact with the Kubernetes cluster at UNB Libraries." \
  org.opencontainers.image.vendor="University of New Brunswick Libraries" \
  org.opencontainers.image.authors="UNB Libraries <libsupport@unb.ca>" \
  org.opencontainers.image.url="https://maximilian.lib.unb.ca" \
  org.opencontainers.image.source="https://github.com/unb-libraries/maximilian.lib.unb.ca" \
  org.opencontainers.image.version="$VERSION" \
  org.opencontainers.image.revision="$VCS_REF" \
  org.opencontainers.image.created="$BUILD_DATE"
