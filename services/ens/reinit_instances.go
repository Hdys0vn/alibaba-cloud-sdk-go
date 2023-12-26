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

// ReinitInstances invokes the ens.ReinitInstances API synchronously
func (client *Client) ReinitInstances(request *ReinitInstancesRequest) (response *ReinitInstancesResponse, err error) {
	response = CreateReinitInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ReinitInstancesWithChan invokes the ens.ReinitInstances API asynchronously
func (client *Client) ReinitInstancesWithChan(request *ReinitInstancesRequest) (<-chan *ReinitInstancesResponse, <-chan error) {
	responseChan := make(chan *ReinitInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReinitInstances(request)
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

// ReinitInstancesWithCallback invokes the ens.ReinitInstances API asynchronously
func (client *Client) ReinitInstancesWithCallback(request *ReinitInstancesRequest, callback func(response *ReinitInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReinitInstancesResponse
		var err error
		defer close(result)
		response, err = client.ReinitInstances(request)
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

// ReinitInstancesRequest is the request struct for api ReinitInstances
type ReinitInstancesRequest struct {
	*requests.RpcRequest
	ImageId     string    `position:"Query" name:"ImageId"`
	Password    string    `position:"Query" name:"Password"`
	InstanceIds *[]string `position:"Query" name:"InstanceIds"  type:"Repeated"`
}

// ReinitInstancesResponse is the response struct for api ReinitInstances
type ReinitInstancesResponse struct {
	*responses.BaseResponse
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	InstanceResponses []InstanceResponsesItem `json:"InstanceResponses" xml:"InstanceResponses"`
}

// CreateReinitInstancesRequest creates a request to invoke ReinitInstances API
func CreateReinitInstancesRequest() (request *ReinitInstancesRequest) {
	request = &ReinitInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ReinitInstances", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReinitInstancesResponse creates a response to parse from ReinitInstances response
func CreateReinitInstancesResponse() (response *ReinitInstancesResponse) {
	response = &ReinitInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
