# ProjectedTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectedActivePoints** | Pointer to **float32** | The active points of the customer when their current tier expires. | 
**StayInTierPoints** | Pointer to **float32** | The number of points the customer needs to stay in the current tier.  **Note**: This is included in the response when the customer is projected to be downgraded.  | [optional] 
**ProjectedTierName** | Pointer to **string** | The name of the tier the user will enter once their current tier expires. | [optional] 

## Methods

### NewProjectedTier

`func NewProjectedTier(projectedActivePoints float32, ) *ProjectedTier`

NewProjectedTier instantiates a new ProjectedTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectedTierWithDefaults

`func NewProjectedTierWithDefaults() *ProjectedTier`

NewProjectedTierWithDefaults instantiates a new ProjectedTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectedActivePoints

`func (o *ProjectedTier) GetProjectedActivePoints() float32`

GetProjectedActivePoints returns the ProjectedActivePoints field if non-nil, zero value otherwise.

### GetProjectedActivePointsOk

`func (o *ProjectedTier) GetProjectedActivePointsOk() (*float32, bool)`

GetProjectedActivePointsOk returns a tuple with the ProjectedActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedActivePoints

`func (o *ProjectedTier) SetProjectedActivePoints(v float32)`

SetProjectedActivePoints sets ProjectedActivePoints field to given value.


### GetStayInTierPoints

`func (o *ProjectedTier) GetStayInTierPoints() float32`

GetStayInTierPoints returns the StayInTierPoints field if non-nil, zero value otherwise.

### GetStayInTierPointsOk

`func (o *ProjectedTier) GetStayInTierPointsOk() (*float32, bool)`

GetStayInTierPointsOk returns a tuple with the StayInTierPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStayInTierPoints

`func (o *ProjectedTier) SetStayInTierPoints(v float32)`

SetStayInTierPoints sets StayInTierPoints field to given value.

### HasStayInTierPoints

`func (o *ProjectedTier) HasStayInTierPoints() bool`

HasStayInTierPoints returns a boolean if a field has been set.

### GetProjectedTierName

`func (o *ProjectedTier) GetProjectedTierName() string`

GetProjectedTierName returns the ProjectedTierName field if non-nil, zero value otherwise.

### GetProjectedTierNameOk

`func (o *ProjectedTier) GetProjectedTierNameOk() (*string, bool)`

GetProjectedTierNameOk returns a tuple with the ProjectedTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedTierName

`func (o *ProjectedTier) SetProjectedTierName(v string)`

SetProjectedTierName sets ProjectedTierName field to given value.

### HasProjectedTierName

`func (o *ProjectedTier) HasProjectedTierName() bool`

HasProjectedTierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


