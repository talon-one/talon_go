# Talon.One Go SDK

This SDK supports all of the operations of Talon.One's Integration API and Management API.

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./talon"
```

## Getting started

### Integration API

The following code shows an example of using the Integration API:

```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"

	talon "github.com/talon-one/talon_go/v8"
)

func main() {
	configuration := talon.NewConfiguration()
	// Set API base path
	configuration.Servers = talon.ServerConfigurations{
		{
			// Notice that there is no trailing '/'
			URL:         "https://yourbaseurl.talon.one",
			Description: "Talon.One's API base URL",
		},
	}
	// If you wish to inject a custom implementation of HTTPClient
	// configuration.HTTPClient = &customHTTPClient

	integrationClient := talon.NewAPIClient(configuration)

	// Create integration authentication context using api key
	integrationAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"Authorization": {
			Prefix: "ApiKey-v1",
			Key:    "fd1fd219b1e953a6b2700e8034de5bfc877462ae106127311ddd710978654312",
		},
	})

	// Instantiating a NewCustomerSessionV2 struct
	newCustomerSession := talon.NewCustomerSessionV2{
		// You can use either struct literals
		ProfileId:   talon.PtrString("DEADBEEF"),
		CouponCodes: &[]string{"Cool-Stuff!"},
	}

	// Or alternatively, using the relevant setter in a later stage in the code
	newCustomerSession.SetCartItems([]talon.CartItem{
		{
			Name:     talon.PtrString("Pad Thai - Veggie"),
			Sku:      "pad-332",
			Quantity: 1,
			Price:    talon.PtrFloat32(5.5),
			Category: talon.PtrString("Noodles"),
		},
		{
			Name:     talon.PtrString("Chang"),
			Sku:      "chang-br-42",
			Quantity: 1,
			Price:    talon.PtrFloat32(2.3),
			Category: talon.PtrString("Beverages"),
		},
	})

	// Instantiating a new IntegrationRequest
	integrationRequest := talon.IntegrationRequest{
		CustomerSession: newCustomerSession,
	}

	// Optional list of requested information to be present on the response.
  // See docs/IntegrationRequest.md for full list of supported values
	// integrationRequest.SetResponseContent([]string{
	// 	"customerSession",
	// 	"customerProfile",
	// 	"loyalty",
	// })

	// Create/update a customer session using `UpdateCustomerSessionV2` function
	integrationState, _, err := integrationClient.IntegrationApi.
		UpdateCustomerSessionV2(integrationAuthContext, "deetdoot_2").
		Body(integrationRequest).
		Execute()

	if err != nil {
		fmt.Printf("ERROR while calling UpdateCustomerSessionV2: %s\n", err)
		return
	}
	fmt.Printf("%#v\n", integrationState)

	// Parsing the returned effects list, please consult https://developers.talon.one/Integration-API/handling-effects-v2 for the full list of effects and their corresponding properties
	for _, effect := range integrationState.GetEffects() {
		effectType := effect.GetEffectType()
		switch {
		case "setDiscount" == effectType:
			// Initiating right props instance according to the effect type
			effectProps := talon.SetDiscountEffectProps{}
			if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
				fmt.Printf("ERROR while decoding 'setDiscount' props: %s\n", err)
				continue
			}

			// Access the specific effect's properties
			fmt.Printf("Set a discount '%s' of %2.3f\n", effectProps.GetName(), effectProps.GetValue())
		case "acceptCoupon" == effectType:
			// Initiating right props instance according to the effect type
			effectProps := talon.AcceptCouponEffectProps{}
			if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
				fmt.Printf("ERROR while decoding props: %s\n", err)
				continue
			}

			// Work with AcceptCouponEffectProps' properties
			// ...
		default:
			fmt.Printf("Encounter unknown effect type: %s\n", effectType)
		}
	}
}

// quick decoding of props-map into our library structures using JSON marshaling,
// or alternatively using a library like https://github.com/mitchellh/mapstructure
func decodeHelper(propsMap map[string]interface{}, v interface{}) error {
	propsJSON, err := json.Marshal(propsMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(propsJSON, v)
}
```

### Management API

The following code shows an example of using the Management API:

```golang
package main

import (
	"context"
	"fmt"

	talon "github.com/talon-one/talon_go/v8"
)

func main() {
	configuration := talon.NewConfiguration()
	// Set API base path
	configuration.Servers = talon.ServerConfigurations{
		{
			// Notice that there is no trailing '/'
			URL:         "https://yourbaseurl.talon.one",
			Description: "Talon.One's API base URL",
		},
	}
	// If you wish to inject a custom implementation of HTTPClient
	// configuration.HTTPClient = &customHTTPClient

	managementClient := talon.NewAPIClient(configuration)

	// Create integration authentication context using the logged-in session
	managerAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"Authorization": talon.APIKey{
			Prefix: "ManagementKey-v1",
			Key:    "2f0dce055da01ae595005d7d79154bae7448d319d5fc7c5b2951fadd6ba1ea07",
		},
	})

	// Calling `GetApplication` function with the desired id (7)
	application, response, err := managementClient.ManagementApi.
		GetApplication(managerAuthContext, 7).
		Execute()

	if err != nil {
		fmt.Printf("ERROR while calling GetApplication: %s\n", err)
		return
	}

	fmt.Printf("%#v\n\n", application)
	fmt.Printf("%#v\n\n", response)
}
```

## Documentation for API endpoints

All URLs are relative to `https://yourbaseurl.talon.one`.

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IntegrationApi* | [**CreateAudienceV2**](docs/IntegrationApi.md#createaudiencev2) | **Post** /v2/audiences | Create audience
*IntegrationApi* | [**CreateCouponReservation**](docs/IntegrationApi.md#createcouponreservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
*IntegrationApi* | [**CreateReferral**](docs/IntegrationApi.md#createreferral) | **Post** /v1/referrals | Create referral code for an advocate
*IntegrationApi* | [**CreateReferralsForMultipleAdvocates**](docs/IntegrationApi.md#createreferralsformultipleadvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates
*IntegrationApi* | [**DeleteAudienceMembershipsV2**](docs/IntegrationApi.md#deleteaudiencemembershipsv2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
*IntegrationApi* | [**DeleteAudienceV2**](docs/IntegrationApi.md#deleteaudiencev2) | **Delete** /v2/audiences/{audienceId} | Delete audience
*IntegrationApi* | [**DeleteCouponReservation**](docs/IntegrationApi.md#deletecouponreservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
*IntegrationApi* | [**DeleteCustomerData**](docs/IntegrationApi.md#deletecustomerdata) | **Delete** /v1/customer_data/{integrationId} | Delete customer&#39;s personal data
*IntegrationApi* | [**GenerateLoyaltyCard**](docs/IntegrationApi.md#generateloyaltycard) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards | Generate loyalty card
*IntegrationApi* | [**GetCustomerAchievementHistory**](docs/IntegrationApi.md#getcustomerachievementhistory) | **Get** /v1/customer_profiles/{integrationId}/achievements/{achievementId} | List customer&#39;s achievement history
*IntegrationApi* | [**GetCustomerAchievements**](docs/IntegrationApi.md#getcustomerachievements) | **Get** /v1/customer_profiles/{integrationId}/achievements | List customer&#39;s available achievements
*IntegrationApi* | [**GetCustomerInventory**](docs/IntegrationApi.md#getcustomerinventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
*IntegrationApi* | [**GetCustomerSession**](docs/IntegrationApi.md#getcustomersession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
*IntegrationApi* | [**GetLoyaltyBalances**](docs/IntegrationApi.md#getloyaltybalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/balances | Get customer&#39;s loyalty balances
*IntegrationApi* | [**GetLoyaltyCardBalances**](docs/IntegrationApi.md#getloyaltycardbalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/balances | Get card&#39;s point balances
*IntegrationApi* | [**GetLoyaltyCardPoints**](docs/IntegrationApi.md#getloyaltycardpoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/points | List card&#39;s unused loyalty points
*IntegrationApi* | [**GetLoyaltyCardTransactions**](docs/IntegrationApi.md#getloyaltycardtransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transactions | List card&#39;s transactions
*IntegrationApi* | [**GetLoyaltyProgramProfilePoints**](docs/IntegrationApi.md#getloyaltyprogramprofilepoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/points | List customer&#39;s unused loyalty points
*IntegrationApi* | [**GetLoyaltyProgramProfileTransactions**](docs/IntegrationApi.md#getloyaltyprogramprofiletransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/transactions | List customer&#39;s loyalty transactions
*IntegrationApi* | [**GetReservedCustomers**](docs/IntegrationApi.md#getreservedcustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List customers that have this coupon reserved
*IntegrationApi* | [**LinkLoyaltyCardToProfile**](docs/IntegrationApi.md#linkloyaltycardtoprofile) | **Post** /v2/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/link_profile | Link customer profile to card
*IntegrationApi* | [**ReopenCustomerSession**](docs/IntegrationApi.md#reopencustomersession) | **Put** /v2/customer_sessions/{customerSessionId}/reopen | Reopen customer session
*IntegrationApi* | [**ReturnCartItems**](docs/IntegrationApi.md#returncartitems) | **Post** /v2/customer_sessions/{customerSessionId}/returns | Return cart items
*IntegrationApi* | [**SyncCatalog**](docs/IntegrationApi.md#synccatalog) | **Put** /v1/catalogs/{catalogId}/sync | Sync cart item catalog
*IntegrationApi* | [**TrackEventV2**](docs/IntegrationApi.md#trackeventv2) | **Post** /v2/events | Track event
*IntegrationApi* | [**UpdateAudienceCustomersAttributes**](docs/IntegrationApi.md#updateaudiencecustomersattributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
*IntegrationApi* | [**UpdateAudienceV2**](docs/IntegrationApi.md#updateaudiencev2) | **Put** /v2/audiences/{audienceId} | Update audience name
*IntegrationApi* | [**UpdateCustomerProfileAudiences**](docs/IntegrationApi.md#updatecustomerprofileaudiences) | **Post** /v2/customer_audiences | Update multiple customer profiles&#39; audiences
*IntegrationApi* | [**UpdateCustomerProfileV2**](docs/IntegrationApi.md#updatecustomerprofilev2) | **Put** /v2/customer_profiles/{integrationId} | Update customer profile
*IntegrationApi* | [**UpdateCustomerProfilesV2**](docs/IntegrationApi.md#updatecustomerprofilesv2) | **Put** /v2/customer_profiles | Update multiple customer profiles
*IntegrationApi* | [**UpdateCustomerSessionV2**](docs/IntegrationApi.md#updatecustomersessionv2) | **Put** /v2/customer_sessions/{customerSessionId} | Update customer session
*ManagementApi* | [**ActivateUserByEmail**](docs/ManagementApi.md#activateuserbyemail) | **Post** /v1/users/activate | Enable user by email address
*ManagementApi* | [**AddLoyaltyCardPoints**](docs/ManagementApi.md#addloyaltycardpoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/add_points | Add points to card
*ManagementApi* | [**AddLoyaltyPoints**](docs/ManagementApi.md#addloyaltypoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/add_points | Add points to customer profile
*ManagementApi* | [**CopyCampaignToApplications**](docs/ManagementApi.md#copycampaigntoapplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into the specified Application
*ManagementApi* | [**CreateAccountCollection**](docs/ManagementApi.md#createaccountcollection) | **Post** /v1/collections | Create account-level collection
*ManagementApi* | [**CreateAchievement**](docs/ManagementApi.md#createachievement) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | Create achievement
*ManagementApi* | [**CreateAdditionalCost**](docs/ManagementApi.md#createadditionalcost) | **Post** /v1/additional_costs | Create additional cost
*ManagementApi* | [**CreateAttribute**](docs/ManagementApi.md#createattribute) | **Post** /v1/attributes | Create custom attribute
*ManagementApi* | [**CreateBatchLoyaltyCards**](docs/ManagementApi.md#createbatchloyaltycards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards/batch | Create loyalty cards
*ManagementApi* | [**CreateCampaignFromTemplate**](docs/ManagementApi.md#createcampaignfromtemplate) | **Post** /v1/applications/{applicationId}/create_campaign_from_template | Create campaign from campaign template
*ManagementApi* | [**CreateCollection**](docs/ManagementApi.md#createcollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | Create campaign-level collection
*ManagementApi* | [**CreateCoupons**](docs/ManagementApi.md#createcoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
*ManagementApi* | [**CreateCouponsAsync**](docs/ManagementApi.md#createcouponsasync) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_async | Create coupons asynchronously
*ManagementApi* | [**CreateCouponsDeletionJob**](docs/ManagementApi.md#createcouponsdeletionjob) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_deletion_jobs | Creates a coupon deletion job
*ManagementApi* | [**CreateCouponsForMultipleRecipients**](docs/ManagementApi.md#createcouponsformultiplerecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
*ManagementApi* | [**CreateInviteEmail**](docs/ManagementApi.md#createinviteemail) | **Post** /v1/invite_emails | Resend invitation email
*ManagementApi* | [**CreateInviteV2**](docs/ManagementApi.md#createinvitev2) | **Post** /v2/invites | Invite user
*ManagementApi* | [**CreatePasswordRecoveryEmail**](docs/ManagementApi.md#createpasswordrecoveryemail) | **Post** /v1/password_recovery_emails | Request a password reset
*ManagementApi* | [**CreateSession**](docs/ManagementApi.md#createsession) | **Post** /v1/sessions | Create session
*ManagementApi* | [**CreateStore**](docs/ManagementApi.md#createstore) | **Post** /v1/applications/{applicationId}/stores | Create store
*ManagementApi* | [**DeactivateUserByEmail**](docs/ManagementApi.md#deactivateuserbyemail) | **Post** /v1/users/deactivate | Disable user by email address
*ManagementApi* | [**DeductLoyaltyCardPoints**](docs/ManagementApi.md#deductloyaltycardpoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/deduct_points | Deduct points from card
*ManagementApi* | [**DeleteAccountCollection**](docs/ManagementApi.md#deleteaccountcollection) | **Delete** /v1/collections/{collectionId} | Delete account-level collection
*ManagementApi* | [**DeleteAchievement**](docs/ManagementApi.md#deleteachievement) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Delete achievement
*ManagementApi* | [**DeleteCampaign**](docs/ManagementApi.md#deletecampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
*ManagementApi* | [**DeleteCollection**](docs/ManagementApi.md#deletecollection) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Delete campaign-level collection
*ManagementApi* | [**DeleteCoupon**](docs/ManagementApi.md#deletecoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
*ManagementApi* | [**DeleteCoupons**](docs/ManagementApi.md#deletecoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
*ManagementApi* | [**DeleteLoyaltyCard**](docs/ManagementApi.md#deleteloyaltycard) | **Delete** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Delete loyalty card
*ManagementApi* | [**DeleteReferral**](docs/ManagementApi.md#deletereferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
*ManagementApi* | [**DeleteStore**](docs/ManagementApi.md#deletestore) | **Delete** /v1/applications/{applicationId}/stores/{storeId} | Delete store
*ManagementApi* | [**DeleteUser**](docs/ManagementApi.md#deleteuser) | **Delete** /v1/users/{userId} | Delete user
*ManagementApi* | [**DeleteUserByEmail**](docs/ManagementApi.md#deleteuserbyemail) | **Post** /v1/users/delete | Delete user by email address
*ManagementApi* | [**DestroySession**](docs/ManagementApi.md#destroysession) | **Delete** /v1/sessions | Destroy session
*ManagementApi* | [**DisconnectCampaignStores**](docs/ManagementApi.md#disconnectcampaignstores) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores | Disconnect stores
*ManagementApi* | [**ExportAccountCollectionItems**](docs/ManagementApi.md#exportaccountcollectionitems) | **Get** /v1/collections/{collectionId}/export | Export account-level collection&#39;s items
*ManagementApi* | [**ExportAchievements**](docs/ManagementApi.md#exportachievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId}/export | Export achievement customer data
*ManagementApi* | [**ExportAudiencesMemberships**](docs/ManagementApi.md#exportaudiencesmemberships) | **Get** /v1/audiences/{audienceId}/memberships/export | Export audience members
*ManagementApi* | [**ExportCampaignStores**](docs/ManagementApi.md#exportcampaignstores) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/export | Export stores
*ManagementApi* | [**ExportCollectionItems**](docs/ManagementApi.md#exportcollectionitems) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/export | Export campaign-level collection&#39;s items
*ManagementApi* | [**ExportCoupons**](docs/ManagementApi.md#exportcoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons
*ManagementApi* | [**ExportCustomerSessions**](docs/ManagementApi.md#exportcustomersessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions
*ManagementApi* | [**ExportCustomersTiers**](docs/ManagementApi.md#exportcustomerstiers) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customers_tiers | Export customers&#39; tier data
*ManagementApi* | [**ExportEffects**](docs/ManagementApi.md#exporteffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects
*ManagementApi* | [**ExportLoyaltyBalance**](docs/ManagementApi.md#exportloyaltybalance) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balance | Export customer loyalty balance to CSV
*ManagementApi* | [**ExportLoyaltyBalances**](docs/ManagementApi.md#exportloyaltybalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balances | Export customer loyalty balances
*ManagementApi* | [**ExportLoyaltyCardBalances**](docs/ManagementApi.md#exportloyaltycardbalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_card_balances | Export all card transaction logs
*ManagementApi* | [**ExportLoyaltyCardLedger**](docs/ManagementApi.md#exportloyaltycardledger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/export_log | Export card&#39;s ledger log
*ManagementApi* | [**ExportLoyaltyCards**](docs/ManagementApi.md#exportloyaltycards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/export | Export loyalty cards
*ManagementApi* | [**ExportLoyaltyLedger**](docs/ManagementApi.md#exportloyaltyledger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/export_log | Export customer&#39;s transaction logs
*ManagementApi* | [**ExportPoolGiveaways**](docs/ManagementApi.md#exportpoolgiveaways) | **Get** /v1/giveaways/pools/{poolId}/export | Export giveaway codes of a giveaway pool
*ManagementApi* | [**ExportReferrals**](docs/ManagementApi.md#exportreferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals
*ManagementApi* | [**GetAccessLogsWithoutTotalCount**](docs/ManagementApi.md#getaccesslogswithouttotalcount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
*ManagementApi* | [**GetAccount**](docs/ManagementApi.md#getaccount) | **Get** /v1/accounts/{accountId} | Get account details
*ManagementApi* | [**GetAccountAnalytics**](docs/ManagementApi.md#getaccountanalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
*ManagementApi* | [**GetAccountCollection**](docs/ManagementApi.md#getaccountcollection) | **Get** /v1/collections/{collectionId} | Get account-level collection
*ManagementApi* | [**GetAchievement**](docs/ManagementApi.md#getachievement) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Get achievement
*ManagementApi* | [**GetAdditionalCost**](docs/ManagementApi.md#getadditionalcost) | **Get** /v1/additional_costs/{additionalCostId} | Get additional cost
*ManagementApi* | [**GetAdditionalCosts**](docs/ManagementApi.md#getadditionalcosts) | **Get** /v1/additional_costs | List additional costs
*ManagementApi* | [**GetApplication**](docs/ManagementApi.md#getapplication) | **Get** /v1/applications/{applicationId} | Get Application
*ManagementApi* | [**GetApplicationApiHealth**](docs/ManagementApi.md#getapplicationapihealth) | **Get** /v1/applications/{applicationId}/health_report | Get Application health
*ManagementApi* | [**GetApplicationCustomer**](docs/ManagementApi.md#getapplicationcustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
*ManagementApi* | [**GetApplicationCustomerFriends**](docs/ManagementApi.md#getapplicationcustomerfriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
*ManagementApi* | [**GetApplicationCustomers**](docs/ManagementApi.md#getapplicationcustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
*ManagementApi* | [**GetApplicationCustomersByAttributes**](docs/ManagementApi.md#getapplicationcustomersbyattributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
*ManagementApi* | [**GetApplicationEventTypes**](docs/ManagementApi.md#getapplicationeventtypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
*ManagementApi* | [**GetApplicationEventsWithoutTotalCount**](docs/ManagementApi.md#getapplicationeventswithouttotalcount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
*ManagementApi* | [**GetApplicationSession**](docs/ManagementApi.md#getapplicationsession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
*ManagementApi* | [**GetApplicationSessions**](docs/ManagementApi.md#getapplicationsessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
*ManagementApi* | [**GetApplications**](docs/ManagementApi.md#getapplications) | **Get** /v1/applications | List Applications
*ManagementApi* | [**GetAttribute**](docs/ManagementApi.md#getattribute) | **Get** /v1/attributes/{attributeId} | Get custom attribute
*ManagementApi* | [**GetAttributes**](docs/ManagementApi.md#getattributes) | **Get** /v1/attributes | List custom attributes
*ManagementApi* | [**GetAudienceMemberships**](docs/ManagementApi.md#getaudiencememberships) | **Get** /v1/audiences/{audienceId}/memberships | List audience members
*ManagementApi* | [**GetAudiences**](docs/ManagementApi.md#getaudiences) | **Get** /v1/audiences | List audiences
*ManagementApi* | [**GetAudiencesAnalytics**](docs/ManagementApi.md#getaudiencesanalytics) | **Get** /v1/audiences/analytics | List audience analytics
*ManagementApi* | [**GetCampaign**](docs/ManagementApi.md#getcampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
*ManagementApi* | [**GetCampaignAnalytics**](docs/ManagementApi.md#getcampaignanalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
*ManagementApi* | [**GetCampaignByAttributes**](docs/ManagementApi.md#getcampaignbyattributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
*ManagementApi* | [**GetCampaignGroup**](docs/ManagementApi.md#getcampaigngroup) | **Get** /v1/campaign_groups/{campaignGroupId} | Get campaign access group
*ManagementApi* | [**GetCampaignGroups**](docs/ManagementApi.md#getcampaigngroups) | **Get** /v1/campaign_groups | List campaign access groups
*ManagementApi* | [**GetCampaignTemplates**](docs/ManagementApi.md#getcampaigntemplates) | **Get** /v1/campaign_templates | List campaign templates
*ManagementApi* | [**GetCampaigns**](docs/ManagementApi.md#getcampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
*ManagementApi* | [**GetChanges**](docs/ManagementApi.md#getchanges) | **Get** /v1/changes | Get audit logs for an account
*ManagementApi* | [**GetCollection**](docs/ManagementApi.md#getcollection) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Get campaign-level collection
*ManagementApi* | [**GetCollectionItems**](docs/ManagementApi.md#getcollectionitems) | **Get** /v1/collections/{collectionId}/items | Get collection items
*ManagementApi* | [**GetCouponsWithoutTotalCount**](docs/ManagementApi.md#getcouponswithouttotalcount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
*ManagementApi* | [**GetCustomerActivityReport**](docs/ManagementApi.md#getcustomeractivityreport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
*ManagementApi* | [**GetCustomerActivityReportsWithoutTotalCount**](docs/ManagementApi.md#getcustomeractivityreportswithouttotalcount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
*ManagementApi* | [**GetCustomerAnalytics**](docs/ManagementApi.md#getcustomeranalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
*ManagementApi* | [**GetCustomerProfile**](docs/ManagementApi.md#getcustomerprofile) | **Get** /v1/customers/{customerId} | Get customer profile
*ManagementApi* | [**GetCustomerProfileAchievementProgress**](docs/ManagementApi.md#getcustomerprofileachievementprogress) | **Get** /v1/applications/{applicationId}/achievement_progress/{integrationId} | List customer achievements
*ManagementApi* | [**GetCustomerProfiles**](docs/ManagementApi.md#getcustomerprofiles) | **Get** /v1/customers/no_total | List customer profiles
*ManagementApi* | [**GetCustomersByAttributes**](docs/ManagementApi.md#getcustomersbyattributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
*ManagementApi* | [**GetDashboardStatistics**](docs/ManagementApi.md#getdashboardstatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/dashboard | Get statistics for loyalty dashboard
*ManagementApi* | [**GetEventTypes**](docs/ManagementApi.md#geteventtypes) | **Get** /v1/event_types | List event types
*ManagementApi* | [**GetExports**](docs/ManagementApi.md#getexports) | **Get** /v1/exports | Get exports
*ManagementApi* | [**GetLoyaltyCard**](docs/ManagementApi.md#getloyaltycard) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Get loyalty card
*ManagementApi* | [**GetLoyaltyCardTransactionLogs**](docs/ManagementApi.md#getloyaltycardtransactionlogs) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/logs | List card&#39;s transactions
*ManagementApi* | [**GetLoyaltyCards**](docs/ManagementApi.md#getloyaltycards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards | List loyalty cards
*ManagementApi* | [**GetLoyaltyPoints**](docs/ManagementApi.md#getloyaltypoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId} | Get customer&#39;s full loyalty ledger
*ManagementApi* | [**GetLoyaltyProgram**](docs/ManagementApi.md#getloyaltyprogram) | **Get** /v1/loyalty_programs/{loyaltyProgramId} | Get loyalty program
*ManagementApi* | [**GetLoyaltyProgramTransactions**](docs/ManagementApi.md#getloyaltyprogramtransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/transactions | List loyalty program transactions
*ManagementApi* | [**GetLoyaltyPrograms**](docs/ManagementApi.md#getloyaltyprograms) | **Get** /v1/loyalty_programs | List loyalty programs
*ManagementApi* | [**GetLoyaltyStatistics**](docs/ManagementApi.md#getloyaltystatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/statistics | Get loyalty program statistics
*ManagementApi* | [**GetMessageLogs**](docs/ManagementApi.md#getmessagelogs) | **Get** /v1/message_logs | List message log entries
*ManagementApi* | [**GetReferralsWithoutTotalCount**](docs/ManagementApi.md#getreferralswithouttotalcount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
*ManagementApi* | [**GetRoleV2**](docs/ManagementApi.md#getrolev2) | **Get** /v2/roles/{roleId} | Get role
*ManagementApi* | [**GetRuleset**](docs/ManagementApi.md#getruleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
*ManagementApi* | [**GetRulesets**](docs/ManagementApi.md#getrulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
*ManagementApi* | [**GetStore**](docs/ManagementApi.md#getstore) | **Get** /v1/applications/{applicationId}/stores/{storeId} | Get store
*ManagementApi* | [**GetUser**](docs/ManagementApi.md#getuser) | **Get** /v1/users/{userId} | Get user
*ManagementApi* | [**GetUsers**](docs/ManagementApi.md#getusers) | **Get** /v1/users | List users in account
*ManagementApi* | [**GetWebhook**](docs/ManagementApi.md#getwebhook) | **Get** /v1/webhooks/{webhookId} | Get webhook
*ManagementApi* | [**GetWebhookActivationLogs**](docs/ManagementApi.md#getwebhookactivationlogs) | **Get** /v1/webhook_activation_logs | List webhook activation log entries
*ManagementApi* | [**GetWebhookLogs**](docs/ManagementApi.md#getwebhooklogs) | **Get** /v1/webhook_logs | List webhook log entries
*ManagementApi* | [**GetWebhooks**](docs/ManagementApi.md#getwebhooks) | **Get** /v1/webhooks | List webhooks
*ManagementApi* | [**ImportAccountCollection**](docs/ManagementApi.md#importaccountcollection) | **Post** /v1/collections/{collectionId}/import | Import data into existing account-level collection
*ManagementApi* | [**ImportAllowedList**](docs/ManagementApi.md#importallowedlist) | **Post** /v1/attributes/{attributeId}/allowed_list/import | Import allowed values for attribute
*ManagementApi* | [**ImportAudiencesMemberships**](docs/ManagementApi.md#importaudiencesmemberships) | **Post** /v1/audiences/{audienceId}/memberships/import | Import audience members
*ManagementApi* | [**ImportCampaignStores**](docs/ManagementApi.md#importcampaignstores) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/import | Import stores
*ManagementApi* | [**ImportCollection**](docs/ManagementApi.md#importcollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import data into existing campaign-level collection
*ManagementApi* | [**ImportCoupons**](docs/ManagementApi.md#importcoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons
*ManagementApi* | [**ImportLoyaltyCards**](docs/ManagementApi.md#importloyaltycards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_cards | Import loyalty cards
*ManagementApi* | [**ImportLoyaltyCustomersTiers**](docs/ManagementApi.md#importloyaltycustomerstiers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_customers_tiers | Import customers into loyalty tiers
*ManagementApi* | [**ImportLoyaltyPoints**](docs/ManagementApi.md#importloyaltypoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_points | Import loyalty points
*ManagementApi* | [**ImportPoolGiveaways**](docs/ManagementApi.md#importpoolgiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
*ManagementApi* | [**ImportReferrals**](docs/ManagementApi.md#importreferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals
*ManagementApi* | [**InviteUserExternal**](docs/ManagementApi.md#inviteuserexternal) | **Post** /v1/users/invite | Invite user from identity provider
*ManagementApi* | [**ListAccountCollections**](docs/ManagementApi.md#listaccountcollections) | **Get** /v1/collections | List collections in account
*ManagementApi* | [**ListAchievements**](docs/ManagementApi.md#listachievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | List achievements
*ManagementApi* | [**ListAllRolesV2**](docs/ManagementApi.md#listallrolesv2) | **Get** /v2/roles | List roles
*ManagementApi* | [**ListCatalogItems**](docs/ManagementApi.md#listcatalogitems) | **Get** /v1/catalogs/{catalogId}/items | List items in a catalog
*ManagementApi* | [**ListCollections**](docs/ManagementApi.md#listcollections) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | List collections in campaign
*ManagementApi* | [**ListCollectionsInApplication**](docs/ManagementApi.md#listcollectionsinapplication) | **Get** /v1/applications/{applicationId}/collections | List collections in Application
*ManagementApi* | [**ListStores**](docs/ManagementApi.md#liststores) | **Get** /v1/applications/{applicationId}/stores | List stores
*ManagementApi* | [**OktaEventHandlerChallenge**](docs/ManagementApi.md#oktaeventhandlerchallenge) | **Get** /v1/provisioning/okta | Validate Okta API ownership
*ManagementApi* | [**RemoveLoyaltyPoints**](docs/ManagementApi.md#removeloyaltypoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/deduct_points | Deduct points from customer profile
*ManagementApi* | [**ResetPassword**](docs/ManagementApi.md#resetpassword) | **Post** /v1/reset_password | Reset password
*ManagementApi* | [**ScimCreateUser**](docs/ManagementApi.md#scimcreateuser) | **Post** /v1/provisioning/scim/Users | Create SCIM user
*ManagementApi* | [**ScimDeleteUser**](docs/ManagementApi.md#scimdeleteuser) | **Delete** /v1/provisioning/scim/Users/{userId} | Delete SCIM user
*ManagementApi* | [**ScimGetResourceTypes**](docs/ManagementApi.md#scimgetresourcetypes) | **Get** /v1/provisioning/scim/ResourceTypes | List supported SCIM resource types
*ManagementApi* | [**ScimGetSchemas**](docs/ManagementApi.md#scimgetschemas) | **Get** /v1/provisioning/scim/Schemas | List supported SCIM schemas
*ManagementApi* | [**ScimGetServiceProviderConfig**](docs/ManagementApi.md#scimgetserviceproviderconfig) | **Get** /v1/provisioning/scim/ServiceProviderConfig | Get SCIM service provider configuration
*ManagementApi* | [**ScimGetUser**](docs/ManagementApi.md#scimgetuser) | **Get** /v1/provisioning/scim/Users/{userId} | Get SCIM user
*ManagementApi* | [**ScimGetUsers**](docs/ManagementApi.md#scimgetusers) | **Get** /v1/provisioning/scim/Users | List SCIM users
*ManagementApi* | [**ScimPatchUser**](docs/ManagementApi.md#scimpatchuser) | **Patch** /v1/provisioning/scim/Users/{userId} | Update SCIM user attributes
*ManagementApi* | [**ScimReplaceUserAttributes**](docs/ManagementApi.md#scimreplaceuserattributes) | **Put** /v1/provisioning/scim/Users/{userId} | Update SCIM user
*ManagementApi* | [**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](docs/ManagementApi.md#searchcouponsadvancedapplicationwidewithouttotalcount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes (without total count)
*ManagementApi* | [**SearchCouponsAdvancedWithoutTotalCount**](docs/ManagementApi.md#searchcouponsadvancedwithouttotalcount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes in campaign (without total count)
*ManagementApi* | [**TransferLoyaltyCard**](docs/ManagementApi.md#transferloyaltycard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transfer | Transfer card data
*ManagementApi* | [**UpdateAccountCollection**](docs/ManagementApi.md#updateaccountcollection) | **Put** /v1/collections/{collectionId} | Update account-level collection
*ManagementApi* | [**UpdateAchievement**](docs/ManagementApi.md#updateachievement) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Update achievement
*ManagementApi* | [**UpdateAdditionalCost**](docs/ManagementApi.md#updateadditionalcost) | **Put** /v1/additional_costs/{additionalCostId} | Update additional cost
*ManagementApi* | [**UpdateAttribute**](docs/ManagementApi.md#updateattribute) | **Put** /v1/attributes/{attributeId} | Update custom attribute
*ManagementApi* | [**UpdateCampaign**](docs/ManagementApi.md#updatecampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
*ManagementApi* | [**UpdateCollection**](docs/ManagementApi.md#updatecollection) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Update campaign-level collection&#39;s description
*ManagementApi* | [**UpdateCoupon**](docs/ManagementApi.md#updatecoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
*ManagementApi* | [**UpdateCouponBatch**](docs/ManagementApi.md#updatecouponbatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update coupons
*ManagementApi* | [**UpdateLoyaltyCard**](docs/ManagementApi.md#updateloyaltycard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Update loyalty card status
*ManagementApi* | [**UpdateReferral**](docs/ManagementApi.md#updatereferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
*ManagementApi* | [**UpdateRoleV2**](docs/ManagementApi.md#updaterolev2) | **Put** /v2/roles/{roleId} | Update role
*ManagementApi* | [**UpdateStore**](docs/ManagementApi.md#updatestore) | **Put** /v1/applications/{applicationId}/stores/{storeId} | Update store
*ManagementApi* | [**UpdateUser**](docs/ManagementApi.md#updateuser) | **Put** /v1/users/{userId} | Update user

## Documentation for models

- [AcceptCouponEffectProps](docs/AcceptCouponEffectProps.md)
- [AcceptReferralEffectProps](docs/AcceptReferralEffectProps.md)
- [AccessLogEntry](docs/AccessLogEntry.md)
- [Account](docs/Account.md)
- [AccountAdditionalCost](docs/AccountAdditionalCost.md)
- [AccountAnalytics](docs/AccountAnalytics.md)
- [AccountDashboardStatistic](docs/AccountDashboardStatistic.md)
- [AccountDashboardStatisticCampaigns](docs/AccountDashboardStatisticCampaigns.md)
- [AccountDashboardStatisticDiscount](docs/AccountDashboardStatisticDiscount.md)
- [AccountDashboardStatisticLoyaltyPoints](docs/AccountDashboardStatisticLoyaltyPoints.md)
- [AccountDashboardStatisticReferrals](docs/AccountDashboardStatisticReferrals.md)
- [AccountDashboardStatisticRevenue](docs/AccountDashboardStatisticRevenue.md)
- [AccountEntity](docs/AccountEntity.md)
- [AccountLimits](docs/AccountLimits.md)
- [Achievement](docs/Achievement.md)
- [AchievementAdditionalProperties](docs/AchievementAdditionalProperties.md)
- [AchievementBase](docs/AchievementBase.md)
- [AchievementProgress](docs/AchievementProgress.md)
- [AchievementProgressWithDefinition](docs/AchievementProgressWithDefinition.md)
- [AchievementStatusEntry](docs/AchievementStatusEntry.md)
- [ActivateUserRequest](docs/ActivateUserRequest.md)
- [AddFreeItemEffectProps](docs/AddFreeItemEffectProps.md)
- [AddItemCatalogAction](docs/AddItemCatalogAction.md)
- [AddLoyaltyPoints](docs/AddLoyaltyPoints.md)
- [AddLoyaltyPointsEffectProps](docs/AddLoyaltyPointsEffectProps.md)
- [AddToAudienceEffectProps](docs/AddToAudienceEffectProps.md)
- [AddedDeductedPointsNotificationPolicy](docs/AddedDeductedPointsNotificationPolicy.md)
- [AdditionalCampaignProperties](docs/AdditionalCampaignProperties.md)
- [AdditionalCost](docs/AdditionalCost.md)
- [AnalyticsDataPoint](docs/AnalyticsDataPoint.md)
- [AnalyticsDataPointWithTrend](docs/AnalyticsDataPointWithTrend.md)
- [AnalyticsDataPointWithTrendAndInfluencedRate](docs/AnalyticsDataPointWithTrendAndInfluencedRate.md)
- [AnalyticsDataPointWithTrendAndUplift](docs/AnalyticsDataPointWithTrendAndUplift.md)
- [AnalyticsProduct](docs/AnalyticsProduct.md)
- [AnalyticsSku](docs/AnalyticsSku.md)
- [ApiError](docs/ApiError.md)
- [Application](docs/Application.md)
- [ApplicationAnalyticsDataPoint](docs/ApplicationAnalyticsDataPoint.md)
- [ApplicationApiHealth](docs/ApplicationApiHealth.md)
- [ApplicationApiKey](docs/ApplicationApiKey.md)
- [ApplicationCampaignAnalytics](docs/ApplicationCampaignAnalytics.md)
- [ApplicationCampaignStats](docs/ApplicationCampaignStats.md)
- [ApplicationCif](docs/ApplicationCif.md)
- [ApplicationCifExpression](docs/ApplicationCifExpression.md)
- [ApplicationCifReferences](docs/ApplicationCifReferences.md)
- [ApplicationCustomer](docs/ApplicationCustomer.md)
- [ApplicationCustomerEntity](docs/ApplicationCustomerEntity.md)
- [ApplicationEntity](docs/ApplicationEntity.md)
- [ApplicationEvent](docs/ApplicationEvent.md)
- [ApplicationNotification](docs/ApplicationNotification.md)
- [ApplicationReferee](docs/ApplicationReferee.md)
- [ApplicationSession](docs/ApplicationSession.md)
- [ApplicationSessionEntity](docs/ApplicationSessionEntity.md)
- [ApplicationStoreEntity](docs/ApplicationStoreEntity.md)
- [AsyncCouponCreationResponse](docs/AsyncCouponCreationResponse.md)
- [AsyncCouponDeletionJobResponse](docs/AsyncCouponDeletionJobResponse.md)
- [Attribute](docs/Attribute.md)
- [AttributesMandatory](docs/AttributesMandatory.md)
- [AttributesSettings](docs/AttributesSettings.md)
- [Audience](docs/Audience.md)
- [AudienceAnalytics](docs/AudienceAnalytics.md)
- [AudienceCustomer](docs/AudienceCustomer.md)
- [AudienceIntegrationId](docs/AudienceIntegrationId.md)
- [AudienceMembership](docs/AudienceMembership.md)
- [AwardGiveawayEffectProps](docs/AwardGiveawayEffectProps.md)
- [BaseCampaign](docs/BaseCampaign.md)
- [BaseLoyaltyProgram](docs/BaseLoyaltyProgram.md)
- [BaseNotification](docs/BaseNotification.md)
- [BaseNotificationEntity](docs/BaseNotificationEntity.md)
- [BaseNotificationWebhook](docs/BaseNotificationWebhook.md)
- [BaseNotifications](docs/BaseNotifications.md)
- [BaseSamlConnection](docs/BaseSamlConnection.md)
- [Binding](docs/Binding.md)
- [BulkApplicationNotification](docs/BulkApplicationNotification.md)
- [BulkCampaignNotification](docs/BulkCampaignNotification.md)
- [BulkOperationOnCampaigns](docs/BulkOperationOnCampaigns.md)
- [Campaign](docs/Campaign.md)
- [CampaignActivationRequest](docs/CampaignActivationRequest.md)
- [CampaignAnalytics](docs/CampaignAnalytics.md)
- [CampaignBudget](docs/CampaignBudget.md)
- [CampaignCollection](docs/CampaignCollection.md)
- [CampaignCollectionEditedNotification](docs/CampaignCollectionEditedNotification.md)
- [CampaignCollectionWithoutPayload](docs/CampaignCollectionWithoutPayload.md)
- [CampaignCopy](docs/CampaignCopy.md)
- [CampaignCreatedNotification](docs/CampaignCreatedNotification.md)
- [CampaignDeletedNotification](docs/CampaignDeletedNotification.md)
- [CampaignDetail](docs/CampaignDetail.md)
- [CampaignEditedNotification](docs/CampaignEditedNotification.md)
- [CampaignEntity](docs/CampaignEntity.md)
- [CampaignEvaluationGroup](docs/CampaignEvaluationGroup.md)
- [CampaignEvaluationPosition](docs/CampaignEvaluationPosition.md)
- [CampaignEvaluationTreeChangedNotification](docs/CampaignEvaluationTreeChangedNotification.md)
- [CampaignGroup](docs/CampaignGroup.md)
- [CampaignGroupEntity](docs/CampaignGroupEntity.md)
- [CampaignNotification](docs/CampaignNotification.md)
- [CampaignNotificationPolicy](docs/CampaignNotificationPolicy.md)
- [CampaignRulesetChangedNotification](docs/CampaignRulesetChangedNotification.md)
- [CampaignSearch](docs/CampaignSearch.md)
- [CampaignSet](docs/CampaignSet.md)
- [CampaignSetBranchNode](docs/CampaignSetBranchNode.md)
- [CampaignSetLeafNode](docs/CampaignSetLeafNode.md)
- [CampaignSetNode](docs/CampaignSetNode.md)
- [CampaignStateChangedNotification](docs/CampaignStateChangedNotification.md)
- [CampaignStoreBudget](docs/CampaignStoreBudget.md)
- [CampaignStoreBudgetLimitConfig](docs/CampaignStoreBudgetLimitConfig.md)
- [CampaignTemplate](docs/CampaignTemplate.md)
- [CampaignTemplateCollection](docs/CampaignTemplateCollection.md)
- [CampaignTemplateCouponReservationSettings](docs/CampaignTemplateCouponReservationSettings.md)
- [CampaignTemplateParams](docs/CampaignTemplateParams.md)
- [CampaignVersions](docs/CampaignVersions.md)
- [CardAddedDeductedPointsNotificationPolicy](docs/CardAddedDeductedPointsNotificationPolicy.md)
- [CardExpiringPointsNotificationPolicy](docs/CardExpiringPointsNotificationPolicy.md)
- [CardExpiringPointsNotificationTrigger](docs/CardExpiringPointsNotificationTrigger.md)
- [CardLedgerPointsEntryIntegrationApi](docs/CardLedgerPointsEntryIntegrationApi.md)
- [CardLedgerTransactionLogEntry](docs/CardLedgerTransactionLogEntry.md)
- [CardLedgerTransactionLogEntryIntegrationApi](docs/CardLedgerTransactionLogEntryIntegrationApi.md)
- [CartItem](docs/CartItem.md)
- [Catalog](docs/Catalog.md)
- [CatalogAction](docs/CatalogAction.md)
- [CatalogActionFilter](docs/CatalogActionFilter.md)
- [CatalogItem](docs/CatalogItem.md)
- [CatalogSyncRequest](docs/CatalogSyncRequest.md)
- [CatalogsStrikethroughNotificationPolicy](docs/CatalogsStrikethroughNotificationPolicy.md)
- [Change](docs/Change.md)
- [ChangeLoyaltyTierLevelEffectProps](docs/ChangeLoyaltyTierLevelEffectProps.md)
- [ChangeProfilePassword](docs/ChangeProfilePassword.md)
- [CodeGeneratorSettings](docs/CodeGeneratorSettings.md)
- [Collection](docs/Collection.md)
- [CollectionItem](docs/CollectionItem.md)
- [CollectionWithoutPayload](docs/CollectionWithoutPayload.md)
- [Coupon](docs/Coupon.md)
- [CouponConstraints](docs/CouponConstraints.md)
- [CouponCreatedEffectProps](docs/CouponCreatedEffectProps.md)
- [CouponCreationJob](docs/CouponCreationJob.md)
- [CouponDeletionFilters](docs/CouponDeletionFilters.md)
- [CouponDeletionJob](docs/CouponDeletionJob.md)
- [CouponLimitConfigs](docs/CouponLimitConfigs.md)
- [CouponRejectionReason](docs/CouponRejectionReason.md)
- [CouponReservations](docs/CouponReservations.md)
- [CouponSearch](docs/CouponSearch.md)
- [CouponValue](docs/CouponValue.md)
- [CouponsNotificationPolicy](docs/CouponsNotificationPolicy.md)
- [CreateAchievement](docs/CreateAchievement.md)
- [CreateApplicationApiKey](docs/CreateApplicationApiKey.md)
- [CreateManagementKey](docs/CreateManagementKey.md)
- [CreateTemplateCampaign](docs/CreateTemplateCampaign.md)
- [CreateTemplateCampaignResponse](docs/CreateTemplateCampaignResponse.md)
- [CustomEffect](docs/CustomEffect.md)
- [CustomEffectProps](docs/CustomEffectProps.md)
- [CustomerActivityReport](docs/CustomerActivityReport.md)
- [CustomerAnalytics](docs/CustomerAnalytics.md)
- [CustomerInventory](docs/CustomerInventory.md)
- [CustomerProfile](docs/CustomerProfile.md)
- [CustomerProfileAudienceRequest](docs/CustomerProfileAudienceRequest.md)
- [CustomerProfileAudienceRequestItem](docs/CustomerProfileAudienceRequestItem.md)
- [CustomerProfileIntegrationRequestV2](docs/CustomerProfileIntegrationRequestV2.md)
- [CustomerProfileIntegrationResponseV2](docs/CustomerProfileIntegrationResponseV2.md)
- [CustomerProfileSearchQuery](docs/CustomerProfileSearchQuery.md)
- [CustomerProfileUpdateV2Response](docs/CustomerProfileUpdateV2Response.md)
- [CustomerSession](docs/CustomerSession.md)
- [CustomerSessionV2](docs/CustomerSessionV2.md)
- [DeactivateUserRequest](docs/DeactivateUserRequest.md)
- [DeductLoyaltyPoints](docs/DeductLoyaltyPoints.md)
- [DeductLoyaltyPointsEffectProps](docs/DeductLoyaltyPointsEffectProps.md)
- [DeleteUserRequest](docs/DeleteUserRequest.md)
- [Effect](docs/Effect.md)
- [EffectEntity](docs/EffectEntity.md)
- [EmailEntity](docs/EmailEntity.md)
- [Endpoint](docs/Endpoint.md)
- [Entity](docs/Entity.md)
- [EntityWithTalangVisibleId](docs/EntityWithTalangVisibleId.md)
- [Environment](docs/Environment.md)
- [ErrorEffectProps](docs/ErrorEffectProps.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [ErrorResponseWithStatus](docs/ErrorResponseWithStatus.md)
- [ErrorSource](docs/ErrorSource.md)
- [EvaluableCampaignIds](docs/EvaluableCampaignIds.md)
- [Event](docs/Event.md)
- [EventType](docs/EventType.md)
- [EventV2](docs/EventV2.md)
- [ExpiringCouponsNotificationPolicy](docs/ExpiringCouponsNotificationPolicy.md)
- [ExpiringCouponsNotificationTrigger](docs/ExpiringCouponsNotificationTrigger.md)
- [ExpiringPointsNotificationPolicy](docs/ExpiringPointsNotificationPolicy.md)
- [ExpiringPointsNotificationTrigger](docs/ExpiringPointsNotificationTrigger.md)
- [Export](docs/Export.md)
- [FeatureFlag](docs/FeatureFlag.md)
- [FeaturesFeed](docs/FeaturesFeed.md)
- [FuncArgDef](docs/FuncArgDef.md)
- [FunctionDef](docs/FunctionDef.md)
- [GenerateCampaignDescription](docs/GenerateCampaignDescription.md)
- [GenerateCampaignTags](docs/GenerateCampaignTags.md)
- [GenerateItemFilterDescription](docs/GenerateItemFilterDescription.md)
- [GenerateLoyaltyCard](docs/GenerateLoyaltyCard.md)
- [GenerateRuleTitle](docs/GenerateRuleTitle.md)
- [GenerateRuleTitleRule](docs/GenerateRuleTitleRule.md)
- [GetIntegrationCouponRequest](docs/GetIntegrationCouponRequest.md)
- [Giveaway](docs/Giveaway.md)
- [GiveawaysPool](docs/GiveawaysPool.md)
- [HiddenConditionsEffects](docs/HiddenConditionsEffects.md)
- [IdentifiableEntity](docs/IdentifiableEntity.md)
- [Import](docs/Import.md)
- [ImportEntity](docs/ImportEntity.md)
- [IncreaseAchievementProgressEffectProps](docs/IncreaseAchievementProgressEffectProps.md)
- [InlineResponse200](docs/InlineResponse200.md)
- [InlineResponse2001](docs/InlineResponse2001.md)
- [InlineResponse20010](docs/InlineResponse20010.md)
- [InlineResponse20011](docs/InlineResponse20011.md)
- [InlineResponse20012](docs/InlineResponse20012.md)
- [InlineResponse20013](docs/InlineResponse20013.md)
- [InlineResponse20014](docs/InlineResponse20014.md)
- [InlineResponse20015](docs/InlineResponse20015.md)
- [InlineResponse20016](docs/InlineResponse20016.md)
- [InlineResponse20017](docs/InlineResponse20017.md)
- [InlineResponse20018](docs/InlineResponse20018.md)
- [InlineResponse20019](docs/InlineResponse20019.md)
- [InlineResponse2002](docs/InlineResponse2002.md)
- [InlineResponse20020](docs/InlineResponse20020.md)
- [InlineResponse20021](docs/InlineResponse20021.md)
- [InlineResponse20022](docs/InlineResponse20022.md)
- [InlineResponse20023](docs/InlineResponse20023.md)
- [InlineResponse20024](docs/InlineResponse20024.md)
- [InlineResponse20025](docs/InlineResponse20025.md)
- [InlineResponse20026](docs/InlineResponse20026.md)
- [InlineResponse20027](docs/InlineResponse20027.md)
- [InlineResponse20028](docs/InlineResponse20028.md)
- [InlineResponse20029](docs/InlineResponse20029.md)
- [InlineResponse2003](docs/InlineResponse2003.md)
- [InlineResponse20030](docs/InlineResponse20030.md)
- [InlineResponse20031](docs/InlineResponse20031.md)
- [InlineResponse20032](docs/InlineResponse20032.md)
- [InlineResponse20033](docs/InlineResponse20033.md)
- [InlineResponse20034](docs/InlineResponse20034.md)
- [InlineResponse20035](docs/InlineResponse20035.md)
- [InlineResponse20036](docs/InlineResponse20036.md)
- [InlineResponse20037](docs/InlineResponse20037.md)
- [InlineResponse20038](docs/InlineResponse20038.md)
- [InlineResponse20039](docs/InlineResponse20039.md)
- [InlineResponse2004](docs/InlineResponse2004.md)
- [InlineResponse20040](docs/InlineResponse20040.md)
- [InlineResponse20041](docs/InlineResponse20041.md)
- [InlineResponse20042](docs/InlineResponse20042.md)
- [InlineResponse20043](docs/InlineResponse20043.md)
- [InlineResponse20044](docs/InlineResponse20044.md)
- [InlineResponse20045](docs/InlineResponse20045.md)
- [InlineResponse20046](docs/InlineResponse20046.md)
- [InlineResponse20047](docs/InlineResponse20047.md)
- [InlineResponse20048](docs/InlineResponse20048.md)
- [InlineResponse20049](docs/InlineResponse20049.md)
- [InlineResponse2005](docs/InlineResponse2005.md)
- [InlineResponse2006](docs/InlineResponse2006.md)
- [InlineResponse2007](docs/InlineResponse2007.md)
- [InlineResponse2008](docs/InlineResponse2008.md)
- [InlineResponse2009](docs/InlineResponse2009.md)
- [InlineResponse201](docs/InlineResponse201.md)
- [IntegrationCoupon](docs/IntegrationCoupon.md)
- [IntegrationCustomerSessionResponse](docs/IntegrationCustomerSessionResponse.md)
- [IntegrationEntity](docs/IntegrationEntity.md)
- [IntegrationEvent](docs/IntegrationEvent.md)
- [IntegrationEventV2Request](docs/IntegrationEventV2Request.md)
- [IntegrationProfileEntity](docs/IntegrationProfileEntity.md)
- [IntegrationRequest](docs/IntegrationRequest.md)
- [IntegrationState](docs/IntegrationState.md)
- [IntegrationStateV2](docs/IntegrationStateV2.md)
- [IntegrationStoreEntity](docs/IntegrationStoreEntity.md)
- [InventoryCoupon](docs/InventoryCoupon.md)
- [InventoryReferral](docs/InventoryReferral.md)
- [ItemAttribute](docs/ItemAttribute.md)
- [LedgerEntry](docs/LedgerEntry.md)
- [LedgerInfo](docs/LedgerInfo.md)
- [LedgerPointsEntryIntegrationApi](docs/LedgerPointsEntryIntegrationApi.md)
- [LedgerTransactionLogEntryIntegrationApi](docs/LedgerTransactionLogEntryIntegrationApi.md)
- [LibraryAttribute](docs/LibraryAttribute.md)
- [LimitConfig](docs/LimitConfig.md)
- [LimitCounter](docs/LimitCounter.md)
- [ListCampaignStoreBudgets](docs/ListCampaignStoreBudgets.md)
- [ListCampaignStoreBudgetsStore](docs/ListCampaignStoreBudgetsStore.md)
- [LoginParams](docs/LoginParams.md)
- [Loyalty](docs/Loyalty.md)
- [LoyaltyBalance](docs/LoyaltyBalance.md)
- [LoyaltyBalanceWithTier](docs/LoyaltyBalanceWithTier.md)
- [LoyaltyBalances](docs/LoyaltyBalances.md)
- [LoyaltyBalancesWithTiers](docs/LoyaltyBalancesWithTiers.md)
- [LoyaltyCard](docs/LoyaltyCard.md)
- [LoyaltyCardBalances](docs/LoyaltyCardBalances.md)
- [LoyaltyCardBatch](docs/LoyaltyCardBatch.md)
- [LoyaltyCardBatchResponse](docs/LoyaltyCardBatchResponse.md)
- [LoyaltyCardProfileRegistration](docs/LoyaltyCardProfileRegistration.md)
- [LoyaltyCardRegistration](docs/LoyaltyCardRegistration.md)
- [LoyaltyDashboardData](docs/LoyaltyDashboardData.md)
- [LoyaltyDashboardPointsBreakdown](docs/LoyaltyDashboardPointsBreakdown.md)
- [LoyaltyLedger](docs/LoyaltyLedger.md)
- [LoyaltyLedgerEntry](docs/LoyaltyLedgerEntry.md)
- [LoyaltyLedgerEntryFlags](docs/LoyaltyLedgerEntryFlags.md)
- [LoyaltyLedgerTransactions](docs/LoyaltyLedgerTransactions.md)
- [LoyaltyMembership](docs/LoyaltyMembership.md)
- [LoyaltyProgram](docs/LoyaltyProgram.md)
- [LoyaltyProgramBalance](docs/LoyaltyProgramBalance.md)
- [LoyaltyProgramEntity](docs/LoyaltyProgramEntity.md)
- [LoyaltyProgramLedgers](docs/LoyaltyProgramLedgers.md)
- [LoyaltyProgramTransaction](docs/LoyaltyProgramTransaction.md)
- [LoyaltySubLedger](docs/LoyaltySubLedger.md)
- [LoyaltyTier](docs/LoyaltyTier.md)
- [ManagementKey](docs/ManagementKey.md)
- [ManagerConfig](docs/ManagerConfig.md)
- [MessageLogEntries](docs/MessageLogEntries.md)
- [MessageLogEntry](docs/MessageLogEntry.md)
- [MessageLogRequest](docs/MessageLogRequest.md)
- [MessageLogResponse](docs/MessageLogResponse.md)
- [MessageTest](docs/MessageTest.md)
- [Meta](docs/Meta.md)
- [MultiApplicationEntity](docs/MultiApplicationEntity.md)
- [MultipleAttribute](docs/MultipleAttribute.md)
- [MultipleAudiences](docs/MultipleAudiences.md)
- [MultipleAudiencesItem](docs/MultipleAudiencesItem.md)
- [MultipleCustomerProfileIntegrationRequest](docs/MultipleCustomerProfileIntegrationRequest.md)
- [MultipleCustomerProfileIntegrationRequestItem](docs/MultipleCustomerProfileIntegrationRequestItem.md)
- [MultipleCustomerProfileIntegrationResponseV2](docs/MultipleCustomerProfileIntegrationResponseV2.md)
- [MultipleNewAttribute](docs/MultipleNewAttribute.md)
- [MultipleNewAudiences](docs/MultipleNewAudiences.md)
- [MutableEntity](docs/MutableEntity.md)
- [NewAccount](docs/NewAccount.md)
- [NewAccountSignUp](docs/NewAccountSignUp.md)
- [NewAdditionalCost](docs/NewAdditionalCost.md)
- [NewAppWideCouponDeletionJob](docs/NewAppWideCouponDeletionJob.md)
- [NewApplication](docs/NewApplication.md)
- [NewApplicationApiKey](docs/NewApplicationApiKey.md)
- [NewApplicationCif](docs/NewApplicationCif.md)
- [NewApplicationCifExpression](docs/NewApplicationCifExpression.md)
- [NewAttribute](docs/NewAttribute.md)
- [NewAudience](docs/NewAudience.md)
- [NewBaseNotification](docs/NewBaseNotification.md)
- [NewCampaign](docs/NewCampaign.md)
- [NewCampaignCollection](docs/NewCampaignCollection.md)
- [NewCampaignEvaluationGroup](docs/NewCampaignEvaluationGroup.md)
- [NewCampaignGroup](docs/NewCampaignGroup.md)
- [NewCampaignSet](docs/NewCampaignSet.md)
- [NewCampaignStoreBudget](docs/NewCampaignStoreBudget.md)
- [NewCampaignStoreBudgetStoreLimit](docs/NewCampaignStoreBudgetStoreLimit.md)
- [NewCampaignTemplate](docs/NewCampaignTemplate.md)
- [NewCatalog](docs/NewCatalog.md)
- [NewCollection](docs/NewCollection.md)
- [NewCouponCreationJob](docs/NewCouponCreationJob.md)
- [NewCouponDeletionJob](docs/NewCouponDeletionJob.md)
- [NewCoupons](docs/NewCoupons.md)
- [NewCouponsForMultipleRecipients](docs/NewCouponsForMultipleRecipients.md)
- [NewCustomEffect](docs/NewCustomEffect.md)
- [NewCustomerProfile](docs/NewCustomerProfile.md)
- [NewCustomerSession](docs/NewCustomerSession.md)
- [NewCustomerSessionV2](docs/NewCustomerSessionV2.md)
- [NewEvent](docs/NewEvent.md)
- [NewEventType](docs/NewEventType.md)
- [NewExternalInvitation](docs/NewExternalInvitation.md)
- [NewGiveawaysPool](docs/NewGiveawaysPool.md)
- [NewInternalAudience](docs/NewInternalAudience.md)
- [NewInvitation](docs/NewInvitation.md)
- [NewInviteEmail](docs/NewInviteEmail.md)
- [NewLoyaltyProgram](docs/NewLoyaltyProgram.md)
- [NewLoyaltyTier](docs/NewLoyaltyTier.md)
- [NewManagementKey](docs/NewManagementKey.md)
- [NewMessageTest](docs/NewMessageTest.md)
- [NewMultipleAudiencesItem](docs/NewMultipleAudiencesItem.md)
- [NewNotificationWebhook](docs/NewNotificationWebhook.md)
- [NewOutgoingIntegrationWebhook](docs/NewOutgoingIntegrationWebhook.md)
- [NewPassword](docs/NewPassword.md)
- [NewPasswordEmail](docs/NewPasswordEmail.md)
- [NewPicklist](docs/NewPicklist.md)
- [NewReferral](docs/NewReferral.md)
- [NewReferralsForMultipleAdvocates](docs/NewReferralsForMultipleAdvocates.md)
- [NewReturn](docs/NewReturn.md)
- [NewRevisionVersion](docs/NewRevisionVersion.md)
- [NewRole](docs/NewRole.md)
- [NewRoleV2](docs/NewRoleV2.md)
- [NewRuleset](docs/NewRuleset.md)
- [NewSamlConnection](docs/NewSamlConnection.md)
- [NewStore](docs/NewStore.md)
- [NewTemplateDef](docs/NewTemplateDef.md)
- [NewUser](docs/NewUser.md)
- [NewWebhook](docs/NewWebhook.md)
- [Notification](docs/Notification.md)
- [NotificationActivation](docs/NotificationActivation.md)
- [NotificationListItem](docs/NotificationListItem.md)
- [OktaEvent](docs/OktaEvent.md)
- [OktaEventPayload](docs/OktaEventPayload.md)
- [OktaEventPayloadData](docs/OktaEventPayloadData.md)
- [OktaEventTarget](docs/OktaEventTarget.md)
- [OneTimeCode](docs/OneTimeCode.md)
- [OutgoingIntegrationBrazePolicy](docs/OutgoingIntegrationBrazePolicy.md)
- [OutgoingIntegrationCleverTapPolicy](docs/OutgoingIntegrationCleverTapPolicy.md)
- [OutgoingIntegrationConfiguration](docs/OutgoingIntegrationConfiguration.md)
- [OutgoingIntegrationIterablePolicy](docs/OutgoingIntegrationIterablePolicy.md)
- [OutgoingIntegrationMoEngagePolicy](docs/OutgoingIntegrationMoEngagePolicy.md)
- [OutgoingIntegrationTemplate](docs/OutgoingIntegrationTemplate.md)
- [OutgoingIntegrationTemplateWithConfigurationDetails](docs/OutgoingIntegrationTemplateWithConfigurationDetails.md)
- [OutgoingIntegrationTemplates](docs/OutgoingIntegrationTemplates.md)
- [OutgoingIntegrationType](docs/OutgoingIntegrationType.md)
- [OutgoingIntegrationTypes](docs/OutgoingIntegrationTypes.md)
- [PatchItemCatalogAction](docs/PatchItemCatalogAction.md)
- [PatchManyItemsCatalogAction](docs/PatchManyItemsCatalogAction.md)
- [PendingPointsNotificationPolicy](docs/PendingPointsNotificationPolicy.md)
- [Picklist](docs/Picklist.md)
- [Product](docs/Product.md)
- [ProductSearchMatch](docs/ProductSearchMatch.md)
- [ProductUnitAnalytics](docs/ProductUnitAnalytics.md)
- [ProductUnitAnalyticsDataPoint](docs/ProductUnitAnalyticsDataPoint.md)
- [ProductUnitAnalyticsTotals](docs/ProductUnitAnalyticsTotals.md)
- [ProfileAudiencesChanges](docs/ProfileAudiencesChanges.md)
- [ProjectedTier](docs/ProjectedTier.md)
- [RedeemReferralEffectProps](docs/RedeemReferralEffectProps.md)
- [Referral](docs/Referral.md)
- [ReferralConstraints](docs/ReferralConstraints.md)
- [ReferralCreatedEffectProps](docs/ReferralCreatedEffectProps.md)
- [ReferralRejectionReason](docs/ReferralRejectionReason.md)
- [RejectCouponEffectProps](docs/RejectCouponEffectProps.md)
- [RejectReferralEffectProps](docs/RejectReferralEffectProps.md)
- [RemoveFromAudienceEffectProps](docs/RemoveFromAudienceEffectProps.md)
- [RemoveItemCatalogAction](docs/RemoveItemCatalogAction.md)
- [RemoveManyItemsCatalogAction](docs/RemoveManyItemsCatalogAction.md)
- [ReopenSessionResponse](docs/ReopenSessionResponse.md)
- [ReserveCouponEffectProps](docs/ReserveCouponEffectProps.md)
- [Return](docs/Return.md)
- [ReturnIntegrationRequest](docs/ReturnIntegrationRequest.md)
- [ReturnedCartItem](docs/ReturnedCartItem.md)
- [Revision](docs/Revision.md)
- [RevisionActivation](docs/RevisionActivation.md)
- [RevisionActivationRequest](docs/RevisionActivationRequest.md)
- [RevisionVersion](docs/RevisionVersion.md)
- [Role](docs/Role.md)
- [RoleAssign](docs/RoleAssign.md)
- [RoleMembership](docs/RoleMembership.md)
- [RoleV2](docs/RoleV2.md)
- [RoleV2ApplicationDetails](docs/RoleV2ApplicationDetails.md)
- [RoleV2Base](docs/RoleV2Base.md)
- [RoleV2PermissionSet](docs/RoleV2PermissionSet.md)
- [RoleV2Permissions](docs/RoleV2Permissions.md)
- [RoleV2RolesGroup](docs/RoleV2RolesGroup.md)
- [RollbackAddedLoyaltyPointsEffectProps](docs/RollbackAddedLoyaltyPointsEffectProps.md)
- [RollbackCouponEffectProps](docs/RollbackCouponEffectProps.md)
- [RollbackDeductedLoyaltyPointsEffectProps](docs/RollbackDeductedLoyaltyPointsEffectProps.md)
- [RollbackDiscountEffectProps](docs/RollbackDiscountEffectProps.md)
- [RollbackIncreasedAchievementProgressEffectProps](docs/RollbackIncreasedAchievementProgressEffectProps.md)
- [RollbackReferralEffectProps](docs/RollbackReferralEffectProps.md)
- [Rule](docs/Rule.md)
- [RuleFailureReason](docs/RuleFailureReason.md)
- [Ruleset](docs/Ruleset.md)
- [SamlConnection](docs/SamlConnection.md)
- [SamlConnectionInternal](docs/SamlConnectionInternal.md)
- [SamlConnectionMetadata](docs/SamlConnectionMetadata.md)
- [SamlLoginEndpoint](docs/SamlLoginEndpoint.md)
- [ScimBaseUser](docs/ScimBaseUser.md)
- [ScimBaseUserName](docs/ScimBaseUserName.md)
- [ScimNewUser](docs/ScimNewUser.md)
- [ScimPatchOperation](docs/ScimPatchOperation.md)
- [ScimPatchRequest](docs/ScimPatchRequest.md)
- [ScimResource](docs/ScimResource.md)
- [ScimResourceTypesListResponse](docs/ScimResourceTypesListResponse.md)
- [ScimSchemaResource](docs/ScimSchemaResource.md)
- [ScimSchemasListResponse](docs/ScimSchemasListResponse.md)
- [ScimServiceProviderConfigResponse](docs/ScimServiceProviderConfigResponse.md)
- [ScimServiceProviderConfigResponseBulk](docs/ScimServiceProviderConfigResponseBulk.md)
- [ScimServiceProviderConfigResponseChangePassword](docs/ScimServiceProviderConfigResponseChangePassword.md)
- [ScimServiceProviderConfigResponseFilter](docs/ScimServiceProviderConfigResponseFilter.md)
- [ScimServiceProviderConfigResponsePatch](docs/ScimServiceProviderConfigResponsePatch.md)
- [ScimServiceProviderConfigResponseSort](docs/ScimServiceProviderConfigResponseSort.md)
- [ScimUser](docs/ScimUser.md)
- [ScimUsersListResponse](docs/ScimUsersListResponse.md)
- [Session](docs/Session.md)
- [SetDiscountEffectProps](docs/SetDiscountEffectProps.md)
- [SetDiscountPerAdditionalCostEffectProps](docs/SetDiscountPerAdditionalCostEffectProps.md)
- [SetDiscountPerAdditionalCostPerItemEffectProps](docs/SetDiscountPerAdditionalCostPerItemEffectProps.md)
- [SetDiscountPerItemEffectProps](docs/SetDiscountPerItemEffectProps.md)
- [ShowBundleMetadataEffectProps](docs/ShowBundleMetadataEffectProps.md)
- [ShowNotificationEffectProps](docs/ShowNotificationEffectProps.md)
- [SkuUnitAnalytics](docs/SkuUnitAnalytics.md)
- [SkuUnitAnalyticsDataPoint](docs/SkuUnitAnalyticsDataPoint.md)
- [SlotDef](docs/SlotDef.md)
- [SsoConfig](docs/SsoConfig.md)
- [Store](docs/Store.md)
- [StrikethroughChangedItem](docs/StrikethroughChangedItem.md)
- [StrikethroughCustomEffectPerItemProps](docs/StrikethroughCustomEffectPerItemProps.md)
- [StrikethroughDebugResponse](docs/StrikethroughDebugResponse.md)
- [StrikethroughEffect](docs/StrikethroughEffect.md)
- [StrikethroughLabelingNotification](docs/StrikethroughLabelingNotification.md)
- [StrikethroughSetDiscountPerItemEffectProps](docs/StrikethroughSetDiscountPerItemEffectProps.md)
- [StrikethroughTrigger](docs/StrikethroughTrigger.md)
- [SummaryCampaignStoreBudget](docs/SummaryCampaignStoreBudget.md)
- [TalangAttribute](docs/TalangAttribute.md)
- [TalangAttributeVisibility](docs/TalangAttributeVisibility.md)
- [TemplateArgDef](docs/TemplateArgDef.md)
- [TemplateDef](docs/TemplateDef.md)
- [TemplateLimitConfig](docs/TemplateLimitConfig.md)
- [Tier](docs/Tier.md)
- [TierDowngradeNotificationPolicy](docs/TierDowngradeNotificationPolicy.md)
- [TierUpgradeNotificationPolicy](docs/TierUpgradeNotificationPolicy.md)
- [TierWillDowngradeNotificationPolicy](docs/TierWillDowngradeNotificationPolicy.md)
- [TierWillDowngradeNotificationTrigger](docs/TierWillDowngradeNotificationTrigger.md)
- [TimePoint](docs/TimePoint.md)
- [TrackEventV2Response](docs/TrackEventV2Response.md)
- [TransferLoyaltyCard](docs/TransferLoyaltyCard.md)
- [TriggerWebhookEffectProps](docs/TriggerWebhookEffectProps.md)
- [TwoFaConfig](docs/TwoFaConfig.md)
- [UpdateAccount](docs/UpdateAccount.md)
- [UpdateAchievement](docs/UpdateAchievement.md)
- [UpdateApplication](docs/UpdateApplication.md)
- [UpdateApplicationApiKey](docs/UpdateApplicationApiKey.md)
- [UpdateApplicationCif](docs/UpdateApplicationCif.md)
- [UpdateAttributeEffectProps](docs/UpdateAttributeEffectProps.md)
- [UpdateAudience](docs/UpdateAudience.md)
- [UpdateCampaign](docs/UpdateCampaign.md)
- [UpdateCampaignCollection](docs/UpdateCampaignCollection.md)
- [UpdateCampaignEvaluationGroup](docs/UpdateCampaignEvaluationGroup.md)
- [UpdateCampaignGroup](docs/UpdateCampaignGroup.md)
- [UpdateCampaignTemplate](docs/UpdateCampaignTemplate.md)
- [UpdateCatalog](docs/UpdateCatalog.md)
- [UpdateCollection](docs/UpdateCollection.md)
- [UpdateCoupon](docs/UpdateCoupon.md)
- [UpdateCouponBatch](docs/UpdateCouponBatch.md)
- [UpdateCustomEffect](docs/UpdateCustomEffect.md)
- [UpdateLoyaltyCard](docs/UpdateLoyaltyCard.md)
- [UpdateLoyaltyProgram](docs/UpdateLoyaltyProgram.md)
- [UpdateLoyaltyProgramTier](docs/UpdateLoyaltyProgramTier.md)
- [UpdatePicklist](docs/UpdatePicklist.md)
- [UpdateReferral](docs/UpdateReferral.md)
- [UpdateReferralBatch](docs/UpdateReferralBatch.md)
- [UpdateRole](docs/UpdateRole.md)
- [UpdateRoleV2](docs/UpdateRoleV2.md)
- [UpdateStore](docs/UpdateStore.md)
- [UpdateUser](docs/UpdateUser.md)
- [User](docs/User.md)
- [UserEntity](docs/UserEntity.md)
- [ValueMap](docs/ValueMap.md)
- [Webhook](docs/Webhook.md)
- [WebhookActivationLogEntry](docs/WebhookActivationLogEntry.md)
- [WebhookLogEntry](docs/WebhookLogEntry.md)
- [WebhookWithOutgoingIntegrationDetails](docs/WebhookWithOutgoingIntegrationDetails.md)
- [WillAwardGiveawayEffectProps](docs/WillAwardGiveawayEffectProps.md)

## Authorization

### api_key_v1

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

### management_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

### manager_auth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

## Utility methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
