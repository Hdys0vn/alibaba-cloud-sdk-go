package cms

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

// DeleteGroupMonitoringAgentProcess invokes the cms.DeleteGroupMonitoringAgentProcess API synchronously
func (client *Client) DeleteGroupMonitoringAgentProcess(request *DeleteGroupMonitoringAgentProcessRequest) (response *DeleteGroupMonitoringAgentProcessResponse, err error) {
	response = CreateDeleteGroupMonitoringAgentProcessResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGroupMonitoringAgentProcessWithChan invokes the cms.DeleteGroupMonitoringAgentProcess API asynchronously
func (client *Client) DeleteGroupMonitoringAgentProcessWithChan(request *DeleteGroupMonitoringAgentProcessRequest) (<-chan *DeleteGroupMonitoringAgentProcessResponse, <-chan error) {
	responseChan := make(chan *DeleteGroupMonitoringAgentProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGroupMonitoringAgentProcess(request)
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

// DeleteGroupMonitoringAgentProcessWithCallback invokes the cms.DeleteGroupMonitoringAgentProcess API asynchronously
func (client *Client) DeleteGroupMonitoringAgentProcessWithCallback(request *DeleteGroupMonitoringAgentProcessRequest, callback func(response *DeleteGroupMonitoringAgentProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGroupMonitoringAgentProcessResponse
		var err error
		defer close(result)
		response, err = client.DeleteGroupMonitoringAgentProcess(request)
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

// DeleteGroupMonitoringAgentProcessRequest is the request struct for api DeleteGroupMonitoringAgentProcess
type DeleteGroupMonitoringAgentProcessRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	Id      string `position:"Query" name:"Id"`
}

// DeleteGroupMonitoringAgentProcessResponse is the response struct for api DeleteGroupMonitoringAgentProcess
type DeleteGroupMonitoringAgentProcessResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteGroupMonitoringAgentProcessRequest creates a request to invoke DeleteGroupMonitoringAgentProcess API
func CreateDeleteGroupMonitoringAgentProcessRequest() (request *DeleteGroupMonitoringAgentProcessRequest) {
	request = &DeleteGroupMonitoringAgentProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteGroupMonitoringAgentProcess", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGroupMonitoringAgentProcessResponse creates a response to parse from DeleteGroupMonitoringAgentProcess response
func CreateDeleteGroupMonitoringAgentProcessResponse() (response *DeleteGroupMonitoringAgentProcessResponse) {
	response = &DeleteGroupMonitoringAgentProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
