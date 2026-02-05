package main

import (
	"net/http"

	httpapi "github.com/minhtungle/GiapTech.GoTemp/backend/src/GiapTech.GoTemp.HttpApi"
	"github.com/minhtungle/GiapTech.GoTemp/backend/src/GiapTech.GoTemp.Shared.Hosting/module"
)

func main() {
	ctx := &module.Context{Values: map[string]any{}}

	loader := module.NewLoader(
		httpapi.HttpApiModule{},
	)
	if err := loader.Init(ctx); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", mux)
}
