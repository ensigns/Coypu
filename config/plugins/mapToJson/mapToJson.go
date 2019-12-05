// Converts a map to a json string
// Context Interaction
// - m2jFrom -- what context field to convert to a json string
// - m2jTo -- what context field in which to put the json string

package main

import "github.com/fatih/color"
import "encoding/json"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] mapToJson")
    var convertFrom string = context["j2mFrom"].(string)
    var convertTo string = context["j2mTo"].(string)
    jsonString, err := json.Marshal(context[convertFrom])
    if err != nil {
      context["error"] = string(err.Error())
    }
    context[convertTo] = jsonString
    return context
  }
}
