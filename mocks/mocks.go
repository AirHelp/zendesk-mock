package mocks

import (
	"time"
)

func Id() int {
	return int(time.Now().Unix())
}
