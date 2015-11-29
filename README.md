PiCloud
====================

A websocket server for simple Raspberry Pi pub/sub.



Build and deploy with Fabric
====================

```
fab --user=username --hosts=remote_hostname build_deploy
```



Environment Variables
====================

* `PORT` - Port for server. Defaults to `9000`
* `API_KEY` - API Key required for authenticated requests.
* `USE_TLS` - Set this variable to enable TLS
* `KEY_FILE` - Path to TLS key file
* `CERT_FILE` - Path to TLS cert file
* `DEPLOY_USER` - User for SSH connection (Fabric deployment only)
* `DEPLOY_HOST` - Host for SSH connection (Fabric deployment only)



Websocket
====================

### Websocket subscription route

This route requires authentication with an API key passed in via the `X-API-Key` header or `apiKey` query parameter. Additionally, each connected client must send a query parameter `clientName`.

```
/subscribe?clientName=Something
```

#### Subscribe to an event

Send the following JSON message to **subscribe** to an event. After subscribing, your connection will receive data from any other connection that publishes that event.

```json
{
    "action": "subscribe",
    "event": "whatever"
}
```


### Websocket info route

This route requires authentication with an API key passed in via the `X-API-Key` header or `apiKey` query parameter. This socket does not handle any incoming messages. It pushes a JSON message every 5 seconds with info about the clients connected to the server. This is used to run the [picloud-info-web](https://github.com/exitcodezero/picloud-info-web) app.

```
/socket/info
```

```json
{
    "subscriptions": [
        {
            "name": "test",
            "connections": [        
                {   
                    "ip_address": "192.168.99.1:62178",     
                    "connected_at": "2015-11-25T03:40:13.433493056Z"
                }
            ]
        }
    ],
    "all_connections": [
        {
            "ip_address": "192.168.99.1:62178",
            "connected_at": "2015-11-25T03:40:13.433493056Z"
        },
        {
            "ip_address":"192.168.99.1:62184",
            "connected_at":"2015-11-25T03:40:31.147654724Z"
        }
    ],
    "created_at":"2015-11-25T03:40:39.889395863Z"
}
```



HTTP Routes
====================

### Publish data

Each request must send a query parameter `clientName` or a header `X-API-Client-Name`.

**POST:**
```
/publish?clientName=Something
```

**Body:**
```json
{
    "event": "whatever",
    "data": "howdy"
}
```

**Response:** None


**Status Codes:**
* `201` if successful
* `400` if incorrect data provided
* `401` if invalid `X-API-Key`
