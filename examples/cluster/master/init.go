package master

import (
	"github.com/popfendi/nano/component"
	"github.com/popfendi/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	// Topic service
	topicService = newTopicService()
	// ... other services
)

func init() {
	Services.Register(topicService)
}

func OnSessionClosed(s *session.Session) {
	topicService.userDisconnected(s)
}
