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

// ReinitInstance invokes the ens.ReinitInstance API synchronously
func (client *Client) ReinitInstance(request *ReinitInstanceRequest) (response *ReinitInstanceResponse, err error) {
	response = CreateReinitInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReinitInstanceWithChan invokes the ens.ReinitInstance API asynchronously
func (client *Client) ReinitInstanceWithChan(request *ReinitInstanceRequest) (<-chan *ReinitInstanceResponse, <-chan error) {
	responseChan := make(chan *ReinitInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReinitInstance(request)
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

// ReinitInstanceWithCallback invokes the ens.ReinitInstance API asynchronously
func (client *Client) ReinitInstanceWithCallback(request *ReinitInstanceRequest, callback func(response *ReinitInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReinitInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReinitInstance(request)
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

// ReinitInstanceRequest is the request struct for api ReinitInstance
type ReinitInstanceRequest struct {
	*requests.RpcRequest
	ImageId    string `position:"Body" name:"ImageId"`
	Password   string `position:"Body" name:"Password"`
	InstanceId string `position:"Body" name:"InstanceId"`
}

// ReinitInstanceResponse is the response struct for api ReinitInstance
type ReinitInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReinitInstanceRequest creates a request to invoke ReinitInstance API
func CreateReinitInstanceRequest() (request *ReinitInstanceRequest) {
	request = &ReinitInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ReinitInstance", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReinitInstanceResponse creates a response to parse from ReinitInstance response
func CreateReinitInstanceResponse() (response *ReinitInstanceResponse) {
	response = &ReinitInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
