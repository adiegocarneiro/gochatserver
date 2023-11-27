package routes

import (
	"encoding/json"
	"fmt"
	"gochatserver/app/database/entities"
	"gochatserver/app/database/repositories"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Message struct {
	UserID   *string `json:"id_usuario"`
	Function *string `json:"funcao"`
	Params   *string `json:"parametros"`
}

func wsRouter(db *gorm.DB, app *fiber.App) {
	auxctx := &fiber.Ctx{}
	app.Use("/ws", func(c *fiber.Ctx) error {
		auxctx = c
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(wsc *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(wsc.Locals("allowed"))  // true
		log.Println(wsc.Query("v"))         // 1.0
		log.Println(wsc.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = wsc.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			WebsocketController(wsc, auxctx, db, mt, msg)
			// if err = c.WriteMessage(mt, msg); err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }
		}

	}))
}

func WebsocketController(wsc *websocket.Conn, ctx *fiber.Ctx, db *gorm.DB, messageType int, msg []byte) {
	message := Message{}
	err := json.Unmarshal(msg, &message)
	repo := &repositories.Repository{
		DB: db,
	}

	if err != nil {
		wsc.WriteMessage(messageType, []byte("Mensagem não está no formato correto."))
		return
	}

	if message.UserID == nil {
		wsc.WriteMessage(messageType, []byte("Erro! Você precisa enviar seu id de usuário!"))
		return
	}
	if message.Function == nil {
		wsc.WriteMessage(messageType, []byte("Erro! Você precisa enviar a função desejada!"))
		return
	}

	switch *message.Function {
	case "getUserData": //TODO: Preencher com as funções do sistema
		wsc.WriteMessage(messageType, []byte("Você acesso a função getUserData! Parabéns!"))
		return
	case "sendMessage":
		params := strings.Split(*message.Params, "&")
		var roomId uint
		var userId uint
		var messageToSend string

		value, err := strconv.ParseUint(*message.UserID, 0, 32)
		if err != nil {
			fmt.Println(err)
		}

		userId = uint(value)

		for _, param := range params {
			if strings.Contains(param, "id_sala") {
				value, err := strconv.ParseUint(strings.Split(param, "=")[1], 0, 32)
				if err != nil {
					fmt.Println(err)
				}
				roomId = uint(value)
			}
			if strings.Contains(param, "mensagem") {
				messageToSend = strings.Split(param, "=")[1]
			}
		}

		chatMessage := entities.ChatMessage{
			UserId:  &userId,
			RoomId:  &roomId,
			Message: messageToSend,
		}
		response := repo.CreateMessage(ctx, &chatMessage)

		wsc.WriteMessage(messageType, []byte(response.Message))
		return
	default:
		wsc.WriteMessage(messageType, []byte("Função desconhecida!"))
		return
	}

}
