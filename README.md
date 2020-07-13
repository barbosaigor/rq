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
```golang
func (rq *RQ) Endpoint(endpoint string) *RQ
```  
**JSON** sets the JSON to send within the request  
```golang
func (rq *RQ) JSON(data interface{}) *RQ
```  
**Text** sets the text to send within the request  
```golang
func (rq *RQ) Text(data string) *RQ
```  
**Cookies** sets Cookies which is used to send within the request  
```golang
func (rq *RQ) Cookies(cookies []*http.Cookie) *RQ
```  
**Cookie** sets a cookie which is used to send within the request  
```golang
func (rq *RQ) Cookie(cookie *http.Cookie) *RQ
```  
**SetHeader** sets fields for HTTP header  
```golang
func (rq *RQ) SetHeader(name, value string) *RQ
```  

After that, you can make a request using REST verbs on these methods **Post**, **Get**, **Put**, **Delete**, **Head**, and **Patch**  
```golang
func (rq *RQ) Post() *RQ  
```
```golang
func (rq *RQ) Get() *RQ  
```
```golang
func (rq *RQ) Put() *RQ  
```
```golang
func (rq *RQ) Delete() *RQ  
```
```golang
func (rq *RQ) Head() *RQ  
```
```golang
func (rq *RQ) Patch() *RQ  
```  

After the request is done you can get the response  
**ToJSON** unmarshal the response body to JSON  
```golang
func (rq *RQ) ToJSON(data interface{}) *RQ
```  
**ToString** convert the response body to string  
```golang
func (rq *RQ) ToString(str *string) *RQ
```  
**StatusCode** returns the status code response  
```golang
func (rq *RQ) StatusCode() int
```  
**ResponseBody** returns the response body as []byte  
```golang
func (rq *RQ) ResponseBody() []byte
```  
**Response** returns the net/http Response  
```golang
func (rq *RQ) Response() *http.Response
```  

## Usage
e.g fetch products using Get method  
```golang
var products Products
rq.Endpoint("my-api.com/products").Get().ToJSON(products)
```  
e.g create product using the Post method  
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
