# go-builder
Easily build Go applications to multiple platforms with one command

## How it works
go-builder simply sets the adequate environment variables and executes the `go build` command for each specified platform/architecture.

Here's an example of go-builder in action: https://asciinema.org/a/369411

## Usage
`./go-builder -s <source.go> -o <output_%> -t <win64[,lin386[,linarm64]]>`

You can use source files from another directory: `-s ../src/main.go`. You can also build into another directory: `-o bin/app_%`

Output must contain the `%` character, which will be replaced with the target OS and architecture in the format `os-arch`.

If the target OS is `windows`, the file name will be automatically suffixed with `.exe`.

Targets must be sepparated by comma. A full list of target codes is available below.

## Build targets
The way the target codes are defined is as follows:

Begin with the first 3 letters of the OS. `win` for windows and `lin` for linux.

If the OS has less than 3 letters, all letters should be used (like `js`).

Then, if the architecture is `amd64`, add `64`. Otherwise, use the full Arch name: `386`, `arm`, `arm64`, `ppc64`, `mips64le`.

So if you want to build for `android/arm64` you should use `andarm64`.

Full list of build codes:
- aixppc64 (aix/ppc64)
- and386 (android/386)
- and64 (android/amd64)
- andarm (android/arm)
- andarm64 (android/arm64)
- dar64 (darwin/amd64)
- dararm64 (darwin/arm64)
- dra64 (dragonfly/amd64)
- fre386 (freebsd/386)
- fre64 (freebsd/amd64)
- frearm (freebsd/arm)
- frearm64 (freebsd/arm64)
- ill64 (illumos/amd64)
- jswasm (js/wasm)
- lin386 (linux/386)
- lin64 (linux/amd64)
- linarm (linux/arm)
- linarm64 (linux/arm64)
- linmips (linux/mips)
- linmips64 (linux/mips64)
- linmips64le (linux/mips64le)
- linmipsle (linux/mipsle)
- linppc64 (linux/ppc64)
- linppc64le (linux/ppc64le)
- linriscv64 (linux/riscv64)
- lins390x (linux/s390x)
- net386 (netbsd/386)
- net64 (netbsd/amd64)
- netarm (netbsd/arm)
- netarm64 (netbsd/arm64)
- ope386 (openbsd/386)
- ope64 (openbsd/amd64)
- opearm (openbsd/arm)
- opearm64 (openbsd/arm64)
- pla386 (plan9/386)
- pla64 (plan9/amd64)
- plaarm (plan9/arm)
- sol64 (solaris/amd64)
- win386 (windows/386)
- win64 (windows/amd64)
- winarm (windows/arm)
