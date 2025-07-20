package logic

import (
	"context"

	"github.com/go-study-lab/go-mall/common/logger"
)

type WordLogic struct {
	ctx context.Context
}

func NewWordLogic(ctx context.Context) *WordLogic {
	return &WordLogic{
		ctx: ctx,
	}
}

func (l *WordLogic) ParseWord(input, output string) error {
	log := logger.New(l.ctx)
	log.Info("开始解析Word文档", "input", input, "output", output)
	log.Info("解析文档中", "input", input, "output", output)
	log.Info("解析文档完成", "input", input, "output", output)
	return nil
}
