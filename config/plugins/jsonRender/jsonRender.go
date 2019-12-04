package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] jsonRender")
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)

    context["resHeaders"].(map[string]string)["Content-Type"] = "application/json"
    if context["error"] != nil {
      context["resBody"] = "{\"data\" :\"" + renderData + "}"
    } else {
      context["resBody"] = "{\"err\" :\"" + context["error"].(string) + "\"}"
    }
    return context
  }
}
