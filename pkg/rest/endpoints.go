package rest

import (
	"github.com/kjbreil/hass-ws/model"
	"net/url"
)

func (c *Client) GetApi() (*Api, error) {
	return Get[Api](c, "", url.Values{})
}

func (c *Client) GetStates() (*model.States, error) {
	return Get[model.States](c, "states", url.Values{})
}

func (c *Client) GetState(domainEntity string) (*model.State, error) {
	return Get[model.State](c, "states/"+domainEntity, url.Values{})
}
