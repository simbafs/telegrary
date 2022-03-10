package util

import (
	"fmt"
	"github.com/simba-fs/telegrary/config"
)

func Path(year, month, day string) string {
	return fmt.Sprintf("%s/%s/%s/%s.md", config.Config.Root, year, month, day)
}
