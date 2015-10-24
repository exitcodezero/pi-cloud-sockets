picloud
====================

A websocket server for simple Raspberry Pi pub/sub.

Environment Variables
====================

* `PORT` - Port for server. Defaults to `9000`
* `API_KEY` - API Key used for authenticated requests. Passed via `X-API-Key` header.
* `USE_TSL` - Set this variable to enable TSL
* `KEY_FILE` - Path to TSL key file
* `CERT_FILE` - Path to TSL cert file


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
