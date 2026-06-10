.PHONY: generate check test build

generate:
	uv run scripts/generate.py

check:
	uv run scripts/generate.py --check

test: build
	cd packages/typescript && npm test
	cd packages/python && uv run --with pytest pytest tests/ -v
	cd packages/go && go test ./... -v

build:
	cd packages/typescript && npm install && npm run build
	cd packages/python && uv build
	cd packages/go && go build ./...
