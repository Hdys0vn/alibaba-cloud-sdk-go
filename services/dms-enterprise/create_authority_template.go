package dms_enterprise

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

// CreateAuthorityTemplate invokes the dms_enterprise.CreateAuthorityTemplate API synchronously
func (client *Client) CreateAuthorityTemplate(request *CreateAuthorityTemplateRequest) (response *CreateAuthorityTemplateResponse, err error) {
	response = CreateCreateAuthorityTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAuthorityTemplateWithChan invokes the dms_enterprise.CreateAuthorityTemplate API asynchronously
func (client *Client) CreateAuthorityTemplateWithChan(request *CreateAuthorityTemplateRequest) (<-chan *CreateAuthorityTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateAuthorityTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAuthorityTemplate(request)
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

// CreateAuthorityTemplateWithCallback invokes the dms_enterprise.CreateAuthorityTemplate API asynchronously
func (client *Client) CreateAuthorityTemplateWithCallback(request *CreateAuthorityTemplateRequest, callback func(response *CreateAuthorityTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAuthorityTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateAuthorityTemplate(request)
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

// CreateAuthorityTemplateRequest is the request struct for api CreateAuthorityTemplate
type CreateAuthorityTemplateRequest struct {
	*requests.RpcRequest
	Description string           `position:"Query" name:"Description"`
	Tid         requests.Integer `position:"Query" name:"Tid"`
	Name        string           `position:"Query" name:"Name"`
}

// CreateAuthorityTemplateResponse is the response struct for api CreateAuthorityTemplate
type CreateAuthorityTemplateResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	ErrorCode             string                `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage          string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Success               bool                  `json:"Success" xml:"Success"`
	Tid                   int64                 `json:"Tid" xml:"Tid"`
	AuthorityTemplateView AuthorityTemplateView `json:"AuthorityTemplateView" xml:"AuthorityTemplateView"`
}

// CreateCreateAuthorityTemplateRequest creates a request to invoke CreateAuthorityTemplate API
func CreateCreateAuthorityTemplateRequest() (request *CreateAuthorityTemplateRequest) {
	request = &CreateAuthorityTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateAuthorityTemplate", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAuthorityTemplateResponse creates a response to parse from CreateAuthorityTemplate response
func CreateCreateAuthorityTemplateResponse() (response *CreateAuthorityTemplateResponse) {
	response = &CreateAuthorityTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
