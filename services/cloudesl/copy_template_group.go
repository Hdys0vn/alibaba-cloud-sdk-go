package cloudesl

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

// CopyTemplateGroup invokes the cloudesl.CopyTemplateGroup API synchronously
func (client *Client) CopyTemplateGroup(request *CopyTemplateGroupRequest) (response *CopyTemplateGroupResponse, err error) {
	response = CreateCopyTemplateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CopyTemplateGroupWithChan invokes the cloudesl.CopyTemplateGroup API asynchronously
func (client *Client) CopyTemplateGroupWithChan(request *CopyTemplateGroupRequest) (<-chan *CopyTemplateGroupResponse, <-chan error) {
	responseChan := make(chan *CopyTemplateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyTemplateGroup(request)
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

// CopyTemplateGroupWithCallback invokes the cloudesl.CopyTemplateGroup API asynchronously
func (client *Client) CopyTemplateGroupWithCallback(request *CopyTemplateGroupRequest, callback func(response *CopyTemplateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyTemplateGroupResponse
		var err error
		defer close(result)
		response, err = client.CopyTemplateGroup(request)
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

// CopyTemplateGroupRequest is the request struct for api CopyTemplateGroup
type CopyTemplateGroupRequest struct {
	*requests.RpcRequest
	GroupId         string `position:"Body" name:"GroupId"`
	TemplateVersion string `position:"Body" name:"TemplateVersion"`
	EslModelId      string `position:"Body" name:"EslModelId"`
}

// CopyTemplateGroupResponse is the response struct for api CopyTemplateGroup
type CopyTemplateGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code           string `json:"Code" xml:"Code"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateCopyTemplateGroupRequest creates a request to invoke CopyTemplateGroup API
func CreateCopyTemplateGroupRequest() (request *CopyTemplateGroupRequest) {
	request = &CopyTemplateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "CopyTemplateGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCopyTemplateGroupResponse creates a response to parse from CopyTemplateGroup response
func CreateCopyTemplateGroupResponse() (response *CopyTemplateGroupResponse) {
	response = &CopyTemplateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}