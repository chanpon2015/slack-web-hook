# slack-web-hook

## Section
```
webhook := slack.Webhook{
    URL: "",
}
webhook.Msg.CreateSection()
if err := webhook.Send(); err != nil {
    fmt.Println(err)
}
```

### SectionOption
```
webhook.Msg.CreateSection(
    WithSetMrkdwnText(),
    WithSetImageAccessory(url, text),
)
```

# Button
```
webhook := slack.Webhook{
    URL: "",
}
webhook.Msg.CreateActions().
    AddButton("", "", "")
if err := webhook.Send(); err != nil {
    fmt.Println(err)
}
```

# Select
```
webhook := slack.Webhook{
    URL: "",
}
webhook.Msg.CreateActions().
    AddUserSelect()
if err := webhook.Send(); err != nil {
    fmt.Println(err)
}
```