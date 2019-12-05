// Reports errors to user and log
// BROKEN FOR UNKNOWN REASON!! -- do not use
// context interaction
// - error -- the error text
// - resBody -- replaces with error text if an error is present
// - resStatus -- replaces with 500 if error is present

package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  color.Blue("0")
  return func(context map[string]interface{}) map[string]interface{} {
    // if there is an error
    color.Blue("1")
    if context["error"] != nil{
      color.Blue("2")
       // report it as json string to user
       context["resBody"] = "{\"err\": \""+context["error"].(string) + "\"}"
       context["resStatus"] = 500
       color.Blue("3")
       // log it
       color.Blue("[ERR] " + context["error"].(string))
    }
    // otherwise don't modify anything

    return context
  }
}
