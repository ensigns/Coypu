// Renames context fields
// Context Interaction
// - renameFrom -- comma separated list of things to rename
// - renameTo -- comma separated list of what to rename to

package main

import "strings"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    s := strings.Split(context["renameFrom"].(string), ",")
    t := strings.Split(context["renameTo"].(string), ",")
    if (len(s)!=len(t)){
      context["error"] = "context renamer, fields must be of the same number of elements"
    } else {
      for i := 0; i < len(s); i++ {
        context[t[i]] = context[s[i]]
      }
    }
    return context
  }
}
