package services

import (
	"fmt"
	"github.com/astaxie/beego"
)

func DecodeError(
	controller *beego.Controller,
	err error,
) {
	controller.Data["json"] = map[string]interface{}{
		"error": fmt.Sprintf("json decode error: %v", err),
	}
	controller.ServeJSON()
}
