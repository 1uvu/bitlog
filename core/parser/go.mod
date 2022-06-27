module github.com/1uvu/bitlog/core/parser

go 1.18

replace (
	github.com/1uvu/bitlog/core => ../
	github.com/1uvu/bitlog/core/collector => ./
)

require (
	github.com/1uvu/bitlog v0.0.0-20220407081013-ba4ecf84323a
	github.com/btcsuite/btcd v0.22.0-beta.0.20210803133449-f5a1fb9965e4
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
)

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/schollz/progressbar/v3 v3.8.5 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)
