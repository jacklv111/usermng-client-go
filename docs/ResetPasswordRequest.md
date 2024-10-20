# ResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**NewPassword** | **string** |  | 
**ConfirmPassword** | **string** |  | 

## Methods

### NewResetPasswordRequest

`func NewResetPasswordRequest(username string, password string, newPassword string, confirmPassword string, ) *ResetPasswordRequest`

NewResetPasswordRequest instantiates a new ResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordRequestWithDefaults

`func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest`

NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ResetPasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResetPasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResetPasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ResetPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ResetPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetNewPassword

`func (o *ResetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetConfirmPassword

`func (o *ResetPasswordRequest) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *ResetPasswordRequest) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *ResetPasswordRequest) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


