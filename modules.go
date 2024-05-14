package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveUser.Id
	text := fmt.Sprintf(`👋 Hello <b>%s</b>, I am an Auto-Approval Bot. 
	
	<i>✨ Please add me to your channel or group as admin to automatically approve member join requests.</i> 

	⚡ Powered by <b>RioTraders™.</b>`, ctx.EffectiveUser.FirstName)

	_, err := b.SendMessage(
		chatId,
		text,
		&gotgbot.SendMessageOpts{
			ParseMode: "HTML",
		},
	)
	return err

}

func autoApproveChatJoinRequest(b *gotgbot.Bot, ctx *ext.Context) error {
	userId := ctx.EffectiveUser.Id
	chatId := ctx.EffectiveChat.Id
	ok, err := b.ApproveChatJoinRequest(chatId, userId, &gotgbot.ApproveChatJoinRequestOpts{})
	text := fmt.Sprintf(`<b>🎉 Your request to join on <b>%s</b> has been approved. Enjoy!</b>

	<i>👥 You can add me to your channel or group as admin to automatically approve chat join requests.</i>


	<b>⚡ Powered by RioTraders™.</b>
	`, ctx.EffectiveChat.Title)
	if ok {
		_, err := b.SendMessage(userId, text, &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
		return err

	}

	return err
}
