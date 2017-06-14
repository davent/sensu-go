package testing

import "github.com/sensu/sensu-go/types"

// ListChecks for use with mock lib
func (c *MockClient) ListChecks() ([]types.CheckConfig, error) {
	args := c.Called()
	return args.Get(0).([]types.CheckConfig), args.Error(1)
}

// CreateCheck for use with mock lib
func (c *MockClient) CreateCheck(check *types.CheckConfig) error {
	args := c.Called(check)
	return args.Error(0)
}

// DeleteCheck for use with mock lib
func (c *MockClient) DeleteCheck(check *types.CheckConfig) error {
	args := c.Called(check)
	return args.Error(0)
}