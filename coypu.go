package main

import (
  "github.com/fatih/color"
   "plugin"
   "gopkg.in/yaml.v2"
   "io/ioutil"
   "net/http"
   "log"
   "fmt"
   "os"
   "time"
   "strings"
)

type PluginConf struct {
    Plugins map[string] struct{
        Path string
        Config map[string]interface{}
      }
}

type SingleRoute struct{
    Config map[string]interface{}
    Plugins []string
  }

type RouteConf struct {
    Routes map[string] SingleRoute
}

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

func getPlugins() map[string]func(map[string]interface{})map[string]interface{} {
  var plugin_conf PluginConf
  pluginConfRaw, err := ioutil.ReadFile("./config/plugins.yaml")
  if err != nil {
      panic(err)
  }
  err = yaml.Unmarshal(pluginConfRaw, &plugin_conf)
  if err != nil {
      panic(err)
  }
  Plugins := make(map[string]func(map[string]interface{})map[string]interface{},0)
  for k, v := range plugin_conf.Plugins {
    color.Yellow("[Plugin] " + k + " at " + v.Path)
    p, err := plugin.Open(v.Path)
    if err != nil {
      panic(err)
    }
    pl, err := p.Lookup("New")
    if err != nil {
      panic(err)
    }
    x, ok := pl.(func(map[string]interface{}) func(map[string]interface{}) map[string]interface{})
    if !ok {
      panic("OH NO")
    }
    Plugins[k] = x(v.Config)
  }
  return Plugins
}

func getRoutes() RouteConf{
  var route_conf RouteConf
  routeConfRaw, err := ioutil.ReadFile("./config/routes.yaml")
  if err != nil {
      panic(err)
  }
  err = yaml.Unmarshal(routeConfRaw, &route_conf)
  if err != nil {
      panic(err)
  }
  //Routes := make(map[string]interface{})
  return route_conf
}

func main() {
    color.Green("[STARTUP] Setting Up Plugins")
    var Plugins = getPlugins()
    _ = Plugins // REMOVE
    color.Green("[STARTUP] Setting Up Routes")
    var Routes = getRoutes().Routes
    var http_port = getEnv("port", "8080")
    color.Green("[STARTUP] Up on port " + http_port)
    var handler = func(w http.ResponseWriter, r *http.Request) {
      var pathArray = strings.Split(r.URL.Path, "/")
      if (len(pathArray)>1){
        var subdir = strings.ToLower(pathArray[1])
        if route, ok := Routes[subdir]; ok {
          var pls = route.Plugins
          for _, pl := range pls {
            fmt.Println(pl)
          }
          // assemble context
          // run the things in the route in order
          // return output from context
          // ERROR HANDLER
        } else {
          // default handler, echo for now
          fmt.Fprintf(w, "Tried to use and could not find route named %q", subdir)
        }
      } else {
        fmt.Fprintf(w, "No subdir, no route selected")
      }




    }
    var s = &http.Server{
    	Addr:           ":"+http_port,
    	Handler:        http.HandlerFunc(handler),
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())

}
