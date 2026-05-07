# CouponFailureSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the evaluation record. | 
**EventID** | Pointer to **int64** | ID of the event. | 
**SessionID** | Pointer to **string** | ID of the customer session set by your integration layer. | [optional] 
**ProfileID** | Pointer to **string** | ID of the customer profile set by your integration layer. | [optional] 
**Status** | Pointer to **string** | Status defines if the coupon code was applied or rejected. | 
**CouponCode** | Pointer to **string** | Coupon code passed for evaluation. | 
**Language** | Pointer to **string** | Language of the summary. | 
**Summary** | Pointer to **string** | A summary of the reasons for coupon redemption failure. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the request was made. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the request was last updated. | 

## Methods

### NewCouponFailureSummary

`func NewCouponFailureSummary(id int64, eventID int64, status string, couponCode string, language string, summary string, createdAt time.Time, updatedAt time.Time, ) *CouponFailureSummary`

NewCouponFailureSummary instantiates a new CouponFailureSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponFailureSummaryWithDefaults

`func NewCouponFailureSummaryWithDefaults() *CouponFailureSummary`

NewCouponFailureSummaryWithDefaults instantiates a new CouponFailureSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponFailureSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponFailureSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponFailureSummary) SetId(v int64)`

SetId sets Id field to given value.


### GetEventID

`func (o *CouponFailureSummary) GetEventID() int64`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *CouponFailureSummary) GetEventIDOk() (*int64, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *CouponFailureSummary) SetEventID(v int64)`

SetEventID sets EventID field to given value.


### GetSessionID

`func (o *CouponFailureSummary) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *CouponFailureSummary) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *CouponFailureSummary) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.

### HasSessionID

`func (o *CouponFailureSummary) HasSessionID() bool`

HasSessionID returns a boolean if a field has been set.

### GetProfileID

`func (o *CouponFailureSummary) GetProfileID() string`

GetProfileID returns the ProfileID field if non-nil, zero value otherwise.

### GetProfileIDOk

`func (o *CouponFailureSummary) GetProfileIDOk() (*string, bool)`

GetProfileIDOk returns a tuple with the ProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileID

`func (o *CouponFailureSummary) SetProfileID(v string)`

SetProfileID sets ProfileID field to given value.

### HasProfileID

`func (o *CouponFailureSummary) HasProfileID() bool`

HasProfileID returns a boolean if a field has been set.

### GetStatus

`func (o *CouponFailureSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponFailureSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponFailureSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCouponCode

`func (o *CouponFailureSummary) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CouponFailureSummary) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CouponFailureSummary) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetLanguage

`func (o *CouponFailureSummary) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CouponFailureSummary) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CouponFailureSummary) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetSummary

`func (o *CouponFailureSummary) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CouponFailureSummary) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CouponFailureSummary) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetCreatedAt

`func (o *CouponFailureSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponFailureSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponFailureSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CouponFailureSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponFailureSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponFailureSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


