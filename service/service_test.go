package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learnGO/service/mocks"
	"testing"
)

func TestSpammyMakser(t *testing.T) {
	testTable := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "Ссылка среди текста",
			text:     "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			name:     "2 ссылки среди текста",
			text:     "Here's my spammy page: http://hehefouls.netHAHAHA see you. http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you. http://******************* see you.",
		},
		{
			name:     "Ссылка с цифрами",
			text:     "Here's my spammy page: http://hehefo23452444uls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://*************************** see you.",
		},
		{
			name:     "Ссылка вначале",
			text:     "http://hehefouls.netHAHAHA Here's my spammy page: see you.",
			expected: "http://******************* Here's my spammy page: see you.",
		},
		{
			name:     "No links!",
			text:     "Hello world! My program is the best!I'm Golang developer!",
			expected: "Hello world! My program is the best!I'm Golang developer!",
		},
		{
			name:     "Fake http",
			text:     "http:/Hello world! My program is the http:best!I'm Golang developer!",
			expected: "http:/Hello world! My program is the http:best!I'm Golang developer!",
		},
		{
			name:     "Empty text",
			text:     "",
			expected: "",
		},
	}
	for _, testCase := range testTable {
		result := spammyMasker([]byte(testCase.text))
		assert.Equal(t, testCase.expected, string(result), testCase.name)
	}
}

func TestService_Run(t *testing.T) {
	mockProd := mocks.NewProducer(t)
	mockPres := mocks.NewPresenter(t)

	mockProd.On("Produce").Return([]byte(mock.Anything), nil)
	mockPres.On("Present", []byte(mock.Anything)).Return(nil)
	srv := NewService(mockProd, mockPres)

	assert.NotNil(t, srv)

	srv.Run()

	mockProd.AssertExpectations(t)
	mockPres.AssertExpectations(t)
}
