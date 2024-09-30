# ScimServiceProviderConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bulk** | Pointer to [**ScimServiceProviderConfigResponseBulk**](ScimServiceProviderConfigResponse_bulk.md) |  | [optional] 
**ChangePassword** | Pointer to [**ScimServiceProviderConfigResponseChangePassword**](ScimServiceProviderConfigResponse_changePassword.md) |  | [optional] 
**DocumentationUri** | Pointer to **string** | The URI that points to the SCIM service provider&#39;s documentation, providing further details about the service&#39;s capabilities and usage. | [optional] 
**Filter** | Pointer to [**ScimServiceProviderConfigResponseFilter**](ScimServiceProviderConfigResponse_filter.md) |  | [optional] 
**Patch** | Pointer to [**ScimServiceProviderConfigResponsePatch**](ScimServiceProviderConfigResponse_patch.md) |  | [optional] 
**Schemas** | Pointer to **[]string** | A list of SCIM schemas that define the structure and data types supported by the service provider. | [optional] 

## Methods

### GetBulk

`func (o *ScimServiceProviderConfigResponse) GetBulk() ScimServiceProviderConfigResponseBulk`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *ScimServiceProviderConfigResponse) GetBulkOk() (ScimServiceProviderConfigResponseBulk, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBulk

`func (o *ScimServiceProviderConfigResponse) HasBulk() bool`

HasBulk returns a boolean if a field has been set.

### SetBulk

`func (o *ScimServiceProviderConfigResponse) SetBulk(v ScimServiceProviderConfigResponseBulk)`

SetBulk gets a reference to the given ScimServiceProviderConfigResponseBulk and assigns it to the Bulk field.

### GetChangePassword

`func (o *ScimServiceProviderConfigResponse) GetChangePassword() ScimServiceProviderConfigResponseChangePassword`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *ScimServiceProviderConfigResponse) GetChangePasswordOk() (ScimServiceProviderConfigResponseChangePassword, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangePassword

`func (o *ScimServiceProviderConfigResponse) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### SetChangePassword

`func (o *ScimServiceProviderConfigResponse) SetChangePassword(v ScimServiceProviderConfigResponseChangePassword)`

SetChangePassword gets a reference to the given ScimServiceProviderConfigResponseChangePassword and assigns it to the ChangePassword field.

### GetDocumentationUri

`func (o *ScimServiceProviderConfigResponse) GetDocumentationUri() string`

GetDocumentationUri returns the DocumentationUri field if non-nil, zero value otherwise.

### GetDocumentationUriOk

`func (o *ScimServiceProviderConfigResponse) GetDocumentationUriOk() (string, bool)`

GetDocumentationUriOk returns a tuple with the DocumentationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDocumentationUri

`func (o *ScimServiceProviderConfigResponse) HasDocumentationUri() bool`

HasDocumentationUri returns a boolean if a field has been set.

### SetDocumentationUri

`func (o *ScimServiceProviderConfigResponse) SetDocumentationUri(v string)`

SetDocumentationUri gets a reference to the given string and assigns it to the DocumentationUri field.

### GetFilter

`func (o *ScimServiceProviderConfigResponse) GetFilter() ScimServiceProviderConfigResponseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScimServiceProviderConfigResponse) GetFilterOk() (ScimServiceProviderConfigResponseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilter

`func (o *ScimServiceProviderConfigResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilter

`func (o *ScimServiceProviderConfigResponse) SetFilter(v ScimServiceProviderConfigResponseFilter)`

SetFilter gets a reference to the given ScimServiceProviderConfigResponseFilter and assigns it to the Filter field.

### GetPatch

`func (o *ScimServiceProviderConfigResponse) GetPatch() ScimServiceProviderConfigResponsePatch`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ScimServiceProviderConfigResponse) GetPatchOk() (ScimServiceProviderConfigResponsePatch, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatch

`func (o *ScimServiceProviderConfigResponse) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### SetPatch

`func (o *ScimServiceProviderConfigResponse) SetPatch(v ScimServiceProviderConfigResponsePatch)`

SetPatch gets a reference to the given ScimServiceProviderConfigResponsePatch and assigns it to the Patch field.

### GetSchemas

`func (o *ScimServiceProviderConfigResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimServiceProviderConfigResponse) GetSchemasOk() ([]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchemas

`func (o *ScimServiceProviderConfigResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemas

`func (o *ScimServiceProviderConfigResponse) SetSchemas(v []string)`

SetSchemas gets a reference to the given []string and assigns it to the Schemas field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


