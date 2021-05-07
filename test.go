package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	talon "github.com/talon-one/talon_go"
)

func main() {
	configuration := talon.NewConfiguration()
	configuration.Servers = talon.ServerConfigurations{
		{
			URL:         "http://localhost:9000",
			Description: "Talon.One's API base URL",
		},
	}
	// configuration.Debug = true
	configuration.DefaultHeader = map[string]string{
		"something": "something",
	}
	// configuration.HTTPClient = &customHTTPClient

	/**********************************************************************
	**************************** Management API ***************************
	**********************************************************************/
	// managementClient := talon.NewAPIClient(configuration)

	// session, _, err := managementClient.ManagementApi.
	// 	CreateSession(context.Background()).
	// 	Body(talon.LoginParams{
	// 		Email:    "demo@talon.one",
	// 		Password: "Demo1234",
	// 	}).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while creating a new session using CreateSession: %s\n", err)
	// 	return
	// }

	// // create manager context
	// managerAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
	// 	"Authorization": {
	// 		Prefix: "Bearer",
	// 		Key:    session.GetToken(),
	// 	},
	// })

	// // Get Application
	// application, _, err := managementClient.ManagementApi.
	// 	GetApplication(managerAuthContext, 1).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling GetApplication: %s\n", err)
	// 	return
	// }

	// fmt.Println(">> Management API call `GetApplication` response:")
	// pprint, _ := json.MarshalIndent(application, "", "  ")
	// fmt.Println(string(pprint))
	// fmt.Println("")

	// // Export loyalty balance
	// export, res, err := managementClient.ManagementApi.
	// 	ExportLoyaltyBalance(managerAuthContext, "1").
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling ExportLoyaltyBalance: %s\n", err)
	// 	printErrorWitheResponse(err, res)
	// 	return
	// }

	// exportCSVReader := csv.NewReader(strings.NewReader(export))
	// records, err := exportCSVReader.ReadAll()
	// if err != nil {
	// 	fmt.Printf("ERROR while reading export CSV: %s\n", err)
	// 	return
	// }

	// fmt.Println(">> Management API call `ExportLoyaltyBalance` records:")
	// fmt.Printf("%+v", records)
	// fmt.Println("")

	// // Import referrals
	// referralsFile, err := ioutil.ReadFile("/Users/altjake/Downloads/import-referrals-example.csv")
	// if err != nil {
	// 	fmt.Printf("ERROR while reading referrals import file: %s\n", err)
	// 	return
	// }

	// importObject, res, err := managementClient.ManagementApi.
	// 	ImportReferrals(managerAuthContext, 1, 181).
	// 	UpFile(string(referralsFile)).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling ImportReferrals: %s\n", err)
	// 	printErrorWitheResponse(err, res)
	// 	return
	// }

	// fmt.Println(">> Management API call `ImportReferrals`:")
	// fmt.Println(importObject)
	// fmt.Println(res)
	// fmt.Println("")

	// // Get Rulesets
	// rulesets, _, err := managementClient.ManagementApi.
	// 	GetRulesets(managerAuthContext, 1, 1).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling GetRulesets: %s\n", err)
	// 	return
	// }

	// fmt.Println(">> Management API call `GetRulesets` response:")
	// pprint, _ := json.MarshalIndent(rulesets, "", "  ")
	// fmt.Println(string(pprint))
	// fmt.Println("")

	// // Add Loyalty Points
	// programID := "1"
	// profileIntegrationID := "Cool_Dude"

	// body := talon.LoyaltyPoints{}
	// body.SetPoints(42.42)
	// body.SetName("Points that expire after 3 days")
	// body.SetValidityDuration("72h")

	// _, err = managementClient.ManagementApi.
	// 	AddLoyaltyPoints(managerAuthContext, programID, profileIntegrationID).
	// 	Body(body).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling AddLoyaltyPoints: %s\n", err)
	// 	return
	// }

	/**********************************************************************
	*************************** Integration API ***************************
	**********************************************************************/

	// // f10e9ee8463785b1aa0f40fa64bfed336253bddf2f3b55d76cb65055e638fdc9

	integrationClient := talon.NewAPIClient(configuration)
	integrationAuth := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"Authorization": {
			Prefix: "ApiKey-v1",
			Key:    "f10e9ee8463785b1aa0f40fa64bfed336253bddf2f3b55d76cb65055e638fdc9",
		},
	})

	// // Customer Session
	// // Instantiating a NewCustomerSession struct
	// customerSession := talon.NewCustomerSession{
	// 	// You can use both struct literals
	// 	ProfileId: talon.PtrString("DEADBEEF"),
	// 	State:     talon.PtrString("open"),
	// }

	// // Or alternatively, using the relevant setter in a later stage in the code
	// customerSession.SetTotal(42.0)

	// integrationState, _, err := integrationClient.IntegrationApi.
	// 	UpdateCustomerSession(integrationAuth, "deetdoot").
	// 	Body(customerSession).
	// 	Dry(false).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling UpdateCustomerSession: %s\n", err)
	// 	return
	// }
	// fmt.Println(">> Integration API call `UpdateCustomerSession` response:")
	// pprint, _ := json.MarshalIndent(integrationState, "", "  ")
	// fmt.Println(string(pprint))
	// fmt.Println("")

	////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////

	// // Customer Profile
	// newCustomerProfile := talon.NewCustomerProfile{}
	// newCustomerProfile.SetAttributes(map[string]interface{}{
	// 	"Name":  "HoopAlA",
	// 	"Email": "WOHOPOOL@talon.one",
	// })

	// updateCustomerProfileRequest := integrationClient.IntegrationApi.
	// 	UpdateCustomerProfile(integrationAuth, "DEADBEEF").
	// 	Body(newCustomerProfile)

	// integrationState, res, err := updateCustomerProfileRequest.Execute()
	// if err != nil {
	// 	printErrorWitheResponse(err, res)
	// 	return
	// }

	// fmt.Println(">> Integration API call `UpdateCustomerProfile` response:")
	// pprint, _ := json.MarshalIndent(integrationState, "", "  ")
	// fmt.Println(string(pprint))
	// fmt.Println("")

	////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////

	// // Customer Profile v2
	// profileV2Request := talon.CustomerProfileIntegrationRequestV2{}
	// profileV2Request.SetAttributes(map[string]interface{}{
	// 	"Name":  "HoopAlA25",
	// 	"Email": "WOHOPOOL5@talon.one",
	// })
	// profileV2Request.SetResponseContent([]string{
	// 	"customerProfile",
	// 	"loyalty",
	// })

	// updateCustomerProfileRequest := integrationClient.IntegrationApi.
	// 	UpdateCustomerProfileV2(integrationAuth, "DEADBEEF5").
	// 	Body(profileV2Request).
	// 	RunRuleEngine(true).
	// 	Dry(true)

	// integrationState, res, err := updateCustomerProfileRequest.Execute()
	// if err != nil {
	// 	printErrorWitheResponse(err, res)
	// 	return
	// }

	// fmt.Println(">> Integration API call `UpdateCustomerProfile` response:")
	// pprint, _ := json.MarshalIndent(integrationState, "", "  ")
	// fmt.Println(string(pprint))
	// fmt.Println("")

	////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////

	// Customer Inventory
	getCustomerInventoryRequest := integrationClient.IntegrationApi.
		GetCustomerInventory(integrationAuth, "my_unique_profile_id").
		Profile(true).
		Coupons(true)

	customerInventory, res, err := getCustomerInventoryRequest.Execute()
	if err != nil {
		printErrorWitheResponse(err, res)
		return
	}

	for i, coupon := range customerInventory.GetCoupons() {
		fmt.Println(i, coupon)
	}

	fmt.Println(">> Integration API call `GetCustomerInventory` response:")
	pprint, _ := json.MarshalIndent(customerInventory, "", "  ")
	fmt.Println(string(pprint))
	fmt.Println("")

	////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////

	// // Customer Session V2
	// newCustomerSession := talon.NewCustomerSessionV2{
	// 	ProfileId:   talon.PtrString("DEADBEEF"),
	// 	CouponCodes: &[]string{"Cool-Stuff!"},
	// }
	// newCustomerSession.SetCartItems([]talon.CartItem{
	// 	talon.CartItem{
	// 		Name:     "Pad Thai - Veggie",
	// 		Sku:      "pad-332",
	// 		Quantity: 1,
	// 		Price:    5.5,
	// 		Category: talon.PtrString("Noodles"),
	// 	},
	// 	talon.CartItem{
	// 		Name:     "Chang",
	// 		Sku:      "chang-br-42",
	// 		Quantity: 1,
	// 		Price:    2.3,
	// 		Category: talon.PtrString("Beverages"),
	// 	},
	// })

	// integrationRequest := talon.IntegrationRequest{
	// 	CustomerSession: newCustomerSession,
	// }
	// integrationRequest.SetResponseContent([]string{
	// 	// "customerSession",
	// 	// "customerProfile",
	// 	// "loyalty",
	// 	"ruleFailureReasons",
	// })

	// // Create/update a customer session using `UpdateCustomerSessionV2` function
	// integrationState, _, err := integrationClient.IntegrationApi.
	// 	UpdateCustomerSessionV2(integrationAuth, "deetdoot_2").
	// 	Body(integrationRequest).
	// 	Execute()

	// if err != nil {
	// 	fmt.Printf("ERROR while calling UpdateCustomerSessionV2: %s\n", err)
	// 	return
	// }
	// fmt.Println(">> Integration API call `UpdateCustomerSessionV2` response:")
	// pprint, _ := json.MarshalIndent(integrationState, "", "  ")
	// fmt.Printf("%s\n\n", string(pprint))

	// // Parsing the returned effects list, please consult https://developers.talon.one/Integration-API/handling-effects-v2 for the full list of effects and their corresponding properties
	// for _, effect := range integrationState.GetEffects() {
	// 	effectType := effect.GetEffectType()
	// 	switch {
	// 	case "setDiscount" == effectType:
	// 		// Initiating right props instance according to the effect type
	// 		effectProps := talon.SetDiscountEffectProps{}
	// 		if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
	// 			fmt.Printf("Error while decoding 'setDiscount' props: %s\n", err)
	// 			continue
	// 		}

	// 		// Access the specific effect's properties
	// 		fmt.Printf("Set a discount '%s' of %2.3f\n", effectProps.GetName(), effectProps.GetValue())
	// 	case "acceptCoupon" == effectType:
	// 		// Initiating right props instance according to the effect type
	// 		effectProps := talon.AcceptCouponEffectProps{}
	// 		if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
	// 			fmt.Printf("Error while decoding props: %s\n", err)
	// 			continue
	// 		}

	// 		// Work with AcceptCouponEffectProps' properties
	// 		// ...
	// 	case "rejectCoupon" == effectType:
	// 		// Initiating right props instance according to the effect type
	// 		effectProps := talon.RejectCouponEffectProps{}
	// 		if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
	// 			fmt.Printf("Error while decoding props: %s\n", err)
	// 			continue
	// 		}

	// 		fmt.Printf(">>>>>: %+v\n", effectProps)

	// 		// Work with RejectCouponEffectProps' properties
	// 		// ...
	// 	default:
	// 		fmt.Printf("Encounter unknown effect type: %s\n", effectType)
	// 	}
	// }
}

// quick decoding of props-map into our library structures using JSON marshaling, could also be done using a library like https://github.com/mitchellh/mapstructure
func decodeHelper(propsMap map[string]interface{}, v interface{}) error {
	propsJSON, err := json.Marshal(propsMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(propsJSON, v)
}

func printErrorWitheResponse(err error, res *http.Response) {
	fmt.Println(err)
	defer res.Body.Close()
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		// fmt.Println("Body issues...")
		// fmt.Println(e)
		return
	}

	fmt.Printf(string(body))
}
