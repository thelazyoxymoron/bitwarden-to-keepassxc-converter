# Bitwarden to KeepassXC Converter
KeepassXC does have a native import feature from Bitwarden starting with [version 2.7.7](https://github.com/keepassxreboot/keepassxc/releases/tag/2.7.7) but it does not ([yet support organization exports](https://github.com/keepassxreboot/keepassxc/pull/10499)) correctly. This is a simple script which converts a bitwarden export to a csv format which can be imported in KeepassXC.<br>

It supports exporting your normal website logins including TOTP, secure notes and credit/debit cards. Since there are no username/password fields for credit/debit card type, the script adds a new group called "Cards" and imports all the card details in "notes" field.

## Requirements
You need to have [golang](https://go.dev/doc/install) installed.

## Usage
- Create an export from bitwarden: Tools -> Export vault -> .json format
- Run the script from the main folder: `go run json2csv.go <exported_file.json>`
- Import the generated file into KeepassXC, check "First line has field names", match appropriate columns and proceed.
