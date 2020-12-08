//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
)

type DeviceProfileClient struct {
	baseUrl string
}

// NewDeviceProfileClient creates an instance of DeviceProfileClient
func NewDeviceProfileClient(baseUrl string) interfaces.DeviceProfileClient {
	return &DeviceProfileClient{
		baseUrl: baseUrl,
	}
}

func (client *DeviceProfileClient) Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	var responses []common.BaseWithIdResponse
	err := utils.PostRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

func (client *DeviceProfileClient) Put(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseResponse, errors.EdgeX) {
	var responses []common.BaseResponse
	err := utils.PutRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}
