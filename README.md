picloud
====================

A websocket server for simple Raspberry Pi pub/sub.

Compile with Docker
====================

```
docker run --rm -it -v $PWD/app/:/go/src/app -w /go/src/app golang:1.5 /bin/bash
```

```
go get && go build -v
```

The new binary named `app` can be found in the `/app` directory.


Environment Variables
====================

* `PORT` - Port for server. Defaults to `9000`
* `API_KEY` - API Key used for authenticated requests. Passed via `X-API-Key` header.
* `USE_TLS` - Set this variable to enable TLS
* `KEY_FILE` - Path to TLS key file
* `CERT_FILE` - Path to TLS cert file


Messages
====================

### Subscribe to an event

Send the following JSON message to **subscribe** to an event. After subscribing, your connection will receive data from any other connection that publishes that event.

```json
{
    "action": "subscribe",
    "event": "whatever"
}
```

### Publish data for an event

Send the following JSON message to **publish** data for an event. Any connection that has subscribed to this event will receive this JSON message.

```json
{
    "action": "publish",
    "event": "whatever",
    "data": "howdy"
}
```
