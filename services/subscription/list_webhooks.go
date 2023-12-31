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

// ListWebhooks invokes the subscription.ListWebhooks API synchronously
// api document: https://help.aliyun.com/api/subscription/listwebhooks.html
func (client *Client) ListWebhooks(request *ListWebhooksRequest) (response *ListWebhooksResponse, err error) {
	response = CreateListWebhooksResponse()
	err = client.DoAction(request, response)
	return
}

// ListWebhooksWithChan invokes the subscription.ListWebhooks API asynchronously
// api document: https://help.aliyun.com/api/subscription/listwebhooks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListWebhooksWithChan(request *ListWebhooksRequest) (<-chan *ListWebhooksResponse, <-chan error) {
	responseChan := make(chan *ListWebhooksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWebhooks(request)
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

// ListWebhooksWithCallback invokes the subscription.ListWebhooks API asynchronously
// api document: https://help.aliyun.com/api/subscription/listwebhooks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListWebhooksWithCallback(request *ListWebhooksRequest, callback func(response *ListWebhooksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWebhooksResponse
		var err error
		defer close(result)
		response, err = client.ListWebhooks(request)
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

// ListWebhooksRequest is the request struct for api ListWebhooks
type ListWebhooksRequest struct {
	*requests.RpcRequest
	WebhookId  requests.Integer `position:"Query" name:"WebhookId"`
	Locale     string           `position:"Query" name:"Locale"`
	Filter     string           `position:"Query" name:"Filter"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// ListWebhooksResponse is the response struct for api ListWebhooks
type ListWebhooksResponse struct {
	*responses.BaseResponse
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	Message    string    `json:"Message" xml:"Message"`
	NextToken  int       `json:"NextToken" xml:"NextToken"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	Code       string    `json:"Code" xml:"Code"`
	Success    bool      `json:"Success" xml:"Success"`
	Webhooks   []Webhook `json:"Webhooks" xml:"Webhooks"`
}

// CreateListWebhooksRequest creates a request to invoke ListWebhooks API
func CreateListWebhooksRequest() (request *ListWebhooksRequest) {
	request = &ListWebhooksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "ListWebhooks", "", "")
	return
}

// CreateListWebhooksResponse creates a response to parse from ListWebhooks response
func CreateListWebhooksResponse() (response *ListWebhooksResponse) {
	response = &ListWebhooksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
