package main

import "github.com/fatih/color"
import "plugin";

func main() {

    color.Green("[STARTUP] Reading Config")

    color.Green("[STARTUP] Fetching Plugins")
    p, err := plugin.Open("./config/plugins/sampleData.so")
    if err != nil {
      panic(err)
    }
    //mod, err = plugin.open("config/plugins/filter.so")
    //out, err = plugin.open("config/plugins/json.so")
    //auth, err = plugin.open("config/plugins/json.so")
    pl, err := p.Lookup("New")
    if err != nil {
    	panic(err)
    }
    var context = map[string]interface{}{}
    x, ok := pl.(func(map[string]interface{}) func(map[string]interface{}) map[string]interface{})
    if !ok {
      panic("OH NO")
    }
    x(context)(context)
    color.Green("[STARTUP] Starting Server")

    color.Green("[STARTUP] Up at")
}
