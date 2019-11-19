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

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

func main() {

    color.Green("[STARTUP] Reading Config")

    color.Green("[STARTUP] Fetching Plugins")
    type PluginConf struct {
        Plugins map[string] struct{
            Path string
            Config map[string]interface{}
          }
    }
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
    // TODO routes
    // function to generate initial context
    // pass to list in this route
    // function to return output from context.
    var http_port = getEnv("port", "8080")
    color.Green("[STARTUP] Up on port " + http_port)
    var handler = func(w http.ResponseWriter, r *http.Request) {
      if(r.URL.Path == "/err"){
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
