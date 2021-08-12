# Owncast Demo-Boot

This is a simple bot supposed to help people visiting the [Owncast Demo Instance](https://demo.owncast.online). It's also here, to show how easy it is to build on top of Owncast. This bot is written in [Go](https://golang.org/).

## Features

- receive Webhooks from Owncast with catch-all webhook endpoint
- send random chat messages on Stream Start / Stream Stop
- greet newly joined users
- give new users hints to change their name
- rudimentary question detection, causing the bot to reply with a handful of Owncast links
- support for !bang commands / custom commands
- running as container

## Configuration

This bot can be configured via the `config.yaml` configuration file and/or Environment variables. Environment variables take higher precedence and will override any values configured in the configuration.

### Configuration file

- ListenAddress
    Hostname/Port the bot is listening for incoming HTTP Requests from Owncast. The expected format is `hostname:port`.
    If you don't care to set the hostname, just set the value to the port like so: `:port` (e.g. `:8080`)

- AccessToken
    required to respond to chat-messages (you can create a new AccessToken in the Admin UI of your owncast instance)

- OwncastAddress
    the address to the Owncast instance that the bot should be running on (e.g. `http://demo.owncast.online`)
    make sure to include the scheme (`http` or `https`)


### Environment variables

| Name                                | Description                                                                       | Example Value |
|:------------------------------------|:----------------------------------------------------------------------------------|:--------------|
| OWNCAST_DEMOBOT_LISTENADDRESS       | Listen Address the Bot is listening on, for incoming HTTP Webhooks from Owncast   | `:8100`       |
| OWNCAST_DEMOBOT_ACCESSTOKEN         | Access token allowing to post to `/api/integrations/chat/system`                  | `123-asd-234` |
| OWNCAST_DEMOBOT_OWNCASTADDRESS      | To which Owncast instance this bot should reply to                                | `https://https://watch.owncast.online `

Environment variables WILL override values in the `config.yaml`. Environment variables are supposed to be used for configuring the bot, when run as Docker container.

## How to run this bot?

### Build and run from Source Code as executable

1. `git clone https://github.com/owncast/owncast-examples`
1. `cd owncast-examples/owncast-demobot`
1. edit the `config.yaml` to match your needs
1. `go run *.go`

### Build and run from Source Code as container

1. `git clone https://github.com/owncast/owncast-examples`
1. `cd owncast-examples/owncast-demobot`
1. `docker build -t owncastdemobot .`
1. `docker run --name owncastdemobot --network host -e OWNCAST_DEMOBOT_LISTENADDRESS=":8100" -e OWNCAST_DEMOBOT_ACCESSTOKEN="<our AccessToken> -e OWNCAST_DEMOBOT_OWNCASTADDRESS="http://localhost:8080" owncastdemobot`
