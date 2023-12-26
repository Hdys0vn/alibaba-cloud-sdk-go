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

// GetAITemplate invokes the vod.GetAITemplate API synchronously
func (client *Client) GetAITemplate(request *GetAITemplateRequest) (response *GetAITemplateResponse, err error) {
	response = CreateGetAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetAITemplateWithChan invokes the vod.GetAITemplate API asynchronously
func (client *Client) GetAITemplateWithChan(request *GetAITemplateRequest) (<-chan *GetAITemplateResponse, <-chan error) {
	responseChan := make(chan *GetAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAITemplate(request)
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

// GetAITemplateWithCallback invokes the vod.GetAITemplate API asynchronously
func (client *Client) GetAITemplateWithCallback(request *GetAITemplateRequest, callback func(response *GetAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAITemplateResponse
		var err error
		defer close(result)
		response, err = client.GetAITemplate(request)
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

// GetAITemplateRequest is the request struct for api GetAITemplate
type GetAITemplateRequest struct {
	*requests.RpcRequest
	TemplateId string `position:"Query" name:"TemplateId"`
}

// GetAITemplateResponse is the response struct for api GetAITemplate
type GetAITemplateResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TemplateInfo TemplateInfo `json:"TemplateInfo" xml:"TemplateInfo"`
}

// CreateGetAITemplateRequest creates a request to invoke GetAITemplate API
func CreateGetAITemplateRequest() (request *GetAITemplateRequest) {
	request = &GetAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAITemplate", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAITemplateResponse creates a response to parse from GetAITemplate response
func CreateGetAITemplateResponse() (response *GetAITemplateResponse) {
	response = &GetAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
