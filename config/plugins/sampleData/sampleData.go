package main

import "github.com/fatih/color"
import "math/rand"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] sampleData")
    context["random"] = rand.Intn(100)
    context["randomFloat"] = rand.Float64()
    return context
  }
}
