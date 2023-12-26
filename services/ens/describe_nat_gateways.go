package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeNatGateways invokes the ens.DescribeNatGateways API synchronously
func (client *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	response = CreateDescribeNatGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNatGatewaysWithChan invokes the ens.DescribeNatGateways API asynchronously
func (client *Client) DescribeNatGatewaysWithChan(request *DescribeNatGatewaysRequest) (<-chan *DescribeNatGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeNatGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNatGateways(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeNatGatewaysWithCallback invokes the ens.DescribeNatGateways API asynchronously
func (client *Client) DescribeNatGatewaysWithCallback(request *DescribeNatGatewaysRequest, callback func(response *DescribeNatGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNatGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeNatGateways(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeNatGatewaysRequest is the request struct for api DescribeNatGateways
type DescribeNatGatewaysRequest struct {
	*requests.RpcRequest
	EnsRegionId  string           `position:"Query" name:"EnsRegionId"`
	VSwitchId    string           `position:"Query" name:"VSwitchId"`
	Name         string           `position:"Query" name:"Name"`
	NetworkId    string           `position:"Query" name:"NetworkId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	NatGatewayId string           `position:"Query" name:"NatGatewayId"`
}

// DescribeNatGatewaysResponse is the response struct for api DescribeNatGateways
type DescribeNatGatewaysResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	PageNumber  int          `json:"PageNumber" xml:"PageNumber"`
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	PageSize    int          `json:"PageSize" xml:"PageSize"`
	NatGateways []NatGateway `json:"NatGateways" xml:"NatGateways"`
}

// CreateDescribeNatGatewaysRequest creates a request to invoke DescribeNatGateways API
func CreateDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeNatGateways", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeNatGatewaysResponse creates a response to parse from DescribeNatGateways response
func CreateDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
