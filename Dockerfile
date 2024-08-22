FROM scratch
COPY bin/linux_amd64/parser /
ENTRYPOINT ["/parser"]
