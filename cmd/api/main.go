package main

import (
	"fmt"
	"net/http"

	"github.com/chi-go/chi"
	"github.com/izzuddinafif/coin-api-go/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportcaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
   ____      _            _    ____ ___ 
  / ___|___ (_)_ __      / \  |  _ \_ _|
 | |   / _ \| | '_ \    / _ \ | |_) | | 
 | |__| (_) | | | | |  / ___ \|  __/| | 
  \____\___/|_|_| |_| /_/   \_\_|  |___|
                                        `)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		return nil, err
	}
}
