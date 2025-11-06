# GenerateCouponFailureSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **int64** | The ID of the event. | 
**Language** | Pointer to **string** | The language the summary will be generated in. | [optional] 

## Methods

### NewGenerateCouponFailureSummary

`func NewGenerateCouponFailureSummary(eventID int64, ) *GenerateCouponFailureSummary`

NewGenerateCouponFailureSummary instantiates a new GenerateCouponFailureSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCouponFailureSummaryWithDefaults

`func NewGenerateCouponFailureSummaryWithDefaults() *GenerateCouponFailureSummary`

NewGenerateCouponFailureSummaryWithDefaults instantiates a new GenerateCouponFailureSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *GenerateCouponFailureSummary) GetEventID() int64`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *GenerateCouponFailureSummary) GetEventIDOk() (*int64, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *GenerateCouponFailureSummary) SetEventID(v int64)`

SetEventID sets EventID field to given value.


### GetLanguage

`func (o *GenerateCouponFailureSummary) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GenerateCouponFailureSummary) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GenerateCouponFailureSummary) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GenerateCouponFailureSummary) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


