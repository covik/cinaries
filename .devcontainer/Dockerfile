# syntax=docker/dockerfile:1@sha256:dabfc0969b935b2080555ace70ee69a5261af8a8f1b4df97b9e7fbcf6722eddf

FROM ghcr.io/project-zot/zot:v2.1.8@sha256:cd2aea942f428630bcb4190542be6abd35e14177aab84fc7ccad0dca8ecb363d AS zot
FROM golang:1.25.1-trixie@sha256:c8c8d55e02aa456e99a4fd0d2a224c0fb487d609c86a2550e8d64c3021b1df9a AS golang

FROM node:24.8.0-trixie-slim@sha256:a7f91166324d17fa1029718c876f3b02fb542e64d3cf40208ae6d31ff33c90c9 AS base
ENV TERM="xterm" \
    COREPACK_ENABLE_DOWNLOAD_PROMPT=0 \
    DO_NOT_TRACK=1 \
    TURBO_TELEMETRY_DISABLED=1
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get --no-install-recommends install -y \
        ca-certificates \
        curl \
        git \
        gnupg2 \
        less \
        jq \
        openssh-client \
        skopeo \
        sudo \
        tzdata \
        umoci \
        vim \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean \
    && corepack enable
COPY --from=zot /usr/local/bin/zot-* /usr/local/bin/zot
COPY --from=golang /usr/local/go /usr/local/go
ENV GOTOOLCHAIN=local \
    CGO_ENABLED=0 \
    GOROOT=/usr/local/go
ENV PATH="$GOROOT/bin:$PATH"

FROM base AS dev
ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ENV GOPATH=/home/$USERNAME/go
ENV PATH="$GOPATH/bin:$PATH"
RUN deluser --remove-home node \
    && groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID --create-home $USERNAME --shell /bin/bash \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && mkdir -p /home/$USERNAME/.cache/JetBrains/RemoteDev/dist /home/$USERNAME/.codex $GOPATH/pkg \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME \
    && rm -f /usr/local/bin/{npm,yarn}
USER $USERNAME
RUN git config --global commit.gpgsign true
