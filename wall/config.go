package wall

import (
	"github.com/paked/configure"
)

var (
	conf        = configure.New()
	mongoString = conf.String("MONGO_STRING", "", "")
)

func init() {
	conf.Use(configure.NewEnvironment())
}
