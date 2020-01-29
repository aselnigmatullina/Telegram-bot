package main 


import (
   //чтобы получить обновление
   "net/http"
   "io/ioutil"
   "encoding/json"
   "fmt"
   "log"
   "bytes"
   "strconv"


)
//точка входа программы
func main(){
   //сюда запишем токен, который получили в телеграм

   botToken := "1010071670:AAF9mK4itljnMnD7cVk1qwl8pzK81-o0jn0"
   botApi   := "https://api.telegram.org/bot"
   botUrl   := botApi + botToken
   offset   := 0

   //бесконечный цикл
   for ;; {
	   updates, err := getUpdates(botUrl, offset)
	   if err != nil {
		   log.Println("Smth went wrong...", err.Error())
	   }
   }
   for _, update := range updates {
	 err = respond(botUrl, update)
	 offset = update.UpdateId + 1
   }
   fmt.Println(updates)
}


//функция, которая запрашивает обновления

func getUpdates(botUrl string, offset int) ( []Update, error ){

   resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
   if err != nil {
	   return nil, err
   }
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
	   return nil, err
   }
   var restResponse RestResponse
   err := json.Unmarshal(body, &restResponse)
   if err != nil {
	   return nil, err
   }
   return restResponse.Result, nil

}


//функция, которая отвечает на обновления 

func respond(botUrl string, update Update){
 var botMessage BotMessage
 botMessage.ChatId = update.Message.Chat.ChatId
 botMessage.Text   = update.Message.Text
 buf, err := json.Marshal(botMessage)
 if err != nil {
	 return nil
 }
 _, err = http.Post(botUrl + "/sendMessage", "application/json" , bytes.NewBuffer(buf))
 if err != nil {
	 return err
 }
 return nil
}


