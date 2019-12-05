// Simple Raw Rendering
// Context Interaction
// - renderFrom -- what field (string) to write to resBody
// - resBody -- this plugin writes resBody


package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    context["resBody"] = renderData
    context["resStatus"] = 200
    color.Blue(context["resBody"].(string))
    return context
  }
}
