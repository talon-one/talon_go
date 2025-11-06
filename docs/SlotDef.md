# SlotDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The dot-separated path to this slot for use in Talang. | 
**Type** | Pointer to **string** | The type of this slot, one of string, number, boolean, or list[type]. | 
**Title** | Pointer to **string** | Campaigner-friendly name for the slot. | 
**Description** | Pointer to **string** | A short description of the slot. | [optional] 
**Help** | Pointer to **string** | Extended help text for the slot. | [optional] 
**Writable** | Pointer to **bool** | Whether or not this slot can be updated by rule effects. | 

## Methods

### NewSlotDef

`func NewSlotDef(name string, type_ string, title string, writable bool, ) *SlotDef`

NewSlotDef instantiates a new SlotDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlotDefWithDefaults

`func NewSlotDefWithDefaults() *SlotDef`

NewSlotDefWithDefaults instantiates a new SlotDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlotDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlotDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlotDef) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SlotDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlotDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlotDef) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *SlotDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SlotDef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SlotDef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SlotDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlotDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlotDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlotDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHelp

`func (o *SlotDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *SlotDef) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *SlotDef) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *SlotDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetWritable

`func (o *SlotDef) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *SlotDef) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *SlotDef) SetWritable(v bool)`

SetWritable sets Writable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


