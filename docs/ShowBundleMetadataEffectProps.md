# ShowBundleMetadataEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the product bundle. | 
**BundleAttributes** | Pointer to **[]string** | The cart item attributes that determined which items are being bundled together. | 
**ItemsIndices** | Pointer to **[]float32** | The indices in the cart items array of the bundled items. | 

## Methods

### NewShowBundleMetadataEffectProps

`func NewShowBundleMetadataEffectProps(description string, bundleAttributes []string, itemsIndices []float32, ) *ShowBundleMetadataEffectProps`

NewShowBundleMetadataEffectProps instantiates a new ShowBundleMetadataEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowBundleMetadataEffectPropsWithDefaults

`func NewShowBundleMetadataEffectPropsWithDefaults() *ShowBundleMetadataEffectProps`

NewShowBundleMetadataEffectPropsWithDefaults instantiates a new ShowBundleMetadataEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ShowBundleMetadataEffectProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShowBundleMetadataEffectProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShowBundleMetadataEffectProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBundleAttributes

`func (o *ShowBundleMetadataEffectProps) GetBundleAttributes() []string`

GetBundleAttributes returns the BundleAttributes field if non-nil, zero value otherwise.

### GetBundleAttributesOk

`func (o *ShowBundleMetadataEffectProps) GetBundleAttributesOk() (*[]string, bool)`

GetBundleAttributesOk returns a tuple with the BundleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleAttributes

`func (o *ShowBundleMetadataEffectProps) SetBundleAttributes(v []string)`

SetBundleAttributes sets BundleAttributes field to given value.


### GetItemsIndices

`func (o *ShowBundleMetadataEffectProps) GetItemsIndices() []float32`

GetItemsIndices returns the ItemsIndices field if non-nil, zero value otherwise.

### GetItemsIndicesOk

`func (o *ShowBundleMetadataEffectProps) GetItemsIndicesOk() (*[]float32, bool)`

GetItemsIndicesOk returns a tuple with the ItemsIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsIndices

`func (o *ShowBundleMetadataEffectProps) SetItemsIndices(v []float32)`

SetItemsIndices sets ItemsIndices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


