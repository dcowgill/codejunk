package partial_json

import (
	"encoding/json"
	"log"
	"testing"
)

type skipDecode struct{}

func (v *skipDecode) UnmarshalJSON(data []byte) error {
	return nil
}

func fullDecode(data []byte) interface{} {
	var result map[string]interface{}
	must(json.Unmarshal(data, &result))
	return result
}

func partialDecode(data []byte) interface{} {
	var result map[string]skipDecode
	must(json.Unmarshal(data, &result))
	return result
}

var (
	input1 = []byte(`{"Param1": [1, "Hello", 42], "Param2": {"a": ["x", "y"], "b": [2, 4, 6, 8], "c": null}}`)
	input2 = []byte(`{"AddressID": "araqbcdgv9b0t", "BusinessID": null, "BusinessAddressID": null, "HubPickupAddressID": null, "PostalCode": "USA-10010", "Street": "22 West 21st Street", "Apt": "3rd fl", "City": "New York", "RegionCode": "NY", "Geocoded": { "PostalCode": "USA-10010", "Street": "22 W 21st St", "GPS": { "Lat": 0.7110612744294311, "Lon": -1.2914050673331954 }, "Notes": "", "DeliveryProtocols": null, "RequiresFreightElevator": false, "RequiresGovernmentID": false, "RequiresSecurityCheckIn": false }, "Instructions": "Drop on 3rd floor", "EcoFriendly": false, "Meals": [ "lunch" ], "HubID": "568a92b2b4c5066a690000d9", "HubName": "119 W 31st St", "HubExpires": "2016-08-24T22:00:00-04:00", "Created": "2015-09-02T13:42:55.198Z", "Updated": "2016-08-24T23:30:35.833798222Z", "Ordered": "2016-08-24T23:30:35.833798222Z", "CanDeliver": true, "Default": true}`)
)

func BenchmarkFullDecode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fullDecode(input1)
	}
}

func BenchmarkPartialDecode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partialDecode(input1)
	}
}

func BenchmarkFullDecode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fullDecode(input2)
	}
}

func BenchmarkPartialDecode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partialDecode(input2)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
