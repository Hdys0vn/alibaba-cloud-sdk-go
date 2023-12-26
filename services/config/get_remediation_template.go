package config

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

// GetRemediationTemplate invokes the config.GetRemediationTemplate API synchronously
func (client *Client) GetRemediationTemplate(request *GetRemediationTemplateRequest) (response *GetRemediationTemplateResponse, err error) {
	response = CreateGetRemediationTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetRemediationTemplateWithChan invokes the config.GetRemediationTemplate API asynchronously
func (client *Client) GetRemediationTemplateWithChan(request *GetRemediationTemplateRequest) (<-chan *GetRemediationTemplateResponse, <-chan error) {
	responseChan := make(chan *GetRemediationTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRemediationTemplate(request)
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

// GetRemediationTemplateWithCallback invokes the config.GetRemediationTemplate API asynchronously
func (client *Client) GetRemediationTemplateWithCallback(request *GetRemediationTemplateRequest, callback func(response *GetRemediationTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRemediationTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetRemediationTemplate(request)
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

// GetRemediationTemplateRequest is the request struct for api GetRemediationTemplate
type GetRemediationTemplateRequest struct {
	*requests.RpcRequest
	TemplateIdentifier string `position:"Query" name:"TemplateIdentifier"`
}

// GetRemediationTemplateResponse is the response struct for api GetRemediationTemplate
type GetRemediationTemplateResponse struct {
	*responses.BaseResponse
	RequestId            string                `json:"RequestId" xml:"RequestId"`
	RemediationTemplates []RemediationTemplate `json:"RemediationTemplates" xml:"RemediationTemplates"`
}

// CreateGetRemediationTemplateRequest creates a request to invoke GetRemediationTemplate API
func CreateGetRemediationTemplateRequest() (request *GetRemediationTemplateRequest) {
	request = &GetRemediationTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetRemediationTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRemediationTemplateResponse creates a response to parse from GetRemediationTemplate response
func CreateGetRemediationTemplateResponse() (response *GetRemediationTemplateResponse) {
	response = &GetRemediationTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}