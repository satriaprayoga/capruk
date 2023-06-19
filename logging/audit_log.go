package logging

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/capruk/utils"
)

type AuditLog struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Level     string    `json:"level"`
	UUID      string    `json:"uuid"`
	FuncName  string    `json:"func_name"`
	FileName  string    `json:"file_name"`
	Line      int       `json:"line"`
	Time      string    `json:"time"`
	Message   string    `json:"message"`
}

func (a *AuditLog) SaveAudit() {

	a.ID = utils.GetTimeNow().Unix()
	a.Message = "API User : " + a.Message
	fmt.Printf("Calling the Logs: %v", a.Message)

}
