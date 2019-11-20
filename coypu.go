package main

import "github.com/fatih/color"
import "plugin";
import "gopkg.in/yaml.v2";
import "io/ioutil";
import "net/http"
import "log"
import "fmt"
import "os"
import "time"
import "strings"

type PluginConf struct {
    Plugins map[string] struct{
        Path string
        Config map[string]interface{}
      }
}
type RouteConf struct {
    Routes map[string] struct{
        Path string
        Config map[string]interface{}
        Plugins []string
      }
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

func getRoutes()  map[string]func(map[string]interface{})map[string]interface{}{
  var route_conf RouteConf
  routeConfRaw, err := ioutil.ReadFile("./config/routes.yaml")
  if err != nil {
      panic(err)
  }
  err = yaml.Unmarshal(routeConfRaw, &route_conf)
  if err != nil {
      panic(err)
  }
  Routes := make(map[string]func(map[string]interface{})map[string]interface{},0)
  return Routes
}

func main() {
    color.Green("[STARTUP] Setting Up Plugins")
    var Plugins = getPlugins()
    color.Green("[STARTUP] Setting Up Routes")
    var Routes = getRoutes()
    // TODO routes, using route_conf variable
    // check if the url prefix matches
    // function to generate initial context
    // pass to list in this route
    // function to return output from context.
    var http_port = getEnv("port", "8080")
    color.Green("[STARTUP] Up on port " + http_port)
    var handler = func(w http.ResponseWriter, r *http.Request) {
      if(strings.HasPrefix(r.URL.Path, "/err")){
        w.Header().Set("snail", "SNAIL")
        http.Error(w, "snails are seriously everywhere", http.StatusInternalServerError)
      } else {
        fmt.Fprintf(w, "Hello, %q", r.URL.Path)
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
