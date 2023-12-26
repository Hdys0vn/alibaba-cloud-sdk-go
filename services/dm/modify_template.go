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

// ModifyTemplate invokes the dm.ModifyTemplate API synchronously
// api document: https://help.aliyun.com/api/dm/modifytemplate.html
func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (response *ModifyTemplateResponse, err error) {
	response = CreateModifyTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTemplateWithChan invokes the dm.ModifyTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/modifytemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTemplateWithChan(request *ModifyTemplateRequest) (<-chan *ModifyTemplateResponse, <-chan error) {
	responseChan := make(chan *ModifyTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTemplate(request)
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

// ModifyTemplateWithCallback invokes the dm.ModifyTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/modifytemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTemplateWithCallback(request *ModifyTemplateRequest, callback func(response *ModifyTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTemplateResponse
		var err error
		defer close(result)
		response, err = client.ModifyTemplate(request)
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

// ModifyTemplateRequest is the request struct for api ModifyTemplate
type ModifyTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	TemplateSubject      string           `position:"Query" name:"TemplateSubject"`
	TemplateNickName     string           `position:"Query" name:"TemplateNickName"`
	TemplateText         string           `position:"Query" name:"TemplateText"`
	SmsType              requests.Integer `position:"Query" name:"SmsType"`
	SmsContent           string           `position:"Query" name:"SmsContent"`
	Remark               string           `position:"Query" name:"Remark"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// ModifyTemplateResponse is the response struct for api ModifyTemplate
type ModifyTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyTemplateRequest creates a request to invoke ModifyTemplate API
func CreateModifyTemplateRequest() (request *ModifyTemplateRequest) {
	request = &ModifyTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "ModifyTemplate", "", "")
	return
}

// CreateModifyTemplateResponse creates a response to parse from ModifyTemplate response
func CreateModifyTemplateResponse() (response *ModifyTemplateResponse) {
	response = &ModifyTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
