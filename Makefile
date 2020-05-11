APP_NAME 	= favorite-tree
IMPORT_PATH = github.com/manue1/${APP_NAME}

build:
	@echo "==> Building the app..."
	docker build --pull --tag ${APP_NAME} --build-arg importPath=${IMPORT_PATH} .

up:
	@echo "==> Starting the app..."
	docker run -p 8000:8000 ${APP_NAME}
