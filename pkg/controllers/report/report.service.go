package report

import (
	"strings"
	"time"

	"github.com/Kaciis/go-reporter/pkg/controllers/app"
	"github.com/bwmarrin/discordgo"
)

func SendMessage(message discordgo.MessageEmbed) {

	app.App.Session.ChannelMessageSendEmbed(app.App.ChannelID, &message)

}
func GenerateReport(report ReportBody) (msg discordgo.MessageEmbed) {

	msg = discordgo.MessageEmbed{
		Title:       generateTitle(report.Type, report.Iniciator),
		Color:       generateColor(report.Type),
		Description: generateDescription(report.Message),
	}

	return msg
}

func generateTitle(reportType, iniciator string) (title string) {
	var emote string

	switch reportType {
	case "error":
		emote = ":bangbang:"
	case "warning":
		emote = ":warning:"
	default:
		emote = ":information_source:"
	}

	title = emote + " " + strings.ToTitle(reportType) + " " + emote + " => " + iniciator + " <= " + emote + " " + strings.ToTitle(reportType) + " " + emote

	return title
}

func generateColor(reportType string) (color int) {
	switch reportType {
	case "error":
		color = 16711680
	case "warning":
		color = 16776960
	default:
		color = 65280
	}

	return color
}

func generateDescription(message string) (description string) {
	description = "```" + time.Now().Format(time.UnixDate) + "\n" + message + "```"

	return description
}
