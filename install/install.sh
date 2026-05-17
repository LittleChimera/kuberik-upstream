#!/usr/bin/env bash
# install.sh - Download and install the kuberik CLI.
#
#   curl -s https://raw.githubusercontent.com/kuberik/kuberik/main/install/install.sh | sudo bash
#
# Environment variables:
#   KUBERIK_VERSION  Specific release to install (default: latest)
#   KUBERIK_BINDIR   Install destination (default: /usr/local/bin)
#
set -euo pipefail

REPO="kuberik/kuberik"
BINDIR="${KUBERIK_BINDIR:-/usr/local/bin}"
VERSION="${KUBERIK_VERSION:-}"

require() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "error: required tool '$1' not found on PATH" >&2
    exit 1
  }
}

require curl
require tar
require uname

if [[ -z "${VERSION}" ]]; then
  VERSION=$(curl -sSL -o /dev/null -w '%{url_effective}' "https://github.com/${REPO}/releases/latest" | sed 's|.*/||')
  if [[ -z "${VERSION}" ]]; then
    echo "error: could not determine latest release" >&2
    exit 1
  fi
fi
VERSION="${VERSION#v}"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "${OS}" in
  darwin|linux) ;;
  *) echo "error: unsupported OS: ${OS}" >&2; exit 1 ;;
esac

ARCH=$(uname -m)
case "${ARCH}" in
  x86_64|amd64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  armv7l|armv7)  ARCH="arm" ;;
  *) echo "error: unsupported arch: ${ARCH}" >&2; exit 1 ;;
esac

ARCHIVE="kuberik_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${ARCHIVE}"
CHECKSUMS_URL="https://github.com/${REPO}/releases/download/v${VERSION}/checksums.txt"

TMP=$(mktemp -d)
trap 'rm -rf "${TMP}"' EXIT

echo "downloading ${URL}"
curl -fsSL -o "${TMP}/${ARCHIVE}" "${URL}"
curl -fsSL -o "${TMP}/checksums.txt" "${CHECKSUMS_URL}"

(cd "${TMP}" && grep " ${ARCHIVE}\$" checksums.txt | sha256sum -c -)

tar -xzf "${TMP}/${ARCHIVE}" -C "${TMP}" kuberik

if [[ -w "${BINDIR}" ]]; then
  install -m 0755 "${TMP}/kuberik" "${BINDIR}/kuberik"
else
  echo "installing to ${BINDIR} requires elevated privileges"
  sudo install -m 0755 "${TMP}/kuberik" "${BINDIR}/kuberik"
fi

echo "installed kuberik ${VERSION} to ${BINDIR}/kuberik"
"${BINDIR}/kuberik" version
