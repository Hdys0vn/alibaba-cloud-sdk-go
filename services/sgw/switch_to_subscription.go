package sgw

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

// SwitchToSubscription invokes the sgw.SwitchToSubscription API synchronously
func (client *Client) SwitchToSubscription(request *SwitchToSubscriptionRequest) (response *SwitchToSubscriptionResponse, err error) {
	response = CreateSwitchToSubscriptionResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchToSubscriptionWithChan invokes the sgw.SwitchToSubscription API asynchronously
func (client *Client) SwitchToSubscriptionWithChan(request *SwitchToSubscriptionRequest) (<-chan *SwitchToSubscriptionResponse, <-chan error) {
	responseChan := make(chan *SwitchToSubscriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchToSubscription(request)
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

// SwitchToSubscriptionWithCallback invokes the sgw.SwitchToSubscription API asynchronously
func (client *Client) SwitchToSubscriptionWithCallback(request *SwitchToSubscriptionRequest, callback func(response *SwitchToSubscriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchToSubscriptionResponse
		var err error
		defer close(result)
		response, err = client.SwitchToSubscription(request)
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

// SwitchToSubscriptionRequest is the request struct for api SwitchToSubscription
type SwitchToSubscriptionRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// SwitchToSubscriptionResponse is the response struct for api SwitchToSubscription
type SwitchToSubscriptionResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	Success         bool   `json:"Success" xml:"Success"`
	Code            string `json:"Code" xml:"Code"`
	Message         string `json:"Message" xml:"Message"`
	SubscriptionURL string `json:"SubscriptionURL" xml:"SubscriptionURL"`
}

// CreateSwitchToSubscriptionRequest creates a request to invoke SwitchToSubscription API
func CreateSwitchToSubscriptionRequest() (request *SwitchToSubscriptionRequest) {
	request = &SwitchToSubscriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "SwitchToSubscription", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchToSubscriptionResponse creates a response to parse from SwitchToSubscription response
func CreateSwitchToSubscriptionResponse() (response *SwitchToSubscriptionResponse) {
	response = &SwitchToSubscriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
