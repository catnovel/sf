# SF API

The SF API is a Go package that provides a client for interacting with the SFACG (SF轻小说) website's API. It allows users to retrieve book information, search for novels, and retrieve user account information. The package provides methods for making API requests and handling the responses. Customization options, such as setting a device token or API key, are also available. Contributions to the package are welcome, and it is licensed under the MIT License.

## Installation

To use the SF API package in your Go project, you can use the following command to install it:

```shell
go get github.com/catnovelapi/sf
```

## Usage

To use the SF API package, you need to import it in your Go code:

```go
import "github.com/catnovelapi/sf"
```

### Creating a Client

To start using the SF API, you need to create a client instance. You can customize the client by providing options such as the device token, API key, and API base URL. Here's an example of creating a client with the default options:

```go
client := sf.NewSfClient()
```

### Making API Requests

Once you have a client instance, you can use it to make API requests. The SF API package provides methods for various API endpoints, such as getting book information, searching for novels, and retrieving user information. Here are some examples:

```go
// Get book information
bookInfo := client.GetBookInfoApi(bookId)

// Search for novels
searchResults := client.SearchNovelsResultApi(keyword, page)

// Get user account information
accountInfo := client.GetAccountInApi()
```

### Handling API Responses

The API methods return the response data in the form of `gjson.Result`, which is a type provided by the `github.com/tidwall/gjson` package. You can use the methods provided by `gjson.Result` to extract the desired data from the response. Here's an example:

```go
// Get the book title from the book information response
title := bookInfo.Get("data.title").String()
```

### Customizing Client Options

If you want to customize the client options, such as setting a device token or API key, you can use the provided options when creating the client. Here's an example:

```go
client := sf.NewSfClient(
	sf.DeviceToken("your-device-token"),
	sf.ApiKey("your-api-key"),
)
```

## Contributing

Contributions to the SF API package are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
