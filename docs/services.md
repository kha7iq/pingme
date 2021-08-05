# Configuration

All the flags have corresponding environment variables assosiated with it. You
can either provide the value with flags or export to a variable. You can view
the corresponding variable to each with --help flag.

*Flags* take precedence over *variables*

*Default* value for message title is current *time*

## Telegram

Telegram uses bot token to authenticate & send messages to defined channels.
Multiple channel IDs can be used separated by comma ','.

```bash
pingme  telegram \
--token "0125:AAFHvnYf_ABC" \
--msg "This is a new message ✈" \
--channel="-1001001001,-1002002001"
```

- GitHub Action

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
          TELEGRAM_TITLE: 'Reference: ${{ github.ref }}'
          TELEGRAM_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email
          service: telegram
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| TELEGRAM_TITLE         | ""                 |  
| TELEGRAM_TOKEN             | ""                 |
| TELEGRAM_CHANNELS          | ""                 |
| TELEGRAM_MESSAGE           | ""                 |

## Gotify

With gotify you can send push messages to any instance of Gotify server.

```bash
pingme gotify \
--url 'https://example.com' \
--token 'tokenabc' \
--title 'some title' \
--msg 'some message' \
--priority 5
```

- GitHub Action

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
          GOTIFY_TOKEN: ${{ secrets.TELEGRAM_TOKEN }}
          GOTIFY_TITLE: 'Reference: ${{ github.ref }}'
          GOTIFY_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / pushover / discord / email
          service: gotify
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| GOTIFY_URL       | ""                 |  
| GOTIFY_TOKEN             | ""                 |
| GOTIFY_PRIORITY          | "5"                 |
| GOTIFY_MESSAGE           | ""                 |
| GOTIFY_TITLE         | "Server Time"                 |  

## RocketChat

RocketChat uses token & userID to authenticate and send messages to defined
channels. Multiple channel IDs can be used separated by comma ','.

```bash
pingme rocketchat \
  --channel "general,Pingme" \
  --msg ":wave: rocketchat from cli" \
  --userid "123" \
  --token "abcxyz" \
  --url 'localhost:3000' \
  --scheme "http"
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
          ROCKETCHAT_TITLE: 'Reference: ${{ github.ref }}'
          ROCKETCHAT_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / 
          # pushover / discord / email / mattermost
          service: rocketchat
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| ROCKETCHAT_USERID          | ""                 |
| ROCKETCHAT_TOKEN           | ""                 |
| ROCKETCHAT_SERVER_URL      | ""                 |
| ROCKETCHAT_URL_SCHEME      | "https"            |  
| RTOCKETCHAT_MESSAGE        | ""                 |
| ROCKETCHAT_TITLE           | ""                 |
| ROCKETCHAT_CHANNELS        | ""                 |  

## Pushover

```bash
pingme pushover \
  --token '123' \
  --user '12345567' \
  --title 'some title' \
  --msg 'some message'
```

- GitHub Action

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
          PUSHOVER_TITLE: 'Reference: ${{ github.ref }}'
          PUSHOVER_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / 
          # pushover / discord / email
          service: pushover
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| PUSHOVER_TOKEN             | ""                 |
| PUSHOVER_USER              | ""                 |
| PUSHOVER_MESSAGE           | ""                 |
| PUSHOVER_TITLE             | ""                 |

## Mattermost

Mattermost uses token to authenticate and channel IDs for targets. Destination
server can be specified as 'example.com' by default the 'https' is used, you
can change this with --scheme flag and set it to 'http'. Latest api  version 4
is used for interacting with server, this can also be changed with --api flag.
You can specify multiple channels by separating the value with ','.

```bash
pingme mattermost \
  --token '123' \
  --channel '12345,567' \
  --url 'localhost' \
  --scheme 'http' \
  --msg 'some message'
```

- GitHub Action

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
          MATTERMOST_SERVER_URL: ${{ secrets.MATTERMOST_SERVER_URL }}
          MATTERMOST_CHANNELS: ${{ secrets.MATTERMOST_CHANNELS }}
          MATTERMOST_CHANNELS: ${{ secrets.MATTERMOST_CHANNELS }}
          MATTERMOST_TITLE: 'Reference: ${{ github.ref }}'
          MATTERMOST_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / 
          # pushover / discord / email / mattermost
          service: mattermost
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| MATTERMOST_API_URL         | "/api/v4/posts"    |
| MATTERMOST_TOKEN           | ""                 |
| MATTERMOST_SERVER_URL      | ""                 |
| MATTERMOST_SCHEME          | "https"            |  
| MATTERMOST_MESSAGE         | ""                 |
| MATTERMOST_TITLE           | ""                 |
| MATTERMOST_CHANNELS        | ""                 |  

## Slack

Slack uses token to authenticate and send messages to defined channels.
Multiple channel IDs can be used separated by comma ','.

```bash
pingme slack \
  --token '123' \
  --channel '1234567890' \
  --msg 'some message'
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
          SLACK_TOKEN: ${{ secrets.SLACK_TOKEN }}
          SLACK_CHANNELS: ${{ secrets.SLACK_CHANNELS }}
          SLACK_MSG_TITLE: 'Reference: ${{ github.ref }}'
          SLACK_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / 
          # pushover / discord / email
          service: slack
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| SLACK_TOKEN                | ""                 |
| SLACK_CHANNELS             | ""                 |
| SLACK_MESSAGE              | ""                 |

## Discord

Discord uses bot token to authenticate & send messages to defined channels.
Multiple channel IDs can be used separated by comma ','.

```bash
 pingme discord \
 --token '123' \
 --channel '1234567890' \
 --msg 'some message'
```

- GitHub Action

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
          DISCORD_TITLE: 'Reference: ${{ github.ref }}'
          DISCORD_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams / 
          # pushover / discord / email / mattermost
          service: discord
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| DISCORD_TOKEN              | ""                 |
| DISCORD_CHANNELS           | ""                 |
| DISCORD_MESSAGE            | ""                 |
| DISCORD_MSG_TITLE          | ""                 |  

## Microsoft Teams

Teams uses webhooks to send messages, you can add multiple webhooks separated
by comma ',' or you can add permissions for multiple channels to single webhook.

```bash
pingme teams \
--webhook 'https://example.webhook.office.com/xx' \
--msg 'some message'
``` 

- GitHub Action

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
          TEAMS_MSG_TITLE: 'Reference: ${{ github.ref }}'
          TEAMS_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams /
          # pushover / discord / email / mattermost
          service: teams
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| TEAMS_WEBHOOK              | ""                 |
| TEAMS_MESSAGE              | ""                 |
| TEAMS_MSG_TITLE            | ""                 |  

## Pushbullet

- SMS

```bash
pingme pushbullet \
  --sms true \
  --token "abcdefg" \
  -d "adnroid" \
  --msg "some message" \
  --number "00123456789"
```

- Push notification

```bash
pingme pushbullet --token "abcdefg" -d "adnroid" --msg "some message"
```

- GitHub Action

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
          PUSHBULLET_TOKEN: ${{ secrets.PUSHBULLET_TOKEN }}
          PUSHBULLET_DEVICE: ${{ secrets.PUSHBULLET_DEVICE }}
          PUSHBULLET_TITLE: 'Reference: ${{ github.ref }}'
          PUSHBULLET_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams /
          # pushover / discord / email
          service: pushbullet
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| PUSHBULLET_TOKEN           | ""                 |
| PUSHBULLET_DEVICE          | ""                 |
| PUSHBULLET_NUMBER          | ""                 |
| PUSHBULLET_MESSAGE         | ""                 |
| PUSHBULLET_SMS             | "false"            |
| PUSHBULLET_TITLE           | ""                 |

## Twillio SMS

SMS can be sent via twillio to multiple numbers, you can add multiple receivers
separated by a comma.

```bash
 pingme twillio \
   --token 'tokenabc' \
   --account 'sid123' \
   --sender '+140001442' \
   --receiver '+140001442' \
   --msg 'some message'
```

- GitHub Action

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
          TWILLIO_TOKEN: ${{ secrets.TWILLIO_TOKEN }}
          TWILLIO_ACCOUNT_SID: ${{ secrets.TWILLIO_ACCOUNT_SID }}
          TWILLIO_SENDER: ${{ secrets.TWILLIO_SENDER }}
          TWILLIO_RECEIVER: ${{ secrets.TWILLIO_RECEIVER }}
          TWILLIO_TITLE: 'Reference: ${{ github.ref }}'
          TWILLIO_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams /
          # pushover / discord / email / mattermost / twillio
          service: twillio
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| TWILLIO_TOKEN              | ""                 |
| TWILLIO_ACCOUNT_SID        | ""                 |
| TWILLIO_SENDER             | ""                 |  
| TWILLIO_RECEIVER           | ""                 |
| TWILLIO_TITLE              | ""                 |
| TWILLIO_MESSAGE            | ""                 |

## Mastodon

Mastodon uses application token to authorize and set status.

```bash
mastodon --url "mastodon.social" --msg "some message" --title "PingMe CLI" --token "123"
```

- GitHub Action

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
          MASTODON_TOKEN: ${{ secrets.MASTODON_TOKEN }}
          MASTODON_SERVER: 'mastodon.social'
          MASTODON_TITLE: 'Reference: ${{ github.ref }}'
          MASTODON_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          service: mastodon
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| MASTODON_TOKEN              | ""                 |
| MASTODON_SERVER              | ""                 |
| MASTODON_TITLE            | ""                 |
| MASTODON_MESSAGE            | ""                 |  

## Zulip

Zulip uses bot email and token for authentication, and sends messages to particular topic.

```bash
pingme zulip 
 --email 'john.doe@email.com' \ 
 --api-key '12345567' \ 
 --to 'london' \
 --type 'stream' \
 --topic 'some topic' \
 --msg 'content of message'
```

- GitHub Action

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
          ZULIP_DOMAIN: ${{ secrets.ZULIP_DOMAIN }}
          ZULIP_BOT_EMAIL_ADDRESS: ${{ secrets.ZULIP_BOT_EMAIL_ADDRESS }}
          ZULIP_BOT_API_KEY: ${{ secrets.ZULIP_BOT_API_KEY }}
          ZULIP_MSG_TYPE: 'stream'
          ZULIP_STREAM_NAME: 'general'
          ZULIP_TOPIC: 'Reference: ${{ github.ref }}'
          ZULIP_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          service: zulip
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| ZULIP_DOMAIN              | ""                 |
| ZULIP_BOT_EMAIL_ADDRESS              | ""                 |
| ZULIP_BOT_API_KEY            | ""                 |
| ZULIP_MSG_TYPE            | ""                 |
| ZULIP_STREAM_NAME            | ""                 |
| ZULIP_TOPIC            | ""                 |  
| ZULIP_MESSAGE            | ""                 |  

## Line

Line uses channel secret and token for authentication, and sends messages.

```bash
pingme line 
 --secret 'secretxxx' \ 
 --token '12345567' \ 
 --receivers 'ab1234545xx' \
 --msg 'content of message'
```

- GitHub Action

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
          LINE_SECRET: ${{ secrets.LINE_SECRET }}
          LINE_TOKEN: ${{ secrets.LINE_TOKEN }}
          LINE_RECEIVER_IDS: 'ab1235xxx8'
          LINE_MSG_TITLE: 'Reference: ${{ github.ref }}'
          LINE_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          service: line
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| LINE_SECRET              | ""                 |
| LINE_TOKEN              | ""                 |
| LINE_RECEIVER_IDS            | ""                 |
| LINE_MSG_TITLE            | ""                 |
| LINE_MESSAGE            | ""                 |

## Email

Email uses username & password to authenticate for sending emails. SMTP
hostname i.e smtp.gmail.com and port i.e (587) should be provided as well for
the server. Multiple email IDs can be used separated by comma ',' as receiver
email address. All configuration options are also available via environment
variables. Check configuration section.

```bash
pingme email \
  --rec "example@gmail.com,example@outlook.com" \
  --msg "This is an email from PingMe CLI" \
  --sub "Email from PingMe CLI" \
  --sender "sender@gmail.com" \
  --host "smtp.gmail.com" \
  --port "587" \
  --pass "secretPassword"
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| EMAIL_SENDER               | ""                 |
| EMAIL_PASSWORD             | ""                 |
| EMAIL_RECEIVER             | ""                 |
| EMAIL_IDENTITY             | ""                 |
| EMAIL_HOST                 | "smtp.gmail.com"   |
| EMAIL_PORT                 | "587"              |
| EMAIL_MESSAGE              | ""                 |  
| EMAIL_SUBJECT              | ""                 |


## Wechat Official Account

Wechat uses appid, appsecret, chatbot server token and encoding AES key for authentication, and sends messages.

```bash
pingme wechat
--appid "xxxxxxxx" \
--appsecret 'xxxxxxxxxx' \
--token 'xxxxxxxxxx' \
--aes 'IGNORED-IN-SANDBOX' \
--msg 'content of message' \
--receivers 'some receivers'
```

- GitHub Action

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
          WECHAT_APPID: ${{ secrets.WECHAT_APPID }}
          WECHAT_APPSECRET: ${{ secrets.WECHAT_APPSECRET }}
          WECHAT_TOKEN: ${{ secrets.WECHAT_TOKEN }}
          WECHAT_ENCODINGAESKEY: ${{ secrets.WECHAT_ENCODINGAESKEY }}
          WECHAT_RECEIVER_IDS: ${{ secrets.WECHAT_RECEIVER_IDS }}
          WECHAT_MSG_TITLE: 'Reference: ${{ github.ref }}'
          WECHAT_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
        
        with:
          service: wechat
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| WECHAT_APPID               | ""                 |
| WECHAT_APPSECRET           | ""                 |
| WECHAT_TOKEN               | ""                 |
| WECHAT_ENCODINGAESKEY      | ""                 |
| WECHAT_RECEIVER_IDS        | ""                 |
| WECHAT_MSG_TITLE           | ""                 |
| WECHAT_MESSAGE             | ""                 |

## TextMagic

TextMagic allows you to send SMS to multiple receivers provided by comma separated string.

```bash
 pingme textmagic \
   --token 'tokenabc' \
   --user 'sid123' \
   --title 'message title' \
   --receiver '+140001442' \
   --msg 'some message'
```

- GitHub Action

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
          TEXTMAGIC_USER: ${{ secrets.TEXTMAGIC_USER }}
          TEXTMAGIC_TOKEN: ${{ secrets.TEXTMAGIC_TOKEN }}
          TEXTMAGIC_TITLE: 'Reference: ${{ github.ref }}'
          TEXTMAGIC_MESSAGE: 'Event is triggered by ${{ github.event_name }}'
          TEXTMAGIC_RECEIVER:  ${{ secrets.TEXTMAGIC_RECEIVER }}
        
        with:
          # Chose the messaging platform. 
          # slack / telegram / rocketchat / teams /
          # pushover / discord / email / mattermost / textmagic
          service: textmagic
```

- **Variables**

|          Variables         | Default Value      |
| -------------------------- | :----------------: |
| TEXTMAGIC_USER             | ""                 |
| TEXTMAGIC_TOKEN            | ""                 |
| TEXTMAGIC_TITLE          | ""                 | 
| TEXTMAGIC_MESSAGE          | ""                 | 
| TEXTMAGIC_RECEIVER         | ""                 | 
