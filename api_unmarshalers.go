package tgo

import (
	"bytes"
	"encoding/json"
	"errors"
)

func unmarshalMaybeInaccessibleMessage(rawBytes json.RawMessage) (data MaybeInaccessibleMessage, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Date int64 `json:"date"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Date {
	case 0:
		data = &InaccessibleMessage{}
	default:
		data = &Message{}
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalMessageOrigin(rawBytes json.RawMessage) (data MessageOrigin, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "user":
		data = &MessageOriginUser{}
	case "hidden_user":
		data = &MessageOriginHiddenUser{}
	case "chat":
		data = &MessageOriginChat{}
	case "channel":
		data = &MessageOriginChannel{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalBackgroundFill(rawBytes json.RawMessage) (data BackgroundFill, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "solid":
		data = &BackgroundFillSolid{}
	case "gradient":
		data = &BackgroundFillGradient{}
	case "freeform_gradient":
		data = &BackgroundFillFreeformGradient{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalBackgroundType(rawBytes json.RawMessage) (data BackgroundType, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "fill":
		data = &BackgroundTypeFill{}
	case "wallpaper":
		data = &BackgroundTypeWallpaper{}
	case "pattern":
		data = &BackgroundTypePattern{}
	case "chat_theme":
		data = &BackgroundTypeChatTheme{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalChatMember(rawBytes json.RawMessage) (data ChatMember, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

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

func unmarshalReactionType(rawBytes json.RawMessage) (data ReactionType, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "emoji":
		data = &ReactionTypeEmoji{}
	case "custom_emoji":
		data = &ReactionTypeCustomEmoji{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalMenuButton(rawBytes json.RawMessage) (data MenuButton, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

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

func unmarshalChatBoostSource(rawBytes json.RawMessage) (data ChatBoostSource, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Source string `json:"source"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Source {
	case "premium":
		data = &ChatBoostSourcePremium{}
	case "gift_code":
		data = &ChatBoostSourceGiftCode{}
	case "giveaway":
		data = &ChatBoostSourceGiveaway{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

// Note: I have no idea if it's implemented in a good way or not.
func unmarshalInputMessageContent(rawBytes json.RawMessage) (data InputMessageContent, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

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

func unmarshalRevenueWithdrawalState(rawBytes json.RawMessage) (data RevenueWithdrawalState, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "pending":
		data = &RevenueWithdrawalStatePending{}
	case "succeeded":
		data = &RevenueWithdrawalStateSucceeded{}
	case "failed":
		data = &RevenueWithdrawalStateFailed{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalTransactionPartner(rawBytes json.RawMessage) (data TransactionPartner, err error) {
	if len(rawBytes) == 0 {
		return nil, nil
	}

	var temp struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(rawBytes, &temp); err != nil {
		return nil, err
	}

	switch temp.Type {
	case "fragment":
		data = &TransactionPartnerFragment{}
	case "user":
		data = &TransactionPartnerUser{}
	case "other":
		data = &TransactionPartnerOther{}
	default:
		return nil, errors.New("unknown type")
	}

	err = json.Unmarshal(rawBytes, data)
	return data, err
}

func unmarshalChatID(data json.RawMessage) (ChatID, error) {
	if bytes.HasPrefix(data, []byte("\"")) {
		var username Username
		err := json.Unmarshal(data, &username)
		return username, err
	} else {
		var id ID
		err := json.Unmarshal(data, &id)
		return id, err
	}
}
