module github.com/42wim/matterbridge

require (
	github.com/42wim/go-gitter v0.0.0-20170828205020-017310c2d557
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f
	github.com/Jeffail/gabs v1.1.1 // indirect
	github.com/Philipp15b/go-steam v1.0.1-0.20190816133340-b04c5a83c1c0
	github.com/Rhymen/go-whatsapp v0.1.0
	github.com/d5/tengo/v2 v2.0.2
	github.com/dfordsoft/golib v0.0.0-20180902042739-76ee6ab99bec
	github.com/ethereum/go-ethereum v1.9.5
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.5-0.20181225215658-ec221ba9ea45+incompatible
	github.com/gomarkdown/markdown v0.0.0-20200127000047-1813ea067497
	github.com/google/gops v0.3.6
	github.com/google/uuid v1.1.1
	github.com/gopackage/ddp v0.0.0-20170117053602-652027933df4 // indirect
	github.com/gorilla/schema v1.1.0
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/golang-lru v0.5.3
	github.com/jpillora/backoff v1.0.0
	github.com/keybase/go-keybase-chat-bot v0.0.0-20200226211841-4e48f3eaef3e
	github.com/labstack/echo/v4 v4.1.13
	github.com/lrstanley/girc v0.0.0-20190801035559-4fc93959e1a7
	github.com/matterbridge/Rocket.Chat.Go.SDK v0.0.0-20190210153444-cc9d05784d5d
	github.com/matterbridge/discordgo v0.18.1-0.20200308151012-aa40f01cbcc3
	github.com/matterbridge/emoji v2.1.1-0.20191117213217-af507f6b02db+incompatible
	github.com/matterbridge/go-xmpp v0.0.0-20180529212104-cd19799fba91
	github.com/matterbridge/gomatrix v0.0.0-20200209224845-c2104d7936a6
	github.com/matterbridge/gozulipbot v0.0.0-20190212232658-7aa251978a18
	github.com/matterbridge/logrus-prefixed-formatter v0.0.0-20180806162718-01618749af61
	github.com/matterbridge/msgraph.go v0.0.0-20200308150230-9e043fe9dbaa
	github.com/mattermost/mattermost-server v5.5.0+incompatible
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/mattn/godown v0.0.0-20180312012330-2e9e17e0ea51
	github.com/mreiferson/go-httpclient v0.0.0-20160630210159-31f0106b4474 // indirect
	github.com/mrexodia/wray v0.0.0-20160318003008-78a2c1f284ff // indirect
	github.com/nelsonken/gomf v0.0.0-20180504123937-a9dd2f9deae9
	github.com/nicksnyder/go-i18n v1.4.0 // indirect
	github.com/paulrosania/go-charset v0.0.0-20190326053356-55c9d7a5834c
	github.com/pkg/errors v0.9.1
	github.com/rs/xid v1.2.1
	github.com/russross/blackfriday v1.5.2
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca
	github.com/shazow/ssh-chat v1.8.3-0.20200308224626-80ddf1f43a98
	github.com/sirupsen/logrus v1.4.2
	github.com/slack-go/slack v0.6.3-0.20200228121756-f56d616d5901
	github.com/spf13/viper v1.6.1
	github.com/status-im/status-go v0.49.0
	github.com/stretchr/testify v1.5.1
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/zfjagann/golang-ring v0.0.0-20190106091943-a88bb6aef447
	go.uber.org/zap v1.13.0
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace github.com/nlopes/slack v0.6.0 => github.com/matterbridge/slack v0.1.1-0.20191208194820-95190f11bfb6

replace github.com/bwmarrin/discordgo v0.19.0 => github.com/matterbridge/discordgo v0.0.0-20191026232317-01823f4ebba4

replace github.com/ethereum/go-ethereum v1.9.5 => github.com/status-im/go-ethereum v1.9.5-status.7

go 1.13
