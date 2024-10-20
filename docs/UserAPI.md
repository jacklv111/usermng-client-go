# \UserAPI

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindDevice**](UserAPI.md#BindDevice) | **Post** /device/bind | user bind a device
[**CreateDevice**](UserAPI.md#CreateDevice) | **Post** /device | Create device
[**GetTokenInfo**](UserAPI.md#GetTokenInfo) | **Get** /token-info | Get token info
[**Login**](UserAPI.md#Login) | **Post** /login | User login
[**ResetPassword**](UserAPI.md#ResetPassword) | **Post** /resetpwd | Reset password
[**Signup**](UserAPI.md#Signup) | **Post** /signup | User signup
[**UnbindDevice**](UserAPI.md#UnbindDevice) | **Post** /device/unbind | user unbind a device
[**UpdateUser**](UserAPI.md#UpdateUser) | **Put** /user | Update user



## BindDevice

> BindDevice(ctx).Authorization(authorization).BindDeviceRequest(bindDeviceRequest).Execute()

user bind a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	authorization := "authorization_example" // string | 
	bindDeviceRequest := *openapiclient.NewBindDeviceRequest("DeviceId_example") // BindDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.BindDevice(context.Background()).Authorization(authorization).BindDeviceRequest(bindDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.BindDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **bindDeviceRequest** | [**BindDeviceRequest**](BindDeviceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> CreateDevice201Response CreateDevice(ctx).Authorization(authorization).Execute()

Create device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateDevice(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevice`: CreateDevice201Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**CreateDevice201Response**](CreateDevice201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenInfo

> TokenInfo GetTokenInfo(ctx).Authorization(authorization).Execute()

Get token info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTokenInfo(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTokenInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenInfo`: TokenInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTokenInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> Login200Response Login(ctx).LoginRequest(loginRequest).Execute()

User login



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	loginRequest := *openapiclient.NewLoginRequest("Username_example", "Password_example") // LoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Login(context.Background()).LoginRequest(loginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: Login200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**Login200Response**](Login200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("Username_example", "Password_example", "NewPassword_example", "ConfirmPassword_example") // ResetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ResetPassword(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> Signup201Response Signup(ctx).CreateUserRequest(createUserRequest).Execute()

User signup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	createUserRequest := *openapiclient.NewCreateUserRequest("Username_example", "Password_example") // CreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Signup(context.Background()).CreateUserRequest(createUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Signup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Signup`: Signup201Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**Signup201Response**](Signup201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindDevice

> UnbindDevice(ctx).Authorization(authorization).Execute()

user unbind a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UnbindDevice(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnbindDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnbindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx).Authorization(authorization).UpdateUserRequest(updateUserRequest).Execute()

Update user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/usermng-client-go"
)

func main() {
	authorization := "authorization_example" // string | 
	updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUser(context.Background()).Authorization(authorization).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

