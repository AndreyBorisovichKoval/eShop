// C:\GoProject\src\eShop\server\server.go

package server

import (
	"context"
	"net/http"
	"time"
)

// Структура Server представляет HTTP сервер с полем для http.Server...
type Server struct {
	httpServer *http.Server // Встроенный HTTP сервер...
}

// Run запускает HTTP сервер на указанном порту с переданным обработчиком...
// Настраивает сервер с максимальным размером заголовка 1 MB и таймаутами чтения/записи по 10 секунд...
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB
		ReadTimeout:    10 * time.Second, // 10 секунд
		WriteTimeout:   10 * time.Second, // 10 секунд
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown корректно завершает работу сервера, закрывая все активные соединения...
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
