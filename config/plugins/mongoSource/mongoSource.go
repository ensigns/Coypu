// SQLITE interface for Coypu
// UNTESTED -- Hopefully not for long
// Plugin config
// - mongoUri -- what mongo instance to open
// Context Interaction
// - mongoAllowWrite -- if this route should be able to write
// - mongoDb -- which db to use
// - mongoCollection -- which collection to use
// - mongoInsertBody -- data to insert into mongo
// - mongoStatus -- 500 if error, 200 if ok
// - mongoRes -- some indication of the result/error, or results


package main

import (
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "time"
   "context"
)

func New(config map[string]interface{}) func(coypuContext map[string]interface{}) map[string]interface{} {
  return func(coypuContext map[string]interface{}) map[string]interface{} {
    MongoClient, Err := mongo.NewClient(options.Client().ApplyURI(config["mongoUri"].(string)))
    if Err!=nil{
      coypuContext["mongoStatus"] = 500
      coypuContext["mongoRes"] = string(Err.Error())
      coypuContext["error"] = string(Err.Error())
    }
    mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := MongoClient.Connect(mongoCtx)
    if err != nil {
      coypuContext["mongoStatus"] = 500
      coypuContext["mongoRes"] = string(err.Error())
      coypuContext["error"] = string(err.Error())
    } else {
      collection := MongoClient.Database(coypuContext["mongoDb"].(string)).Collection(coypuContext["mongoCollection"].(string))
      // TODO are we allowed to write and Is this a post?
      if (coypuContext["method"]=="POST" && coypuContext["mongoAllowWrite"] !=nil){
        mongoBody := coypuContext["mongoInsertBody"].(map[string]interface{});
        res, err := collection.InsertOne(mongoCtx, &mongoBody)
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
        var results []map[string]interface{}
        cur, err := collection.Find(mongoCtx, bson.D{})
        _ = cur
        if err != nil{
          coypuContext["mongoStatus"] = 500
          coypuContext["mongoRes"] = string(err.Error())
          coypuContext["error"] = string(err.Error())
        } else {
          cur.All(mongoCtx, &results)
          coypuContext["mongoStatus"] = 200
          // debug, look at results
          coypuContext["mongoRes"] = results
        }
      }
    }
    return coypuContext

  }
}
