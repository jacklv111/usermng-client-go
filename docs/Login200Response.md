# Login200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The token to be used for subsequent authenticated requests. | [optional] 

## Methods

### NewLogin200Response

`func NewLogin200Response() *Login200Response`

NewLogin200Response instantiates a new Login200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogin200ResponseWithDefaults

`func NewLogin200ResponseWithDefaults() *Login200Response`

NewLogin200ResponseWithDefaults instantiates a new Login200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *Login200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Login200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Login200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Login200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


