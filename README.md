# Go Akeneo SDK

The Go Akeneo SDK is a client library for interacting with the Akeneo API in Go applications.

## Installation
You can install the library using the go get command:
```bash
go get github.com/ezifyio/go-akeneo
```

To use the Go Akeneo SDK in your Go project, you need to import it:

```go
import "github.com/ezifyio/go-akeneo"
```

## Usage

To get started with the Go Akeneo SDK, you need to create a new client by providing the necessary configuration and options:

```go
connector := goakeneo.Connector{
ClientID:   "your-client-id",
Secret:     "your-secret",
UserName:   "your-username",
Password:   "your-password",
}

client, err := goakeneo.NewClient(connector,
	goakeneo.WithBaseURL("https://your-akeneo-instance.com"),
	goakeneo.WithRateLimit(10, 1*time.Second),
)
if err != nil {
// Handle error
}
```

Once you have a client instance, you can use it to interact with the Akeneo API. The client provides various services for different API endpoints, such as AuthService, ProductService, FamilyService, etc. You can access these services from the client and make API calls:

```go
// Example: Get products
products,links, err := client.Product.ListWithPagination(nil)
if err != nil {
// Handle error
}

for product := range products {
// Process each product
}
```


Refer to the Go Akeneo SDK documentation and API reference for more information on available services and methods.

## Contributing
If you would like to contribute to the Go Akeneo SDK, feel free to submit pull requests or open issues on the GitHub repository: https://github.com/ezifyio/go-akeneo

## License
The Go Akeneo SDK is open-source and available under the MIT License.