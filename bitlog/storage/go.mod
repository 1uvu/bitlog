module github.com/1uvu/bitlog/storage

go 1.18

replace (
	github.com/1uvu/bitlog => ../
	github.com/1uvu/bitlog/parser => ../parser
)

require (
	github.com/1uvu/bitlog v0.0.0-00010101000000-000000000000
	github.com/1uvu/bitlog/parser v0.0.0-00010101000000-000000000000
)
