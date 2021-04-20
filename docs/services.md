

### Configuration 

All the flags have crosponding enviornment variables assosiated with it. You can either provide the value with flags

or export to a variable. You can view the crosponding variable to each with --help flag. 

*Flags* take presedance over *variables*

*Default* value for message title is current *time*


## Telegram
Telegram uses bot token to authenticate & send messages to defined channels.
Multiple channel ids can be used separated by comma ','.

```bash
pingme  telegram  --token "0125:AAFHvnYf_ABC"  --msg "This is a new message ✈" --channel="-1001001001,-1002002001"
```

- Github Action

```yaml
on: [push]

jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          TELEGRAM_TOKEN: ${{ secrets.TELEGRAM_TOKEN }}
          TELEGRAM_CHANNELS: ${{ secrets.TELEGRAM_CHANNELS }}
          TELEGRAM_TITLE: 'Refrence: ${{ github.ref }}'
          TELEGRAM_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email
          service: telegram
```
- **Variables**


|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| TELEGRAM_MSG_TITLE        |         ""         |  
| TELEGRAM_TOKEN           |         ""        |
| TELEGRAM_CHANNELS   |         ""         | 
| TELEGRAM_MESSAGE            |         ""         | 
| TELEGRAM_MSG_TITLE        |         ""         |  


## RocketChat
RocketChat uses token & userID to authenticate and send messages to defined channels.
Multiple channel ids can be used separated by comma ','.

```bash
pingme rocketchat --channel "general,Pingme" --msg ":wave: rocketchat from cli" --userid "123" --token "abcxyz" \
--url 'localhost:3000' --scheme "http"
```

- Github Action

```yaml
on:
  release:
    types: [published]
jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      
      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          ROCKETCHAT_USERID: ${{ secrets.ROCKETCHAT_USERID }}
          ROCKETCHAT_TOKEN: ${{ secrets.ROCKETCHAT_TOKEN }}
          ROCKETCHAT_SERVER_URL: ${{ secrets.ROCKETCHAT_SERVER_URL }}
          ROCKETCHAT_CHANNELS: ${{ secrets.ROCKETCHAT_CHANNELS }}
          ROCKETCHAT_URL_SCHEME: "https"
          ROCKETCHAT_TITLE: 'Refrence: ${{ github.ref }}'
          ROCKETCHAT_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email / mattermost
          service: rocketchat
```
- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| ROCKETCHAT_USERID           |         ""        |
| ROCKETCHAT_TOKEN   |         ""         | 
| ROCKETCHAT_SERVER_URL            |         ""         | 
| ROCKETCHAT_URL_SCHEME        |         "https"         |  
| RTOCKETCHAT_MESSAGE            |         ""         | 
| ROCKETCHAT_TITLE            |         ""         | 
| ROCKETCHAT_CHANNELS        |         ""         |  


## Pushover

```bash
pingme pushover --token '123' --user '12345567' --title 'some title' --message 'some message'
```

- Github Action

```yaml
on: [push]

jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          PUSHOVER_TOKEN: ${{ secrets.PUSHOVER_TOKEN }}
          PUSHOVER_USER: ${{ secrets.PUSHOVER_USER }}
          PUSHOVER_TITLE: 'Refrence: ${{ github.ref }}'
          PUSHOVER_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email
          service: pushover
```

- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| PUSHOVER_TOKEN           |         ""        |
| PUSHOVER_USER   |         ""         | 
| PUSHOVER_MESSAGE            |         ""         |
| PUSHOVER_TITLE            |         ""         | 

## Mattermost
Mattermost uses token to authenticate and channel ids for targets.
Destination server can be specified as 'example.com' by default the 'https' is used, you
can change this with --scheme flag and set it to 'http'.
Latest api  version 4 is used for interacting with server, this can also be changes with --api flag.
You can specify multiple channels by separating the value with ','.

```bash
pingme mattermost --token '123' --channel '12345,567' --url 'localhost' --scheme 'http' --message 'some message'
```

- Github Action

```yaml
on:
  release:
    types: [published]
jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      
      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          MATTERMOST_TOKEN: ${{ secrets.MATTERMOST_TOKEN }}
          ROCKETCHAT_SERVER_URL: ${{ secrets.ROCKETCHAT_SERVER_URL }}
          MATTERMOST_CHANNELS: ${{ secrets.MATTERMOST_CHANNELS }}
          MATTERMOST_CHANNELS: ${{ secrets.MATTERMOST_CHANNELS }}
          MATTERMOST_TITLE: 'Refrence: ${{ github.ref }}'
          MATTERMOST_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email / mattermost
          service: mattermost
```

- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| MATTERMOST_API_URL           |         "/api/v4/posts"        |
| MATTERMOST_TOKEN   |         ""         | 
| MATTERMOST_SERVER_URL            |         ""         | 
| MATTERMOST_SCHEME        |         "https"         |  
| MATTERMOST_MESSAGE            |         ""         | 
| MATTERMOST_TITLE            |         ""         | 
| MATTERMOST_CHANNELS        |         ""         |  

## Slack
Slack uses token to authenticate and send messages to defined channels.
Multiple channel ids can be used separated by comma ','.

```bash
pingme slack --token '123' --channel '1234567890' --message 'some message'
```

- Github Action

```yaml
on:
  release:
    types: [published]
jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      
      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          PUSHOVER_TOKEN: ${{ secrets.SLACK_TOKEN }}
          SLACK_CHANNELS: ${{ secrets.SLACK_CHANNELS }}
          SLACK_MSG_TITLE: 'Refrence: ${{ github.ref }}'
          SLACK_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email
          service: slack
```

- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| SLACK_TOKEN           |         ""        |
| SLACK_CHANNELS   |         ""         | 
| SLACK_MESSAGE            |         ""         | 


## Discord
Discord uses bot token to authenticate & send messages to defined channels.
Multiple channel ids can be used separated by comma ','.

```bash
 pingme discord --token '123' --channel '1234567890' --message 'some message'
```

- Github Action

```yaml
on:
  release:
    types: [published]
jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      
      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          DISCORD_CHANNELS: ${{ secrets.DISCORD_CHANNELS }}
          DISCORD_TOKEN: ${{ secrets.DISCORD_TOKEN }}
          DISCORD_TITLE: 'Refrence: ${{ github.ref }}'
          DISCORD_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email / mattermost
          service: discord
```
- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| DISCORD_TOKEN           |         ""        |
| DISCORD_CHANNELS   |         ""         | 
| DISCORD_MESSAGE            |         ""         | 
| DISCORD_MSG_TITLE        |         ""         |  


## Microsoft Teams
Teams uses webhooks to send messages, you can add multiple webhooks separated by comma ',' or
you can add permissions for multiple channels to single webhook.

```bash
pingme teams --webhook 'https://example.webhook.office.com/xx' --message 'some message'
```

- Github Action

```yaml
on: [push]

jobs:
  pingme-job:
    runs-on: ubuntu-latest
    name: PingMe
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Ping me On
        uses: kha7iq/pingme-action@v1
        env:
          TEAMS_WEBHOOK: ${{ secrets.TEAMS_WEBHOOK }}
          TELEGRAM_CHANNELS: ${{ secrets.TELEGRAM_CHANNELS }}
          TEAMS_MSG_TITLE: 'Refrence: ${{ github.ref }}'
          TEAMS_MESSAGE: 'Event is triggerd by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email / mattermost
          service: teams
```
- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| TEAMS_WEBHOOK           |         ""        |
| TEAMS_MESSAGE            |         ""         | 
| TEAMS_MSG_TITLE        |         ""         |  


## Email
Email uses username  & password to authenticate for sending emails.
SMTP hostname i.e smtp.gmail.com and port i.e (587) should be provided as well for the server.
Multiple email ids can be used separated by comma ',' as receiver email address.
All configuration options are also available via environment variables check configuration section.

```bash
pingme email --rec "example@gmail.com,example@outlook.com" --msg "This is an email from PingMe CLI" --sub "Email from PingMe CLI" \
 --sender "sender@gmail.com" --host "smtp.gmail.com" --port "587" --pass "secretPassword"

```
- **Variables**

|          Variables         | Default Value  | 
| -------------------------- | :----------------: |
| EMAIL_SENDER           |         ""        |
| EMAIL_PASSWORD   |         ""         | 
| EMAIL_RECEIVER            |         ""         | 
| EMAIL_IDENTITY        |         ""         |  
| EMAIL_HOST            |         "smtp.gmail.com"         | 
| EMAIL_PORT            |         "587"         | 
| EMAIL_MESSAGE        |         ""         |  
| EMAIL_SUBJECT            |         ""         | 