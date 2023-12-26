package retailcloud

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

// DeleteAppEnvironment invokes the retailcloud.DeleteAppEnvironment API synchronously
func (client *Client) DeleteAppEnvironment(request *DeleteAppEnvironmentRequest) (response *DeleteAppEnvironmentResponse, err error) {
	response = CreateDeleteAppEnvironmentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAppEnvironmentWithChan invokes the retailcloud.DeleteAppEnvironment API asynchronously
func (client *Client) DeleteAppEnvironmentWithChan(request *DeleteAppEnvironmentRequest) (<-chan *DeleteAppEnvironmentResponse, <-chan error) {
	responseChan := make(chan *DeleteAppEnvironmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAppEnvironment(request)
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

// DeleteAppEnvironmentWithCallback invokes the retailcloud.DeleteAppEnvironment API asynchronously
func (client *Client) DeleteAppEnvironmentWithCallback(request *DeleteAppEnvironmentRequest, callback func(response *DeleteAppEnvironmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAppEnvironmentResponse
		var err error
		defer close(result)
		response, err = client.DeleteAppEnvironment(request)
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

// DeleteAppEnvironmentRequest is the request struct for api DeleteAppEnvironment
type DeleteAppEnvironmentRequest struct {
	*requests.RpcRequest
	AppId requests.Integer `position:"Query" name:"AppId"`
	Force requests.Boolean `position:"Query" name:"Force"`
	EnvId requests.Integer `position:"Query" name:"EnvId"`
}

// DeleteAppEnvironmentResponse is the response struct for api DeleteAppEnvironment
type DeleteAppEnvironmentResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeleteAppEnvironmentRequest creates a request to invoke DeleteAppEnvironment API
func CreateDeleteAppEnvironmentRequest() (request *DeleteAppEnvironmentRequest) {
	request = &DeleteAppEnvironmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DeleteAppEnvironment", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAppEnvironmentResponse creates a response to parse from DeleteAppEnvironment response
func CreateDeleteAppEnvironmentResponse() (response *DeleteAppEnvironmentResponse) {
	response = &DeleteAppEnvironmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}