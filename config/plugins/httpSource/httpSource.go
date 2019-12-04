package main

import "github.com/fatih/color"
import "io/ioutil"
import "net/http"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] httpSource")
    // get content from whatever is set as httpUrl
    var client http.Client
    resp, err := client.Get(context[config["urlFieldName"].(string)].(string))
    if err != nil {
      context["httpStatus"] = 500
      context["httpRes"] = string(err.Error())
      color.Red(string(err.Error()))
      context["error"] = string(err.Error())
    } else {
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        context["httpStatus"] = 500
        context["httpRes"] = string(err.Error())
        color.Red(string(err.Error()))
        context["error"] = string(err.Error())
      }
      context["httpStatus"] = 200
      context["httpRes"] = string(body)
    }
    return context
  }
}
