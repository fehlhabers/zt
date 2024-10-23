package timer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/model"
)

type Timer struct {
	Timer int    `json:"timer,omitempty"`
	User  string `json:"user,omitempty"`
}

type BreakTimer struct {
	BreakTimer int    `json:"breaktimer,omitempty"`
	User       string `json:"user,omitempty"`
}

const timerBaseUrl = "https://timer.mob.sh"

func SetTimer(zt model.Ztream, team model.Team) error {
	client := http.DefaultClient
	timerRoomUrl := fmt.Sprintf("%s/%s-%s", timerBaseUrl, team.Name, zt.Name)

	body, _ := json.Marshal(Timer{
		Timer: 10,
		User:  "test",
	})
	requestWriter := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPut, timerRoomUrl, requestWriter)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Warn("Timer service responded with non-OK", "status", res.Status)
			return nil
		}
		log.Warn("Timer service responded with non-OK", "status", res.Status, "message", string(body))
		return nil
	}

	log.Infof("Started timer at %s", timerRoomUrl)
	return nil
}
