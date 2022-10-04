# builder image
FROM golang:1.18 as builder

COPY /build/src /build
WORKDIR /build

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o maximilian .


FROM ghcr.io/unb-libraries/base:2.x
MAINTAINER UNB Libraries <libsupport@unb.ca>

COPY --from=builder /build/maximilian /app/maximilian
COPY ./build /build

ENTRYPOINT /app/maximilian

# Container metadata.
LABEL ca.unb.lib.generator="go" \
  com.microscaling.docker.dockerfile="/Dockerfile" \
  com.microscaling.license="MIT" \
  org.label-schema.build-date=$BUILD_DATE \
  org.label-schema.description="maximilian.lib.unb.ca is a Slack ChatOps app to interact with the Kubernetes cluster at UNB Libraries." \
  org.label-schema.name="maximilian.lib.unb.ca" \
  org.label-schema.schema-version="1.0" \
  org.label-schema.url="https://maximilian.lib.unb.ca" \
  org.label-schema.vcs-ref=$VCS_REF \
  org.label-schema.vcs-url="https://github.com/unb-libraries/maximilian.lib.unb.ca" \
  org.label-schema.vendor="University of New Brunswick Libraries" \
  org.label-schema.version=$VERSION \
  org.opencontainers.image.source="https://github.com/unb-libraries/maximilian.lib.unb.ca"
