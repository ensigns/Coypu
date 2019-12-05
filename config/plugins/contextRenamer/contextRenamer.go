// Renames context fields
// Context Interaction
// - m2jFrom -- what context field to convert to a json string
// - m2jTo -- what context field in which to put the json string

package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] contextRenamer")
    // TODO!!
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    context["resBody"] = renderData
    return context
  }
}
