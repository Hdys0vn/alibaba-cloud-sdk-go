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

// GetAuthorityTemplate invokes the dms_enterprise.GetAuthorityTemplate API synchronously
func (client *Client) GetAuthorityTemplate(request *GetAuthorityTemplateRequest) (response *GetAuthorityTemplateResponse, err error) {
	response = CreateGetAuthorityTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuthorityTemplateWithChan invokes the dms_enterprise.GetAuthorityTemplate API asynchronously
func (client *Client) GetAuthorityTemplateWithChan(request *GetAuthorityTemplateRequest) (<-chan *GetAuthorityTemplateResponse, <-chan error) {
	responseChan := make(chan *GetAuthorityTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuthorityTemplate(request)
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

// GetAuthorityTemplateWithCallback invokes the dms_enterprise.GetAuthorityTemplate API asynchronously
func (client *Client) GetAuthorityTemplateWithCallback(request *GetAuthorityTemplateRequest, callback func(response *GetAuthorityTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuthorityTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetAuthorityTemplate(request)
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

// GetAuthorityTemplateRequest is the request struct for api GetAuthorityTemplate
type GetAuthorityTemplateRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	TemplateId requests.Integer `position:"Query" name:"TemplateId"`
}

// GetAuthorityTemplateResponse is the response struct for api GetAuthorityTemplate
type GetAuthorityTemplateResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	ErrorCode             string                `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage          string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Success               bool                  `json:"Success" xml:"Success"`
	Tid                   int64                 `json:"Tid" xml:"Tid"`
	AuthorityTemplateView AuthorityTemplateView `json:"AuthorityTemplateView" xml:"AuthorityTemplateView"`
}

// CreateGetAuthorityTemplateRequest creates a request to invoke GetAuthorityTemplate API
func CreateGetAuthorityTemplateRequest() (request *GetAuthorityTemplateRequest) {
	request = &GetAuthorityTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetAuthorityTemplate", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAuthorityTemplateResponse creates a response to parse from GetAuthorityTemplate response
func CreateGetAuthorityTemplateResponse() (response *GetAuthorityTemplateResponse) {
	response = &GetAuthorityTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
