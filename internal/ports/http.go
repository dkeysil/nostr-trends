package ports

import (
	"net/http"

	"github.com/dkeysil/nostr-trends/internal/app"
	"github.com/dkeysil/nostr-trends/internal/app/query"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Handler(queries app.Queries) http.Handler {
	router := chi.NewRouter()

	router.Get("/trends", trendsHandler(queries.WordTrends))

	return router
}

type trendsHandlerResponse struct {
	Words []struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
	} `json:"words"`
}

func trendsHandler(handler query.WordTrends) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		words, err := handler.Handle(r.Context(), query.Params{
			Count: 10,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := trendsHandlerResponse{
			Words: make([]struct {
				Word  string `json:"word"`
				Count int    `json:"count"`
			}, len(*words)),
		}

		for i, word := range *words {
			response.Words[i] = struct {
				Word  string `json:"word"`
				Count int    `json:"count"`
			}{
				Word:  word.Word,
				Count: word.Count,
			}
		}

		render.JSON(w, r, response)
	}
}
