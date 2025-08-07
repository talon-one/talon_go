package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	talon "github.com/talon-one/talon_go/v25"
)

func main() {
	configuration := talon.NewConfiguration()
	// Set API base path
	configuration.Servers = talon.ServerConfigurations{
		{
			// Notice that there is no trailing '/'
			URL:         "http://localhost:9000",
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
			Key:    os.Getenv("TALON_API_KEY"),
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
