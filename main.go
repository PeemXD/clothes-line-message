// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	log.Println("Open!")
	bot, err := linebot.New(
		"22bad19b24bc87abc675aabed2ad0e74",
		"OuFHr2lU/8wFp/RVDLKaNQYNjbJB1T7FLZFvT1I3f6pmQ171lJ+c1A4PHgq+OdXc3pG+CTwFMp+sW8WUTpKt6DdZwdaWfIkpIn/IY1Ux5uAWNebTeAGH+ahbWxiKidFD7NOmpifb4Apoki2IOL4IMwdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "swap" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("My token\n80")).Do(); err != nil {
							log.Print(err)
						}
					}
				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Eiei")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {

	// })
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+"8080", nil); err != nil {
		log.Fatal(err)
	}
}
