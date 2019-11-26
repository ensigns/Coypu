package main

import "github.com/fatih/color"
import "math/rand"
import "strconv"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] sampleData")
    context["random"] = strconv.FormatInt(rand.Int63n(100), 10)
    context["randomFloat"] = strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
    return context
  }
}
