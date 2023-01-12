# LinkScreen Repository

## Description
The LinkScreen repository contains a golang project that implements a telegram bot that sends a page image in response to a link sent by a user.
The purpose of this bot is to allow the user to get a visual representation of the site, the link of which seems suspicious.

<img src="/repositoryDescription/images/usage_example_en.jpg" width="250" height="550">

## Notes
* The program needs to read a special telegram bot token. In my implementation, the token is stored as an operating system environment variable.
The name of variable is `LINK_SCREEN_TGBOT_TOKEN`. You can get your own token in the official way, which is described in detail in the [telegram documentation](https://core.telegram.org/bots/api).

## Used packages
* [Viper](https://github.com/spf13/viper)
* [Golang bindings for the Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)
