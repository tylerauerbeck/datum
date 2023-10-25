package fga

import (
	"context"

	"github.com/openfga/go-sdk/client"
)

// CheckTuple checks the openFGA store for provided relationship tuple
func (c *Client) CheckTuple(ctx context.Context, check client.ClientCheckRequest) (bool, error) {
	data, err := c.O.Check(ctx).Body(check).Execute()
	if err != nil {
		c.Logger.Infof("GetCheck error: [%s]", err.Error())

		return false, err
	}

	return *data.Allowed, nil
}
