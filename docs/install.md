
## Linux & MacOs
```bash
brew install kha7iq/tap/pingme
```

## Go Get
```bash
go get -u github.com/kha7iq/pingme
```


## Windows
```powershell
scoop bucket add pingme https://github.com/kha7iq/scoop-bucket.git
scoop install pingme
```

Alternatively you can head over to [release pages](https://github.com/kha7iq/pingme/releases) and download the binary for windows & all other supported platforms.


## Docker
Docker container is also available on both dockerhub and github container registry.

`latest` tage will always pull the latest version avaialbe, you can also download specific version.
Checkout [release](https://github.com/kha7iq/pingme/releases) page for available versions.

- Docker Registry
```bash
docker pull khaliq/pingme:latest
```
- Github Registry
```bash
docker pull ghcr.io/kha7iq/pingme:latest
```
- Run
```bash
docker run ghcr.io/kha7iq/pingme:latest
```


## Github Action
A github action is also available now for this app, you can find it on [Github Market Place](https://github.com/marketplace/actions/pingme-action) or from this [repository](https://github.com/kha7iq/pingme-action) on github.

Usage examples for workflow are available in the repo.
