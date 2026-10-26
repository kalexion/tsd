package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/znwng/tsd/internal/cli"
	"github.com/znwng/tsd/internal/config"
	"github.com/znwng/tsd/internal/executor"
	"github.com/znwng/tsd/internal/printer"
)

func main() {
	show := flag.Bool("show", false, "show the passed config")
	help := flag.Bool("help", false, "show help")
	attach := flag.String("attach", "", "attach to a session after setup")
	single := flag.Bool("single", false, "works only when exactly one session is defined")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "init":
			if len(args) < 2 {
				fmt.Println("usage: tsd init <filename.hcl>")
				return
			}

			currentDirectory, err := os.Getwd()
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			if filepath.Ext(args[1]) != ".hcl" {
				fmt.Println("error: filename must have a .hcl extension")
				return
			}

			if err := executor.Init(currentDirectory, args[1]); err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			return
		}
	}

	if *help {
		cli.PrintHelp()
		return
	}

	if flag.NArg() == 0 {
		fmt.Println("error: no configuration file provided")
		fmt.Println("usage: tsd [options] <config.hcl>")
		return
	}

	fileName := flag.Arg(0)

	if filepath.Ext(fileName) != ".hcl" {
		fmt.Printf("error: invalid configuration file %q\n", fileName)
		fmt.Println("only .hcl files are accepted")
		return
	}

	var cfg config.Config

	if err := hclsimple.DecodeFile(fileName, nil, &cfg); err != nil {
		fmt.Printf("error: failed to parse configuration file %q\n", fileName)
		fmt.Println(err)
		return
	}

	if *show {
		printer.PrintConfig(cfg)
		return
	}

	if *single {
		if len(cfg.Sessions) != 1 {
			fmt.Println("error: --single flag must be used only when exactly one session is defined")
			return
		}
		*attach = cfg.Sessions[0].Name
	}

	if err := executor.ExecuteConfig(cfg, *attach); err != nil {
		fmt.Printf("error: failed to execute configuration: %v\n", err)
	}
}
