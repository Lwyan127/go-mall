package commands

import (
	"github.com/go-study-lab/go-mall/cmd/logic"
	"github.com/urfave/cli/v2"
)

var Word = &cli.Command{
	// 这个命令是样式模板, 以后其他命令都可以这么添加
	// export ENV=dev;go run cmd/main.go word parse -i input.docx -o output.txt
	// export ENV=dev;go build -o gm-tools cmd/main.go;./gm-tools word parse -i input.docx -o output.txt
	Name:    "word",
	Aliases: []string{"w"},
	Usage:   "Word文档处理相关命令",
	Subcommands: []*cli.Command{
		{
			Name:  "parse",
			Usage: "解析Word文档",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "input",
					Aliases:  []string{"i"},
					Usage:    "输入文件路径",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "output",
					Aliases:  []string{"o"},
					Usage:    "输出文件路径",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				return logic.NewWordLogic(c.Context).ParseWord(c.String("input"), c.String("output"))
			},
		},
	},
}
