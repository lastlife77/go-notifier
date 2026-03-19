// Package domain defines core models.
package domain

// Notify represents a notification that should send the message [Msg]
// at the specified time [Time].
type Notify struct {
	Id   string `json:"id" example:"1"`
	Msg  string `json:"msg" example:"Message 1"`
	Time string `json:"time" example:"yyyy-mm-dd HH:MM"` //yyyy-mm-dd HH:MM
}
