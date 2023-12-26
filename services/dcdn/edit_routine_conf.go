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

// EditRoutineConf invokes the dcdn.EditRoutineConf API synchronously
func (client *Client) EditRoutineConf(request *EditRoutineConfRequest) (response *EditRoutineConfResponse, err error) {
	response = CreateEditRoutineConfResponse()
	err = client.DoAction(request, response)
	return
}

// EditRoutineConfWithChan invokes the dcdn.EditRoutineConf API asynchronously
func (client *Client) EditRoutineConfWithChan(request *EditRoutineConfRequest) (<-chan *EditRoutineConfResponse, <-chan error) {
	responseChan := make(chan *EditRoutineConfResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditRoutineConf(request)
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

// EditRoutineConfWithCallback invokes the dcdn.EditRoutineConf API asynchronously
func (client *Client) EditRoutineConfWithCallback(request *EditRoutineConfRequest, callback func(response *EditRoutineConfResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditRoutineConfResponse
		var err error
		defer close(result)
		response, err = client.EditRoutineConf(request)
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

// EditRoutineConfRequest is the request struct for api EditRoutineConf
type EditRoutineConfRequest struct {
	*requests.RpcRequest
	EnvConf     string `position:"Body" name:"EnvConf"`
	Name        string `position:"Body" name:"Name"`
	Description string `position:"Body" name:"Description"`
}

// EditRoutineConfResponse is the response struct for api EditRoutineConf
type EditRoutineConfResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateEditRoutineConfRequest creates a request to invoke EditRoutineConf API
func CreateEditRoutineConfRequest() (request *EditRoutineConfRequest) {
	request = &EditRoutineConfRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "EditRoutineConf", "", "")
	request.Method = requests.POST
	return
}

// CreateEditRoutineConfResponse creates a response to parse from EditRoutineConf response
func CreateEditRoutineConfResponse() (response *EditRoutineConfResponse) {
	response = &EditRoutineConfResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
