package tgo

// GetUpdatesOptions GetUpdatesOptions contains GetUpdates's optional params
type GetUpdatesOptions struct {
	Offset         int64    `json:"offset,omitempty"`          // Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          int64    `json:"limit,omitempty"`           // Optional. Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int64    `json:"timeout,omitempty"`         // Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

// getUpdatesParams getUpdatesParams contains GetUpdates's params
type getUpdatesParams struct {
	*GetUpdatesOptions
}

// GetUpdates Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
func (b *Bot) GetUpdates(optionalParams *GetUpdatesOptions) ([]*Update, error) {
	params := &getUpdatesParams{}

	params.GetUpdatesOptions = optionalParams

	return doHTTP[[]*Update](b.client, b.url, "getUpdates", params)
}

// SetWebhookOptions SetWebhookOptions contains SetWebhook's optional params
type SetWebhookOptions struct {
	Certificate        InputFile `json:"certificate,omitempty"`          // Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string    `json:"ip_address,omitempty"`           // Optional. The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int64     `json:"max_connections,omitempty"`      // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string  `json:"allowed_updates,omitempty"`      // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool      `json:"drop_pending_updates,omitempty"` // Optional. Pass True to drop all pending updates
	SecretToken        string    `json:"secret_token,omitempty"`         // Optional. A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

// setWebhookParams setWebhookParams contains SetWebhook's params
type setWebhookParams struct {
	*SetWebhookOptions
	Url string `json:"url"` // HTTPS URL to send updates to. Use an empty string to remove webhook integration
}

// SetWebhook Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
func (b *Bot) SetWebhook(url string, optionalParams *SetWebhookOptions) (bool, error) {
	params := &setWebhookParams{}

	params.Url = url
	params.SetWebhookOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setWebhook", params)
}

func (params *setWebhookParams) HasUploadable() bool {
	return params.Certificate.NeedsUpload()
}

// DeleteWebhookOptions DeleteWebhookOptions contains DeleteWebhook's optional params
type DeleteWebhookOptions struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"` // Optional. Pass True to drop all pending updates
}

// deleteWebhookParams deleteWebhookParams contains DeleteWebhook's params
type deleteWebhookParams struct {
	*DeleteWebhookOptions
}

// DeleteWebhook Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (b *Bot) DeleteWebhook(optionalParams *DeleteWebhookOptions) (bool, error) {
	params := &deleteWebhookParams{}

	params.DeleteWebhookOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "deleteWebhook", params)
}

// getWebhookInfoParams getWebhookInfoParams contains GetWebhookInfo's params
type getWebhookInfoParams struct {
}

// GetWebhookInfo Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (b *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	params := &getWebhookInfoParams{}

	return doHTTP[*WebhookInfo](b.client, b.url, "getWebhookInfo", params)
}

// getMeParams getMeParams contains GetMe's params
type getMeParams struct {
}

// GetMe A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (b *Bot) GetMe() (*User, error) {
	params := &getMeParams{}

	return doHTTP[*User](b.client, b.url, "getMe", params)
}

// logOutParams logOutParams contains LogOut's params
type logOutParams struct {
}

// LogOut Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (b *Bot) LogOut() (bool, error) {
	params := &logOutParams{}

	return doHTTP[bool](b.client, b.url, "logOut", params)
}

// closeParams closeParams contains Close's params
type closeParams struct {
}

// Close Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (b *Bot) Close() (bool, error) {
	params := &closeParams{}

	return doHTTP[bool](b.client, b.url, "close", params)
}

// SendMessageOptions SendMessageOptions contains SendMessage's optional params
type SendMessageOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities                 []*MessageEntity `json:"entities,omitempty"`                    // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`    // Optional. Disables link previews for links in this message
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendMessageParams sendMessageParams contains SendMessage's params
type sendMessageParams struct {
	*SendMessageOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text   string `json:"text"`    // Text of the message to be sent, 1-4096 characters after entities parsing
}

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (b *Bot) SendMessage(chatId ChatID, text string, optionalParams *SendMessageOptions) (*Message, error) {
	params := &sendMessageParams{}

	params.ChatId = chatId
	params.Text = text
	params.SendMessageOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendMessage", params)
}

// ForwardMessageOptions ForwardMessageOptions contains ForwardMessage's optional params
type ForwardMessageOptions struct {
	MessageThreadId     int64 `json:"message_thread_id,omitempty"`    // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	DisableNotification bool  `json:"disable_notification,omitempty"` // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool  `json:"protect_content,omitempty"`      // Optional. Protects the contents of the forwarded message from forwarding and saving
}

// forwardMessageParams forwardMessageParams contains ForwardMessage's params
type forwardMessageParams struct {
	*ForwardMessageOptions
	ChatId     ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId ChatID `json:"from_chat_id"` // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId  int64  `json:"message_id"`   // Message identifier in the chat specified in from_chat_id
}

// ForwardMessage Use this method to forward messages of any kind. Service messages can't be forwarded. On success, the sent Message is returned.
func (b *Bot) ForwardMessage(chatId ChatID, fromChatId ChatID, messageId int64, optionalParams *ForwardMessageOptions) (*Message, error) {
	params := &forwardMessageParams{}

	params.ChatId = chatId
	params.FromChatId = fromChatId
	params.MessageId = messageId
	params.ForwardMessageOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "forwardMessage", params)
}

// CopyMessageOptions CopyMessageOptions contains CopyMessage's optional params
type CopyMessageOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Caption                  string           `json:"caption,omitempty"`                     // Optional. New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// copyMessageParams copyMessageParams contains CopyMessage's params
type copyMessageParams struct {
	*CopyMessageOptions
	ChatId     ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId ChatID `json:"from_chat_id"` // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId  int64  `json:"message_id"`   // Message identifier in the chat specified in from_chat_id
}

// CopyMessage Use this method to copy messages of any kind. Service messages and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (b *Bot) CopyMessage(chatId ChatID, fromChatId ChatID, messageId int64, optionalParams *CopyMessageOptions) (*MessageId, error) {
	params := &copyMessageParams{}

	params.ChatId = chatId
	params.FromChatId = fromChatId
	params.MessageId = messageId
	params.CopyMessageOptions = optionalParams

	return doHTTP[*MessageId](b.client, b.url, "copyMessage", params)
}

// SendPhotoOptions SendPhotoOptions contains SendPhoto's optional params
type SendPhotoOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	HasSpoiler               bool             `json:"has_spoiler,omitempty"`                 // Optional. Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendPhotoParams sendPhotoParams contains SendPhoto's params
type sendPhotoParams struct {
	*SendPhotoOptions
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile `json:"photo"`   // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (b *Bot) SendPhoto(chatId ChatID, photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	params := &sendPhotoParams{}

	params.ChatId = chatId
	params.Photo = photo
	params.SendPhotoOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendPhoto", params)
}

func (params *sendPhotoParams) HasUploadable() bool {
	return params.Photo.NeedsUpload()
}

// SendAudioOptions SendAudioOptions contains SendAudio's optional params
type SendAudioOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
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
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendAudioParams sendAudioParams contains SendAudio's params
type sendAudioParams struct {
	*SendAudioOptions
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Audio  InputFile `json:"audio"`   // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
}

// SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendAudio(chatId ChatID, audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	params := &sendAudioParams{}

	params.ChatId = chatId
	params.Audio = audio
	params.SendAudioOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendAudio", params)
}

func (params *sendAudioParams) HasUploadable() bool {
	return params.Audio.NeedsUpload() || params.Thumb.NeedsUpload()
}

// SendDocumentOptions SendDocumentOptions contains SendDocument's optional params
type SendDocumentOptions struct {
	MessageThreadId             int64            `json:"message_thread_id,omitempty"`              // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Thumb                       InputFile        `json:"thumb,omitempty"`                          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             `json:"protect_content,omitempty"`                // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId            int64            `json:"reply_to_message_id,omitempty"`            // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply,omitempty"`    // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`                   // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendDocumentParams sendDocumentParams contains SendDocument's params
type sendDocumentParams struct {
	*SendDocumentOptions
	ChatId   ChatID    `json:"chat_id"`  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Document InputFile `json:"document"` // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
}

// SendDocument Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendDocument(chatId ChatID, document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	params := &sendDocumentParams{}

	params.ChatId = chatId
	params.Document = document
	params.SendDocumentOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendDocument", params)
}

func (params *sendDocumentParams) HasUploadable() bool {
	return params.Document.NeedsUpload() || params.Thumb.NeedsUpload()
}

// SendVideoOptions SendVideoOptions contains SendVideo's optional params
type SendVideoOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
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
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendVideoParams sendVideoParams contains SendVideo's params
type sendVideoParams struct {
	*SendVideoOptions
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Video  InputFile `json:"video"`   // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files »
}

// SendVideo Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVideo(chatId ChatID, video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	params := &sendVideoParams{}

	params.ChatId = chatId
	params.Video = video
	params.SendVideoOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendVideo", params)
}

func (params *sendVideoParams) HasUploadable() bool {
	return params.Video.NeedsUpload() || params.Thumb.NeedsUpload()
}

// SendAnimationOptions SendAnimationOptions contains SendAnimation's optional params
type SendAnimationOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
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
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendAnimationParams sendAnimationParams contains SendAnimation's params
type sendAnimationParams struct {
	*SendAnimationOptions
	ChatId    ChatID    `json:"chat_id"`   // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Animation InputFile `json:"animation"` // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files »
}

// SendAnimation Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendAnimation(chatId ChatID, animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	params := &sendAnimationParams{}

	params.ChatId = chatId
	params.Animation = animation
	params.SendAnimationOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendAnimation", params)
}

func (params *sendAnimationParams) HasUploadable() bool {
	return params.Animation.NeedsUpload() || params.Thumb.NeedsUpload()
}

// SendVoiceOptions SendVoiceOptions contains SendVoice's optional params
type SendVoiceOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Caption                  string           `json:"caption,omitempty"`                     // Optional. Voice message caption, 0-1024 characters after entities parsing
	ParseMode                ParseMode        `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`            // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int64            `json:"duration,omitempty"`                    // Optional. Duration of the voice message in seconds
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendVoiceParams sendVoiceParams contains SendVoice's params
type sendVoiceParams struct {
	*SendVoiceOptions
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice  InputFile `json:"voice"`   // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
}

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVoice(chatId ChatID, voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	params := &sendVoiceParams{}

	params.ChatId = chatId
	params.Voice = voice
	params.SendVoiceOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendVoice", params)
}

func (params *sendVoiceParams) HasUploadable() bool {
	return params.Voice.NeedsUpload()
}

// SendVideoNoteOptions SendVideoNoteOptions contains SendVideoNote's optional params
type SendVideoNoteOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Duration                 int64       `json:"duration,omitempty"`                    // Optional. Duration of sent video in seconds
	Length                   int64       `json:"length,omitempty"`                      // Optional. Video width and height, i.e. diameter of the video message
	Thumb                    InputFile   `json:"thumb,omitempty"`                       // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendVideoNoteParams sendVideoNoteParams contains SendVideoNote's params
type sendVideoNoteParams struct {
	*SendVideoNoteOptions
	ChatId    ChatID    `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	VideoNote InputFile `json:"video_note"` // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files ». Sending video notes by a URL is currently unsupported
}

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
func (b *Bot) SendVideoNote(chatId ChatID, videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	params := &sendVideoNoteParams{}

	params.ChatId = chatId
	params.VideoNote = videoNote
	params.SendVideoNoteOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendVideoNote", params)
}

func (params *sendVideoNoteParams) HasUploadable() bool {
	return params.VideoNote.NeedsUpload() || params.Thumb.NeedsUpload()
}

// SendMediaGroupOptions SendMediaGroupOptions contains SendMediaGroup's optional params
type SendMediaGroupOptions struct {
	MessageThreadId          int64 `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	DisableNotification      bool  `json:"disable_notification,omitempty"`        // Optional. Sends messages silently. Users will receive a notification with no sound.
	ProtectContent           bool  `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent messages from forwarding and saving
	ReplyToMessageId         int64 `json:"reply_to_message_id,omitempty"`         // Optional. If the messages are a reply, ID of the original message
	AllowSendingWithoutReply bool  `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
}

// sendMediaGroupParams sendMediaGroupParams contains SendMediaGroup's params
type sendMediaGroupParams struct {
	*SendMediaGroupOptions
	ChatId ChatID       `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Media  []InputMedia `json:"media"`   // A JSON-serialized array describing messages to be sent, must include 2-10 items
}

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
func (b *Bot) SendMediaGroup(chatId ChatID, media []InputMedia, optionalParams *SendMediaGroupOptions) ([]*Message, error) {
	params := &sendMediaGroupParams{}

	params.ChatId = chatId
	params.Media = media
	params.SendMediaGroupOptions = optionalParams

	return doHTTP[[]*Message](b.client, b.url, "sendMediaGroup", params)
}

// SendLocationOptions SendLocationOptions contains SendLocation's optional params
type SendLocationOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	HorizontalAccuracy       float64     `json:"horizontal_accuracy,omitempty"`         // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod               int64       `json:"live_period,omitempty"`                 // Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	Heading                  int64       `json:"heading,omitempty"`                     // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius     int64       `json:"proximity_alert_radius,omitempty"`      // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendLocationParams sendLocationParams contains SendLocation's params
type sendLocationParams struct {
	*SendLocationOptions
	ChatId    ChatID  `json:"chat_id"`   // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude  float64 `json:"latitude"`  // Latitude of the location
	Longitude float64 `json:"longitude"` // Longitude of the location
}

// SendLocation Use this method to send point on the map. On success, the sent Message is returned.
func (b *Bot) SendLocation(chatId ChatID, latitude float64, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	params := &sendLocationParams{}

	params.ChatId = chatId
	params.Latitude = latitude
	params.Longitude = longitude
	params.SendLocationOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendLocation", params)
}

// EditMessageLiveLocationOptions EditMessageLiveLocationOptions contains EditMessageLiveLocation's optional params
type EditMessageLiveLocationOptions struct {
	ChatId               ChatID                `json:"chat_id,omitempty"`                // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int64                 `json:"message_id,omitempty"`             // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id,omitempty"`      // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int64                 `json:"heading,omitempty"`                // Optional. Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional. A JSON-serialized object for a new inline keyboard.
}

// editMessageLiveLocationParams editMessageLiveLocationParams contains EditMessageLiveLocation's params
type editMessageLiveLocationParams struct {
	*EditMessageLiveLocationOptions
	Latitude  float64 `json:"latitude"`  // Latitude of new location
	Longitude float64 `json:"longitude"` // Longitude of new location
}

// EditMessageLiveLocation Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageLiveLocation(latitude float64, longitude float64, optionalParams *EditMessageLiveLocationOptions) (*Message, error) {
	params := &editMessageLiveLocationParams{}

	params.Latitude = latitude
	params.Longitude = longitude
	params.EditMessageLiveLocationOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "editMessageLiveLocation", params)
}

// StopMessageLiveLocationOptions StopMessageLiveLocationOptions contains StopMessageLiveLocation's optional params
type StopMessageLiveLocationOptions struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for a new inline keyboard.
}

// stopMessageLiveLocationParams stopMessageLiveLocationParams contains StopMessageLiveLocation's params
type stopMessageLiveLocationParams struct {
	*StopMessageLiveLocationOptions
}

// StopMessageLiveLocation Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(optionalParams *StopMessageLiveLocationOptions) (*Message, error) {
	params := &stopMessageLiveLocationParams{}

	params.StopMessageLiveLocationOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "stopMessageLiveLocation", params)
}

// SendVenueOptions SendVenueOptions contains SendVenue's optional params
type SendVenueOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FoursquareId             string      `json:"foursquare_id,omitempty"`               // Optional. Foursquare identifier of the venue
	FoursquareType           string      `json:"foursquare_type,omitempty"`             // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId            string      `json:"google_place_id,omitempty"`             // Optional. Google Places identifier of the venue
	GooglePlaceType          string      `json:"google_place_type,omitempty"`           // Optional. Google Places type of the venue. (See supported types.)
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendVenueParams sendVenueParams contains SendVenue's params
type sendVenueParams struct {
	*SendVenueOptions
	ChatId    ChatID  `json:"chat_id"`   // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude  float64 `json:"latitude"`  // Latitude of the venue
	Longitude float64 `json:"longitude"` // Longitude of the venue
	Title     string  `json:"title"`     // Name of the venue
	Address   string  `json:"address"`   // Address of the venue
}

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
func (b *Bot) SendVenue(chatId ChatID, latitude float64, longitude float64, title string, address string, optionalParams *SendVenueOptions) (*Message, error) {
	params := &sendVenueParams{}

	params.ChatId = chatId
	params.Latitude = latitude
	params.Longitude = longitude
	params.Title = title
	params.Address = address
	params.SendVenueOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendVenue", params)
}

// SendContactOptions SendContactOptions contains SendContact's optional params
type SendContactOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	LastName                 string      `json:"last_name,omitempty"`                   // Optional. Contact's last name
	Vcard                    string      `json:"vcard,omitempty"`                       // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendContactParams sendContactParams contains SendContact's params
type sendContactParams struct {
	*SendContactOptions
	ChatId      ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
}

// SendContact Use this method to send phone contacts. On success, the sent Message is returned.
func (b *Bot) SendContact(chatId ChatID, phoneNumber string, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	params := &sendContactParams{}

	params.ChatId = chatId
	params.PhoneNumber = phoneNumber
	params.FirstName = firstName
	params.SendContactOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendContact", params)
}

// SendPollOptions SendPollOptions contains SendPoll's optional params
type SendPollOptions struct {
	MessageThreadId          int64            `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
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
	ReplyMarkup              ReplyMarkup      `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendPollParams sendPollParams contains SendPoll's params
type sendPollParams struct {
	*SendPollOptions
	ChatId   ChatID   `json:"chat_id"`  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question string   `json:"question"` // Poll question, 1-300 characters
	Options  []string `json:"options"`  // A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
}

// SendPoll Use this method to send a native poll. On success, the sent Message is returned.
func (b *Bot) SendPoll(chatId ChatID, question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	params := &sendPollParams{}

	params.ChatId = chatId
	params.Question = question
	params.Options = options
	params.SendPollOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendPoll", params)
}

// SendDiceOptions SendDiceOptions contains SendDice's optional params
type SendDiceOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                    string      `json:"emoji,omitempty"`                       // Optional. Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64 for “”. Defaults to “”
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendDiceParams sendDiceParams contains SendDice's params
type sendDiceParams struct {
	*SendDiceOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// SendDice Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (b *Bot) SendDice(chatId ChatID, optionalParams *SendDiceOptions) (*Message, error) {
	params := &sendDiceParams{}

	params.ChatId = chatId
	params.SendDiceOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendDice", params)
}

// SendChatActionOptions SendChatActionOptions contains SendChatAction's optional params
type SendChatActionOptions struct {
	MessageThreadId int64 `json:"message_thread_id,omitempty"` // Optional. Unique identifier for the target message thread; supergroups only
}

// sendChatActionParams sendChatActionParams contains SendChatAction's params
type sendChatActionParams struct {
	*SendChatActionOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

// SendChatAction Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
func (b *Bot) SendChatAction(chatId ChatID, action string, optionalParams *SendChatActionOptions) (bool, error) {
	params := &sendChatActionParams{}

	params.ChatId = chatId
	params.Action = action
	params.SendChatActionOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "sendChatAction", params)
}

// GetUserProfilePhotosOptions GetUserProfilePhotosOptions contains GetUserProfilePhotos's optional params
type GetUserProfilePhotosOptions struct {
	Offset int64 `json:"offset,omitempty"` // Optional. Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int64 `json:"limit,omitempty"`  // Optional. Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

// getUserProfilePhotosParams getUserProfilePhotosParams contains GetUserProfilePhotos's params
type getUserProfilePhotosParams struct {
	*GetUserProfilePhotosOptions
	UserId int64 `json:"user_id"` // Unique identifier of the target user
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (b *Bot) GetUserProfilePhotos(userId int64, optionalParams *GetUserProfilePhotosOptions) (*UserProfilePhotos, error) {
	params := &getUserProfilePhotosParams{}

	params.UserId = userId
	params.GetUserProfilePhotosOptions = optionalParams

	return doHTTP[*UserProfilePhotos](b.client, b.url, "getUserProfilePhotos", params)
}

// getFileParams getFileParams contains GetFile's params
type getFileParams struct {
	FileId string `json:"file_id"` // File identifier to get information about
}

// GetFile Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
func (b *Bot) GetFile(fileId string) (*File, error) {
	params := &getFileParams{}

	params.FileId = fileId

	return doHTTP[*File](b.client, b.url, "getFile", params)
}

// BanChatMemberOptions BanChatMemberOptions contains BanChatMember's optional params
type BanChatMemberOptions struct {
	UntilDate      int64 `json:"until_date,omitempty"`      // Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool  `json:"revoke_messages,omitempty"` // Optional. Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

// banChatMemberParams banChatMemberParams contains BanChatMember's params
type banChatMemberParams struct {
	*BanChatMemberOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// BanChatMember Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatMember(chatId ChatID, userId int64, optionalParams *BanChatMemberOptions) (bool, error) {
	params := &banChatMemberParams{}

	params.ChatId = chatId
	params.UserId = userId
	params.BanChatMemberOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "banChatMember", params)
}

// UnbanChatMemberOptions UnbanChatMemberOptions contains UnbanChatMember's optional params
type UnbanChatMemberOptions struct {
	OnlyIfBanned bool `json:"only_if_banned,omitempty"` // Optional. Do nothing if the user is not banned
}

// unbanChatMemberParams unbanChatMemberParams contains UnbanChatMember's params
type unbanChatMemberParams struct {
	*UnbanChatMemberOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// UnbanChatMember Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (b *Bot) UnbanChatMember(chatId ChatID, userId int64, optionalParams *UnbanChatMemberOptions) (bool, error) {
	params := &unbanChatMemberParams{}

	params.ChatId = chatId
	params.UserId = userId
	params.UnbanChatMemberOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "unbanChatMember", params)
}

// RestrictChatMemberOptions RestrictChatMemberOptions contains RestrictChatMember's optional params
type RestrictChatMemberOptions struct {
	UseIndependentChatPermissions bool  `json:"use_independent_chat_permissions,omitempty"` // Optional. Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UntilDate                     int64 `json:"until_date,omitempty"`                       // Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

// restrictChatMemberParams restrictChatMemberParams contains RestrictChatMember's params
type restrictChatMemberParams struct {
	*RestrictChatMemberOptions
	ChatId      ChatID           `json:"chat_id"`     // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int64            `json:"user_id"`     // Unique identifier of the target user
	Permissions *ChatPermissions `json:"permissions"` // A JSON-serialized object for new user permissions
}

// RestrictChatMember Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(chatId ChatID, userId int64, permissions *ChatPermissions, optionalParams *RestrictChatMemberOptions) (bool, error) {
	params := &restrictChatMemberParams{}

	params.ChatId = chatId
	params.UserId = userId
	params.Permissions = permissions
	params.RestrictChatMemberOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "restrictChatMember", params)
}

// PromoteChatMemberOptions PromoteChatMemberOptions contains PromoteChatMember's optional params
type PromoteChatMemberOptions struct {
	IsAnonymous         bool `json:"is_anonymous,omitempty"`           // Optional. Pass True if the administrator's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat,omitempty"`        // Optional. Pass True if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanPostMessages     bool `json:"can_post_messages,omitempty"`      // Optional. Pass True if the administrator can create channel posts, channels only
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`      // Optional. Pass True if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages   bool `json:"can_delete_messages,omitempty"`    // Optional. Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"` // Optional. Pass True if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members,omitempty"`   // Optional. Pass True if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool `json:"can_promote_members,omitempty"`    // Optional. Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool `json:"can_change_info,omitempty"`        // Optional. Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users,omitempty"`       // Optional. Pass True if the administrator can invite new users to the chat
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`       // Optional. Pass True if the administrator can pin messages, supergroups only
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`      // Optional. Pass True if the user is allowed to create, rename, close, and reopen forum topics, supergroups only
}

// promoteChatMemberParams promoteChatMemberParams contains PromoteChatMember's params
type promoteChatMemberParams struct {
	*PromoteChatMemberOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(chatId ChatID, userId int64, optionalParams *PromoteChatMemberOptions) (bool, error) {
	params := &promoteChatMemberParams{}

	params.ChatId = chatId
	params.UserId = userId
	params.PromoteChatMemberOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "promoteChatMember", params)
}

// setChatAdministratorCustomTitleParams setChatAdministratorCustomTitleParams contains SetChatAdministratorCustomTitle's params
type setChatAdministratorCustomTitleParams struct {
	ChatId      ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int64  `json:"user_id"`      // Unique identifier of the target user
	CustomTitle string `json:"custom_title"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

// SetChatAdministratorCustomTitle Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (b *Bot) SetChatAdministratorCustomTitle(chatId ChatID, userId int64, customTitle string) (bool, error) {
	params := &setChatAdministratorCustomTitleParams{}

	params.ChatId = chatId
	params.UserId = userId
	params.CustomTitle = customTitle

	return doHTTP[bool](b.client, b.url, "setChatAdministratorCustomTitle", params)
}

// banChatSenderChatParams banChatSenderChatParams contains BanChatSenderChat's params
type banChatSenderChatParams struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatSenderChat(chatId ChatID, senderChatId int64) (bool, error) {
	params := &banChatSenderChatParams{}

	params.ChatId = chatId
	params.SenderChatId = senderChatId

	return doHTTP[bool](b.client, b.url, "banChatSenderChat", params)
}

// unbanChatSenderChatParams unbanChatSenderChatParams contains UnbanChatSenderChat's params
type unbanChatSenderChatParams struct {
	ChatId       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatId int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

// UnbanChatSenderChat Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) UnbanChatSenderChat(chatId ChatID, senderChatId int64) (bool, error) {
	params := &unbanChatSenderChatParams{}

	params.ChatId = chatId
	params.SenderChatId = senderChatId

	return doHTTP[bool](b.client, b.url, "unbanChatSenderChat", params)
}

// SetChatPermissionsOptions SetChatPermissionsOptions contains SetChatPermissions's optional params
type SetChatPermissionsOptions struct {
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"` // Optional. Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
}

// setChatPermissionsParams setChatPermissionsParams contains SetChatPermissions's params
type setChatPermissionsParams struct {
	*SetChatPermissionsOptions
	ChatId      ChatID           `json:"chat_id"`     // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions *ChatPermissions `json:"permissions"` // A JSON-serialized object for new default chat permissions
}

// SetChatPermissions Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (b *Bot) SetChatPermissions(chatId ChatID, permissions *ChatPermissions, optionalParams *SetChatPermissionsOptions) (bool, error) {
	params := &setChatPermissionsParams{}

	params.ChatId = chatId
	params.Permissions = permissions
	params.SetChatPermissionsOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setChatPermissions", params)
}

// exportChatInviteLinkParams exportChatInviteLinkParams contains ExportChatInviteLink's params
type exportChatInviteLinkParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
func (b *Bot) ExportChatInviteLink(chatId ChatID) (string, error) {
	params := &exportChatInviteLinkParams{}

	params.ChatId = chatId

	return doHTTP[string](b.client, b.url, "exportChatInviteLink", params)
}

// CreateChatInviteLinkOptions CreateChatInviteLinkOptions contains CreateChatInviteLink's optional params
type CreateChatInviteLinkOptions struct {
	Name               string `json:"name,omitempty"`                 // Optional. Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Optional. Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // Optional. True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// createChatInviteLinkParams createChatInviteLinkParams contains CreateChatInviteLink's params
type createChatInviteLinkParams struct {
	*CreateChatInviteLinkOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// CreateChatInviteLink Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (b *Bot) CreateChatInviteLink(chatId ChatID, optionalParams *CreateChatInviteLinkOptions) (*ChatInviteLink, error) {
	params := &createChatInviteLinkParams{}

	params.ChatId = chatId
	params.CreateChatInviteLinkOptions = optionalParams

	return doHTTP[*ChatInviteLink](b.client, b.url, "createChatInviteLink", params)
}

// EditChatInviteLinkOptions EditChatInviteLinkOptions contains EditChatInviteLink's optional params
type EditChatInviteLinkOptions struct {
	Name               string `json:"name,omitempty"`                 // Optional. Invite link name; 0-32 characters
	ExpireDate         int64  `json:"expire_date,omitempty"`          // Optional. Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  `json:"member_limit,omitempty"`         // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // Optional. True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// editChatInviteLinkParams editChatInviteLinkParams contains EditChatInviteLink's params
type editChatInviteLinkParams struct {
	*EditChatInviteLinkOptions
	ChatId     ChatID `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"` // The invite link to edit
}

// EditChatInviteLink Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (b *Bot) EditChatInviteLink(chatId ChatID, inviteLink string, optionalParams *EditChatInviteLinkOptions) (*ChatInviteLink, error) {
	params := &editChatInviteLinkParams{}

	params.ChatId = chatId
	params.InviteLink = inviteLink
	params.EditChatInviteLinkOptions = optionalParams

	return doHTTP[*ChatInviteLink](b.client, b.url, "editChatInviteLink", params)
}

// revokeChatInviteLinkParams revokeChatInviteLinkParams contains RevokeChatInviteLink's params
type revokeChatInviteLinkParams struct {
	ChatId     ChatID `json:"chat_id"`     // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"` // The invite link to revoke
}

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
func (b *Bot) RevokeChatInviteLink(chatId ChatID, inviteLink string) (*ChatInviteLink, error) {
	params := &revokeChatInviteLinkParams{}

	params.ChatId = chatId
	params.InviteLink = inviteLink

	return doHTTP[*ChatInviteLink](b.client, b.url, "revokeChatInviteLink", params)
}

// approveChatJoinRequestParams approveChatJoinRequestParams contains ApproveChatJoinRequest's params
type approveChatJoinRequestParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// ApproveChatJoinRequest Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) ApproveChatJoinRequest(chatId ChatID, userId int64) (bool, error) {
	params := &approveChatJoinRequestParams{}

	params.ChatId = chatId
	params.UserId = userId

	return doHTTP[bool](b.client, b.url, "approveChatJoinRequest", params)
}

// declineChatJoinRequestParams declineChatJoinRequestParams contains DeclineChatJoinRequest's params
type declineChatJoinRequestParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// DeclineChatJoinRequest Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) DeclineChatJoinRequest(chatId ChatID, userId int64) (bool, error) {
	params := &declineChatJoinRequestParams{}

	params.ChatId = chatId
	params.UserId = userId

	return doHTTP[bool](b.client, b.url, "declineChatJoinRequest", params)
}

// setChatPhotoParams setChatPhotoParams contains SetChatPhoto's params
type setChatPhotoParams struct {
	ChatId ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

// SetChatPhoto Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatPhoto(chatId ChatID, photo InputFile) (bool, error) {
	params := &setChatPhotoParams{}

	params.ChatId = chatId
	params.Photo = photo

	return doHTTP[bool](b.client, b.url, "setChatPhoto", params)
}

func (params *setChatPhotoParams) HasUploadable() bool {
	return params.Photo.NeedsUpload()
}

// deleteChatPhotoParams deleteChatPhotoParams contains DeleteChatPhoto's params
type deleteChatPhotoParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// DeleteChatPhoto Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) DeleteChatPhoto(chatId ChatID) (bool, error) {
	params := &deleteChatPhotoParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "deleteChatPhoto", params)
}

// setChatTitleParams setChatTitleParams contains SetChatTitle's params
type setChatTitleParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-128 characters
}

// SetChatTitle Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatTitle(chatId ChatID, title string) (bool, error) {
	params := &setChatTitleParams{}

	params.ChatId = chatId
	params.Title = title

	return doHTTP[bool](b.client, b.url, "setChatTitle", params)
}

// SetChatDescriptionOptions SetChatDescriptionOptions contains SetChatDescription's optional params
type SetChatDescriptionOptions struct {
	Description string `json:"description,omitempty"` // Optional. New chat description, 0-255 characters
}

// setChatDescriptionParams setChatDescriptionParams contains SetChatDescription's params
type setChatDescriptionParams struct {
	*SetChatDescriptionOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// SetChatDescription Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatDescription(chatId ChatID, optionalParams *SetChatDescriptionOptions) (bool, error) {
	params := &setChatDescriptionParams{}

	params.ChatId = chatId
	params.SetChatDescriptionOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setChatDescription", params)
}

// PinChatMessageOptions PinChatMessageOptions contains PinChatMessage's optional params
type PinChatMessageOptions struct {
	DisableNotification bool `json:"disable_notification,omitempty"` // Optional. Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

// pinChatMessageParams pinChatMessageParams contains PinChatMessage's params
type pinChatMessageParams struct {
	*PinChatMessageOptions
	ChatId    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id"` // Identifier of a message to pin
}

// PinChatMessage Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) PinChatMessage(chatId ChatID, messageId int64, optionalParams *PinChatMessageOptions) (bool, error) {
	params := &pinChatMessageParams{}

	params.ChatId = chatId
	params.MessageId = messageId
	params.PinChatMessageOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "pinChatMessage", params)
}

// UnpinChatMessageOptions UnpinChatMessageOptions contains UnpinChatMessage's optional params
type UnpinChatMessageOptions struct {
	MessageId int64 `json:"message_id,omitempty"` // Optional. Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

// unpinChatMessageParams unpinChatMessageParams contains UnpinChatMessage's params
type unpinChatMessageParams struct {
	*UnpinChatMessageOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// UnpinChatMessage Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) UnpinChatMessage(chatId ChatID, optionalParams *UnpinChatMessageOptions) (bool, error) {
	params := &unpinChatMessageParams{}

	params.ChatId = chatId
	params.UnpinChatMessageOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "unpinChatMessage", params)
}

// unpinAllChatMessagesParams unpinAllChatMessagesParams contains UnpinAllChatMessages's params
type unpinAllChatMessagesParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// UnpinAllChatMessages Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) UnpinAllChatMessages(chatId ChatID) (bool, error) {
	params := &unpinAllChatMessagesParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "unpinAllChatMessages", params)
}

// leaveChatParams leaveChatParams contains LeaveChat's params
type leaveChatParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(chatId ChatID) (bool, error) {
	params := &leaveChatParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "leaveChat", params)
}

// getChatParams getChatParams contains GetChat's params
type getChatParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (b *Bot) GetChat(chatId ChatID) (*Chat, error) {
	params := &getChatParams{}

	params.ChatId = chatId

	return doHTTP[*Chat](b.client, b.url, "getChat", params)
}

// getChatAdministratorsParams getChatAdministratorsParams contains GetChatAdministrators's params
type getChatAdministratorsParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatAdministrators Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
func (b *Bot) GetChatAdministrators(chatId ChatID) ([]*ChatMember, error) {
	params := &getChatAdministratorsParams{}

	params.ChatId = chatId

	return doHTTP[[]*ChatMember](b.client, b.url, "getChatAdministrators", params)
}

// getChatMemberCountParams getChatMemberCountParams contains GetChatMemberCount's params
type getChatMemberCountParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMemberCount Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMemberCount(chatId ChatID) (int64, error) {
	params := &getChatMemberCountParams{}

	params.ChatId = chatId

	return doHTTP[int64](b.client, b.url, "getChatMemberCount", params)
}

// getChatMemberParams getChatMemberParams contains GetChatMember's params
type getChatMemberParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserId int64  `json:"user_id"` // Unique identifier of the target user
}

// GetChatMember Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
func (b *Bot) GetChatMember(chatId ChatID, userId int64) (*ChatMember, error) {
	params := &getChatMemberParams{}

	params.ChatId = chatId
	params.UserId = userId

	return doHTTP[*ChatMember](b.client, b.url, "getChatMember", params)
}

// setChatStickerSetParams setChatStickerSetParams contains SetChatStickerSet's params
type setChatStickerSetParams struct {
	ChatId         ChatID `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

// SetChatStickerSet Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) SetChatStickerSet(chatId ChatID, stickerSetName string) (bool, error) {
	params := &setChatStickerSetParams{}

	params.ChatId = chatId
	params.StickerSetName = stickerSetName

	return doHTTP[bool](b.client, b.url, "setChatStickerSet", params)
}

// deleteChatStickerSetParams deleteChatStickerSetParams contains DeleteChatStickerSet's params
type deleteChatStickerSetParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) DeleteChatStickerSet(chatId ChatID) (bool, error) {
	params := &deleteChatStickerSetParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "deleteChatStickerSet", params)
}

// getForumTopicIconStickersParams getForumTopicIconStickersParams contains GetForumTopicIconStickers's params
type getForumTopicIconStickersParams struct {
}

// GetForumTopicIconStickers Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
func (b *Bot) GetForumTopicIconStickers() ([]*Sticker, error) {
	params := &getForumTopicIconStickersParams{}

	return doHTTP[[]*Sticker](b.client, b.url, "getForumTopicIconStickers", params)
}

// CreateForumTopicOptions CreateForumTopicOptions contains CreateForumTopic's optional params
type CreateForumTopicOptions struct {
	IconColor         int64  `json:"icon_color,omitempty"`           // Optional. Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
}

// createForumTopicParams createForumTopicParams contains CreateForumTopic's params
type createForumTopicParams struct {
	*CreateForumTopicOptions
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // Topic name, 1-128 characters
}

// CreateForumTopic Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
func (b *Bot) CreateForumTopic(chatId ChatID, name string, optionalParams *CreateForumTopicOptions) (*ForumTopic, error) {
	params := &createForumTopicParams{}

	params.ChatId = chatId
	params.Name = name
	params.CreateForumTopicOptions = optionalParams

	return doHTTP[*ForumTopic](b.client, b.url, "createForumTopic", params)
}

// EditForumTopicOptions EditForumTopicOptions contains EditForumTopic's optional params
type EditForumTopicOptions struct {
	Name              string `json:"name,omitempty"`                 // Optional. New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"` // Optional. New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

// editForumTopicParams editForumTopicParams contains EditForumTopic's params
type editForumTopicParams struct {
	*EditForumTopicOptions
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// EditForumTopic Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) EditForumTopic(chatId ChatID, messageThreadId int64, optionalParams *EditForumTopicOptions) (bool, error) {
	params := &editForumTopicParams{}

	params.ChatId = chatId
	params.MessageThreadId = messageThreadId
	params.EditForumTopicOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "editForumTopic", params)
}

// closeForumTopicParams closeForumTopicParams contains CloseForumTopic's params
type closeForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// CloseForumTopic Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) CloseForumTopic(chatId ChatID, messageThreadId int64) (bool, error) {
	params := &closeForumTopicParams{}

	params.ChatId = chatId
	params.MessageThreadId = messageThreadId

	return doHTTP[bool](b.client, b.url, "closeForumTopic", params)
}

// reopenForumTopicParams reopenForumTopicParams contains ReopenForumTopic's params
type reopenForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// ReopenForumTopic Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) ReopenForumTopic(chatId ChatID, messageThreadId int64) (bool, error) {
	params := &reopenForumTopicParams{}

	params.ChatId = chatId
	params.MessageThreadId = messageThreadId

	return doHTTP[bool](b.client, b.url, "reopenForumTopic", params)
}

// deleteForumTopicParams deleteForumTopicParams contains DeleteForumTopic's params
type deleteForumTopicParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// DeleteForumTopic Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (b *Bot) DeleteForumTopic(chatId ChatID, messageThreadId int64) (bool, error) {
	params := &deleteForumTopicParams{}

	params.ChatId = chatId
	params.MessageThreadId = messageThreadId

	return doHTTP[bool](b.client, b.url, "deleteForumTopic", params)
}

// unpinAllForumTopicMessagesParams unpinAllForumTopicMessagesParams contains UnpinAllForumTopicMessages's params
type unpinAllForumTopicMessagesParams struct {
	ChatId          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadId int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

// UnpinAllForumTopicMessages Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllForumTopicMessages(chatId ChatID, messageThreadId int64) (bool, error) {
	params := &unpinAllForumTopicMessagesParams{}

	params.ChatId = chatId
	params.MessageThreadId = messageThreadId

	return doHTTP[bool](b.client, b.url, "unpinAllForumTopicMessages", params)
}

// editGeneralForumTopicParams editGeneralForumTopicParams contains EditGeneralForumTopic's params
type editGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // New topic name, 1-128 characters
}

// EditGeneralForumTopic Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
func (b *Bot) EditGeneralForumTopic(chatId ChatID, name string) (bool, error) {
	params := &editGeneralForumTopicParams{}

	params.ChatId = chatId
	params.Name = name

	return doHTTP[bool](b.client, b.url, "editGeneralForumTopic", params)
}

// closeGeneralForumTopicParams closeGeneralForumTopicParams contains CloseGeneralForumTopic's params
type closeGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// CloseGeneralForumTopic Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) CloseGeneralForumTopic(chatId ChatID) (bool, error) {
	params := &closeGeneralForumTopicParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "closeGeneralForumTopic", params)
}

// reopenGeneralForumTopicParams reopenGeneralForumTopicParams contains ReopenGeneralForumTopic's params
type reopenGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// ReopenGeneralForumTopic Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (b *Bot) ReopenGeneralForumTopic(chatId ChatID) (bool, error) {
	params := &reopenGeneralForumTopicParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "reopenGeneralForumTopic", params)
}

// hideGeneralForumTopicParams hideGeneralForumTopicParams contains HideGeneralForumTopic's params
type hideGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// HideGeneralForumTopic Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (b *Bot) HideGeneralForumTopic(chatId ChatID) (bool, error) {
	params := &hideGeneralForumTopicParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "hideGeneralForumTopic", params)
}

// unhideGeneralForumTopicParams unhideGeneralForumTopicParams contains UnhideGeneralForumTopic's params
type unhideGeneralForumTopicParams struct {
	ChatId ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// UnhideGeneralForumTopic Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) UnhideGeneralForumTopic(chatId ChatID) (bool, error) {
	params := &unhideGeneralForumTopicParams{}

	params.ChatId = chatId

	return doHTTP[bool](b.client, b.url, "unhideGeneralForumTopic", params)
}

// AnswerCallbackQueryOptions AnswerCallbackQueryOptions contains AnswerCallbackQuery's optional params
type AnswerCallbackQueryOptions struct {
	Text      string `json:"text,omitempty"`       // Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert bool   `json:"show_alert,omitempty"` // Optional. If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url       string `json:"url,omitempty"`        // Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime int64  `json:"cache_time,omitempty"` // Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

// answerCallbackQueryParams answerCallbackQueryParams contains AnswerCallbackQuery's params
type answerCallbackQueryParams struct {
	*AnswerCallbackQueryOptions
	CallbackQueryId string `json:"callback_query_id"` // Unique identifier for the query to be answered
}

// AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (b *Bot) AnswerCallbackQuery(callbackQueryId string, optionalParams *AnswerCallbackQueryOptions) (bool, error) {
	params := &answerCallbackQueryParams{}

	params.CallbackQueryId = callbackQueryId
	params.AnswerCallbackQueryOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "answerCallbackQuery", params)
}

// SetMyCommandsOptions SetMyCommandsOptions contains SetMyCommands's optional params
type SetMyCommandsOptions struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// setMyCommandsParams setMyCommandsParams contains SetMyCommands's params
type setMyCommandsParams struct {
	*SetMyCommandsOptions
	Commands []*BotCommand `json:"commands"` // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
}

// SetMyCommands Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (b *Bot) SetMyCommands(commands []*BotCommand, optionalParams *SetMyCommandsOptions) (bool, error) {
	params := &setMyCommandsParams{}

	params.Commands = commands
	params.SetMyCommandsOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setMyCommands", params)
}

// DeleteMyCommandsOptions DeleteMyCommandsOptions contains DeleteMyCommands's optional params
type DeleteMyCommandsOptions struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

// deleteMyCommandsParams deleteMyCommandsParams contains DeleteMyCommands's params
type deleteMyCommandsParams struct {
	*DeleteMyCommandsOptions
}

// DeleteMyCommands Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (b *Bot) DeleteMyCommands(optionalParams *DeleteMyCommandsOptions) (bool, error) {
	params := &deleteMyCommandsParams{}

	params.DeleteMyCommandsOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "deleteMyCommands", params)
}

// GetMyCommandsOptions GetMyCommandsOptions contains GetMyCommands's optional params
type GetMyCommandsOptions struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // Optional. A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // Optional. A two-letter ISO 639-1 language code or an empty string
}

// getMyCommandsParams getMyCommandsParams contains GetMyCommands's params
type getMyCommandsParams struct {
	*GetMyCommandsOptions
}

// GetMyCommands Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
func (b *Bot) GetMyCommands(optionalParams *GetMyCommandsOptions) ([]*BotCommand, error) {
	params := &getMyCommandsParams{}

	params.GetMyCommandsOptions = optionalParams

	return doHTTP[[]*BotCommand](b.client, b.url, "getMyCommands", params)
}

// SetChatMenuButtonOptions SetChatMenuButtonOptions contains SetChatMenuButton's optional params
type SetChatMenuButtonOptions struct {
	ChatId     int64       `json:"chat_id,omitempty"`     // Optional. Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton *MenuButton `json:"menu_button,omitempty"` // Optional. A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}

// setChatMenuButtonParams setChatMenuButtonParams contains SetChatMenuButton's params
type setChatMenuButtonParams struct {
	*SetChatMenuButtonOptions
}

// SetChatMenuButton Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (b *Bot) SetChatMenuButton(optionalParams *SetChatMenuButtonOptions) (bool, error) {
	params := &setChatMenuButtonParams{}

	params.SetChatMenuButtonOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setChatMenuButton", params)
}

// GetChatMenuButtonOptions GetChatMenuButtonOptions contains GetChatMenuButton's optional params
type GetChatMenuButtonOptions struct {
	ChatId int64 `json:"chat_id,omitempty"` // Optional. Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
}

// getChatMenuButtonParams getChatMenuButtonParams contains GetChatMenuButton's params
type getChatMenuButtonParams struct {
	*GetChatMenuButtonOptions
}

// GetChatMenuButton Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
func (b *Bot) GetChatMenuButton(optionalParams *GetChatMenuButtonOptions) (*MenuButton, error) {
	params := &getChatMenuButtonParams{}

	params.GetChatMenuButtonOptions = optionalParams

	return doHTTP[*MenuButton](b.client, b.url, "getChatMenuButton", params)
}

// SetMyDefaultAdministratorRightsOptions SetMyDefaultAdministratorRightsOptions contains SetMyDefaultAdministratorRights's optional params
type SetMyDefaultAdministratorRightsOptions struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`       // Optional. A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	ForChannels bool                     `json:"for_channels,omitempty"` // Optional. Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
}

// setMyDefaultAdministratorRightsParams setMyDefaultAdministratorRightsParams contains SetMyDefaultAdministratorRights's params
type setMyDefaultAdministratorRightsParams struct {
	*SetMyDefaultAdministratorRightsOptions
}

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
func (b *Bot) SetMyDefaultAdministratorRights(optionalParams *SetMyDefaultAdministratorRightsOptions) (bool, error) {
	params := &setMyDefaultAdministratorRightsParams{}

	params.SetMyDefaultAdministratorRightsOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setMyDefaultAdministratorRights", params)
}

// GetMyDefaultAdministratorRightsOptions GetMyDefaultAdministratorRightsOptions contains GetMyDefaultAdministratorRights's optional params
type GetMyDefaultAdministratorRightsOptions struct {
	ForChannels bool `json:"for_channels,omitempty"` // Optional. Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
}

// getMyDefaultAdministratorRightsParams getMyDefaultAdministratorRightsParams contains GetMyDefaultAdministratorRights's params
type getMyDefaultAdministratorRightsParams struct {
	*GetMyDefaultAdministratorRightsOptions
}

// GetMyDefaultAdministratorRights Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
func (b *Bot) GetMyDefaultAdministratorRights(optionalParams *GetMyDefaultAdministratorRightsOptions) (*ChatAdministratorRights, error) {
	params := &getMyDefaultAdministratorRightsParams{}

	params.GetMyDefaultAdministratorRightsOptions = optionalParams

	return doHTTP[*ChatAdministratorRights](b.client, b.url, "getMyDefaultAdministratorRights", params)
}

// EditMessageTextOptions EditMessageTextOptions contains EditMessageText's optional params
type EditMessageTextOptions struct {
	ChatId                ChatID                `json:"chat_id,omitempty"`                  // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int64                 `json:"message_id,omitempty"`               // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string                `json:"inline_message_id,omitempty"`        // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity      `json:"entities,omitempty"`                 // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. A JSON-serialized object for an inline keyboard.
}

// editMessageTextParams editMessageTextParams contains EditMessageText's params
type editMessageTextParams struct {
	*EditMessageTextOptions
	Text string `json:"text"` // New text of the message, 1-4096 characters after entities parsing
}

// EditMessageText Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageText(text string, optionalParams *EditMessageTextOptions) (*Message, error) {
	params := &editMessageTextParams{}

	params.Text = text
	params.EditMessageTextOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "editMessageText", params)
}

// EditMessageCaptionOptions EditMessageCaptionOptions contains EditMessageCaption's optional params
type EditMessageCaptionOptions struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string                `json:"caption,omitempty"`           // Optional. New caption of the message, 0-1024 characters after entities parsing
	ParseMode       ParseMode             `json:"parse_mode,omitempty"`        // Optional. Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities []*MessageEntity      `json:"caption_entities,omitempty"`  // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for an inline keyboard.
}

// editMessageCaptionParams editMessageCaptionParams contains EditMessageCaption's params
type editMessageCaptionParams struct {
	*EditMessageCaptionOptions
}

// EditMessageCaption Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageCaption(optionalParams *EditMessageCaptionOptions) (*Message, error) {
	params := &editMessageCaptionParams{}

	params.EditMessageCaptionOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "editMessageCaption", params)
}

// EditMessageMediaOptions EditMessageMediaOptions contains EditMessageMedia's optional params
type EditMessageMediaOptions struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for a new inline keyboard.
}

// editMessageMediaParams editMessageMediaParams contains EditMessageMedia's params
type editMessageMediaParams struct {
	*EditMessageMediaOptions
	Media *InputMedia `json:"media"` // A JSON-serialized object for a new media content of the message
}

// EditMessageMedia Use this method to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageMedia(media *InputMedia, optionalParams *EditMessageMediaOptions) (*Message, error) {
	params := &editMessageMediaParams{}

	params.Media = media
	params.EditMessageMediaOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "editMessageMedia", params)
}

// EditMessageReplyMarkupOptions EditMessageReplyMarkupOptions contains EditMessageReplyMarkup's optional params
type EditMessageReplyMarkupOptions struct {
	ChatId          ChatID                `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int64                 `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional. A JSON-serialized object for an inline keyboard.
}

// editMessageReplyMarkupParams editMessageReplyMarkupParams contains EditMessageReplyMarkup's params
type editMessageReplyMarkupParams struct {
	*EditMessageReplyMarkupOptions
}

// EditMessageReplyMarkup Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageReplyMarkup(optionalParams *EditMessageReplyMarkupOptions) (*Message, error) {
	params := &editMessageReplyMarkupParams{}

	params.EditMessageReplyMarkupOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "editMessageReplyMarkup", params)
}

// StopPollOptions StopPollOptions contains StopPoll's optional params
type StopPollOptions struct {
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. A JSON-serialized object for a new message inline keyboard.
}

// stopPollParams stopPollParams contains StopPoll's params
type stopPollParams struct {
	*StopPollOptions
	ChatId    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id"` // Identifier of the original message with the poll
}

// StopPoll Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
func (b *Bot) StopPoll(chatId ChatID, messageId int64, optionalParams *StopPollOptions) (*Poll, error) {
	params := &stopPollParams{}

	params.ChatId = chatId
	params.MessageId = messageId
	params.StopPollOptions = optionalParams

	return doHTTP[*Poll](b.client, b.url, "stopPoll", params)
}

// deleteMessageParams deleteMessageParams contains DeleteMessage's params
type deleteMessageParams struct {
	ChatId    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int64  `json:"message_id"` // Identifier of the message to delete
}

// DeleteMessage Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can't be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
func (b *Bot) DeleteMessage(chatId ChatID, messageId int64) (bool, error) {
	params := &deleteMessageParams{}

	params.ChatId = chatId
	params.MessageId = messageId

	return doHTTP[bool](b.client, b.url, "deleteMessage", params)
}

// SendStickerOptions SendStickerOptions contains SendSticker's optional params
type SendStickerOptions struct {
	MessageThreadId          int64       `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	DisableNotification      bool        `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool        `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64       `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// sendStickerParams sendStickerParams contains SendSticker's params
type sendStickerParams struct {
	*SendStickerOptions
	ChatId  ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker InputFile `json:"sticker"` // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
}

// SendSticker Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
func (b *Bot) SendSticker(chatId ChatID, sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	params := &sendStickerParams{}

	params.ChatId = chatId
	params.Sticker = sticker
	params.SendStickerOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendSticker", params)
}

func (params *sendStickerParams) HasUploadable() bool {
	return params.Sticker.NeedsUpload()
}

// getStickerSetParams getStickerSetParams contains GetStickerSet's params
type getStickerSetParams struct {
	Name string `json:"name"` // Name of the sticker set
}

// GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
func (b *Bot) GetStickerSet(name string) (*StickerSet, error) {
	params := &getStickerSetParams{}

	params.Name = name

	return doHTTP[*StickerSet](b.client, b.url, "getStickerSet", params)
}

// getCustomEmojiStickersParams getCustomEmojiStickersParams contains GetCustomEmojiStickers's params
type getCustomEmojiStickersParams struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"` // List of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

// GetCustomEmojiStickers Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (b *Bot) GetCustomEmojiStickers(customEmojiIds []string) ([]*Sticker, error) {
	params := &getCustomEmojiStickersParams{}

	params.CustomEmojiIds = customEmojiIds

	return doHTTP[[]*Sticker](b.client, b.url, "getCustomEmojiStickers", params)
}

// uploadStickerFileParams uploadStickerFileParams contains UploadStickerFile's params
type uploadStickerFileParams struct {
	UserId     int64     `json:"user_id"`     // User identifier of sticker file owner
	PngSticker InputFile `json:"png_sticker"` // PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More information on Sending Files »
}

// UploadStickerFile Use this method to upload a .PNG file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
func (b *Bot) UploadStickerFile(userId int64, pngSticker InputFile) (*File, error) {
	params := &uploadStickerFileParams{}

	params.UserId = userId
	params.PngSticker = pngSticker

	return doHTTP[*File](b.client, b.url, "uploadStickerFile", params)
}

func (params *uploadStickerFileParams) HasUploadable() bool {
	return params.PngSticker.NeedsUpload()
}

// CreateNewStickerSetOptions CreateNewStickerSetOptions contains CreateNewStickerSet's optional params
type CreateNewStickerSetOptions struct {
	PngSticker   InputFile     `json:"png_sticker,omitempty"`   // Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	TgsSticker   InputFile     `json:"tgs_sticker,omitempty"`   // Optional. TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#animated-sticker-requirements for technical requirements
	WebmSticker  InputFile     `json:"webm_sticker,omitempty"`  // Optional. WEBM video with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#video-sticker-requirements for technical requirements
	StickerType  string        `json:"sticker_type,omitempty"`  // Optional. Type of stickers in the set, pass “regular” or “mask”. Custom emoji sticker sets can't be created via the Bot API at the moment. By default, a regular sticker set is created.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. A JSON-serialized object for position where the mask should be placed on faces
}

// createNewStickerSetParams createNewStickerSetParams contains CreateNewStickerSet's params
type createNewStickerSetParams struct {
	*CreateNewStickerSetOptions
	UserId int64  `json:"user_id"` // User identifier of created sticker set owner
	Name   string `json:"name"`    // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Title  string `json:"title"`   // Sticker set title, 1-64 characters
	Emojis string `json:"emojis"`  // One or more emoji corresponding to the sticker
}

// CreateNewStickerSet Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Returns True on success.
func (b *Bot) CreateNewStickerSet(userId int64, name string, title string, emojis string, optionalParams *CreateNewStickerSetOptions) (bool, error) {
	params := &createNewStickerSetParams{}

	params.UserId = userId
	params.Name = name
	params.Title = title
	params.Emojis = emojis
	params.CreateNewStickerSetOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "createNewStickerSet", params)
}

func (params *createNewStickerSetParams) HasUploadable() bool {
	return params.PngSticker.NeedsUpload() || params.TgsSticker.NeedsUpload() || params.WebmSticker.NeedsUpload()
}

// AddStickerToSetOptions AddStickerToSetOptions contains AddStickerToSet's optional params
type AddStickerToSetOptions struct {
	PngSticker   InputFile     `json:"png_sticker,omitempty"`   // Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	TgsSticker   InputFile     `json:"tgs_sticker,omitempty"`   // Optional. TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#animated-sticker-requirements for technical requirements
	WebmSticker  InputFile     `json:"webm_sticker,omitempty"`  // Optional. WEBM video with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/stickers#video-sticker-requirements for technical requirements
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. A JSON-serialized object for position where the mask should be placed on faces
}

// addStickerToSetParams addStickerToSetParams contains AddStickerToSet's params
type addStickerToSetParams struct {
	*AddStickerToSetOptions
	UserId int64  `json:"user_id"` // User identifier of sticker set owner
	Name   string `json:"name"`    // Sticker set name
	Emojis string `json:"emojis"`  // One or more emoji corresponding to the sticker
}

// AddStickerToSet Use this method to add a new sticker to a set created by the bot. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Animated stickers can be added to animated sticker sets and only to them. Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
func (b *Bot) AddStickerToSet(userId int64, name string, emojis string, optionalParams *AddStickerToSetOptions) (bool, error) {
	params := &addStickerToSetParams{}

	params.UserId = userId
	params.Name = name
	params.Emojis = emojis
	params.AddStickerToSetOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "addStickerToSet", params)
}

func (params *addStickerToSetParams) HasUploadable() bool {
	return params.PngSticker.NeedsUpload() || params.TgsSticker.NeedsUpload() || params.WebmSticker.NeedsUpload()
}

// setStickerPositionInSetParams setStickerPositionInSetParams contains SetStickerPositionInSet's params
type setStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int64  `json:"position"` // New sticker position in the set, zero-based
}

// SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(sticker string, position int64) (bool, error) {
	params := &setStickerPositionInSetParams{}

	params.Sticker = sticker
	params.Position = position

	return doHTTP[bool](b.client, b.url, "setStickerPositionInSet", params)
}

// deleteStickerFromSetParams deleteStickerFromSetParams contains DeleteStickerFromSet's params
type deleteStickerFromSetParams struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

// DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (b *Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	params := &deleteStickerFromSetParams{}

	params.Sticker = sticker

	return doHTTP[bool](b.client, b.url, "deleteStickerFromSet", params)
}

// SetStickerSetThumbOptions SetStickerSetThumbOptions contains SetStickerSetThumb's optional params
type SetStickerSetThumbOptions struct {
	Thumb InputFile `json:"thumb,omitempty"` // Optional. A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and height exactly 100px, or a TGS animation with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#animated-sticker-requirements for animated sticker technical requirements, or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-sticker-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated sticker set thumbnails can't be uploaded via HTTP URL.
}

// setStickerSetThumbParams setStickerSetThumbParams contains SetStickerSetThumb's params
type setStickerSetThumbParams struct {
	*SetStickerSetThumbOptions
	Name   string `json:"name"`    // Sticker set name
	UserId int64  `json:"user_id"` // User identifier of the sticker set owner
}

// SetStickerSetThumb Use this method to set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets only. Video thumbnails can be set only for video sticker sets only. Returns True on success.
func (b *Bot) SetStickerSetThumb(name string, userId int64, optionalParams *SetStickerSetThumbOptions) (bool, error) {
	params := &setStickerSetThumbParams{}

	params.Name = name
	params.UserId = userId
	params.SetStickerSetThumbOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "setStickerSetThumb", params)
}

func (params *setStickerSetThumbParams) HasUploadable() bool {
	return params.Thumb.NeedsUpload()
}

// AnswerInlineQueryOptions AnswerInlineQueryOptions contains AnswerInlineQuery's optional params
type AnswerInlineQueryOptions struct {
	CacheTime         int64  `json:"cache_time,omitempty"`          // Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool   `json:"is_personal,omitempty"`         // Optional. Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string `json:"next_offset,omitempty"`         // Optional. Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	SwitchPmText      string `json:"switch_pm_text,omitempty"`      // Optional. If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"` // Optional. Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

// answerInlineQueryParams answerInlineQueryParams contains AnswerInlineQuery's params
type answerInlineQueryParams struct {
	*AnswerInlineQueryOptions
	InlineQueryId string               `json:"inline_query_id"` // Unique identifier for the answered query
	Results       []*InlineQueryResult `json:"results"`         // A JSON-serialized array of results for the inline query
}

// AnswerInlineQuery Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
func (b *Bot) AnswerInlineQuery(inlineQueryId string, results []*InlineQueryResult, optionalParams *AnswerInlineQueryOptions) (bool, error) {
	params := &answerInlineQueryParams{}

	params.InlineQueryId = inlineQueryId
	params.Results = results
	params.AnswerInlineQueryOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "answerInlineQuery", params)
}

// answerWebAppQueryParams answerWebAppQueryParams contains AnswerWebAppQuery's params
type answerWebAppQueryParams struct {
	WebAppQueryId string             `json:"web_app_query_id"` // Unique identifier for the query to be answered
	Result        *InlineQueryResult `json:"result"`           // A JSON-serialized object describing the message to be sent
}

// AnswerWebAppQuery Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
func (b *Bot) AnswerWebAppQuery(webAppQueryId string, result *InlineQueryResult) (*SentWebAppMessage, error) {
	params := &answerWebAppQueryParams{}

	params.WebAppQueryId = webAppQueryId
	params.Result = result

	return doHTTP[*SentWebAppMessage](b.client, b.url, "answerWebAppQuery", params)
}

// SendInvoiceOptions SendInvoiceOptions contains SendInvoice's optional params
type SendInvoiceOptions struct {
	MessageThreadId           int64                 `json:"message_thread_id,omitempty"`             // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
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

// sendInvoiceParams sendInvoiceParams contains SendInvoice's params
type sendInvoiceParams struct {
	*SendInvoiceOptions
	ChatId        ChatID          `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title         string          `json:"title"`          // Product name, 1-32 characters
	Description   string          `json:"description"`    // Product description, 1-255 characters
	Payload       string          `json:"payload"`        // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken string          `json:"provider_token"` // Payment provider token, obtained via @BotFather
	Currency      string          `json:"currency"`       // Three-letter ISO 4217 currency code, see more on currencies
	Prices        []*LabeledPrice `json:"prices"`         // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
}

// SendInvoice Use this method to send invoices. On success, the sent Message is returned.
func (b *Bot) SendInvoice(chatId ChatID, title string, description string, payload string, providerToken string, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	params := &sendInvoiceParams{}

	params.ChatId = chatId
	params.Title = title
	params.Description = description
	params.Payload = payload
	params.ProviderToken = providerToken
	params.Currency = currency
	params.Prices = prices
	params.SendInvoiceOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendInvoice", params)
}

// CreateInvoiceLinkOptions CreateInvoiceLinkOptions contains CreateInvoiceLink's optional params
type CreateInvoiceLinkOptions struct {
	MaxTipAmount              int64   `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int64 `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string  `json:"provider_data,omitempty"`                 // Optional. JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string  `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64   `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int64   `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int64   `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool    `json:"need_name,omitempty"`                     // Optional. Pass True if you require the user's full name to complete the order
	NeedPhoneNumber           bool    `json:"need_phone_number,omitempty"`             // Optional. Pass True if you require the user's phone number to complete the order
	NeedEmail                 bool    `json:"need_email,omitempty"`                    // Optional. Pass True if you require the user's email address to complete the order
	NeedShippingAddress       bool    `json:"need_shipping_address,omitempty"`         // Optional. Pass True if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool    `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True if the user's phone number should be sent to the provider
	SendEmailToProvider       bool    `json:"send_email_to_provider,omitempty"`        // Optional. Pass True if the user's email address should be sent to the provider
	IsFlexible                bool    `json:"is_flexible,omitempty"`                   // Optional. Pass True if the final price depends on the shipping method
}

// createInvoiceLinkParams createInvoiceLinkParams contains CreateInvoiceLink's params
type createInvoiceLinkParams struct {
	*CreateInvoiceLinkOptions
	Title         string          `json:"title"`          // Product name, 1-32 characters
	Description   string          `json:"description"`    // Product description, 1-255 characters
	Payload       string          `json:"payload"`        // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken string          `json:"provider_token"` // Payment provider token, obtained via BotFather
	Currency      string          `json:"currency"`       // Three-letter ISO 4217 currency code, see more on currencies
	Prices        []*LabeledPrice `json:"prices"`         // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
}

// CreateInvoiceLink Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (b *Bot) CreateInvoiceLink(title string, description string, payload string, providerToken string, currency string, prices []*LabeledPrice, optionalParams *CreateInvoiceLinkOptions) (string, error) {
	params := &createInvoiceLinkParams{}

	params.Title = title
	params.Description = description
	params.Payload = payload
	params.ProviderToken = providerToken
	params.Currency = currency
	params.Prices = prices
	params.CreateInvoiceLinkOptions = optionalParams

	return doHTTP[string](b.client, b.url, "createInvoiceLink", params)
}

// AnswerShippingQueryOptions AnswerShippingQueryOptions contains AnswerShippingQuery's optional params
type AnswerShippingQueryOptions struct {
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"` // Optional. Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string            `json:"error_message,omitempty"`    // Optional. Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

// answerShippingQueryParams answerShippingQueryParams contains AnswerShippingQuery's params
type answerShippingQueryParams struct {
	*AnswerShippingQueryOptions
	ShippingQueryId string `json:"shipping_query_id"` // Unique identifier for the query to be answered
	Ok              bool   `json:"ok"`                // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
}

// AnswerShippingQuery If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
func (b *Bot) AnswerShippingQuery(shippingQueryId string, ok bool, optionalParams *AnswerShippingQueryOptions) (bool, error) {
	params := &answerShippingQueryParams{}

	params.ShippingQueryId = shippingQueryId
	params.Ok = ok
	params.AnswerShippingQueryOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "answerShippingQuery", params)
}

// AnswerPreCheckoutQueryOptions AnswerPreCheckoutQueryOptions contains AnswerPreCheckoutQuery's optional params
type AnswerPreCheckoutQueryOptions struct {
	ErrorMessage string `json:"error_message,omitempty"` // Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

// answerPreCheckoutQueryParams answerPreCheckoutQueryParams contains AnswerPreCheckoutQuery's params
type answerPreCheckoutQueryParams struct {
	*AnswerPreCheckoutQueryOptions
	PreCheckoutQueryId string `json:"pre_checkout_query_id"` // Unique identifier for the query to be answered
	Ok                 bool   `json:"ok"`                    // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
}

// AnswerPreCheckoutQuery Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (b *Bot) AnswerPreCheckoutQuery(preCheckoutQueryId string, ok bool, optionalParams *AnswerPreCheckoutQueryOptions) (bool, error) {
	params := &answerPreCheckoutQueryParams{}

	params.PreCheckoutQueryId = preCheckoutQueryId
	params.Ok = ok
	params.AnswerPreCheckoutQueryOptions = optionalParams

	return doHTTP[bool](b.client, b.url, "answerPreCheckoutQuery", params)
}

// setPassportDataErrorsParams setPassportDataErrorsParams contains SetPassportDataErrors's params
type setPassportDataErrorsParams struct {
	UserId int64                   `json:"user_id"` // User identifier
	Errors []*PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
func (b *Bot) SetPassportDataErrors(userId int64, errors []*PassportElementError) (bool, error) {
	params := &setPassportDataErrorsParams{}

	params.UserId = userId
	params.Errors = errors

	return doHTTP[bool](b.client, b.url, "setPassportDataErrors", params)
}

// SendGameOptions SendGameOptions contains SendGame's optional params
type SendGameOptions struct {
	MessageThreadId          int64                 `json:"message_thread_id,omitempty"`           // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	DisableNotification      bool                  `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool                  `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageId         int64                 `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                // Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

// sendGameParams sendGameParams contains SendGame's params
type sendGameParams struct {
	*SendGameOptions
	ChatId        int64  `json:"chat_id"`         // Unique identifier for the target chat
	GameShortName string `json:"game_short_name"` // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
}

// SendGame Use this method to send a game. On success, the sent Message is returned.
func (b *Bot) SendGame(chatId int64, gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	params := &sendGameParams{}

	params.ChatId = chatId
	params.GameShortName = gameShortName
	params.SendGameOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "sendGame", params)
}

// SetGameScoreOptions SetGameScoreOptions contains SetGameScore's optional params
type SetGameScoreOptions struct {
	Force              bool   `json:"force,omitempty"`                // Optional. Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Optional. Pass True if the game message should not be automatically edited to include the current scoreboard
	ChatId             int64  `json:"chat_id,omitempty"`              // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId          int64  `json:"message_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId    string `json:"inline_message_id,omitempty"`    // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
}

// setGameScoreParams setGameScoreParams contains SetGameScore's params
type setGameScoreParams struct {
	*SetGameScoreOptions
	UserId int64 `json:"user_id"` // User identifier
	Score  int64 `json:"score"`   // New score, must be non-negative
}

// SetGameScore Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (b *Bot) SetGameScore(userId int64, score int64, optionalParams *SetGameScoreOptions) (*Message, error) {
	params := &setGameScoreParams{}

	params.UserId = userId
	params.Score = score
	params.SetGameScoreOptions = optionalParams

	return doHTTP[*Message](b.client, b.url, "setGameScore", params)
}

// GetGameHighScoresOptions GetGameHighScoresOptions contains GetGameHighScores's optional params
type GetGameHighScoresOptions struct {
	ChatId          int64  `json:"chat_id,omitempty"`           // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int64  `json:"message_id,omitempty"`        // Optional. Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id,omitempty"` // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
}

// getGameHighScoresParams getGameHighScoresParams contains GetGameHighScores's params
type getGameHighScoresParams struct {
	*GetGameHighScoresOptions
	UserId int64 `json:"user_id"` // Target user id
}

// GetGameHighScores Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
func (b *Bot) GetGameHighScores(userId int64, optionalParams *GetGameHighScoresOptions) ([]*GameHighScore, error) {
	params := &getGameHighScoresParams{}

	params.UserId = userId
	params.GetGameHighScoresOptions = optionalParams

	return doHTTP[[]*GameHighScore](b.client, b.url, "getGameHighScores", params)
}
