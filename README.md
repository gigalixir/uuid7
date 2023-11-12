UUID7
---
A very simple go program that prints a UUIDv7


## Development in Docker
``` bash
docker build -t uuid7-dev .
docker run -v $(pwd):/app uuid7-dev
```

The binary will be compiled as `uuid7`.
