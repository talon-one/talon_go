# BestPriorTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | Pointer to **string** | The type of price target. | 
**AudienceID** | Pointer to **int64** | The AudienceID of an audience. Must be used with \&quot;AUDIENCE\&quot; target type. | [optional] 

## Methods

### NewBestPriorTarget

`func NewBestPriorTarget(targetType string, ) *BestPriorTarget`

NewBestPriorTarget instantiates a new BestPriorTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorTargetWithDefaults

`func NewBestPriorTargetWithDefaults() *BestPriorTarget`

NewBestPriorTargetWithDefaults instantiates a new BestPriorTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *BestPriorTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *BestPriorTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *BestPriorTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetAudienceID

`func (o *BestPriorTarget) GetAudienceID() int64`

GetAudienceID returns the AudienceID field if non-nil, zero value otherwise.

### GetAudienceIDOk

`func (o *BestPriorTarget) GetAudienceIDOk() (*int64, bool)`

GetAudienceIDOk returns a tuple with the AudienceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceID

`func (o *BestPriorTarget) SetAudienceID(v int64)`

SetAudienceID sets AudienceID field to given value.

### HasAudienceID

`func (o *BestPriorTarget) HasAudienceID() bool`

HasAudienceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


