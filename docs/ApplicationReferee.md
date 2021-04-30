# ApplicationReferee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**SessionId** | Pointer to **string** | Integration ID of the session in which the customer redeemed the referral | 
**AdvocateIntegrationId** | Pointer to **string** | Integration ID of the Advocate&#39;s Profile | 
**FriendIntegrationId** | Pointer to **string** | Integration ID of the Friend&#39;s Profile | 
**Code** | Pointer to **string** | Advocate&#39;s referral code. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment the customer redeemed the referral | 

## Methods

### GetApplicationId

`func (o *ApplicationReferee) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationReferee) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationReferee) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationReferee) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetSessionId

`func (o *ApplicationReferee) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationReferee) GetSessionIdOk() (string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *ApplicationReferee) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *ApplicationReferee) SetSessionId(v string)`

SetSessionId gets a reference to the given string and assigns it to the SessionId field.

### GetAdvocateIntegrationId

`func (o *ApplicationReferee) GetAdvocateIntegrationId() string`

GetAdvocateIntegrationId returns the AdvocateIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateIntegrationIdOk

`func (o *ApplicationReferee) GetAdvocateIntegrationIdOk() (string, bool)`

GetAdvocateIntegrationIdOk returns a tuple with the AdvocateIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdvocateIntegrationId

`func (o *ApplicationReferee) HasAdvocateIntegrationId() bool`

HasAdvocateIntegrationId returns a boolean if a field has been set.

### SetAdvocateIntegrationId

`func (o *ApplicationReferee) SetAdvocateIntegrationId(v string)`

SetAdvocateIntegrationId gets a reference to the given string and assigns it to the AdvocateIntegrationId field.

### GetFriendIntegrationId

`func (o *ApplicationReferee) GetFriendIntegrationId() string`

GetFriendIntegrationId returns the FriendIntegrationId field if non-nil, zero value otherwise.

### GetFriendIntegrationIdOk

`func (o *ApplicationReferee) GetFriendIntegrationIdOk() (string, bool)`

GetFriendIntegrationIdOk returns a tuple with the FriendIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFriendIntegrationId

`func (o *ApplicationReferee) HasFriendIntegrationId() bool`

HasFriendIntegrationId returns a boolean if a field has been set.

### SetFriendIntegrationId

`func (o *ApplicationReferee) SetFriendIntegrationId(v string)`

SetFriendIntegrationId gets a reference to the given string and assigns it to the FriendIntegrationId field.

### GetCode

`func (o *ApplicationReferee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplicationReferee) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *ApplicationReferee) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *ApplicationReferee) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetCreated

`func (o *ApplicationReferee) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationReferee) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationReferee) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationReferee) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


