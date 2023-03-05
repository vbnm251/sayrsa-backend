package handler

import (
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/cha1l/sayrsa-2.0/pkg/service"
	"github.com/gorilla/mux"
)

const (
	createConversationAction = "create_conv"
	sendMessageAction        = "send_message"
)

type Handler struct {
	service *service.Service
	clients map[string]Client
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
		clients: make(map[string]Client),
	}
}

type Client struct {
	connection *websocket.Conn
}

func NewClient(conn *websocket.Conn) Client {
	return Client{
		connection: conn,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	//Authorization
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-up", h.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/sign-in", h.SignIn).Methods(http.MethodPost)

	//WebSockets handler
	api := r.PathPrefix("/api").Subrouter()
	api.Use(h.AuthorizationMiddleware)
	api.HandleFunc("/", h.wsHandler)

	return r
}
