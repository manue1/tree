# favorite-tree

This web server displays your favorite tree in a HTML document if provided as a query parameter.

## Build & run

In order to run this web server you need to first build the image. Then when running it you may find the endpoint exposed on port `:8000`.

```sh
make build
make up
```

## Pending improvements

Things that need more attention to reach the desired state:

- [ ] Main tests
- [ ] Move handler out of main package
- [ ] Add metrics
- [ ] Create API documentation
- [ ] Multi-stage build of Docker image to serve the binary in a scratch container
- [X] Serve actual HTML
- [X] Add query parameter check to handler
- [X] Gracefully shut down the server
