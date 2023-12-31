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

// DeleteContact invokes the subscription.DeleteContact API synchronously
// api document: https://help.aliyun.com/api/subscription/deletecontact.html
func (client *Client) DeleteContact(request *DeleteContactRequest) (response *DeleteContactResponse, err error) {
	response = CreateDeleteContactResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteContactWithChan invokes the subscription.DeleteContact API asynchronously
// api document: https://help.aliyun.com/api/subscription/deletecontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteContactWithChan(request *DeleteContactRequest) (<-chan *DeleteContactResponse, <-chan error) {
	responseChan := make(chan *DeleteContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteContact(request)
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

// DeleteContactWithCallback invokes the subscription.DeleteContact API asynchronously
// api document: https://help.aliyun.com/api/subscription/deletecontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteContactWithCallback(request *DeleteContactRequest, callback func(response *DeleteContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteContactResponse
		var err error
		defer close(result)
		response, err = client.DeleteContact(request)
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

// DeleteContactRequest is the request struct for api DeleteContact
type DeleteContactRequest struct {
	*requests.RpcRequest
	ContactId requests.Integer `position:"Body" name:"ContactId"`
	Locale    string           `position:"Query" name:"Locale"`
}

// DeleteContactResponse is the response struct for api DeleteContact
type DeleteContactResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteContactRequest creates a request to invoke DeleteContact API
func CreateDeleteContactRequest() (request *DeleteContactRequest) {
	request = &DeleteContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "DeleteContact", "", "")
	return
}

// CreateDeleteContactResponse creates a response to parse from DeleteContact response
func CreateDeleteContactResponse() (response *DeleteContactResponse) {
	response = &DeleteContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
