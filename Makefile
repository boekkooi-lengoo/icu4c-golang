
tools_path=$(CURDIR)/.tools
tools_vendor=$(tools_path)/vendor

c-for-go=$(tools_path)/bin/c-for-go
icu4c_path="$(CURDIR)/icu"

run: build-icu4c
	go run main.go

test:
	go test "$(CURDIR)/..."

test-fuzz:
	go test -fuzz=. "$(CURDIR)/icu" -fuzztime 60s

benchmark:
	go test -bench=. "$(CURDIR)/..."

build: build-icu4c
	go build main.go

build-icu4c: $(c-for-go) clean-icu4c
	$(c-for-go) -nostamp "$(CURDIR)/icu4c.yaml"
	sed -i -e 's/NewbreakIterator/newBreakIterator/g' -e 's/Newtext/newText/g' "$(CURDIR)/icu/cgo_helpers.go"
	sed -i -e 's/NewbreakIterator/newBreakIterator/g' -e 's/Newtext/newText/g' "$(CURDIR)/icu/icu.go"

clean-icu4c:
	@rm -f \
		"$(icu4c_path)/cgo_helpers.go" \
		"$(icu4c_path)/cgo_helpers.h" \
		"$(icu4c_path)/const.go" \
	 	"$(icu4c_path)/doc.go" \
	 	"$(icu4c_path)/types.go" \
 		"$(icu4c_path)/icu.go"

tools: $(c-for-go)

$(tools_vendor):
	cd $(tools_path); \
	go mod vendor

$(c-for-go): $(tools_vendor)
	cd $(tools_path); \
	go build -o "$(c-for-go)" $(tools_vendor)/github.com/xlab/c-for-go/
