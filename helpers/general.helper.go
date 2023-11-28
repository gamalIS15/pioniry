package helpers

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReportError(err error) {
	if viper.GetString("ENVIRONMENT") == "Dev" {
		fmt.Println(err)
	}
}
