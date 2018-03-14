//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package cni

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

type CNI interface {
	Info(ctx context.Context) *types.Subnet
	Create(ctx context.Context, network *types.Subnet) error
	Destroy(ctx context.Context, network *types.Subnet) error
	Replace(ctx context.Context, current *types.Subnet, proposal *types.Subnet) error
	Subnets(ctx context.Context) (map[string]*types.Subnet, error)
}
