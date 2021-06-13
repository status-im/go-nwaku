.PHONY: all build

build:
	go build -o build/nwaku nwaku.go

all: build

# TODO Integrate with nim-waku wrapper
# TODO Assume we have libwaku.so already
# libraries for dynamic linking of non-Nim objects
EXTRA_LIBS_DYNAMIC := -L"$(CURDIR)/build" -lwaku -lm
wrapper2: # | build deps libwaku.so
	echo -e $(BUILD_MSG) "build/wrapper2" && \
		go build -ldflags "-linkmode external -extldflags '$(EXTRA_LIBS_DYNAMIC)'" -o build/wrapper2 nwaku/wrapper2.go
