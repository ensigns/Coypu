// Converts a json string to a map[string]interface{}
// Context Interaction
// - m2jFrom -- what context field to convert to a map
// - m2jTo -- what context field in which to put the map

package main

import "encoding/json"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    var convertFrom string = context["m2jFrom"].(string)
    var convertTo string = context["m2jTo"].(string)
    var msgMapTemplate interface{}
    err := json.Unmarshal([]byte(context[convertFrom].([]byte)), &msgMapTemplate)
    if err != nil {
      context["error"] = string(err.Error())
    }
    msgMap := msgMapTemplate.(map[string]interface{})
    context[convertTo] = msgMap
    return context
  }
}
