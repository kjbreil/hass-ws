package model

import (
	"time"
)

type Histories []*History

type History struct {
	Entity string
	Domain string
	Start  time.Time
	End    time.Time
	States []*State
}

func (hs *Histories) Get(domain, entity string) *History {
	for _, h := range *hs {
		if h.Entity == entity && h.Domain == domain {
			return h
		}
	}
	return nil
}
