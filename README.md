picloud
====================

A websocket server for simple Raspberry Pi pub/sub.



Build and deploy with Fabric
====================

`Fabfile.py` is located in `/deploy`

```
fab build_deploy
```



Environment Variables
====================

* `PORT` - Port for server. Defaults to `9000`
* `API_KEY` - API Key used for authenticated requests. Passed via `X-API-Key` header.
* `USE_TLS` - Set this variable to enable TLS
* `KEY_FILE` - Path to TLS key file
* `CERT_FILE` - Path to TLS cert file
* `DEPLOY_USER` - User for SSH connection (Fabric deployment only)
* `DEPLOY_HOST` - Host for SSH connection (Fabric deployment only)



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
