define bench
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench . $(1)
endef

define check-escape
	go build -gcflags '-m' $(1)
endef

all: bench-poollog check-escape-poollog bench-normallog check-escape-normallog

bench-poollog:
	$(call bench,./poollog/...)

check-escape-poollog:
	$(call check-escape,./poollog/...)

bench-normallog:
	$(call bench,./normallog/...)

check-escape-normallog:
	$(call check-escape,./normallog/...)

