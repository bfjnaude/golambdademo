package main_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	main "github.com/bfjnaude/golambdademo"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds ErrNameNotProvided when no name is given in the HTTP body
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     main.ErrNameNotProvided,
		},
		{
			request: events.APIGatewayProxyRequest{Body: "Bob"},
			expect:  "Hello Bob",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}
