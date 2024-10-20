# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTokenInfo

`func NewTokenInfo() *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *TokenInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExpiredAt

`func (o *TokenInfo) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *TokenInfo) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *TokenInfo) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *TokenInfo) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


