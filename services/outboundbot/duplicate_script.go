package outboundbot

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

// DuplicateScript invokes the outboundbot.DuplicateScript API synchronously
func (client *Client) DuplicateScript(request *DuplicateScriptRequest) (response *DuplicateScriptResponse, err error) {
	response = CreateDuplicateScriptResponse()
	err = client.DoAction(request, response)
	return
}

// DuplicateScriptWithChan invokes the outboundbot.DuplicateScript API asynchronously
func (client *Client) DuplicateScriptWithChan(request *DuplicateScriptRequest) (<-chan *DuplicateScriptResponse, <-chan error) {
	responseChan := make(chan *DuplicateScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DuplicateScript(request)
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

// DuplicateScriptWithCallback invokes the outboundbot.DuplicateScript API asynchronously
func (client *Client) DuplicateScriptWithCallback(request *DuplicateScriptRequest, callback func(response *DuplicateScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DuplicateScriptResponse
		var err error
		defer close(result)
		response, err = client.DuplicateScript(request)
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

// DuplicateScriptRequest is the request struct for api DuplicateScript
type DuplicateScriptRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Query" name:"InstanceId"`
	SourceScriptId string `position:"Query" name:"SourceScriptId"`
	Name           string `position:"Query" name:"Name"`
}

// DuplicateScriptResponse is the response struct for api DuplicateScript
type DuplicateScriptResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	ScriptId       string `json:"ScriptId" xml:"ScriptId"`
}

// CreateDuplicateScriptRequest creates a request to invoke DuplicateScript API
func CreateDuplicateScriptRequest() (request *DuplicateScriptRequest) {
	request = &DuplicateScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DuplicateScript", "", "")
	request.Method = requests.POST
	return
}

// CreateDuplicateScriptResponse creates a response to parse from DuplicateScript response
func CreateDuplicateScriptResponse() (response *DuplicateScriptResponse) {
	response = &DuplicateScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
