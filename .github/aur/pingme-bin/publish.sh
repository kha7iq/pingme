#!/usr/bin/env bash


set -e

WD=$(cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)
PKGNAME=$(basename $WD)
repo="kha7iq/pingme"
export _PKGNAME="${PKGNAME%-bin}"
ROOT=${WD%/.github/aur/$PKGNAME}

LOCKFILE=/tmp/aur-$PKGNAME.lock
exec 100>$LOCKFILE || exit 0
flock -n 100 || exit 0
trap "rm -f $LOCKFILE" EXIT

export PKGVER=$(curl -s https://api.github.com/repos/$repo/releases/latest | grep -o '"tag_name": "v[^"]*' | cut -d'"' -f4 | sed 's/^v//')

echo "Publishing to AUR as version ${PKGVER}"

cd $WD

export GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no"

if [ -z "$AUR_SSH_PRIVATE_KEY" ]
then
    echo "\$AUR_SSH_PRIVATE_KEY is not set"
else
    eval $(ssh-agent -s)
    ssh-add <(echo "$AUR_SSH_PRIVATE_KEY")
fi

GITDIR=$(mktemp -d /tmp/aur-$PKGNAME-XXX)
trap "rm -rf $GITDIR" EXIT
git clone aur@aur.archlinux.org:$PKGNAME $GITDIR 2>&1

CURRENT_PKGVER=$(cat $GITDIR/.SRCINFO | grep pkgver | awk '{ print $3 }')
CURRENT_PKGREL=$(cat $GITDIR/.SRCINFO | grep pkgrel | awk '{ print $3 }')

if [[ "${CURRENT_PKGVER}" == "${PKGVER}" ]]; then
    export PKGREL=$((CURRENT_PKGREL+1))
else
    export PKGREL=1
fi

export SHA256SUM_X86=$(curl -sL https://github.com/kha7iq/${_PKGNAME}/releases/download/v${PKGVER}/${_PKGNAME}_linux_x86_64.tar.gz | sha256sum | awk '{ print $1 }')

export SHA256SUM_i686=$(curl -sL https://github.com/kha7iq/${_PKGNAME}/releases/download/v${PKGVER}/${_PKGNAME}_linux_i386.tar.gz | sha256sum | awk '{ print $1 }')

export SHA256SUM_AARCH64=$(curl -sL https://github.com/kha7iq/$_PKGNAME/releases/download/v${PKGVER}/${_PKGNAME}_linux_arm64.tar.gz | sha256sum | awk '{ print $1 }') 

envsubst '$PKGVER $PKGREL $_PKGNAME $SHA256SUM_X86 $SHA256SUM_i686 $SHA256SUM_AARCH64' < .SRCINFO.template > $GITDIR/.SRCINFO
envsubst '$PKGVER $PKGREL $SHA256SUM_X86 $SHA256SUM_i686 $SHA256SUM_AARCH64' < PKGBUILD.template > $GITDIR/PKGBUILD

cd $GITDIR
git config user.name "kha7iq"
git config user.email "a.khaliq@outlook.my"
git add -A
if [ -z "$(git status --porcelain)" ]; then
  echo "No changes."
else
  git commit -m "Updated to version v${PKGVER} release ${PKGREL}"
  git push origin master
fi
