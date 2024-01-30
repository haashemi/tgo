# TGO Filters

TGO Filters are set of functions that can be used with routers to determine which handler should be called and which should not.

## Example:

```go
package main

import (
    ...
    "github.com/haashemi/tgo/filters"
)

func main() {
    // ...

    mr := messages.NewRouter()

    // startCommandHandler will only get called if the Command filter approves it.
    mr.Handle(filters.Command("start", botUsername), startCommandHandler)

    bot.AddRouter(mr)

    // ...
}
```

## Built-in filters:

### Logical filters:

- `filters.True()`
- `filters.False()`
- `filters.And(...)`
- `filters.Or(...)`
- `filters.Not(...)`

### Message filters:

- `filters.IsPrivate()`
- `filters.Command(...)`
- `filters.Commands(...)`

### General filters:

General filters are currently working with `Message`, `CallbackQuery`, and `InlineQuery`.

- `filters.Text(...)`
- `filters.Texts(...)`
- `filters.WithPrefix(...)`
- `filters.WithSuffix(...)`
- `filters.Regex(...)`
- `filters.Whitelist(...)`

### Update filters:

- `filters.HasMessage()`
- `filters.IsMessage()`
- `filters.IsEditedMessage()`
- `filters.IsChannelPost()`
- `filters.IsEditedChannelPost()`
- `filters.IsInlineQuery()`
- `filters.IsChosenInlineResult()`
- `filters.IsCallbackQuery()`
- `filters.IsShippingQuery()`
- `filters.IsPreCheckoutQuery()`
- `filters.IsPoll()`
- `filters.IsPollAnswer()`
- `filters.IsMyChatMember()`
- `filters.IsChatMember()`
- `filters.IsChatJoinRequest()`

## How to implement your own filter?

It's simple! just pass your filter function to `filters.NewFilter` and you're done!

Here is an example:

```go
// It can be with a variable, fastest way.
var myInlineFilter = filters.NewFilter(func(update *tgo.Update) bool {
    // your filter's logic goes here
})

// Or it can be its own separated function.
func MyFilter() *tgo.Filter {
    return filters.NewFilter(func(update *tgo.Update) bool {
        // your filter's logic goes here
    })
}
```
