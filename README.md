[![Build Status][travis-image]][travis-url]
[![Github Tag][githubtag-image]][githubtag-url]

[![Coverage Status][coveralls-image]][coveralls-url]
[![Maintainability][codeclimate-image]][codeclimate-url]
[![codecov][codecov-image]][codecov-url]

[![Go Report Card][goreport-image]][goreport-url]
[![GoDoc][godoc-image]][godoc-url]
[![License][license-image]][license-url]

***

# go-teams-notify

> A package to send messages to Microsoft Teams (channels)

...

# Usage

To get the package, execute:

```
go get gopkg.in/dasrick/go-teams-notify.v1
```

To import this package, add the following line to your code:

```
import "gopkg.in/dasrick/go-teams-notify.v1"
```

And this is an example of a simple implementation ...

```
import (
	"gopkg.in/dasrick/go-teams-notify.v1"
)

func main() {
	_ = sendTheMessage()
}

func sendTheMessage() error {
	// init the client
	mstClient, err := goteamsnotify.NewClient()
	if err != nil {
		return err
	}
	// setup webhook url
	webhookUrl := "https://outlook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"
	// setup message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = "Hello world"
	msgCard.Text = "Here are some examples of formatted stuff like <br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
	msgCard.ThemeColor = "#DF813D"
	// firestarter
	return mstClient.Send(webhookUrl, msgCard)
}
```

# <a id="links"></a>some useful links

* [MS Teams - adaptive cards](https://docs.microsoft.com/de-de/outlook/actionable-messages/adaptive-card)
* [MS Teams - send via connectors](https://docs.microsoft.com/de-de/outlook/actionable-messages/send-via-connectors)
* [adaptivecards.io](https://adaptivecards.io/designer)

***

[travis-image]: https://travis-ci.org/dasrick/go-teams-notify.svg?branch=master
[travis-url]: https://travis-ci.org/dasrick/go-teams-notify

[githubtag-image]: https://img.shields.io/github/tag/dasrick/go-teams-notify.svg?style=flat
[githubtag-url]: https://github.com/dasrick/go-teams-notify

[coveralls-image]: https://coveralls.io/repos/github/dasrick/go-teams-notify/badge.svg?branch=master
[coveralls-url]: https://coveralls.io/github/dasrick/go-teams-notify?branch=master

[codeclimate-image]: https://api.codeclimate.com/v1/badges/fe69cc992370b3f97d94/maintainability
[codeclimate-url]: https://codeclimate.com/github/dasrick/go-teams-notify/maintainability

[codecov-image]: https://codecov.io/gh/dasrick/go-teams-notify/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/dasrick/go-teams-notify

[goreport-image]: https://goreportcard.com/badge/github.com/dasrick/go-teams-notify
[goreport-url]: https://goreportcard.com/report/github.com/dasrick/go-teams-notify

[godoc-image]: https://godoc.org/github.com/dasrick/go-teams-notify?status.svg
[godoc-url]: https://godoc.org/github.com/dasrick/go-teams-notify

[license-image]: https://img.shields.io/github/license/dasrick/go-teams-notify.svg?style=flat
[license-url]: https://github.com/dasrick/go-teams-notify/blob/master/LICENSE
