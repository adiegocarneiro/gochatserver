package routes

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Username *string `json:"username"`
	Function *string `json:"function"`
	Params   *string `json:"params,required"`
}

func wsRouter(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			WebsocketController(c, mt, msg)
			// if err = c.WriteMessage(mt, msg); err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }
		}

	}))
}

func WebsocketController(c *websocket.Conn, messageType int, msg []byte) {
	message := Message{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		c.WriteMessage(messageType, []byte("Mensagem não está no formato correto."))
		return
	}

	if message.Username == nil {
		c.WriteMessage(messageType, []byte("Erro! Você precisa enviar seu nome de usuário!"))
		return
	}
	if message.Function == nil {
		c.WriteMessage(messageType, []byte("Erro! Você precisa enviar a função desejada!"))
		return
	}

	switch *message.Function {
	case "getUserData": //TODO: Preencher com as funções do sistema
		c.WriteMessage(messageType, []byte("Você acesso a função getUserData! Parabéns!"))
		return
	default:
		c.WriteMessage(messageType, []byte("Função desconhecida!"))
		return
	}

}
