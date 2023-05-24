package server

import (
	"fmt"

	"github.com/dadadam/sono-backend/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%d", c.Port))
}
