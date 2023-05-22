# network-uci-bridge
Bridge to a network UCI engine for chess GUIs that support only a local binary<br/>

Chess GUI <- exec -> uci-bridge (binary) <- network connection -> UCI engine

## Usage

At the moment, compiled binaries are available from the prebuilt docker images.
```
ghcr.io/n0rthernl1ghts/uci-bridge:latest
```

Example usage within your own Dockerfile:
```Dockefile
FROM scratch AS rootfs

# Copy compiled binary from the image
COPY --from=ghcr.io/n0rthernl1ghts/uci-bridge:latest ["/usr/local/bin/uci-bridge", "/usr/local/bin/"]


FROM alpine:3.18 AS my-chess-gui

COPY --from=rootfs ["/", "/"]

(...)
```

CLI Usage:
```shell
echo "isready" | docker run --init --rm -e "UCI_TCP_HOST=192.168.1.20" -e "UCI_TCP_PORT=3333" -i ghcr.io/n0rthernl1ghts/uci-bridge:latest
```
If everything is working correctly, you should see `readyok` response from the engine.<br/>
The same applies to any other UCI command.

### Environment variables

Configuration is very simple and done via environment variables.
```dotenv
UCI_TCP_HOST=uci-engine.example.com
UCI_TCP_PORT=3333
```

When running outside docker environment, you can use `.env` file to set environment variables. <br/>
You can also export them to the environment or run command like this:
```shell
echo "isready" | UCI_TCP_HOST=uci-engine.example.com UCI_TCP_PORT=3333 uci-bridge
```

## Project status
This project is in early development stage. As is, it should be considered experimental. <br/>