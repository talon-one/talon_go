# ApplicationRefereeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | Integration ID of the session in which the customer redeemed the referral. | 
**AdvocateIntegrationId** | Pointer to **string** | Integration ID of the Advocate&#39;s Profile. | 
**FriendIntegrationId** | Pointer to **string** | Integration ID of the Friend&#39;s Profile. | 
**Code** | Pointer to **string** | Advocate&#39;s referral code. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment the customer redeemed the referral. | 

## Methods

### GetSessionId

`func (o *ApplicationRefereeAllOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationRefereeAllOf) GetSessionIdOk() (string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *ApplicationRefereeAllOf) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *ApplicationRefereeAllOf) SetSessionId(v string)`

SetSessionId gets a reference to the given string and assigns it to the SessionId field.

### GetAdvocateIntegrationId

`func (o *ApplicationRefereeAllOf) GetAdvocateIntegrationId() string`

GetAdvocateIntegrationId returns the AdvocateIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateIntegrationIdOk

`func (o *ApplicationRefereeAllOf) GetAdvocateIntegrationIdOk() (string, bool)`

GetAdvocateIntegrationIdOk returns a tuple with the AdvocateIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdvocateIntegrationId

`func (o *ApplicationRefereeAllOf) HasAdvocateIntegrationId() bool`

HasAdvocateIntegrationId returns a boolean if a field has been set.

### SetAdvocateIntegrationId

`func (o *ApplicationRefereeAllOf) SetAdvocateIntegrationId(v string)`

SetAdvocateIntegrationId gets a reference to the given string and assigns it to the AdvocateIntegrationId field.

### GetFriendIntegrationId

`func (o *ApplicationRefereeAllOf) GetFriendIntegrationId() string`

GetFriendIntegrationId returns the FriendIntegrationId field if non-nil, zero value otherwise.

### GetFriendIntegrationIdOk

`func (o *ApplicationRefereeAllOf) GetFriendIntegrationIdOk() (string, bool)`

GetFriendIntegrationIdOk returns a tuple with the FriendIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFriendIntegrationId

`func (o *ApplicationRefereeAllOf) HasFriendIntegrationId() bool`

HasFriendIntegrationId returns a boolean if a field has been set.

### SetFriendIntegrationId

`func (o *ApplicationRefereeAllOf) SetFriendIntegrationId(v string)`

SetFriendIntegrationId gets a reference to the given string and assigns it to the FriendIntegrationId field.

### GetCode

`func (o *ApplicationRefereeAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplicationRefereeAllOf) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *ApplicationRefereeAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *ApplicationRefereeAllOf) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetCreated

`func (o *ApplicationRefereeAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationRefereeAllOf) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationRefereeAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationRefereeAllOf) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


