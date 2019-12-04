package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] rawRender")
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    context["resBody"] = renderData
    return context
  }
}
