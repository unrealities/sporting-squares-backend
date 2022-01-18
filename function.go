package function

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

// GetGameDataByWeek returns useful (to Sporting Squares) game information
// ex. POST request:
// https://us-central1-sporting-squares-backend.cloudfunctions.net/GetGameData
func GetGameDataByWeek(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusOK)
		return
	}

	ctx := r.Context()
	s, err := InitService(ctx) // Execution Time: ~300ms
	if err != nil {
		s.HandleFatalError("error initializing service", err)
	}
	defer s.ErrorReporter.Close()
	defer s.FirestoreClient.Close()
	defer s.Logger.Close()
	defer s.TraceSpan.End()

	// Extract
	games, err := nflStats.GetGames()
	if err != nil {
		s.HandleFatalError("error getting the game data", err)
	}
	s.DebugMsg("successfully fetched data")

	// Transform
	// TODO: Decouple Extract & Transform

	// Loade
	dataKey := fmt.Sprintf("y%dw%dt%d", games[0].Year, games[0].Week, games[0].SeasonType)
	_, err = s.FirestoreClient.Collection(s.DBCollection).Doc(dataKey).Set(ctx, games) // Execution Time: ~ 3500ms
	if err != nil {
		s.HandleFatalError("error persisting data to Firebase", err)
	}

	// Send Response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
