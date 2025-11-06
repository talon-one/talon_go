# CampaignTemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the campaign template parameter. | 
**Type** | Pointer to **string** | Defines the type of parameter value. | 
**Description** | Pointer to **string** | Explains the meaning of this template parameter and the placeholder value that will define it. It is used on campaign creation from this template. | 
**AttributeId** | Pointer to **int64** | ID of the corresponding attribute. | [optional] 

## Methods

### NewCampaignTemplateParams

`func NewCampaignTemplateParams(name string, type_ string, description string, ) *CampaignTemplateParams`

NewCampaignTemplateParams instantiates a new CampaignTemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignTemplateParamsWithDefaults

`func NewCampaignTemplateParamsWithDefaults() *CampaignTemplateParams`

NewCampaignTemplateParamsWithDefaults instantiates a new CampaignTemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignTemplateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignTemplateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignTemplateParams) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CampaignTemplateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignTemplateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignTemplateParams) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CampaignTemplateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignTemplateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignTemplateParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAttributeId

`func (o *CampaignTemplateParams) GetAttributeId() int64`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *CampaignTemplateParams) GetAttributeIdOk() (*int64, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *CampaignTemplateParams) SetAttributeId(v int64)`

SetAttributeId sets AttributeId field to given value.

### HasAttributeId

`func (o *CampaignTemplateParams) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


