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

### GetName

`func (o *SlotDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlotDef) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SlotDef) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SlotDef) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetType

`func (o *SlotDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlotDef) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *SlotDef) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *SlotDef) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetTitle

`func (o *SlotDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SlotDef) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *SlotDef) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *SlotDef) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *SlotDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlotDef) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *SlotDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *SlotDef) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetHelp

`func (o *SlotDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *SlotDef) GetHelpOk() (string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHelp

`func (o *SlotDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### SetHelp

`func (o *SlotDef) SetHelp(v string)`

SetHelp gets a reference to the given string and assigns it to the Help field.

### GetWritable

`func (o *SlotDef) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *SlotDef) GetWritableOk() (bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWritable

`func (o *SlotDef) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### SetWritable

`func (o *SlotDef) SetWritable(v bool)`

SetWritable gets a reference to the given bool and assigns it to the Writable field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


