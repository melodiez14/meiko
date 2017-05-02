
all: test build start

run:
	@echo " >> running apps"
	@go run app.go

test:
	@echo " >> running tests"
	@go test -v -race ./... -cover

build:
	@echo " >> building binaries"
	@go build -o bin/meiko app.go

start:
	@echo " >> starting binaries"
	@./bin/meiko

pre-deploy:
	sudo mv bin/meiko /usr/local/bin/.
	sudo cp -r files/etc/meiko/. /etc/meiko/.

profile-build:
	go build -tags profile -o bin/meiko app.go

profile: profile-build
	export LCENV=development
	sudo ./bin/meiko
