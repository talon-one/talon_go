# ShowNotificationEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of notification that should be shown (e.g. error/warning/info). | 
**Title** | Pointer to **string** | Title of the notification. | 
**Body** | Pointer to **string** | Body of the notification. | 

## Methods

### NewShowNotificationEffectProps

`func NewShowNotificationEffectProps(notificationType string, title string, body string, ) *ShowNotificationEffectProps`

NewShowNotificationEffectProps instantiates a new ShowNotificationEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowNotificationEffectPropsWithDefaults

`func NewShowNotificationEffectPropsWithDefaults() *ShowNotificationEffectProps`

NewShowNotificationEffectPropsWithDefaults instantiates a new ShowNotificationEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *ShowNotificationEffectProps) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ShowNotificationEffectProps) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ShowNotificationEffectProps) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTitle

`func (o *ShowNotificationEffectProps) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShowNotificationEffectProps) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShowNotificationEffectProps) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ShowNotificationEffectProps) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ShowNotificationEffectProps) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ShowNotificationEffectProps) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


