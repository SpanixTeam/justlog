package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

// RandomQuoteJSON response when request a random message
type RandomQuoteJSON struct {
	Channel     string    `json:"channel"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Message     string    `json:"message"`
	Timestamp   timestamp `json:"timestamp"`
}

// swagger:route GET /channel/{channel}/user/{username}/random logs channelUserLogsRandom
//
// Obtén una línea aleatoria de un usuario en un canal determinado
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/userid/{userid}/random logs channelIdUserIdLogsRandom
//
// Obtén una línea aleatoria de un usuario en un canal determinado
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/user/{user}/random logs channelIdUserLogsRandom
//
// Obtén una línea aleatoria de un usuario en un canal determinado
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channel/{channel}/userid/{userid}/random logs channelUserIdLogsRandom
//
// Obtén una línea aleatoria de un usuario en un canal determinado
//
//	Produces:
//	- application/json
//	- text/plain
//
//	Responses:
//	  200: chatLog
func (s *Server) getRandomQuote(request logRequest) (*chatLog, error) {
	rawMessage, err := s.fileLogger.ReadRandomMessageForUser(request.channelid, request.userid)
	if err != nil {
		return &chatLog{}, err
	}
	parsedMessage := twitch.ParseMessage(rawMessage)

	chatMsg := createChatMessage(parsedMessage)

	return &chatLog{Messages: []chatMessage{chatMsg}}, nil
}

// swagger:route GET /list logs list
//
// Lista los registros disponibles de un usuario o canal, la respuesta del canal también incluye el día. OpenAPI 2 no admite ahora mismo múltiples respuestas con el mismo código http.
//
//	Produces:
//	- application/json
//	- text/plain
//
//	Schemes: https
//
//	Responses:
//	  200: logList
func (s *Server) writeAvailableLogs(w http.ResponseWriter, r *http.Request, q url.Values) {
	channelid := q.Get("channelid")
	userid := q.Get("userid")

	if userid == "" {
		logs, err := s.fileLogger.GetAvailableLogsForChannel(channelid)
		if err != nil {
			http.Error(w, "failed to get available channel logs: "+err.Error(), http.StatusNotFound)
			return
		}

		writeCacheControl(w, r, time.Hour)
		writeJSON(&channelLogList{logs}, http.StatusOK, w, r)
		return
	}

	logs, err := s.fileLogger.GetAvailableLogsForUser(channelid, userid)
	if err != nil {
		http.Error(w, "failed to get available user logs: "+err.Error(), http.StatusNotFound)
		return
	}

	writeCacheControl(w, r, time.Hour)
	writeJSON(&logList{logs}, http.StatusOK, w, r)
}

// swagger:route GET /channel/{channel}/user/{username} logs channelUserLogs
//
// Obtén los registros del usuario en el canal del mes actual
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/userid/{userid} logs channelIdUserIdLogs
//
// Obtén los registros del usuario en el canal del mes actual
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/user/{username} logs channelIdUserLogs
//
// Obtén los registros del usuario en el canal del mes actual
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channel/{channel}/userid/{userid} logs channelUserIdLogs
//
// Obtén los registros del usuario en el canal del mes actual
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channel/{channel}/user/{username}/{year}/{month} logs channelUserLogsYearMonth
//
// Obtén los registros del usuario en el canal del año y mes dados
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/userid/{userid}/{year}/{month} logs channelIdUserIdLogsYearMonth
//
// Obtén los registros del usuario en el canal del año y mes dados
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channelid/{channelid}/user/{username}/{year}/{month} logs channelIdUserLogsYearMonth
//
// Obtén los registros del usuario en el canal del año y mes dados
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Responses:
//       200: chatLog

// swagger:route GET /channel/{channel}/userid/{userid}/{year}/{month} logs channelUserIdLogsYearMonth
//
// Obtén los registros del usuario en el canal del año y mes dados
//
//	Produces:
//	- application/json
//	- text/plain
//
//	Responses:
//	  200: chatLog
func (s *Server) getUserLogs(request logRequest) (*chatLog, error) {
	logMessages, err := s.fileLogger.ReadLogForUser(request.channelid, request.userid, request.time.year, request.time.month)
	if err != nil {
		return &chatLog{}, err
	}

	if request.reverse {
		reverseSlice(logMessages)
	}

	logResult := createLogResult()

	for _, rawMessage := range logMessages {
		logResult.Messages = append(logResult.Messages, createChatMessage(twitch.ParseMessage(rawMessage)))
	}

	return &logResult, nil
}

func (s *Server) getUserLogsRange(request logRequest) (*chatLog, error) {

	fromTime, toTime, err := parseFromTo(request.time.from, request.time.to, userHourLimit)
	if err != nil {
		return &chatLog{}, err
	}

	var logMessages []string

	logMessages, _ = s.fileLogger.ReadLogForUser(request.channelid, request.userid, fmt.Sprintf("%d", fromTime.Year()), fmt.Sprintf("%d", int(fromTime.Month())))

	if fromTime.Month() != toTime.Month() {
		additionalMessages, _ := s.fileLogger.ReadLogForUser(request.channelid, request.userid, fmt.Sprint(toTime.Year()), fmt.Sprint(toTime.Month()))

		logMessages = append(logMessages, additionalMessages...)
	}

	if request.reverse {
		reverseSlice(logMessages)
	}

	logResult := createLogResult()

	for _, rawMessage := range logMessages {
		parsedMessage := twitch.ParseMessage(rawMessage)

		switch message := parsedMessage.(type) {
		case *twitch.PrivateMessage:
			if message.Time.Unix() < fromTime.Unix() || message.Time.Unix() > toTime.Unix() {
				continue
			}
		case *twitch.ClearChatMessage:
			if message.Time.Unix() < fromTime.Unix() || message.Time.Unix() > toTime.Unix() {
				continue
			}
		case *twitch.UserNoticeMessage:
			if message.Time.Unix() < fromTime.Unix() || message.Time.Unix() > toTime.Unix() {
				continue
			}
		}

		logResult.Messages = append(logResult.Messages, createChatMessage(parsedMessage))
	}

	return &logResult, nil
}
