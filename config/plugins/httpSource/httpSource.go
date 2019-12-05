// http interface for Coypu
// Plugin config
// - urlFieldName -- what context field to use for url fiels
// Context Interaction
// - (as set by config urlFieldName above) -- which url to use
// - httpStatus -- 200 if ok, 500 if error. (todo -- use source heder info)
// - httpRes -- response or error text


package main

import "io/ioutil"
import "net/http"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    // get content from whatever is set as httpUrl
    var client http.Client
    resp, err := client.Get(context[config["urlFieldName"].(string)].(string))
    if err != nil {
      context["httpStatus"] = 500
      context["httpRes"] = string(err.Error())
      context["error"] = string(err.Error())
    } else {
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        context["httpStatus"] = 500
        context["httpRes"] = string(err.Error())
        context["error"] = string(err.Error())
      }
      context["httpStatus"] = 200
      context["httpRes"] = string(body)
    }
    return context
  }
}
