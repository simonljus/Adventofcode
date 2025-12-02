FROM denoland/deno:bin-2.5.6 AS deno
FROM mcr.microsoft.com/devcontainers/typescript-node:24
COPY --from=deno /deno /usr/local/bin/deno