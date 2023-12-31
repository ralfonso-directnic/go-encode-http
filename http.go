package encodehttp

import(
 "net/http"
 "encoding/json"
 "io/ioutil"
 "io"
)


/*This Wrapper Allows you to handle the functions that are in an http request struct without panicing*/

type HttpRequest struct {
	// this assumes Body is always a string
	Body    string `json:"Body,omitempty"`
	// the types of variables `GetBody` and `Cancel` do not
	// really matter since I don't want them in my final json output
	GetBody string `json:"GetBody,omitempty"`
	Cancel  string `json:"Cancel,omitempty"`
	*http.Request
}


func Encode(r *http.Request) (*HttpRequest,error) {

  wrapper := &HttpRequest{Request: r}

  body, rerr := ioutil.ReadAll(r.Body)

  if rerr != nil { 
    return wrapper,rerr
  }

  wrapper.Body = string(body)

  return wrapper,nil

}


func Decode(req_reader io.Reader) (*http.Request,error) { 


  var req http.Request

  perr := json.NewDecoder(req_reader).Decode(&req)

  if perr!=nil { 
   return &req,perr
  }

  return &req,nil


}
