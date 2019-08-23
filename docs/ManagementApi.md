# \ManagementApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLoyaltyPoints**](ManagementApi.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/add_points | Add points in a certain loyalty program for the specified customer
[**CopyCampaignToApplications**](ManagementApi.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into every specified application
[**CreateCampaign**](ManagementApi.md#CreateCampaign) | **Post** /v1/applications/{applicationId}/campaigns | Create a Campaign
[**CreateCoupons**](ManagementApi.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create Coupons
[**CreatePasswordRecoveryEmail**](ManagementApi.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateRuleset**](ManagementApi.md#CreateRuleset) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | Create a Ruleset
[**CreateSession**](ManagementApi.md#CreateSession) | **Post** /v1/sessions | Create a Session
[**DeleteCampaign**](ManagementApi.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete a Campaign
[**DeleteCoupon**](ManagementApi.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete one Coupon
[**DeleteCoupons**](ManagementApi.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete Coupons
[**DeleteReferral**](ManagementApi.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete one Referral
[**DeleteRuleset**](ManagementApi.md#DeleteRuleset) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Delete a Ruleset
[**GetAccessLogs**](ManagementApi.md#GetAccessLogs) | **Get** /v1/applications/{applicationId}/access_logs | Get access logs for application
[**GetAccessLogsWithoutTotalCount**](ManagementApi.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for application
[**GetAccount**](ManagementApi.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get Account Details
[**GetAccountAnalytics**](ManagementApi.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get Account Analytics
[**GetAccountLimits**](ManagementApi.md#GetAccountLimits) | **Get** /v1/accounts/{accountId}/limits | Get Account Limits
[**GetAllAccessLogs**](ManagementApi.md#GetAllAccessLogs) | **Get** /v1/access_logs | Get all access logs
[**GetAllRoles**](ManagementApi.md#GetAllRoles) | **Get** /v1/roles | Get all roles.
[**GetApplication**](ManagementApi.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementApi.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get report of health of application API
[**GetApplicationCustomer**](ManagementApi.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get Application Customer
[**GetApplicationCustomers**](ManagementApi.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List Application Customers
[**GetApplicationCustomersByAttributes**](ManagementApi.md#GetApplicationCustomersByAttributes) | **Post** /v1/application_customer_search | Get a list of the customer profiles that match the given attributes
[**GetApplicationEventTypes**](ManagementApi.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications Event Types
[**GetApplicationEvents**](ManagementApi.md#GetApplicationEvents) | **Get** /v1/applications/{applicationId}/events | List Applications Events
[**GetApplicationEventsWithoutTotalCount**](ManagementApi.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications Events
[**GetApplicationSession**](ManagementApi.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application Session
[**GetApplicationSessions**](ManagementApi.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application Sessions
[**GetApplications**](ManagementApi.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementApi.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get a custom attribute
[**GetCampaign**](ManagementApi.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get a Campaign
[**GetCampaignAnalytics**](ManagementApi.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementApi.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | Get a list of all campaigns that match the given attributes
[**GetCampaignSet**](ManagementApi.md#GetCampaignSet) | **Get** /v1/applications/{applicationId}/campaign_set | List CampaignSet
[**GetCampaigns**](ManagementApi.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List your Campaigns
[**GetChanges**](ManagementApi.md#GetChanges) | **Get** /v1/changes | Get audit log for an account
[**GetCoupons**](ManagementApi.md#GetCoupons) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | List Coupons
[**GetCouponsByAttributes**](ManagementApi.md#GetCouponsByAttributes) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search | Get a list of the coupons that match the given attributes
[**GetCouponsByAttributesApplicationWide**](ManagementApi.md#GetCouponsByAttributesApplicationWide) | **Post** /v1/applications/{applicationId}/coupons_search | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**GetCouponsWithoutTotalCount**](ManagementApi.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List Coupons
[**GetCustomerActivityReport**](ManagementApi.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get Activity Report for Single Customer
[**GetCustomerActivityReports**](ManagementApi.md#GetCustomerActivityReports) | **Get** /v1/applications/{applicationId}/customer_activity_reports | Get Activity Reports for Application Customers
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementApi.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementApi.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get Analytics Report for a Customer
[**GetCustomerProfile**](ManagementApi.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get Customer Profile
[**GetCustomerProfiles**](ManagementApi.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List Customer Profiles
[**GetCustomersByAttributes**](ManagementApi.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | Get a list of the customer profiles that match the given attributes
[**GetEventTypes**](ManagementApi.md#GetEventTypes) | **Get** /v1/event_types | List Event Types
[**GetExports**](ManagementApi.md#GetExports) | **Get** /v1/exports | Get Exports
[**GetImports**](ManagementApi.md#GetImports) | **Get** /v1/imports | Get Imports
[**GetLoyaltyPoints**](ManagementApi.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{programID}/profile/{integrationID} | get the Loyalty Ledger for this integrationID
[**GetLoyaltyProgram**](ManagementApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{programID} | Get a loyalty program
[**GetLoyaltyPrograms**](ManagementApi.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List all loyalty Programs
[**GetReferrals**](ManagementApi.md#GetReferrals) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals | List Referrals
[**GetReferralsWithoutTotalCount**](ManagementApi.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List Referrals
[**GetRole**](ManagementApi.md#GetRole) | **Get** /v1/roles/{roleId} | Get information for the specified role.
[**GetRuleset**](ManagementApi.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get a Ruleset
[**GetRulesets**](ManagementApi.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List Campaign Rulesets
[**GetUser**](ManagementApi.md#GetUser) | **Get** /v1/users/{userId} | Get a single User
[**GetUsers**](ManagementApi.md#GetUsers) | **Get** /v1/users | List Users in your account
[**GetWebhook**](ManagementApi.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get Webhook
[**GetWebhookActivationLogs**](ManagementApi.md#GetWebhookActivationLogs) | **Get** /v1/webhook_activation_logs | List Webhook activation Log Entries
[**GetWebhookLogs**](ManagementApi.md#GetWebhookLogs) | **Get** /v1/webhook_logs | List Webhook Log Entries
[**GetWebhooks**](ManagementApi.md#GetWebhooks) | **Get** /v1/webhooks | List Webhooks
[**RefreshAnalytics**](ManagementApi.md#RefreshAnalytics) | **Post** /v1/refresh_analytics | Trigger refresh on stale analytics.
[**RemoveLoyaltyPoints**](ManagementApi.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/deduct_points | Deduct points in a certain loyalty program for the specified customer
[**ResetPassword**](ManagementApi.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**SearchCouponsAdvanced**](ManagementApi.md#SearchCouponsAdvanced) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced | Get a list of the coupons that match the given attributes
[**SearchCouponsAdvancedApplicationWide**](ManagementApi.md#SearchCouponsAdvancedApplicationWide) | **Post** /v1/applications/{applicationId}/coupons_search_advanced | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | Get a list of the coupons that match the given attributes
[**SetAccountLimits**](ManagementApi.md#SetAccountLimits) | **Put** /v1/accounts/{accountId}/limits | Set account limits
[**UpdateCampaign**](ManagementApi.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update a Campaign
[**UpdateCampaignSet**](ManagementApi.md#UpdateCampaignSet) | **Put** /v1/applications/{applicationId}/campaign_set | Update a Campaign Set
[**UpdateCoupon**](ManagementApi.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update a Coupon
[**UpdateCouponBatch**](ManagementApi.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update a Batch of Coupons
[**UpdateRuleset**](ManagementApi.md#UpdateRuleset) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Update a Ruleset


# **AddLoyaltyPoints**
> AddLoyaltyPoints(ctx, programID, integrationID, body)
Add points in a certain loyalty program for the specified customer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programID** | **string**|  | 
  **integrationID** | **string**|  | 
  **body** | [**LoyaltyPoints**](LoyaltyPoints.md)|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyCampaignToApplications**
> InlineResponse2003 CopyCampaignToApplications(ctx, applicationId, campaignId, body)
Copy the campaign into every specified application

Copy the campaign into every specified application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**CampaignCopy**](CampaignCopy.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCampaign**
> Campaign CreateCampaign(ctx, applicationId, body)
Create a Campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**NewCampaign**](NewCampaign.md)|  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCoupons**
> InlineResponse2001 CreateCoupons(ctx, applicationId, campaignId, body, optional)
Create Coupons

Create coupons according to some pattern. Up to 20.000 coupons can be created without a unique prefix. When a unique prefix is provided, up to 200.000 coupns can be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**NewCoupons**](NewCoupons.md)|  | 
 **optional** | ***CreateCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateCouponsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **silent** | **optional.String**| If set to &#39;yes&#39;, response will be an empty 204, otherwise a list of the coupons generated (to to 1000). | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePasswordRecoveryEmail**
> NewPasswordEmail CreatePasswordRecoveryEmail(ctx, body)
Request a password reset

Sends an email with a password recovery link to the email of an existing account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewPasswordEmail**](NewPasswordEmail.md)|  | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRuleset**
> Ruleset CreateRuleset(ctx, applicationId, campaignId, body)
Create a Ruleset



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**NewRuleset**](NewRuleset.md)|  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSession**
> Session CreateSession(ctx, body)
Create a Session



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginParams**](LoginParams.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCampaign**
> DeleteCampaign(ctx, applicationId, campaignId)
Delete a Campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCoupon**
> DeleteCoupon(ctx, applicationId, campaignId, couponId)
Delete one Coupon



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **couponId** | **string**| The ID of the coupon code to delete | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCoupons**
> DeleteCoupons(ctx, applicationId, campaignId, optional)
Delete Coupons



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***DeleteCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCouponsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReferral**
> DeleteReferral(ctx, applicationId, campaignId, referralId)
Delete one Referral



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **referralId** | **string**| The ID of the referral code to delete | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRuleset**
> DeleteRuleset(ctx, applicationId, campaignId, rulesetId)
Delete a Ruleset



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **rulesetId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessLogs**
> InlineResponse2009 GetAccessLogs(ctx, applicationId, rangeStart, rangeEnd, optional)
Get access logs for application



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **optional** | ***GetAccessLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAccessLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **optional.String**| Only return results where the request path matches the given regular expresssion. | 
 **method** | **optional.String**| Only return results where the request method matches the given regular expresssion. | 
 **status** | **optional.String**| Filter results by HTTP status codes. | 
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessLogsWithoutTotalCount**
> InlineResponse20010 GetAccessLogsWithoutTotalCount(ctx, applicationId, rangeStart, rangeEnd, optional)
Get access logs for application



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **optional** | ***GetAccessLogsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAccessLogsWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **optional.String**| Only return results where the request path matches the given regular expresssion. | 
 **method** | **optional.String**| Only return results where the request method matches the given regular expresssion. | 
 **status** | **optional.String**| Filter results by HTTP status codes. | 
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccount**
> Account GetAccount(ctx, accountId)
Get Account Details

Return the details of your companies Talon.One account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**|  | 

### Return type

[**Account**](Account.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountAnalytics**
> AccountAnalytics GetAccountAnalytics(ctx, accountId)
Get Account Analytics

Return the analytics of your companies Talon.One account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**|  | 

### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountLimits**
> AccountLimits GetAccountLimits(ctx, accountId)
Get Account Limits

Returns a list of all account limits set 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**|  | 

### Return type

[**AccountLimits**](AccountLimits.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAccessLogs**
> InlineResponse2009 GetAllAccessLogs(ctx, rangeStart, rangeEnd, optional)
Get all access logs

Fetches the access logs for the entire account. Sensitive requests (logins) are _always_ filtered from the logs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **optional** | ***GetAllAccessLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllAccessLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **optional.String**| Only return results where the request path matches the given regular expresssion. | 
 **method** | **optional.String**| Only return results where the request method matches the given regular expresssion. | 
 **status** | **optional.String**| Filter results by HTTP status codes. | 
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRoles**
> InlineResponse20028 GetAllRoles(ctx, )
Get all roles.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplication**
> Application GetApplication(ctx, applicationId)
Get Application

Get the application specified by the ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 

### Return type

[**Application**](Application.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationApiHealth**
> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId)
Get report of health of application API



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 

### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationCustomer**
> ApplicationCustomer GetApplicationCustomer(ctx, applicationId, customerId)
Get Application Customer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **customerId** | **int32**|  | 

### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationCustomers**
> InlineResponse20012 GetApplicationCustomers(ctx, applicationId)
List Application Customers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationCustomersByAttributes**
> InlineResponse20013 GetApplicationCustomersByAttributes(ctx, body)
Get a list of the customer profiles that match the given attributes

Gets a list of all the customer profiles for the account that exactly match a set of attributes.  The match is successful if all the attributes of the request are found in a profile, even if the profile has more attributes that are not present on the request.  [Customer Profile]: http://help.talon.one/customer/en/portal/articles/2525263-data-model?b_id=14115#customer-profile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationCustomerSearch**](ApplicationCustomerSearch.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[integration_auth](../README.md#integration_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationEventTypes**
> InlineResponse20019 GetApplicationEventTypes(ctx, applicationId, optional)
List Applications Event Types

Get all of the distinct values of the Event `type` property for events recorded in the application.  See also: [Track an event](/integration-api/reference/#trackEvent) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
 **optional** | ***GetApplicationEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationEventTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationEvents**
> InlineResponse20017 GetApplicationEvents(ctx, applicationId, optional)
List Applications Events

Lists all events recorded for an application. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
 **optional** | ***GetApplicationEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **type_** | **optional.String**| Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **optional.Time**| Only return events created before this date | 
 **createdAfter** | **optional.Time**| Only return events created after this date | 
 **session** | **optional.String**| Session integration ID filter for events. Must be exact match. | 
 **profile** | **optional.String**| Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **optional.String**| Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **optional.String**| Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **effectsQuery** | **optional.String**| Effects filter for events. Will perform a full-text search on the text content of the events effects, if any. | 
 **attributesQuery** | **optional.String**| Attributes filter for events. Will perform a full-text search on the text content of the events attributes, both keys and values. | 
 **ruleQuery** | **optional.String**| Rule name filter for events | 
 **campaignQuery** | **optional.String**| Campaign name filter for events | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationEventsWithoutTotalCount**
> InlineResponse20018 GetApplicationEventsWithoutTotalCount(ctx, applicationId, optional)
List Applications Events

Lists all events recorded for an application. Instead of having the total number of results in the response, this endpoint only if there are more results. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
 **optional** | ***GetApplicationEventsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationEventsWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **type_** | **optional.String**| Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **optional.Time**| Only return events created before this date | 
 **createdAfter** | **optional.Time**| Only return events created after this date | 
 **session** | **optional.String**| Session integration ID filter for events. Must be exact match. | 
 **profile** | **optional.String**| Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **optional.String**| Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **optional.String**| Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **effectsQuery** | **optional.String**| Effects filter for events. Will perform a full-text search on the text content of the events effects, if any. | 
 **attributesQuery** | **optional.String**| Attributes filter for events. Will perform a full-text search on the text content of the events attributes, both keys and values. | 
 **ruleQuery** | **optional.String**| Rule name filter for events | 
 **campaignQuery** | **optional.String**| Campaign name filter for events | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationSession**
> ApplicationSession GetApplicationSession(ctx, applicationId, sessionId)
Get Application Session



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **sessionId** | **int32**|  | 

### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationSessions**
> InlineResponse20016 GetApplicationSessions(ctx, applicationId, optional)
List Application Sessions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
 **optional** | ***GetApplicationSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **profile** | **optional.String**| Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **optional.String**| Filter by sessions with this state. Must be exact match. | 
 **coupon** | **optional.String**| Filter by sessions with this coupon. Must be exact match. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplications**
> InlineResponse2002 GetApplications(ctx, optional)
List Applications

List all application in the current account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttribute**
> Attribute GetAttribute(ctx, attributeId)
Get a custom attribute

Returns custom attribute for the account by its id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeId** | **int32**|  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaign**
> Campaign GetCampaign(ctx, applicationId, campaignId)
Get a Campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignAnalytics**
> InlineResponse20011 GetCampaignAnalytics(ctx, applicationId, campaignId, rangeStart, rangeEnd, optional)
Get analytics of campaigns



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**| The identifier for the application | 
  **campaignId** | **int32**|  | 
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **optional** | ***GetCampaignAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCampaignAnalyticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **granularity** | **optional.String**| The time interval between the results in the returned time-series. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignByAttributes**
> InlineResponse2003 GetCampaignByAttributes(ctx, applicationId, body, optional)
Get a list of all campaigns that match the given attributes

Gets a list of all the campaigns that exactly match a set of attributes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**CampaignSearch**](CampaignSearch.md)|  | 
 **optional** | ***GetCampaignByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCampaignByAttributesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **optional.String**| Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignSet**
> CampaignSet GetCampaignSet(ctx, applicationId)
List CampaignSet



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 

### Return type

[**CampaignSet**](CampaignSet.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaigns**
> InlineResponse2003 GetCampaigns(ctx, applicationId, optional)
List your Campaigns



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
 **optional** | ***GetCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **optional.String**| Filter results by the state of the campaign. | 
 **name** | **optional.String**| Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **optional.String**| Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChanges**
> InlineResponse20025 GetChanges(ctx, optional)
Get audit log for an account

Get list of changes caused by API calls for an account. Only accessible for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCoupons**
> InlineResponse2001 GetCoupons(ctx, applicationId, campaignId, optional)
List Coupons



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***GetCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCouponsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponsByAttributes**
> InlineResponse2001 GetCouponsByAttributes(ctx, applicationId, campaignId, body, optional)
Get a list of the coupons that match the given attributes

Gets a list of all the coupons that exactly match a set of attributes.  The match is successful if all the attributes of the request are found in a coupon, even if the coupon has more attributes that are not present on the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**CouponSearch**](CouponSearch.md)|  | 
 **optional** | ***GetCouponsByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCouponsByAttributesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **optional.String**| Filter results by batches of coupons | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponsByAttributesApplicationWide**
> InlineResponse2001 GetCouponsByAttributesApplicationWide(ctx, applicationId, body, optional)
Get a list of the coupons that match the given attributes in all active campaigns of an application

Gets a list of all the coupons with attributes matching the query criteria Application wide 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**CouponSearch**](CouponSearch.md)|  | 
 **optional** | ***GetCouponsByAttributesApplicationWideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCouponsByAttributesApplicationWideOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **optional.String**| Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponsWithoutTotalCount**
> InlineResponse2005 GetCouponsWithoutTotalCount(ctx, applicationId, campaignId, optional)
List Coupons



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***GetCouponsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCouponsWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerActivityReport**
> CustomerActivityReport GetCustomerActivityReport(ctx, rangeStart, rangeEnd, applicationId, customerId, optional)
Get Activity Report for Single Customer

Fetch summary report for single application customer based on a time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
  **applicationId** | **int32**|  | 
  **customerId** | **int32**|  | 
 **optional** | ***GetCustomerActivityReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerActivityReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerActivityReports**
> InlineResponse20014 GetCustomerActivityReports(ctx, rangeStart, rangeEnd, applicationId, optional)
Get Activity Reports for Application Customers

Fetch summary reports for all application customers based on a time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
  **applicationId** | **int32**|  | 
 **optional** | ***GetCustomerActivityReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerActivityReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **name** | **optional.String**| Only return reports matching the customer name | 
 **integrationId** | **optional.String**| Only return reports matching the integrationId | 
 **campaignName** | **optional.String**| Only return reports matching the campaignName | 
 **advocateName** | **optional.String**| Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerActivityReportsWithoutTotalCount**
> InlineResponse20015 GetCustomerActivityReportsWithoutTotalCount(ctx, rangeStart, rangeEnd, applicationId, optional)
Get Activity Reports for Application Customers

Fetch summary reports for all application customers based on a time range. Instead of having the total number of results in the response, this endpoint only if there are more results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rangeStart** | **time.Time**| Only return results from after this timestamp, must be an RFC3339 timestamp string | 
  **rangeEnd** | **time.Time**| Only return results from before this timestamp, must be an RFC3339 timestamp string | 
  **applicationId** | **int32**|  | 
 **optional** | ***GetCustomerActivityReportsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerActivityReportsWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **name** | **optional.String**| Only return reports matching the customer name | 
 **integrationId** | **optional.String**| Only return reports matching the integrationId | 
 **campaignName** | **optional.String**| Only return reports matching the campaignName | 
 **advocateName** | **optional.String**| Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAnalytics**
> CustomerAnalytics GetCustomerAnalytics(ctx, applicationId, customerId, optional)
Get Analytics Report for a Customer

Fetch analytics for single application customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **customerId** | **int32**|  | 
 **optional** | ***GetCustomerAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerAnalyticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerProfile**
> ApplicationCustomer GetCustomerProfile(ctx, applicationId, customerId)
Get Customer Profile



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **customerId** | **int32**|  | 

### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerProfiles**
> InlineResponse20013 GetCustomerProfiles(ctx, optional)
List Customer Profiles



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCustomerProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersByAttributes**
> InlineResponse20013 GetCustomersByAttributes(ctx, body, optional)
Get a list of the customer profiles that match the given attributes

Gets a list of all the customer profiles for the account that exactly match a set of attributes.  The match is successful if all the attributes of the request are found in a profile, even if the profile has more attributes that are not present on the request.  [Customer Profile]: http://help.talon.one/customer/en/portal/articles/2525263-data-model?b_id=14115#customer-profile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationCustomerSearch**](ApplicationCustomerSearch.md)|  | 
 **optional** | ***GetCustomersByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomersByAttributesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventTypes**
> InlineResponse20023 GetEventTypes(ctx, optional)
List Event Types

Fetch all event type definitions for your account. Each event type can be 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEventTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **optional.String**| Filter by one or more application ids separated by comma | 
 **name** | **optional.String**| Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **optional.Bool**| Include all versions of every event type. | [default to false]
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExports**
> InlineResponse20026 GetExports(ctx, optional)
Get Exports

Get a list of all past exports 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetExportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **applicationId** | **optional.Int32**|  | 
 **campaignId** | **optional.Int32**|  | 
 **entity** | **optional.String**| The name of the entity type that was exported. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImports**
> InlineResponse20027 GetImports(ctx, optional)
Get Imports

Get a list of all past imports 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetImportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetImportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoyaltyPoints**
> LoyaltyLedger GetLoyaltyPoints(ctx, programID, integrationID)
get the Loyalty Ledger for this integrationID

Get the Loyalty Ledger for this profile integration ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programID** | **string**| The identifier for the application, must be unique within the account. | 
  **integrationID** | **string**| The identifier for the application, must be unique within the account. | 

### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoyaltyProgram**
> LoyaltyProgram GetLoyaltyProgram(ctx, programID)
Get a loyalty program



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programID** | **string**|  | 

### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoyaltyPrograms**
> InlineResponse2008 GetLoyaltyPrograms(ctx, )
List all loyalty Programs



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReferrals**
> InlineResponse2006 GetReferrals(ctx, applicationId, campaignId, optional)
List Referrals



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***GetReferralsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReferralsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **code** | **optional.String**| Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **optional.String**| Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReferralsWithoutTotalCount**
> InlineResponse2007 GetReferralsWithoutTotalCount(ctx, applicationId, campaignId, optional)
List Referrals



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***GetReferralsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReferralsWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **code** | **optional.String**| Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **optional.String**| Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, roleId)
Get information for the specified role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleset**
> Ruleset GetRuleset(ctx, applicationId, campaignId, rulesetId)
Get a Ruleset



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **rulesetId** | **int32**|  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRulesets**
> InlineResponse2004 GetRulesets(ctx, applicationId, campaignId, optional)
List Campaign Rulesets



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
 **optional** | ***GetRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRulesetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, userId)
Get a single User

Retrieves the data (including an invitation code) for a user. Non-admin users can only get themselves. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**|  | 

### Return type

[**User**](User.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> InlineResponse20024 GetUsers(ctx, optional)
List Users in your account

Retrieve all users in your account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> Webhook GetWebhook(ctx, webhookId)
Get Webhook

Returns an webhook by its id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int32**|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookActivationLogs**
> InlineResponse20021 GetWebhookActivationLogs(ctx, optional)
List Webhook activation Log Entries

Webhook activation log entries would be created as soon as an integration request triggered an effect with a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookActivationLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWebhookActivationLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **integrationRequestUuid** | **optional.String**| Filter results by integration request UUID. | 
 **webhookId** | **optional.Float32**| Filter results by Webhook. | 
 **applicationId** | **optional.Float32**|  | 
 **campaignId** | **optional.Float32**| Filter results by campaign. | 
 **createdBefore** | **optional.Time**| Only return events created before this date. | 
 **createdAfter** | **optional.Time**| Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookLogs**
> InlineResponse20022 GetWebhookLogs(ctx, optional)
List Webhook Log Entries



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWebhookLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **status** | **optional.String**| Filter results by HTTP status codes. | 
 **webhookId** | **optional.Float32**| Filter results by Webhook. | 
 **applicationId** | **optional.Float32**|  | 
 **campaignId** | **optional.Float32**| Filter results by campaign. | 
 **requestUuid** | **optional.String**| Filter results by request UUID. | 
 **createdBefore** | **optional.Time**| Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. | 
 **createdAfter** | **optional.Time**| Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> InlineResponse20020 GetWebhooks(ctx, optional)
List Webhooks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWebhooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **optional.String**| Filter by one or more application ids separated by comma | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshAnalytics**
> RefreshAnalytics(ctx, )
Trigger refresh on stale analytics.

Should be used to trigger a manual refresh of analytics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLoyaltyPoints**
> RemoveLoyaltyPoints(ctx, programID, integrationID, body)
Deduct points in a certain loyalty program for the specified customer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programID** | **string**|  | 
  **integrationID** | **string**|  | 
  **body** | [**LoyaltyPoints**](LoyaltyPoints.md)|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> NewPassword ResetPassword(ctx, body)
Reset password

Consumes the supplied password reset token and updates the password for the associated account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewPassword**](NewPassword.md)|  | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCouponsAdvanced**
> InlineResponse2001 SearchCouponsAdvanced(ctx, applicationId, campaignId, body, optional)
Get a list of the coupons that match the given attributes

Gets a list of all the coupons with attributes matching the query criteria 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**AttributeQuery**](AttributeQuery.md)|  | 
 **optional** | ***SearchCouponsAdvancedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchCouponsAdvancedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **optional.String**| Filter results by batches of coupons | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCouponsAdvancedApplicationWide**
> InlineResponse2001 SearchCouponsAdvancedApplicationWide(ctx, applicationId, body, optional)
Get a list of the coupons that match the given attributes in all active campaigns of an application

Gets a list of all the coupons with attributes matching the query criteria in all active campaigns of an application 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**AttributeQuery**](AttributeQuery.md)|  | 
 **optional** | ***SearchCouponsAdvancedApplicationWideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchCouponsAdvancedApplicationWideOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **optional.String**| Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCouponsAdvancedApplicationWideWithoutTotalCount**
> InlineResponse2005 SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId, body, optional)
Get a list of the coupons that match the given attributes in all active campaigns of an application

Gets a list of all the coupons with attributes matching the query criteria in all active campaigns of an application 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**AttributeQuery**](AttributeQuery.md)|  | 
 **optional** | ***SearchCouponsAdvancedApplicationWideWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchCouponsAdvancedApplicationWideWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **optional.String**| Filter results by batches of coupons | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **optional.String**| Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCouponsAdvancedWithoutTotalCount**
> InlineResponse2005 SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId, body, optional)
Get a list of the coupons that match the given attributes

Gets a list of all the coupons with attributes matching the query criteria 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**AttributeQuery**](AttributeQuery.md)|  | 
 **optional** | ***SearchCouponsAdvancedWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchCouponsAdvancedWithoutTotalCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **optional.Int32**| Skips the given number of items when paging through large result sets. | 
 **sort** | **optional.String**| The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **optional.String**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **optional.Time**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **optional.String**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.String**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.Int32**| Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **optional.String**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **optional.Bool**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **optional.String**| Filter results by batches of coupons | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAccountLimits**
> SetAccountLimits(ctx, accountId, body)
Set account limits

sets account limits  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**|  | 
  **body** | [**AccountLimits**](AccountLimits.md)|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaign**
> Campaign UpdateCampaign(ctx, applicationId, campaignId, body)
Update a Campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**UpdateCampaign**](UpdateCampaign.md)|  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignSet**
> CampaignSet UpdateCampaignSet(ctx, applicationId, body)
Update a Campaign Set



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **body** | [**NewCampaignSet**](NewCampaignSet.md)|  | 

### Return type

[**CampaignSet**](CampaignSet.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCoupon**
> Coupon UpdateCoupon(ctx, applicationId, campaignId, couponId, body)
Update a Coupon



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **couponId** | **string**| The ID of the coupon code to update | 
  **body** | [**UpdateCoupon**](UpdateCoupon.md)|  | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCouponBatch**
> UpdateCouponBatch(ctx, applicationId, campaignId, body)
Update a Batch of Coupons



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **body** | [**UpdateCouponBatch**](UpdateCouponBatch.md)|  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRuleset**
> Ruleset UpdateRuleset(ctx, applicationId, campaignId, rulesetId, body)
Update a Ruleset



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**|  | 
  **campaignId** | **int32**|  | 
  **rulesetId** | **int32**|  | 
  **body** | [**NewRuleset**](NewRuleset.md)|  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

