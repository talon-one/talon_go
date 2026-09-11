# GiveawayPoolReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the giveaway pool. | 
**Name** | Pointer to **string** | The display name of the giveaway pool. | [readonly] 

## Methods

### NewGiveawayPoolReference

`func NewGiveawayPoolReference(id int64, name string, ) *GiveawayPoolReference`

NewGiveawayPoolReference instantiates a new GiveawayPoolReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayPoolReferenceWithDefaults

`func NewGiveawayPoolReferenceWithDefaults() *GiveawayPoolReference`

NewGiveawayPoolReferenceWithDefaults instantiates a new GiveawayPoolReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GiveawayPoolReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiveawayPoolReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiveawayPoolReference) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *GiveawayPoolReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiveawayPoolReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiveawayPoolReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


