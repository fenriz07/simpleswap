#BaseImage arg
FROM ignitehq/cli:v28.3.0 as builder
USER root
COPY . /app
WORKDIR /app
ENTRYPOINT ["ignite", "chain", "serve"]