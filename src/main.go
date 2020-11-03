package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*

aixppc64 (aix/ppc64)

and386 (android/386)
and64 (android/amd64)
andarm (android/arm)
andarm64 (android/arm64)

dar64 (darwin/amd64)
dararm64 (darwin/arm64)

dra64 (dragonfly/amd64)

fre386 (freebsd/386)
fre64 (freebsd/amd64)
frearm (freebsd/arm)
frearm64 (freebsd/arm64)

ill64 (illumos/amd64)

jswasm (js/wasm)

lin386 (linux/386)
lin64 (linux/amd64)
linarm (linux/arm)
linarm64 (linux/arm64)
linmips (linux/mips)
linmips64 (linux/mips64)
linmips64le (linux/mips64le)
linmipsle (linux/mipsle)
linppc64 (linux/ppc64)
linppc64le (linux/ppc64le)
linriscv64 (linux/riscv64)
lins390x (linux/s390x)

net386 (netbsd/386)
net64 (netbsd/amd64)
netarm (netbsd/arm)
netarm64 (netbsd/arm64)

ope386 (openbsd/386)
ope64 (openbsd/amd64)
opearm (openbsd/arm)
opearm64 (openbsd/arm64)

pla386 (plan9/386)
pla64 (plan9/amd64)
plaarm (plan9/arm)

sol64 (solaris/amd64)

win386 (windows/386)
win64 (windows/amd64)
winarm (windows/arm)

*/

var (
	targetsFlag    string
	targets        []string
	source, output string
)

func main() {
	fmt.Println("go-builder v1.1 by github.com/NinoM4ster" + "\n")
	flag.StringVar(&targetsFlag, "t", "", "Target platforms: <win64[,lin386[,linarm64]]>")
	flag.StringVar(&source, "s", "", "Source: <main.go>")
	flag.StringVar(&output, "o", "", "Output: <bin/MyApp_%>")
	flag.Parse()
	if targetsFlag == "" || source == "" || output == "" {
		flag.Usage()
		os.Exit(1)
	}
	if !strings.Contains(output, "%") {
		fmt.Println("Output must contain % for replacing it with platform and arch.")
		os.Exit(1)
	}
	targets = strings.Split(targetsFlag, ",")
	for _, a := range targets {
		goos, goarch, err := fetchOsArch(a)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = build(goos, goarch, source, strings.ReplaceAll(output, "%", goos+"-"+goarch))
		if err != nil {
			fmt.Println("\nError when compiling.", err)
			os.Exit(2)
		}
	}
	fmt.Println("\nBuilding complete!")
}

func build(goos, goarch, source, output string) error {
	if goos == "windows" {
		output = output + ".exe"
	}
	fmt.Print("Building '" + source + "' into '" + output + "'...")
	err := os.Setenv("GOOS", goos)
	if err != nil {
		return err
	}
	err = os.Setenv("GOARCH", goarch)
	if err != nil {
		return err
	}
	cmdOutput, err := exec.Command("go", "build", "-o", output, source).CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + "\n" + string(cmdOutput))
	}
	fmt.Println(" done!")
	return nil
}

func fetchOsArch(code string) (string, string, error) {
	switch code {
	case "aixppc64":
		return "aix", "ppc64", nil
	case "and386":
		return "android", "386", nil
	case "and64":
		return "android", "amd64", nil
	case "andarm":
		return "android", "arm", nil
	case "andarm64":
		return "android", "arm64", nil
	case "dar64":
		return "darwin", "amd64", nil
	case "dararm64":
		return "darwin", "arm64", nil
	case "dra64":
		return "dragonfly", "amd64", nil
	case "fre386":
		return "freebsd", "386", nil
	case "fre64":
		return "freebsd", "amd64", nil
	case "frearm":
		return "freebsd", "arm", nil
	case "frearm64":
		return "freebsd", "arm64", nil
	case "ill64":
		return "illumos", "amd64", nil
	case "jswasm":
		return "js", "wasm", nil
	case "lin386":
		return "linux", "386", nil
	case "lin64":
		return "linux", "amd64", nil
	case "linarm":
		return "linux", "arm", nil
	case "linarm64":
		return "linux", "arm64", nil
	case "linmips":
		return "linux", "mips", nil
	case "linmips64":
		return "linux", "mips64", nil
	case "linmips64le":
		return "linux", "mips64le", nil
	case "linmipsle":
		return "linux", "mipsle", nil
	case "linppc64":
		return "linux", "ppc64", nil
	case "linppc64le":
		return "linux", "ppc64le", nil
	case "linriscv64":
		return "linux", "riscv64", nil
	case "lins390x":
		return "linux", "s390x", nil
	case "net386":
		return "netbsd", "386", nil
	case "net64":
		return "netbsd", "amd64", nil
	case "netarm":
		return "netbsd", "arm", nil
	case "netarm64":
		return "netbsd", "arm64", nil
	case "ope386":
		return "openbsd", "386", nil
	case "ope64":
		return "openbsd", "amd64", nil
	case "opearm":
		return "openbsd", "arm", nil
	case "opearm64":
		return "openbsd", "arm64", nil
	case "pla386":
		return "plan9", "386", nil
	case "pla64":
		return "plan9", "amd64", nil
	case "plaarm":
		return "plan9", "arm", nil
	case "sol64":
		return "solaris", "amd64", nil
	case "win386":
		return "windows", "386", nil
	case "win64":
		return "windows", "amd64", nil
	case "winarm":
		return "windows", "arm", nil
	}
	return "", "", errors.New("unknown target  " + code)
}
