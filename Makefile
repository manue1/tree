APP_NAME 	= favorite-tree
IMPORT_PATH = github.com/manue1/${APP_NAME}

up:
	@echo "==> Building the app..."
	docker build --pull --tag ${APP_NAME} --build-arg importPath=${IMPORT_PATH} .
	@echo "==> Starting the app..."
	docker run -p 8000:8000 ${APP_NAME}

test:
	@echo "==> Building unit-tests..."
	docker build --pull --file unittest.Dockerfile \
		--tag ${APP_NAME}-unittest \
		--build-arg importPath=${IMPORT_PATH} .

	@echo "==> Running unit-tests..."
	docker run --rm ${APP_NAME}-unittest go test -v -parallel=4 -race -cover ./...
