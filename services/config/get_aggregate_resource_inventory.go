package config

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

// GetAggregateResourceInventory invokes the config.GetAggregateResourceInventory API synchronously
func (client *Client) GetAggregateResourceInventory(request *GetAggregateResourceInventoryRequest) (response *GetAggregateResourceInventoryResponse, err error) {
	response = CreateGetAggregateResourceInventoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateResourceInventoryWithChan invokes the config.GetAggregateResourceInventory API asynchronously
func (client *Client) GetAggregateResourceInventoryWithChan(request *GetAggregateResourceInventoryRequest) (<-chan *GetAggregateResourceInventoryResponse, <-chan error) {
	responseChan := make(chan *GetAggregateResourceInventoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateResourceInventory(request)
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

// GetAggregateResourceInventoryWithCallback invokes the config.GetAggregateResourceInventory API asynchronously
func (client *Client) GetAggregateResourceInventoryWithCallback(request *GetAggregateResourceInventoryRequest, callback func(response *GetAggregateResourceInventoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateResourceInventoryResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateResourceInventory(request)
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

// GetAggregateResourceInventoryRequest is the request struct for api GetAggregateResourceInventory
type GetAggregateResourceInventoryRequest struct {
	*requests.RpcRequest
	AggregatorId string `position:"Query" name:"AggregatorId"`
}

// GetAggregateResourceInventoryResponse is the response struct for api GetAggregateResourceInventory
type GetAggregateResourceInventoryResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ResourceInventory ResourceInventory `json:"ResourceInventory" xml:"ResourceInventory"`
}

// CreateGetAggregateResourceInventoryRequest creates a request to invoke GetAggregateResourceInventory API
func CreateGetAggregateResourceInventoryRequest() (request *GetAggregateResourceInventoryRequest) {
	request = &GetAggregateResourceInventoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateResourceInventory", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAggregateResourceInventoryResponse creates a response to parse from GetAggregateResourceInventory response
func CreateGetAggregateResourceInventoryResponse() (response *GetAggregateResourceInventoryResponse) {
	response = &GetAggregateResourceInventoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
