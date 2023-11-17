package tgo

import (
	"encoding/json"
	"errors"
)

func unmarshalChatMember(rawBytes json.RawMessage) (data ChatMember, err error) {
	var temp struct {
		Status string `json:"status"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Status {
	case "creator":
		data = &ChatMemberOwner{}
	case "administrator":
		data = &ChatMemberAdministrator{}
	case "member":
		data = &ChatMemberMember{}
	case "restricted":
		data = &ChatMemberRestricted{}
	case "left":
		data = &ChatMemberLeft{}
	case "kicked":
		data = &ChatMemberBanned{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalMenuButton(rawBytes json.RawMessage) (data MenuButton, err error) {
	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "commands":
		data = &MenuButtonCommands{}
	case "web_app":
		data = &MenuButtonWebApp{}
	case "default":
		data = &MenuButtonDefault{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

// ToDo: how the hell should I find out what type it is ???????
func unmarshalInputMessageContent(rawBytes json.RawMessage) (data InputMessageContent, err error) {
	// data = &InputTextMessageContent{}
	// data = &InputLocationMessageContent{}
	// data = &InputVenueMessageContent{}
	// data = &InputContactMessageContent{}
	// data = &InputInvoiceMessageContent{}
	return nil, errors.New("work in progress. sorry")
}
