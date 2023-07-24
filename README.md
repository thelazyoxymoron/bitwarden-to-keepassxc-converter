# Bitwarden to KeepassXC Converter
KeepassXC doesn't have a native import feature from Bitwarden ([yet](https://github.com/keepassxreboot/keepassxc/issues/8367)). This is a simple script which converts a bitwarden export to a csv format which can be imported in KeepassXC.<br>

It supports exporting your normal website logins, secure notes and credit/debit cards. Credit/debit cards are parsed as a different group and all the card details get added in the notes.

## Requirements
You need to have [golang](https://go.dev/doc/install) installed.

## Usage
- Create an export from bitwarden: Tools -> Export vault -> .json format
- Run the script from the main folder: `go run main.go <exported_file.json>`
- Import the generated file into KeepassXC, check "First line has field names", match and appropriate columns and proceed.
