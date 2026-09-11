# CampaignEligibilityExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the experiment. | 
**VariantId** | Pointer to **int64** | The ID of the variant assigned to the customer profile. | 

## Methods

### NewCampaignEligibilityExperiment

`func NewCampaignEligibilityExperiment(id int64, variantId int64, ) *CampaignEligibilityExperiment`

NewCampaignEligibilityExperiment instantiates a new CampaignEligibilityExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEligibilityExperimentWithDefaults

`func NewCampaignEligibilityExperimentWithDefaults() *CampaignEligibilityExperiment`

NewCampaignEligibilityExperimentWithDefaults instantiates a new CampaignEligibilityExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignEligibilityExperiment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEligibilityExperiment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignEligibilityExperiment) SetId(v int64)`

SetId sets Id field to given value.


### GetVariantId

`func (o *CampaignEligibilityExperiment) GetVariantId() int64`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *CampaignEligibilityExperiment) GetVariantIdOk() (*int64, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *CampaignEligibilityExperiment) SetVariantId(v int64)`

SetVariantId sets VariantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


