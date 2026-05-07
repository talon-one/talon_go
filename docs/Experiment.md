# Experiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally.  | [optional] 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | [optional] 
**Activated** | Pointer to [**time.Time**](time.Time.md) | The date and time the experiment was activated.  | [optional] 
**State** | Pointer to **string** | A disabled experiment is not evaluated for rules or coupons.  | [default to "disabled"]
**Variants** | Pointer to [**[]ExperimentVariant**](ExperimentVariant.md) |  | [optional] 
**Deletedat** | Pointer to [**time.Time**](time.Time.md) | The date and time the experiment was deleted.  | [optional] 

## Methods

### NewExperiment

`func NewExperiment(id int64, created time.Time, applicationId int64, state string, ) *Experiment`

NewExperiment instantiates a new Experiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentWithDefaults

`func NewExperimentWithDefaults() *Experiment`

NewExperimentWithDefaults instantiates a new Experiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Experiment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Experiment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Experiment) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Experiment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Experiment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Experiment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Experiment) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Experiment) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Experiment) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetIsVariantAssignmentExternal

`func (o *Experiment) GetIsVariantAssignmentExternal() bool`

GetIsVariantAssignmentExternal returns the IsVariantAssignmentExternal field if non-nil, zero value otherwise.

### GetIsVariantAssignmentExternalOk

`func (o *Experiment) GetIsVariantAssignmentExternalOk() (*bool, bool)`

GetIsVariantAssignmentExternalOk returns a tuple with the IsVariantAssignmentExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariantAssignmentExternal

`func (o *Experiment) SetIsVariantAssignmentExternal(v bool)`

SetIsVariantAssignmentExternal sets IsVariantAssignmentExternal field to given value.

### HasIsVariantAssignmentExternal

`func (o *Experiment) HasIsVariantAssignmentExternal() bool`

HasIsVariantAssignmentExternal returns a boolean if a field has been set.

### GetCampaign

`func (o *Experiment) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Experiment) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Experiment) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *Experiment) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetActivated

`func (o *Experiment) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Experiment) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Experiment) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *Experiment) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetState

`func (o *Experiment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Experiment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Experiment) SetState(v string)`

SetState sets State field to given value.


### GetVariants

`func (o *Experiment) GetVariants() []ExperimentVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Experiment) GetVariantsOk() (*[]ExperimentVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Experiment) SetVariants(v []ExperimentVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *Experiment) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetDeletedat

`func (o *Experiment) GetDeletedat() time.Time`

GetDeletedat returns the Deletedat field if non-nil, zero value otherwise.

### GetDeletedatOk

`func (o *Experiment) GetDeletedatOk() (*time.Time, bool)`

GetDeletedatOk returns a tuple with the Deletedat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedat

`func (o *Experiment) SetDeletedat(v time.Time)`

SetDeletedat sets Deletedat field to given value.

### HasDeletedat

`func (o *Experiment) HasDeletedat() bool`

HasDeletedat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


