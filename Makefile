TARGET := onsenradio-podcast

$(TARGET):
	go build

test:
	go test -v

clean:
	$(RM) $(TARGET)
