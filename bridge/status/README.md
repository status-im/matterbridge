# Description

This is an implementation of support for the [Status](https://status.im/) chat protocol.

It mostly makes use of the [status-go](https://github.com/status-im/status-go) and [status-protocol-go](https://github.com/status-im/status-protocol-go) packages.

# TODO

* Drop usage of the SQLite database entirely
* Improve handling of logs for the Whisper node
* Handle properly verifying successful delivery

# Known Issues

* Sometimes the message doesn't show up in Status Desktop, just in notifications
