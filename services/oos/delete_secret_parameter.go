package oos

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

// DeleteSecretParameter invokes the oos.DeleteSecretParameter API synchronously
func (client *Client) DeleteSecretParameter(request *DeleteSecretParameterRequest) (response *DeleteSecretParameterResponse, err error) {
	response = CreateDeleteSecretParameterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSecretParameterWithChan invokes the oos.DeleteSecretParameter API asynchronously
func (client *Client) DeleteSecretParameterWithChan(request *DeleteSecretParameterRequest) (<-chan *DeleteSecretParameterResponse, <-chan error) {
	responseChan := make(chan *DeleteSecretParameterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSecretParameter(request)
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

// DeleteSecretParameterWithCallback invokes the oos.DeleteSecretParameter API asynchronously
func (client *Client) DeleteSecretParameterWithCallback(request *DeleteSecretParameterRequest, callback func(response *DeleteSecretParameterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSecretParameterResponse
		var err error
		defer close(result)
		response, err = client.DeleteSecretParameter(request)
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

// DeleteSecretParameterRequest is the request struct for api DeleteSecretParameter
type DeleteSecretParameterRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// DeleteSecretParameterResponse is the response struct for api DeleteSecretParameter
type DeleteSecretParameterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSecretParameterRequest creates a request to invoke DeleteSecretParameter API
func CreateDeleteSecretParameterRequest() (request *DeleteSecretParameterRequest) {
	request = &DeleteSecretParameterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "DeleteSecretParameter", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSecretParameterResponse creates a response to parse from DeleteSecretParameter response
func CreateDeleteSecretParameterResponse() (response *DeleteSecretParameterResponse) {
	response = &DeleteSecretParameterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
