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
	app.Usage = "generate agenda template in date directory"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "template, t",
			Value: "template.md",
			Usage: "path of agenda template",
		},
	}

	app.Action = func(c *cli.Context) error {

		from_file_name := c.String("template")
		to_file_name := c.Args().Get(0)

		dir_path := ""

		if to_file_name == "" {
			t := time.Now()
			year, month, day := t.Date()
			dir_path = strconv.Itoa(year) + "/" + strconv.Itoa(int(month))
			to_file_name = dir_path + "/" + strconv.Itoa(day) + ".md"
		} else {
			path := strings.Split(to_file_name, "/")
			for _, v := range path[:len(path)-1] {
				dir_path = dir_path + v + "/"
			}
		}

		if err := os.MkdirAll(dir_path, os.ModeDir|0755); err != nil {
			if !os.IsExist(err) {
				panic(err)
			}
		}

		from_f, err := os.Open(from_file_name)
		if err != nil {
			panic(err)
		}
		defer from_f.Close()

		if _, err := os.Stat(to_file_name); err == nil {
			fmt.Println(to_file_name + " is alreadly file exist")
			return nil
		}

		to_f, err := os.Create(to_file_name)
		if err != nil {
			panic(err)
		}
		defer to_f.Close()

		_, err = io.Copy(to_f, from_f)
		if err != nil {
			panic(err)
		}

		return nil
	}

	app.Run(os.Args)
}
