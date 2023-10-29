# c3pcid - Discord Engagement Bot

![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)
![Golang Version](https://img.shields.io/badge/golang-1.16-green.svg)

![c3pcid](https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/3ef5ebfb-6e80-43e9-970a-5370301b3278/de5jgee-ab550100-8e1c-4de1-b043-bdbe4f251fb9.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzNlZjVlYmZiLTZlODAtNDNlOS05NzBhLTUzNzAzMDFiMzI3OFwvZGU1amdlZS1hYjU1MDEwMC04ZTFjLTRkZTEtYjA0My1iZGJlNGYyNTFmYjkucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.g5Biha_1n1WxTRVtnWg24bt7k2x6l9fFJqZJR5SIopU)

**c3pcid** is a Discord bot developed to foster engagement within Discord servers. It allows for the collection, storage, and customization of channel-specific information, tailored to the server's preferences. This information can be utilized to create interactive and engaging content among server members, such as displaying funny phrases from a streamer during a live stream, on a website, or even in an Alexa skill.

## Features

- Collects and stores information from Discord channels.
- Customization using RegEx or emoji reaction types.
- Filters and stores specific information based on defined criteria.
- Integration with Twitch, websites, or other applications.

## Prerequisites

Before using **c3pcid**, you will need to set up a few things:

1. **Create a Discord Bot**: Follow the official Discord guide on creating a bot - [Creating a Bot Account](https://discordpy.readthedocs.io/en/stable/discord.html).

2. **Obtain a Bot Token**: After creating the bot, obtain the access token to authenticate the bot on Discord.

3. **Channel IDs**: Identify the Discord channels you want the bot to monitor and collect information from.

## Configuration

1. Clone the repository:

   ```
   git clone https://github.com/vlsouza/c3pcid.git
   cd c3pcid
   ```

2. Set up environment variables:

   Create a `.env` file in the project's root directory with the following variables:

   ```
   BOT_TOKEN=value
   your_channel_REGEX=value
   your_channel_CHANNEL_NAME=value
   your_channel_CHANNEL_ID=value
   MESSAGE_AMOUNT_TO_GET=value
   MANY_TIMES_TO_GET_MESSAGE=value
   ```

   Replace `BOT_TOKEN` with your bot's token and list the channel IDs you wish to monitor, separated by commas.

3. Run the bot:

   ```
   go run main.go
   ```

   The bot is now active and collecting information from the configured channels.

## Customization

To customize the information collection and filtering criteria, you can modify the source code in `main.go` according to your specific needs.

## License

This project is distributed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it according to the terms of the license.

## Contributions

Contributions are welcome! Please feel free to open issues or submit pull requests to improve the project.
