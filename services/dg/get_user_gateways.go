package dg

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

// GetUserGateways invokes the dg.GetUserGateways API synchronously
func (client *Client) GetUserGateways(request *GetUserGatewaysRequest) (response *GetUserGatewaysResponse, err error) {
	response = CreateGetUserGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserGatewaysWithChan invokes the dg.GetUserGateways API asynchronously
func (client *Client) GetUserGatewaysWithChan(request *GetUserGatewaysRequest) (<-chan *GetUserGatewaysResponse, <-chan error) {
	responseChan := make(chan *GetUserGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserGateways(request)
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

// GetUserGatewaysWithCallback invokes the dg.GetUserGateways API asynchronously
func (client *Client) GetUserGatewaysWithCallback(request *GetUserGatewaysRequest, callback func(response *GetUserGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserGatewaysResponse
		var err error
		defer close(result)
		response, err = client.GetUserGateways(request)
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

// GetUserGatewaysRequest is the request struct for api GetUserGateways
type GetUserGatewaysRequest struct {
	*requests.RpcRequest
	SearchKey  string           `position:"Body" name:"SearchKey"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
}

// GetUserGatewaysResponse is the response struct for api GetUserGateways
type GetUserGatewaysResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
	Count     int    `json:"Count" xml:"Count"`
}

// CreateGetUserGatewaysRequest creates a request to invoke GetUserGateways API
func CreateGetUserGatewaysRequest() (request *GetUserGatewaysRequest) {
	request = &GetUserGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "GetUserGateways", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserGatewaysResponse creates a response to parse from GetUserGateways response
func CreateGetUserGatewaysResponse() (response *GetUserGatewaysResponse) {
	response = &GetUserGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}