all: test

test: 
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench .
