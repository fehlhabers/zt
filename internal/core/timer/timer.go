package timer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/domain"
)

const (
	timerUrl = "https://timer.mob.sh"
)

type TimerReq struct {
	Timer int    `json:"timer,omitempty"`
	User  string `json:"user,omitempty"`
}

func Start(zt *domain.ZtState) {

	client := http.DefaultClient
	safeZtream := strings.ReplaceAll(zt.CurZtream.Name, "/", "-")
	url := fmt.Sprintf("%s/%s-%s", timerUrl, zt.TeamName, safeZtream)
	reqBody := &TimerReq{
		Timer: zt.TeamConfig.SessionDurMins,
		User:  zt.User,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Error("Unable to create timer request")
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil || res.StatusCode != 202 {
		log.Error("Could not start timer")
		return
	}

	log.Infof("Started timer at: %s", url)
}
