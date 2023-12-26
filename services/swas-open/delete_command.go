package swas_open

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

// DeleteCommand invokes the swas_open.DeleteCommand API synchronously
func (client *Client) DeleteCommand(request *DeleteCommandRequest) (response *DeleteCommandResponse, err error) {
	response = CreateDeleteCommandResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCommandWithChan invokes the swas_open.DeleteCommand API asynchronously
func (client *Client) DeleteCommandWithChan(request *DeleteCommandRequest) (<-chan *DeleteCommandResponse, <-chan error) {
	responseChan := make(chan *DeleteCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCommand(request)
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

// DeleteCommandWithCallback invokes the swas_open.DeleteCommand API asynchronously
func (client *Client) DeleteCommandWithCallback(request *DeleteCommandRequest, callback func(response *DeleteCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCommandResponse
		var err error
		defer close(result)
		response, err = client.DeleteCommand(request)
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

// DeleteCommandRequest is the request struct for api DeleteCommand
type DeleteCommandRequest struct {
	*requests.RpcRequest
	CommandId string `position:"Query" name:"CommandId"`
}

// DeleteCommandResponse is the response struct for api DeleteCommand
type DeleteCommandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCommandRequest creates a request to invoke DeleteCommand API
func CreateDeleteCommandRequest() (request *DeleteCommandRequest) {
	request = &DeleteCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DeleteCommand", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCommandResponse creates a response to parse from DeleteCommand response
func CreateDeleteCommandResponse() (response *DeleteCommandResponse) {
	response = &DeleteCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
