# RQ

Package RQ is a lightweight REST (over HTTP) request library. 
The main goal of this library is to offer simplicity and ease of use.

## Installation
```bash
go get github.com/barbosaigor/rq
```

## Documentation
For more details, you can read the complete documentation on [pkg.go.dev](https://pkg.go.dev/github.com/barbosaigor/rq).  
You can configure your requests with these methods  
**Endpoint** defines the request endpoint (if endpoint hasn't the prefix http://, it will attach it)  
**JSON** sets the JSON to send within the request  
**Text** sets the text to send within the request  
**Cookies** sets Cookies which is used to send within the request  
**Cookie** sets a cookie which is used to send within the request  
**SetHeader** sets fields for HTTP header  

After that, you can make a request using REST verbs on these methods  
**Post**, **Get**, **Put**, **Delete**, **Head**, and **Patch**  

After the request is done you can get the response  
**ToJSON** unmarshal the response body to JSON  
**ToString** convert the response body to string  
**StatusCode** returns the status code response  
**ResponseBody** returns the response body as []byte  
**Response** returns the net/http Response  

## Usage
e.g fetch products using Get method  
```golang
var products Products
rq.Endpoint("my-api.com/products").Get().ToJSON(products)
```  
e.g update products using the Post method  
```golang
product := Product{...}
rq.Endpoint("my-api.com/product").JSON(product).Post()
```  

RQ has an interesting error handling if an operation within the pipeline fails,
then all subsequent operations will silently forward the error, and no operations are done.
For error handling, you can use _Err_ which stores the last error entry inside the pipeline.
```golang
if rq.Endpoint("my-api.com/products").Get().Err != nil {
    ...
}
```  
