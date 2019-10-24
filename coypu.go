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
    input, err := p.Lookup("Input")
    if err != nil {
    	panic(err)
    }
    var query = map[string]string{}
    var context = map[string]string{}
    var body = ""
    var method = "GET"
    x, ok := input.(func(map[string]string, map[string]string, string, string) map[string]string)
    if !ok {
      panic("OH NO")
    }
    x(context, query, method, body)
    color.Green("[STARTUP] Starting Server")

    color.Green("[STARTUP] Up at")
}
