package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()
	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}
	s.handler3 = TestEventHandler{ID: 3}
	s.event = TestEvent{Name: "Test", Payload: "Test"}
	s.event2 = TestEvent{Name: "Test2", Payload: "Test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.EqualError(err, ErrEventHandlerAlreadyRegistered.Error())
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event2.GetName()], 1)

	suite.eventDispatcher.Clear()
	suite.Len(suite.eventDispatcher.handlers, 0)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Len(suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
