// Adds random numbers to context
// Context Interaction
// - random -- a random integer
// - randomFloat -- a random float

package main

import "math/rand"
import "strconv"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    context["random"] = strconv.FormatInt(rand.Int63n(100), 10)
    context["randomFloat"] = strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
    return context
  }
}
