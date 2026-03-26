FROM mcr.microsoft.com/devcontainers/go:1.26@sha256:39c9b6eec0229417a7a27f24a686c4f226c92d42519b9845c3e571188f88e2a4
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get --no-install-recommends install -y \
        ca-certificates \
        jq \
        skopeo \
        tzdata \
        umoci \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean
