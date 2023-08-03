package hass_ws

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/kjbreil/hass-ws/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (c *Client) GetHistory(start, end time.Time, entities ...string) (*model.Histories, error) {

	u := url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%d", c.config.Host, c.config.Port), Path: fmt.Sprintf("/api/history/period/%s", start.Local().Format(time.RFC3339))}

	values := u.Query()
	values.Set("end_time", end.Local().Format(time.RFC3339))
	values.Set("filter_entity_id", strings.Join(entities, ","))

	u.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+c.config.Token)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	fmt.Println(u.String())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var entitiesStates [][]*model.State
	err = json.Unmarshal(body, &entitiesStates)
	if err != nil {
		return nil, err
	}

	var histories model.Histories
	for _, states := range entitiesStates {
		if len(states) == 0 {
			continue
		}
		histories = append(histories, &model.History{
			Entity: states[0].Entity(),
			Domain: states[0].Domain(),
			Start:  start,
			End:    end,
			States: states,
		})
	}

	return &histories, nil
}
