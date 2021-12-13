# Installation

## MacOS & Linux Homebrew

```bash
brew install kha7iq/tap/pingme
```

## Shell Script
By default pingme is going to be installed at `/usr/bin/` sudo is requried for this operation.

If you would like to provide a custom install path you can do so as input to script. i.e `./install.sh $HOME/bin`
```bash
curl -s https://raw.githubusercontent.com/kha7iq/pingme/master/install.sh | sudo sh
```
or
```bash
curl -sL https://bit.ly/installpm | sudo sh
```

## Linux

* AUR
```bash
yay -S pingme-bin
```

## Manual
```bash
# Chose desired version, architecture & target os
export PINGME_VERSION="0.2.4"
export ARCH="x86_64"
export OS="Linux"
wget -q https://github.com/kha7iq/pingme/releases/download/v${PINGME_VERSION}/pingme_${OS}_${ARCH}.tar.gz && \
tar -xf pingme_${OS}_${ARCH}.tar.gz && \
chmod +x pingme && \
sudo mv pingme /usr/local/bin/pingme
```

## Windows

```powershell
scoop bucket add pingme https://github.com/kha7iq/scoop-bucket.git
scoop install pingme
```

Alternatively you can head over to [release pages](https://github.com/kha7iq/pingme/releases)
and download the binary for windows & all other supported platforms.

## Docker

Docker container is also available on both dockerhub and github container registry.

`latest` tage will always pull the latest version avaialbe, you can also
download specific version. Checkout [release](https://github.com/kha7iq/pingme/releases)
page for available versions.

- Docker Registry

```bash
docker pull khaliq/pingme:latest
```

- GitHub Registry

```bash
docker pull ghcr.io/kha7iq/pingme:latest
```

- Run

```bash
docker run ghcr.io/kha7iq/pingme:latest
```

## GitHub Action

A github action is also available now for this app, you can find it on
[Github Market Place](https://github.com/marketplace/actions/pingme-action)
or from this [repository](https://github.com/kha7iq/pingme-action) on github.

Usage examples for workflow are available in the repo.
