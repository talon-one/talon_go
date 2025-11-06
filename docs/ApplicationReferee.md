# ApplicationReferee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**SessionId** | Pointer to **string** | Integration ID of the session in which the customer redeemed the referral. | 
**AdvocateIntegrationId** | Pointer to **string** | Integration ID of the Advocate&#39;s Profile. | 
**FriendIntegrationId** | Pointer to **string** | Integration ID of the Friend&#39;s Profile. | 
**Code** | Pointer to **string** | Advocate&#39;s referral code. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment the customer redeemed the referral. | 

## Methods

### NewApplicationReferee

`func NewApplicationReferee(applicationId int64, sessionId string, advocateIntegrationId string, friendIntegrationId string, code string, created time.Time, ) *ApplicationReferee`

NewApplicationReferee instantiates a new ApplicationReferee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRefereeWithDefaults

`func NewApplicationRefereeWithDefaults() *ApplicationReferee`

NewApplicationRefereeWithDefaults instantiates a new ApplicationReferee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationReferee) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationReferee) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationReferee) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetSessionId

`func (o *ApplicationReferee) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationReferee) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApplicationReferee) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetAdvocateIntegrationId

`func (o *ApplicationReferee) GetAdvocateIntegrationId() string`

GetAdvocateIntegrationId returns the AdvocateIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateIntegrationIdOk

`func (o *ApplicationReferee) GetAdvocateIntegrationIdOk() (*string, bool)`

GetAdvocateIntegrationIdOk returns a tuple with the AdvocateIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocateIntegrationId

`func (o *ApplicationReferee) SetAdvocateIntegrationId(v string)`

SetAdvocateIntegrationId sets AdvocateIntegrationId field to given value.


### GetFriendIntegrationId

`func (o *ApplicationReferee) GetFriendIntegrationId() string`

GetFriendIntegrationId returns the FriendIntegrationId field if non-nil, zero value otherwise.

### GetFriendIntegrationIdOk

`func (o *ApplicationReferee) GetFriendIntegrationIdOk() (*string, bool)`

GetFriendIntegrationIdOk returns a tuple with the FriendIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendIntegrationId

`func (o *ApplicationReferee) SetFriendIntegrationId(v string)`

SetFriendIntegrationId sets FriendIntegrationId field to given value.


### GetCode

`func (o *ApplicationReferee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplicationReferee) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApplicationReferee) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreated

`func (o *ApplicationReferee) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationReferee) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationReferee) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


