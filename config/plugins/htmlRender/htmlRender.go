package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] htmlRender")
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    if context["error"] != nil {
      context["resBody"] = "<h1>" + renderData + "</h1>"
    } else {
      context["resBody"] = "<h1>" + context["error"].(string) + "</h1>"
    }

    return context
  }
}
