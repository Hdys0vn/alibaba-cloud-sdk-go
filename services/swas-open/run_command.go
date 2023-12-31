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

// RunCommand invokes the swas_open.RunCommand API synchronously
func (client *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
	response = CreateRunCommandResponse()
	err = client.DoAction(request, response)
	return
}

// RunCommandWithChan invokes the swas_open.RunCommand API asynchronously
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

// RunCommandWithCallback invokes the swas_open.RunCommand API asynchronously
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
	WorkingDir          string            `position:"Query" name:"WorkingDir"`
	Type                string            `position:"Query" name:"Type"`
	CommandContent      string            `position:"Query" name:"CommandContent"`
	Timeout             requests.Integer  `position:"Query" name:"Timeout"`
	WindowsPasswordName string            `position:"Query" name:"WindowsPasswordName"`
	InstanceId          string            `position:"Query" name:"InstanceId"`
	WorkingUser         string            `position:"Query" name:"WorkingUser"`
	Name                string            `position:"Query" name:"Name"`
	Parameters          map[string]string `position:"Query" name:"Parameters"  type:"Map"`
	EnableParameter     requests.Boolean  `position:"Query" name:"EnableParameter"`
}

// RunCommandResponse is the response struct for api RunCommand
type RunCommandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
}

// CreateRunCommandRequest creates a request to invoke RunCommand API
func CreateRunCommandRequest() (request *RunCommandRequest) {
	request = &RunCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "RunCommand", "SWAS-OPEN", "openAPI")
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
