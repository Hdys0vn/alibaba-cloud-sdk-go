package ecs

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

// RunCommand invokes the ecs.RunCommand API synchronously
func (client *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
	response = CreateRunCommandResponse()
	err = client.DoAction(request, response)
	return
}

// RunCommandWithChan invokes the ecs.RunCommand API asynchronously
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

// RunCommandWithCallback invokes the ecs.RunCommand API asynchronously
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
	ResourceOwnerId      requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string                 `position:"Query" name:"WorkingDir"`
	Type                 string                 `position:"Query" name:"Type"`
	Frequency            string                 `position:"Query" name:"Frequency"`
	ResourceGroupId      string                 `position:"Query" name:"ResourceGroupId"`
	RepeatMode           string                 `position:"Query" name:"RepeatMode"`
	Tag                  *[]RunCommandTag       `position:"Query" name:"Tag"  type:"Repeated"`
	KeepCommand          requests.Boolean       `position:"Query" name:"KeepCommand"`
	Timed                requests.Boolean       `position:"Query" name:"Timed"`
	OwnerId              requests.Integer       `position:"Query" name:"OwnerId"`
	InstanceId           *[]string              `position:"Query" name:"InstanceId"  type:"Repeated"`
	Name                 string                 `position:"Query" name:"Name"`
	ContainerId          string                 `position:"Query" name:"ContainerId"`
	Parameters           map[string]interface{} `position:"Query" name:"Parameters"`
	ContainerName        string                 `position:"Query" name:"ContainerName"`
	ClientToken          string                 `position:"Query" name:"ClientToken"`
	Description          string                 `position:"Query" name:"Description"`
	CommandContent       string                 `position:"Query" name:"CommandContent"`
	Timeout              requests.Integer       `position:"Query" name:"Timeout"`
	ContentEncoding      string                 `position:"Query" name:"ContentEncoding"`
	WindowsPasswordName  string                 `position:"Query" name:"WindowsPasswordName"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                 `position:"Query" name:"OwnerAccount"`
	EnableParameter      requests.Boolean       `position:"Query" name:"EnableParameter"`
	Username             string                 `position:"Query" name:"Username"`
}

// RunCommandTag is a repeated param struct in RunCommandRequest
type RunCommandTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// RunCommandResponse is the response struct for api RunCommand
type RunCommandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CommandId string `json:"CommandId" xml:"CommandId"`
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
}

// CreateRunCommandRequest creates a request to invoke RunCommand API
func CreateRunCommandRequest() (request *RunCommandRequest) {
	request = &RunCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RunCommand", "ecs", "openAPI")
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
