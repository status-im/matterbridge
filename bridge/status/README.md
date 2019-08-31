# Description

This is an implementation of support for the [Status](https://status.im/) chat protocol.

It makes use of the [status-go](https://github.com/status-im/status-go).

# Configuration

Here's a basic config using a randomly generated private key(`Token`) for bridge identity:
```yaml
---
Nick: "bridge-bot"

status:
  bridge:
    Nick: 'mybridge.stateofus.eth'
    Token: '0xeb87e5780fef3a83fa6a7f5c19fb206715a66e1c10aab50471686c6347b1ede4'
    RemoteNickFormat: '**{NICK}**@*{PROTOCOL}*: '

gateway:
  - name: "status-bridge-test"
    enable: true
    inout:
      - account: "status.bridge"
        channel: "test-channel-1"

      - account: "status.bridge"
        channel: "test-channel-2"
```

# TODO

* Drop usage of the SQLite database entirely
* Improve handling of logs for the Whisper node
* Handle properly verifying successful delivery
