package main

import (
	"github.com/89z/rosso/crypto"
	"github.com/89z/rosso/http"
	"io"
	"net/url"
	"strings"
	"fmt"
	"encoding/json"
	"os"
 )
var Client = http.Default_Client

type Response struct {
   *http.Response
}

type Body struct {
	SID string
	LSID string
	Token string
	Services string
 }
 
func New_Auth(email, password string) (*Response, error) {
   req_body := url.Values{
      "Email": {email},
      "Passwd": {password},
      "client_sig": {""},
      "droidguard_results": {"."},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth",
      strings.NewReader(req_body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   hello, err := crypto.Parse_JA3(crypto.Android_API_26)
   if err != nil {
      return nil, err
   }
   tr := crypto.Transport(hello)
   res, err := Client.Transport(tr).Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}
func main() {
   email := os.Args[1]
   password := os.Args[2]
   output := "auth.json"
   if len(os.Args) < 3 {
      fmt.Println("Usage: aas <email> <password> [output]")
      os.Exit(1)
   }
   if len(os.Args) == 4 {
      output = os.Args[3]
   }
   
   res, err := New_Auth(email, password)
   if err != nil {
	  panic(err)
   }
   defer res.Body.Close()
   _body, _ := io.ReadAll(res.Body)
   fmt.Println(string(_body))

   lines := strings.Split(string(_body), "\n")
   data := make(map[string]string)
   for _, line := range lines {
	   v := strings.Split(line, "=")
	   data[v[0]] = v[1]
   }
   body := Body{
	   SID: data["SID"],
	   LSID: data["LSID"],
	   Token: data["Token"],
	   Services: data["services"],
	}
	json, _ := json.Marshal(body)
	os.WriteFile(output, json, 0644)
}