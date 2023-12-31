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

// ListSubscriptionItemGroupDetails invokes the subscription.ListSubscriptionItemGroupDetails API synchronously
// api document: https://help.aliyun.com/api/subscription/listsubscriptionitemgroupdetails.html
func (client *Client) ListSubscriptionItemGroupDetails(request *ListSubscriptionItemGroupDetailsRequest) (response *ListSubscriptionItemGroupDetailsResponse, err error) {
	response = CreateListSubscriptionItemGroupDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSubscriptionItemGroupDetailsWithChan invokes the subscription.ListSubscriptionItemGroupDetails API asynchronously
// api document: https://help.aliyun.com/api/subscription/listsubscriptionitemgroupdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSubscriptionItemGroupDetailsWithChan(request *ListSubscriptionItemGroupDetailsRequest) (<-chan *ListSubscriptionItemGroupDetailsResponse, <-chan error) {
	responseChan := make(chan *ListSubscriptionItemGroupDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSubscriptionItemGroupDetails(request)
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

// ListSubscriptionItemGroupDetailsWithCallback invokes the subscription.ListSubscriptionItemGroupDetails API asynchronously
// api document: https://help.aliyun.com/api/subscription/listsubscriptionitemgroupdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSubscriptionItemGroupDetailsWithCallback(request *ListSubscriptionItemGroupDetailsRequest, callback func(response *ListSubscriptionItemGroupDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSubscriptionItemGroupDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListSubscriptionItemGroupDetails(request)
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

// ListSubscriptionItemGroupDetailsRequest is the request struct for api ListSubscriptionItemGroupDetails
type ListSubscriptionItemGroupDetailsRequest struct {
	*requests.RpcRequest
	Locale string `position:"Query" name:"Locale"`
}

// ListSubscriptionItemGroupDetailsResponse is the response struct for api ListSubscriptionItemGroupDetails
type ListSubscriptionItemGroupDetailsResponse struct {
	*responses.BaseResponse
	Message                      string                        `json:"Message" xml:"Message"`
	RequestId                    string                        `json:"RequestId" xml:"RequestId"`
	Code                         string                        `json:"Code" xml:"Code"`
	Success                      bool                          `json:"Success" xml:"Success"`
	SubscriptionItemGroupDetails []SubscriptionItemGroupDetail `json:"SubscriptionItemGroupDetails" xml:"SubscriptionItemGroupDetails"`
}

// CreateListSubscriptionItemGroupDetailsRequest creates a request to invoke ListSubscriptionItemGroupDetails API
func CreateListSubscriptionItemGroupDetailsRequest() (request *ListSubscriptionItemGroupDetailsRequest) {
	request = &ListSubscriptionItemGroupDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "ListSubscriptionItemGroupDetails", "", "")
	return
}

// CreateListSubscriptionItemGroupDetailsResponse creates a response to parse from ListSubscriptionItemGroupDetails response
func CreateListSubscriptionItemGroupDetailsResponse() (response *ListSubscriptionItemGroupDetailsResponse) {
	response = &ListSubscriptionItemGroupDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
