# Development Server Commands
.PHONY: run
run:
	mkdir -p ./bin && chmod -R 777 ./public
	make -j4 reload tailwind-dev templ-watch air

.PHONY: reload
reload:
	go run ./cmd/utils/reload/main.go

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --watch
	
.PHONY: tailwind-dev
tailwind-dev:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css

.PHONY: templ-watch
templ-watch:
	@templ generate -keep-orphaned-files --watch

.PHONY: air
air:
	air


.PHONY: kill-process
kill-process:
	@kill -15 $(lsof -ti:3000)



# Production Server Commands
.PHONY: build
build:
	go mod tidy
	mkdir -p ./bin && chmod -R 777 ./public
	make templ tailwind vet staticcheck test
	go build -o ./bin/app ./cmd/app/main.go

.PHONY: templ
templ:
	templ generate -keep-orphaned-files

.PHONY: templ-clean
templ-clean:
	find . -type f \( -name '*_templ.txt' \) -exec rm -f {} +

.PHONY: tailwind
tailwind:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.min.css --minify

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck: 
	staticcheck ./...
	
.PHONY: test
test:
	go test -race -v -timeout 30s ./...



# Database Migration-Seeding-Parsing Commands
.PHONY: migrate
migrate:
	go run ./cmd/db/migrate/main.go

.PHONY: drop
drop:
	go run ./cmd/db/migrate/drop/main.go

.PHONY: seed
seed:
	go run ./cmd/db/seed/main.go

.PHONY: parse
parse:
	mkdir -p ./public/uploads/products && chmod 777 ./public/uploads/products/*
	python3 ./cmd/db/parser/products.py


.PHONY: env
env:
	cp -r .example.env .env

	
#