// Unpacks context fields from a map in context
// Context Interaction
// - unpackSource -- what map to use as the unpacking source
// - unpackFrom -- comma separated list of things to unpack
// - unpackTo -- comma separated list of what name contents to

package main

import "github.com/fatih/color"
import "strings"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    src := context["unpackSource"].(map[string]interface{})
    s := strings.Split(context["unpackFrom"].(string), ",")
    t := strings.Split(context["unpackTo"].(string), ",")
    if (len(s)!=len(t)){
      context["error"] = "context renamer, fields must be of the same number of elements"
    } else {
      for i := 0; i < len(s); i++ {
        context[t[i]] = src[s[i]]
      }
    }
    return context
  }
}
