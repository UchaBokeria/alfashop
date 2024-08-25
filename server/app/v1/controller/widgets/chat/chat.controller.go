package chat

import (
	"fmt"
	"main/internal/model"
	"main/pkg/engine/controller"
	mailer "main/pkg/services/mail"
	"main/pkg/storage"
	chat "main/server/app/v1/view/widgets"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
    Mu   sync.Mutex
    Conn *websocket.Conn
}

var (
	mu sync.Mutex
	chats = make(map[*Client] int)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// CheckOrigin: func(r *http.Request) bool { return true },
	}
	PickUpLine = "მოგესალმებით, გთხოვთ დაგველოდოთ მალე გიპასუხებთ."
)

func index(ctx *controller.Context, Parameters *MailStrategyDto) error {
	if Parameters.Email == "" || Parameters.Fullname == "" {
		return fmt.Errorf("parameters are not provided :: Email: `" + Parameters.Email +"`, Fullname: `"+ Parameters.Fullname + "`")
	}

	return ctx.Html(chat.ChatMessages(Parameters.Email, Parameters.Fullname, ""))
}

func MailStrategy(ctx *controller.Context, Parameters *MailStrategyDto) error {
	if Parameters.Email == "" || Parameters.Fullname == "" {
		return fmt.Errorf("parameters are not provided :: Email: `" + Parameters.Email +"`, Fullname: `"+ Parameters.Fullname + "`")
	}

	sent, _ := mailer.Send(mailer.Config{
		Subject: Parameters.Fullname + " გწერთ ელფოსტა ( " + Parameters.Email + " )",
		Body: Parameters.Message,
		To: "ucha2bokeria@gmail.com",
	})

	if sent { return ctx.Html(chat.NewMessage(Parameters.Fullname, Parameters.Message)) }
	return ctx.Html(chat.NewMessageError(Parameters.Fullname, Parameters.Message))
}

func NewChat(ctx *controller.Context, Parameters *NewChatDto) error {
	/* Check And Transform Provided Parameters For New Chat */
	if Parameters.Email == "" || Parameters.Fullname == "" {
		return fmt.Errorf("parameters are not provided :: Email: `" + Parameters.Email +"`, Fullname: `"+ Parameters.Fullname + "`")
	}

	client, _ := SetupWS(ctx, Parameters)

    defer func() {
        mu.Lock()
        defer mu.Unlock()
        delete(chats, client)
        client.Conn.Close()
    }()

	for { GetMsg(client) }
}

func SetupWS(ctx *controller.Context, Parameters *NewChatDto) (*Client, error){
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		fmt.Print("Websocket Upgrading Error: ", err)
		return nil, err
	}

	/* Create New Record In DB For New Chat */
	var ChatRecord model.Chat = model.Chat{
		TypeID: 1,
		ChatStatusID: 1,
		Email: Parameters.Email,
		Fullname: Parameters.Fullname,
	}

    client := &Client{Conn: ws}

    mu.Lock()
	if err := storage.DB.Create(&ChatRecord).Error; err != nil {
		fmt.Print("Creating New Chat Error: ", err)
		return nil, err
    }

    chats[client] = int(ChatRecord.ID)
	SendMsg(client, PickUpLine)
    mu.Unlock()

	return client, nil
}

func OldChat(ctx *controller.Context) error {
	return ctx.String(200, "")
}

func GetMsg(client *Client) error {
	_, msg, err := client.Conn.ReadMessage()
	if err != nil {
		fmt.Print("Websocket Reading Error: ", err)
		return err
	} else if string(msg) != "" {
		storage.DB.Create(&model.Chat_letters{
			Body: string(msg),
			From: "Client",
			To: "Admin",
			LetterStatusID: 1,
			ChatID: uint(chats[client]),
		})
		
		// fmt.Printf("%s\n", msg)
	}

	return nil
}

func SendMsg(client *Client, msg string) error {
	err := client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))

	if err != nil {
		fmt.Print("Websocket Writing Error: ", err)
		return err
	} else if msg != "" && msg != PickUpLine {
		storage.DB.Create(&model.Chat_letters{
			Body: string(msg),
			From: "Admin",
			To: "Client",
			LetterStatusID: 1,
			ChatID: uint(chats[client]),
		})
	}

	return nil
}