package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// 帮助信息模板
	cli.AppHelpTemplate = fmt.Sprintf(`%s
AUTHOR : TigaTeam

GitHub: https://github.com/tigateam/mebius
`, cli.AppHelpTemplate)

	app := &cli.App{
		Name:  "version",
		Usage: "show current binary version info | 显示当前版本信息",
		Action: func(c *cli.Context) error {
			fmt.Printf(`Mebius v0.0.3 功能正在开发中，敬请期待。。。`)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
