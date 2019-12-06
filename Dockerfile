from golang:alpine
run apk add musl-dev
run apk add git
run apk add gcc
run mkdir /coypu/
workdir /coypu/
copy ./ /coypu/
run sh get_deps.sh
workdir /coypu/config/plugins/
run sh build.sh
workdir /coypu
run go build
cmd ./coypu
