package ecd

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

// RunCommand invokes the ecd.RunCommand API synchronously
func (client *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
	response = CreateRunCommandResponse()
	err = client.DoAction(request, response)
	return
}

// RunCommandWithChan invokes the ecd.RunCommand API asynchronously
func (client *Client) RunCommandWithChan(request *RunCommandRequest) (<-chan *RunCommandResponse, <-chan error) {
	responseChan := make(chan *RunCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunCommand(request)
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

// RunCommandWithCallback invokes the ecd.RunCommand API asynchronously
func (client *Client) RunCommandWithCallback(request *RunCommandRequest, callback func(response *RunCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunCommandResponse
		var err error
		defer close(result)
		response, err = client.RunCommand(request)
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

// RunCommandRequest is the request struct for api RunCommand
type RunCommandRequest struct {
	*requests.RpcRequest
	Type            string           `position:"Query" name:"Type"`
	CommandContent  string           `position:"Query" name:"CommandContent"`
	Timeout         requests.Integer `position:"Query" name:"Timeout"`
	ContentEncoding string           `position:"Query" name:"ContentEncoding"`
	EndUserId       string           `position:"Query" name:"EndUserId"`
	DesktopId       *[]string        `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// RunCommandResponse is the response struct for api RunCommand
type RunCommandResponse struct {
	*responses.BaseResponse
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRunCommandRequest creates a request to invoke RunCommand API
func CreateRunCommandRequest() (request *RunCommandRequest) {
	request = &RunCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "RunCommand", "", "")
	request.Method = requests.POST
	return
}

// CreateRunCommandResponse creates a response to parse from RunCommand response
func CreateRunCommandResponse() (response *RunCommandResponse) {
	response = &RunCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
