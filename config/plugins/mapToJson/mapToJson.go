// Converts a map to a json string
// Context Interaction
// - m2jFrom -- what context field to convert to a json string
// - m2jTo -- what context field in which to put the json string

package main

import "encoding/json"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    var convertFrom string = context["m2jFrom"].(string)
    var convertTo string = context["m2jTo"].(string)
    jsonString, err := json.Marshal(context[convertFrom])
    if err != nil {
      context["error"] = string(err.Error())
    }
    context[convertTo] = string(jsonString)
    return context
  }
}
