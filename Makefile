TARGET = simple_db
VERSION="1.0.0"
BUILD_DIR="build"
BUILD_TARGET="$(BUILD_DIR)/$(TARGET)"

default: $(BUILD_TARGET)
all: clean test

$(BUILD_TARGET):
	go build -o $(BUILD_TARGET)

test: $(BUILD_TARGET)
	go test -cover github.com/kejne/simple_database/db
	./integration-tests.sh

install:
	go get

clean:
	-rm -r $(BUILD_DIR)
	-rm -r *.exe