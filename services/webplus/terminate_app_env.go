package webplus

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

// TerminateAppEnv invokes the webplus.TerminateAppEnv API synchronously
// api document: https://help.aliyun.com/api/webplus/terminateappenv.html
func (client *Client) TerminateAppEnv(request *TerminateAppEnvRequest) (response *TerminateAppEnvResponse, err error) {
	response = CreateTerminateAppEnvResponse()
	err = client.DoAction(request, response)
	return
}

// TerminateAppEnvWithChan invokes the webplus.TerminateAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/terminateappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateAppEnvWithChan(request *TerminateAppEnvRequest) (<-chan *TerminateAppEnvResponse, <-chan error) {
	responseChan := make(chan *TerminateAppEnvResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TerminateAppEnv(request)
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

// TerminateAppEnvWithCallback invokes the webplus.TerminateAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/terminateappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateAppEnvWithCallback(request *TerminateAppEnvRequest, callback func(response *TerminateAppEnvResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TerminateAppEnvResponse
		var err error
		defer close(result)
		response, err = client.TerminateAppEnv(request)
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

// TerminateAppEnvRequest is the request struct for api TerminateAppEnv
type TerminateAppEnvRequest struct {
	*requests.RoaRequest
	DryRun string `position:"Body" name:"DryRun"`
	EnvId  string `position:"Body" name:"EnvId"`
}

// TerminateAppEnvResponse is the response struct for api TerminateAppEnv
type TerminateAppEnvResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	EnvChangeDetail EnvChangeDetail `json:"EnvChangeDetail" xml:"EnvChangeDetail"`
}

// CreateTerminateAppEnvRequest creates a request to invoke TerminateAppEnv API
func CreateTerminateAppEnvRequest() (request *TerminateAppEnvRequest) {
	request = &TerminateAppEnvRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "TerminateAppEnv", "/pop/v1/wam/appEnv/terminate", "", "")
	request.Method = requests.POST
	return
}

// CreateTerminateAppEnvResponse creates a response to parse from TerminateAppEnv response
func CreateTerminateAppEnvResponse() (response *TerminateAppEnvResponse) {
	response = &TerminateAppEnvResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
