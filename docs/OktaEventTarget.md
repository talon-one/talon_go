# OktaEventTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the event target. | 
**AlternateId** | Pointer to **string** | Identifier of the event target, depending on its type. | 
**DisplayName** | Pointer to **string** | Display name of the event target. | 

## Methods

### NewOktaEventTarget

`func NewOktaEventTarget(type_ string, alternateId string, displayName string, ) *OktaEventTarget`

NewOktaEventTarget instantiates a new OktaEventTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaEventTargetWithDefaults

`func NewOktaEventTargetWithDefaults() *OktaEventTarget`

NewOktaEventTargetWithDefaults instantiates a new OktaEventTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OktaEventTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OktaEventTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OktaEventTarget) SetType(v string)`

SetType sets Type field to given value.


### GetAlternateId

`func (o *OktaEventTarget) GetAlternateId() string`

GetAlternateId returns the AlternateId field if non-nil, zero value otherwise.

### GetAlternateIdOk

`func (o *OktaEventTarget) GetAlternateIdOk() (*string, bool)`

GetAlternateIdOk returns a tuple with the AlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateId

`func (o *OktaEventTarget) SetAlternateId(v string)`

SetAlternateId sets AlternateId field to given value.


### GetDisplayName

`func (o *OktaEventTarget) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OktaEventTarget) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OktaEventTarget) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


