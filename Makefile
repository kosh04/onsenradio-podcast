TARGET := onsenradio-podcast
GO_FILES := $(shell go list -f '{{ join .GoFiles "\n"}}')

.PHONY: test run clean govendor

$(TARGET): $(GO_FILES)
	go build

run: $(TARGET)
	./$(TARGET) $(ARGS)

test:
	go test -v

clean:
	go clean
