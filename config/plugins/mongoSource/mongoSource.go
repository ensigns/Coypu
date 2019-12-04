package main

import "github.com/fatih/color"
import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "time"
import "context"

func New(config map[string]interface{}) func(coypuContext map[string]interface{}) map[string]interface{} {
  MongoClient, Err := mongo.NewClient(options.Client().ApplyURI(config["mongoUri"].(string)))
  return func(coypuContext map[string]interface{}) map[string]interface{} {
    if Err!=nil{
      coypuContext["mongoStatus"] = 500
      coypuContext["mongoRes"] = string(Err.Error())
      coypuContext["error"] = string(Err.Error())
    }
    color.Red("[PKG] mongoSource")
    // get content from whatever is set as httpUrl
    mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := MongoClient.Connect(mongoCtx)
    if err != nil {
      coypuContext["mongoStatus"] = 500
      coypuContext["mongoRes"] = string(err.Error())
      coypuContext["error"] = string(err.Error())
    } else {
      collection := MongoClient.Database(coypuContext["mongoDb"].(string)).Collection(coypuContext["mongoCollection"].(string))
      // TODO are we allowed to write and Is this a post?
      if 1==2 {
        // TODO assemble body
        res, err := collection.InsertOne(mongoCtx, bson.M{"name": "pi", "value": 3.14159})
        if err != nil{
          coypuContext["mongoStatus"] = 500
          coypuContext["mongoRes"] = string(err.Error())
          coypuContext["error"] = string(err.Error())
        } else {
          coypuContext["mongoStatus"] = 200
          coypuContext["mongoRes"] = string(res.InsertedID.(string))
        }
      } else {
        // TODO assemble query
        var results []*map[string]interface{}
        cur, err := collection.Find(mongoCtx, bson.D{})
        if err != nil{
          coypuContext["mongoStatus"] = 500
          coypuContext["mongoRes"] = string(err.Error())
          coypuContext["error"] = string(err.Error())
        } else {
          cur.All(mongoCtx, &results)
          coypuContext["mongoStatus"] = 200
          coypuContext["mongoRes"] = results
        }
      }
    }
    return coypuContext
  }
}
