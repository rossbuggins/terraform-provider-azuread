package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnAuthenticationMethodLoadStartHandler interface {
	OnAuthenticationMethodLoadStartHandler() BaseOnAuthenticationMethodLoadStartHandlerImpl
}

var _ OnAuthenticationMethodLoadStartHandler = BaseOnAuthenticationMethodLoadStartHandlerImpl{}

type BaseOnAuthenticationMethodLoadStartHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnAuthenticationMethodLoadStartHandlerImpl) OnAuthenticationMethodLoadStartHandler() BaseOnAuthenticationMethodLoadStartHandlerImpl {
	return s
}

var _ OnAuthenticationMethodLoadStartHandler = RawOnAuthenticationMethodLoadStartHandlerImpl{}

// RawOnAuthenticationMethodLoadStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnAuthenticationMethodLoadStartHandlerImpl struct {
	onAuthenticationMethodLoadStartHandler BaseOnAuthenticationMethodLoadStartHandlerImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawOnAuthenticationMethodLoadStartHandlerImpl) OnAuthenticationMethodLoadStartHandler() BaseOnAuthenticationMethodLoadStartHandlerImpl {
	return s.onAuthenticationMethodLoadStartHandler
}

func UnmarshalOnAuthenticationMethodLoadStartHandlerImplementation(input []byte) (OnAuthenticationMethodLoadStartHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAuthenticationMethodLoadStartHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp") {
		var out OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnAuthenticationMethodLoadStartHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnAuthenticationMethodLoadStartHandlerImpl: %+v", err)
	}

	return RawOnAuthenticationMethodLoadStartHandlerImpl{
		onAuthenticationMethodLoadStartHandler: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}