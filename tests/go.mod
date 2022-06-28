module github.com/1uvu/bitlog/tests

go 1.17

replace (
	github.com/1uvu/bitlog => ./../bitlog
	github.com/1uvu/bitlog/collector => ./../bitlog/collector
)

require (
	github.com/1uvu/bitlog v0.0.0 // indirect
	github.com/1uvu/bitlog/collector v0.0.0-00010101000000-000000000000 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
)
