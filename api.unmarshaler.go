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

// Note: I have no idea if it's implemented in a good way or not.
func unmarshalInputMessageContent(rawBytes json.RawMessage) (data InputMessageContent, err error) {
	var temp struct {
		MessageText *string  `json:"message_text,omitempty"`
		Latitude    *float64 `json:"latitude,omitempty"`
		Address     *string  `json:"address,omitempty"`
		PhoneNumber *string  `json:"phone_number,omitempty"`
		Description *string  `json:"description,omitempty"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch {
	case temp.MessageText != nil:
		data = &InputTextMessageContent{}
	case temp.Address != nil:
		data = &InputVenueMessageContent{}
	case temp.Latitude != nil:
		data = &InputLocationMessageContent{}
	case temp.PhoneNumber != nil:
		data = &InputContactMessageContent{}
	case temp.Description != nil:
		data = &InputInvoiceMessageContent{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}
