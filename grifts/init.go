package grifts

import (
	"github.com/6uss/k_manage/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
