package arms

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

// CreateOrUpdateWebhookContact invokes the arms.CreateOrUpdateWebhookContact API synchronously
func (client *Client) CreateOrUpdateWebhookContact(request *CreateOrUpdateWebhookContactRequest) (response *CreateOrUpdateWebhookContactResponse, err error) {
	response = CreateCreateOrUpdateWebhookContactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateWebhookContactWithChan invokes the arms.CreateOrUpdateWebhookContact API asynchronously
func (client *Client) CreateOrUpdateWebhookContactWithChan(request *CreateOrUpdateWebhookContactRequest) (<-chan *CreateOrUpdateWebhookContactResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateWebhookContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateWebhookContact(request)
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

// CreateOrUpdateWebhookContactWithCallback invokes the arms.CreateOrUpdateWebhookContact API asynchronously
func (client *Client) CreateOrUpdateWebhookContactWithCallback(request *CreateOrUpdateWebhookContactRequest, callback func(response *CreateOrUpdateWebhookContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateWebhookContactResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateWebhookContact(request)
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

// CreateOrUpdateWebhookContactRequest is the request struct for api CreateOrUpdateWebhookContact
type CreateOrUpdateWebhookContactRequest struct {
	*requests.RpcRequest
	WebhookId   requests.Integer `position:"Body" name:"WebhookId"`
	Method      string           `position:"Body" name:"Method"`
	WebhookName string           `position:"Body" name:"WebhookName"`
	BizParams   string           `position:"Body" name:"BizParams"`
	Body        string           `position:"Body" name:"Body"`
	ProxyUserId string           `position:"Body" name:"ProxyUserId"`
	Url         string           `position:"Body" name:"Url"`
	BizHeaders  string           `position:"Body" name:"BizHeaders"`
	RecoverBody string           `position:"Body" name:"RecoverBody"`
}

// CreateOrUpdateWebhookContactResponse is the response struct for api CreateOrUpdateWebhookContact
type CreateOrUpdateWebhookContactResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	WebhookContact WebhookContact `json:"WebhookContact" xml:"WebhookContact"`
}

// CreateCreateOrUpdateWebhookContactRequest creates a request to invoke CreateOrUpdateWebhookContact API
func CreateCreateOrUpdateWebhookContactRequest() (request *CreateOrUpdateWebhookContactRequest) {
	request = &CreateOrUpdateWebhookContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateOrUpdateWebhookContact", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOrUpdateWebhookContactResponse creates a response to parse from CreateOrUpdateWebhookContact response
func CreateCreateOrUpdateWebhookContactResponse() (response *CreateOrUpdateWebhookContactResponse) {
	response = &CreateOrUpdateWebhookContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
