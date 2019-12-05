// Simple JSON Rendering
// Context Interaction
// - renderFrom -- what field (string) to write to resBody
// - resBody -- this plugin writes resBody


package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] jsonRender")
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    context["resBody"] = "{\"data\" :" + renderData + "}"
    context["resHeaders"].(map[string]string)["Content-Type"] = "application/json"
    return context
  }
}
