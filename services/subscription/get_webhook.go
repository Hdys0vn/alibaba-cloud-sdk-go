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

// GetWebhook invokes the subscription.GetWebhook API synchronously
// api document: https://help.aliyun.com/api/subscription/getwebhook.html
func (client *Client) GetWebhook(request *GetWebhookRequest) (response *GetWebhookResponse, err error) {
	response = CreateGetWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// GetWebhookWithChan invokes the subscription.GetWebhook API asynchronously
// api document: https://help.aliyun.com/api/subscription/getwebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWebhookWithChan(request *GetWebhookRequest) (<-chan *GetWebhookResponse, <-chan error) {
	responseChan := make(chan *GetWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWebhook(request)
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

// GetWebhookWithCallback invokes the subscription.GetWebhook API asynchronously
// api document: https://help.aliyun.com/api/subscription/getwebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWebhookWithCallback(request *GetWebhookRequest, callback func(response *GetWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWebhookResponse
		var err error
		defer close(result)
		response, err = client.GetWebhook(request)
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

// GetWebhookRequest is the request struct for api GetWebhook
type GetWebhookRequest struct {
	*requests.RpcRequest
	WebhookId requests.Integer `position:"Query" name:"WebhookId"`
	Locale    string           `position:"Query" name:"Locale"`
}

// GetWebhookResponse is the response struct for api GetWebhook
type GetWebhookResponse struct {
	*responses.BaseResponse
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Code      string  `json:"Code" xml:"Code"`
	Success   bool    `json:"Success" xml:"Success"`
	Webhook   Webhook `json:"Webhook" xml:"Webhook"`
}

// CreateGetWebhookRequest creates a request to invoke GetWebhook API
func CreateGetWebhookRequest() (request *GetWebhookRequest) {
	request = &GetWebhookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "GetWebhook", "", "")
	return
}

// CreateGetWebhookResponse creates a response to parse from GetWebhook response
func CreateGetWebhookResponse() (response *GetWebhookResponse) {
	response = &GetWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
