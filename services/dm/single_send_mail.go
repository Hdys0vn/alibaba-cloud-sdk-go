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

// SingleSendMail invokes the dm.SingleSendMail API synchronously
// api document: https://help.aliyun.com/api/dm/singlesendmail.html
func (client *Client) SingleSendMail(request *SingleSendMailRequest) (response *SingleSendMailResponse, err error) {
	response = CreateSingleSendMailResponse()
	err = client.DoAction(request, response)
	return
}

// SingleSendMailWithChan invokes the dm.SingleSendMail API asynchronously
// api document: https://help.aliyun.com/api/dm/singlesendmail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SingleSendMailWithChan(request *SingleSendMailRequest) (<-chan *SingleSendMailResponse, <-chan error) {
	responseChan := make(chan *SingleSendMailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SingleSendMail(request)
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

// SingleSendMailWithCallback invokes the dm.SingleSendMail API asynchronously
// api document: https://help.aliyun.com/api/dm/singlesendmail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SingleSendMailWithCallback(request *SingleSendMailRequest, callback func(response *SingleSendMailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SingleSendMailResponse
		var err error
		defer close(result)
		response, err = client.SingleSendMail(request)
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

// SingleSendMailRequest is the request struct for api SingleSendMail
type SingleSendMailRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	AddressType          requests.Integer `position:"Query" name:"AddressType"`
	TagName              string           `position:"Query" name:"TagName"`
	ReplyToAddress       requests.Boolean `position:"Query" name:"ReplyToAddress"`
	ToAddress            string           `position:"Query" name:"ToAddress"`
	Subject              string           `position:"Query" name:"Subject"`
	HtmlBody             string           `position:"Query" name:"HtmlBody"`
	TextBody             string           `position:"Query" name:"TextBody"`
	FromAlias            string           `position:"Query" name:"FromAlias"`
	ReplyAddress         string           `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string           `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string           `position:"Query" name:"ClickTrace"`
}

// SingleSendMailResponse is the response struct for api SingleSendMail
type SingleSendMailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	EnvId     string `json:"EnvId" xml:"EnvId"`
}

// CreateSingleSendMailRequest creates a request to invoke SingleSendMail API
func CreateSingleSendMailRequest() (request *SingleSendMailRequest) {
	request = &SingleSendMailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "SingleSendMail", "", "")
	return
}

// CreateSingleSendMailResponse creates a response to parse from SingleSendMail response
func CreateSingleSendMailResponse() (response *SingleSendMailResponse) {
	response = &SingleSendMailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
