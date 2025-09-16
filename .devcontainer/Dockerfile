# syntax=docker/dockerfile:1@sha256:dabfc0969b935b2080555ace70ee69a5261af8a8f1b4df97b9e7fbcf6722eddf

FROM gcr.io/go-containerregistry/crane:latest@sha256:13952b31a107e8bc4e4c170b0097edb708cd468486a28fc20bbf0d49f71c74e1 AS crane
FROM ghcr.io/project-zot/zot:v2.1.8@sha256:cd2aea942f428630bcb4190542be6abd35e14177aab84fc7ccad0dca8ecb363d AS zot

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
        umoci \
        vim \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean \
    && corepack enable
COPY --from=crane /ko-app/crane /usr/local/bin/crane
COPY --from=zot /usr/local/bin/zot-* /usr/local/bin/zot

FROM base AS dev
ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN deluser --remove-home node \
    && groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID --create-home $USERNAME --shell /bin/bash \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && mkdir -p /home/$USERNAME/.cache/JetBrains/RemoteDev/dist /home/$USERNAME/.codex \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME \
    && rm -f /usr/local/bin/{npm,yarn}
USER $USERNAME
RUN git config --global commit.gpgsign true
