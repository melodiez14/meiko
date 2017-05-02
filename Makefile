
all: test build start

run:
	@echo " >> running apps"
	@go run app.go

test:
	@echo " >> running tests"
	@go test -v -race ./... -cover

build:
	@echo " >> building binaries"
	@go build -o bin/lastcake app.go

start:
	@echo " >> starting binaries"
	@./bin/lastcake

pre-deploy:
	sudo mv bin/lastcake /usr/local/bin/.
	sudo cp -r files/etc/lastcake/. /etc/lastcake/.

profile-build:
	go build -tags profile -o bin/lastcake app.go

profile: profile-build
	export LCENV=development
	sudo ./bin/lastcake
