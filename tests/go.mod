module github.com/1uvu/bitlog/tests

go 1.17

replace (
	github.com/1uvu/bitlog/core => ./../core
	github.com/1uvu/bitlog/core/collector => ./../core/collector
)

require (
	github.com/1uvu/bitlog/core v0.0.0 // indirect
	github.com/1uvu/bitlog/core/collector v0.0.0-00010101000000-000000000000 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
)
