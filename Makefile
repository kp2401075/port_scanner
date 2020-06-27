.POSIX:
.SUFFIXES:

GO?=go
GOFLAGS?=
GOSRC!=find . -name '*.go'

portscanner: $(GOSRC)
	$(GO) build $(GOFLAGS) \
		-o port_scanner port_scanner.go

all: portscanner

RM?=rm -f

clean:
	$(RM) -r port_scanner