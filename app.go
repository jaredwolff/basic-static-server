// +build darwin freebsd netbsd openbsd

package main

import (
		"github.com/urfave/cli/v2"
    "github.com/julienschmidt/httprouter"
    "net/http"
		"log"
		"fmt"
		"os"
)

func main() {

  app := &cli.App{
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "port",
        Value: "3001",
        Usage: "port for the server (default 3001)",
			},
			&cli.StringFlag{
        Name: "dir",
        Value: "/var/www/",
        Usage: "directory where to serve (default /var/www/)",
      },
    },
    Action: func(c *cli.Context) error {
			directory := c.String("dir")
			port :=c.String("port")

			fmt.Println(directory)
			fmt.Println(port)

			router := httprouter.New()
			router.ServeFiles("/*filepath",http.Dir(directory))

			log.Fatal(http.ListenAndServe(":3001", router))

      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }

}