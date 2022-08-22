package client

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/urfave/cli/v2"
	"helang-go/helang/core"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

//go:embed vm.template
var helangVMFile string

func fileValidator(file string) (string, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("%w: '%s' not found", core.HeLangCompilerException, file)
	} else if err != nil {
		return "", err
	}
	if matched, err := regexp.Match(".he$", []byte(file)); !matched || err != nil {
		return "", fmt.Errorf("%w: '%s' must be end of '.he'", core.HeLangCompilerException, file)
	}
	absPath, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return absPath, err
}

func buildCodes(ctx *cli.Context) error {
	file := ctx.Args().Get(0)
	if file == "" {
		return cli.ShowCommandHelp(ctx, "build")
	}

	absPath, err := fileValidator(file)
	if err != nil {
		return err
	}

	fileName := strings.Replace(file, ".he", "", -1)
	workDir := path.Dir(absPath)


	codeName := fmt.Sprintf("compile_%s.go", fileName)
	targetCodePath := path.Join(workDir, codeName)
	err = ioutil.WriteFile(targetCodePath, []byte(helangVMFile), 0755)
	if err != nil {
		return err
	}
	fmt.Printf("[Build] %s ...\n", file)

	outputName := ctx.String("output")
	if ctx.String("output") == "" {
		outputName = fileName
		if runtime.GOOS == "windows" {
			outputName = fileName + ".exe"
		}
	}

	timeoutCtx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cmd := exec.CommandContext(timeoutCtx, "go", "build", "-o", outputName, codeName)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir
	err = cmd.Start()
	if err != nil { panic(err) }
	err = cmd.Wait()
	if err != nil { panic(err) }

	_ = os.Remove(targetCodePath)

	fmt.Printf("Done. see './%s'\n", outputName)
	return nil
}