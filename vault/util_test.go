package vault

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func testCheckResourceAttrJSON(name, key, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %q", name)
		}
		instanceState := resourceState.Primary
		if instanceState == nil {
			return fmt.Errorf("%q has no primary instance state", name)
		}
		v, ok := instanceState.Attributes[key]
		if !ok {
			return fmt.Errorf("%s: attribute not found %q", name, key)
		}
		var stateJSON, valueJSON interface{}
		err := json.Unmarshal([]byte(v), &stateJSON)
		if err != nil {
			return fmt.Errorf("%s: attribute %q not JSON: %s", name, key, err)
		}
		err = json.Unmarshal([]byte(value), &valueJSON)
		if err != nil {
			return fmt.Errorf("expected value %q not JSON: %s", value, err)
		}
		if !reflect.DeepEqual(stateJSON, valueJSON) {
			return fmt.Errorf("%s: attribute %q expected %#v, got %#v", name, key, stateJSON, valueJSON)
		}
		return nil

	}
}
