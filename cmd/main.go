package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-study-lab/go-mall/cmd/commands"
	"github.com/go-study-lab/go-mall/common/enum"

	"github.com/go-study-lab/go-mall/common/logger"
	"github.com/go-study-lab/go-mall/common/util"
)

func main() {
	app := &cli.App{
		Name:  "gm-tools",
		Usage: "Go Mall 工具集",
		Before: func(c *cli.Context) error {
			// 为每个命令创建带有追踪信息的上下文
			ctx := context.Background()
			spanId := util.GenerateSpanID(util.GetLocalIP())
			ctx = context.WithValue(ctx, "traceid", spanId)
			ctx = context.WithValue(ctx, "spanid", spanId)
			c.Context = ctx

			cmdName := strings.Join(c.Args().Slice(), " ")
			logger.Info(ctx, fmt.Sprintf("定时任务【%s】开始执行. 时间=【%s】)", cmdName, time.Now().Format(enum.TimeFormatHyphenedYMDHIS)))
			return nil
		},
		After: func(c *cli.Context) error {
			cmdName := strings.Join(c.Args().Slice(), " ")
			// 记录执行的错误
			if c.Context.Err() != nil {
				logger.Error(c.Context, fmt.Sprintf("定时任务【%s】执行失败", cmdName), "err", c.Context.Err())
				return c.Context.Err()
			}

			logger.Info(c.Context, fmt.Sprintf("定时任务【%s】执行完成. 时间=【%s", cmdName, time.Now().Format(enum.TimeFormatHyphenedYMDHIS)))
			return nil
		},
		Commands: []*cli.Command{
			commands.Word,
			// 添加更多工具命令
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
