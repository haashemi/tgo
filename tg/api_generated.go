package tg

import (
	"encoding/json"
	"strconv"
)

// Update represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId                int64                        `json:"update_id"`                           // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message                 *Message                     `json:"message,omitempty"`                   // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage           *Message                     `json:"edited_message,omitempty"`            // Optional. New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	ChannelPost             *Message                     `json:"channel_post,omitempty"`              // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`       // Optional. New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`       // Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessMessage         *Message                     `json:"business_message,omitempty"`          // Optional. New message from a connected business account
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`   // Optional. New version of a message from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"` // Optional. Messages were deleted from a connected business account
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`          // Optional. A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`    // Optional. Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`              // Optional. New incoming inline query
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`      // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`            // Optional. New incoming callback query
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`            // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`        // Optional. New incoming pre-checkout query. Contains full information about checkout
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`      // Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	Poll                    *Poll                        `json:"poll,omitempty"`                      // Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`               // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`            // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`               // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`         // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`                // Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`        // Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
}

// getUpdates is used to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
type GetUpdates struct {
	Offset         int64    `json:"offset,omitempty"`          // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
	Limit          int64    `json:"limit,omitempty"`           // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int64    `json:"timeout,omitempty"`         // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

// getUpdates is used to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (api *API) GetUpdates(payload *GetUpdates) ([]*Update, error) {
	return callJson[[]*Update](api, "getUpdates", payload)
}

// setWebhook is used to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
//
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.
// If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.
type SetWebhook struct {
	Url                string     `json:"url"`                            // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        *InputFile `json:"certificate,omitempty"`          // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string     `json:"ip_address,omitempty"`           // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int64      `json:"max_connections,omitempty"`      // The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string   `json:"allowed_updates,omitempty"`      // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool       `json:"drop_pending_updates,omitempty"` // Pass True to drop all pending updates
	SecretToken        string     `json:"secret_token,omitempty"`         // A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

func (x *SetWebhook) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Certificate != nil {
		if x.Certificate.IsUploadable() {
			media["certificate"] = x.Certificate
		}
	}

	return media
}

func (x *SetWebhook) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["url"] = x.Url
	if x.IpAddress != "" {
		payload["ip_address"] = x.IpAddress
	}
	if x.MaxConnections != 0 {
		payload["max_connections"] = strconv.FormatInt(x.MaxConnections, 10)
	}
	if x.AllowedUpdates != nil {
		if bb, err := json.Marshal(x.AllowedUpdates); err != nil {
			return nil, err
		} else {
			payload["allowed_updates"] = string(bb)
		}
	}
	if x.DropPendingUpdates {
		payload["drop_pending_updates"] = strconv.FormatBool(x.DropPendingUpdates)
	}
	if x.SecretToken != "" {
		payload["secret_token"] = x.SecretToken
	}

	return payload, nil
}

// setWebhook is used to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
//
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.
// If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.
func (api *API) SetWebhook(payload *SetWebhook) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "setWebhook", params, files)
	}
	return callJson[bool](api, "setWebhook", payload)
}

// deleteWebhook is used to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
type DeleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"` // Pass True to drop all pending updates
}

// deleteWebhook is used to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (api *API) DeleteWebhook(payload *DeleteWebhook) (bool, error) {
	return callJson[bool](api, "deleteWebhook", payload)
}

// getWebhookInfo is used to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (api *API) GetWebhookInfo() (*WebhookInfo, error) {
	return callJson[*WebhookInfo](api, "getWebhookInfo", nil)
}

// Describes the current status of a webhook.
type WebhookInfo struct {
	Url                          string   `json:"url"`                                       // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate         bool     `json:"has_custom_certificate"`                    // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount           int64    `json:"pending_update_count"`                      // Number of updates awaiting delivery
	IpAddress                    string   `json:"ip_address,omitempty"`                      // Optional. Currently used webhook IP address
	LastErrorDate                int64    `json:"last_error_date,omitempty"`                 // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage             string   `json:"last_error_message,omitempty"`              // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"` // Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
	MaxConnections               int64    `json:"max_connections,omitempty"`                 // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`                 // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
}

// User represents a Telegram user or bot.
type User struct {
	Id                      int64  `json:"id"`                                    // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	IsBot                   bool   `json:"is_bot"`                                // True, if this user is a bot
	FirstName               string `json:"first_name"`                            // User's or bot's first name
	LastName                string `json:"last_name,omitempty"`                   // Optional. User's or bot's last name
	Username                string `json:"username,omitempty"`                    // Optional. User's or bot's username
	LanguageCode            string `json:"language_code,omitempty"`               // Optional. IETF language tag of the user's language
	IsPremium               bool   `json:"is_premium,omitempty"`                  // Optional. True, if this user is a Telegram Premium user
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`    // Optional. True, if this user added the bot to the attachment menu
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`     // Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`            // Optional. True, if the bot has a main Web App. Returned only in getMe.
}

// Chat represents a chat.
type Chat struct {
	Id        int64  `json:"id"`                   // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type      string `json:"type"`                 // Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Title     string `json:"title,omitempty"`      // Optional. Title, for supergroups, channels and group chats
	Username  string `json:"username,omitempty"`   // Optional. Username, for private chats, supergroups and channels if available
	FirstName string `json:"first_name,omitempty"` // Optional. First name of the other party in a private chat
	LastName  string `json:"last_name,omitempty"`  // Optional. Last name of the other party in a private chat
	IsForum   bool   `json:"is_forum,omitempty"`   // Optional. True, if the supergroup chat is a forum (has topics enabled)
}

// ChatFullInfo contains full information about a chat.
type ChatFullInfo struct {
	Id                                 int64                 `json:"id"`                                                // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                               string                `json:"type"`                                              // Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                              string                `json:"title,omitempty"`                                   // Optional. Title, for supergroups, channels and group chats
	Username                           string                `json:"username,omitempty"`                                // Optional. Username, for private chats, supergroups and channels if available
	FirstName                          string                `json:"first_name,omitempty"`                              // Optional. First name of the other party in a private chat
	LastName                           string                `json:"last_name,omitempty"`                               // Optional. Last name of the other party in a private chat
	IsForum                            bool                  `json:"is_forum,omitempty"`                                // Optional. True, if the supergroup chat is a forum (has topics enabled)
	AccentColorId                      int64                 `json:"accent_color_id"`                                   // Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See accent colors for more details.
	MaxReactionCount                   int64                 `json:"max_reaction_count"`                                // The maximum number of reactions that can be set on a message in the chat
	Photo                              *ChatPhoto            `json:"photo,omitempty"`                                   // Optional. Chat photo
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`                        // Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`                               // Optional. For private chats, the date of birth of the user
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`                          // Optional. For private chats with business accounts, the intro of the business
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`                       // Optional. For private chats with business accounts, the location of the business
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`                  // Optional. For private chats with business accounts, the opening hours of the business
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`                           // Optional. For private chats, the personal channel of the user
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`                     // Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed.
	BackgroundCustomEmojiId            string                `json:"background_custom_emoji_id,omitempty"`              // Optional. Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background
	ProfileAccentColorId               int64                 `json:"profile_accent_color_id,omitempty"`                 // Optional. Identifier of the accent color for the chat's profile background. See profile accent colors for more details.
	ProfileBackgroundCustomEmojiId     string                `json:"profile_background_custom_emoji_id,omitempty"`      // Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background
	EmojiStatusCustomEmojiId           string                `json:"emoji_status_custom_emoji_id,omitempty"`            // Optional. Custom emoji identifier of the emoji status of the chat or the other party in a private chat
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`            // Optional. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any
	Bio                                string                `json:"bio,omitempty"`                                     // Optional. Bio of the other party in a private chat
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`                    // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"` // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`                   // Optional. True, if users need to join the supergroup before they can send messages
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`                         // Optional. True, if all users directly joining the supergroup without using an invite link need to be approved by supergroup administrators
	Description                        string                `json:"description,omitempty"`                             // Optional. Description, for groups, supergroups and channel chats
	InviteLink                         string                `json:"invite_link,omitempty"`                             // Optional. Primary invite link, for groups, supergroups and channel chats
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`                          // Optional. The most recent pinned message (by sending date)
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`                             // Optional. Default chat member permissions, for groups and supergroups
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`                     // Optional. True, if paid media messages can be sent or forwarded to the channel chat. The field is available only for channel chats.
	SlowModeDelay                      int64                 `json:"slow_mode_delay,omitempty"`                         // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds
	UnrestrictBoostCount               int64                 `json:"unrestrict_boost_count,omitempty"`                  // Optional. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions
	MessageAutoDeleteTime              int64                 `json:"message_auto_delete_time,omitempty"`                // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`        // Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators.
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`                      // Optional. True, if non-administrators can only get the list of bots and administrators in the chat
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`                   // Optional. True, if messages from the chat can't be forwarded to other chats
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`                     // Optional. True, if new chat members will have access to old messages; available only to chat administrators
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`                        // Optional. For supergroups, name of the group sticker set
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`                     // Optional. True, if the bot can change the group sticker set
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`           // Optional. For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group.
	LinkedChatId                       int64                 `json:"linked_chat_id,omitempty"`                          // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Location                           *ChatLocation         `json:"location,omitempty"`                                // Optional. For supergroups, the location to which the supergroup is connected
}

// Message represents a message.
type Message struct {
	MessageId                     int64                          `json:"message_id"`                                  // Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageThreadId               int64                          `json:"message_thread_id,omitempty"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from,omitempty"`                              // Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderBoostCount              int64                          `json:"sender_boost_count,omitempty"`                // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`               // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	Date                          int64                          `json:"date"`                                        // Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	BusinessConnectionId          string                         `json:"business_connection_id,omitempty"`            // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	Chat                          Chat                           `json:"chat"`                                        // Chat the message belongs to
	ForwardOrigin                 MessageOrigin                  `json:"forward_origin,omitempty"`                    // Optional. Information about the original message for forwarded messages
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                    // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	Quote                         *TextQuote                     `json:"quote,omitempty"`                             // Optional. For replies that quote part of the original message, the quoted part of the message
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                    // Optional. For replies to a story, the original story
	ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot through which the message was sent
	EditDate                      int64                          `json:"edit_date,omitempty"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True, if the message can't be forwarded
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`                   // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text,omitempty"`                              // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      []*MessageEntity               `json:"entities,omitempty"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`              // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	EffectId                      string                         `json:"effect_id,omitempty"`                         // Optional. Unique identifier of the message effect added to the message
	Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document,omitempty"`                          // Optional. Message is a general file, information about the file
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`                        // Optional. Message contains paid media; information about the paid media
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story,omitempty"`                             // Optional. Message is a forwarded story
	Video                         *Video                         `json:"video,omitempty"`                             // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`          // Optional. True, if the caption must be shown above the message media
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`                 // Optional. True, if the message media is covered by a spoiler animation
	Contact                       *Contact                       `json:"contact,omitempty"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice,omitempty"`                              // Optional. Message is a dice with random value
	Game                          *Game                          `json:"game,omitempty"`                              // Optional. Message is a game, information about the game. More about games »
	Poll                          *Poll                          `json:"poll,omitempty"`                              // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue,omitempty"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location,omitempty"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`                    // Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`                // Optional. Service message: the group has been created
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatId               int64                          `json:"migrate_to_chat_id,omitempty"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId             int64                          `json:"migrate_from_chat_id,omitempty"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage                 MaybeInaccessibleMessage       `json:"pinned_message,omitempty"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`                  // Optional. Message is a service message about a refunded payment, information about the payment. More about payments »
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                      // Optional. Service message: users were shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                       // Optional. Service message: user boosted the chat
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`               // Optional. Service message: chat background set
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Service message: forum topic edited
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Service message: forum topic closed
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Service message: forum topic reopened
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. Service message: the 'General' forum topic unhidden
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                  // Optional. Service message: a scheduled giveaway was created
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                          // Optional. The message is a scheduled giveaway message
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                  // Optional. A giveaway with public winners was completed
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`                // Optional. Service message: a giveaway without public winners was completed
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Service message: video chat ended
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

func (Message) IsMaybeInaccessibleMessage() {}

func (x *Message) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		MessageId                     int64                          `json:"message_id"`                                  // Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
		MessageThreadId               int64                          `json:"message_thread_id,omitempty"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
		From                          *User                          `json:"from,omitempty"`                              // Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
		SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
		SenderBoostCount              int64                          `json:"sender_boost_count,omitempty"`                // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
		SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`               // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
		Date                          int64                          `json:"date"`                                        // Date the message was sent in Unix time. It is always a positive number, representing a valid date.
		BusinessConnectionId          string                         `json:"business_connection_id,omitempty"`            // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
		Chat                          Chat                           `json:"chat"`                                        // Chat the message belongs to
		ForwardOrigin                 json.RawMessage                `json:"forward_origin,omitempty"`                    // Optional. Information about the original message for forwarded messages
		IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True, if the message is sent to a forum topic
		IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
		ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
		ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                    // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
		Quote                         *TextQuote                     `json:"quote,omitempty"`                             // Optional. For replies that quote part of the original message, the quoted part of the message
		ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                    // Optional. For replies to a story, the original story
		ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot through which the message was sent
		EditDate                      int64                          `json:"edit_date,omitempty"`                         // Optional. Date the message was last edited in Unix time
		HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True, if the message can't be forwarded
		IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`                   // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
		MediaGroupId                  string                         `json:"media_group_id,omitempty"`                    // Optional. The unique identifier of a media message group this message belongs to
		AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
		Text                          string                         `json:"text,omitempty"`                              // Optional. For text messages, the actual UTF-8 text of the message
		Entities                      []*MessageEntity               `json:"entities,omitempty"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
		LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`              // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
		EffectId                      string                         `json:"effect_id,omitempty"`                         // Optional. Unique identifier of the message effect added to the message
		Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
		Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Message is an audio file, information about the file
		Document                      *Document                      `json:"document,omitempty"`                          // Optional. Message is a general file, information about the file
		PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`                        // Optional. Message contains paid media; information about the paid media
		Photo                         []*PhotoSize                   `json:"photo,omitempty"`                             // Optional. Message is a photo, available sizes of the photo
		Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Message is a sticker, information about the sticker
		Story                         *Story                         `json:"story,omitempty"`                             // Optional. Message is a forwarded story
		Video                         *Video                         `json:"video,omitempty"`                             // Optional. Message is a video, information about the video
		VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Message is a video note, information about the video message
		Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Message is a voice message, information about the file
		Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for the animation, audio, document, paid media, photo, video or voice
		CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
		ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`          // Optional. True, if the caption must be shown above the message media
		HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`                 // Optional. True, if the message media is covered by a spoiler animation
		Contact                       *Contact                       `json:"contact,omitempty"`                           // Optional. Message is a shared contact, information about the contact
		Dice                          *Dice                          `json:"dice,omitempty"`                              // Optional. Message is a dice with random value
		Game                          *Game                          `json:"game,omitempty"`                              // Optional. Message is a game, information about the game. More about games »
		Poll                          *Poll                          `json:"poll,omitempty"`                              // Optional. Message is a native poll, information about the poll
		Venue                         *Venue                         `json:"venue,omitempty"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
		Location                      *Location                      `json:"location,omitempty"`                          // Optional. Message is a shared location, information about the location
		NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
		LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
		NewChatTitle                  string                         `json:"new_chat_title,omitempty"`                    // Optional. A chat title was changed to this value
		NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`                    // Optional. A chat photo was change to this value
		DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`                 // Optional. Service message: the chat photo was deleted
		GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`                // Optional. Service message: the group has been created
		SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
		ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
		MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"` // Optional. Service message: auto-delete timer settings changed in the chat
		MigrateToChatId               int64                          `json:"migrate_to_chat_id,omitempty"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
		MigrateFromChatId             int64                          `json:"migrate_from_chat_id,omitempty"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
		PinnedMessage                 json.RawMessage                `json:"pinned_message,omitempty"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
		Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
		SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
		RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`                  // Optional. Message is a service message about a refunded payment, information about the payment. More about payments »
		UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                      // Optional. Service message: users were shared with the bot
		ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message: a chat was shared with the bot
		ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
		WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
		PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data
		ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
		BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                       // Optional. Service message: user boosted the chat
		ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`               // Optional. Service message: chat background set
		ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Service message: forum topic created
		ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Service message: forum topic edited
		ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Service message: forum topic closed
		ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Service message: forum topic reopened
		GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. Service message: the 'General' forum topic hidden
		GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. Service message: the 'General' forum topic unhidden
		GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                  // Optional. Service message: a scheduled giveaway was created
		Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                          // Optional. The message is a scheduled giveaway message
		GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                  // Optional. A giveaway with public winners was completed
		GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`                // Optional. Service message: a giveaway without public winners was completed
		VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Service message: video chat scheduled
		VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Service message: video chat started
		VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Service message: video chat ended
		VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. Service message: new participants invited to a video chat
		WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Service message: data sent by a Web App
		ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalMessageOrigin(raw.ForwardOrigin); err != nil {
		return err
	} else {
		x.ForwardOrigin = data
	}

	if data, err := unmarshalMaybeInaccessibleMessage(raw.PinnedMessage); err != nil {
		return err
	} else {
		x.PinnedMessage = data
	}
	x.MessageId = raw.MessageId
	x.MessageThreadId = raw.MessageThreadId
	x.From = raw.From
	x.SenderChat = raw.SenderChat
	x.SenderBoostCount = raw.SenderBoostCount
	x.SenderBusinessBot = raw.SenderBusinessBot
	x.Date = raw.Date
	x.BusinessConnectionId = raw.BusinessConnectionId
	x.Chat = raw.Chat

	x.IsTopicMessage = raw.IsTopicMessage
	x.IsAutomaticForward = raw.IsAutomaticForward
	x.ReplyToMessage = raw.ReplyToMessage
	x.ExternalReply = raw.ExternalReply
	x.Quote = raw.Quote
	x.ReplyToStory = raw.ReplyToStory
	x.ViaBot = raw.ViaBot
	x.EditDate = raw.EditDate
	x.HasProtectedContent = raw.HasProtectedContent
	x.IsFromOffline = raw.IsFromOffline
	x.MediaGroupId = raw.MediaGroupId
	x.AuthorSignature = raw.AuthorSignature
	x.Text = raw.Text
	x.Entities = raw.Entities
	x.LinkPreviewOptions = raw.LinkPreviewOptions
	x.EffectId = raw.EffectId
	x.Animation = raw.Animation
	x.Audio = raw.Audio
	x.Document = raw.Document
	x.PaidMedia = raw.PaidMedia
	x.Photo = raw.Photo
	x.Sticker = raw.Sticker
	x.Story = raw.Story
	x.Video = raw.Video
	x.VideoNote = raw.VideoNote
	x.Voice = raw.Voice
	x.Caption = raw.Caption
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.HasMediaSpoiler = raw.HasMediaSpoiler
	x.Contact = raw.Contact
	x.Dice = raw.Dice
	x.Game = raw.Game
	x.Poll = raw.Poll
	x.Venue = raw.Venue
	x.Location = raw.Location
	x.NewChatMembers = raw.NewChatMembers
	x.LeftChatMember = raw.LeftChatMember
	x.NewChatTitle = raw.NewChatTitle
	x.NewChatPhoto = raw.NewChatPhoto
	x.DeleteChatPhoto = raw.DeleteChatPhoto
	x.GroupChatCreated = raw.GroupChatCreated
	x.SupergroupChatCreated = raw.SupergroupChatCreated
	x.ChannelChatCreated = raw.ChannelChatCreated
	x.MessageAutoDeleteTimerChanged = raw.MessageAutoDeleteTimerChanged
	x.MigrateToChatId = raw.MigrateToChatId
	x.MigrateFromChatId = raw.MigrateFromChatId

	x.Invoice = raw.Invoice
	x.SuccessfulPayment = raw.SuccessfulPayment
	x.RefundedPayment = raw.RefundedPayment
	x.UsersShared = raw.UsersShared
	x.ChatShared = raw.ChatShared
	x.ConnectedWebsite = raw.ConnectedWebsite
	x.WriteAccessAllowed = raw.WriteAccessAllowed
	x.PassportData = raw.PassportData
	x.ProximityAlertTriggered = raw.ProximityAlertTriggered
	x.BoostAdded = raw.BoostAdded
	x.ChatBackgroundSet = raw.ChatBackgroundSet
	x.ForumTopicCreated = raw.ForumTopicCreated
	x.ForumTopicEdited = raw.ForumTopicEdited
	x.ForumTopicClosed = raw.ForumTopicClosed
	x.ForumTopicReopened = raw.ForumTopicReopened
	x.GeneralForumTopicHidden = raw.GeneralForumTopicHidden
	x.GeneralForumTopicUnhidden = raw.GeneralForumTopicUnhidden
	x.GiveawayCreated = raw.GiveawayCreated
	x.Giveaway = raw.Giveaway
	x.GiveawayWinners = raw.GiveawayWinners
	x.GiveawayCompleted = raw.GiveawayCompleted
	x.VideoChatScheduled = raw.VideoChatScheduled
	x.VideoChatStarted = raw.VideoChatStarted
	x.VideoChatEnded = raw.VideoChatEnded
	x.VideoChatParticipantsInvited = raw.VideoChatParticipantsInvited
	x.WebAppData = raw.WebAppData
	x.ReplyMarkup = raw.ReplyMarkup
	return nil
}

// MessageId represents a unique message identifier.
type MessageId struct {
	MessageId int64 `json:"message_id"` // Unique message identifier. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
}

// InaccessibleMessage describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`       // Chat the message belonged to
	MessageId int64 `json:"message_id"` // Unique message identifier inside the chat
	Date      int64 `json:"date"`       // Always 0. The field can be used to differentiate regular and inaccessible messages.
}

func (InaccessibleMessage) IsMaybeInaccessibleMessage() {}

// MaybeInaccessibleMessage describes a message that can be inaccessible to the bot. It can be one of
// Message, InaccessibleMessage
type MaybeInaccessibleMessage interface {
	// IsMaybeInaccessibleMessage does nothing and is only used to enforce type-safety
	IsMaybeInaccessibleMessage()
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type          string `json:"type"`                      // Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or #hashtag@chatusername), “cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “expandable_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Offset        int64  `json:"offset"`                    // Offset in UTF-16 code units to the start of the entity
	Length        int64  `json:"length"`                    // Length of the entity in UTF-16 code units
	Url           string `json:"url,omitempty"`             // Optional. For “text_link” only, URL that will be opened after user taps on the text
	User          *User  `json:"user,omitempty"`            // Optional. For “text_mention” only, the mentioned user
	Language      string `json:"language,omitempty"`        // Optional. For “pre” only, the programming language of the entity text
	CustomEmojiId string `json:"custom_emoji_id,omitempty"` // Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

// TextQuote contains information about the quoted part of a message that is replied to by the given message.
type TextQuote struct {
	Text     string           `json:"text"`                // Text of the quoted part of a message that is replied to by the given message
	Entities []*MessageEntity `json:"entities,omitempty"`  // Optional. Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Position int64            `json:"position"`            // Approximate quote position in the original message in UTF-16 code units as specified by the sender
	IsManual bool             `json:"is_manual,omitempty"` // Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
}

// ExternalReplyInfo contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`                         // Origin of the message replied to by the given message
	Chat               *Chat               `json:"chat,omitempty"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	MessageId          int64               `json:"message_id,omitempty"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Options used for link preview generation for the original message, if it is a text message
	Animation          *Animation          `json:"animation,omitempty"`            // Optional. Message is an animation, information about the animation
	Audio              *Audio              `json:"audio,omitempty"`                // Optional. Message is an audio file, information about the file
	Document           *Document           `json:"document,omitempty"`             // Optional. Message is a general file, information about the file
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Optional. Message contains paid media; information about the paid media
	Photo              []*PhotoSize        `json:"photo,omitempty"`                // Optional. Message is a photo, available sizes of the photo
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional. Message is a sticker, information about the sticker
	Story              *Story              `json:"story,omitempty"`                // Optional. Message is a forwarded story
	Video              *Video              `json:"video,omitempty"`                // Optional. Message is a video, information about the video
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional. Message is a video note, information about the video message
	Voice              *Voice              `json:"voice,omitempty"`                // Optional. Message is a voice message, information about the file
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // Optional. True, if the message media is covered by a spoiler animation
	Contact            *Contact            `json:"contact,omitempty"`              // Optional. Message is a shared contact, information about the contact
	Dice               *Dice               `json:"dice,omitempty"`                 // Optional. Message is a dice with random value
	Game               *Game               `json:"game,omitempty"`                 // Optional. Message is a game, information about the game. More about games »
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional. Message is a scheduled giveaway, information about the giveaway
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional. A giveaway with public winners was completed
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	Location           *Location           `json:"location,omitempty"`             // Optional. Message is a shared location, information about the location
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. Message is a native poll, information about the poll
	Venue              *Venue              `json:"venue,omitempty"`                // Optional. Message is a venue, information about the venue
}

func (x *ExternalReplyInfo) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Origin             json.RawMessage     `json:"origin"`                         // Origin of the message replied to by the given message
		Chat               *Chat               `json:"chat,omitempty"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
		MessageId          int64               `json:"message_id,omitempty"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
		LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Options used for link preview generation for the original message, if it is a text message
		Animation          *Animation          `json:"animation,omitempty"`            // Optional. Message is an animation, information about the animation
		Audio              *Audio              `json:"audio,omitempty"`                // Optional. Message is an audio file, information about the file
		Document           *Document           `json:"document,omitempty"`             // Optional. Message is a general file, information about the file
		PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Optional. Message contains paid media; information about the paid media
		Photo              []*PhotoSize        `json:"photo,omitempty"`                // Optional. Message is a photo, available sizes of the photo
		Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional. Message is a sticker, information about the sticker
		Story              *Story              `json:"story,omitempty"`                // Optional. Message is a forwarded story
		Video              *Video              `json:"video,omitempty"`                // Optional. Message is a video, information about the video
		VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional. Message is a video note, information about the video message
		Voice              *Voice              `json:"voice,omitempty"`                // Optional. Message is a voice message, information about the file
		HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // Optional. True, if the message media is covered by a spoiler animation
		Contact            *Contact            `json:"contact,omitempty"`              // Optional. Message is a shared contact, information about the contact
		Dice               *Dice               `json:"dice,omitempty"`                 // Optional. Message is a dice with random value
		Game               *Game               `json:"game,omitempty"`                 // Optional. Message is a game, information about the game. More about games »
		Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional. Message is a scheduled giveaway, information about the giveaway
		GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional. A giveaway with public winners was completed
		Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
		Location           *Location           `json:"location,omitempty"`             // Optional. Message is a shared location, information about the location
		Poll               *Poll               `json:"poll,omitempty"`                 // Optional. Message is a native poll, information about the poll
		Venue              *Venue              `json:"venue,omitempty"`                // Optional. Message is a venue, information about the venue
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalMessageOrigin(raw.Origin); err != nil {
		return err
	} else {
		x.Origin = data
	}

	x.Chat = raw.Chat
	x.MessageId = raw.MessageId
	x.LinkPreviewOptions = raw.LinkPreviewOptions
	x.Animation = raw.Animation
	x.Audio = raw.Audio
	x.Document = raw.Document
	x.PaidMedia = raw.PaidMedia
	x.Photo = raw.Photo
	x.Sticker = raw.Sticker
	x.Story = raw.Story
	x.Video = raw.Video
	x.VideoNote = raw.VideoNote
	x.Voice = raw.Voice
	x.HasMediaSpoiler = raw.HasMediaSpoiler
	x.Contact = raw.Contact
	x.Dice = raw.Dice
	x.Game = raw.Game
	x.Giveaway = raw.Giveaway
	x.GiveawayWinners = raw.GiveawayWinners
	x.Invoice = raw.Invoice
	x.Location = raw.Location
	x.Poll = raw.Poll
	x.Venue = raw.Venue
	return nil
}

// Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	MessageId                int64            `json:"message_id"`                            // Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	ChatId                   ChatID           `json:"chat_id,omitempty"`                     // Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername). Not supported for messages sent on behalf of a business account.
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified message to be replied to is not found. Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
	Quote                    string           `json:"quote,omitempty"`                       // Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
	QuoteParseMode           string           `json:"quote_parse_mode,omitempty"`            // Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteEntities            []*MessageEntity `json:"quote_entities,omitempty"`              // Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
	QuotePosition            int64            `json:"quote_position,omitempty"`              // Optional. Position of the quote in the original message in UTF-16 code units
}

func (x *ReplyParameters) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		MessageId                int64            `json:"message_id"`                            // Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
		ChatId                   json.RawMessage  `json:"chat_id,omitempty"`                     // Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername). Not supported for messages sent on behalf of a business account.
		AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified message to be replied to is not found. Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
		Quote                    string           `json:"quote,omitempty"`                       // Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
		QuoteParseMode           string           `json:"quote_parse_mode,omitempty"`            // Optional. Mode for parsing entities in the quote. See formatting options for more details.
		QuoteEntities            []*MessageEntity `json:"quote_entities,omitempty"`              // Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
		QuotePosition            int64            `json:"quote_position,omitempty"`              // Optional. Position of the quote in the original message in UTF-16 code units
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatID(raw.ChatId); err != nil {
		return err
	} else {
		x.ChatId = data
	}
	x.MessageId = raw.MessageId

	x.AllowSendingWithoutReply = raw.AllowSendingWithoutReply
	x.Quote = raw.Quote
	x.QuoteParseMode = raw.QuoteParseMode
	x.QuoteEntities = raw.QuoteEntities
	x.QuotePosition = raw.QuotePosition
	return nil
}

// MessageOrigin describes the origin of a message. It can be one of
// MessageOriginUser, MessageOriginHiddenUser, MessageOriginChat, MessageOriginChannel
type MessageOrigin interface {
	// IsMessageOrigin does nothing and is only used to enforce type-safety
	IsMessageOrigin()
}

// The message was originally sent by a known user.
type MessageOriginUser struct {
	Type       string `json:"type"`        // Type of the message origin, always “user”
	Date       int64  `json:"date"`        // Date the message was sent originally in Unix time
	SenderUser User   `json:"sender_user"` // User that sent the message originally
}

func (MessageOriginUser) IsMessageOrigin() {}

// The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always “hidden_user”
	Date           int64  `json:"date"`             // Date the message was sent originally in Unix time
	SenderUserName string `json:"sender_user_name"` // Name of the user that sent the message originally
}

func (MessageOriginHiddenUser) IsMessageOrigin() {}

// The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	Type            string `json:"type"`                       // Type of the message origin, always “chat”
	Date            int64  `json:"date"`                       // Date the message was sent originally in Unix time
	SenderChat      Chat   `json:"sender_chat"`                // Chat that sent the message originally
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. For messages originally sent by an anonymous chat administrator, original message author signature
}

func (MessageOriginChat) IsMessageOrigin() {}

// The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	Type            string `json:"type"`                       // Type of the message origin, always “channel”
	Date            int64  `json:"date"`                       // Date the message was sent originally in Unix time
	Chat            Chat   `json:"chat"`                       // Channel chat to which the message was originally sent
	MessageId       int64  `json:"message_id"`                 // Unique message identifier inside the chat
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. Signature of the original post author
}

func (MessageOriginChannel) IsMessageOrigin() {}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileId       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int64  `json:"width"`               // Photo width
	Height       int64  `json:"height"`              // Photo height
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int64      `json:"width"`               // Video width as defined by the sender
	Height       int64      `json:"height"`              // Video height as defined by the sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Animation thumbnail as defined by the sender
	FileName     string     `json:"file_name,omitempty"` // Optional. Original animation filename as defined by the sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64      `json:"duration"`            // Duration of the audio in seconds as defined by the sender
	Performer    string     `json:"performer,omitempty"` // Optional. Performer of the audio as defined by the sender or by audio tags
	Title        string     `json:"title,omitempty"`     // Optional. Title of the audio as defined by the sender or by audio tags
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the album cover to which the music file belongs
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Document thumbnail as defined by the sender
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Story represents a story.
type Story struct {
	Chat Chat  `json:"chat"` // Chat that posted the story
	Id   int64 `json:"id"`   // Unique identifier for the story in the chat
}

// Video represents a video file.
type Video struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int64      `json:"width"`               // Video width as defined by the sender
	Height       int64      `json:"height"`              // Video height as defined by the sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int64      `json:"length"`              // Video width and height (diameter of the video message) as defined by the sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes
}

// Voice represents a voice note.
type Voice struct {
	FileId       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64  `json:"duration"`            // Duration of the audio in seconds as defined by the sender
	MimeType     string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Describes the paid media added to a message.
type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"` // The number of Telegram Stars that must be paid to buy access to the media
	PaidMedia []PaidMedia `json:"paid_media"` // Information about the paid media
}

// PaidMedia describes paid media. Currently, it can be one of
// PaidMediaPreview, PaidMediaPhoto, PaidMediaVideo
type PaidMedia interface {
	// IsPaidMedia does nothing and is only used to enforce type-safety
	IsPaidMedia()
}

// The paid media isn't available before the payment.
type PaidMediaPreview struct {
	Type     string `json:"type"`               // Type of the paid media, always “preview”
	Width    int64  `json:"width,omitempty"`    // Optional. Media width as defined by the sender
	Height   int64  `json:"height,omitempty"`   // Optional. Media height as defined by the sender
	Duration int64  `json:"duration,omitempty"` // Optional. Duration of the media in seconds as defined by the sender
}

func (PaidMediaPreview) IsPaidMedia() {}

// The paid media is a photo.
type PaidMediaPhoto struct {
	Type  string       `json:"type"`  // Type of the paid media, always “photo”
	Photo []*PhotoSize `json:"photo"` // The photo
}

func (PaidMediaPhoto) IsPaidMedia() {}

// The paid media is a video.
type PaidMediaVideo struct {
	Type  string `json:"type"`  // Type of the paid media, always “video”
	Video Video  `json:"video"` // The video
}

func (PaidMediaVideo) IsPaidMedia() {}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	UserId      int64  `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Vcard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int64  `json:"value"` // Value of the dice, 1-6 for “”, “” and “” base emoji, 1-5 for “” and “” base emoji, 1-64 for “” base emoji
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	Text         string           `json:"text"`                    // Option text, 1-100 characters
	TextEntities []*MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount   int64            `json:"voter_count"`             // Number of users that voted for this option
}

// InputPollOption contains information about one answer option in a poll to be sent.
type InputPollOption struct {
	Text          string           `json:"text"`                      // Option text, 1-100 characters
	TextParseMode string           `json:"text_parse_mode,omitempty"` // Optional. Mode for parsing entities in the text. See formatting options for more details. Currently, only custom emoji entities are allowed
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`   // Optional. A JSON-serialized list of special entities that appear in the poll option text. It can be specified instead of text_parse_mode
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollId    string  `json:"poll_id"`              // Unique poll identifier
	VoterChat *Chat   `json:"voter_chat,omitempty"` // Optional. The chat that changed the answer to the poll, if the voter is anonymous
	User      *User   `json:"user,omitempty"`       // Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	OptionIds []int64 `json:"option_ids"`           // 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}

// Poll contains information about a poll.
type Poll struct {
	Id                    string           `json:"id"`                             // Unique poll identifier
	Question              string           `json:"question"`                       // Poll question, 1-300 characters
	QuestionEntities      []*MessageEntity `json:"question_entities,omitempty"`    // Optional. Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions
	Options               []*PollOption    `json:"options"`                        // List of poll options
	TotalVoterCount       int64            `json:"total_voter_count"`              // Total number of users that voted in the poll
	IsClosed              bool             `json:"is_closed"`                      // True, if the poll is closed
	IsAnonymous           bool             `json:"is_anonymous"`                   // True, if the poll is anonymous
	Type                  string           `json:"type"`                           // Poll type, currently can be “regular” or “quiz”
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`        // True, if the poll allows multiple answers
	CorrectOptionId       int64            `json:"correct_option_id,omitempty"`    // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string           `json:"explanation,omitempty"`          // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"` // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int64            `json:"open_period,omitempty"`          // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int64            `json:"close_date,omitempty"`           // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

// Location represents a point on the map.
type Location struct {
	Latitude             float64 `json:"latitude"`                         // Latitude as defined by the sender
	Longitude            float64 `json:"longitude"`                        // Longitude as defined by the sender
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64   `json:"live_period,omitempty"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              int64   `json:"heading,omitempty"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

// Venue represents a venue.
type Venue struct {
	Location        Location `json:"location"`                    // Venue location. Can't be a live location
	Title           string   `json:"title"`                       // Name of the venue
	Address         string   `json:"address"`                     // Address of the venue
	FoursquareId    string   `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string   `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId   string   `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string   `json:"google_place_type,omitempty"` // Optional. Google Places type of the venue. (See supported types.)
}

// Describes data sent from a Web App to the bot.
type WebAppData struct {
	Data       string `json:"data"`        // The data. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"` // Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
}

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler User  `json:"traveler"` // User that triggered the alert
	Watcher  User  `json:"watcher"`  // User that set the alert
	Distance int64 `json:"distance"` // The distance between the users
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}

// ChatBoostAdded represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int64 `json:"boost_count"` // Number of boosts added by the user
}

// BackgroundFill describes the way a background is filled based on the selected colors. Currently, it can be one of
// BackgroundFillSolid, BackgroundFillGradient, BackgroundFillFreeformGradient
type BackgroundFill interface {
	// IsBackgroundFill does nothing and is only used to enforce type-safety
	IsBackgroundFill()
}

// The background is filled using the selected color.
type BackgroundFillSolid struct {
	Type  string `json:"type"`  // Type of the background fill, always “solid”
	Color int64  `json:"color"` // The color of the background fill in the RGB24 format
}

func (BackgroundFillSolid) IsBackgroundFill() {}

// The background is a gradient fill.
type BackgroundFillGradient struct {
	Type          string `json:"type"`           // Type of the background fill, always “gradient”
	TopColor      int64  `json:"top_color"`      // Top color of the gradient in the RGB24 format
	BottomColor   int64  `json:"bottom_color"`   // Bottom color of the gradient in the RGB24 format
	RotationAngle int64  `json:"rotation_angle"` // Clockwise rotation angle of the background fill in degrees; 0-359
}

func (BackgroundFillGradient) IsBackgroundFill() {}

// The background is a freeform gradient that rotates after every message in the chat.
type BackgroundFillFreeformGradient struct {
	Type   string  `json:"type"`   // Type of the background fill, always “freeform_gradient”
	Colors []int64 `json:"colors"` // A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
}

func (BackgroundFillFreeformGradient) IsBackgroundFill() {}

// BackgroundType describes the type of a background. Currently, it can be one of
// BackgroundTypeFill, BackgroundTypeWallpaper, BackgroundTypePattern, BackgroundTypeChatTheme
type BackgroundType interface {
	// IsBackgroundType does nothing and is only used to enforce type-safety
	IsBackgroundType()
}

// The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	Type             string         `json:"type"`               // Type of the background, always “fill”
	Fill             BackgroundFill `json:"fill"`               // The background fill
	DarkThemeDimming int64          `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
}

func (BackgroundTypeFill) IsBackgroundType() {}

func (x *BackgroundTypeFill) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type             string          `json:"type"`               // Type of the background, always “fill”
		Fill             json.RawMessage `json:"fill"`               // The background fill
		DarkThemeDimming int64           `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalBackgroundFill(raw.Fill); err != nil {
		return err
	} else {
		x.Fill = data
	}
	x.Type = raw.Type

	x.DarkThemeDimming = raw.DarkThemeDimming
	return nil
}

// The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`                 // Type of the background, always “wallpaper”
	Document         Document `json:"document"`             // Document with the wallpaper
	DarkThemeDimming int64    `json:"dark_theme_dimming"`   // Dimming of the background in dark themes, as a percentage; 0-100
	IsBlurred        bool     `json:"is_blurred,omitempty"` // Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsMoving         bool     `json:"is_moving,omitempty"`  // Optional. True, if the background moves slightly when the device is tilted
}

func (BackgroundTypeWallpaper) IsBackgroundType() {}

// The background is a PNG or TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”) pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	Type       string         `json:"type"`                  // Type of the background, always “pattern”
	Document   Document       `json:"document"`              // Document with the pattern
	Fill       BackgroundFill `json:"fill"`                  // The background fill that is combined with the pattern
	Intensity  int64          `json:"intensity"`             // Intensity of the pattern when it is shown above the filled background; 0-100
	IsInverted bool           `json:"is_inverted,omitempty"` // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsMoving   bool           `json:"is_moving,omitempty"`   // Optional. True, if the background moves slightly when the device is tilted
}

func (BackgroundTypePattern) IsBackgroundType() {}

func (x *BackgroundTypePattern) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type       string          `json:"type"`                  // Type of the background, always “pattern”
		Document   Document        `json:"document"`              // Document with the pattern
		Fill       json.RawMessage `json:"fill"`                  // The background fill that is combined with the pattern
		Intensity  int64           `json:"intensity"`             // Intensity of the pattern when it is shown above the filled background; 0-100
		IsInverted bool            `json:"is_inverted,omitempty"` // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
		IsMoving   bool            `json:"is_moving,omitempty"`   // Optional. True, if the background moves slightly when the device is tilted
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalBackgroundFill(raw.Fill); err != nil {
		return err
	} else {
		x.Fill = data
	}
	x.Type = raw.Type
	x.Document = raw.Document

	x.Intensity = raw.Intensity
	x.IsInverted = raw.IsInverted
	x.IsMoving = raw.IsMoving
	return nil
}

// The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`       // Type of the background, always “chat_theme”
	ThemeName string `json:"theme_name"` // Name of the chat theme, which is usually an emoji
}

func (BackgroundTypeChatTheme) IsBackgroundType() {}

// ChatBackground represents a chat background.
type ChatBackground struct {
	Type BackgroundType `json:"type"` // Type of the background
}

func (x *ChatBackground) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type json.RawMessage `json:"type"` // Type of the background
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalBackgroundType(raw.Type); err != nil {
		return err
	} else {
		x.Type = data
	}

	return nil
}

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string `json:"name"`                           // Name of the topic
	IconColor         int64  `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat. Currently holds no information.
type ForumTopicClosed struct{}

// ForumTopicEdited represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`                 // Optional. New name of the topic, if it was edited
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat. Currently holds no information.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat. Currently holds no information.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat. Currently holds no information.
type GeneralForumTopicUnhidden struct{}

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
type SharedUser struct {
	UserId    int64        `json:"user_id"`              // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	FirstName string       `json:"first_name,omitempty"` // Optional. First name of the user, if the name was requested by the bot
	LastName  string       `json:"last_name,omitempty"`  // Optional. Last name of the user, if the name was requested by the bot
	Username  string       `json:"username,omitempty"`   // Optional. Username of the user, if the username was requested by the bot
	Photo     []*PhotoSize `json:"photo,omitempty"`      // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// UsersShared contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	RequestId int64         `json:"request_id"` // Identifier of the request
	Users     []*SharedUser `json:"users"`      // Information about users shared with the bot.
}

// ChatShared contains information about a chat that was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	RequestId int64        `json:"request_id"`         // Identifier of the request
	ChatId    int64        `json:"chat_id"`            // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	Title     string       `json:"title,omitempty"`    // Optional. Title of the chat, if the title was requested by the bot.
	Username  string       `json:"username,omitempty"` // Optional. Username of the chat, if the username was requested by the bot and available.
	Photo     []*PhotoSize `json:"photo,omitempty"`    // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// WriteAccessAllowed represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`         // Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	WebAppName         string `json:"web_app_name,omitempty"`         // Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"` // Optional. True, if the access was granted when the bot was added to the attachment or side menu
}

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

// VideoChatStarted represents a service message about a video chat started in the chat. Currently holds no information.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int64 `json:"duration"` // Video chat duration in seconds
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"` // New members that were invited to the video chat
}

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount int64 `json:"prize_star_count,omitempty"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
}

// Giveaway represents a message about a scheduled giveaway.
type Giveaway struct {
	Chats                         []*Chat  `json:"chats"`                                      // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int64    `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int64    `json:"winner_count"`                               // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`                 // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              string   `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PrizeStarCount                int64    `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int64    `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
}

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`                                       // The chat that created the giveaway
	GiveawayMessageId             int64   `json:"giveaway_message_id"`                        // Identifier of the message with the giveaway in the chat
	WinnersSelectionDate          int64   `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int64   `json:"winner_count"`                               // Total number of winners in the giveaway
	Winners                       []*User `json:"winners"`                                    // List of up to 100 winners of the giveaway
	AdditionalChatCount           int64   `json:"additional_chat_count,omitempty"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PrizeStarCount                int64   `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int64   `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	UnclaimedPrizeCount           int64   `json:"unclaimed_prize_count,omitempty"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                bool    `json:"only_new_members,omitempty"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   bool    `json:"was_refunded,omitempty"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              string  `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
}

// GiveawayCompleted represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	WinnerCount         int64    `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount int64    `json:"unclaimed_prize_count,omitempty"` // Optional. Number of undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`      // Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
}

// Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`        // Optional. True, if the link preview is disabled
	Url              string `json:"url,omitempty"`                // Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"` // Optional. True, if the media in the link preview is supposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"` // Optional. True, if the media in the link preview is supposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	ShowAboveText    bool   `json:"show_above_text,omitempty"`    // Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64          `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]*PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
//
// The maximum file size to download is 20 MB
type File struct {
	FileId       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FilePath     string `json:"file_path,omitempty"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// Describes a Web App.
type WebAppInfo struct {
	Url string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          bool                `json:"is_persistent,omitempty"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool                `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool                `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

func (ReplyKeyboardMarkup) IsReplyMarkup() {}

// KeyboardButton represents one button of the reply keyboard. At most one of the optional fields must be used to specify type of the button. For simple text buttons, String can be used instead of this object to specify the button text.
// Note: request_users and request_chat options will only work in Telegram versions released after 3 February, 2023. Older clients will display unsupported message.
type KeyboardButton struct {
	Text            string                      `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`    // Optional. If specified, pressing the button will open a list of suitable users. Identifiers of selected users will be sent to the bot in a “users_shared” service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestContact  bool                        `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation bool                        `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

// KeyboardButtonRequestUsers defines the criteria used to request suitable users. Information about the selected users will be shared with the bot when the corresponding button is pressed. More about requesting users »
type KeyboardButtonRequestUsers struct {
	RequestId       int64 `json:"request_id"`                 // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	UserIsBot       bool  `json:"user_is_bot,omitempty"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   bool  `json:"user_is_premium,omitempty"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     int64 `json:"max_quantity,omitempty"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool  `json:"request_name,omitempty"`     // Optional. Pass True to request the users' first and last names
	RequestUsername bool  `json:"request_username,omitempty"` // Optional. Pass True to request the users' usernames
	RequestPhoto    bool  `json:"request_photo,omitempty"`    // Optional. Pass True to request the users' photos
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable chat. Information about the selected chat will be shared with the bot when the corresponding button is pressed. The bot will be granted requested rights in the chat if appropriate. More about requesting chats ».
type KeyboardButtonRequestChat struct {
	RequestId               int64                    `json:"request_id"`                          // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel           bool                     `json:"chat_is_channel"`                     // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestTitle            bool                     `json:"request_title,omitempty"`             // Optional. Pass True to request the chat's title
	RequestUsername         bool                     `json:"request_username,omitempty"`          // Optional. Pass True to request the chat's username
	RequestPhoto            bool                     `json:"request_photo,omitempty"`             // Optional. Pass True to request the chat's photo
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

func (ReplyKeyboardRemove) IsReplyMarkup() {}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

func (InlineKeyboardMarkup) IsReplyMarkup() {}

// InlineKeyboardButton represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify type of the button.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	Url                          string                       `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	CallbackData                 string                       `json:"callback_data,omitempty"`                    // Optional. Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field. Not supported for messages sent on behalf of a Telegram Business account.
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`                        // Optional. Description of the button that copies the specified text to the clipboard.
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
	Pay                          bool                         `json:"pay,omitempty"`                              // Optional. Specify True, to send a Pay button. Substrings “” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
// Telegram apps support these buttons as of version 5.7.
//
// Sample bot: @discussbot
type LoginUrl struct {
	Url                string `json:"url"`                            // An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        string `json:"forward_text,omitempty"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username,omitempty"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access,omitempty"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`               // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`    // Optional. True, if private chats with users can be chosen
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"` // Optional. True, if channel chats can be chosen
}

// CopyTextButton represents an inline keyboard button that copies specified text to the clipboard.
type CopyTextButton struct {
	Text string `json:"text"` // The text to be copied to the clipboard; 1-256 characters
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
//
// NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	Id              string                   `json:"id"`                          // Unique identifier for this query
	From            User                     `json:"from"`                        // Sender
	Message         MaybeInaccessibleMessage `json:"message,omitempty"`           // Optional. Message sent by the bot with the callback button that originated the query
	InlineMessageId string                   `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string                   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string                   `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   string                   `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

func (x *CallbackQuery) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Id              string          `json:"id"`                          // Unique identifier for this query
		From            User            `json:"from"`                        // Sender
		Message         json.RawMessage `json:"message,omitempty"`           // Optional. Message sent by the bot with the callback button that originated the query
		InlineMessageId string          `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
		ChatInstance    string          `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
		Data            string          `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
		GameShortName   string          `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalMaybeInaccessibleMessage(raw.Message); err != nil {
		return err
	} else {
		x.Message = data
	}
	x.Id = raw.Id
	x.From = raw.From

	x.InlineMessageId = raw.InlineMessageId
	x.ChatInstance = raw.ChatInstance
	x.Data = raw.Data
	x.GameShortName = raw.GameShortName
	return nil
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode. Not supported in channels and for messages sent on behalf of a Telegram Business account.
//
// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions). There could be two ways to create a new poll:
//
// Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing for hardcore users but lacks modern day polish.
// Guide the user through a step-by-step process. 'Please send me your question', 'Cool, now let's add the first answer option', 'Great. Keep adding answer options, then send /done when you're ready'.
//
// The last option is definitely more attractive. And if you use ForceReply in your bot's questions, it will receive the user's answers even if it only receives replies, commands and mentions - without any extra work for the user.
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`                       // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool   `json:"selective,omitempty"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
}

func (ForceReply) IsReplyMarkup() {}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueId string `json:"small_file_unique_id"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileId         string `json:"big_file_id"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueId   string `json:"big_file_unique_id"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

// Represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                          // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	Creator                 User   `json:"creator"`                              // Creator of the link
	CreatesJoinRequest      bool   `json:"creates_join_request"`                 // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool   `json:"is_primary"`                           // True, if the link is primary
	IsRevoked               bool   `json:"is_revoked"`                           // True, if the link is revoked
	Name                    string `json:"name,omitempty"`                       // Optional. Invite link name
	ExpireDate              int64  `json:"expire_date,omitempty"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             int64  `json:"member_limit,omitempty"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount int64  `json:"pending_join_request_count,omitempty"` // Optional. Number of pending join requests created using this link
	SubscriptionPeriod      int64  `json:"subscription_period,omitempty"`        // Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPrice       int64  `json:"subscription_price,omitempty"`         // Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`                // True, if the user's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat"`             // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool `json:"can_delete_messages"`         // True, if the administrator can delete messages of other users
	CanManageVideoChats bool `json:"can_manage_video_chats"`      // True, if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members"`        // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool `json:"can_promote_members"`         // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool `json:"can_change_info"`             // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users"`            // True, if the user is allowed to invite new users to the chat
	CanPostStories      bool `json:"can_post_stories"`            // True, if the administrator can post stories to the chat
	CanEditStories      bool `json:"can_edit_stories"`            // True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool `json:"can_delete_stories"`          // True, if the administrator can delete stories posted by other users
	CanPostMessages     bool `json:"can_post_messages,omitempty"` // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool `json:"can_edit_messages,omitempty"` // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`  // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     bool `json:"can_manage_topics,omitempty"` // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                                  // Chat the user belongs to
	From                    User            `json:"from"`                                  // Performer of the action, which resulted in the change
	Date                    int64           `json:"date"`                                  // Date the change was done in Unix time
	OldChatMember           ChatMember      `json:"old_chat_member"`                       // Previous information about the chat member
	NewChatMember           ChatMember      `json:"new_chat_member"`                       // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
}

func (x *ChatMemberUpdated) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Chat                    Chat            `json:"chat"`                                  // Chat the user belongs to
		From                    User            `json:"from"`                                  // Performer of the action, which resulted in the change
		Date                    int64           `json:"date"`                                  // Date the change was done in Unix time
		OldChatMember           json.RawMessage `json:"old_chat_member"`                       // Previous information about the chat member
		NewChatMember           json.RawMessage `json:"new_chat_member"`                       // New information about the chat member
		InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
		ViaJoinRequest          bool            `json:"via_join_request,omitempty"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
		ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatMember(raw.OldChatMember); err != nil {
		return err
	} else {
		x.OldChatMember = data
	}

	if data, err := unmarshalChatMember(raw.NewChatMember); err != nil {
		return err
	} else {
		x.NewChatMember = data
	}
	x.Chat = raw.Chat
	x.From = raw.From
	x.Date = raw.Date

	x.InviteLink = raw.InviteLink
	x.ViaJoinRequest = raw.ViaJoinRequest
	x.ViaChatFolderInviteLink = raw.ViaChatFolderInviteLink
	return nil
}

// ChatMember contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
// ChatMemberOwner, ChatMemberAdministrator, ChatMemberMember, ChatMemberRestricted, ChatMemberLeft, ChatMemberBanned
type ChatMember interface {
	// IsChatMember does nothing and is only used to enforce type-safety
	IsChatMember()
}

// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	Status      string `json:"status"`                 // The member's status in the chat, always “creator”
	User        User   `json:"user"`                   // Information about the user
	IsAnonymous bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CustomTitle string `json:"custom_title,omitempty"` // Optional. Custom title for this user
}

func (ChatMemberOwner) IsChatMember() {}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	Status              string `json:"status"`                      // The member's status in the chat, always “administrator”
	User                User   `json:"user"`                        // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`               // True, if the bot is allowed to edit administrator privileges of that user
	IsAnonymous         bool   `json:"is_anonymous"`                // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`             // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool   `json:"can_delete_messages"`         // True, if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats"`      // True, if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`        // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members"`         // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`             // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`            // True, if the user is allowed to invite new users to the chat
	CanPostStories      bool   `json:"can_post_stories"`            // True, if the administrator can post stories to the chat
	CanEditStories      bool   `json:"can_edit_stories"`            // True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool   `json:"can_delete_stories"`          // True, if the administrator can delete stories posted by other users
	CanPostMessages     bool   `json:"can_post_messages,omitempty"` // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"` // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`  // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"` // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CustomTitle         string `json:"custom_title,omitempty"`      // Optional. Custom title for this user
}

func (ChatMemberAdministrator) IsChatMember() {}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status    string `json:"status"`               // The member's status in the chat, always “member”
	User      User   `json:"user"`                 // Information about the user
	UntilDate int64  `json:"until_date,omitempty"` // Optional. Date when the user's subscription will expire; Unix time
}

func (ChatMemberMember) IsChatMember() {}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always “restricted”
	User                  User   `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendAudios         bool   `json:"can_send_audios"`           // True, if the user is allowed to send audios
	CanSendDocuments      bool   `json:"can_send_documents"`        // True, if the user is allowed to send documents
	CanSendPhotos         bool   `json:"can_send_photos"`           // True, if the user is allowed to send photos
	CanSendVideos         bool   `json:"can_send_videos"`           // True, if the user is allowed to send videos
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`      // True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`      // True, if the user is allowed to send voice notes
	CanSendPolls          bool   `json:"can_send_polls"`            // True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True, if the user is allowed to add web page previews to their messages
	CanChangeInfo         bool   `json:"can_change_info"`           // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users"`          // True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages"`          // True, if the user is allowed to pin messages
	CanManageTopics       bool   `json:"can_manage_topics"`         // True, if the user is allowed to create forum topics
	UntilDate             int64  `json:"until_date"`                // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
}

func (ChatMemberRestricted) IsChatMember() {}

// Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always “left”
	User   User   `json:"user"`   // Information about the user
}

func (ChatMemberLeft) IsChatMember() {}

// Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always “kicked”
	User      User   `json:"user"`       // Information about the user
	UntilDate int64  `json:"until_date"` // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
}

func (ChatMemberBanned) IsChatMember() {}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`                  // Chat to which the request was sent
	From       User            `json:"from"`                  // User that sent the join request
	UserChatId int64           `json:"user_chat_id"`          // Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	Date       int64           `json:"date"`                  // Date the request was sent in Unix time
	Bio        string          `json:"bio,omitempty"`         // Optional. Bio of the user.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // Optional. Chat invite link that was used by the user to send the join request
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`         // Optional. True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendAudios         bool `json:"can_send_audios,omitempty"`           // Optional. True, if the user is allowed to send audios
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`        // Optional. True, if the user is allowed to send documents
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`           // Optional. True, if the user is allowed to send photos
	CanSendVideos         bool `json:"can_send_videos,omitempty"`           // Optional. True, if the user is allowed to send videos
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`      // Optional. True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`      // Optional. True, if the user is allowed to send voice notes
	CanSendPolls          bool `json:"can_send_polls,omitempty"`            // Optional. True, if the user is allowed to send polls
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"` // Optional. True, if the user is allowed to add web page previews to their messages
	CanChangeInfo         bool `json:"can_change_info,omitempty"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`         // Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages
}

// Describes the birthdate of a user.
type Birthdate struct {
	Day   int64 `json:"day"`            // Day of the user's birth; 1-31
	Month int64 `json:"month"`          // Month of the user's birth; 1-12
	Year  int64 `json:"year,omitempty"` // Optional. Year of the user's birth
}

// Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`   // Optional. Title text of the business intro
	Message string   `json:"message,omitempty"` // Optional. Message text of the business intro
	Sticker *Sticker `json:"sticker,omitempty"` // Optional. Sticker of the business intro
}

// Contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	Address  string    `json:"address"`            // Address of the business
	Location *Location `json:"location,omitempty"` // Optional. Location of the business
}

// Describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	OpeningMinute int64 `json:"opening_minute"` // The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	ClosingMinute int64 `json:"closing_minute"` // The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
}

// Describes the opening hours of a business.
type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"` // Unique name of the time zone for which the opening hours are defined
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours"`  // List of time intervals describing business opening hours
}

// Represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string   `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

// ReactionType describes the type of a reaction. Currently, it can be one of
// ReactionTypeEmoji, ReactionTypeCustomEmoji, ReactionTypePaid
type ReactionType interface {
	// IsReactionType does nothing and is only used to enforce type-safety
	IsReactionType()
}

// The reaction is based on an emoji.
type ReactionTypeEmoji struct {
	Type  string `json:"type"`  // Type of the reaction, always “emoji”
	Emoji string `json:"emoji"` // Reaction emoji. Currently, it can be one of "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func (ReactionTypeEmoji) IsReactionType() {}

// The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`            // Type of the reaction, always “custom_emoji”
	CustomEmojiId string `json:"custom_emoji_id"` // Custom emoji identifier
}

func (ReactionTypeCustomEmoji) IsReactionType() {}

// The reaction is paid.
type ReactionTypePaid struct {
	Type string `json:"type"` // Type of the reaction, always “paid”
}

func (ReactionTypePaid) IsReactionType() {}

// Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	Type       ReactionType `json:"type"`        // Type of the reaction
	TotalCount int64        `json:"total_count"` // Number of times the reaction was added
}

func (x *ReactionCount) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type       json.RawMessage `json:"type"`        // Type of the reaction
		TotalCount int64           `json:"total_count"` // Number of times the reaction was added
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalReactionType(raw.Type); err != nil {
		return err
	} else {
		x.Type = data
	}

	x.TotalCount = raw.TotalCount
	return nil
}

// MessageReactionUpdated represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`                 // The chat containing the message the user reacted to
	MessageId   int64          `json:"message_id"`           // Unique identifier of the message inside the chat
	User        *User          `json:"user,omitempty"`       // Optional. The user that changed the reaction, if the user isn't anonymous
	ActorChat   *Chat          `json:"actor_chat,omitempty"` // Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	Date        int64          `json:"date"`                 // Date of the change in Unix time
	OldReaction []ReactionType `json:"old_reaction"`         // Previous list of reaction types that were set by the user
	NewReaction []ReactionType `json:"new_reaction"`         // New list of reaction types that have been set by the user
}

// MessageReactionCountUpdated represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	Chat      Chat             `json:"chat"`       // The chat containing the message
	MessageId int64            `json:"message_id"` // Unique message identifier inside the chat
	Date      int64            `json:"date"`       // Date of the change in Unix time
	Reactions []*ReactionCount `json:"reactions"`  // List of reactions that are present on the message
}

// ForumTopic represents a forum topic.
type ForumTopic struct {
	MessageThreadId   int64  `json:"message_thread_id"`              // Unique identifier of the forum topic
	Name              string `json:"name"`                           // Name of the topic
	IconColor         int64  `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// BotCommand represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`     // Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description"` // Description of the command; 1-256 characters.
}

// BotCommandScope represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
// BotCommandScopeDefault, BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators, BotCommandScopeChat, BotCommandScopeChatAdministrators, BotCommandScopeChatMember
type BotCommandScope interface {
	// IsBotCommandScope does nothing and is only used to enforce type-safety
	IsBotCommandScope()
}

// Represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	Type string `json:"type"` // Scope type, must be default
}

func (BotCommandScopeDefault) IsBotCommandScope() {}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // Scope type, must be all_private_chats
}

func (BotCommandScopeAllPrivateChats) IsBotCommandScope() {}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // Scope type, must be all_group_chats
}

func (BotCommandScopeAllGroupChats) IsBotCommandScope() {}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // Scope type, must be all_chat_administrators
}

func (BotCommandScopeAllChatAdministrators) IsBotCommandScope() {}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	Type   string `json:"type"`    // Scope type, must be chat
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (BotCommandScopeChat) IsBotCommandScope() {}

func (x *BotCommandScopeChat) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type   string          `json:"type"`    // Scope type, must be chat
		ChatId json.RawMessage `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatID(raw.ChatId); err != nil {
		return err
	} else {
		x.ChatId = data
	}
	x.Type = raw.Type

	return nil
}

// Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`    // Scope type, must be chat_administrators
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (BotCommandScopeChatAdministrators) IsBotCommandScope() {}

func (x *BotCommandScopeChatAdministrators) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type   string          `json:"type"`    // Scope type, must be chat_administrators
		ChatId json.RawMessage `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatID(raw.ChatId); err != nil {
		return err
	} else {
		x.ChatId = data
	}
	x.Type = raw.Type

	return nil
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`    // Scope type, must be chat_member
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

func (BotCommandScopeChatMember) IsBotCommandScope() {}

func (x *BotCommandScopeChatMember) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type   string          `json:"type"`    // Scope type, must be chat_member
		ChatId json.RawMessage `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
		UserId int64           `json:"user_id"` // Unique identifier of the target user
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatID(raw.ChatId); err != nil {
		return err
	} else {
		x.ChatId = data
	}
	x.Type = raw.Type

	x.UserId = raw.UserId
	return nil
}

// BotName represents the bot's name.
type BotName struct {
	Name string `json:"name"` // The bot's name
}

// BotDescription represents the bot's description.
type BotDescription struct {
	Description string `json:"description"` // The bot's description
}

// BotShortDescription represents the bot's short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"` // The bot's short description
}

// MenuButton describes the bot's menu button in a private chat. It should be one of
// MenuButtonCommands, MenuButtonWebApp, MenuButtonDefault
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
type MenuButton interface {
	// IsMenuButton does nothing and is only used to enforce type-safety
	IsMenuButton()
}

// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	Type string `json:"type"` // Type of the button, must be commands
}

func (MenuButtonCommands) IsMenuButton() {}

// Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	Type   string     `json:"type"`    // Type of the button, must be web_app
	Text   string     `json:"text"`    // Text on the button
	WebApp WebAppInfo `json:"web_app"` // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Alternatively, a t.me link to a Web App of the bot can be specified in the object instead of the Web App's URL, in which case the Web App will be opened as if the user pressed the link.
}

func (MenuButtonWebApp) IsMenuButton() {}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	Type string `json:"type"` // Type of the button, must be default
}

func (MenuButtonDefault) IsMenuButton() {}

// ChatBoostSource describes the source of a chat boost. It can be one of
// ChatBoostSourcePremium, ChatBoostSourceGiftCode, ChatBoostSourceGiveaway
type ChatBoostSource interface {
	// IsChatBoostSource does nothing and is only used to enforce type-safety
	IsChatBoostSource()
}

// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	Source string `json:"source"` // Source of the boost, always “premium”
	User   User   `json:"user"`   // User that boosted the chat
}

func (ChatBoostSourcePremium) IsChatBoostSource() {}

// The boost was obtained by the creation of Telegram Premium gift codes to boost a chat. Each such code boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // Source of the boost, always “gift_code”
	User   User   `json:"user"`   // User for which the gift code was created
}

func (ChatBoostSourceGiftCode) IsChatBoostSource() {}

// The boost was obtained by the creation of a Telegram Premium or a Telegram Star giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription for Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`                     // Source of the boost, always “giveaway”
	GiveawayMessageId int64  `json:"giveaway_message_id"`        // Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	User              *User  `json:"user,omitempty"`             // Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	PrizeStarCount    int64  `json:"prize_star_count,omitempty"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`     // Optional. True, if the giveaway was completed, but there was no user to win the prize
}

func (ChatBoostSourceGiveaway) IsChatBoostSource() {}

// ChatBoost contains information about a chat boost.
type ChatBoost struct {
	BoostId        string          `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int64           `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int64           `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	Source         ChatBoostSource `json:"source"`          // Source of the added boost
}

func (x *ChatBoost) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		BoostId        string          `json:"boost_id"`        // Unique identifier of the boost
		AddDate        int64           `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
		ExpirationDate int64           `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
		Source         json.RawMessage `json:"source"`          // Source of the added boost
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatBoostSource(raw.Source); err != nil {
		return err
	} else {
		x.Source = data
	}
	x.BoostId = raw.BoostId
	x.AddDate = raw.AddDate
	x.ExpirationDate = raw.ExpirationDate

	return nil
}

// ChatBoostUpdated represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Information about the chat boost
}

// ChatBoostRemoved represents a boost removed from a chat.
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`        // Chat which was boosted
	BoostId    string          `json:"boost_id"`    // Unique identifier of the boost
	RemoveDate int64           `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
	Source     ChatBoostSource `json:"source"`      // Source of the removed boost
}

func (x *ChatBoostRemoved) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Chat       Chat            `json:"chat"`        // Chat which was boosted
		BoostId    string          `json:"boost_id"`    // Unique identifier of the boost
		RemoveDate int64           `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
		Source     json.RawMessage `json:"source"`      // Source of the removed boost
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalChatBoostSource(raw.Source); err != nil {
		return err
	} else {
		x.Source = data
	}
	x.Chat = raw.Chat
	x.BoostId = raw.BoostId
	x.RemoveDate = raw.RemoveDate

	return nil
}

// UserChatBoosts represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	Boosts []*ChatBoost `json:"boosts"` // The list of boosts added to the chat by the user
}

// Describes the connection of the bot with a business account.
type BusinessConnection struct {
	Id         string `json:"id"`           // Unique identifier of the business connection
	User       User   `json:"user"`         // Business account user that created the business connection
	UserChatId int64  `json:"user_chat_id"` // Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Date       int64  `json:"date"`         // Date the connection was established in Unix time
	CanReply   bool   `json:"can_reply"`    // True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	IsEnabled  bool   `json:"is_enabled"`   // True, if the connection is active
}

// BusinessMessagesDeleted is received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	BusinessConnectionId string  `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 Chat    `json:"chat"`                   // Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	MessageIds           []int64 `json:"message_ids"`            // The list of identifiers of deleted messages in the chat of the business account
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int64 `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// InputMedia represents the content of a media message to be sent. It should be one of
// InputMediaAnimation, InputMediaDocument, InputMediaAudio, InputMediaPhoto, InputMediaVideo
type InputMedia interface {
	// IsInputMedia does nothing and is only used to enforce type-safety
	IsInputMedia()

	getFiles() map[string]*InputFile
}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	Type                  string           `json:"type"`                               // Type of the result, must be photo
	Media                 *InputFile       `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the photo needs to be covered with a spoiler animation
}

func (InputMediaPhoto) IsInputMedia() {}

func (x *InputMediaPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}

	return media
}

// Represents a video to be sent.
type InputMediaVideo struct {
	Type                  string           `json:"type"`                               // Type of the result, must be video
	Media                 *InputFile       `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	Width                 int64            `json:"width,omitempty"`                    // Optional. Video width
	Height                int64            `json:"height,omitempty"`                   // Optional. Video height
	Duration              int64            `json:"duration,omitempty"`                 // Optional. Video duration in seconds
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`       // Optional. Pass True if the uploaded video is suitable for streaming
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the video needs to be covered with a spoiler animation
}

func (InputMediaVideo) IsInputMedia() {}

func (x *InputMediaVideo) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media[x.Thumbnail.Value] = x.Thumbnail
		}
	}

	return media
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type                  string           `json:"type"`                               // Type of the result, must be animation
	Media                 *InputFile       `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	Width                 int64            `json:"width,omitempty"`                    // Optional. Animation width
	Height                int64            `json:"height,omitempty"`                   // Optional. Animation height
	Duration              int64            `json:"duration,omitempty"`                 // Optional. Animation duration in seconds
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the animation needs to be covered with a spoiler animation
}

func (InputMediaAnimation) IsInputMedia() {}

func (x *InputMediaAnimation) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media[x.Thumbnail.Value] = x.Thumbnail
		}
	}

	return media
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type            string           `json:"type"`                       // Type of the result, must be audio
	Media           *InputFile       `json:"media"`                      // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail       *InputFile       `json:"thumbnail,omitempty"`        // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int64            `json:"duration,omitempty"`         // Optional. Duration of the audio in seconds
	Performer       string           `json:"performer,omitempty"`        // Optional. Performer of the audio
	Title           string           `json:"title,omitempty"`            // Optional. Title of the audio
}

func (InputMediaAudio) IsInputMedia() {}

func (x *InputMediaAudio) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media[x.Thumbnail.Value] = x.Thumbnail
		}
	}

	return media
}

// Represents a general file to be sent.
type InputMediaDocument struct {
	Type                        string           `json:"type"`                                     // Type of the result, must be document
	Media                       *InputFile       `json:"media"`                                    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`                      // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
}

func (InputMediaDocument) IsInputMedia() {}

func (x *InputMediaDocument) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media[x.Thumbnail.Value] = x.Thumbnail
		}
	}

	return media
}

// InputPaidMedia describes the paid media to be sent. Currently, it can be one of
// InputPaidMediaPhoto, InputPaidMediaVideo
type InputPaidMedia interface {
	// IsInputPaidMedia does nothing and is only used to enforce type-safety
	IsInputPaidMedia()

	getFiles() map[string]*InputFile
}

// The paid media to send is a photo.
type InputPaidMediaPhoto struct {
	Type  string     `json:"type"`  // Type of the media, must be photo
	Media *InputFile `json:"media"` // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
}

func (InputPaidMediaPhoto) IsInputPaidMedia() {}

func (x *InputPaidMediaPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}

	return media
}

// The paid media to send is a video.
type InputPaidMediaVideo struct {
	Type              string     `json:"type"`                         // Type of the media, must be video
	Media             *InputFile `json:"media"`                        // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail         *InputFile `json:"thumbnail,omitempty"`          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Width             int64      `json:"width,omitempty"`              // Optional. Video width
	Height            int64      `json:"height,omitempty"`             // Optional. Video height
	Duration          int64      `json:"duration,omitempty"`           // Optional. Video duration in seconds
	SupportsStreaming bool       `json:"supports_streaming,omitempty"` // Optional. Pass True if the uploaded video is suitable for streaming
}

func (InputPaidMediaVideo) IsInputPaidMedia() {}

func (x *InputPaidMediaVideo) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media != nil {
		if x.Media.IsUploadable() {
			media[x.Media.Value] = x.Media
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media[x.Thumbnail.Value] = x.Thumbnail
		}
	}

	return media
}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (api *API) GetMe() (*User, error) {
	return callJson[*User](api, "getMe", nil)
}

// logOut is used to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (api *API) LogOut() (bool, error) {
	return callJson[bool](api, "logOut", nil)
}

// close is used to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (api *API) Close() (bool, error) {
	return callJson[bool](api, "close", nil)
}

// sendMessage is used to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	BusinessConnectionId string              `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID              `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64               `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                 string              `json:"text"`                             // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`             // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []*MessageEntity    `json:"entities,omitempty"`               // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`   // Link preview generation options for the message
	DisableNotification  bool                `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string              `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters    `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendMessage is used to send text messages. On success, the sent Message is returned.
func (api *API) SendMessage(payload *SendMessage) (*Message, error) {
	return callJson[*Message](api, "sendMessage", payload)
}

// forwardMessage is used to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
type ForwardMessage struct {
	ChatId              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId     int64  `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId          ChatID `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool   `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Protects the contents of the forwarded message from forwarding and saving
	MessageId           int64  `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

// forwardMessage is used to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
func (api *API) ForwardMessage(payload *ForwardMessage) (*Message, error) {
	return callJson[*Message](api, "forwardMessage", payload)
}

// forwardMessages is used to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
type ForwardMessages struct {
	ChatId              ChatID  `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId     int64   `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId          ChatID  `json:"from_chat_id"`                   // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIds          []int64 `json:"message_ids"`                    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool    `json:"disable_notification,omitempty"` // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool    `json:"protect_content,omitempty"`      // Protects the contents of the forwarded messages from forwarding and saving
}

// forwardMessages is used to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
func (api *API) ForwardMessages(payload *ForwardMessages) ([]*MessageId, error) {
	return callJson[[]*MessageId](api, "forwardMessages", payload)
}

// copyMessage is used to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessage struct {
	ChatId                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId            ChatID           `json:"from_chat_id"`                       // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId             int64            `json:"message_id"`                         // Message identifier in the chat specified in from_chat_id
	Caption               string           `json:"caption,omitempty"`                  // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// copyMessage is used to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (api *API) CopyMessage(payload *CopyMessage) (*MessageId, error) {
	return callJson[*MessageId](api, "copyMessage", payload)
}

// copyMessages is used to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
type CopyMessages struct {
	ChatId              ChatID  `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId     int64   `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId          ChatID  `json:"from_chat_id"`                   // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIds          []int64 `json:"message_ids"`                    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool    `json:"disable_notification,omitempty"` // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool    `json:"protect_content,omitempty"`      // Protects the contents of the sent messages from forwarding and saving
	RemoveCaption       bool    `json:"remove_caption,omitempty"`       // Pass True to copy the messages without their captions
}

// copyMessages is used to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
func (api *API) CopyMessages(payload *CopyMessages) ([]*MessageId, error) {
	return callJson[[]*MessageId](api, "copyMessages", payload)
}

// sendPhoto is used to send photos. On success, the sent Message is returned.
type SendPhoto struct {
	BusinessConnectionId  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Photo                 *InputFile       `json:"photo"`                              // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Photo != nil {
		if x.Photo.IsUploadable() {
			media["photo"] = x.Photo
		}
	}

	return media
}

func (x *SendPhoto) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.ShowCaptionAboveMedia {
		payload["show_caption_above_media"] = strconv.FormatBool(x.ShowCaptionAboveMedia)
	}
	if x.HasSpoiler {
		payload["has_spoiler"] = strconv.FormatBool(x.HasSpoiler)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendPhoto is used to send photos. On success, the sent Message is returned.
func (api *API) SendPhoto(payload *SendPhoto) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendPhoto", params, files)
	}
	return callJson[*Message](api, "sendPhoto", payload)
}

// sendAudio is used to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Audio                *InputFile       `json:"audio"`                            // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption              string           `json:"caption,omitempty"`                // Audio caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`             // Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities      []*MessageEntity `json:"caption_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int64            `json:"duration,omitempty"`               // Duration of the audio in seconds
	Performer            string           `json:"performer,omitempty"`              // Performer
	Title                string           `json:"title,omitempty"`                  // Track name
	Thumbnail            *InputFile       `json:"thumbnail,omitempty"`              // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendAudio) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Audio != nil {
		if x.Audio.IsUploadable() {
			media["audio"] = x.Audio
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SendAudio) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.Duration != 0 {
		payload["duration"] = strconv.FormatInt(x.Duration, 10)
	}
	if x.Performer != "" {
		payload["performer"] = x.Performer
	}
	if x.Title != "" {
		payload["title"] = x.Title
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendAudio is used to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
func (api *API) SendAudio(payload *SendAudio) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendAudio", params, files)
	}
	return callJson[*Message](api, "sendAudio", payload)
}

// sendDocument is used to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	BusinessConnectionId        string           `json:"business_connection_id,omitempty"`         // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                      ChatID           `json:"chat_id"`                                  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId             int64            `json:"message_thread_id,omitempty"`              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Document                    *InputFile       `json:"document"`                                 // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`                      // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                     string           `json:"caption,omitempty"`                        // Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`                     // Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             `json:"protect_content,omitempty"`                // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast          bool             `json:"allow_paid_broadcast,omitempty"`           // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId             string           `json:"message_effect_id,omitempty"`              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`               // Description of the message to reply to
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`                   // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendDocument) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Document != nil {
		if x.Document.IsUploadable() {
			media["document"] = x.Document
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SendDocument) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.DisableContentTypeDetection {
		payload["disable_content_type_detection"] = strconv.FormatBool(x.DisableContentTypeDetection)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendDocument is used to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (api *API) SendDocument(payload *SendDocument) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendDocument", params, files)
	}
	return callJson[*Message](api, "sendDocument", payload)
}

// sendVideo is used to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	BusinessConnectionId  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Video                 *InputFile       `json:"video"`                              // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files »
	Duration              int64            `json:"duration,omitempty"`                 // Duration of sent video in seconds
	Width                 int64            `json:"width,omitempty"`                    // Video width
	Height                int64            `json:"height,omitempty"`                   // Video height
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the video needs to be covered with a spoiler animation
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`       // Pass True if the uploaded video is suitable for streaming
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendVideo) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Video != nil {
		if x.Video.IsUploadable() {
			media["video"] = x.Video
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SendVideo) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Duration != 0 {
		payload["duration"] = strconv.FormatInt(x.Duration, 10)
	}
	if x.Width != 0 {
		payload["width"] = strconv.FormatInt(x.Width, 10)
	}
	if x.Height != 0 {
		payload["height"] = strconv.FormatInt(x.Height, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.ShowCaptionAboveMedia {
		payload["show_caption_above_media"] = strconv.FormatBool(x.ShowCaptionAboveMedia)
	}
	if x.HasSpoiler {
		payload["has_spoiler"] = strconv.FormatBool(x.HasSpoiler)
	}
	if x.SupportsStreaming {
		payload["supports_streaming"] = strconv.FormatBool(x.SupportsStreaming)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendVideo is used to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (api *API) SendVideo(payload *SendVideo) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendVideo", params, files)
	}
	return callJson[*Message](api, "sendVideo", payload)
}

// sendAnimation is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	BusinessConnectionId  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Animation             *InputFile       `json:"animation"`                          // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files »
	Duration              int64            `json:"duration,omitempty"`                 // Duration of sent animation in seconds
	Width                 int64            `json:"width,omitempty"`                    // Animation width
	Height                int64            `json:"height,omitempty"`                   // Animation height
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the animation needs to be covered with a spoiler animation
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendAnimation) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Animation != nil {
		if x.Animation.IsUploadable() {
			media["animation"] = x.Animation
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SendAnimation) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Duration != 0 {
		payload["duration"] = strconv.FormatInt(x.Duration, 10)
	}
	if x.Width != 0 {
		payload["width"] = strconv.FormatInt(x.Width, 10)
	}
	if x.Height != 0 {
		payload["height"] = strconv.FormatInt(x.Height, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.ShowCaptionAboveMedia {
		payload["show_caption_above_media"] = strconv.FormatBool(x.ShowCaptionAboveMedia)
	}
	if x.HasSpoiler {
		payload["has_spoiler"] = strconv.FormatBool(x.HasSpoiler)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendAnimation is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (api *API) SendAnimation(payload *SendAnimation) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendAnimation", params, files)
	}
	return callJson[*Message](api, "sendAnimation", payload)
}

// sendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Voice                *InputFile       `json:"voice"`                            // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption              string           `json:"caption,omitempty"`                // Voice message caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`             // Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities      []*MessageEntity `json:"caption_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int64            `json:"duration,omitempty"`               // Duration of the voice message in seconds
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendVoice) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Voice != nil {
		if x.Voice.IsUploadable() {
			media["voice"] = x.Voice
		}
	}

	return media
}

func (x *SendVoice) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.Duration != 0 {
		payload["duration"] = strconv.FormatInt(x.Duration, 10)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (api *API) SendVoice(payload *SendVoice) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendVoice", params, files)
	}
	return callJson[*Message](api, "sendVoice", payload)
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. sendVideoNote is used to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	VideoNote            *InputFile       `json:"video_note"`                       // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration             int64            `json:"duration,omitempty"`               // Duration of sent video in seconds
	Length               int64            `json:"length,omitempty"`                 // Video width and height, i.e. diameter of the video message
	Thumbnail            *InputFile       `json:"thumbnail,omitempty"`              // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendVideoNote) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.VideoNote != nil {
		if x.VideoNote.IsUploadable() {
			media["video_note"] = x.VideoNote
		}
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SendVideoNote) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Duration != 0 {
		payload["duration"] = strconv.FormatInt(x.Duration, 10)
	}
	if x.Length != 0 {
		payload["length"] = strconv.FormatInt(x.Length, 10)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. sendVideoNote is used to send video messages. On success, the sent Message is returned.
func (api *API) SendVideoNote(payload *SendVideoNote) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendVideoNote", params, files)
	}
	return callJson[*Message](api, "sendVideoNote", payload)
}

// sendPaidMedia is used to send paid media. On success, the sent Message is returned.
type SendPaidMedia struct {
	BusinessConnectionId  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
	StarCount             int64            `json:"star_count"`                         // The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	Media                 []InputPaidMedia `json:"media"`                              // A JSON-serialized array describing the media to be sent; up to 10 items
	Payload               string           `json:"payload,omitempty"`                  // Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Caption               string           `json:"caption,omitempty"`                  // Media caption, 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the media caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendPaidMedia) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for _, m := range x.Media {
		for key, value := range m.getFiles() {
			media[key] = value
		}
	}

	return media
}

func (x *SendPaidMedia) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	payload["star_count"] = strconv.FormatInt(x.StarCount, 10)
	if bb, err := json.Marshal(x.Media); err != nil {
		return nil, err
	} else {
		payload["media"] = string(bb)
	}
	if x.Payload != "" {
		payload["payload"] = x.Payload
	}
	if x.Caption != "" {
		payload["caption"] = x.Caption
	}
	if x.ParseMode != ParseModeNone {
		payload["parse_mode"] = string(x.ParseMode)
	}
	if x.CaptionEntities != nil {
		if bb, err := json.Marshal(x.CaptionEntities); err != nil {
			return nil, err
		} else {
			payload["caption_entities"] = string(bb)
		}
	}
	if x.ShowCaptionAboveMedia {
		payload["show_caption_above_media"] = strconv.FormatBool(x.ShowCaptionAboveMedia)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendPaidMedia is used to send paid media. On success, the sent Message is returned.
func (api *API) SendPaidMedia(payload *SendPaidMedia) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendPaidMedia", params, files)
	}
	return callJson[*Message](api, "sendPaidMedia", payload)
}

// sendMediaGroup is used to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Media                []InputMedia     `json:"media"`                            // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends messages silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent messages from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
}

func (x *SendMediaGroup) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for _, m := range x.Media {
		for key, value := range m.getFiles() {
			media[key] = value
		}
	}

	return media
}

func (x *SendMediaGroup) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if bb, err := json.Marshal(x.Media); err != nil {
		return nil, err
	} else {
		payload["media"] = string(bb)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}

	return payload, nil
}

// sendMediaGroup is used to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
func (api *API) SendMediaGroup(payload *SendMediaGroup) ([]*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[[]*Message](api, "sendMediaGroup", params, files)
	}
	return callJson[[]*Message](api, "sendMediaGroup", payload)
}

// sendLocation is used to send point on the map. On success, the sent Message is returned.
type SendLocation struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          `json:"latitude"`                         // Latitude of the location
	Longitude            float64          `json:"longitude"`                        // Longitude of the location
	HorizontalAccuracy   float64          `json:"horizontal_accuracy,omitempty"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64            `json:"live_period,omitempty"`            // Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64            `json:"heading,omitempty"`                // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64            `json:"proximity_alert_radius,omitempty"` // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendLocation is used to send point on the map. On success, the sent Message is returned.
func (api *API) SendLocation(payload *SendLocation) (*Message, error) {
	return callJson[*Message](api, "sendLocation", payload)
}

// sendVenue is used to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          `json:"latitude"`                         // Latitude of the venue
	Longitude            float64          `json:"longitude"`                        // Longitude of the venue
	Title                string           `json:"title"`                            // Name of the venue
	Address              string           `json:"address"`                          // Address of the venue
	FoursquareId         string           `json:"foursquare_id,omitempty"`          // Foursquare identifier of the venue
	FoursquareType       string           `json:"foursquare_type,omitempty"`        // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId        string           `json:"google_place_id,omitempty"`        // Google Places identifier of the venue
	GooglePlaceType      string           `json:"google_place_type,omitempty"`      // Google Places type of the venue. (See supported types.)
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendVenue is used to send information about a venue. On success, the sent Message is returned.
func (api *API) SendVenue(payload *SendVenue) (*Message, error) {
	return callJson[*Message](api, "sendVenue", payload)
}

// sendContact is used to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	PhoneNumber          string           `json:"phone_number"`                     // Contact's phone number
	FirstName            string           `json:"first_name"`                       // Contact's first name
	LastName             string           `json:"last_name,omitempty"`              // Contact's last name
	Vcard                string           `json:"vcard,omitempty"`                  // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendContact is used to send phone contacts. On success, the sent Message is returned.
func (api *API) SendContact(payload *SendContact) (*Message, error) {
	return callJson[*Message](api, "sendContact", payload)
}

// sendPoll is used to send a native poll. On success, the sent Message is returned.
type SendPoll struct {
	BusinessConnectionId  string             `json:"business_connection_id,omitempty"`  // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId                ChatID             `json:"chat_id"`                           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId       int64              `json:"message_thread_id,omitempty"`       // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Question              string             `json:"question"`                          // Poll question, 1-300 characters
	QuestionParseMode     string             `json:"question_parse_mode,omitempty"`     // Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed
	QuestionEntities      []*MessageEntity   `json:"question_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of question_parse_mode
	Options               []*InputPollOption `json:"options"`                           // A JSON-serialized list of 2-10 answer options
	IsAnonymous           bool               `json:"is_anonymous,omitempty"`            // True, if the poll needs to be anonymous, defaults to True
	Type                  string             `json:"type,omitempty"`                    // Poll type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers bool               `json:"allows_multiple_answers,omitempty"` // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId       int64              `json:"correct_option_id,omitempty"`       // 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation           string             `json:"explanation,omitempty"`             // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode  string             `json:"explanation_parse_mode,omitempty"`  // Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities   []*MessageEntity   `json:"explanation_entities,omitempty"`    // A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode
	OpenPeriod            int64              `json:"open_period,omitempty"`             // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate             int64              `json:"close_date,omitempty"`              // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed              bool               `json:"is_closed,omitempty"`               // Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool               `json:"disable_notification,omitempty"`    // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool               `json:"protect_content,omitempty"`         // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool               `json:"allow_paid_broadcast,omitempty"`    // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId       string             `json:"message_effect_id,omitempty"`       // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters   `json:"reply_parameters,omitempty"`        // Description of the message to reply to
	ReplyMarkup           ReplyMarkup        `json:"reply_markup,omitempty"`            // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendPoll is used to send a native poll. On success, the sent Message is returned.
func (api *API) SendPoll(payload *SendPoll) (*Message, error) {
	return callJson[*Message](api, "sendPoll", payload)
}

// sendDice is used to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDice struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                string           `json:"emoji,omitempty"`                  // Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64 for “”. Defaults to “”
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// sendDice is used to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (api *API) SendDice(payload *SendDice) (*Message, error) {
	return callJson[*Message](api, "sendDice", payload)
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
type SendChatAction struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the action will be sent
	ChatId               ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64  `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread; for supergroups only
	Action               string `json:"action"`                           // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (api *API) SendChatAction(payload *SendChatAction) (bool, error) {
	return callJson[bool](api, "sendChatAction", payload)
}

// setMessageReaction is used to change the chosen reactions on a message. Service messages can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can't use paid reactions. Returns True on success.
type SetMessageReaction struct {
	ChatId    ChatID         `json:"chat_id"`            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64          `json:"message_id"`         // Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	Reaction  []ReactionType `json:"reaction,omitempty"` // A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	IsBig     bool           `json:"is_big,omitempty"`   // Pass True to set the reaction with a big animation
}

// setMessageReaction is used to change the chosen reactions on a message. Service messages can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can't use paid reactions. Returns True on success.
func (api *API) SetMessageReaction(payload *SetMessageReaction) (bool, error) {
	return callJson[bool](api, "setMessageReaction", payload)
}

// getUserProfilePhotos is used to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
	UserId int64 `json:"user_id"`          // Unique identifier of the target user
	Offset int64 `json:"offset,omitempty"` // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int64 `json:"limit,omitempty"`  // Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

// getUserProfilePhotos is used to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (api *API) GetUserProfilePhotos(payload *GetUserProfilePhotos) (*UserProfilePhotos, error) {
	return callJson[*UserProfilePhotos](api, "getUserProfilePhotos", payload)
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
type SetUserEmojiStatus struct {
	UserId                    int64  `json:"user_id"`                                // Unique identifier of the target user
	EmojiStatusCustomEmojiId  string `json:"emoji_status_custom_emoji_id,omitempty"` // Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusExpirationDate int64  `json:"emoji_status_expiration_date,omitempty"` // Expiration date of the emoji status, if any
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
func (api *API) SetUserEmojiStatus(payload *SetUserEmojiStatus) (bool, error) {
	return callJson[bool](api, "setUserEmojiStatus", payload)
}

// getFile is used to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	FileId string `json:"file_id"` // File identifier to get information about
}

// getFile is used to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (api *API) GetFile(payload *GetFile) (*File, error) {
	return callJson[*File](api, "getFile", payload)
}

// banChatMember is used to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatMember struct {
	ChatId         ChatID `json:"chat_id"`                   // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId         int64  `json:"user_id"`                   // Unique identifier of the target user
	UntilDate      int64  `json:"until_date,omitempty"`      // Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool   `json:"revoke_messages,omitempty"` // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

// banChatMember is used to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) BanChatMember(payload *BanChatMember) (bool, error) {
	return callJson[bool](api, "banChatMember", payload)
}

// unbanChatMember is used to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
type UnbanChatMember struct {
	ChatId       ChatID `json:"chat_id"`                  // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId       int64  `json:"user_id"`                  // Unique identifier of the target user
	OnlyIfBanned bool   `json:"only_if_banned,omitempty"` // Do nothing if the user is not banned
}

// unbanChatMember is used to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (api *API) UnbanChatMember(payload *UnbanChatMember) (bool, error) {
	return callJson[bool](api, "unbanChatMember", payload)
}

// restrictChatMember is used to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
	ChatId                        ChatID          `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId                        int64           `json:"user_id"`                                    // Unique identifier of the target user
	Permissions                   ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new user permissions
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"` // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UntilDate                     int64           `json:"until_date,omitempty"`                       // Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

// restrictChatMember is used to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (api *API) RestrictChatMember(payload *RestrictChatMember) (bool, error) {
	return callJson[bool](api, "restrictChatMember", payload)
}

// promoteChatMember is used to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMember struct {
	ChatId              ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId              int64  `json:"user_id"`                          // Unique identifier of the target user
	IsAnonymous         bool   `json:"is_anonymous,omitempty"`           // Pass True if the administrator's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat,omitempty"`        // Pass True if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool   `json:"can_delete_messages,omitempty"`    // Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats,omitempty"` // Pass True if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members,omitempty"`   // Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members,omitempty"`    // Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool   `json:"can_change_info,omitempty"`        // Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users,omitempty"`       // Pass True if the administrator can invite new users to the chat
	CanPostStories      bool   `json:"can_post_stories,omitempty"`       // Pass True if the administrator can post stories to the chat
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`       // Pass True if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"`     // Pass True if the administrator can delete stories posted by other users
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`      // Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`      // Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`       // Pass True if the administrator can pin messages; for supergroups only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`      // Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

// promoteChatMember is used to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (api *API) PromoteChatMember(payload *PromoteChatMember) (bool, error) {
	return callJson[bool](api, "promoteChatMember", payload)
}

// setChatAdministratorCustomTitle is used to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitle struct {
	ChatId      ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int64  `json:"user_id"`      // Unique identifier of the target user
	CustomTitle string `json:"custom_title"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

// setChatAdministratorCustomTitle is used to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (api *API) SetChatAdministratorCustomTitle(payload *SetChatAdministratorCustomTitle) (bool, error) {
	return callJson[bool](api, "setChatAdministratorCustomTitle", payload)
}

// banChatSenderChat is used to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatSenderChat struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// banChatSenderChat is used to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) BanChatSenderChat(payload *BanChatSenderChat) (bool, error) {
	return callJson[bool](api, "banChatSenderChat", payload)
}

// unbanChatSenderChat is used to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
type UnbanChatSenderChat struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// unbanChatSenderChat is used to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) UnbanChatSenderChat(payload *UnbanChatSenderChat) (bool, error) {
	return callJson[bool](api, "unbanChatSenderChat", payload)
}

// setChatPermissions is used to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
type SetChatPermissions struct {
	ChatId                        ChatID          `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions                   ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new default chat permissions
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"` // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
}

// setChatPermissions is used to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (api *API) SetChatPermissions(payload *SetChatPermissions) (bool, error) {
	return callJson[bool](api, "setChatPermissions", payload)
}

// exportChatInviteLink is used to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLink struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// exportChatInviteLink is used to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
func (api *API) ExportChatInviteLink(payload *ExportChatInviteLink) (string, error) {
	return callJson[string](api, "exportChatInviteLink", payload)
}

// createChatInviteLink is used to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLink struct {
	ChatId             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`                 // Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// createChatInviteLink is used to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (api *API) CreateChatInviteLink(payload *CreateChatInviteLink) (*ChatInviteLink, error) {
	return callJson[*ChatInviteLink](api, "createChatInviteLink", payload)
}

// editChatInviteLink is used to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLink struct {
	ChatId             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink         string `json:"invite_link"`                    // The invite link to edit
	Name               string `json:"name,omitempty"`                 // Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// editChatInviteLink is used to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (api *API) EditChatInviteLink(payload *EditChatInviteLink) (*ChatInviteLink, error) {
	return callJson[*ChatInviteLink](api, "editChatInviteLink", payload)
}

// createChatSubscriptionInviteLink is used to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
type CreateChatSubscriptionInviteLink struct {
	ChatId             ChatID `json:"chat_id"`             // Unique identifier for the target channel chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`      // Invite link name; 0-32 characters
	SubscriptionPeriod int64  `json:"subscription_period"` // The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPrice  int64  `json:"subscription_price"`  // The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
}

// createChatSubscriptionInviteLink is used to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
func (api *API) CreateChatSubscriptionInviteLink(payload *CreateChatSubscriptionInviteLink) (*ChatInviteLink, error) {
	return callJson[*ChatInviteLink](api, "createChatSubscriptionInviteLink", payload)
}

// editChatSubscriptionInviteLink is used to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatSubscriptionInviteLink struct {
	ChatId     ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"`    // The invite link to edit
	Name       string `json:"name,omitempty"` // Invite link name; 0-32 characters
}

// editChatSubscriptionInviteLink is used to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (api *API) EditChatSubscriptionInviteLink(payload *EditChatSubscriptionInviteLink) (*ChatInviteLink, error) {
	return callJson[*ChatInviteLink](api, "editChatSubscriptionInviteLink", payload)
}

// revokeChatInviteLink is used to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLink struct {
	ChatId     ChatID `json:"chat_id"`     // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"` // The invite link to revoke
}

// revokeChatInviteLink is used to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
func (api *API) RevokeChatInviteLink(payload *RevokeChatInviteLink) (*ChatInviteLink, error) {
	return callJson[*ChatInviteLink](api, "revokeChatInviteLink", payload)
}

// approveChatJoinRequest is used to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type ApproveChatJoinRequest struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// approveChatJoinRequest is used to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (api *API) ApproveChatJoinRequest(payload *ApproveChatJoinRequest) (bool, error) {
	return callJson[bool](api, "approveChatJoinRequest", payload)
}

// declineChatJoinRequest is used to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type DeclineChatJoinRequest struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// declineChatJoinRequest is used to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (api *API) DeclineChatJoinRequest(payload *DeclineChatJoinRequest) (bool, error) {
	return callJson[bool](api, "declineChatJoinRequest", payload)
}

// setChatPhoto is used to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatPhoto struct {
	ChatId ChatID     `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  *InputFile `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

func (x *SetChatPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Photo != nil {
		if x.Photo.IsUploadable() {
			media["photo"] = x.Photo
		}
	}

	return media
}

func (x *SetChatPhoto) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}

	return payload, nil
}

// setChatPhoto is used to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) SetChatPhoto(payload *SetChatPhoto) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "setChatPhoto", params, files)
	}
	return callJson[bool](api, "setChatPhoto", payload)
}

// deleteChatPhoto is used to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type DeleteChatPhoto struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// deleteChatPhoto is used to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) DeleteChatPhoto(payload *DeleteChatPhoto) (bool, error) {
	return callJson[bool](api, "deleteChatPhoto", payload)
}

// setChatTitle is used to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatTitle struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-128 characters
}

// setChatTitle is used to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) SetChatTitle(payload *SetChatTitle) (bool, error) {
	return callJson[bool](api, "setChatTitle", payload)
}

// setChatDescription is used to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatDescription struct {
	ChatId      ChatID `json:"chat_id"`               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string `json:"description,omitempty"` // New chat description, 0-255 characters
}

// setChatDescription is used to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (api *API) SetChatDescription(payload *SetChatDescription) (bool, error) {
	return callJson[bool](api, "setChatDescription", payload)
}

// pinChatMessage is used to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type PinChatMessage struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be pinned
	ChatId               ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64  `json:"message_id"`                       // Identifier of a message to pin
	DisableNotification  bool   `json:"disable_notification,omitempty"`   // Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

// pinChatMessage is used to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (api *API) PinChatMessage(payload *PinChatMessage) (bool, error) {
	return callJson[bool](api, "pinChatMessage", payload)
}

// unpinChatMessage is used to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinChatMessage struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be unpinned
	ChatId               ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64  `json:"message_id,omitempty"`             // Identifier of the message to unpin. Required if business_connection_id is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

// unpinChatMessage is used to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (api *API) UnpinChatMessage(payload *UnpinChatMessage) (bool, error) {
	return callJson[bool](api, "unpinChatMessage", payload)
}

// unpinAllChatMessages is used to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinAllChatMessages struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// unpinAllChatMessages is used to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (api *API) UnpinAllChatMessages(payload *UnpinAllChatMessages) (bool, error) {
	return callJson[bool](api, "unpinAllChatMessages", payload)
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (api *API) LeaveChat(payload *LeaveChat) (bool, error) {
	return callJson[bool](api, "leaveChat", payload)
}

// getChat is used to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
type GetChat struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// getChat is used to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
func (api *API) GetChat(payload *GetChat) (*ChatFullInfo, error) {
	return callJson[*ChatFullInfo](api, "getChat", payload)
}

// getChatAdministrators is used to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
type GetChatAdministrators struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// getChatAdministrators is used to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
func (api *API) GetChatAdministrators(payload *GetChatAdministrators) ([]ChatMember, error) {
	return callJson[[]ChatMember](api, "getChatAdministrators", payload)
}

// getChatMemberCount is used to get the number of members in a chat. Returns Int on success.
type GetChatMemberCount struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// getChatMemberCount is used to get the number of members in a chat. Returns Int on success.
func (api *API) GetChatMemberCount(payload *GetChatMemberCount) (int64, error) {
	return callJson[int64](api, "getChatMemberCount", payload)
}

// getChatMember is used to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
type GetChatMember struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// getChatMember is used to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
func (api *API) GetChatMember(payload *GetChatMember) (ChatMember, error) {
	resp, err := callJson[json.RawMessage](api, "getChatMember", payload)
	if err != nil {
		return nil, err
	}
	return unmarshalChatMember(resp)
}

// setChatStickerSet is used to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSet struct {
	ChatId         ChatID `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

// setChatStickerSet is used to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (api *API) SetChatStickerSet(payload *SetChatStickerSet) (bool, error) {
	return callJson[bool](api, "setChatStickerSet", payload)
}

// deleteChatStickerSet is used to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSet struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// deleteChatStickerSet is used to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (api *API) DeleteChatStickerSet(payload *DeleteChatStickerSet) (bool, error) {
	return callJson[bool](api, "deleteChatStickerSet", payload)
}

// getForumTopicIconStickers is used to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
func (api *API) GetForumTopicIconStickers() ([]*Sticker, error) {
	return callJson[[]*Sticker](api, "getForumTopicIconStickers", nil)
}

// createForumTopic is used to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
type CreateForumTopic struct {
	ChatId            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name              string `json:"name"`                           // Topic name, 1-128 characters
	IconColor         int64  `json:"icon_color,omitempty"`           // Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
}

// createForumTopic is used to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
func (api *API) CreateForumTopic(payload *CreateForumTopic) (*ForumTopic, error) {
	return callJson[*ForumTopic](api, "createForumTopic", payload)
}

// editForumTopic is used to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type EditForumTopic struct {
	ChatId            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId   int64  `json:"message_thread_id"`              // Unique identifier for the target message thread of the forum topic
	Name              string `json:"name,omitempty"`                 // New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

// editForumTopic is used to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (api *API) EditForumTopic(payload *EditForumTopic) (bool, error) {
	return callJson[bool](api, "editForumTopic", payload)
}

// closeForumTopic is used to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type CloseForumTopic struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// closeForumTopic is used to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (api *API) CloseForumTopic(payload *CloseForumTopic) (bool, error) {
	return callJson[bool](api, "closeForumTopic", payload)
}

// reopenForumTopic is used to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type ReopenForumTopic struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// reopenForumTopic is used to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (api *API) ReopenForumTopic(payload *ReopenForumTopic) (bool, error) {
	return callJson[bool](api, "reopenForumTopic", payload)
}

// deleteForumTopic is used to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
type DeleteForumTopic struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// deleteForumTopic is used to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (api *API) DeleteForumTopic(payload *DeleteForumTopic) (bool, error) {
	return callJson[bool](api, "deleteForumTopic", payload)
}

// unpinAllForumTopicMessages is used to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllForumTopicMessages struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// unpinAllForumTopicMessages is used to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (api *API) UnpinAllForumTopicMessages(payload *UnpinAllForumTopicMessages) (bool, error) {
	return callJson[bool](api, "unpinAllForumTopicMessages", payload)
}

// editGeneralForumTopic is used to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type EditGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // New topic name, 1-128 characters
}

// editGeneralForumTopic is used to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (api *API) EditGeneralForumTopic(payload *EditGeneralForumTopic) (bool, error) {
	return callJson[bool](api, "editGeneralForumTopic", payload)
}

// closeGeneralForumTopic is used to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type CloseGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// closeGeneralForumTopic is used to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (api *API) CloseGeneralForumTopic(payload *CloseGeneralForumTopic) (bool, error) {
	return callJson[bool](api, "closeGeneralForumTopic", payload)
}

// reopenGeneralForumTopic is used to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
type ReopenGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// reopenGeneralForumTopic is used to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (api *API) ReopenGeneralForumTopic(payload *ReopenGeneralForumTopic) (bool, error) {
	return callJson[bool](api, "reopenGeneralForumTopic", payload)
}

// hideGeneralForumTopic is used to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
type HideGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// hideGeneralForumTopic is used to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (api *API) HideGeneralForumTopic(payload *HideGeneralForumTopic) (bool, error) {
	return callJson[bool](api, "hideGeneralForumTopic", payload)
}

// unhideGeneralForumTopic is used to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type UnhideGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// unhideGeneralForumTopic is used to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (api *API) UnhideGeneralForumTopic(payload *UnhideGeneralForumTopic) (bool, error) {
	return callJson[bool](api, "unhideGeneralForumTopic", payload)
}

// unpinAllGeneralForumTopicMessages is used to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllGeneralForumTopicMessages struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// unpinAllGeneralForumTopicMessages is used to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (api *API) UnpinAllGeneralForumTopicMessages(payload *UnpinAllGeneralForumTopicMessages) (bool, error) {
	return callJson[bool](api, "unpinAllGeneralForumTopicMessages", payload)
}

// answerCallbackQuery is used to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`    // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`       // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert,omitempty"` // If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url,omitempty"`        // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int64  `json:"cache_time,omitempty"` // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

// answerCallbackQuery is used to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
func (api *API) AnswerCallbackQuery(payload *AnswerCallbackQuery) (bool, error) {
	return callJson[bool](api, "answerCallbackQuery", payload)
}

// getUserChatBoosts is used to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
type GetUserChatBoosts struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the chat or username of the channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// getUserChatBoosts is used to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
func (api *API) GetUserChatBoosts(payload *GetUserChatBoosts) (*UserChatBoosts, error) {
	return callJson[*UserChatBoosts](api, "getUserChatBoosts", payload)
}

// getBusinessConnection is used to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
type GetBusinessConnection struct {
	BusinessConnectionId string `json:"business_connection_id"` // Unique identifier of the business connection
}

// getBusinessConnection is used to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
func (api *API) GetBusinessConnection(payload *GetBusinessConnection) (*BusinessConnection, error) {
	return callJson[*BusinessConnection](api, "getBusinessConnection", payload)
}

// setMyCommands is used to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
type SetMyCommands struct {
	Commands     []*BotCommand   `json:"commands"`                // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string          `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// setMyCommands is used to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (api *API) SetMyCommands(payload *SetMyCommands) (bool, error) {
	return callJson[bool](api, "setMyCommands", payload)
}

// deleteMyCommands is used to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommands struct {
	Scope        BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string          `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// deleteMyCommands is used to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (api *API) DeleteMyCommands(payload *DeleteMyCommands) (bool, error) {
	return callJson[bool](api, "deleteMyCommands", payload)
}

// getMyCommands is used to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	Scope        BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string          `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

// getMyCommands is used to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
func (api *API) GetMyCommands(payload *GetMyCommands) ([]*BotCommand, error) {
	return callJson[[]*BotCommand](api, "getMyCommands", payload)
}

// setMyName is used to change the bot's name. Returns True on success.
type SetMyName struct {
	Name         string `json:"name,omitempty"`          // New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
}

// setMyName is used to change the bot's name. Returns True on success.
func (api *API) SetMyName(payload *SetMyName) (bool, error) {
	return callJson[bool](api, "setMyName", payload)
}

// getMyName is used to get the current bot name for the given user language. Returns BotName on success.
type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

// getMyName is used to get the current bot name for the given user language. Returns BotName on success.
func (api *API) GetMyName(payload *GetMyName) (*BotName, error) {
	return callJson[*BotName](api, "getMyName", payload)
}

// setMyDescription is used to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
type SetMyDescription struct {
	Description  string `json:"description,omitempty"`   // New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
}

// setMyDescription is used to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
func (api *API) SetMyDescription(payload *SetMyDescription) (bool, error) {
	return callJson[bool](api, "setMyDescription", payload)
}

// getMyDescription is used to get the current bot description for the given user language. Returns BotDescription on success.
type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

// getMyDescription is used to get the current bot description for the given user language. Returns BotDescription on success.
func (api *API) GetMyDescription(payload *GetMyDescription) (*BotDescription, error) {
	return callJson[*BotDescription](api, "getMyDescription", payload)
}

// setMyShortDescription is used to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"` // New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	LanguageCode     string `json:"language_code,omitempty"`     // A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
}

// setMyShortDescription is used to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
func (api *API) SetMyShortDescription(payload *SetMyShortDescription) (bool, error) {
	return callJson[bool](api, "setMyShortDescription", payload)
}

// getMyShortDescription is used to get the current bot short description for the given user language. Returns BotShortDescription on success.
type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

// getMyShortDescription is used to get the current bot short description for the given user language. Returns BotShortDescription on success.
func (api *API) GetMyShortDescription(payload *GetMyShortDescription) (*BotShortDescription, error) {
	return callJson[*BotShortDescription](api, "getMyShortDescription", payload)
}

// setChatMenuButton is used to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
type SetChatMenuButton struct {
	ChatId     int64      `json:"chat_id,omitempty"`     // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton MenuButton `json:"menu_button,omitempty"` // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}

// setChatMenuButton is used to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (api *API) SetChatMenuButton(payload *SetChatMenuButton) (bool, error) {
	return callJson[bool](api, "setChatMenuButton", payload)
}

// getChatMenuButton is used to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
type GetChatMenuButton struct {
	ChatId int64 `json:"chat_id,omitempty"` // Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
}

// getChatMenuButton is used to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
func (api *API) GetChatMenuButton(payload *GetChatMenuButton) (MenuButton, error) {
	resp, err := callJson[json.RawMessage](api, "getChatMenuButton", payload)
	if err != nil {
		return nil, err
	}
	return unmarshalMenuButton(resp)
}

// setMyDefaultAdministratorRights is used to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
type SetMyDefaultAdministratorRights struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`       // A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	ForChannels bool                     `json:"for_channels,omitempty"` // Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
}

// setMyDefaultAdministratorRights is used to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
func (api *API) SetMyDefaultAdministratorRights(payload *SetMyDefaultAdministratorRights) (bool, error) {
	return callJson[bool](api, "setMyDefaultAdministratorRights", payload)
}

// getMyDefaultAdministratorRights is used to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"` // Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
}

// getMyDefaultAdministratorRights is used to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
func (api *API) GetMyDefaultAdministratorRights(payload *GetMyDefaultAdministratorRights) (*ChatAdministratorRights, error) {
	return callJson[*ChatAdministratorRights](api, "getMyDefaultAdministratorRights", payload)
}

// editMessageText is used to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageText struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                 string                `json:"text"`                             // New text of the message, 1-4096 characters after entities parsing
	ParseMode            ParseMode             `json:"parse_mode,omitempty"`             // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []*MessageEntity      `json:"entities,omitempty"`               // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions   `json:"link_preview_options,omitempty"`   // Link preview generation options for the message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard.
}

// editMessageText is used to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (api *API) EditMessageText(payload *EditMessageText) (*Message, error) {
	return callJson[*Message](api, "editMessageText", payload)
}

// editMessageCaption is used to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageCaption struct {
	BusinessConnectionId  string                `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId                ChatID                `json:"chat_id,omitempty"`                  // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int64                 `json:"message_id,omitempty"`               // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string                `json:"inline_message_id,omitempty"`        // Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption               string                `json:"caption,omitempty"`                  // New caption of the message, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // A JSON-serialized object for an inline keyboard.
}

// editMessageCaption is used to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (api *API) EditMessageCaption(payload *EditMessageCaption) (*Message, error) {
	return callJson[*Message](api, "editMessageCaption", payload)
}

// editMessageMedia is used to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageMedia struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Media                InputMedia            `json:"media"`                            // A JSON-serialized object for a new media content of the message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

func (x *EditMessageMedia) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for key, value := range x.Media.getFiles() {
		media[key] = value
	}

	return media
}

func (x *EditMessageMedia) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if x.ChatId != nil {
		if bb, err := json.Marshal(x.ChatId); err != nil {
			return nil, err
		} else {
			payload["chat_id"] = string(bb)
		}
	}
	if x.MessageId != 0 {
		payload["message_id"] = strconv.FormatInt(x.MessageId, 10)
	}
	if x.InlineMessageId != "" {
		payload["inline_message_id"] = x.InlineMessageId
	}
	if bb, err := json.Marshal(x.Media); err != nil {
		return nil, err
	} else {
		payload["media"] = string(bb)
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// editMessageMedia is used to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (api *API) EditMessageMedia(payload *EditMessageMedia) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "editMessageMedia", params, files)
	}
	return callJson[*Message](api, "editMessageMedia", payload)
}

// editMessageLiveLocation is used to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               `json:"latitude"`                         // Latitude of new location
	Longitude            float64               `json:"longitude"`                        // Longitude of new location
	LivePeriod           int64                 `json:"live_period,omitempty"`            // New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current live_period by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then live_period remains unchanged
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int64                 `json:"heading,omitempty"`                // Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

// editMessageLiveLocation is used to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) EditMessageLiveLocation(payload *EditMessageLiveLocation) (*Message, error) {
	return callJson[*Message](api, "editMessageLiveLocation", payload)
}

// stopMessageLiveLocation is used to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

// stopMessageLiveLocation is used to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) StopMessageLiveLocation(payload *StopMessageLiveLocation) (*Message, error) {
	return callJson[*Message](api, "stopMessageLiveLocation", payload)
}

// editMessageReplyMarkup is used to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageReplyMarkup struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard.
}

// editMessageReplyMarkup is used to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (api *API) EditMessageReplyMarkup(payload *EditMessageReplyMarkup) (*Message, error) {
	return callJson[*Message](api, "editMessageReplyMarkup", payload)
}

// stopPoll is used to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
type StopPoll struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatId               ChatID                `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id"`                       // Identifier of the original message with the poll
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new message inline keyboard.
}

// stopPoll is used to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
func (api *API) StopPoll(payload *StopPoll) (*Poll, error) {
	return callJson[*Poll](api, "stopPoll", payload)
}

// deleteMessage is used to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can't be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
type DeleteMessage struct {
	ChatId    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id"` // Identifier of the message to delete
}

// deleteMessage is used to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can't be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
func (api *API) DeleteMessage(payload *DeleteMessage) (bool, error) {
	return callJson[bool](api, "deleteMessage", payload)
}

// deleteMessages is used to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
type DeleteMessages struct {
	ChatId     ChatID  `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageIds []int64 `json:"message_ids"` // A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted
}

// deleteMessages is used to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
func (api *API) DeleteMessages(payload *DeleteMessages) (bool, error) {
	return callJson[bool](api, "deleteMessages", payload)
}

// Sticker represents a sticker.
type Sticker struct {
	FileId           string        `json:"file_id"`                     // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId     string        `json:"file_unique_id"`              // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type             string        `json:"type"`                        // Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Width            int64         `json:"width"`                       // Sticker width
	Height           int64         `json:"height"`                      // Sticker height
	IsAnimated       bool          `json:"is_animated"`                 // True, if the sticker is animated
	IsVideo          bool          `json:"is_video"`                    // True, if the sticker is a video sticker
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`         // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji            string        `json:"emoji,omitempty"`             // Optional. Emoji associated with the sticker
	SetName          string        `json:"set_name,omitempty"`          // Optional. Name of the sticker set to which the sticker belongs
	PremiumAnimation *File         `json:"premium_animation,omitempty"` // Optional. For premium regular stickers, premium animation for the sticker
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`     // Optional. For mask stickers, the position where the mask should be placed
	CustomEmojiId    string        `json:"custom_emoji_id,omitempty"`   // Optional. For custom emoji stickers, unique identifier of the custom emoji
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`  // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	FileSize         int64         `json:"file_size,omitempty"`         // Optional. File size in bytes
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name        string     `json:"name"`                // Sticker set name
	Title       string     `json:"title"`               // Sticker set title
	StickerType string     `json:"sticker_type"`        // Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	Stickers    []*Sticker `json:"stickers"`            // List of all set stickers
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"` // Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}

// InputSticker describes a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      *InputFile    `json:"sticker"`                 // The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one using multipart/form-data, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
	Format       string        `json:"format"`                  // Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a WEBM video
	EmojiList    []string      `json:"emoji_list"`              // List of 1-20 emoji associated with the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	Keywords     []string      `json:"keywords,omitempty"`      // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
}

func (x *InputSticker) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker != nil {
		if x.Sticker.IsUploadable() {
			media[x.Sticker.Value] = x.Sticker
		}
	}

	return media
}

// sendSticker is used to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendSticker struct {
	BusinessConnectionId string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Sticker              *InputFile       `json:"sticker"`                          // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files ». Video and animated stickers can't be sent via an HTTP URL.
	Emoji                string           `json:"emoji,omitempty"`                  // Emoji associated with the sticker; only for just uploaded stickers
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (x *SendSticker) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker != nil {
		if x.Sticker.IsUploadable() {
			media["sticker"] = x.Sticker
		}
	}

	return media
}

func (x *SendSticker) getParams() (map[string]string, error) {
	payload := map[string]string{}

	if x.BusinessConnectionId != "" {
		payload["business_connection_id"] = x.BusinessConnectionId
	}
	if bb, err := json.Marshal(x.ChatId); err != nil {
		return nil, err
	} else {
		payload["chat_id"] = string(bb)
	}
	if x.MessageThreadId != 0 {
		payload["message_thread_id"] = strconv.FormatInt(x.MessageThreadId, 10)
	}
	if x.Emoji != "" {
		payload["emoji"] = x.Emoji
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.AllowPaidBroadcast {
		payload["allow_paid_broadcast"] = strconv.FormatBool(x.AllowPaidBroadcast)
	}
	if x.MessageEffectId != "" {
		payload["message_effect_id"] = x.MessageEffectId
	}
	if x.ReplyParameters != nil {
		if bb, err := json.Marshal(x.ReplyParameters); err != nil {
			return nil, err
		} else {
			payload["reply_parameters"] = string(bb)
		}
	}
	if x.ReplyMarkup != nil {
		if bb, err := json.Marshal(x.ReplyMarkup); err != nil {
			return nil, err
		} else {
			payload["reply_markup"] = string(bb)
		}
	}

	return payload, nil
}

// sendSticker is used to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
func (api *API) SendSticker(payload *SendSticker) (*Message, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*Message](api, "sendSticker", params, files)
	}
	return callJson[*Message](api, "sendSticker", payload)
}

// getStickerSet is used to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSet struct {
	Name string `json:"name"` // Name of the sticker set
}

// getStickerSet is used to get a sticker set. On success, a StickerSet object is returned.
func (api *API) GetStickerSet(payload *GetStickerSet) (*StickerSet, error) {
	return callJson[*StickerSet](api, "getStickerSet", payload)
}

// getCustomEmojiStickers is used to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
type GetCustomEmojiStickers struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"` // A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

// getCustomEmojiStickers is used to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (api *API) GetCustomEmojiStickers(payload *GetCustomEmojiStickers) ([]*Sticker, error) {
	return callJson[[]*Sticker](api, "getCustomEmojiStickers", payload)
}

// uploadStickerFile is used to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
	UserId        int64      `json:"user_id"`        // User identifier of sticker file owner
	Sticker       *InputFile `json:"sticker"`        // A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files »
	StickerFormat string     `json:"sticker_format"` // Format of the sticker, must be one of “static”, “animated”, “video”
}

func (x *UploadStickerFile) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker != nil {
		if x.Sticker.IsUploadable() {
			media["sticker"] = x.Sticker
		}
	}

	return media
}

func (x *UploadStickerFile) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["sticker_format"] = x.StickerFormat

	return payload, nil
}

// uploadStickerFile is used to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
func (api *API) UploadStickerFile(payload *UploadStickerFile) (*File, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return nil, err
		}
		return callMultipart[*File](api, "uploadStickerFile", params, files)
	}
	return callJson[*File](api, "uploadStickerFile", payload)
}

// createNewStickerSet is used to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
type CreateNewStickerSet struct {
	UserId          int64           `json:"user_id"`                    // User identifier of created sticker set owner
	Name            string          `json:"name"`                       // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Title           string          `json:"title"`                      // Sticker set title, 1-64 characters
	Stickers        []*InputSticker `json:"stickers"`                   // A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	StickerType     string          `json:"sticker_type,omitempty"`     // Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a regular sticker set is created.
	NeedsRepainting bool            `json:"needs_repainting,omitempty"` // Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
}

func (x *CreateNewStickerSet) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for _, m := range x.Stickers {
		for key, value := range m.getFiles() {
			media[key] = value
		}
	}

	return media
}

func (x *CreateNewStickerSet) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["name"] = x.Name
	payload["title"] = x.Title
	if bb, err := json.Marshal(x.Stickers); err != nil {
		return nil, err
	} else {
		payload["stickers"] = string(bb)
	}
	if x.StickerType != "" {
		payload["sticker_type"] = x.StickerType
	}
	if x.NeedsRepainting {
		payload["needs_repainting"] = strconv.FormatBool(x.NeedsRepainting)
	}

	return payload, nil
}

// createNewStickerSet is used to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
func (api *API) CreateNewStickerSet(payload *CreateNewStickerSet) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "createNewStickerSet", params, files)
	}
	return callJson[bool](api, "createNewStickerSet", payload)
}

// addStickerToSet is used to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
type AddStickerToSet struct {
	UserId  int64        `json:"user_id"` // User identifier of sticker set owner
	Name    string       `json:"name"`    // Sticker set name
	Sticker InputSticker `json:"sticker"` // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
}

func (x *AddStickerToSet) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for key, value := range x.Sticker.getFiles() {
		media[key] = value
	}

	return media
}

func (x *AddStickerToSet) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["name"] = x.Name
	if bb, err := json.Marshal(x.Sticker); err != nil {
		return nil, err
	} else {
		payload["sticker"] = string(bb)
	}

	return payload, nil
}

// addStickerToSet is used to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
func (api *API) AddStickerToSet(payload *AddStickerToSet) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "addStickerToSet", params, files)
	}
	return callJson[bool](api, "addStickerToSet", payload)
}

// setStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position. Returns True on success.
type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int64  `json:"position"` // New sticker position in the set, zero-based
}

// setStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position. Returns True on success.
func (api *API) SetStickerPositionInSet(payload *SetStickerPositionInSet) (bool, error) {
	return callJson[bool](api, "setStickerPositionInSet", payload)
}

// deleteStickerFromSet is used to delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

// deleteStickerFromSet is used to delete a sticker from a set created by the bot. Returns True on success.
func (api *API) DeleteStickerFromSet(payload *DeleteStickerFromSet) (bool, error) {
	return callJson[bool](api, "deleteStickerFromSet", payload)
}

// replaceStickerInSet is used to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
type ReplaceStickerInSet struct {
	UserId     int64        `json:"user_id"`     // User identifier of the sticker set owner
	Name       string       `json:"name"`        // Sticker set name
	OldSticker string       `json:"old_sticker"` // File identifier of the replaced sticker
	Sticker    InputSticker `json:"sticker"`     // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
}

func (x *ReplaceStickerInSet) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for key, value := range x.Sticker.getFiles() {
		media[key] = value
	}

	return media
}

func (x *ReplaceStickerInSet) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["name"] = x.Name
	payload["old_sticker"] = x.OldSticker
	if bb, err := json.Marshal(x.Sticker); err != nil {
		return nil, err
	} else {
		payload["sticker"] = string(bb)
	}

	return payload, nil
}

// replaceStickerInSet is used to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
func (api *API) ReplaceStickerInSet(payload *ReplaceStickerInSet) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "replaceStickerInSet", params, files)
	}
	return callJson[bool](api, "replaceStickerInSet", payload)
}

// setStickerEmojiList is used to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker"`    // File identifier of the sticker
	EmojiList []string `json:"emoji_list"` // A JSON-serialized list of 1-20 emoji associated with the sticker
}

// setStickerEmojiList is used to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (api *API) SetStickerEmojiList(payload *SetStickerEmojiList) (bool, error) {
	return callJson[bool](api, "setStickerEmojiList", payload)
}

// setStickerKeywords is used to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerKeywords struct {
	Sticker  string   `json:"sticker"`            // File identifier of the sticker
	Keywords []string `json:"keywords,omitempty"` // A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
}

// setStickerKeywords is used to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (api *API) SetStickerKeywords(payload *SetStickerKeywords) (bool, error) {
	return callJson[bool](api, "setStickerKeywords", payload)
}

// setStickerMaskPosition is used to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
type SetStickerMaskPosition struct {
	Sticker      string        `json:"sticker"`                 // File identifier of the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
}

// setStickerMaskPosition is used to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
func (api *API) SetStickerMaskPosition(payload *SetStickerMaskPosition) (bool, error) {
	return callJson[bool](api, "setStickerMaskPosition", payload)
}

// setStickerSetTitle is used to set the title of a created sticker set. Returns True on success.
type SetStickerSetTitle struct {
	Name  string `json:"name"`  // Sticker set name
	Title string `json:"title"` // Sticker set title, 1-64 characters
}

// setStickerSetTitle is used to set the title of a created sticker set. Returns True on success.
func (api *API) SetStickerSetTitle(payload *SetStickerSetTitle) (bool, error) {
	return callJson[bool](api, "setStickerSetTitle", payload)
}

// setStickerSetThumbnail is used to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
type SetStickerSetThumbnail struct {
	Name      string     `json:"name"`                // Sticker set name
	UserId    int64      `json:"user_id"`             // User identifier of the sticker set owner
	Thumbnail *InputFile `json:"thumbnail,omitempty"` // A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Format    string     `json:"format"`              // Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, or “video” for a WEBM video
}

func (x *SetStickerSetThumbnail) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

func (x *SetStickerSetThumbnail) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["name"] = x.Name
	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["format"] = x.Format

	return payload, nil
}

// setStickerSetThumbnail is used to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
func (api *API) SetStickerSetThumbnail(payload *SetStickerSetThumbnail) (bool, error) {
	if files := payload.getFiles(); len(files) != 0 {
		params, err := payload.getParams()
		if err != nil {
			return false, err
		}
		return callMultipart[bool](api, "setStickerSetThumbnail", params, files)
	}
	return callJson[bool](api, "setStickerSetThumbnail", payload)
}

// setCustomEmojiStickerSetThumbnail is used to set the thumbnail of a custom emoji sticker set. Returns True on success.
type SetCustomEmojiStickerSetThumbnail struct {
	Name          string `json:"name"`                      // Sticker set name
	CustomEmojiId string `json:"custom_emoji_id,omitempty"` // Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
}

// setCustomEmojiStickerSetThumbnail is used to set the thumbnail of a custom emoji sticker set. Returns True on success.
func (api *API) SetCustomEmojiStickerSetThumbnail(payload *SetCustomEmojiStickerSetThumbnail) (bool, error) {
	return callJson[bool](api, "setCustomEmojiStickerSetThumbnail", payload)
}

// deleteStickerSet is used to delete a sticker set that was created by the bot. Returns True on success.
type DeleteStickerSet struct {
	Name string `json:"name"` // Sticker set name
}

// deleteStickerSet is used to delete a sticker set that was created by the bot. Returns True on success.
func (api *API) DeleteStickerSet(payload *DeleteStickerSet) (bool, error) {
	return callJson[bool](api, "deleteStickerSet", payload)
}

// Gift represents a gift that can be sent by the bot.
type Gift struct {
	Id             string  `json:"id"`                        // Unique identifier of the gift
	Sticker        Sticker `json:"sticker"`                   // The sticker that represents the gift
	StarCount      int64   `json:"star_count"`                // The number of Telegram Stars that must be paid to send the sticker
	TotalCount     int64   `json:"total_count,omitempty"`     // Optional. The total number of the gifts of this type that can be sent; for limited gifts only
	RemainingCount int64   `json:"remaining_count,omitempty"` // Optional. The number of remaining gifts of this type that can be sent; for limited gifts only
}

// Gifts represent a list of gifts.
type Gifts struct {
	Gifts []*Gift `json:"gifts"` // The list of gifts
}

// Returns the list of gifts that can be sent by the bot to users. Requires no parameters. Returns a Gifts object.
func (api *API) GetAvailableGifts() (*Gifts, error) {
	return callJson[*Gifts](api, "getAvailableGifts", nil)
}

// Sends a gift to the given user. The gift can't be converted to Telegram Stars by the user. Returns True on success.
type SendGift struct {
	UserId        int64            `json:"user_id"`                   // Unique identifier of the target user that will receive the gift
	GiftId        string           `json:"gift_id"`                   // Identifier of the gift
	Text          string           `json:"text,omitempty"`            // Text that will be shown along with the gift; 0-255 characters
	TextParseMode string           `json:"text_parse_mode,omitempty"` // Mode for parsing entities in the text. See formatting options for more details. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`   // A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
}

// Sends a gift to the given user. The gift can't be converted to Telegram Stars by the user. Returns True on success.
func (api *API) SendGift(payload *SendGift) (bool, error) {
	return callJson[bool](api, "sendGift", payload)
}

// InlineQuery represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	Id       string    `json:"id"`                  // Unique identifier for this query
	From     User      `json:"from"`                // Sender
	Query    string    `json:"query"`               // Text of the query (up to 256 characters)
	Offset   string    `json:"offset"`              // Offset of the results to be returned, can be controlled by the bot
	ChatType string    `json:"chat_type,omitempty"` // Optional. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	Location *Location `json:"location,omitempty"`  // Optional. Sender location, only for bots that request user location
}

// answerInlineQuery is used to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
	InlineQueryId string                    `json:"inline_query_id"`       // Unique identifier for the answered query
	Results       []InlineQueryResult       `json:"results"`               // A JSON-serialized array of results for the inline query
	CacheTime     int64                     `json:"cache_time,omitempty"`  // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal    bool                      `json:"is_personal,omitempty"` // Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	NextOffset    string                    `json:"next_offset,omitempty"` // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	Button        *InlineQueryResultsButton `json:"button,omitempty"`      // A JSON-serialized object describing a button to be shown above inline query results
}

// answerInlineQuery is used to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
func (api *API) AnswerInlineQuery(payload *AnswerInlineQuery) (bool, error) {
	return callJson[bool](api, "answerInlineQuery", payload)
}

// InlineQueryResultsButton represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`                      // Label text on the button
	WebApp         *WebAppInfo `json:"web_app,omitempty"`         // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
	StartParameter string      `json:"start_parameter,omitempty"` // Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

// InlineQueryResult represents one result of an inline query. Telegram clients currently support results of the following 20 types:
// InlineQueryResultCachedAudio, InlineQueryResultCachedDocument, InlineQueryResultCachedGif, InlineQueryResultCachedMpeg4Gif, InlineQueryResultCachedPhoto, InlineQueryResultCachedSticker, InlineQueryResultCachedVideo, InlineQueryResultCachedVoice, InlineQueryResultArticle, InlineQueryResultAudio, InlineQueryResultContact, InlineQueryResultGame, InlineQueryResultDocument, InlineQueryResultGif, InlineQueryResultLocation, InlineQueryResultMpeg4Gif, InlineQueryResultPhoto, InlineQueryResultVenue, InlineQueryResultVideo, InlineQueryResultVoice
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult interface {
	// IsInlineQueryResult does nothing and is only used to enforce type-safety
	IsInlineQueryResult()
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                       // Type of the result, must be article
	Id                  string                `json:"id"`                         // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                      // Title of the result
	InputMessageContent InputMessageContent   `json:"input_message_content"`      // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`     // Optional. Inline keyboard attached to the message
	Url                 string                `json:"url,omitempty"`              // Optional. URL of the result
	HideUrl             bool                  `json:"hide_url,omitempty"`         // Optional. Pass True if you don't want the URL to be shown in the message
	Description         string                `json:"description,omitempty"`      // Optional. Short description of the result
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`    // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`  // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"` // Optional. Thumbnail height
}

func (InlineQueryResultArticle) IsInlineQueryResult() {}

func (x *InlineQueryResultArticle) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                       // Type of the result, must be article
		Id                  string                `json:"id"`                         // Unique identifier for this result, 1-64 Bytes
		Title               string                `json:"title"`                      // Title of the result
		InputMessageContent json.RawMessage       `json:"input_message_content"`      // Content of the message to be sent
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`     // Optional. Inline keyboard attached to the message
		Url                 string                `json:"url,omitempty"`              // Optional. URL of the result
		HideUrl             bool                  `json:"hide_url,omitempty"`         // Optional. Pass True if you don't want the URL to be shown in the message
		Description         string                `json:"description,omitempty"`      // Optional. Short description of the result
		ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`    // Optional. Url of the thumbnail for the result
		ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`  // Optional. Thumbnail width
		ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"` // Optional. Thumbnail height
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Title = raw.Title

	x.ReplyMarkup = raw.ReplyMarkup
	x.Url = raw.Url
	x.HideUrl = raw.HideUrl
	x.Description = raw.Description
	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailWidth = raw.ThumbnailWidth
	x.ThumbnailHeight = raw.ThumbnailHeight
	return nil
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`                               // Type of the result, must be photo
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	PhotoUrl              string                `json:"photo_url"`                          // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the thumbnail for the photo
	PhotoWidth            int64                 `json:"photo_width,omitempty"`              // Optional. Width of the photo
	PhotoHeight           int64                 `json:"photo_height,omitempty"`             // Optional. Height of the photo
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
}

func (InlineQueryResultPhoto) IsInlineQueryResult() {}

func (x *InlineQueryResultPhoto) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be photo
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		PhotoUrl              string                `json:"photo_url"`                          // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
		ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the thumbnail for the photo
		PhotoWidth            int64                 `json:"photo_width,omitempty"`              // Optional. Width of the photo
		PhotoHeight           int64                 `json:"photo_height,omitempty"`             // Optional. Height of the photo
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.PhotoUrl = raw.PhotoUrl
	x.ThumbnailUrl = raw.ThumbnailUrl
	x.PhotoWidth = raw.PhotoWidth
	x.PhotoHeight = raw.PhotoHeight
	x.Title = raw.Title
	x.Description = raw.Description
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                  string                `json:"type"`                               // Type of the result, must be gif
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	GifUrl                string                `json:"gif_url"`                            // A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth              int64                 `json:"gif_width,omitempty"`                // Optional. Width of the GIF
	GifHeight             int64                 `json:"gif_height,omitempty"`               // Optional. Height of the GIF
	GifDuration           int64                 `json:"gif_duration,omitempty"`             // Optional. Duration of the GIF in seconds
	ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`      // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the GIF animation
}

func (InlineQueryResultGif) IsInlineQueryResult() {}

func (x *InlineQueryResultGif) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be gif
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		GifUrl                string                `json:"gif_url"`                            // A valid URL for the GIF file. File size must not exceed 1MB
		GifWidth              int64                 `json:"gif_width,omitempty"`                // Optional. Width of the GIF
		GifHeight             int64                 `json:"gif_height,omitempty"`               // Optional. Height of the GIF
		GifDuration           int64                 `json:"gif_duration,omitempty"`             // Optional. Duration of the GIF in seconds
		ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
		ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`      // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the GIF animation
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.GifUrl = raw.GifUrl
	x.GifWidth = raw.GifWidth
	x.GifHeight = raw.GifHeight
	x.GifDuration = raw.GifDuration
	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailMimeType = raw.ThumbnailMimeType
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`                               // Type of the result, must be mpeg4_gif
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	Mpeg4Url              string                `json:"mpeg4_url"`                          // A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4Width            int64                 `json:"mpeg4_width,omitempty"`              // Optional. Video width
	Mpeg4Height           int64                 `json:"mpeg4_height,omitempty"`             // Optional. Video height
	Mpeg4Duration         int64                 `json:"mpeg4_duration,omitempty"`           // Optional. Video duration in seconds
	ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`      // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video animation
}

func (InlineQueryResultMpeg4Gif) IsInlineQueryResult() {}

func (x *InlineQueryResultMpeg4Gif) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be mpeg4_gif
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		Mpeg4Url              string                `json:"mpeg4_url"`                          // A valid URL for the MPEG4 file. File size must not exceed 1MB
		Mpeg4Width            int64                 `json:"mpeg4_width,omitempty"`              // Optional. Video width
		Mpeg4Height           int64                 `json:"mpeg4_height,omitempty"`             // Optional. Video height
		Mpeg4Duration         int64                 `json:"mpeg4_duration,omitempty"`           // Optional. Video duration in seconds
		ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
		ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`      // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video animation
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Mpeg4Url = raw.Mpeg4Url
	x.Mpeg4Width = raw.Mpeg4Width
	x.Mpeg4Height = raw.Mpeg4Height
	x.Mpeg4Duration = raw.Mpeg4Duration
	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailMimeType = raw.ThumbnailMimeType
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`                               // Type of the result, must be video
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	VideoUrl              string                `json:"video_url"`                          // A valid URL for the embedded video player or video file
	MimeType              string                `json:"mime_type"`                          // MIME type of the content of the video URL, “text/html” or “video/mp4”
	ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the thumbnail (JPEG only) for the video
	Title                 string                `json:"title"`                              // Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	VideoWidth            int64                 `json:"video_width,omitempty"`              // Optional. Video width
	VideoHeight           int64                 `json:"video_height,omitempty"`             // Optional. Video height
	VideoDuration         int64                 `json:"video_duration,omitempty"`           // Optional. Video duration in seconds
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}

func (InlineQueryResultVideo) IsInlineQueryResult() {}

func (x *InlineQueryResultVideo) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be video
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		VideoUrl              string                `json:"video_url"`                          // A valid URL for the embedded video player or video file
		MimeType              string                `json:"mime_type"`                          // MIME type of the content of the video URL, “text/html” or “video/mp4”
		ThumbnailUrl          string                `json:"thumbnail_url"`                      // URL of the thumbnail (JPEG only) for the video
		Title                 string                `json:"title"`                              // Title for the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		VideoWidth            int64                 `json:"video_width,omitempty"`              // Optional. Video width
		VideoHeight           int64                 `json:"video_height,omitempty"`             // Optional. Video height
		VideoDuration         int64                 `json:"video_duration,omitempty"`           // Optional. Video duration in seconds
		Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.VideoUrl = raw.VideoUrl
	x.MimeType = raw.MimeType
	x.ThumbnailUrl = raw.ThumbnailUrl
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.VideoWidth = raw.VideoWidth
	x.VideoHeight = raw.VideoHeight
	x.VideoDuration = raw.VideoDuration
	x.Description = raw.Description
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioUrl            string                `json:"audio_url"`                       // A valid URL for the audio file
	Title               string                `json:"title"`                           // Title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Performer           string                `json:"performer,omitempty"`             // Optional. Performer
	AudioDuration       int64                 `json:"audio_duration,omitempty"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

func (InlineQueryResultAudio) IsInlineQueryResult() {}

func (x *InlineQueryResultAudio) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be audio
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		AudioUrl            string                `json:"audio_url"`                       // A valid URL for the audio file
		Title               string                `json:"title"`                           // Title
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		Performer           string                `json:"performer,omitempty"`             // Optional. Performer
		AudioDuration       int64                 `json:"audio_duration,omitempty"`        // Optional. Audio duration in seconds
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.AudioUrl = raw.AudioUrl
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.Performer = raw.Performer
	x.AudioDuration = raw.AudioDuration
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`                            // Type of the result, must be voice
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceUrl            string                `json:"voice_url"`                       // A valid URL for the voice recording
	Title               string                `json:"title"`                           // Recording title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	VoiceDuration       int64                 `json:"voice_duration,omitempty"`        // Optional. Recording duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice recording
}

func (InlineQueryResultVoice) IsInlineQueryResult() {}

func (x *InlineQueryResultVoice) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be voice
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		VoiceUrl            string                `json:"voice_url"`                       // A valid URL for the voice recording
		Title               string                `json:"title"`                           // Recording title
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		VoiceDuration       int64                 `json:"voice_duration,omitempty"`        // Optional. Recording duration in seconds
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice recording
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.VoiceUrl = raw.VoiceUrl
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.VoiceDuration = raw.VoiceDuration
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                            // Type of the result, must be document
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DocumentUrl         string                `json:"document_url"`                    // A valid URL for the file
	MimeType            string                `json:"mime_type"`                       // MIME type of the content of the file, either “application/pdf” or “application/zip”
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultDocument) IsInlineQueryResult() {}

func (x *InlineQueryResultDocument) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be document
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		Title               string                `json:"title"`                           // Title for the result
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		DocumentUrl         string                `json:"document_url"`                    // A valid URL for the file
		MimeType            string                `json:"mime_type"`                       // MIME type of the content of the file, either “application/pdf” or “application/zip”
		Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
		ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. URL of the thumbnail (JPEG only) for the file
		ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
		ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.DocumentUrl = raw.DocumentUrl
	x.MimeType = raw.MimeType
	x.Description = raw.Description
	x.ReplyMarkup = raw.ReplyMarkup

	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailWidth = raw.ThumbnailWidth
	x.ThumbnailHeight = raw.ThumbnailHeight
	return nil
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`                             // Type of the result, must be location
	Id                   string                `json:"id"`                               // Unique identifier for this result, 1-64 Bytes
	Latitude             float64               `json:"latitude"`                         // Location latitude in degrees
	Longitude            float64               `json:"longitude"`                        // Location longitude in degrees
	Title                string                `json:"title"`                            // Location title
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64                 `json:"live_period,omitempty"`            // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64                 `json:"heading,omitempty"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional. Inline keyboard attached to the message
	InputMessageContent  InputMessageContent   `json:"input_message_content,omitempty"`  // Optional. Content of the message to be sent instead of the location
	ThumbnailUrl         string                `json:"thumbnail_url,omitempty"`          // Optional. Url of the thumbnail for the result
	ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`        // Optional. Thumbnail width
	ThumbnailHeight      int64                 `json:"thumbnail_height,omitempty"`       // Optional. Thumbnail height
}

func (InlineQueryResultLocation) IsInlineQueryResult() {}

func (x *InlineQueryResultLocation) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                 string                `json:"type"`                             // Type of the result, must be location
		Id                   string                `json:"id"`                               // Unique identifier for this result, 1-64 Bytes
		Latitude             float64               `json:"latitude"`                         // Location latitude in degrees
		Longitude            float64               `json:"longitude"`                        // Location longitude in degrees
		Title                string                `json:"title"`                            // Location title
		HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
		LivePeriod           int64                 `json:"live_period,omitempty"`            // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
		Heading              int64                 `json:"heading,omitempty"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
		ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
		ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional. Inline keyboard attached to the message
		InputMessageContent  json.RawMessage       `json:"input_message_content,omitempty"`  // Optional. Content of the message to be sent instead of the location
		ThumbnailUrl         string                `json:"thumbnail_url,omitempty"`          // Optional. Url of the thumbnail for the result
		ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`        // Optional. Thumbnail width
		ThumbnailHeight      int64                 `json:"thumbnail_height,omitempty"`       // Optional. Thumbnail height
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Latitude = raw.Latitude
	x.Longitude = raw.Longitude
	x.Title = raw.Title
	x.HorizontalAccuracy = raw.HorizontalAccuracy
	x.LivePeriod = raw.LivePeriod
	x.Heading = raw.Heading
	x.ProximityAlertRadius = raw.ProximityAlertRadius
	x.ReplyMarkup = raw.ReplyMarkup

	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailWidth = raw.ThumbnailWidth
	x.ThumbnailHeight = raw.ThumbnailHeight
	return nil
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                            // Type of the result, must be venue
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`                        // Latitude of the venue location in degrees
	Longitude           float64               `json:"longitude"`                       // Longitude of the venue location in degrees
	Title               string                `json:"title"`                           // Title of the venue
	Address             string                `json:"address"`                         // Address of the venue
	FoursquareId        string                `json:"foursquare_id,omitempty"`         // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type,omitempty"`       // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId       string                `json:"google_place_id,omitempty"`       // Optional. Google Places identifier of the venue
	GooglePlaceType     string                `json:"google_place_type,omitempty"`     // Optional. Google Places type of the venue. (See supported types.)
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the venue
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultVenue) IsInlineQueryResult() {}

func (x *InlineQueryResultVenue) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be venue
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
		Latitude            float64               `json:"latitude"`                        // Latitude of the venue location in degrees
		Longitude           float64               `json:"longitude"`                       // Longitude of the venue location in degrees
		Title               string                `json:"title"`                           // Title of the venue
		Address             string                `json:"address"`                         // Address of the venue
		FoursquareId        string                `json:"foursquare_id,omitempty"`         // Optional. Foursquare identifier of the venue if known
		FoursquareType      string                `json:"foursquare_type,omitempty"`       // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
		GooglePlaceId       string                `json:"google_place_id,omitempty"`       // Optional. Google Places identifier of the venue
		GooglePlaceType     string                `json:"google_place_type,omitempty"`     // Optional. Google Places type of the venue. (See supported types.)
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the venue
		ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
		ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
		ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Latitude = raw.Latitude
	x.Longitude = raw.Longitude
	x.Title = raw.Title
	x.Address = raw.Address
	x.FoursquareId = raw.FoursquareId
	x.FoursquareType = raw.FoursquareType
	x.GooglePlaceId = raw.GooglePlaceId
	x.GooglePlaceType = raw.GooglePlaceType
	x.ReplyMarkup = raw.ReplyMarkup

	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailWidth = raw.ThumbnailWidth
	x.ThumbnailHeight = raw.ThumbnailHeight
	return nil
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                            // Type of the result, must be contact
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`                    // Contact's phone number
	FirstName           string                `json:"first_name"`                      // Contact's first name
	LastName            string                `json:"last_name,omitempty"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard,omitempty"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the contact
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultContact) IsInlineQueryResult() {}

func (x *InlineQueryResultContact) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be contact
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
		PhoneNumber         string                `json:"phone_number"`                    // Contact's phone number
		FirstName           string                `json:"first_name"`                      // Contact's first name
		LastName            string                `json:"last_name,omitempty"`             // Optional. Contact's last name
		Vcard               string                `json:"vcard,omitempty"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the contact
		ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
		ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
		ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.PhoneNumber = raw.PhoneNumber
	x.FirstName = raw.FirstName
	x.LastName = raw.LastName
	x.Vcard = raw.Vcard
	x.ReplyMarkup = raw.ReplyMarkup

	x.ThumbnailUrl = raw.ThumbnailUrl
	x.ThumbnailWidth = raw.ThumbnailWidth
	x.ThumbnailHeight = raw.ThumbnailHeight
	return nil
}

// Represents a Game.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`                   // Type of the result, must be game
	Id            string                `json:"id"`                     // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"`        // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
}

func (InlineQueryResultGame) IsInlineQueryResult() {}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`                               // Type of the result, must be photo
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	PhotoFileId           string                `json:"photo_file_id"`                      // A valid file identifier of the photo
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
}

func (InlineQueryResultCachedPhoto) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedPhoto) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be photo
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		PhotoFileId           string                `json:"photo_file_id"`                      // A valid file identifier of the photo
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.PhotoFileId = raw.PhotoFileId
	x.Title = raw.Title
	x.Description = raw.Description
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`                               // Type of the result, must be gif
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	GifFileId             string                `json:"gif_file_id"`                        // A valid file identifier for the GIF file
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the GIF animation
}

func (InlineQueryResultCachedGif) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedGif) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be gif
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		GifFileId             string                `json:"gif_file_id"`                        // A valid file identifier for the GIF file
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the GIF animation
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.GifFileId = raw.GifFileId
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`                               // Type of the result, must be mpeg4_gif
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	Mpeg4FileId           string                `json:"mpeg4_file_id"`                      // A valid file identifier for the MPEG4 file
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video animation
}

func (InlineQueryResultCachedMpeg4Gif) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedMpeg4Gif) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be mpeg4_gif
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		Mpeg4FileId           string                `json:"mpeg4_file_id"`                      // A valid file identifier for the MPEG4 file
		Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video animation
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Mpeg4FileId = raw.Mpeg4FileId
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                            // Type of the result, must be sticker
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	StickerFileId       string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
}

func (InlineQueryResultCachedSticker) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedSticker) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be sticker
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		StickerFileId       string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.StickerFileId = raw.StickerFileId
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`                            // Type of the result, must be document
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // Title for the result
	DocumentFileId      string                `json:"document_file_id"`                // A valid file identifier for the file
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
}

func (InlineQueryResultCachedDocument) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedDocument) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be document
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		Title               string                `json:"title"`                           // Title for the result
		DocumentFileId      string                `json:"document_file_id"`                // A valid file identifier for the file
		Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.Title = raw.Title
	x.DocumentFileId = raw.DocumentFileId
	x.Description = raw.Description
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`                               // Type of the result, must be video
	Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	VideoFileId           string                `json:"video_file_id"`                      // A valid file identifier for the video file
	Title                 string                `json:"title"`                              // Title for the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video
}

func (InlineQueryResultCachedVideo) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedVideo) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                  string                `json:"type"`                               // Type of the result, must be video
		Id                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
		VideoFileId           string                `json:"video_file_id"`                      // A valid file identifier for the video file
		Title                 string                `json:"title"`                              // Title for the result
		Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
		Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
		ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
		CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
		InputMessageContent   json.RawMessage       `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.VideoFileId = raw.VideoFileId
	x.Title = raw.Title
	x.Description = raw.Description
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ShowCaptionAboveMedia = raw.ShowCaptionAboveMedia
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                            // Type of the result, must be voice
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceFileId         string                `json:"voice_file_id"`                   // A valid file identifier for the voice message
	Title               string                `json:"title"`                           // Voice message title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice message
}

func (InlineQueryResultCachedVoice) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedVoice) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be voice
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		VoiceFileId         string                `json:"voice_file_id"`                   // A valid file identifier for the voice message
		Title               string                `json:"title"`                           // Voice message title
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice message
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.VoiceFileId = raw.VoiceFileId
	x.Title = raw.Title
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioFileId         string                `json:"audio_file_id"`                   // A valid file identifier for the audio file
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

func (InlineQueryResultCachedAudio) IsInlineQueryResult() {}

func (x *InlineQueryResultCachedAudio) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type                string                `json:"type"`                            // Type of the result, must be audio
		Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
		AudioFileId         string                `json:"audio_file_id"`                   // A valid file identifier for the audio file
		Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
		ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
		CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
		InputMessageContent json.RawMessage       `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalInputMessageContent(raw.InputMessageContent); err != nil {
		return err
	} else {
		x.InputMessageContent = data
	}
	x.Type = raw.Type
	x.Id = raw.Id
	x.AudioFileId = raw.AudioFileId
	x.Caption = raw.Caption
	x.ParseMode = raw.ParseMode
	x.CaptionEntities = raw.CaptionEntities
	x.ReplyMarkup = raw.ReplyMarkup

	return nil
}

// InputMessageContent represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:
// InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent
type InputMessageContent interface {
	// IsInputMessageContent does nothing and is only used to enforce type-safety
	IsInputMessageContent()
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`                   // Text of the message to be sent, 1-4096 characters
	ParseMode          ParseMode           `json:"parse_mode,omitempty"`           // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities           []*MessageEntity    `json:"entities,omitempty"`             // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Link preview generation options for the message
}

func (InputTextMessageContent) IsInputMessageContent() {}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`                         // Latitude of the location in degrees
	Longitude            float64 `json:"longitude"`                        // Longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64   `json:"live_period,omitempty"`            // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64   `json:"heading,omitempty"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
}

func (InputLocationMessageContent) IsInputMessageContent() {}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`                    // Latitude of the venue in degrees
	Longitude       float64 `json:"longitude"`                   // Longitude of the venue in degrees
	Title           string  `json:"title"`                       // Name of the venue
	Address         string  `json:"address"`                     // Address of the venue
	FoursquareId    string  `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue, if known
	FoursquareType  string  `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId   string  `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string  `json:"google_place_type,omitempty"` // Optional. Google Places type of the venue. (See supported types.)
}

func (InputVenueMessageContent) IsInputMessageContent() {}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	Vcard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}

func (InputContactMessageContent) IsInputMessageContent() {}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	Title                     string          `json:"title"`                                   // Product name, 1-32 characters
	Description               string          `json:"description"`                             // Product description, 1-255 characters
	Payload                   string          `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string          `json:"provider_token,omitempty"`                // Optional. Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string          `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Prices                    []*LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount              int64           `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64         `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data,omitempty"`                 // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64           `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int64           `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int64           `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool            `json:"need_name,omitempty"`                     // Optional. Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool            `json:"need_phone_number,omitempty"`             // Optional. Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool            `json:"need_email,omitempty"`                    // Optional. Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool            `json:"need_shipping_address,omitempty"`         // Optional. Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool            `json:"send_email_to_provider,omitempty"`        // Optional. Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool            `json:"is_flexible,omitempty"`                   // Optional. Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
}

func (InputInvoiceMessageContent) IsInputMessageContent() {}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`                   // The unique identifier for the result that was chosen
	From            User      `json:"from"`                        // The user that chose the result
	Location        *Location `json:"location,omitempty"`          // Optional. Sender location, only for bots that require user location
	InlineMessageId string    `json:"inline_message_id,omitempty"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`                       // The query that was used to obtain the result
}

// answerWebAppQuery is used to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
type AnswerWebAppQuery struct {
	WebAppQueryId string            `json:"web_app_query_id"` // Unique identifier for the query to be answered
	Result        InlineQueryResult `json:"result"`           // A JSON-serialized object describing the message to be sent
}

// answerWebAppQuery is used to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
func (api *API) AnswerWebAppQuery(payload *AnswerWebAppQuery) (*SentWebAppMessage, error) {
	return callJson[*SentWebAppMessage](api, "answerWebAppQuery", payload)
}

// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id,omitempty"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
type SavePreparedInlineMessage struct {
	UserId            int64             `json:"user_id"`                       // Unique identifier of the target user that can use the prepared message
	Result            InlineQueryResult `json:"result"`                        // A JSON-serialized object describing the message to be sent
	AllowUserChats    bool              `json:"allow_user_chats,omitempty"`    // Pass True if the message can be sent to private chats with users
	AllowBotChats     bool              `json:"allow_bot_chats,omitempty"`     // Pass True if the message can be sent to private chats with bots
	AllowGroupChats   bool              `json:"allow_group_chats,omitempty"`   // Pass True if the message can be sent to group and supergroup chats
	AllowChannelChats bool              `json:"allow_channel_chats,omitempty"` // Pass True if the message can be sent to channel chats
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
func (api *API) SavePreparedInlineMessage(payload *SavePreparedInlineMessage) (*PreparedInlineMessage, error) {
	return callJson[*PreparedInlineMessage](api, "savePreparedInlineMessage", payload)
}

// Describes an inline message to be sent by a user of a Mini App.
type PreparedInlineMessage struct {
	Id             string `json:"id"`              // Unique identifier of the prepared message
	ExpirationDate int64  `json:"expiration_date"` // Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used
}

// sendInvoice is used to send invoices. On success, the sent Message is returned.
type SendInvoice struct {
	ChatId                    ChatID                `json:"chat_id"`                                 // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId           int64                 `json:"message_thread_id,omitempty"`             // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Title                     string                `json:"title"`                                   // Product name, 1-32 characters
	Description               string                `json:"description"`                             // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string                `json:"provider_token,omitempty"`                // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string                `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Prices                    []*LabeledPrice       `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount              int64                 `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64               `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                `json:"start_parameter,omitempty"`               // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int64                 `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int64                 `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int64                 `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool                  `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool                  `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool                  `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	DisableNotification       bool                  `json:"disable_notification,omitempty"`          // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent            bool                  `json:"protect_content,omitempty"`               // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast        bool                  `json:"allow_paid_broadcast,omitempty"`          // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId           string                `json:"message_effect_id,omitempty"`             // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters           *ReplyParameters      `json:"reply_parameters,omitempty"`              // Description of the message to reply to
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

// sendInvoice is used to send invoices. On success, the sent Message is returned.
func (api *API) SendInvoice(payload *SendInvoice) (*Message, error) {
	return callJson[*Message](api, "sendInvoice", payload)
}

// createInvoiceLink is used to create a link for an invoice. Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
	BusinessConnectionId      string          `json:"business_connection_id,omitempty"`        // Unique identifier of the business connection on behalf of which the link will be created. For payments in Telegram Stars only.
	Title                     string          `json:"title"`                                   // Product name, 1-32 characters
	Description               string          `json:"description"`                             // Product description, 1-255 characters
	Payload                   string          `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string          `json:"provider_token,omitempty"`                // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string          `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Prices                    []*LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	SubscriptionPeriod        int64           `json:"subscription_period,omitempty"`           // The number of seconds the subscription will be active for before the next payment. The currency must be set to “XTR” (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 2500 Telegram Stars.
	MaxTipAmount              int64           `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64         `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64           `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int64           `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int64           `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool            `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool            `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool            `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool            `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool            `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool            `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
}

// createInvoiceLink is used to create a link for an invoice. Returns the created invoice link as String on success.
func (api *API) CreateInvoiceLink(payload *CreateInvoiceLink) (string, error) {
	return callJson[string](api, "createInvoiceLink", payload)
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. answerShippingQuery is used to reply to shipping queries. On success, True is returned.
type AnswerShippingQuery struct {
	ShippingQueryId string            `json:"shipping_query_id"`          // Unique identifier for the query to be answered
	Ok              bool              `json:"ok"`                         // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"` // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string            `json:"error_message,omitempty"`    // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. answerShippingQuery is used to reply to shipping queries. On success, True is returned.
func (api *API) AnswerShippingQuery(payload *AnswerShippingQuery) (bool, error) {
	return callJson[bool](api, "answerShippingQuery", payload)
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. answerPreCheckoutQuery is used to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`   // Unique identifier for the query to be answered
	Ok                 bool   `json:"ok"`                      // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string `json:"error_message,omitempty"` // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. answerPreCheckoutQuery is used to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (api *API) AnswerPreCheckoutQuery(payload *AnswerPreCheckoutQuery) (bool, error) {
	return callJson[bool](api, "answerPreCheckoutQuery", payload)
}

// Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
type GetStarTransactions struct {
	Offset int64 `json:"offset,omitempty"` // Number of transactions to skip in the response
	Limit  int64 `json:"limit,omitempty"`  // The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

// Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
func (api *API) GetStarTransactions(payload *GetStarTransactions) (*StarTransactions, error) {
	return callJson[*StarTransactions](api, "getStarTransactions", payload)
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
type RefundStarPayment struct {
	UserId                  int64  `json:"user_id"`                    // Identifier of the user whose payment will be refunded
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"` // Telegram payment identifier
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
func (api *API) RefundStarPayment(payload *RefundStarPayment) (bool, error) {
	return callJson[bool](api, "refundStarPayment", payload)
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
type EditUserStarSubscription struct {
	UserId                  int64  `json:"user_id"`                    // Identifier of the user whose subscription will be edited
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"` // Telegram payment identifier for the subscription
	IsCanceled              bool   `json:"is_canceled"`                // Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
func (api *API) EditUserStarSubscription(payload *EditUserStarSubscription) (bool, error) {
	return callJson[bool](api, "editUserStarSubscription", payload)
}

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int64  `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount    int64  `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number,omitempty"`     // Optional. User's phone number
	Email           string           `json:"email,omitempty"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` // Optional. User shipping address
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	Id     string          `json:"id"`     // Shipping option identifier
	Title  string          `json:"title"`  // Option title
	Prices []*LabeledPrice `json:"prices"` // List of price portions
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`                               // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount                int64      `json:"total_amount"`                           // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload             string     `json:"invoice_payload"`                        // Bot-specified invoice payload
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date,omitempty"` // Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	IsRecurring                bool       `json:"is_recurring,omitempty"`                 // Optional. True, if the payment is a recurring payment for a subscription
	IsFirstRecurring           bool       `json:"is_first_recurring,omitempty"`           // Optional. True, if the payment is the first payment for a subscription
	ShippingOptionId           string     `json:"shipping_option_id,omitempty"`           // Optional. Identifier of the shipping option chosen by the user
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`                   // Optional. Order information provided by the user
	TelegramPaymentChargeId    string     `json:"telegram_payment_charge_id"`             // Telegram payment identifier
	ProviderPaymentChargeId    string     `json:"provider_payment_charge_id"`             // Provider payment identifier
}

// RefundedPayment contains basic information about a refunded payment.
type RefundedPayment struct {
	Currency                string `json:"currency"`                             // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars. Currently, always “XTR”
	TotalAmount             int64  `json:"total_amount"`                         // Total refunded price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string `json:"invoice_payload"`                      // Bot-specified invoice payload
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`           // Telegram payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id,omitempty"` // Optional. Provider payment identifier
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	Id              string          `json:"id"`               // Unique query identifier
	From            User            `json:"from"`             // User who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Bot-specified invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User specified shipping address
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Id               string     `json:"id"`                           // Unique query identifier
	From             User       `json:"from"`                         // User who sent the query
	Currency         string     `json:"currency"`                     // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount      int64      `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`              // Bot-specified invoice payload
	ShippingOptionId string     `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
}

// PaidMediaPurchased contains information about a paid media purchase.
type PaidMediaPurchased struct {
	From             User   `json:"from"`               // User who purchased the media
	PaidMediaPayload string `json:"paid_media_payload"` // Bot-specified paid media payload
}

// RevenueWithdrawalState describes the state of a revenue withdrawal operation. Currently, it can be one of
// RevenueWithdrawalStatePending, RevenueWithdrawalStateSucceeded, RevenueWithdrawalStateFailed
type RevenueWithdrawalState interface {
	// IsRevenueWithdrawalState does nothing and is only used to enforce type-safety
	IsRevenueWithdrawalState()
}

// The withdrawal is in progress.
type RevenueWithdrawalStatePending struct {
	Type string `json:"type"` // Type of the state, always “pending”
}

func (RevenueWithdrawalStatePending) IsRevenueWithdrawalState() {}

// The withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"` // Type of the state, always “succeeded”
	Date int64  `json:"date"` // Date the withdrawal was completed in Unix time
	Url  string `json:"url"`  // An HTTPS URL that can be used to see transaction details
}

func (RevenueWithdrawalStateSucceeded) IsRevenueWithdrawalState() {}

// The withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"` // Type of the state, always “failed”
}

func (RevenueWithdrawalStateFailed) IsRevenueWithdrawalState() {}

// Contains information about the affiliate that received a commission via this transaction.
type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`  // Optional. The bot or the user that received an affiliate commission if it was received by a bot or a user
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`  // Optional. The chat that received an affiliate commission if it was received by a chat
	CommissionPerMille int64 `json:"commission_per_mille"`      // The number of Telegram Stars received by the affiliate for each 1000 Telegram Stars received by the bot from referred users
	Amount             int64 `json:"amount"`                    // Integer amount of Telegram Stars received by the affiliate from the transaction, rounded to 0; can be negative for refunds
	NanostarAmount     int64 `json:"nanostar_amount,omitempty"` // Optional. The number of 1/1000000000 shares of Telegram Stars received by the affiliate; from -999999999 to 999999999; can be negative for refunds
}

// TransactionPartner describes the source of a transaction, or its recipient for outgoing transactions. Currently, it can be one of
// TransactionPartnerUser, TransactionPartnerAffiliateProgram, TransactionPartnerFragment, TransactionPartnerTelegramAds, TransactionPartnerTelegramApi, TransactionPartnerOther
type TransactionPartner interface {
	// IsTransactionPartner does nothing and is only used to enforce type-safety
	IsTransactionPartner()
}

// Describes a transaction with a user.
type TransactionPartnerUser struct {
	Type               string         `json:"type"`                          // Type of the transaction partner, always “user”
	User               User           `json:"user"`                          // Information about the user
	Affiliate          *AffiliateInfo `json:"affiliate,omitempty"`           // Optional. Information about the affiliate that received a commission via this transaction
	InvoicePayload     string         `json:"invoice_payload,omitempty"`     // Optional. Bot-specified invoice payload
	SubscriptionPeriod int64          `json:"subscription_period,omitempty"` // Optional. The duration of the paid subscription
	PaidMedia          []PaidMedia    `json:"paid_media,omitempty"`          // Optional. Information about the paid media bought by the user
	PaidMediaPayload   string         `json:"paid_media_payload,omitempty"`  // Optional. Bot-specified paid media payload
	Gift               *Gift          `json:"gift,omitempty"`                // Optional. The gift sent to the user by the bot
}

func (TransactionPartnerUser) IsTransactionPartner() {}

// Describes the affiliate program that issued the affiliate commission received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`                   // Type of the transaction partner, always “affiliate_program”
	SponsorUser        *User  `json:"sponsor_user,omitempty"` // Optional. Information about the bot that sponsored the affiliate program
	CommissionPerMille int64  `json:"commission_per_mille"`   // The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users
}

func (TransactionPartnerAffiliateProgram) IsTransactionPartner() {}

// Describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	Type            string                 `json:"type"`                       // Type of the transaction partner, always “fragment”
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state,omitempty"` // Optional. State of the transaction if the transaction is outgoing
}

func (TransactionPartnerFragment) IsTransactionPartner() {}

func (x *TransactionPartnerFragment) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Type            string          `json:"type"`                       // Type of the transaction partner, always “fragment”
		WithdrawalState json.RawMessage `json:"withdrawal_state,omitempty"` // Optional. State of the transaction if the transaction is outgoing
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalRevenueWithdrawalState(raw.WithdrawalState); err != nil {
		return err
	} else {
		x.WithdrawalState = data
	}
	x.Type = raw.Type

	return nil
}

// Describes a withdrawal transaction to the Telegram Ads platform.
type TransactionPartnerTelegramAds struct {
	Type string `json:"type"` // Type of the transaction partner, always “telegram_ads”
}

func (TransactionPartnerTelegramAds) IsTransactionPartner() {}

// Describes a transaction with payment for paid broadcasting.
type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`          // Type of the transaction partner, always “telegram_api”
	RequestCount int64  `json:"request_count"` // The number of successful requests that exceeded regular limits and were therefore billed
}

func (TransactionPartnerTelegramApi) IsTransactionPartner() {}

// Describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	Type string `json:"type"` // Type of the transaction partner, always “other”
}

func (TransactionPartnerOther) IsTransactionPartner() {}

// Describes a Telegram Star transaction.
type StarTransaction struct {
	Id             string             `json:"id"`                        // Unique identifier of the transaction. Coincides with the identifier of the original transaction for refund transactions. Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
	Amount         int64              `json:"amount"`                    // Integer amount of Telegram Stars transferred by the transaction
	NanostarAmount int64              `json:"nanostar_amount,omitempty"` // Optional. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999
	Date           int64              `json:"date"`                      // Date the transaction was created in Unix time
	Source         TransactionPartner `json:"source,omitempty"`          // Optional. Source of an incoming transaction (e.g., a user purchasing goods or services, Fragment refunding a failed withdrawal). Only for incoming transactions
	Receiver       TransactionPartner `json:"receiver,omitempty"`        // Optional. Receiver of an outgoing transaction (e.g., a user for a purchase refund, Fragment for a withdrawal). Only for outgoing transactions
}

func (x *StarTransaction) UnmarshalJSON(rawBytes []byte) (err error) {
	if len(rawBytes) == 0 {
		return nil
	}

	type temp struct {
		Id             string          `json:"id"`                        // Unique identifier of the transaction. Coincides with the identifier of the original transaction for refund transactions. Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
		Amount         int64           `json:"amount"`                    // Integer amount of Telegram Stars transferred by the transaction
		NanostarAmount int64           `json:"nanostar_amount,omitempty"` // Optional. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999
		Date           int64           `json:"date"`                      // Date the transaction was created in Unix time
		Source         json.RawMessage `json:"source,omitempty"`          // Optional. Source of an incoming transaction (e.g., a user purchasing goods or services, Fragment refunding a failed withdrawal). Only for incoming transactions
		Receiver       json.RawMessage `json:"receiver,omitempty"`        // Optional. Receiver of an outgoing transaction (e.g., a user for a purchase refund, Fragment for a withdrawal). Only for outgoing transactions
	}
	raw := &temp{}

	if err = json.Unmarshal(rawBytes, raw); err != nil {
		return err
	}

	if data, err := unmarshalTransactionPartner(raw.Source); err != nil {
		return err
	} else {
		x.Source = data
	}

	if data, err := unmarshalTransactionPartner(raw.Receiver); err != nil {
		return err
	} else {
		x.Receiver = data
	}
	x.Id = raw.Id
	x.Amount = raw.Amount
	x.NanostarAmount = raw.NanostarAmount
	x.Date = raw.Date

	return nil
}

// Contains a list of Telegram Star transactions.
type StarTransactions struct {
	Transactions []*StarTransaction `json:"transactions"` // The list of transactions
}

// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials EncryptedCredentials        `json:"credentials"` // Encrypted credentials required to decrypt the data
}

// PassportFile represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size"`      // File size in bytes
	FileDate     int64  `json:"file_date"`      // Unix time when the file was uploaded
}

// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string          `json:"type"`                   // Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Data        string          `json:"data,omitempty"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string          `json:"phone_number,omitempty"` // Optional. User's verified phone number; available only for “phone_number” type
	Email       string          `json:"email,omitempty"`        // Optional. User's verified email address; available only for “email” type
	Files       []*PassportFile `json:"files,omitempty"`        // Optional. Array of encrypted files with documents provided by the user; available only for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // Optional. Encrypted file with the front side of the document, provided by the user; available only for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // Optional. Encrypted file with the reverse side of the document, provided by the user; available only for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available if requested for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation,omitempty"`  // Optional. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string          `json:"hash"`                   // Base64-encoded element hash for using in PassportElementErrorUnspecified
}

// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrors struct {
	UserId int64                  `json:"user_id"` // User identifier
	Errors []PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (api *API) SetPassportDataErrors(payload *SetPassportDataErrors) (bool, error) {
	return callJson[bool](api, "setPassportDataErrors", payload)
}

// PassportElementError represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
// PassportElementErrorDataField, PassportElementErrorFrontSide, PassportElementErrorReverseSide, PassportElementErrorSelfie, PassportElementErrorFile, PassportElementErrorFiles, PassportElementErrorTranslationFile, PassportElementErrorTranslationFiles, PassportElementErrorUnspecified
type PassportElementError interface {
	// IsPassportElementError does nothing and is only used to enforce type-safety
	IsPassportElementError()
}

// Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	Source    string `json:"source"`     // Error source, must be data
	Type      string `json:"type"`       // The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	FieldName string `json:"field_name"` // Name of the data field which has the error
	DataHash  string `json:"data_hash"`  // Base64-encoded data hash
	Message   string `json:"message"`    // Error message
}

func (PassportElementErrorDataField) IsPassportElementError() {}

// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`    // Error source, must be front_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message"`   // Error message
}

func (PassportElementErrorFrontSide) IsPassportElementError() {}

// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`    // Error source, must be reverse_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the reverse side of the document
	Message  string `json:"message"`   // Error message
}

func (PassportElementErrorReverseSide) IsPassportElementError() {}

// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`    // Error source, must be selfie
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the selfie
	Message  string `json:"message"`   // Error message
}

func (PassportElementErrorSelfie) IsPassportElementError() {}

// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	Source   string `json:"source"`    // Error source, must be file
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

func (PassportElementErrorFile) IsPassportElementError() {}

// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`      // Error source, must be files
	Type       string   `json:"type"`        // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

func (PassportElementErrorFiles) IsPassportElementError() {}

// Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`    // Error source, must be translation_file
	Type     string `json:"type"`      // Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

func (PassportElementErrorTranslationFile) IsPassportElementError() {}

// Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`      // Error source, must be translation_files
	Type       string   `json:"type"`        // Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

func (PassportElementErrorTranslationFiles) IsPassportElementError() {}

// Represents an issue in an unspecified place. The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`       // Error source, must be unspecified
	Type        string `json:"type"`         // Type of element of the user's Telegram Passport which has the issue
	ElementHash string `json:"element_hash"` // Base64-encoded element hash
	Message     string `json:"message"`      // Error message
}

func (PassportElementErrorUnspecified) IsPassportElementError() {}

// sendGame is used to send a game. On success, the sent Message is returned.
type SendGame struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               int64                 `json:"chat_id"`                          // Unique identifier for the target chat
	MessageThreadId      int64                 `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	GameShortName        string                `json:"game_short_name"`                  // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	DisableNotification  bool                  `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                  `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                  `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      string                `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters      `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

// sendGame is used to send a game. On success, the sent Message is returned.
func (api *API) SendGame(payload *SendGame) (*Message, error) {
	return callJson[*Message](api, "sendGame", payload)
}

// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string           `json:"title"`                   // Title of the game
	Description  string           `json:"description"`             // Description of the game
	Photo        []*PhotoSize     `json:"photo"`                   // Photo that will be displayed in the game message in chats.
	Text         string           `json:"text,omitempty"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation       `json:"animation,omitempty"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// setGameScore is used to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	UserId             int64  `json:"user_id"`                        // User identifier
	Score              int64  `json:"score"`                          // New score, must be non-negative
	Force              bool   `json:"force,omitempty"`                // Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Pass True if the game message should not be automatically edited to include the current scoreboard
	ChatId             int64  `json:"chat_id,omitempty"`              // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId          int64  `json:"message_id,omitempty"`           // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId    string `json:"inline_message_id,omitempty"`    // Required if chat_id and message_id are not specified. Identifier of the inline message
}

// setGameScore is used to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (api *API) SetGameScore(payload *SetGameScore) (*Message, error) {
	return callJson[*Message](api, "setGameScore", payload)
}

// getGameHighScores is used to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
type GetGameHighScores struct {
	UserId          int64  `json:"user_id"`                     // Target user id
	ChatId          int64  `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int64  `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
}

// getGameHighScores is used to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
func (api *API) GetGameHighScores(payload *GetGameHighScores) ([]*GameHighScore, error) {
	return callJson[[]*GameHighScore](api, "getGameHighScores", payload)
}

// GameHighScore represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int64 `json:"position"` // Position in high score table for the game
	User     User  `json:"user"`     // User
	Score    int64 `json:"score"`    // Score
}

// ReplyMarkup is an interface for InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply
type ReplyMarkup interface {
	// IsReplyMarkup does nothing and is only used to enforce type-safety
	IsReplyMarkup()
}

// ChatID is an interface for usernames and chatIDs
type ChatID interface {
	// IsChatID does nothing and is only used to enforce type-safety
	IsChatID()
}
