package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthKeyAuthTestSuite struct {
	suite.Suite
	testable *KeyAuth
}

func (suite *AuthKeyAuthTestSuite) SetupTest() {
	suite.testable = &KeyAuth{"foo", true}
}

func (suite *AuthKeyAuthTestSuite) TestGetRequestMetadata() {
	ctx := context.Background()
	result, err := suite.testable.GetRequestMetadata(ctx, "http://localhost:8000/api")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "apikey foo", result["authorization"])
}

func (suite *AuthKeyAuthTestSuite) TestTransportSecurity() {
	assert.True(suite.T(), suite.testable.RequireTransportSecurity())
}

func TestAuthKeyAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthKeyAuthTestSuite))
}
