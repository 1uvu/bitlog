module github.com/1uvu/bitlog/core/storage

go 1.18

replace (
	github.com/1uvu/bitlog/core => ../
	github.com/1uvu/bitlog/core/parser => ../parser
	github.com/1uvu/bitlog/core/storage => ./
)

require (
	github.com/1uvu/bitlog/core v0.0.0
	github.com/1uvu/bitlog/core/parser v0.0.0
)
