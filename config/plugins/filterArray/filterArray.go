// filter an array of maps based on a string field in each map
// context interaction
// - filterFrom
// - filterOut
// - filterMatchList - comma separated string of items to match
// - filterField -- field in map item to test; should be string

package main
import "strings"

import "github.com/fatih/color"

func stringInSlice(a string, list []string) bool {
  for _, b := range list {
      if b == a {
          return true
      }
  }
  return false
}

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    var filtered []map[string]interface{}
    filterMatchList := strings.Split(context["filterMatchList"].(string), ",")
    filterFromList := context[context["filterFrom"].(string)].([]map[string]interface{})
    for _, i := range filterFromList{
      if stringInSlice(i[context["filterField"].(string)].(string), filterMatchList) {
        filtered = append(filtered, i)
      }
    }
    context[context["filterOut"].(string)] = filtered
    return context
  }
}
