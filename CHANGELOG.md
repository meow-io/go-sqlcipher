# Changelog

## 1.1.0 - Sep 2nd 2026

Upgrade to v4.18.0 of sqlcipher and v1.14.50 of go-sqlite3, and update [libtomcrypt](https://github.com/libtom/libtomcrypt) to digest 6c6d5104de66f3ca0dfd7b68540ef86869982b07.

This raises the minimum Go version to 1.21, as go-sqlite3 now uses `any`, `unsafe.Slice` and `unsafe.StringData`.

## 1.0.3 - Sep 1st 2023

Upgrade to v4.5.5 of sqlcipher and v1.14.17 of go-sqlite3.

## 1.0.2 - Jul 6th 2023

Upgrade to v4.5.4 of sqlcipher.

## 1.0.1 - Jun 16th 2023

Transfers ownership of package to meow-io.

## 1.0.0 - Jun 11th 2023

Initial release. This is based on v1.14.16 of [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) and v4.5.3 of [sqlcipher](https://www.zetetic.net/sqlcipher/). As well this uses [libtomcrypt](https://github.com/libtom/libtomcrypt) at digest 2a1b284677a51f587ab7cd9d97395e0c0c93a447.
