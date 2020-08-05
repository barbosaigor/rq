// Package rq (or requester) is a lightweight REST (over HTTP/HTTPS) request library.
// The main goal of this library is to offer simplicity and ease of use.
//
//
// e.g Fetch products using Get method
//	var products Products
//	rq.Endpoint("my-api.com/products").Get().ToJSON(&products)
//
// e.g create a product using Post method
//	product := Product{...}
//	rq.Endpoint("my-api.com/product").JSON(product).Post()
//
// RQ has an interesting error handling if an operation within the pipeline fails,
// then all subsequent operations will silently forward the error, and no operations are done.
// For error handling, you can use Err which stores the last error entry inside the pipeline.
//     if rq.Endpoint("my-api.com/products").Get().Err != nil {
//        ...
//     }
package rq
