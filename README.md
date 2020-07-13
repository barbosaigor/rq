# RQ

Package RQ is a lightweight REST (over HTTP) request library. 
The main goal of this library is to offer simplicity and ease of use.

## Installation
```bash
go get github.com/barbosaigor/rq
```

## Documentation
For more details, you can read the complete documentation on [gopkg](https://pkg.go.dev/github.com/barbosaigor/rq).  
You can configure your requests with these methods  
_Endpoint_ defines the request endpoint (endpoint must contain the prefix http://)  
_JSON_ sets the JSON to send within the request  
_Text_ sets the text to send within the request  
_Cookies_ sets Cookies which is used to send within the request  
_Cookie_ sets a cookie which is used to send within the request  
_SetHeader_ sets fields for HTTP header  

After that, you can make a request using REST verbs on these methods  
_Post_, _Get_, _Put_, _Delete_, _Head_, _Patch_  

After the request is done you can get the response  
_ToJSON_ unmarshal the response body to JSON  
_ToString_ convert the response body to string  
_StatusCode_ returns the status code response  
_ResponseBody_ returns the response body as []byte  
_Response_ returns the net/http Response  

## Usage
e.g fetch products using Get method  
```golang
var products Products
rq.Endpoint("http://my-api.com/products").Get().ToJSON(products)
```  
e.g update products using the Post method  
```golang
product := Product{...}
rq.Endpoint("http://my-api.com/product").JSON(product).Post()
```  

RQ has an interesting error handling if an operation within the pipeline fails,
then all subsequent operations will silently forward the error, and no operations are done.
For error handling, you can use _Err_ which stores the last error entry inside the pipeline.
```golang
if rq.Endpoint("http://my-api.com/products").Get().Err != nil {
    ...
}
```  
