package dcdn

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

// DeleteRoutineConfEnvs invokes the dcdn.DeleteRoutineConfEnvs API synchronously
func (client *Client) DeleteRoutineConfEnvs(request *DeleteRoutineConfEnvsRequest) (response *DeleteRoutineConfEnvsResponse, err error) {
	response = CreateDeleteRoutineConfEnvsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRoutineConfEnvsWithChan invokes the dcdn.DeleteRoutineConfEnvs API asynchronously
func (client *Client) DeleteRoutineConfEnvsWithChan(request *DeleteRoutineConfEnvsRequest) (<-chan *DeleteRoutineConfEnvsResponse, <-chan error) {
	responseChan := make(chan *DeleteRoutineConfEnvsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRoutineConfEnvs(request)
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

// DeleteRoutineConfEnvsWithCallback invokes the dcdn.DeleteRoutineConfEnvs API asynchronously
func (client *Client) DeleteRoutineConfEnvsWithCallback(request *DeleteRoutineConfEnvsRequest, callback func(response *DeleteRoutineConfEnvsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRoutineConfEnvsResponse
		var err error
		defer close(result)
		response, err = client.DeleteRoutineConfEnvs(request)
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

// DeleteRoutineConfEnvsRequest is the request struct for api DeleteRoutineConfEnvs
type DeleteRoutineConfEnvsRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
	Envs string `position:"Body" name:"Envs"`
}

// DeleteRoutineConfEnvsResponse is the response struct for api DeleteRoutineConfEnvs
type DeleteRoutineConfEnvsResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRoutineConfEnvsRequest creates a request to invoke DeleteRoutineConfEnvs API
func CreateDeleteRoutineConfEnvsRequest() (request *DeleteRoutineConfEnvsRequest) {
	request = &DeleteRoutineConfEnvsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DeleteRoutineConfEnvs", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRoutineConfEnvsResponse creates a response to parse from DeleteRoutineConfEnvs response
func CreateDeleteRoutineConfEnvsResponse() (response *DeleteRoutineConfEnvsResponse) {
	response = &DeleteRoutineConfEnvsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
