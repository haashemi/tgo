package botapi

// GetUpdatesParams contains the method's parameters
type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`          // Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          int64    `json:"limit,omitempty"`           // Optional. Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int64    `json:"timeout,omitempty"`         // Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

// GetUpdates Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
func (c *API) GetUpdates(params GetUpdatesParams) (data []*Update, err error) {
	return doHTTP[[]*Update](c.client, c.url, "getUpdates", params)
}

// SetWebhookParams contains the method's parameters
type SetWebhookParams struct {
	Url                string    `json:"url"`                            // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        InputFile `json:"certificate,omitempty"`          // Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string    `json:"ip_address,omitempty"`           // Optional. The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int64     `json:"max_connections,omitempty"`      // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string  `json:"allowed_updates,omitempty"`      // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool      `json:"drop_pending_updates,omitempty"` // Optional. Pass True to drop all pending updates
	SecretToken        string    `json:"secret_token,omitempty"`         // Optional. A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

func (d SetWebhookParams) HasUploadable() bool {
	return d.Certificate.NeedsUpload()
}

// SetWebhook Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
func (c *API) SetWebhook(params SetWebhookParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setWebhook", params)
}

// DeleteWebhookParams contains the method's parameters
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"` // Optional. Pass True to drop all pending updates
}

// DeleteWebhook Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (c *API) DeleteWebhook(params DeleteWebhookParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteWebhook", params)
}

// GetWebhookInfo Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (c *API) GetWebhookInfo() (data *WebhookInfo, err error) {
	return doHTTP[*WebhookInfo](c.client, c.url, "getWebhookInfo", nil)
}

// GetMe A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (c *API) GetMe() (data *User, err error) {
	return doHTTP[*User](c.client, c.url, "getMe", nil)
}

// LogOut Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (c *API) LogOut() (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "logOut", nil)
}

// Close Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (c *API) Close() (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "close", nil)
}

// SendMessageParams contains the method's parameters
type SendMessageParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                     string           `json:"text"`                                  // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities                 []*MessageEntity `json:"entities,omitempty"`                    // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`    // Optional. Disables link previews for links in this message
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (c *API) SendMessage(params SendMessageParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendMessage", params)
}

// ForwardMessageParams contains the method's parameters
type ForwardMessageParams struct {
	ChatId              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId     int64  `json:"message_thread_id,omitempty"`    // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId          ChatID `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool   `json:"disable_notification,omitempty"` // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Optional. Protects the contents of the forwarded message from forwarding and saving
	MessageId           int64  `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

// ForwardMessage Use this method to forward messages of any kind. Service messages can't be forwarded. On success, the sent Message is returned.
func (c *API) ForwardMessage(params ForwardMessageParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "forwardMessage", params)
}

// CopyMessageParams contains the method's parameters
type CopyMessageParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatId               ChatID           `json:"from_chat_id"`                          // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId                int64            `json:"message_id"`                            // Message identifier in the chat specified in from_chat_id
	Caption                  string           `json:"caption,omitempty"`                     // Optional. New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// CopyMessage Use this method to copy messages of any kind. Service messages and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (c *API) CopyMessage(params CopyMessageParams) (data *MessageId, err error) {
	return doHTTP[*MessageId](c.client, c.url, "copyMessage", params)
}

// SendPhotoParams contains the method's parameters
type SendPhotoParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Photo                    InputFile        `json:"photo"`                                 // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Optional. Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendPhotoParams) HasUploadable() bool {
	return d.Photo.NeedsUpload()
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (c *API) SendPhoto(params SendPhotoParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendPhoto", params)
}

// SendAudioParams contains the method's parameters
type SendAudioParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Audio                    InputFile        `json:"audio"`                                 // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Audio caption, 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int64            `json:"duration,omitempty"`                    // Optional. Duration of the audio in seconds
	Performer                string           `json:"performer,omitempty"`                   // Optional. Performer
	Title                    string           `json:"title,omitempty"`                       // Optional. Track name
	Thumb                    InputFile        `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendAudioParams) HasUploadable() bool {
	return d.Audio.NeedsUpload() || d.Thumb.NeedsUpload()
}

// SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (c *API) SendAudio(params SendAudioParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendAudio", params)
}

// SendDocumentParams contains the method's parameters
type SendDocumentParams struct {
	ChatId                      ChatID           `json:"chat_id"`                                  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId             int64            `json:"message_thread_id,omitempty"`              // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Document                    InputFile        `json:"document"`                                 // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Thumb                       InputFile        `json:"thumb,omitempty"`                          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             `json:"protect_content,omitempty"`                // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId            int64            `json:"reply_to_message_id,omitempty"`            // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply,omitempty"`    // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 InlineKeyboard   `json:"reply_markup,omitempty"`                   // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendDocumentParams) HasUploadable() bool {
	return d.Document.NeedsUpload() || d.Thumb.NeedsUpload()
}

// SendDocument Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (c *API) SendDocument(params SendDocumentParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendDocument", params)
}

// SendVideoParams contains the method's parameters
type SendVideoParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Video                    InputFile        `json:"video"`                                 // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files »
	Duration                 int64            `json:"duration,omitempty"`                    // Optional. Duration of sent video in seconds
	Width                    int64            `json:"width,omitempty"`                       // Optional. Video width
	Height                   int64            `json:"height,omitempty"`                      // Optional. Video height
	Thumb                    InputFile        `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Optional. Pass True if the video needs to be covered with a spoiler animation
	SupportsStreaming        bool             `json:"supports_streaming,omitempty"`          // Optional. Pass True if the uploaded video is suitable for streaming
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendVideoParams) HasUploadable() bool {
	return d.Video.NeedsUpload() || d.Thumb.NeedsUpload()
}

// SendVideo Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (c *API) SendVideo(params SendVideoParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendVideo", params)
}

// SendAnimationParams contains the method's parameters
type SendAnimationParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Animation                InputFile        `json:"animation"`                             // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files »
	Duration                 int64            `json:"duration,omitempty"`                    // Optional. Duration of sent animation in seconds
	Width                    int64            `json:"width,omitempty"`                       // Optional. Animation width
	Height                   int64            `json:"height,omitempty"`                      // Optional. Animation height
	Thumb                    InputFile        `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Optional. Pass True if the animation needs to be covered with a spoiler animation
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendAnimationParams) HasUploadable() bool {
	return d.Animation.NeedsUpload() || d.Thumb.NeedsUpload()
}

// SendAnimation Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (c *API) SendAnimation(params SendAnimationParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendAnimation", params)
}

// SendVoiceParams contains the method's parameters
type SendVoiceParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Voice                    InputFile        `json:"voice"`                                 // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Voice message caption, 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int64            `json:"duration,omitempty"`                    // Optional. Duration of the voice message in seconds
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendVoiceParams) HasUploadable() bool {
	return d.Voice.NeedsUpload()
}

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (c *API) SendVoice(params SendVoiceParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendVoice", params)
}

// SendVideoNoteParams contains the method's parameters
type SendVideoNoteParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	VideoNote                InputFile      `json:"video_note"`                            // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration                 int64          `json:"duration,omitempty"`                    // Optional. Duration of sent video in seconds
	Length                   int64          `json:"length,omitempty"`                      // Optional. Video width and height, i.e. diameter of the video message
	Thumb                    InputFile      `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendVideoNoteParams) HasUploadable() bool {
	return d.VideoNote.NeedsUpload() || d.Thumb.NeedsUpload()
}

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
func (c *API) SendVideoNote(params SendVideoNoteParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendVideoNote", params)
}

// SendMediaGroupParams contains the method's parameters
type SendMediaGroupParams struct {
	ChatId                   ChatID       `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64        `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Media                    []InputMedia `json:"media"`                                 // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification      bool         `json:"disable_notification,omitempty"`        // Optional. Sends messages silently. Users will receive a notification with no sound.
	ProtectContent           bool         `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent messages from forwarding and saving
	ReplyToMessageId         int64        `json:"reply_to_message_id,omitempty"`         // Optional. If the messages are a reply, ID of the original message
	AllowSendingWithoutReply bool         `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
}

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
func (c *API) SendMediaGroup(params SendMediaGroupParams) (data []*Message, err error) {
	return doHTTP[[]*Message](c.client, c.url, "sendMediaGroup", params)
}

// SendLocationParams contains the method's parameters
type SendLocationParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude                 float64        `json:"latitude"`                              // Latitude of the location
	Longitude                float64        `json:"longitude"`                             // Longitude of the location
	HorizontalAccuracy       float64        `json:"horizontal_accuracy,omitempty"`         // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod               int64          `json:"live_period,omitempty"`                 // Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	Heading                  int64          `json:"heading,omitempty"`                     // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius     int64          `json:"proximity_alert_radius,omitempty"`      // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendLocation Use this method to send point on the map. On success, the sent Message is returned.
func (c *API) SendLocation(params SendLocationParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendLocation", params)
}

// EditMessageLiveLocationParams contains the method's parameters
type EditMessageLiveLocationParams struct {
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               `json:"latitude"`                         // Latitude of new location
	Longitude            float64               `json:"longitude"`                        // Longitude of new location
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int64                 `json:"heading,omitempty"`                // Optional. Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional. A JSON-serialized object for a new inline keyboard.
}

// EditMessageLiveLocation Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) EditMessageLiveLocation(params EditMessageLiveLocationParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "editMessageLiveLocation", params)
}

// StopMessageLiveLocationParams contains the method's parameters
type StopMessageLiveLocationParams struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for a new inline keyboard.
}

// StopMessageLiveLocation Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) StopMessageLiveLocation(params StopMessageLiveLocationParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "stopMessageLiveLocation", params)
}

// SendVenueParams contains the method's parameters
type SendVenueParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude                 float64        `json:"latitude"`                              // Latitude of the venue
	Longitude                float64        `json:"longitude"`                             // Longitude of the venue
	Title                    string         `json:"title"`                                 // Name of the venue
	Address                  string         `json:"address"`                               // Address of the venue
	FoursquareId             string         `json:"foursquare_id,omitempty"`               // Optional. Foursquare identifier of the venue
	FoursquareType           string         `json:"foursquare_type,omitempty"`             // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId            string         `json:"google_place_id,omitempty"`             // Optional. Google Places identifier of the venue
	GooglePlaceType          string         `json:"google_place_type,omitempty"`           // Optional. Google Places type of the venue. (See supported types.)
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
func (c *API) SendVenue(params SendVenueParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendVenue", params)
}

// SendContactParams contains the method's parameters
type SendContactParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	PhoneNumber              string         `json:"phone_number"`                          // Contact's phone number
	FirstName                string         `json:"first_name"`                            // Contact's first name
	LastName                 string         `json:"last_name,omitempty"`                   // Optional. Contact's last name
	Vcard                    string         `json:"vcard,omitempty"`                       // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendContact Use this method to send phone contacts. On success, the sent Message is returned.
func (c *API) SendContact(params SendContactParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendContact", params)
}

// SendPollParams contains the method's parameters
type SendPollParams struct {
	ChatId                   ChatID           `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Question                 string           `json:"question"`                              // Poll question, 1-300 characters
	Options                  []string         `json:"options"`                               // A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous              bool             `json:"is_anonymous,omitempty"`                // Optional. True, if the poll needs to be anonymous, defaults to True
	Type                     string           `json:"type,omitempty"`                        // Optional. Poll type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers    bool             `json:"allows_multiple_answers,omitempty"`     // Optional. True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId          int64            `json:"correct_option_id,omitempty"`           // Optional. 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation              string           `json:"explanation,omitempty"`                 // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode     string           `json:"explanation_parse_mode,omitempty"`      // Optional. Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities      []*MessageEntity `json:"explanation_entities,omitempty"`        // Optional. A JSON-serialized list of special entities that appear in the poll explanation, which can be specified instead of parse_mode
	OpenPeriod               int64            `json:"open_period,omitempty"`                 // Optional. Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate                int64            `json:"close_date,omitempty"`                  // Optional. Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed                 bool             `json:"is_closed,omitempty"`                   // Optional. Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendPoll Use this method to send a native poll. On success, the sent Message is returned.
func (c *API) SendPoll(params SendPollParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendPoll", params)
}

// SendDiceParams contains the method's parameters
type SendDiceParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                    string         `json:"emoji,omitempty"`                       // Optional. Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64 for “”. Defaults to “”
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendDice Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (c *API) SendDice(params SendDiceParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendDice", params)
}

// SendChatActionParams contains the method's parameters
type SendChatActionParams struct {
	ChatId          ChatID `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId int64  `json:"message_thread_id,omitempty"` // Optional. Unique identifier for the target message thread; supergroups only
	Action          string `json:"action"`                      // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

// SendChatAction Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
func (c *API) SendChatAction(params SendChatActionParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "sendChatAction", params)
}

// GetUserProfilePhotosParams contains the method's parameters
type GetUserProfilePhotosParams struct {
	UserId int64 `json:"user_id"`          // Unique identifier of the target user
	Offset int64 `json:"offset,omitempty"` // Optional. Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int64 `json:"limit,omitempty"`  // Optional. Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (c *API) GetUserProfilePhotos(params GetUserProfilePhotosParams) (data *UserProfilePhotos, err error) {
	return doHTTP[*UserProfilePhotos](c.client, c.url, "getUserProfilePhotos", params)
}

// GetFileParams contains the method's parameters
type GetFileParams struct {
	FileId string `json:"file_id"` // File identifier to get information about
}

// GetFile Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
func (c *API) GetFile(params GetFileParams) (data *File, err error) {
	return doHTTP[*File](c.client, c.url, "getFile", params)
}

// BanChatMemberParams contains the method's parameters
type BanChatMemberParams struct {
	ChatId         ChatID `json:"chat_id"`                   // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId         int64  `json:"user_id"`                   // Unique identifier of the target user
	UntilDate      int64  `json:"until_date,omitempty"`      // Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool   `json:"revoke_messages,omitempty"` // Optional. Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

// BanChatMember Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) BanChatMember(params BanChatMemberParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "banChatMember", params)
}

// UnbanChatMemberParams contains the method's parameters
type UnbanChatMemberParams struct {
	ChatId       ChatID `json:"chat_id"`                  // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId       int64  `json:"user_id"`                  // Unique identifier of the target user
	OnlyIfBanned bool   `json:"only_if_banned,omitempty"` // Optional. Do nothing if the user is not banned
}

// UnbanChatMember Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (c *API) UnbanChatMember(params UnbanChatMemberParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unbanChatMember", params)
}

// RestrictChatMemberParams contains the method's parameters
type RestrictChatMemberParams struct {
	ChatId                        ChatID           `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId                        int64            `json:"user_id"`                                    // Unique identifier of the target user
	Permissions                   *ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new user permissions
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"` // Optional. Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UntilDate                     int64            `json:"until_date,omitempty"`                       // Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

// RestrictChatMember Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (c *API) RestrictChatMember(params RestrictChatMemberParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "restrictChatMember", params)
}

// PromoteChatMemberParams contains the method's parameters
type PromoteChatMemberParams struct {
	ChatId              ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId              int64  `json:"user_id"`                          // Unique identifier of the target user
	IsAnonymous         bool   `json:"is_anonymous,omitempty"`           // Optional. Pass True if the administrator's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat,omitempty"`        // Optional. Pass True if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`      // Optional. Pass True if the administrator can create channel posts, channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`      // Optional. Pass True if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages   bool   `json:"can_delete_messages,omitempty"`    // Optional. Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats,omitempty"` // Optional. Pass True if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members,omitempty"`   // Optional. Pass True if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool   `json:"can_promote_members,omitempty"`    // Optional. Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool   `json:"can_change_info,omitempty"`        // Optional. Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users,omitempty"`       // Optional. Pass True if the administrator can invite new users to the chat
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`       // Optional. Pass True if the administrator can pin messages, supergroups only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`      // Optional. Pass True if the user is allowed to create, rename, close, and reopen forum topics, supergroups only
}

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (c *API) PromoteChatMember(params PromoteChatMemberParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "promoteChatMember", params)
}

// SetChatAdministratorCustomTitleParams contains the method's parameters
type SetChatAdministratorCustomTitleParams struct {
	ChatId      ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int64  `json:"user_id"`      // Unique identifier of the target user
	CustomTitle string `json:"custom_title"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

// SetChatAdministratorCustomTitle Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (c *API) SetChatAdministratorCustomTitle(params SetChatAdministratorCustomTitleParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatAdministratorCustomTitle", params)
}

// BanChatSenderChatParams contains the method's parameters
type BanChatSenderChatParams struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) BanChatSenderChat(params BanChatSenderChatParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "banChatSenderChat", params)
}

// UnbanChatSenderChatParams contains the method's parameters
type UnbanChatSenderChatParams struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// UnbanChatSenderChat Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) UnbanChatSenderChat(params UnbanChatSenderChatParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unbanChatSenderChat", params)
}

// SetChatPermissionsParams contains the method's parameters
type SetChatPermissionsParams struct {
	ChatId                        ChatID           `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions                   *ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new default chat permissions
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"` // Optional. Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
}

// SetChatPermissions Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (c *API) SetChatPermissions(params SetChatPermissionsParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatPermissions", params)
}

// ExportChatInviteLinkParams contains the method's parameters
type ExportChatInviteLinkParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
func (c *API) ExportChatInviteLink(params ExportChatInviteLinkParams) (data string, err error) {
	return doHTTP[string](c.client, c.url, "exportChatInviteLink", params)
}

// CreateChatInviteLinkParams contains the method's parameters
type CreateChatInviteLinkParams struct {
	ChatId             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`                 // Optional. Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Optional. Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // Optional. True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// CreateChatInviteLink Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (c *API) CreateChatInviteLink(params CreateChatInviteLinkParams) (data *ChatInviteLink, err error) {
	return doHTTP[*ChatInviteLink](c.client, c.url, "createChatInviteLink", params)
}

// EditChatInviteLinkParams contains the method's parameters
type EditChatInviteLinkParams struct {
	ChatId             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink         string `json:"invite_link"`                    // The invite link to edit
	Name               string `json:"name,omitempty"`                 // Optional. Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Optional. Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // Optional. True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// EditChatInviteLink Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (c *API) EditChatInviteLink(params EditChatInviteLinkParams) (data *ChatInviteLink, err error) {
	return doHTTP[*ChatInviteLink](c.client, c.url, "editChatInviteLink", params)
}

// RevokeChatInviteLinkParams contains the method's parameters
type RevokeChatInviteLinkParams struct {
	ChatId     ChatID `json:"chat_id"`     // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"` // The invite link to revoke
}

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
func (c *API) RevokeChatInviteLink(params RevokeChatInviteLinkParams) (data *ChatInviteLink, err error) {
	return doHTTP[*ChatInviteLink](c.client, c.url, "revokeChatInviteLink", params)
}

// ApproveChatJoinRequestParams contains the method's parameters
type ApproveChatJoinRequestParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// ApproveChatJoinRequest Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (c *API) ApproveChatJoinRequest(params ApproveChatJoinRequestParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "approveChatJoinRequest", params)
}

// DeclineChatJoinRequestParams contains the method's parameters
type DeclineChatJoinRequestParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// DeclineChatJoinRequest Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (c *API) DeclineChatJoinRequest(params DeclineChatJoinRequestParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "declineChatJoinRequest", params)
}

// SetChatPhotoParams contains the method's parameters
type SetChatPhotoParams struct {
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

func (d SetChatPhotoParams) HasUploadable() bool {
	return d.Photo.NeedsUpload()
}

// SetChatPhoto Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) SetChatPhoto(params SetChatPhotoParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatPhoto", params)
}

// DeleteChatPhotoParams contains the method's parameters
type DeleteChatPhotoParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// DeleteChatPhoto Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) DeleteChatPhoto(params DeleteChatPhotoParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteChatPhoto", params)
}

// SetChatTitleParams contains the method's parameters
type SetChatTitleParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-128 characters
}

// SetChatTitle Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) SetChatTitle(params SetChatTitleParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatTitle", params)
}

// SetChatDescriptionParams contains the method's parameters
type SetChatDescriptionParams struct {
	ChatId      ChatID `json:"chat_id"`               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string `json:"description,omitempty"` // Optional. New chat description, 0-255 characters
}

// SetChatDescription Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (c *API) SetChatDescription(params SetChatDescriptionParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatDescription", params)
}

// PinChatMessageParams contains the method's parameters
type PinChatMessageParams struct {
	ChatId              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId           int64  `json:"message_id"`                     // Identifier of a message to pin
	DisableNotification bool   `json:"disable_notification,omitempty"` // Optional. Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

// PinChatMessage Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (c *API) PinChatMessage(params PinChatMessageParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "pinChatMessage", params)
}

// UnpinChatMessageParams contains the method's parameters
type UnpinChatMessageParams struct {
	ChatId    ChatID `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id,omitempty"` // Optional. Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

// UnpinChatMessage Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (c *API) UnpinChatMessage(params UnpinChatMessageParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unpinChatMessage", params)
}

// UnpinAllChatMessagesParams contains the method's parameters
type UnpinAllChatMessagesParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// UnpinAllChatMessages Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (c *API) UnpinAllChatMessages(params UnpinAllChatMessagesParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unpinAllChatMessages", params)
}

// LeaveChatParams contains the method's parameters
type LeaveChatParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (c *API) LeaveChat(params LeaveChatParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "leaveChat", params)
}

// GetChatParams contains the method's parameters
type GetChatParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (c *API) GetChat(params GetChatParams) (data *Chat, err error) {
	return doHTTP[*Chat](c.client, c.url, "getChat", params)
}

// GetChatAdministratorsParams contains the method's parameters
type GetChatAdministratorsParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatAdministrators Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
func (c *API) GetChatAdministrators(params GetChatAdministratorsParams) (data []*ChatMember, err error) {
	return doHTTP[[]*ChatMember](c.client, c.url, "getChatAdministrators", params)
}

// GetChatMemberCountParams contains the method's parameters
type GetChatMemberCountParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMemberCount Use this method to get the number of members in a chat. Returns Int on success.
func (c *API) GetChatMemberCount(params GetChatMemberCountParams) (data int64, err error) {
	return doHTTP[int64](c.client, c.url, "getChatMemberCount", params)
}

// GetChatMemberParams contains the method's parameters
type GetChatMemberParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// GetChatMember Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
func (c *API) GetChatMember(params GetChatMemberParams) (data *ChatMember, err error) {
	return doHTTP[*ChatMember](c.client, c.url, "getChatMember", params)
}

// SetChatStickerSetParams contains the method's parameters
type SetChatStickerSetParams struct {
	ChatId         ChatID `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

// SetChatStickerSet Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (c *API) SetChatStickerSet(params SetChatStickerSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatStickerSet", params)
}

// DeleteChatStickerSetParams contains the method's parameters
type DeleteChatStickerSetParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (c *API) DeleteChatStickerSet(params DeleteChatStickerSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteChatStickerSet", params)
}

// GetForumTopicIconStickers Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
func (c *API) GetForumTopicIconStickers() (data []*Sticker, err error) {
	return doHTTP[[]*Sticker](c.client, c.url, "getForumTopicIconStickers", nil)
}

// CreateForumTopicParams contains the method's parameters
type CreateForumTopicParams struct {
	ChatId            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name              string `json:"name"`                           // Topic name, 1-128 characters
	IconColor         int64  `json:"icon_color,omitempty"`           // Optional. Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
}

// CreateForumTopic Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
func (c *API) CreateForumTopic(params CreateForumTopicParams) (data *ForumTopic, err error) {
	return doHTTP[*ForumTopic](c.client, c.url, "createForumTopic", params)
}

// EditForumTopicParams contains the method's parameters
type EditForumTopicParams struct {
	ChatId            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId   int64  `json:"message_thread_id"`              // Unique identifier for the target message thread of the forum topic
	Name              string `json:"name,omitempty"`                 // Optional. New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

// EditForumTopic Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (c *API) EditForumTopic(params EditForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "editForumTopic", params)
}

// CloseForumTopicParams contains the method's parameters
type CloseForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// CloseForumTopic Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (c *API) CloseForumTopic(params CloseForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "closeForumTopic", params)
}

// ReopenForumTopicParams contains the method's parameters
type ReopenForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// ReopenForumTopic Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (c *API) ReopenForumTopic(params ReopenForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "reopenForumTopic", params)
}

// DeleteForumTopicParams contains the method's parameters
type DeleteForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// DeleteForumTopic Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (c *API) DeleteForumTopic(params DeleteForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteForumTopic", params)
}

// UnpinAllForumTopicMessagesParams contains the method's parameters
type UnpinAllForumTopicMessagesParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// UnpinAllForumTopicMessages Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (c *API) UnpinAllForumTopicMessages(params UnpinAllForumTopicMessagesParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unpinAllForumTopicMessages", params)
}

// EditGeneralForumTopicParams contains the method's parameters
type EditGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // New topic name, 1-128 characters
}

// EditGeneralForumTopic Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
func (c *API) EditGeneralForumTopic(params EditGeneralForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "editGeneralForumTopic", params)
}

// CloseGeneralForumTopicParams contains the method's parameters
type CloseGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// CloseGeneralForumTopic Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (c *API) CloseGeneralForumTopic(params CloseGeneralForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "closeGeneralForumTopic", params)
}

// ReopenGeneralForumTopicParams contains the method's parameters
type ReopenGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// ReopenGeneralForumTopic Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (c *API) ReopenGeneralForumTopic(params ReopenGeneralForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "reopenGeneralForumTopic", params)
}

// HideGeneralForumTopicParams contains the method's parameters
type HideGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// HideGeneralForumTopic Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (c *API) HideGeneralForumTopic(params HideGeneralForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "hideGeneralForumTopic", params)
}

// UnhideGeneralForumTopicParams contains the method's parameters
type UnhideGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// UnhideGeneralForumTopic Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (c *API) UnhideGeneralForumTopic(params UnhideGeneralForumTopicParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "unhideGeneralForumTopic", params)
}

// AnswerCallbackQueryParams contains the method's parameters
type AnswerCallbackQueryParams struct {
	CallbackQueryId string `json:"callback_query_id"`    // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`       // Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert,omitempty"` // Optional. If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url,omitempty"`        // Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int64  `json:"cache_time,omitempty"` // Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

// AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (c *API) AnswerCallbackQuery(params AnswerCallbackQueryParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "answerCallbackQuery", params)
}

// SetMyCommandsParams contains the method's parameters
type SetMyCommandsParams struct {
	Commands     []*BotCommand    `json:"commands"`                // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// SetMyCommands Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (c *API) SetMyCommands(params SetMyCommandsParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setMyCommands", params)
}

// DeleteMyCommandsParams contains the method's parameters
type DeleteMyCommandsParams struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// DeleteMyCommands Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (c *API) DeleteMyCommands(params DeleteMyCommandsParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteMyCommands", params)
}

// GetMyCommandsParams contains the method's parameters
type GetMyCommandsParams struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code or an empty string
}

// GetMyCommands Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
func (c *API) GetMyCommands(params GetMyCommandsParams) (data []*BotCommand, err error) {
	return doHTTP[[]*BotCommand](c.client, c.url, "getMyCommands", params)
}

// SetChatMenuButtonParams contains the method's parameters
type SetChatMenuButtonParams struct {
	ChatId     int64       `json:"chat_id,omitempty"`     // Optional. Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton *MenuButton `json:"menu_button,omitempty"` // Optional. A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}

// SetChatMenuButton Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (c *API) SetChatMenuButton(params SetChatMenuButtonParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setChatMenuButton", params)
}

// GetChatMenuButtonParams contains the method's parameters
type GetChatMenuButtonParams struct {
	ChatId int64 `json:"chat_id,omitempty"` // Optional. Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
}

// GetChatMenuButton Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
func (c *API) GetChatMenuButton(params GetChatMenuButtonParams) (data *MenuButton, err error) {
	return doHTTP[*MenuButton](c.client, c.url, "getChatMenuButton", params)
}

// SetMyDefaultAdministratorRightsParams contains the method's parameters
type SetMyDefaultAdministratorRightsParams struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`       // Optional. A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	ForChannels bool                     `json:"for_channels,omitempty"` // Optional. Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
}

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are are free to modify the list before adding the bot. Returns True on success.
func (c *API) SetMyDefaultAdministratorRights(params SetMyDefaultAdministratorRightsParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setMyDefaultAdministratorRights", params)
}

// GetMyDefaultAdministratorRightsParams contains the method's parameters
type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"` // Optional. Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
}

// GetMyDefaultAdministratorRights Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
func (c *API) GetMyDefaultAdministratorRights(params GetMyDefaultAdministratorRightsParams) (data *ChatAdministratorRights, err error) {
	return doHTTP[*ChatAdministratorRights](c.client, c.url, "getMyDefaultAdministratorRights", params)
}

// EditMessageTextParams contains the method's parameters
type EditMessageTextParams struct {
	ChatId                ChatID                `json:"chat_id,omitempty"`                  // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int64                 `json:"message_id,omitempty"`               // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string                `json:"inline_message_id,omitempty"`        // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string                `json:"text"`                               // New text of the message, 1-4096 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity      `json:"entities,omitempty"`                 // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. A JSON-serialized object for an inline keyboard.
}

// EditMessageText Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) EditMessageText(params EditMessageTextParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "editMessageText", params)
}

// EditMessageCaptionParams contains the method's parameters
type EditMessageCaptionParams struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string                `json:"caption,omitempty"`           // Optional. New caption of the message, 0-1024 characters after entities parsing
	ParseMode       ParseMode             `json:"parse_mode,omitempty"`        // Optional. Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities []*MessageEntity      `json:"caption_entities,omitempty"`  // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for an inline keyboard.
}

// EditMessageCaption Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) EditMessageCaption(params EditMessageCaptionParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "editMessageCaption", params)
}

// EditMessageMediaParams contains the method's parameters
type EditMessageMediaParams struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Media           *InputMedia           `json:"media"`                       // A JSON-serialized object for a new media content of the message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for a new inline keyboard.
}

// EditMessageMedia Use this method to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) EditMessageMedia(params EditMessageMediaParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "editMessageMedia", params)
}

// EditMessageReplyMarkupParams contains the method's parameters
type EditMessageReplyMarkupParams struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for an inline keyboard.
}

// EditMessageReplyMarkup Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c *API) EditMessageReplyMarkup(params EditMessageReplyMarkupParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "editMessageReplyMarkup", params)
}

// StopPollParams contains the method's parameters
type StopPollParams struct {
	ChatId      ChatID                `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int64                 `json:"message_id"`             // Identifier of the original message with the poll
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. A JSON-serialized object for a new message inline keyboard.
}

// StopPoll Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
func (c *API) StopPoll(params StopPollParams) (data *Poll, err error) {
	return doHTTP[*Poll](c.client, c.url, "stopPoll", params)
}

// DeleteMessageParams contains the method's parameters
type DeleteMessageParams struct {
	ChatId    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id"` // Identifier of the message to delete
}

// DeleteMessage Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can't be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
func (c *API) DeleteMessage(params DeleteMessageParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteMessage", params)
}

// SendStickerParams contains the method's parameters
type SendStickerParams struct {
	ChatId                   ChatID         `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId          int64          `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Sticker                  InputFile      `json:"sticker"`                               // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	DisableNotification      bool           `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool           `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64          `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              InlineKeyboard `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (d SendStickerParams) HasUploadable() bool {
	return d.Sticker.NeedsUpload()
}

// SendSticker Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
func (c *API) SendSticker(params SendStickerParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendSticker", params)
}

// GetStickerSetParams contains the method's parameters
type GetStickerSetParams struct {
	Name string `json:"name"` // Name of the sticker set
}

// GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
func (c *API) GetStickerSet(params GetStickerSetParams) (data *StickerSet, err error) {
	return doHTTP[*StickerSet](c.client, c.url, "getStickerSet", params)
}

// GetCustomEmojiStickersParams contains the method's parameters
type GetCustomEmojiStickersParams struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"` // List of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

// GetCustomEmojiStickers Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (c *API) GetCustomEmojiStickers(params GetCustomEmojiStickersParams) (data []*Sticker, err error) {
	return doHTTP[[]*Sticker](c.client, c.url, "getCustomEmojiStickers", params)
}

// UploadStickerFileParams contains the method's parameters
type UploadStickerFileParams struct {
	UserId     int64     `json:"user_id"`     // User identifier of sticker file owner
	PngSticker InputFile `json:"png_sticker"` // PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More information on Sending Files »
}

func (d UploadStickerFileParams) HasUploadable() bool {
	return d.PngSticker.NeedsUpload()
}

// UploadStickerFile Use this method to upload a .PNG file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
func (c *API) UploadStickerFile(params UploadStickerFileParams) (data *File, err error) {
	return doHTTP[*File](c.client, c.url, "uploadStickerFile", params)
}

// CreateNewStickerSetParams contains the method's parameters
type CreateNewStickerSetParams struct {
	UserId       int64         `json:"user_id"`                 // User identifier of created sticker set owner
	Name         string        `json:"name"`                    // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Title        string        `json:"title"`                   // Sticker set title, 1-64 characters
	PngSticker   InputFile     `json:"png_sticker,omitempty"`   // Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	TgsSticker   InputFile     `json:"tgs_sticker,omitempty"`   // Optional. TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#animated-sticker-requirements for technical requirements
	WebmSticker  InputFile     `json:"webm_sticker,omitempty"`  // Optional. WEBM video with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#video-sticker-requirements for technical requirements
	StickerType  string        `json:"sticker_type,omitempty"`  // Optional. Type of stickers in the set, pass “regular” or “mask”. Custom emoji sticker sets can't be created via the Bot API at the moment. By default, a regular sticker set is created.
	Emojis       string        `json:"emojis"`                  // One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. A JSON-serialized object for position where the mask should be placed on faces
}

func (d CreateNewStickerSetParams) HasUploadable() bool {
	return d.PngSticker.NeedsUpload() || d.TgsSticker.NeedsUpload() || d.WebmSticker.NeedsUpload()
}

// CreateNewStickerSet Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Returns True on success.
func (c *API) CreateNewStickerSet(params CreateNewStickerSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "createNewStickerSet", params)
}

// AddStickerToSetParams contains the method's parameters
type AddStickerToSetParams struct {
	UserId       int64         `json:"user_id"`                 // User identifier of sticker set owner
	Name         string        `json:"name"`                    // Sticker set name
	PngSticker   InputFile     `json:"png_sticker,omitempty"`   // Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	TgsSticker   InputFile     `json:"tgs_sticker,omitempty"`   // Optional. TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#animated-sticker-requirements for technical requirements
	WebmSticker  InputFile     `json:"webm_sticker,omitempty"`  // Optional. WEBM video with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#video-sticker-requirements for technical requirements
	Emojis       string        `json:"emojis"`                  // One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. A JSON-serialized object for position where the mask should be placed on faces
}

func (d AddStickerToSetParams) HasUploadable() bool {
	return d.PngSticker.NeedsUpload() || d.TgsSticker.NeedsUpload() || d.WebmSticker.NeedsUpload()
}

// AddStickerToSet Use this method to add a new sticker to a set created by the bot. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Animated stickers can be added to animated sticker sets and only to them. Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
func (c *API) AddStickerToSet(params AddStickerToSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "addStickerToSet", params)
}

// SetStickerPositionInSetParams contains the method's parameters
type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int64  `json:"position"` // New sticker position in the set, zero-based
}

// SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
func (c *API) SetStickerPositionInSet(params SetStickerPositionInSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setStickerPositionInSet", params)
}

// DeleteStickerFromSetParams contains the method's parameters
type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

// DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (c *API) DeleteStickerFromSet(params DeleteStickerFromSetParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "deleteStickerFromSet", params)
}

// SetStickerSetThumbParams contains the method's parameters
type SetStickerSetThumbParams struct {
	Name   string    `json:"name"`            // Sticker set name
	UserId int64     `json:"user_id"`         // User identifier of the sticker set owner
	Thumb  InputFile `json:"thumb,omitempty"` // Optional. A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and height exactly 100px, or a TGS animation with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#animated-sticker-requirements for animated sticker technical requirements, or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-sticker-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated sticker set thumbnails can't be uploaded via HTTP URL.
}

func (d SetStickerSetThumbParams) HasUploadable() bool {
	return d.Thumb.NeedsUpload()
}

// SetStickerSetThumb Use this method to set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets only. Video thumbnails can be set only for video sticker sets only. Returns True on success.
func (c *API) SetStickerSetThumb(params SetStickerSetThumbParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setStickerSetThumb", params)
}

// AnswerInlineQueryParams contains the method's parameters
type AnswerInlineQueryParams struct {
	InlineQueryId     string               `json:"inline_query_id"`               // Unique identifier for the answered query
	Results           []*InlineQueryResult `json:"results"`                       // A JSON-serialized array of results for the inline query
	CacheTime         int64                `json:"cache_time,omitempty"`          // Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool                 `json:"is_personal,omitempty"`         // Optional. Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string               `json:"next_offset,omitempty"`         // Optional. Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	SwitchPmText      string               `json:"switch_pm_text,omitempty"`      // Optional. If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string               `json:"switch_pm_parameter,omitempty"` // Optional. Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

// AnswerInlineQuery Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
func (c *API) AnswerInlineQuery(params AnswerInlineQueryParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "answerInlineQuery", params)
}

// AnswerWebAppQueryParams contains the method's parameters
type AnswerWebAppQueryParams struct {
	WebAppQueryId string             `json:"web_app_query_id"` // Unique identifier for the query to be answered
	Result        *InlineQueryResult `json:"result"`           // A JSON-serialized object describing the message to be sent
}

// AnswerWebAppQuery Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
func (c *API) AnswerWebAppQuery(params AnswerWebAppQueryParams) (data *SentWebAppMessage, err error) {
	return doHTTP[*SentWebAppMessage](c.client, c.url, "answerWebAppQuery", params)
}

// SendInvoiceParams contains the method's parameters
type SendInvoiceParams struct {
	ChatId                    ChatID                `json:"chat_id"`                                 // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId           int64                 `json:"message_thread_id,omitempty"`             // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Title                     string                `json:"title"`                                   // Product name, 1-32 characters
	Description               string                `json:"description"`                             // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string                `json:"provider_token"`                          // Payment provider token, obtained via @BotFather
	Currency                  string                `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice       `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int64                 `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64               `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                `json:"start_parameter,omitempty"`               // Optional. Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                `json:"provider_data,omitempty"`                 // Optional. JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int64                 `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int64                 `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int64                 `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool                  `json:"need_name,omitempty"`                     // Optional. Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`             // Optional. Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool                  `json:"need_email,omitempty"`                    // Optional. Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`         // Optional. Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True if the user's phone number should be sent to provider
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`        // Optional. Pass True if the user's email address should be sent to provider
	IsFlexible                bool                  `json:"is_flexible,omitempty"`                   // Optional. Pass True if the final price depends on the shipping method
	DisableNotification       bool                  `json:"disable_notification,omitempty"`          // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent            bool                  `json:"protect_content,omitempty"`               // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId          int64                 `json:"reply_to_message_id,omitempty"`           // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`   // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

// SendInvoice Use this method to send invoices. On success, the sent Message is returned.
func (c *API) SendInvoice(params SendInvoiceParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendInvoice", params)
}

// CreateInvoiceLinkParams contains the method's parameters
type CreateInvoiceLinkParams struct {
	Title                     string          `json:"title"`                                   // Product name, 1-32 characters
	Description               string          `json:"description"`                             // Product description, 1-255 characters
	Payload                   string          `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string          `json:"provider_token"`                          // Payment provider token, obtained via BotFather
	Currency                  string          `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int64           `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64         `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data,omitempty"`                 // Optional. JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64           `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int64           `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int64           `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool            `json:"need_name,omitempty"`                     // Optional. Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool            `json:"need_phone_number,omitempty"`             // Optional. Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool            `json:"need_email,omitempty"`                    // Optional. Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool            `json:"need_shipping_address,omitempty"`         // Optional. Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True if the user's phone number should be sent to the provider
	SendEmailToProvider       bool            `json:"send_email_to_provider,omitempty"`        // Optional. Pass True if the user's email address should be sent to the provider
	IsFlexible                bool            `json:"is_flexible,omitempty"`                   // Optional. Pass True if the final price depends on the shipping method
}

// CreateInvoiceLink Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (c *API) CreateInvoiceLink(params CreateInvoiceLinkParams) (data string, err error) {
	return doHTTP[string](c.client, c.url, "createInvoiceLink", params)
}

// AnswerShippingQueryParams contains the method's parameters
type AnswerShippingQueryParams struct {
	ShippingQueryId string            `json:"shipping_query_id"`          // Unique identifier for the query to be answered
	Ok              bool              `json:"ok"`                         // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"` // Optional. Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string            `json:"error_message,omitempty"`    // Optional. Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

// AnswerShippingQuery If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
func (c *API) AnswerShippingQuery(params AnswerShippingQueryParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "answerShippingQuery", params)
}

// AnswerPreCheckoutQueryParams contains the method's parameters
type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`   // Unique identifier for the query to be answered
	Ok                 bool   `json:"ok"`                      // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string `json:"error_message,omitempty"` // Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

// AnswerPreCheckoutQuery Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (c *API) AnswerPreCheckoutQuery(params AnswerPreCheckoutQueryParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "answerPreCheckoutQuery", params)
}

// SetPassportDataErrorsParams contains the method's parameters
type SetPassportDataErrorsParams struct {
	UserId int64                   `json:"user_id"` // User identifier
	Errors []*PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
func (c *API) SetPassportDataErrors(params SetPassportDataErrorsParams) (data bool, err error) {
	return doHTTP[bool](c.client, c.url, "setPassportDataErrors", params)
}

// SendGameParams contains the method's parameters
type SendGameParams struct {
	ChatId                   int64                 `json:"chat_id"`                               // Unique identifier for the target chat
	MessageThreadId          int64                 `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	GameShortName            string                `json:"game_short_name"`                       // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	DisableNotification      bool                  `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool                  `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64                 `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                // Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

// SendGame Use this method to send a game. On success, the sent Message is returned.
func (c *API) SendGame(params SendGameParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "sendGame", params)
}

// SetGameScoreParams contains the method's parameters
type SetGameScoreParams struct {
	UserId             int64  `json:"user_id"`                        // User identifier
	Score              int64  `json:"score"`                          // New score, must be non-negative
	Force              bool   `json:"force,omitempty"`                // Optional. Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Optional. Pass True if the game message should not be automatically edited to include the current scoreboard
	ChatId             int64  `json:"chat_id,omitempty"`              // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId          int64  `json:"message_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId    string `json:"inline_message_id,omitempty"`    // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
}

// SetGameScore Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (c *API) SetGameScore(params SetGameScoreParams) (data *Message, err error) {
	return doHTTP[*Message](c.client, c.url, "setGameScore", params)
}

// GetGameHighScoresParams contains the method's parameters
type GetGameHighScoresParams struct {
	UserId          int64  `json:"user_id"`                     // Target user id
	ChatId          int64  `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int64  `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
}

// GetGameHighScores Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
func (c *API) GetGameHighScores(params GetGameHighScoresParams) (data []*GameHighScore, err error) {
	return doHTTP[[]*GameHighScore](c.client, c.url, "getGameHighScores", params)
}
