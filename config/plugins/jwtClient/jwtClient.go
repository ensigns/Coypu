// JWT interface for Coypu
// config interactions
// - jwksURL -- where to get the jwks
// context interactions
// - jwtClaims -- where to put claims
// - jwt -- the token itself
package main

import  "github.com/fatih/color"
import  "github.com/dgrijalva/jwt-go"
import  "github.com/lestrrat-go/jwx/jwk"
import  "errors"
import  "fmt"

func New(config map[string]interface{}) func(context map[string]interface{}) map[string]interface{} {
  KeySet, JwksErr := jwk.FetchHTTP(config["jwksURL"].(string))
  return func(context map[string]interface{}) map[string]interface{} {
    color.Red("[PKG] jwtClient")
    if JwksErr != nil {
        context["error"] = string(JwksErr.Error())
    }
    token, err := jwt.Parse(context["jwt"].(string), func(token *jwt.Token) (interface{}, error) {
        keyID, ok := token.Header["kid"].(string)
        if !ok {
            return nil, errors.New("expecting JWT header to have string kid")
        }

        if key := KeySet.LookupKeyID(keyID); len(key) == 1 {
            return key[0].Materialize()
        }

        return nil, fmt.Errorf("unable to find key %q", keyID)
    })
    context["jwtClaims"] = token.Claims.(jwt.MapClaims)
    if err != nil {
      context["error"] = string(err.Error())
    }
    return context
  }
}
