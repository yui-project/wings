package app

import (
	"log"
	"net/http"

	"github.com/ut-issl/wings/internal/core/service"
	"github.com/ut-issl/wings/internal/handler"
	"go.uber.org/fx"
)

func Run() {
	// DIコンテナの作成とエラーハンドリング
	app := fx.New(
		fx.Provide(

			handler.NewHTTPServer,
			handler.Route,

			service.NewCmdService,
		),
		fx.Invoke(func(srv *http.Server) {
			if srv != nil {
				log.Println("HTTP server is ready to handle requests at :8080")
			} else {
				log.Fatal("HTTP server is nil")
			}
		}),
	)

	// DIの初期化エラーをチェック
	if err := app.Err(); err != nil {
		log.Fatalf("Dependency injection setup failed: %v", err)
	}

	// アプリケーション実行
	app.Run()
}
