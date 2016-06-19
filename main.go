package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "agemplate"
	app.Version = Version
	app.Usage = "generate agenda template in date directory"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "template, t",
			Value: "template.md",
			Usage: "path of agenda template",
		},
	}

	app.Action = func(c *cli.Context) error {
		arg := c.Args().Get(0)
		dir, agenda := divide(arg)

		mkdir(dir)

		tamplate := c.String("template")
		copyy(tamplate, agenda)

		return nil
	}
	app.Run(os.Args)
}

func mkdir(path string) {
	if err := os.MkdirAll(path, os.ModeDir|0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
}

func copyy(tamplate, agenda string) {

	from_f, err := os.Open(tamplate)
	if err != nil {
		panic(err)
	}
	defer from_f.Close()

	if _, err := os.Stat(agenda); err == nil {
		fmt.Println(agenda + " is alreadly exist")
		return
	}

	to_f, err := os.Create(agenda)
	if err != nil {
		panic(err)
	}
	defer to_f.Close()

	_, err = io.Copy(to_f, from_f)
	if err != nil {
		panic(err)
	}
}

func divide(arg string) (string, string) {
	dir_path := ""
	agenda := ""
	if arg == "" {
		year, month, day := time.Now().Date()
		dir_path = strconv.Itoa(year) + "/" + strconv.Itoa(int(month))
		agenda = dir_path + "/" + strconv.Itoa(day) + ".md"
	} else {
		path := strings.Split(arg, "/")
		for _, v := range path[:len(path)-1] {
			dir_path = dir_path + v + "/"
		}
		agenda = arg
	}
	return dir_path, agenda
}
