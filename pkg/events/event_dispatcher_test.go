package events

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name string
	Data interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetData() interface{} {
	return e.Data
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.event = TestEvent{Name: "test", Data: "test"}
	suite.event2 = TestEvent{Name: "test2", Data: "test2"}
}

func (suite *EventDispatcherTestSuite) Test_Add() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) Test_Add_AlreadyExists() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.NotNil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) Test_Clear() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event2.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	suite.dispatcher.Clear()
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) Test_Has() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	has := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(has)

	has = suite.dispatcher.Has(suite.event.GetName(), &suite.handler2)
	suite.False(has)
}

func (suite *EventDispatcherTestSuite) Test_Remove() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Add(suite.event2.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	err = suite.dispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Remove(suite.event2.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event2.GetName()]))
}

func (suite *EventDispatcherTestSuite) Test_Remove_NotFound() {
	err := suite.dispatcher.Add(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.NotNil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) Test_Dispatch() {
	mockHandler := &MockHandler{}
	mockHandler.On("Handle", &suite.event)
	suite.dispatcher.Add(suite.event.GetName(), mockHandler)
	suite.dispatcher.Dispatch(&suite.event)
	mockHandler.AssertExpectations(suite.T())
	mockHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
