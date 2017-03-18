A simple cross platform bot that leverages Selly's webhook system for discord chat notifications

#### Setup

Create a discord bot account https://discordapp.com/developers/applications/me and then grab the `Client ID` and then go to the following URL to invite the bot to the Discord server.

```
https://discordapp.com/api/oauth2/authorize?client_id=CLIENT_ID_GOES_HEREscope=bot
```

Enable developer mode on the Discord client and right click on the channel you want the bot to send the message to and click `Copy ID`

Open `config.json` and add the bot's token found on the same page where you found the `Client ID`.

Adjust the `port` (we recommend `80` if possible) and `webhook secret` to your desired values.

Now go to [Selly](https://selly.gg/products) and edit the products you'd like the bot to notify regarding orders. An example webhook URL with a port of `80` and a secret of `piesarenice` would be:

```
http://ip.of.server/webhooks?secret=piesarenice
```

If the port is not `80`, but in this case `123`, it would be:

```
http://ip.of.server:123/webhooks?secret=piesarenice
```

#### Running
Once the setup is completed, you can now run the bot.

If you have Go installed (latest is always recommended), you can run:
```
go build
./executable
/// OR if windows
executable.exe
```
**Make sure** the executable is in the same folder as `config.json`

If you don't have Go installed, you can run one of the precompiled executable/binarys found in the `rel
eases` tab.