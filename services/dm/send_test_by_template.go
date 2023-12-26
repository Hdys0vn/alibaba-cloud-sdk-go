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

// SendTestByTemplate invokes the dm.SendTestByTemplate API synchronously
// api document: https://help.aliyun.com/api/dm/sendtestbytemplate.html
func (client *Client) SendTestByTemplate(request *SendTestByTemplateRequest) (response *SendTestByTemplateResponse, err error) {
	response = CreateSendTestByTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// SendTestByTemplateWithChan invokes the dm.SendTestByTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/sendtestbytemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendTestByTemplateWithChan(request *SendTestByTemplateRequest) (<-chan *SendTestByTemplateResponse, <-chan error) {
	responseChan := make(chan *SendTestByTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendTestByTemplate(request)
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

// SendTestByTemplateWithCallback invokes the dm.SendTestByTemplate API asynchronously
// api document: https://help.aliyun.com/api/dm/sendtestbytemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendTestByTemplateWithCallback(request *SendTestByTemplateRequest, callback func(response *SendTestByTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendTestByTemplateResponse
		var err error
		defer close(result)
		response, err = client.SendTestByTemplate(request)
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

// SendTestByTemplateRequest is the request struct for api SendTestByTemplate
type SendTestByTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	UserName             string           `position:"Query" name:"UserName"`
	NickName             string           `position:"Query" name:"NickName"`
	Birthday             string           `position:"Query" name:"Birthday"`
	Gender               string           `position:"Query" name:"Gender"`
	Mobile               string           `position:"Query" name:"Mobile"`
	Email                string           `position:"Query" name:"Email"`
}

// SendTestByTemplateResponse is the response struct for api SendTestByTemplate
type SendTestByTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendTestByTemplateRequest creates a request to invoke SendTestByTemplate API
func CreateSendTestByTemplateRequest() (request *SendTestByTemplateRequest) {
	request = &SendTestByTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "SendTestByTemplate", "", "")
	return
}

// CreateSendTestByTemplateResponse creates a response to parse from SendTestByTemplate response
func CreateSendTestByTemplateResponse() (response *SendTestByTemplateResponse) {
	response = &SendTestByTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
