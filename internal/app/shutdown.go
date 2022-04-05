package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

// AppShutdown handler отвечающий за завершение работы приложения
func (a *application) AppShutdown(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Shutdown server requested")); err != nil {
		return
	}

	// Ничего не делам, если запрос на отключение уже отправлен
	// if a.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&a.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		a.shutdownReq <- true
	}()
}

// WaitShutdown handler ожидающий завершение сервера
func (a *application) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-a.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}
	log.Printf("Stoping http server ...")

	// Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	err := a.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}
