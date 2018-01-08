package mem

import (
	. "github.com/ankyra/escape-inventory/dao/types"
)

func (a *dao) GetFeedPage() ([]*FeedEvent, error) {
	result := []*FeedEvent{}
	for i := len(a.events) - 1; i >= 0 && len(result) < FeedPageSize; i-- {
		result = append(result, a.events[i])
	}
	return result, nil
}

func (a *dao) GetProjectFeedPage(project string) ([]*FeedEvent, error) {
	result := []*FeedEvent{}
	for i := len(a.events) - 1; i >= 0 && len(result) < FeedPageSize; i-- {
		if a.events[i].Project == project {
			result = append(result, a.events[i])
		}
	}
	return result, nil
}

func (a *dao) GetFeedPageByGroups(readGroups []string) ([]*FeedEvent, error) {
	result := []*FeedEvent{}
	for i := len(a.events) - 1; i >= 0 && len(result) < FeedPageSize; i-- {
		allowedGroups, found := a.acls[a.events[i].Project]
		if !found {
			continue
		}
		for group, _ := range allowedGroups {
			found := false
			for _, g := range readGroups {
				if g == group {
					found = true
				}
			}
			if found {
				result = append(result, a.events[i])
			}
		}
	}
	return result, nil
}

func (a *dao) AddFeedEvent(event *FeedEvent) error {
	a.events = append(a.events, event)
	return nil
}
