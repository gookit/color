# Examples

there are some demo for use gookit/color.

## tests on docker

check ENV info:

```bash
# on manjaro
docker run --rm -ti golang:1.16 env | grep -i term
docker run --rm -ti jonathonf/manjaro bash
docker run --rm -ti -v ~/Workspace/godev/gookit/color:/var/gocolor jonathonf/manjaro bash

# on alpine linux
docker run --rm -ti alpine:3.12 env | grep -i term

# on alpine linux with golang
docker run --rm -ti golang:1.16-alpine env | grep -i term

# on ubuntu linux
docker run --rm -ti ubuntu:19.04 env | grep -i term

# on debian linux with golang
docker run --rm -ti -v ~/Workspace/godev/gookit/color:/var/gocolor golang:1.16 bash
```

run demo on docker with golang:

```bash
# on alpine linux with golang
docker run --rm -ti -v ~/Workspace/godev/gookit/color:/var/gocolor golang:1.16-alpine sh

# on debian linux with golang
docker run --rm -ti -v ~/Workspace/godev/gookit/color:/var/gocolor golang:1.16 bash
```

> TIPS: can set go proxy for fetch deps. eg: `export GOPROXY=https://goproxy.cn,direct`

run example demos:

```bash
/go # cd /var/gocolor/
/var/gocolor # export GOPROXY=https://goproxy.cn,direct
/var/gocolor # go run ./_examples/color_256.go
```

more examples:

```bash
COLOR_DEBUG_MODE=on go run ./_examples/envcheck.go

# examples:
go run ./_examples/color_16.go
go run ./_examples/color_256.go
go run ./_examples/color_tag.go
go run ./_examples/color_rgb.go
```
