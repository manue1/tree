# favorite-tree

This web server displays your favorite tree in a HTML document if provided as a query parameter.

## Build & run

In order to run this web server you need to first build the image. Then when running it you may find the endpoint exposed on port `:8000`.

```sh
make build up
```

## Run tests

For running the available unit-tests the following command can be used. It builds the image before-hand and runs the tests and outputs the given coverage.

```sh
make test
```

## Pending improvements

Improvements that would be needed to reach the desired state:

- [ ] Integration tests (main)
- [ ] Add HTTP request metrics
- [ ] Create API documentation
- [ ] Test coverage of 100% (right now it is 75%)
- [ ] Multi-stage build of Docker image to serve the binary in a scratch container
