GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURERNT_BRANCH := $(shell git branch --show-current)
export GO111MODULE := on

test:
	go test -race ./...

build: *.go cmd/wasataro/*.go go.*
	cd cmd/wasataro && go build -ldflags "-s -w -X main.Version=${GIT_VER} -X main.buildDate=${DATE}" -gcflags="-trimpath=${PWD}"

install: cmd/wasataro/wasataro
	install cmd/wasataro/wasataro ${GOPATH}/bin

packages:
	cd cmd/wasataro && gox -os="linux darwin" -arch="amd64" -output "../../pkg/{{.Dir}}-${GIT_VER}-{{.OS}}-{{.Arch}}" -ldflags "-w -s -X main.Version=${GIT_VER} -X main.buildDate=${DATE}"
	cd pkg && find . -name "*${GIT_VER}*" -type f -exec zip {}.zip {} \;

clean:
	rm -f cmd/wasataro/wasataro
	rm -f pkg/*

release:
	ghr -prerelease -u onigra -r wasataro -n "$(GIT_VER)" $(GIT_VER) pkg/

docker/build:
	docker buildx build --tag ghcr.io/onigra/wasataro:$(CURERNT_BRANCH) --file ./Dockerfile .

docker/run:
	docker run --rm -p 3000:3000 ghcr.io/onigra/wasataro:$(CURERNT_BRANCH); trap 'echo Trapped' 2
