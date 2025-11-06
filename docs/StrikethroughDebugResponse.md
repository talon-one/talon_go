# StrikethroughDebugResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignsIDs** | Pointer to **[]int64** | The campaign IDs that got fetched for the evaluation process. | [optional] 
**Effects** | Pointer to [**[]StrikethroughEffect**](StrikethroughEffect.md) | The strikethrough effects that are returned from the evaluation process. | [optional] 

## Methods

### NewStrikethroughDebugResponse

`func NewStrikethroughDebugResponse() *StrikethroughDebugResponse`

NewStrikethroughDebugResponse instantiates a new StrikethroughDebugResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughDebugResponseWithDefaults

`func NewStrikethroughDebugResponseWithDefaults() *StrikethroughDebugResponse`

NewStrikethroughDebugResponseWithDefaults instantiates a new StrikethroughDebugResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignsIDs

`func (o *StrikethroughDebugResponse) GetCampaignsIDs() []int64`

GetCampaignsIDs returns the CampaignsIDs field if non-nil, zero value otherwise.

### GetCampaignsIDsOk

`func (o *StrikethroughDebugResponse) GetCampaignsIDsOk() (*[]int64, bool)`

GetCampaignsIDsOk returns a tuple with the CampaignsIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignsIDs

`func (o *StrikethroughDebugResponse) SetCampaignsIDs(v []int64)`

SetCampaignsIDs sets CampaignsIDs field to given value.

### HasCampaignsIDs

`func (o *StrikethroughDebugResponse) HasCampaignsIDs() bool`

HasCampaignsIDs returns a boolean if a field has been set.

### GetEffects

`func (o *StrikethroughDebugResponse) GetEffects() []StrikethroughEffect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *StrikethroughDebugResponse) GetEffectsOk() (*[]StrikethroughEffect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *StrikethroughDebugResponse) SetEffects(v []StrikethroughEffect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *StrikethroughDebugResponse) HasEffects() bool`

HasEffects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


