package telegrampkg

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AlekseySauron/tomato/models"
	"github.com/AlekseySauron/tomato/pkg/dbpkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

var bot *tgbotapi.BotAPI
var currentStatus string

type button struct {
	source string
	name   string
	volume int
}

var buttonsTimeouts = [2]button{
	{"buttonsTimeouts", "20 min", 20},
	{"buttonsTimeouts", "25 min", 25},
}

var buttonsNewOrQueue = [2]button{
	{"NewOrQueue", "Есть новая задача?", 0},
	{"NewOrQueue", "Только проверим очередь", 0},
}

func Start(repo dbpkg.DBRepository) {
	if bot == nil {
		initBot()
		currentStatus = "start"
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	// updates, err := bot.GetUpdatesChan(u)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		//if update.Message != nil {
		user, _ := repo.GetUser(getChatId(update))
		defer repo.Save(user)
		if update.Message.IsCommand() {
			commandHandler(user, update.Message.Command())
			continue
		}
		if update.CallbackQuery != nil {
			callbackHandler(user, update.CallbackData())
			continue
		}
	}
}

func commandHandler(u *models.User, command string) {
	go func(u models.User) {
		time.Sleep(25 * time.Minute)
		sendMessage(u.TelegramID, "Task done", nil)
	}(*u)
}

func callbackHandler(u *models.User, data string) {

}

func selectTimeout(update tgbotapi.Update) {
	chatId := getChatId(update)

	msg := tgbotapi.NewMessage(chatId, "Select timeout")
	//bot.Send(msg)

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	//keyboard := tgbotapi.InlineKeyboardButton{}

	for buttonId := range buttonsTimeouts {

		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(buttonsTimeouts[buttonId].name, buttonsTimeouts[buttonId].name)
		row = append(row, btn)
		// keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func initBot() {
	var err error

	token := viper.GetString("telegram.token")
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err, " token", token)
		return
	}
}

func Send(chat_id string, res string) {
	chat_id_int, err := strconv.Atoi(chat_id)
	if err != nil || chat_id_int <= 0 {
		log.Fatal("Ошибка формата chat_id:", err)
		return
	}

	// if bot == nil {
	// 	initBot()
	// }

	bot.Send(tgbotapi.NewMessage(int64(chat_id_int), fmt.Sprint(res)))
}

func addNewOrQueue(update tgbotapi.Update) {
	chatId := getChatId(update)

	msg := tgbotapi.NewMessage(chatId, "Select button")
	//bot.Send(msg)

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	//keyboard := tgbotapi.InlineKeyboardButton{}

	for buttonId := range buttonsNewOrQueue {

		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(buttonsNewOrQueue[buttonId].name, buttonsNewOrQueue[buttonId].name)
		row = append(row, btn)
		// keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func startTask(update tgbotapi.Update, min int) {
	chatId := getChatId(update)

	msg := tgbotapi.NewMessage(chatId, fmt.Sprint("Start ", min))
	bot.Send(msg)

}

func addNew(update tgbotapi.Update) {
	chatId := getChatId(update)

	msg := tgbotapi.NewMessage(chatId, "Введи название новой задачи:")
	bot.Send(msg)

}

func getUserId(update tgbotapi.Update) int64 {
	var userId int64
	if update.Message == nil {
		userId = update.CallbackQuery.Message.Contact.UserID
	} else {
		userId = update.Message.Contact.UserID
	}
	return userId
}

func getChatId(update tgbotapi.Update) int64 {
	var chatId int64
	if update.Message == nil {
		chatId = update.CallbackQuery.Message.Chat.ID
	} else {
		chatId = update.Message.Chat.ID
	}
	return chatId
}
