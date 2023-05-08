package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/sand361049/common/helper"
)

func main(){
	fmt.Println(helper.GreeterText("Text"))
	logger := log.With().Str("version","0.1").Caller().Logger()
	logger.Info().Msg("testing")
}