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
          // inital context is route config
          var ctx = route.Config
          // add empty header item for response headers
          ctx["resHeaders"] = map[string]string{}
          ctx["resStatus"] = 200
          // empty error
          ctx["error"] = nil
          // add method, string
          ctx["method"] = r.Method
          // add body, string
          body, err := ioutil.ReadAll(r.Body)
          if err != nil {
            log.Println("body extraction error:", err)
            body = make([]byte, 0, 0)
          }
          ctx["body"] = string(body)
          // add query, map
          ctx["query"] = r.URL.Query()
          // add headers, map
          ctx["headers"] = r.Header
          for _, pl := range pls {
            // run the plugin, updating context
            defer func() {
              if err := recover(); err != nil {
                  log.Println("plugin execution error:", err)
              }
            }()
            ctx = Plugins[pl](ctx)
          }
          // write all list of resHeaders
          for hdrName, hdrVal := range ctx["resHeaders"].(map[string]string){
            w.Header().Set(hdrName, hdrVal)
          }
          // status code write
          if resCode, ok := ctx["resStatus"]; ok{
            w.WriteHeader(resCode.(int))
          }
          fmt.Fprintf(w, ctx["resBody"].(string))
          // return output from context
        } else {
          w.WriteHeader(http.StatusInternalServerError)
          fmt.Fprintf(w, "Tried to use and could not find route named %q", subdir)
        }
      } else {
        w.WriteHeader(http.StatusInternalServerError)
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
