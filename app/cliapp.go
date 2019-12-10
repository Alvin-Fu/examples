package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/config/source"
	gcli "github.com/micro/go-micro/config/source/cli"
	"os"
)

func main() {
	var src source.Source
	withContext := false

	// setup app
	app := cmd.App()
	app.Name = "testapp"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "db-host"},
	}

	// with context
	if withContext {
		// set action
		app.Action = func(c *cli.Context) {
			src = gcli.WithContext(c)
		}

		// run app
		app.Run([]string{"run", "-db-host", "localhost"})
		// no context
	} else {
		// set args
		os.Args = []string{"run", "-db-host", "localhost"}
		src = gcli.NewSource()
	}

	// test config
	c, err := src.Read()
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	var actual map[string]interface{}
	if err := json.Unmarshal(c.Data, &actual); err != nil {
		fmt.Errorf("%v", err)
		return
	}

	actualDB := actual["db"].(map[string]interface{})
	if actualDB["host"] != "localhost" {
		fmt.Errorf("%v", err)
		return
	}

}
