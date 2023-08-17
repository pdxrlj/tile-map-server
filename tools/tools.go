package tools

import (
	"github.com/spf13/cast"
)

func StringToInt(s string) int {
	return cast.ToInt(s)
}
