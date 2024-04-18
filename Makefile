
tools_path=$(CURDIR)/.tools
tools_vendor=$(tools_path)/vendor

c-for-go=$(tools_path)/bin/c-for-go

run: build-icu4c
	go run main.go

build: build-icu4c
	go build main.go

build-icu4c: $(c-for-go)
	$(c-for-go) -nostamp $(CURDIR)/icu4c.yaml
	sed -i -e 's/NewbreakIterator/newBreakIterator/g' -e 's/Newtext/newText/g' "$(CURDIR)/icu/cgo_helpers.go"
	sed -i -e 's/NewbreakIterator/newBreakIterator/g' -e 's/Newtext/newText/g' "$(CURDIR)/icu/icu.go"

tools: $(c-for-go)

$(tools_vendor):
	cd $(tools_path); \
	go mod vendor

$(c-for-go): $(tools_vendor)
	cd $(tools_path); \
	go build -o "$(c-for-go)" $(tools_vendor)/github.com/xlab/c-for-go/
