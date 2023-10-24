package fga

import (
	"context"
	"fmt"

	openfga "github.com/openfga/go-sdk"
)

// Tuple for openFGA check request
type Tuple struct {
	// Subject contains the type and value for the actor, e.g `user:sarah` or `organization:datum#member`
	Subject string
	// Relation contains the relationship of the subject to the object
	Relation string
	// Object is the type and value being accessed e.g. `domain:datum.net`
	Object string
}

// CheckTuple checks the openFGA store for provided relationship tuple
func (c *Client) CheckTuple(ctx context.Context, check Tuple) (bool, error) {
	body := openfga.CheckRequest{
		TupleKey: openfga.TupleKey{
			User:     openfga.PtrString(check.Subject),
			Relation: openfga.PtrString(check.Relation),
			Object:   openfga.PtrString(check.Object),
		},
	}

	data, response, err := c.o.OpenFgaApi.Check(ctx).Body(body).Execute()
	if err != nil {
		fmt.Printf("GetCheck error: [%s][%v]", err.Error(), response)
		c.logger.Infof("GetCheck error: [%s][%v]", err.Error(), response)

		return false, err
	}

	return *data.Allowed, nil
}
