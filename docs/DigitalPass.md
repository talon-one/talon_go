# DigitalPass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassId** | Pointer to **string** | The ID of the generated digital pass. | 
**PassTemplateId** | Pointer to **string** | The ID of the digital pass template used to generate the pass. | 
**Status** | Pointer to **string** | The status of the digital pass. | 
**PassUrl** | Pointer to **string** | The URL you can use to let the customer add the digital pass to their wallet. | 

## Methods

### NewDigitalPass

`func NewDigitalPass(passId string, passTemplateId string, status string, passUrl string, ) *DigitalPass`

NewDigitalPass instantiates a new DigitalPass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalPassWithDefaults

`func NewDigitalPassWithDefaults() *DigitalPass`

NewDigitalPassWithDefaults instantiates a new DigitalPass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassId

`func (o *DigitalPass) GetPassId() string`

GetPassId returns the PassId field if non-nil, zero value otherwise.

### GetPassIdOk

`func (o *DigitalPass) GetPassIdOk() (*string, bool)`

GetPassIdOk returns a tuple with the PassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassId

`func (o *DigitalPass) SetPassId(v string)`

SetPassId sets PassId field to given value.


### GetPassTemplateId

`func (o *DigitalPass) GetPassTemplateId() string`

GetPassTemplateId returns the PassTemplateId field if non-nil, zero value otherwise.

### GetPassTemplateIdOk

`func (o *DigitalPass) GetPassTemplateIdOk() (*string, bool)`

GetPassTemplateIdOk returns a tuple with the PassTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassTemplateId

`func (o *DigitalPass) SetPassTemplateId(v string)`

SetPassTemplateId sets PassTemplateId field to given value.


### GetStatus

`func (o *DigitalPass) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DigitalPass) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DigitalPass) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPassUrl

`func (o *DigitalPass) GetPassUrl() string`

GetPassUrl returns the PassUrl field if non-nil, zero value otherwise.

### GetPassUrlOk

`func (o *DigitalPass) GetPassUrlOk() (*string, bool)`

GetPassUrlOk returns a tuple with the PassUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassUrl

`func (o *DigitalPass) SetPassUrl(v string)`

SetPassUrl sets PassUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


