//go:build !nosignal
// +build !nosignal

package bridgemap

import (
	bsignal "github.com/42wim/matterbridge/bridge/signal" //nolint:depguard
)

func init() { //nolint:gochecknoinits
	FullMap["signal"] = bsignal.New
}
