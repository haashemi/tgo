// CODE GENERATED. DO NOT EDIT.
package tgo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Update represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId           int64               `json:"update_id"`                      // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message,omitempty"`              // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message,omitempty"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post,omitempty"`         // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`          // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`       // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`          // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`    // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
}

// getUpdates is used to receive incoming updates using long polling (wiki). Returns an Array of Update objects.// // Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.//
type GetUpdates struct {
	Offset         int64    `json:"offset,omitempty"`          // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
	Limit          int64    `json:"limit,omitempty"`           // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int64    `json:"timeout,omitempty"`         // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

// getUpdates is used to receive incoming updates using long polling (wiki). Returns an Array of Update objects.// // Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.//
func (api *API) GetUpdates(payload *GetUpdates) ([]*Update, error) {
	return callJson[[]*Update](api, "getUpdates", payload)
}

// setWebhook is used to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.// // Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.// If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.//
type SetWebhook struct {
	Url                string     `json:"url"`                            // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        *InputFile `json:"certificate,omitempty"`          // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string     `json:"ip_address,omitempty"`           // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int64      `json:"max_connections,omitempty"`      // The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string   `json:"allowed_updates,omitempty"`      // A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
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

// setWebhook is used to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.// // Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.// If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.//
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
}

// Chat represents a chat.
type Chat struct {
	Id                                 int64            `json:"id"`                                                // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                               string           `json:"type"`                                              // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                              string           `json:"title,omitempty"`                                   // Optional. Title, for supergroups, channels and group chats
	Username                           string           `json:"username,omitempty"`                                // Optional. Username, for private chats, supergroups and channels if available
	FirstName                          string           `json:"first_name,omitempty"`                              // Optional. First name of the other party in a private chat
	LastName                           string           `json:"last_name,omitempty"`                               // Optional. Last name of the other party in a private chat
	IsForum                            bool             `json:"is_forum,omitempty"`                                // Optional. True, if the supergroup chat is a forum (has topics enabled)
	Photo                              *ChatPhoto       `json:"photo,omitempty"`                                   // Optional. Chat photo. Returned only in getChat.
	ActiveUsernames                    []string         `json:"active_usernames,omitempty"`                        // Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels. Returned only in getChat.
	EmojiStatusCustomEmojiId           string           `json:"emoji_status_custom_emoji_id,omitempty"`            // Optional. Custom emoji identifier of emoji status of the other party in a private chat. Returned only in getChat.
	EmojiStatusExpirationDate          int64            `json:"emoji_status_expiration_date,omitempty"`            // Optional. Expiration date of the emoji status of the other party in a private chat in Unix time, if any. Returned only in getChat.
	Bio                                string           `json:"bio,omitempty"`                                     // Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards                 bool             `json:"has_private_forwards,omitempty"`                    // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages,omitempty"` // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat. Returned only in getChat.
	JoinToSendMessages                 bool             `json:"join_to_send_messages,omitempty"`                   // Optional. True, if users need to join the supergroup before they can send messages. Returned only in getChat.
	JoinByRequest                      bool             `json:"join_by_request,omitempty"`                         // Optional. True, if all users directly joining the supergroup need to be approved by supergroup administrators. Returned only in getChat.
	Description                        string           `json:"description,omitempty"`                             // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink                         string           `json:"invite_link,omitempty"`                             // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMessage                      *Message         `json:"pinned_message,omitempty"`                          // Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permissions                        *ChatPermissions `json:"permissions,omitempty"`                             // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay                      int64            `json:"slow_mode_delay,omitempty"`                         // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	MessageAutoDeleteTime              int64            `json:"message_auto_delete_time,omitempty"`                // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	HasAggressiveAntiSpamEnabled       bool             `json:"has_aggressive_anti_spam_enabled,omitempty"`        // Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators. Returned only in getChat.
	HasHiddenMembers                   bool             `json:"has_hidden_members,omitempty"`                      // Optional. True, if non-administrators can only get the list of bots and administrators in the chat. Returned only in getChat.
	HasProtectedContent                bool             `json:"has_protected_content,omitempty"`                   // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	StickerSetName                     string           `json:"sticker_set_name,omitempty"`                        // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet                   bool             `json:"can_set_sticker_set,omitempty"`                     // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatId                       int64            `json:"linked_chat_id,omitempty"`                          // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location                           *ChatLocation    `json:"location,omitempty"`                                // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

// Message represents a message.
type Message struct {
	MessageId                     int64                          `json:"message_id"`                                  // Unique message identifier inside this chat
	MessageThreadId               int64                          `json:"message_thread_id,omitempty"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from,omitempty"`                              // Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field from contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	Date                          int64                          `json:"date"`                                        // Date the message was sent in Unix time
	Chat                          Chat                           `json:"chat"`                                        // Conversation the message belongs to
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`                      // Optional. For forwarded messages, sender of the original message
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`                 // Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
	ForwardFromMessageId          int64                          `json:"forward_from_message_id,omitempty"`           // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature              string                         `json:"forward_signature,omitempty"`                 // Optional. For forwarded messages that were originally sent in channels or by an anonymous chat administrator, signature of the message sender if present
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`               // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate                   int64                          `json:"forward_date,omitempty"`                      // Optional. For forwarded messages, date the original message was sent in Unix time
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot through which the message was sent
	EditDate                      int64                          `json:"edit_date,omitempty"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True, if the message can't be forwarded
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text,omitempty"`                              // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      []*MessageEntity               `json:"entities,omitempty"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document,omitempty"`                          // Optional. Message is a general file, information about the file
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story,omitempty"`                             // Optional. Message is a forwarded story
	Video                         *Video                         `json:"video,omitempty"`                             // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for the animation, audio, document, photo, video or voice
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
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
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	UserShared                    *UserShared                    `json:"user_shared,omitempty"`                       // Optional. Service message: a user was shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Service message: forum topic edited
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Service message: forum topic closed
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Service message: forum topic reopened
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. Service message: the 'General' forum topic unhidden
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Service message: video chat ended
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

// MessageId represents a unique message identifier.
type MessageId struct {
	MessageId int64 `json:"message_id"` // Unique message identifier
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type          string `json:"type"`                      // Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Offset        int64  `json:"offset"`                    // Offset in UTF-16 code units to the start of the entity
	Length        int64  `json:"length"`                    // Length of the entity in UTF-16 code units
	Url           string `json:"url,omitempty"`             // Optional. For “text_link” only, URL that will be opened after user taps on the text
	User          *User  `json:"user,omitempty"`            // Optional. For “text_mention” only, the mentioned user
	Language      string `json:"language,omitempty"`        // Optional. For “pre” only, the programming language of the entity text
	CustomEmojiId string `json:"custom_emoji_id,omitempty"` // Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

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
	Width        int64      `json:"width"`               // Video width as defined by sender
	Height       int64      `json:"height"`              // Video height as defined by sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Animation thumbnail as defined by sender
	FileName     string     `json:"file_name,omitempty"` // Optional. Original animation filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64      `json:"duration"`            // Duration of the audio in seconds as defined by sender
	Performer    string     `json:"performer,omitempty"` // Optional. Performer of the audio as defined by sender or by audio tags
	Title        string     `json:"title,omitempty"`     // Optional. Title of the audio as defined by sender or by audio tags
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the album cover to which the music file belongs
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Document thumbnail as defined by sender
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Story represents a message about a forwarded story in the chat. Currently holds no information.
type Story struct{}

// Video represents a video file.
type Video struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int64      `json:"width"`               // Video width as defined by sender
	Height       int64      `json:"height"`              // Video height as defined by sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int64      `json:"length"`              // Video width and height (diameter of the video message) as defined by sender
	Duration     int64      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes
}

// Voice represents a voice note.
type Voice struct {
	FileId       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64  `json:"duration"`            // Duration of the audio in seconds as defined by sender
	MimeType     string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

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
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int64  `json:"voter_count"` // Number of users that voted for this option
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
	Longitude            float64 `json:"longitude"`                        // Longitude as defined by sender
	Latitude             float64 `json:"latitude"`                         // Latitude as defined by sender
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

// UserShared contains information about the user whose identifier was shared with the bot using a KeyboardButtonRequestUser button.
type UserShared struct {
	RequestId int64 `json:"request_id"` // Identifier of the request
	UserId    int64 `json:"user_id"`    // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
}

// ChatShared contains information about the chat whose identifier was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	RequestId int64 `json:"request_id"` // Identifier of the request
	ChatId    int64 `json:"chat_id"`    // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
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

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64          `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]*PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.// // The maximum file size to download is 20 MB//
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

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          bool                `json:"is_persistent,omitempty"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool                `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool                `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

func (ReplyKeyboardMarkup) IsReplyMarkup() {}

// KeyboardButton represents one button of the reply keyboard. For simple text buttons, String can be used instead of this object to specify the button text. The optional fields web_app, request_user, request_chat, request_contact, request_location, and request_poll are mutually exclusive.// Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.Note: request_poll option will only work in Telegram versions released after 23 January, 2020. Older clients will display unsupported message.Note: web_app option will only work in Telegram versions released after 16 April, 2022. Older clients will display unsupported message.Note: request_user and request_chat options will only work in Telegram versions released after 3 February, 2023. Older clients will display unsupported message.
type KeyboardButton struct {
	Text            string                     `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUser     *KeyboardButtonRequestUser `json:"request_user,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable users. Tapping on any user will send their identifier to the bot in a “user_shared” service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat `json:"request_chat,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestContact  bool                       `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation bool                       `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType    `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                `json:"web_app,omitempty"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

// KeyboardButtonRequestUser defines the criteria used to request a suitable user. The identifier of the selected user will be shared with the bot when the corresponding button is pressed. More about requesting users »
type KeyboardButtonRequestUser struct {
	RequestId     int64 `json:"request_id"`                // Signed 32-bit identifier of the request, which will be received back in the UserShared object. Must be unique within the message
	UserIsBot     bool  `json:"user_is_bot,omitempty"`     // Optional. Pass True to request a bot, pass False to request a regular user. If not specified, no additional restrictions are applied.
	UserIsPremium bool  `json:"user_is_premium,omitempty"` // Optional. Pass True to request a premium user, pass False to request a non-premium user. If not specified, no additional restrictions are applied.
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable chat. The identifier of the selected chat will be shared with the bot when the corresponding button is pressed. More about requesting chats »
type KeyboardButtonRequestChat struct {
	RequestId               int64                    `json:"request_id"`                          // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel           bool                     `json:"chat_is_channel"`                     // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

func (ReplyKeyboardRemove) IsReplyMarkup() {}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

func (InlineKeyboardMarkup) IsReplyMarkup() {}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	Url                          string                       `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username, if this is allowed by their privacy settings.
	CallbackData                 string                       `json:"callback_data,omitempty"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
	Pay                          bool                         `json:"pay,omitempty"`                              // Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:// Telegram apps support these buttons as of version 5.7.// // Sample bot: @discussbot//
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

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.// // NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any of the optional parameters).//
type CallbackQuery struct {
	Id              string   `json:"id"`                          // Unique identifier for this query
	From            User     `json:"from"`                        // Sender
	Message         *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageId string   `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string   `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   string   `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.// // Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions). There could be two ways to create a new poll:// // Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing for hardcore users but lacks modern day polish.// Guide the user through a step-by-step process. 'Please send me your question', 'Cool, now let's add the first answer option', 'Great. Keep adding answer options, then send /done when you're ready'.// // The last option is definitely more attractive. And if you use ForceReply in your bot's questions, it will receive the user's answers even if it only receives replies, commands and mentions - without any extra work for the user.//
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`                       // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool   `json:"selective,omitempty"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
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
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`                 // True, if the user's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat"`              // True, if the administrator can access the chat event log, boost list in channels, see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool `json:"can_delete_messages"`          // True, if the administrator can delete messages of other users
	CanManageVideoChats bool `json:"can_manage_video_chats"`       // True, if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members"`         // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool `json:"can_promote_members"`          // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool `json:"can_change_info"`              // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users"`             // True, if the user is allowed to invite new users to the chat
	CanPostMessages     bool `json:"can_post_messages,omitempty"`  // Optional. True, if the administrator can post messages in the channel, or access channel statistics; channels only
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`  // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`   // Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPostStories      bool `json:"can_post_stories,omitempty"`   // Optional. True, if the administrator can post stories in the channel; channels only
	CanEditStories      bool `json:"can_edit_stories,omitempty"`   // Optional. True, if the administrator can edit stories posted by other users; channels only
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"` // Optional. True, if the administrator can delete stories posted by other users; channels only
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`  // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
}

// ChatMember contains information about one member of a chat. Currently, the following 6 types of chat members are supported:// ChatMemberOwner, ChatMemberAdministrator, ChatMemberMember, ChatMemberRestricted, ChatMemberLeft, ChatMemberBanned
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
	Status              string `json:"status"`                       // The member's status in the chat, always “administrator”
	User                User   `json:"user"`                         // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`                // True, if the bot is allowed to edit administrator privileges of that user
	IsAnonymous         bool   `json:"is_anonymous"`                 // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`              // True, if the administrator can access the chat event log, boost list in channels, see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool   `json:"can_delete_messages"`          // True, if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats"`       // True, if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`         // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members"`          // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`              // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`             // True, if the user is allowed to invite new users to the chat
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`  // Optional. True, if the administrator can post messages in the channel, or access channel statistics; channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`  // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`   // Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPostStories      bool   `json:"can_post_stories,omitempty"`   // Optional. True, if the administrator can post stories in the channel; channels only
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`   // Optional. True, if the administrator can edit stories posted by other users; channels only
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"` // Optional. True, if the administrator can delete stories posted by other users; channels only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`  // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
	CustomTitle         string `json:"custom_title,omitempty"`       // Optional. Custom title for this user
}

func (ChatMemberAdministrator) IsChatMember() {}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status string `json:"status"` // The member's status in the chat, always “member”
	User   User   `json:"user"`   // Information about the user
}

func (ChatMemberMember) IsChatMember() {}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always “restricted”
	User                  User   `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, invoices, locations and venues
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

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                                  // Chat the user belongs to
	From                    User            `json:"from"`                                  // Performer of the action, which resulted in the change
	Date                    int64           `json:"date"`                                  // Date the change was done in Unix time
	OldChatMember           ChatMember      `json:"old_chat_member"`                       // Previous information about the chat member
	NewChatMember           ChatMember      `json:"new_chat_member"`                       // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
}

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
	CanSendMessages       bool `json:"can_send_messages,omitempty"`         // Optional. True, if the user is allowed to send text messages, contacts, invoices, locations and venues
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

// Represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string   `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
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

// BotCommandScope represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:// BotCommandScopeDefault, BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators, BotCommandScopeChat, BotCommandScopeChatAdministrators, BotCommandScopeChatMember
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

// Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`    // Scope type, must be chat_administrators
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (BotCommandScopeChatAdministrators) IsBotCommandScope() {}

// Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`    // Scope type, must be chat_member
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

func (BotCommandScopeChatMember) IsBotCommandScope() {}

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

// MenuButton describes the bot's menu button in a private chat. It should be one of// MenuButtonCommands, MenuButtonWebApp, MenuButtonDefault// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
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
	WebApp WebAppInfo `json:"web_app"` // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
}

func (MenuButtonWebApp) IsMenuButton() {}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	Type string `json:"type"` // Type of the button, must be default
}

func (MenuButtonDefault) IsMenuButton() {}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int64 `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// InputMedia represents the content of a media message to be sent. It should be one of// InputMediaAnimation, InputMediaDocument, InputMediaAudio, InputMediaPhoto, InputMediaVideo
type InputMedia interface {
	// IsInputMedia does nothing and is only used to enforce type-safety
	IsInputMedia()

	getFiles() map[string]*InputFile
}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	Type            string           `json:"type"`                       // Type of the result, must be photo
	Media           *InputFile       `json:"media"`                      // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler      bool             `json:"has_spoiler,omitempty"`      // Optional. Pass True if the photo needs to be covered with a spoiler animation
}

func (InputMediaPhoto) IsInputMedia() {}

func (x *InputMediaPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media.IsUploadable() {
		media["media"] = x.Media
	}

	return media
}

// Represents a video to be sent.
type InputMediaVideo struct {
	Type              string           `json:"type"`                         // Type of the result, must be video
	Media             *InputFile       `json:"media"`                        // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail         *InputFile       `json:"thumbnail,omitempty"`          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption           string           `json:"caption,omitempty"`            // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode         ParseMode        `json:"parse_mode,omitempty"`         // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities   []*MessageEntity `json:"caption_entities,omitempty"`   // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width             int64            `json:"width,omitempty"`              // Optional. Video width
	Height            int64            `json:"height,omitempty"`             // Optional. Video height
	Duration          int64            `json:"duration,omitempty"`           // Optional. Video duration in seconds
	SupportsStreaming bool             `json:"supports_streaming,omitempty"` // Optional. Pass True if the uploaded video is suitable for streaming
	HasSpoiler        bool             `json:"has_spoiler,omitempty"`        // Optional. Pass True if the video needs to be covered with a spoiler animation
}

func (InputMediaVideo) IsInputMedia() {}

func (x *InputMediaVideo) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media.IsUploadable() {
		media["media"] = x.Media
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type            string           `json:"type"`                       // Type of the result, must be animation
	Media           *InputFile       `json:"media"`                      // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail       *InputFile       `json:"thumbnail,omitempty"`        // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width           int64            `json:"width,omitempty"`            // Optional. Animation width
	Height          int64            `json:"height,omitempty"`           // Optional. Animation height
	Duration        int64            `json:"duration,omitempty"`         // Optional. Animation duration in seconds
	HasSpoiler      bool             `json:"has_spoiler,omitempty"`      // Optional. Pass True if the animation needs to be covered with a spoiler animation
}

func (InputMediaAnimation) IsInputMedia() {}

func (x *InputMediaAnimation) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Media.IsUploadable() {
		media["media"] = x.Media
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
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

	if x.Media.IsUploadable() {
		media["media"] = x.Media
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
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

	if x.Media.IsUploadable() {
		media["media"] = x.Media
	}
	if x.Thumbnail != nil {
		if x.Thumbnail.IsUploadable() {
			media["thumbnail"] = x.Thumbnail
		}
	}

	return media
}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (api *API) GetMe() (*User, error) {
	return callJson[*User](api, "getMe", nil)
}

// logOut is used to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.

// logOut is used to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (api *API) LogOut() (bool, error) {
	return callJson[bool](api, "logOut", nil)
}

// close is used to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.

// close is used to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (api *API) Close() (bool, error) {
	return callJson[bool](api, "close", nil)
}

// sendMessage is used to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                     string           `json:"text"`                                  // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the message text. See formatting options for more details.
	Entities                 []*MessageEntity `json:"entities,omitempty"`                    // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`    // Disables link previews for links in this message
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendMessage is used to send text messages. On success, the sent Message is returned.
func (api *API) SendMessage(payload *SendMessage) (*Message, error) {
	return callJson[*Message](api, "sendMessage", payload)
}

// forwardMessage is used to forward messages of any kind. Service messages can't be forwarded. On success, the sent Message is returned.
type ForwardMessage struct {
	ChatId              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId     int64  `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId          ChatID `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool   `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Protects the contents of the forwarded message from forwarding and saving
	MessageId           int64  `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

// forwardMessage is used to forward messages of any kind. Service messages can't be forwarded. On success, the sent Message is returned.
func (api *API) ForwardMessage(payload *ForwardMessage) (*Message, error) {
	return callJson[*Message](api, "forwardMessage", payload)
}

// copyMessage is used to copy messages of any kind. Service messages and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessage struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId               ChatID           `json:"from_chat_id"`                          // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId                int64            `json:"message_id"`                            // Message identifier in the chat specified in from_chat_id
	Caption                  string           `json:"caption,omitempty"`                     // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// copyMessage is used to copy messages of any kind. Service messages and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (api *API) CopyMessage(payload *CopyMessage) (*MessageId, error) {
	return callJson[*MessageId](api, "copyMessage", payload)
}

// sendPhoto is used to send photos. On success, the sent Message is returned.
type SendPhoto struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Photo                    *InputFile       `json:"photo"`                                 // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendPhoto) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Photo.IsUploadable() {
		media["photo"] = x.Photo
	}

	return media
}

func (x *SendPhoto) getParams() (map[string]string, error) {
	payload := map[string]string{}

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
	if x.HasSpoiler {
		payload["has_spoiler"] = strconv.FormatBool(x.HasSpoiler)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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

// sendAudio is used to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Audio                    *InputFile       `json:"audio"`                                 // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Audio caption, 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int64            `json:"duration,omitempty"`                    // Duration of the audio in seconds
	Performer                string           `json:"performer,omitempty"`                   // Performer
	Title                    string           `json:"title,omitempty"`                       // Track name
	Thumbnail                *InputFile       `json:"thumbnail,omitempty"`                   // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendAudio) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Audio.IsUploadable() {
		media["audio"] = x.Audio
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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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

// sendAudio is used to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.// For sending voice messages, use the sendVoice method instead.
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
	ReplyToMessageId            int64            `json:"reply_to_message_id,omitempty"`            // If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply,omitempty"`    // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`                   // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendDocument) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Document.IsUploadable() {
		media["document"] = x.Document
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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Video                    *InputFile       `json:"video"`                                 // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files »
	Duration                 int64            `json:"duration,omitempty"`                    // Duration of sent video in seconds
	Width                    int64            `json:"width,omitempty"`                       // Video width
	Height                   int64            `json:"height,omitempty"`                      // Video height
	Thumbnail                *InputFile       `json:"thumbnail,omitempty"`                   // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Pass True if the video needs to be covered with a spoiler animation
	SupportsStreaming        bool             `json:"supports_streaming,omitempty"`          // Pass True if the uploaded video is suitable for streaming
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendVideo) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Video.IsUploadable() {
		media["video"] = x.Video
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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Animation                *InputFile       `json:"animation"`                             // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files »
	Duration                 int64            `json:"duration,omitempty"`                    // Duration of sent animation in seconds
	Width                    int64            `json:"width,omitempty"`                       // Animation width
	Height                   int64            `json:"height,omitempty"`                      // Animation height
	Thumbnail                *InputFile       `json:"thumbnail,omitempty"`                   // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Pass True if the animation needs to be covered with a spoiler animation
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendAnimation) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Animation.IsUploadable() {
		media["animation"] = x.Animation
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
	if x.HasSpoiler {
		payload["has_spoiler"] = strconv.FormatBool(x.HasSpoiler)
	}
	if x.DisableNotification {
		payload["disable_notification"] = strconv.FormatBool(x.DisableNotification)
	}
	if x.ProtectContent {
		payload["protect_content"] = strconv.FormatBool(x.ProtectContent)
	}
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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

// sendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Voice                    *InputFile       `json:"voice"`                                 // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Voice message caption, 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int64            `json:"duration,omitempty"`                    // Duration of the voice message in seconds
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendVoice) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Voice.IsUploadable() {
		media["voice"] = x.Voice
	}

	return media
}

func (x *SendVoice) getParams() (map[string]string, error) {
	payload := map[string]string{}

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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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

// sendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
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
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	VideoNote                *InputFile  `json:"video_note"`                            // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration                 int64       `json:"duration,omitempty"`                    // Duration of sent video in seconds
	Length                   int64       `json:"length,omitempty"`                      // Video width and height, i.e. diameter of the video message
	Thumbnail                *InputFile  `json:"thumbnail,omitempty"`                   // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendVideoNote) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.VideoNote.IsUploadable() {
		media["video_note"] = x.VideoNote
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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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

// sendMediaGroup is used to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
	ChatId                   ChatID       `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64        `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Media                    []InputMedia `json:"media"`                                 // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification      bool         `json:"disable_notification,omitempty"`        // Sends messages silently. Users will receive a notification with no sound.
	ProtectContent           bool         `json:"protect_content,omitempty"`             // Protects the contents of the sent messages from forwarding and saving
	ReplyToMessageId         int64        `json:"reply_to_message_id,omitempty"`         // If the messages are a reply, ID of the original message
	AllowSendingWithoutReply bool         `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
}

func (x *SendMediaGroup) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for idx, m := range x.Media {
		for key, value := range m.getFiles() {
			value.Value = fmt.Sprintf("%d.media.%s", idx, key)
			media[value.Value] = value
		}
	}

	return media
}

func (x *SendMediaGroup) getParams() (map[string]string, error) {
	payload := map[string]string{}

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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude                 float64     `json:"latitude"`                              // Latitude of the location
	Longitude                float64     `json:"longitude"`                             // Longitude of the location
	HorizontalAccuracy       float64     `json:"horizontal_accuracy,omitempty"`         // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod               int64       `json:"live_period,omitempty"`                 // Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	Heading                  int64       `json:"heading,omitempty"`                     // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius     int64       `json:"proximity_alert_radius,omitempty"`      // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendLocation is used to send point on the map. On success, the sent Message is returned.
func (api *API) SendLocation(payload *SendLocation) (*Message, error) {
	return callJson[*Message](api, "sendLocation", payload)
}

// sendVenue is used to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude                 float64     `json:"latitude"`                              // Latitude of the venue
	Longitude                float64     `json:"longitude"`                             // Longitude of the venue
	Title                    string      `json:"title"`                                 // Name of the venue
	Address                  string      `json:"address"`                               // Address of the venue
	FoursquareId             string      `json:"foursquare_id,omitempty"`               // Foursquare identifier of the venue
	FoursquareType           string      `json:"foursquare_type,omitempty"`             // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId            string      `json:"google_place_id,omitempty"`             // Google Places identifier of the venue
	GooglePlaceType          string      `json:"google_place_type,omitempty"`           // Google Places type of the venue. (See supported types.)
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendVenue is used to send information about a venue. On success, the sent Message is returned.
func (api *API) SendVenue(payload *SendVenue) (*Message, error) {
	return callJson[*Message](api, "sendVenue", payload)
}

// sendContact is used to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	PhoneNumber              string      `json:"phone_number"`                          // Contact's phone number
	FirstName                string      `json:"first_name"`                            // Contact's first name
	LastName                 string      `json:"last_name,omitempty"`                   // Contact's last name
	Vcard                    string      `json:"vcard,omitempty"`                       // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendContact is used to send phone contacts. On success, the sent Message is returned.
func (api *API) SendContact(payload *SendContact) (*Message, error) {
	return callJson[*Message](api, "sendContact", payload)
}

// sendPoll is used to send a native poll. On success, the sent Message is returned.
type SendPoll struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Question                 string           `json:"question"`                              // Poll question, 1-300 characters
	Options                  []string         `json:"options"`                               // A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous              bool             `json:"is_anonymous,omitempty"`                // True, if the poll needs to be anonymous, defaults to True
	Type                     string           `json:"type,omitempty"`                        // Poll type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers    bool             `json:"allows_multiple_answers,omitempty"`     // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId          int64            `json:"correct_option_id,omitempty"`           // 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation              string           `json:"explanation,omitempty"`                 // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode     string           `json:"explanation_parse_mode,omitempty"`      // Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities      []*MessageEntity `json:"explanation_entities,omitempty"`        // A JSON-serialized list of special entities that appear in the poll explanation, which can be specified instead of parse_mode
	OpenPeriod               int64            `json:"open_period,omitempty"`                 // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate                int64            `json:"close_date,omitempty"`                  // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed                 bool             `json:"is_closed,omitempty"`                   // Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendPoll is used to send a native poll. On success, the sent Message is returned.
func (api *API) SendPoll(payload *SendPoll) (*Message, error) {
	return callJson[*Message](api, "sendPoll", payload)
}

// sendDice is used to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDice struct {
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                    string      `json:"emoji,omitempty"`                       // Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64 for “”. Defaults to “”
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendDice is used to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (api *API) SendDice(payload *SendDice) (*Message, error) {
	return callJson[*Message](api, "sendDice", payload)
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.// // Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.// // We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
type SendChatAction struct {
	ChatId          ChatID `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId int64  `json:"message_thread_id,omitempty"` // Unique identifier for the target message thread; supergroups only
	Action          string `json:"action"`                      // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.// // Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.// // We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (api *API) SendChatAction(payload *SendChatAction) (bool, error) {
	return callJson[bool](api, "sendChatAction", payload)
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

// getFile is used to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	FileId string `json:"file_id"` // File identifier to get information about
}

// getFile is used to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
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
	CanManageChat       bool   `json:"can_manage_chat,omitempty"`        // Pass True if the administrator can access the chat event log, boost list in channels, see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool   `json:"can_delete_messages,omitempty"`    // Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats,omitempty"` // Pass True if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members,omitempty"`   // Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members,omitempty"`    // Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool   `json:"can_change_info,omitempty"`        // Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users,omitempty"`       // Pass True if the administrator can invite new users to the chat
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`      // Pass True if the administrator can post messages in the channel, or access channel statistics; channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`      // Pass True if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`       // Pass True if the administrator can pin messages, supergroups only
	CanPostStories      bool   `json:"can_post_stories,omitempty"`       // Pass True if the administrator can post stories in the channel; channels only
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`       // Pass True if the administrator can edit stories posted by other users; channels only
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"`     // Pass True if the administrator can delete stories posted by other users; channels only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`      // Pass True if the user is allowed to create, rename, close, and reopen forum topics, supergroups only
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

// exportChatInviteLink is used to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.// // Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.//
type ExportChatInviteLink struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// exportChatInviteLink is used to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.// // Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.//
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

	if x.Photo.IsUploadable() {
		media["photo"] = x.Photo
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
	ChatId              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId           int64  `json:"message_id"`                     // Identifier of a message to pin
	DisableNotification bool   `json:"disable_notification,omitempty"` // Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

// pinChatMessage is used to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (api *API) PinChatMessage(payload *PinChatMessage) (bool, error) {
	return callJson[bool](api, "pinChatMessage", payload)
}

// unpinChatMessage is used to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinChatMessage struct {
	ChatId    ChatID `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id,omitempty"` // Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
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

// getChat is used to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChat struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// getChat is used to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (api *API) GetChat(payload *GetChat) (*Chat, error) {
	return callJson[*Chat](api, "getChat", payload)
}

// getChatAdministrators is used to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
type GetChatAdministrators struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// getChatAdministrators is used to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
func (api *API) GetChatAdministrators(payload *GetChatAdministrators) ([]*ChatMember, error) {
	return callJson[[]*ChatMember](api, "getChatAdministrators", payload)
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
func (api *API) GetChatMember(payload *GetChatMember) (*ChatMember, error) {
	return callJson[*ChatMember](api, "getChatMember", payload)
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

// editForumTopic is used to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type EditForumTopic struct {
	ChatId            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId   int64  `json:"message_thread_id"`              // Unique identifier for the target message thread of the forum topic
	Name              string `json:"name,omitempty"`                 // New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

// editForumTopic is used to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
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

// editGeneralForumTopic is used to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
type EditGeneralForumTopic struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // New topic name, 1-128 characters
}

// editGeneralForumTopic is used to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
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

// answerCallbackQuery is used to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.// // Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.//
type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`    // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`       // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert,omitempty"` // If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url,omitempty"`        // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int64  `json:"cache_time,omitempty"` // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

// answerCallbackQuery is used to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.// // Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.//
func (api *API) AnswerCallbackQuery(payload *AnswerCallbackQuery) (bool, error) {
	return callJson[bool](api, "answerCallbackQuery", payload)
}

// setMyCommands is used to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
type SetMyCommands struct {
	Commands     []*BotCommand    `json:"commands"`                // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// setMyCommands is used to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (api *API) SetMyCommands(payload *SetMyCommands) (bool, error) {
	return callJson[bool](api, "setMyCommands", payload)
}

// deleteMyCommands is used to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommands struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// deleteMyCommands is used to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (api *API) DeleteMyCommands(payload *DeleteMyCommands) (bool, error) {
	return callJson[bool](api, "deleteMyCommands", payload)
}

// getMyCommands is used to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
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
	ChatId     int64       `json:"chat_id,omitempty"`     // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton *MenuButton `json:"menu_button,omitempty"` // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
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
func (api *API) GetChatMenuButton(payload *GetChatMenuButton) (*MenuButton, error) {
	return callJson[*MenuButton](api, "getChatMenuButton", payload)
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

// editMessageText is used to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageText struct {
	ChatId                ChatID                `json:"chat_id,omitempty"`                  // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int64                 `json:"message_id,omitempty"`               // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string                `json:"inline_message_id,omitempty"`        // Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string                `json:"text"`                               // New text of the message, 1-4096 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity      `json:"entities,omitempty"`                 // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"` // Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // A JSON-serialized object for an inline keyboard.
}

// editMessageText is used to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) EditMessageText(payload *EditMessageText) (*Message, error) {
	return callJson[*Message](api, "editMessageText", payload)
}

// editMessageCaption is used to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageCaption struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string                `json:"caption,omitempty"`           // New caption of the message, 0-1024 characters after entities parsing
	ParseMode       ParseMode             `json:"parse_mode,omitempty"`        // Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities []*MessageEntity      `json:"caption_entities,omitempty"`  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for an inline keyboard.
}

// editMessageCaption is used to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) EditMessageCaption(payload *EditMessageCaption) (*Message, error) {
	return callJson[*Message](api, "editMessageCaption", payload)
}

// editMessageMedia is used to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageMedia struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	Media           InputMedia            `json:"media"`                       // A JSON-serialized object for a new media content of the message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

func (x *EditMessageMedia) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	for key, value := range x.Media.getFiles() {
		value.Value = "media." + key
		media[value.Value] = value
	}

	return media
}

func (x *EditMessageMedia) getParams() (map[string]string, error) {
	payload := map[string]string{}

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

// editMessageMedia is used to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
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
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               `json:"latitude"`                         // Latitude of new location
	Longitude            float64               `json:"longitude"`                        // Longitude of new location
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
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

// stopMessageLiveLocation is used to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) StopMessageLiveLocation(payload *StopMessageLiveLocation) (*Message, error) {
	return callJson[*Message](api, "stopMessageLiveLocation", payload)
}

// editMessageReplyMarkup is used to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkup struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for an inline keyboard.
}

// editMessageReplyMarkup is used to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (api *API) EditMessageReplyMarkup(payload *EditMessageReplyMarkup) (*Message, error) {
	return callJson[*Message](api, "editMessageReplyMarkup", payload)
}

// stopPoll is used to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
type StopPoll struct {
	ChatId      ChatID                `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int64                 `json:"message_id"`             // Identifier of the original message with the poll
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // A JSON-serialized object for a new message inline keyboard.
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
	IsAnimated  bool       `json:"is_animated"`         // True, if the sticker set contains animated stickers
	IsVideo     bool       `json:"is_video"`            // True, if the sticker set contains video stickers
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
	EmojiList    []string      `json:"emoji_list"`              // List of 1-20 emoji associated with the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	Keywords     []string      `json:"keywords,omitempty"`      // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
}

func (x *InputSticker) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker.IsUploadable() {
		media["sticker"] = x.Sticker
	}

	return media
}

// sendSticker is used to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendSticker struct {
	ChatId                   ChatID      `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Sticker                  *InputFile  `json:"sticker"`                               // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP or .TGS sticker using multipart/form-data. More information on Sending Files ». Video stickers can only be sent by a file_id. Animated stickers can't be sent via an HTTP URL.
	Emoji                    string      `json:"emoji,omitempty"`                       // Emoji associated with the sticker; only for just uploaded stickers
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (x *SendSticker) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker.IsUploadable() {
		media["sticker"] = x.Sticker
	}

	return media
}

func (x *SendSticker) getParams() (map[string]string, error) {
	payload := map[string]string{}

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
	if x.ReplyToMessageId != 0 {
		payload["reply_to_message_id"] = strconv.FormatInt(x.ReplyToMessageId, 10)
	}
	if x.AllowSendingWithoutReply {
		payload["allow_sending_without_reply"] = strconv.FormatBool(x.AllowSendingWithoutReply)
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
	CustomEmojiIds []string `json:"custom_emoji_ids"` // List of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

// getCustomEmojiStickers is used to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (api *API) GetCustomEmojiStickers(payload *GetCustomEmojiStickers) ([]*Sticker, error) {
	return callJson[[]*Sticker](api, "getCustomEmojiStickers", payload)
}

// uploadStickerFile is used to upload a file with a sticker for later use in the createNewStickerSet and addStickerToSet methods (the file can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
	UserId        int64      `json:"user_id"`        // User identifier of sticker file owner
	Sticker       *InputFile `json:"sticker"`        // A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files »
	StickerFormat string     `json:"sticker_format"` // Format of the sticker, must be one of “static”, “animated”, “video”
}

func (x *UploadStickerFile) getFiles() map[string]*InputFile {
	media := map[string]*InputFile{}

	if x.Sticker.IsUploadable() {
		media["sticker"] = x.Sticker
	}

	return media
}

func (x *UploadStickerFile) getParams() (map[string]string, error) {
	payload := map[string]string{}

	payload["user_id"] = strconv.FormatInt(x.UserId, 10)
	payload["sticker_format"] = x.StickerFormat
	return payload, nil
}

// uploadStickerFile is used to upload a file with a sticker for later use in the createNewStickerSet and addStickerToSet methods (the file can be used multiple times). Returns the uploaded File on success.
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
	StickerFormat   string          `json:"sticker_format"`             // Format of stickers in the set, must be one of “static”, “animated”, “video”
	StickerType     string          `json:"sticker_type,omitempty"`     // Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a regular sticker set is created.
	NeedsRepainting bool            `json:"needs_repainting,omitempty"` // Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
}

// createNewStickerSet is used to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
func (api *API) CreateNewStickerSet(payload *CreateNewStickerSet) (bool, error) {
	return callJson[bool](api, "createNewStickerSet", payload)
}

// addStickerToSet is used to add a new sticker to a set created by the bot. The format of the added sticker must match the format of the other stickers in the set. Emoji sticker sets can have up to 200 stickers. Animated and video sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
type AddStickerToSet struct {
	UserId  int64        `json:"user_id"` // User identifier of sticker set owner
	Name    string       `json:"name"`    // Sticker set name
	Sticker InputSticker `json:"sticker"` // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
}

// addStickerToSet is used to add a new sticker to a set created by the bot. The format of the added sticker must match the format of the other stickers in the set. Emoji sticker sets can have up to 200 stickers. Animated and video sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
func (api *API) AddStickerToSet(payload *AddStickerToSet) (bool, error) {
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
	Thumbnail *InputFile `json:"thumbnail,omitempty"` // A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animated-sticker-requirements for animated sticker technical requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-sticker-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
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
	Results       []*InlineQueryResult      `json:"results"`               // A JSON-serialized array of results for the inline query
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

// InlineQueryResult represents one result of an inline query. Telegram clients currently support results of the following 20 types:// InlineQueryResultCachedAudio, InlineQueryResultCachedDocument, InlineQueryResultCachedGif, InlineQueryResultCachedMpeg4Gif, InlineQueryResultCachedPhoto, InlineQueryResultCachedSticker, InlineQueryResultCachedVideo, InlineQueryResultCachedVoice, InlineQueryResultArticle, InlineQueryResultAudio, InlineQueryResultContact, InlineQueryResultGame, InlineQueryResultDocument, InlineQueryResultGif, InlineQueryResultLocation, InlineQueryResultMpeg4Gif, InlineQueryResultPhoto, InlineQueryResultVenue, InlineQueryResultVideo, InlineQueryResultVoice// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
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

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`                            // Type of the result, must be photo
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	PhotoUrl            string                `json:"photo_url"`                       // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	ThumbnailUrl        string                `json:"thumbnail_url"`                   // URL of the thumbnail for the photo
	PhotoWidth          int64                 `json:"photo_width,omitempty"`           // Optional. Width of the photo
	PhotoHeight         int64                 `json:"photo_height,omitempty"`          // Optional. Height of the photo
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the photo
}

func (InlineQueryResultPhoto) IsInlineQueryResult() {}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                string                `json:"type"`                            // Type of the result, must be gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	GifUrl              string                `json:"gif_url"`                         // A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth            int64                 `json:"gif_width,omitempty"`             // Optional. Width of the GIF
	GifHeight           int64                 `json:"gif_height,omitempty"`            // Optional. Height of the GIF
	GifDuration         int64                 `json:"gif_duration,omitempty"`          // Optional. Duration of the GIF in seconds
	ThumbnailUrl        string                `json:"thumbnail_url"`                   // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`   // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the GIF animation
}

func (InlineQueryResultGif) IsInlineQueryResult() {}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`                            // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Mpeg4Url            string                `json:"mpeg4_url"`                       // A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4Width          int64                 `json:"mpeg4_width,omitempty"`           // Optional. Video width
	Mpeg4Height         int64                 `json:"mpeg4_height,omitempty"`          // Optional. Video height
	Mpeg4Duration       int64                 `json:"mpeg4_duration,omitempty"`        // Optional. Video duration in seconds
	ThumbnailUrl        string                `json:"thumbnail_url"`                   // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`   // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video animation
}

func (InlineQueryResultMpeg4Gif) IsInlineQueryResult() {}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.// // If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using input_message_content.//
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoUrl            string                `json:"video_url"`                       // A valid URL for the embedded video player or video file
	MimeType            string                `json:"mime_type"`                       // MIME type of the content of the video URL, “text/html” or “video/mp4”
	ThumbnailUrl        string                `json:"thumbnail_url"`                   // URL of the thumbnail (JPEG only) for the video
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	VideoWidth          int64                 `json:"video_width,omitempty"`           // Optional. Video width
	VideoHeight         int64                 `json:"video_height,omitempty"`          // Optional. Video height
	VideoDuration       int64                 `json:"video_duration,omitempty"`        // Optional. Video duration in seconds
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}

func (InlineQueryResultVideo) IsInlineQueryResult() {}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
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
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

func (InlineQueryResultAudio) IsInlineQueryResult() {}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
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
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice recording
}

func (InlineQueryResultVoice) IsInlineQueryResult() {}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
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
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultDocument) IsInlineQueryResult() {}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`                             // Type of the result, must be location
	Id                   string                `json:"id"`                               // Unique identifier for this result, 1-64 Bytes
	Latitude             float64               `json:"latitude"`                         // Location latitude in degrees
	Longitude            float64               `json:"longitude"`                        // Location longitude in degrees
	Title                string                `json:"title"`                            // Location title
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64                 `json:"live_period,omitempty"`            // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	Heading              int64                 `json:"heading,omitempty"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional. Inline keyboard attached to the message
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`  // Optional. Content of the message to be sent instead of the location
	ThumbnailUrl         string                `json:"thumbnail_url,omitempty"`          // Optional. Url of the thumbnail for the result
	ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`        // Optional. Thumbnail width
	ThumbnailHeight      int64                 `json:"thumbnail_height,omitempty"`       // Optional. Thumbnail height
}

func (InlineQueryResultLocation) IsInlineQueryResult() {}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
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
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the venue
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultVenue) IsInlineQueryResult() {}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                            // Type of the result, must be contact
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`                    // Contact's phone number
	FirstName           string                `json:"first_name"`                      // Contact's first name
	LastName            string                `json:"last_name,omitempty"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard,omitempty"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the contact
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`         // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`      // Optional. Thumbnail height
}

func (InlineQueryResultContact) IsInlineQueryResult() {}

// Represents a Game.// Note: This will only work in Telegram versions released after October 1, 2016. Older clients will not display any inline results if a game result is among them.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`                   // Type of the result, must be game
	Id            string                `json:"id"`                     // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"`        // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
}

func (InlineQueryResultGame) IsInlineQueryResult() {}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`                            // Type of the result, must be photo
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	PhotoFileId         string                `json:"photo_file_id"`                   // A valid file identifier of the photo
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the photo
}

func (InlineQueryResultCachedPhoto) IsInlineQueryResult() {}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`                            // Type of the result, must be gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	GifFileId           string                `json:"gif_file_id"`                     // A valid file identifier for the GIF file
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the GIF animation
}

func (InlineQueryResultCachedGif) IsInlineQueryResult() {}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`                            // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Mpeg4FileId         string                `json:"mpeg4_file_id"`                   // A valid file identifier for the MPEG4 file
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video animation
}

func (InlineQueryResultCachedMpeg4Gif) IsInlineQueryResult() {}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.// Note: This will only work in Telegram versions released after 9 April, 2016 for static stickers and after 06 July, 2019 for animated stickers. Older clients will ignore them.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                            // Type of the result, must be sticker
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	StickerFileId       string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
}

func (InlineQueryResultCachedSticker) IsInlineQueryResult() {}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
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
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
}

func (InlineQueryResultCachedDocument) IsInlineQueryResult() {}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoFileId         string                `json:"video_file_id"`                   // A valid file identifier for the video file
	Title               string                `json:"title"`                           // Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video
}

func (InlineQueryResultCachedVideo) IsInlineQueryResult() {}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                            // Type of the result, must be voice
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceFileId         string                `json:"voice_file_id"`                   // A valid file identifier for the voice message
	Title               string                `json:"title"`                           // Voice message title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice message
}

func (InlineQueryResultCachedVoice) IsInlineQueryResult() {}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioFileId         string                `json:"audio_file_id"`                   // A valid file identifier for the audio file
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

func (InlineQueryResultCachedAudio) IsInlineQueryResult() {}

// InputMessageContent represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:// InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent
type InputMessageContent interface {
	// IsInputMessageContent does nothing and is only used to enforce type-safety
	IsInputMessageContent()
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`                       // Text of the message to be sent, 1-4096 characters
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity `json:"entities,omitempty"`                 // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in the sent message
}

func (InputTextMessageContent) IsInputMessageContent() {}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`                         // Latitude of the location in degrees
	Longitude            float64 `json:"longitude"`                        // Longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64   `json:"live_period,omitempty"`            // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
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
	Payload                   string          `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string          `json:"provider_token"`                          // Payment provider token, obtained via @BotFather
	Currency                  string          `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int64           `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64         `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data,omitempty"`                 // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64           `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int64           `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int64           `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool            `json:"need_name,omitempty"`                     // Optional. Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool            `json:"need_phone_number,omitempty"`             // Optional. Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool            `json:"need_email,omitempty"`                    // Optional. Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool            `json:"need_shipping_address,omitempty"`         // Optional. Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True if the user's phone number should be sent to provider
	SendEmailToProvider       bool            `json:"send_email_to_provider,omitempty"`        // Optional. Pass True if the user's email address should be sent to provider
	IsFlexible                bool            `json:"is_flexible,omitempty"`                   // Optional. Pass True if the final price depends on the shipping method
}

func (InputInvoiceMessageContent) IsInputMessageContent() {}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
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

// sendInvoice is used to send invoices. On success, the sent Message is returned.
type SendInvoice struct {
	ChatId                    ChatID                `json:"chat_id"`                                 // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId           int64                 `json:"message_thread_id,omitempty"`             // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Title                     string                `json:"title"`                                   // Product name, 1-32 characters
	Description               string                `json:"description"`                             // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string                `json:"provider_token"`                          // Payment provider token, obtained via @BotFather
	Currency                  string                `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice       `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int64                 `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64               `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                `json:"start_parameter,omitempty"`               // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int64                 `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int64                 `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int64                 `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool                  `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool                  `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to provider
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to provider
	IsFlexible                bool                  `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method
	DisableNotification       bool                  `json:"disable_notification,omitempty"`          // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent            bool                  `json:"protect_content,omitempty"`               // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId          int64                 `json:"reply_to_message_id,omitempty"`           // If the message is a reply, ID of the original message
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`   // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

// sendInvoice is used to send invoices. On success, the sent Message is returned.
func (api *API) SendInvoice(payload *SendInvoice) (*Message, error) {
	return callJson[*Message](api, "sendInvoice", payload)
}

// createInvoiceLink is used to create a link for an invoice. Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
	Title                     string          `json:"title"`                                   // Product name, 1-32 characters
	Description               string          `json:"description"`                             // Product description, 1-255 characters
	Payload                   string          `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string          `json:"provider_token"`                          // Payment provider token, obtained via BotFather
	Currency                  string          `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int64           `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64         `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64           `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int64           `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int64           `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool            `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool            `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool            `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool            `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider
	SendEmailToProvider       bool            `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider
	IsFlexible                bool            `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method
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
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code
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
	Currency                string     `json:"currency"`                     // Three-letter ISO 4217 currency code
	TotalAmount             int64      `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload"`              // Bot specified invoice payload
	ShippingOptionId        string     `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`   // Telegram payment identifier
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`   // Provider payment identifier
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	Id              string          `json:"id"`               // Unique query identifier
	From            User            `json:"from"`             // User who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User specified shipping address
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Id               string     `json:"id"`                           // Unique query identifier
	From             User       `json:"from"`                         // User who sent the query
	Currency         string     `json:"currency"`                     // Three-letter ISO 4217 currency code
	TotalAmount      int64      `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`              // Bot specified invoice payload
	ShippingOptionId string     `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
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
	Data        string          `json:"data,omitempty"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string          `json:"phone_number,omitempty"` // Optional. User's verified phone number, available only for “phone_number” type
	Email       string          `json:"email,omitempty"`        // Optional. User's verified email address, available only for “email” type
	Files       []*PassportFile `json:"files,omitempty"`        // Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation,omitempty"`  // Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string          `json:"hash"`                   // Base64-encoded element hash for using in PassportElementErrorUnspecified
}

// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrors struct {
	UserId int64                   `json:"user_id"` // User identifier
	Errors []*PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (api *API) SetPassportDataErrors(payload *SetPassportDataErrors) (bool, error) {
	return callJson[bool](api, "setPassportDataErrors", payload)
}

// PassportElementError represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:// PassportElementErrorDataField, PassportElementErrorFrontSide, PassportElementErrorReverseSide, PassportElementErrorSelfie, PassportElementErrorFile, PassportElementErrorFiles, PassportElementErrorTranslationFile, PassportElementErrorTranslationFiles, PassportElementErrorUnspecified
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
	ChatId                   int64                 `json:"chat_id"`                               // Unique identifier for the target chat
	MessageThreadId          int64                 `json:"message_thread_id,omitempty"`           // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	GameShortName            string                `json:"game_short_name"`                       // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	DisableNotification      bool                  `json:"disable_notification,omitempty"`        // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool                  `json:"protect_content,omitempty"`             // Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64                 `json:"reply_to_message_id,omitempty"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"` // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
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

// getGameHighScores is used to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.// // This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.//
type GetGameHighScores struct {
	UserId          int64  `json:"user_id"`                     // Target user id
	ChatId          int64  `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int64  `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
}

// getGameHighScores is used to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.// // This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.//
func (api *API) GetGameHighScores(payload *GetGameHighScores) ([]*GameHighScore, error) {
	return callJson[[]*GameHighScore](api, "getGameHighScores", payload)
}

// GameHighScore represents one row of the high scores table for a game.// And that's about all we've got for now.If you've got any questions, please check out our Bot FAQ »
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
