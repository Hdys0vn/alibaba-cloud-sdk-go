package vod

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

// GetTranscodeTemplateGroup invokes the vod.GetTranscodeTemplateGroup API synchronously
func (client *Client) GetTranscodeTemplateGroup(request *GetTranscodeTemplateGroupRequest) (response *GetTranscodeTemplateGroupResponse, err error) {
	response = CreateGetTranscodeTemplateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetTranscodeTemplateGroupWithChan invokes the vod.GetTranscodeTemplateGroup API asynchronously
func (client *Client) GetTranscodeTemplateGroupWithChan(request *GetTranscodeTemplateGroupRequest) (<-chan *GetTranscodeTemplateGroupResponse, <-chan error) {
	responseChan := make(chan *GetTranscodeTemplateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTranscodeTemplateGroup(request)
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

// GetTranscodeTemplateGroupWithCallback invokes the vod.GetTranscodeTemplateGroup API asynchronously
func (client *Client) GetTranscodeTemplateGroupWithCallback(request *GetTranscodeTemplateGroupRequest, callback func(response *GetTranscodeTemplateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTranscodeTemplateGroupResponse
		var err error
		defer close(result)
		response, err = client.GetTranscodeTemplateGroup(request)
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

// GetTranscodeTemplateGroupRequest is the request struct for api GetTranscodeTemplateGroup
type GetTranscodeTemplateGroupRequest struct {
	*requests.RpcRequest
	TranscodeTemplateGroupId string `position:"Query" name:"TranscodeTemplateGroupId"`
}

// GetTranscodeTemplateGroupResponse is the response struct for api GetTranscodeTemplateGroup
type GetTranscodeTemplateGroupResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	TranscodeTemplateGroup TranscodeTemplateGroup `json:"TranscodeTemplateGroup" xml:"TranscodeTemplateGroup"`
}

// CreateGetTranscodeTemplateGroupRequest creates a request to invoke GetTranscodeTemplateGroup API
func CreateGetTranscodeTemplateGroupRequest() (request *GetTranscodeTemplateGroupRequest) {
	request = &GetTranscodeTemplateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetTranscodeTemplateGroup", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTranscodeTemplateGroupResponse creates a response to parse from GetTranscodeTemplateGroup response
func CreateGetTranscodeTemplateGroupResponse() (response *GetTranscodeTemplateGroupResponse) {
	response = &GetTranscodeTemplateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
