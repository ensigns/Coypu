// filter a map based on a field
// UNFINISHED - Do not use as is. Put in an issue if you're interested in using this concept.

package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] filterMap")
    // TODO!!
    var renderFrom string = context["renderFrom"].(string)
    var renderData string = context[renderFrom].(string)
    context["resBody"] = renderData
    return context
  }
}
