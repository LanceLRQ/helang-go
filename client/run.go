package client

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"helang-go/helang"
	"helang-go/helang/core"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const helpText = `:help  Print this help message
:exit  Exit the shell
:env   Print current environments`

var exitKeyPressTimes = 0

func runCodes(ctx *cli.Context) error {
	file := ctx.Args().Get(0)
	if file == "" {
		return cli.ShowCommandHelp(ctx, "build")
	}

	absPath, err := fileValidator(file)
	if err != nil {
		return err
	}

	helangFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	lexer := helang.NewLex(helangFile)
	tokens, err := lexer.Lex()
	if err != nil {
		log.Fatal(err)
	}

	parser := helang.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	env := map[string]*core.U8{}
	_, err = ast.Evaluate(env)
	if err != nil {
		return fmt.Errorf("%w: %s", core.HeLangException, err.Error())
	}

	return nil
}

func shellKeywordProc(keyword string, env map[string]*core.U8, arguments... interface{}) {
	switch keyword {
	case "exit":
		fmt.Println("Saint He bless you.")
		os.Exit(0)
	case "env":
		for k, v := range env {
			fmt.Printf("%s: %s\n", k, v)
		}
		break
	case "help":
		fmt.Println(helpText)
		break
	default:
		fmt.Printf("Unknown shell keyword: %s\n", keyword)
	}
}

func exitHandler(exitChan chan os.Signal) {
	for {
		select {
		case _ = <-exitChan:
			exitKeyPressTimes++
			if exitKeyPressTimes > 1 {
				fmt.Println("")
				os.Exit(0)
			} else {
				fmt.Print("\n(To exit, press ^C again or ^D or type :exit)\nSpeak to Saint He > ")
			}
		}
	}

}

func runShell(ctx *cli.Context) error {

	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	go exitHandler(exitChan)

	env := map[string]*core.U8{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Speak to Saint He > ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(io.EOF, err) {
				fmt.Println()
				return nil
			} else {
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}
		}

		exitKeyPressTimes = 0

		cmdString = strings.TrimSuffix(cmdString, "\n")
		if cmdString == "" { continue }

		if cmdString[0] == ':' {
			shellKeywordProc(cmdString[1:], env)
			continue
		}

		if cmdString[len(cmdString) - 1] != ';' {
			cmdString += ","
		}

		lexer := helang.NewLex([]byte(cmdString))
		tokens, err := lexer.Lex()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		parser := helang.NewParser(tokens)
		ast, err := parser.Parse()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		_, err = ast.Evaluate(env)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}