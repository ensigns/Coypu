// SQLITE interface for Coypu
// UNFINISHED - Do not use as is. Put in an issue if you're interested in using this concept.
// Plugin config
// - sqliteDb -- what file to open for sqlite db.
// Context Interaction
// - sqliteAllowWrite -- if this route should be able to write
// - TODO

package main

import (
  "github.com/fatih/color"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  )

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  Database, _ := sql.Open("sqlite3", config["sqliteDb"].(string))
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] sqliteSource")
    _ = Database
    if (context["method"] == "POST" && context["sqliteAllowWrite"] !=nil){
      var a string;
      _ = a;
      // TODO write
    } else {
      var b string;
      _ = b;
      // TODO read
    }
    return context
  }
}
