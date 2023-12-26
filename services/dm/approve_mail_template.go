package dm

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

// ApproveMailTemplate invokes the dm.ApproveMailTemplate API synchronously
// api document: https://help.aliyun.com/api/dm/approvemailtemplate.html
func (client *Client) ApproveMailTemplate(request *ApproveMailTemplateRequest) (response *ApproveMailTemplateResponse, err error) {
	response = CreateApproveMailTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ApproveMailTemplateWithChan invokes the dm.ApproveMailTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/approvemailtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveMailTemplateWithChan(request *ApproveMailTemplateRequest) (<-chan *ApproveMailTemplateResponse, <-chan error) {
	responseChan := make(chan *ApproveMailTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApproveMailTemplate(request)
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

// ApproveMailTemplateWithCallback invokes the dm.ApproveMailTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/approvemailtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveMailTemplateWithCallback(request *ApproveMailTemplateRequest, callback func(response *ApproveMailTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApproveMailTemplateResponse
		var err error
		defer close(result)
		response, err = client.ApproveMailTemplate(request)
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

// ApproveMailTemplateRequest is the request struct for api ApproveMailTemplate
type ApproveMailTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// ApproveMailTemplateResponse is the response struct for api ApproveMailTemplate
type ApproveMailTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateApproveMailTemplateRequest creates a request to invoke ApproveMailTemplate API
func CreateApproveMailTemplateRequest() (request *ApproveMailTemplateRequest) {
	request = &ApproveMailTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "ApproveMailTemplate", "", "")
	return
}

// CreateApproveMailTemplateResponse creates a response to parse from ApproveMailTemplate response
func CreateApproveMailTemplateResponse() (response *ApproveMailTemplateResponse) {
	response = &ApproveMailTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
