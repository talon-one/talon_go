# GeoJSONGeometryCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The geometry type discriminator. | 
**Geometries** | Pointer to **[]map[string]interface{}** | The shapes contained in this group. | 

## Methods

### NewGeoJSONGeometryCollection

`func NewGeoJSONGeometryCollection(type_ string, geometries []map[string]interface{}, ) *GeoJSONGeometryCollection`

NewGeoJSONGeometryCollection instantiates a new GeoJSONGeometryCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONGeometryCollectionWithDefaults

`func NewGeoJSONGeometryCollectionWithDefaults() *GeoJSONGeometryCollection`

NewGeoJSONGeometryCollectionWithDefaults instantiates a new GeoJSONGeometryCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONGeometryCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONGeometryCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONGeometryCollection) SetType(v string)`

SetType sets Type field to given value.


### GetGeometries

`func (o *GeoJSONGeometryCollection) GetGeometries() []map[string]interface{}`

GetGeometries returns the Geometries field if non-nil, zero value otherwise.

### GetGeometriesOk

`func (o *GeoJSONGeometryCollection) GetGeometriesOk() (*[]map[string]interface{}, bool)`

GetGeometriesOk returns a tuple with the Geometries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometries

`func (o *GeoJSONGeometryCollection) SetGeometries(v []map[string]interface{})`

SetGeometries sets Geometries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


