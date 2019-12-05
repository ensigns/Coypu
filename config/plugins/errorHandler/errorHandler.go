// Reports errors to user and log
// context interaction
// - error -- the error text
// - resBody -- replaces with error text if an error is present
// - resStatus -- replaces with 500 if error is present

package main

import "github.com/fatih/color"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  return func(context map[string]interface{}) map[string]interface{} {
    // if there is an error
    if context["error"] !=nil{
       // report it as json string to user
       context["resBody"] = "{\"err\": \""+context["error"].(string) + "\"}"
       context["resStatus"] = 500
       // log it
       color.Blue("[ERR] " + context["error"].(string))
    }
    // otherwise don't modify anything

    return context
  }
}
