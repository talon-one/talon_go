# GenerateCouponFailureDetailedSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationID** | Pointer to **float32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**SessionID** | Pointer to **string** | ID of the customer session where the coupon redemption failed. | 
**EventID** | Pointer to **int64** | The ID of the event for which the coupon redemption failed. | 
**Coupon** | Pointer to **string** | The coupon code that could not be redeemed. | 
**Language** | Pointer to **string** | The language of the summary. | [optional] 

## Methods

### NewGenerateCouponFailureDetailedSummary

`func NewGenerateCouponFailureDetailedSummary(applicationID float32, sessionID string, eventID int64, coupon string, ) *GenerateCouponFailureDetailedSummary`

NewGenerateCouponFailureDetailedSummary instantiates a new GenerateCouponFailureDetailedSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCouponFailureDetailedSummaryWithDefaults

`func NewGenerateCouponFailureDetailedSummaryWithDefaults() *GenerateCouponFailureDetailedSummary`

NewGenerateCouponFailureDetailedSummaryWithDefaults instantiates a new GenerateCouponFailureDetailedSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationID

`func (o *GenerateCouponFailureDetailedSummary) GetApplicationID() float32`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *GenerateCouponFailureDetailedSummary) GetApplicationIDOk() (*float32, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *GenerateCouponFailureDetailedSummary) SetApplicationID(v float32)`

SetApplicationID sets ApplicationID field to given value.


### GetSessionID

`func (o *GenerateCouponFailureDetailedSummary) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *GenerateCouponFailureDetailedSummary) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *GenerateCouponFailureDetailedSummary) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetEventID

`func (o *GenerateCouponFailureDetailedSummary) GetEventID() int64`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *GenerateCouponFailureDetailedSummary) GetEventIDOk() (*int64, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *GenerateCouponFailureDetailedSummary) SetEventID(v int64)`

SetEventID sets EventID field to given value.


### GetCoupon

`func (o *GenerateCouponFailureDetailedSummary) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *GenerateCouponFailureDetailedSummary) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *GenerateCouponFailureDetailedSummary) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.


### GetLanguage

`func (o *GenerateCouponFailureDetailedSummary) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GenerateCouponFailureDetailedSummary) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GenerateCouponFailureDetailedSummary) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GenerateCouponFailureDetailedSummary) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


