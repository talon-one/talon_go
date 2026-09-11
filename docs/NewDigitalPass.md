# NewDigitalPass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoyaltyProgramId** | Pointer to **int64** | The ID of the associated loyalty program. | 
**PassTemplateId** | Pointer to **string** | The ID of the digital pass template used to generate the pass.  | 
**ProfileId** | Pointer to **string** | The integration ID of the customer profile the pass is issued for. | 
**LoyaltyCardId** | Pointer to **string** | The identifier of the loyalty card the pass is issued for.  **Note**: Only applicable for card-based loyalty programs.  | [optional] 
**Platform** | Pointer to **string** | The wallet platform the pass is generated for. | 
**Attributes** | Pointer to **map[string]string** | A map of placeholder values that you provide to fill in the pass template. These values are not validated against the template.  | [optional] 

## Methods

### NewNewDigitalPass

`func NewNewDigitalPass(loyaltyProgramId int64, passTemplateId string, profileId string, platform string, ) *NewDigitalPass`

NewNewDigitalPass instantiates a new NewDigitalPass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDigitalPassWithDefaults

`func NewNewDigitalPassWithDefaults() *NewDigitalPass`

NewNewDigitalPassWithDefaults instantiates a new NewDigitalPass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoyaltyProgramId

`func (o *NewDigitalPass) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *NewDigitalPass) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *NewDigitalPass) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.


### GetPassTemplateId

`func (o *NewDigitalPass) GetPassTemplateId() string`

GetPassTemplateId returns the PassTemplateId field if non-nil, zero value otherwise.

### GetPassTemplateIdOk

`func (o *NewDigitalPass) GetPassTemplateIdOk() (*string, bool)`

GetPassTemplateIdOk returns a tuple with the PassTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassTemplateId

`func (o *NewDigitalPass) SetPassTemplateId(v string)`

SetPassTemplateId sets PassTemplateId field to given value.


### GetProfileId

`func (o *NewDigitalPass) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *NewDigitalPass) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *NewDigitalPass) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetLoyaltyCardId

`func (o *NewDigitalPass) GetLoyaltyCardId() string`

GetLoyaltyCardId returns the LoyaltyCardId field if non-nil, zero value otherwise.

### GetLoyaltyCardIdOk

`func (o *NewDigitalPass) GetLoyaltyCardIdOk() (*string, bool)`

GetLoyaltyCardIdOk returns a tuple with the LoyaltyCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCardId

`func (o *NewDigitalPass) SetLoyaltyCardId(v string)`

SetLoyaltyCardId sets LoyaltyCardId field to given value.

### HasLoyaltyCardId

`func (o *NewDigitalPass) HasLoyaltyCardId() bool`

HasLoyaltyCardId returns a boolean if a field has been set.

### GetPlatform

`func (o *NewDigitalPass) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NewDigitalPass) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NewDigitalPass) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAttributes

`func (o *NewDigitalPass) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewDigitalPass) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewDigitalPass) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewDigitalPass) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


