package cloudwf

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

// DelPageConfigTemplate invokes the cloudwf.DelPageConfigTemplate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/delpageconfigtemplate.html
func (client *Client) DelPageConfigTemplate(request *DelPageConfigTemplateRequest) (response *DelPageConfigTemplateResponse, err error) {
	response = CreateDelPageConfigTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DelPageConfigTemplateWithChan invokes the cloudwf.DelPageConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delpageconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelPageConfigTemplateWithChan(request *DelPageConfigTemplateRequest) (<-chan *DelPageConfigTemplateResponse, <-chan error) {
	responseChan := make(chan *DelPageConfigTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelPageConfigTemplate(request)
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

// DelPageConfigTemplateWithCallback invokes the cloudwf.DelPageConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delpageconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelPageConfigTemplateWithCallback(request *DelPageConfigTemplateRequest, callback func(response *DelPageConfigTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelPageConfigTemplateResponse
		var err error
		defer close(result)
		response, err = client.DelPageConfigTemplate(request)
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

// DelPageConfigTemplateRequest is the request struct for api DelPageConfigTemplate
type DelPageConfigTemplateRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// DelPageConfigTemplateResponse is the response struct for api DelPageConfigTemplate
type DelPageConfigTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateDelPageConfigTemplateRequest creates a request to invoke DelPageConfigTemplate API
func CreateDelPageConfigTemplateRequest() (request *DelPageConfigTemplateRequest) {
	request = &DelPageConfigTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "DelPageConfigTemplate", "cloudwf", "openAPI")
	return
}

// CreateDelPageConfigTemplateResponse creates a response to parse from DelPageConfigTemplate response
func CreateDelPageConfigTemplateResponse() (response *DelPageConfigTemplateResponse) {
	response = &DelPageConfigTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
