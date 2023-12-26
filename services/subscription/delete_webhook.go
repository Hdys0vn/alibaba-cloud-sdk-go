package subscription

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

// DeleteWebhook invokes the subscription.DeleteWebhook API synchronously
// api document: https://help.aliyun.com/api/subscription/deletewebhook.html
func (client *Client) DeleteWebhook(request *DeleteWebhookRequest) (response *DeleteWebhookResponse, err error) {
	response = CreateDeleteWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWebhookWithChan invokes the subscription.DeleteWebhook API asynchronously
// api document: https://help.aliyun.com/api/subscription/deletewebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteWebhookWithChan(request *DeleteWebhookRequest) (<-chan *DeleteWebhookResponse, <-chan error) {
	responseChan := make(chan *DeleteWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWebhook(request)
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

// DeleteWebhookWithCallback invokes the subscription.DeleteWebhook API asynchronously
// api document: https://help.aliyun.com/api/subscription/deletewebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteWebhookWithCallback(request *DeleteWebhookRequest, callback func(response *DeleteWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWebhookResponse
		var err error
		defer close(result)
		response, err = client.DeleteWebhook(request)
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

// DeleteWebhookRequest is the request struct for api DeleteWebhook
type DeleteWebhookRequest struct {
	*requests.RpcRequest
	WebhookId requests.Integer `position:"Body" name:"WebhookId"`
	Locale    string           `position:"Query" name:"Locale"`
}

// DeleteWebhookResponse is the response struct for api DeleteWebhook
type DeleteWebhookResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteWebhookRequest creates a request to invoke DeleteWebhook API
func CreateDeleteWebhookRequest() (request *DeleteWebhookRequest) {
	request = &DeleteWebhookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "DeleteWebhook", "", "")
	return
}

// CreateDeleteWebhookResponse creates a response to parse from DeleteWebhook response
func CreateDeleteWebhookResponse() (response *DeleteWebhookResponse) {
	response = &DeleteWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}