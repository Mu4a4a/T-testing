package main

import "testing"

func TestSpam(t *testing.T) {
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
	}
	for _, testCase := range testTable {
		result := spammyMasker(testCase.text)
		if result != testCase.expected {
			t.Errorf("Некорректный результат. Ожидалось: %v, Получили: %v", testCase.expected, result)
		}
	}
}
