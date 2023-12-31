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

// CreateTemplate invokes the dm.CreateTemplate API synchronously
// api document: https://help.aliyun.com/api/dm/createtemplate.html
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	response = CreateCreateTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTemplateWithChan invokes the dm.CreateTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/createtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTemplateWithChan(request *CreateTemplateRequest) (<-chan *CreateTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTemplate(request)
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

// CreateTemplateWithCallback invokes the dm.CreateTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/createtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTemplateWithCallback(request *CreateTemplateRequest, callback func(response *CreateTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateTemplate(request)
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

// CreateTemplateRequest is the request struct for api CreateTemplate
type CreateTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateType         requests.Integer `position:"Query" name:"TemplateType"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	TemplateSubject      string           `position:"Query" name:"TemplateSubject"`
	TemplateNickName     string           `position:"Query" name:"TemplateNickName"`
	TemplateText         string           `position:"Query" name:"TemplateText"`
	SmsType              requests.Integer `position:"Query" name:"SmsType"`
	SmsContent           string           `position:"Query" name:"SmsContent"`
	Remark               string           `position:"Query" name:"Remark"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// CreateTemplateResponse is the response struct for api CreateTemplate
type CreateTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId int    `json:"TemplateId" xml:"TemplateId"`
}

// CreateCreateTemplateRequest creates a request to invoke CreateTemplate API
func CreateCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "CreateTemplate", "", "")
	return
}

// CreateCreateTemplateResponse creates a response to parse from CreateTemplate response
func CreateCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
