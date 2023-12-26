package dypnsapi_intl

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

// SearchVerification invokes the dypnsapi_intl.SearchVerification API synchronously
func (client *Client) SearchVerification(request *SearchVerificationRequest) (response *SearchVerificationResponse, err error) {
	response = CreateSearchVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// SearchVerificationWithChan invokes the dypnsapi_intl.SearchVerification API asynchronously
func (client *Client) SearchVerificationWithChan(request *SearchVerificationRequest) (<-chan *SearchVerificationResponse, <-chan error) {
	responseChan := make(chan *SearchVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchVerification(request)
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

// SearchVerificationWithCallback invokes the dypnsapi_intl.SearchVerification API asynchronously
func (client *Client) SearchVerificationWithCallback(request *SearchVerificationRequest, callback func(response *SearchVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchVerificationResponse
		var err error
		defer close(result)
		response, err = client.SearchVerification(request)
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

// SearchVerificationRequest is the request struct for api SearchVerification
type SearchVerificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Code                 string           `position:"Query" name:"Code"`
	ServiceSid           string           `position:"Query" name:"ServiceSid"`
	SendDate             requests.Integer `position:"Query" name:"SendDate"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RouteName            string           `position:"Query" name:"RouteName"`
	VerificationId       string           `position:"Query" name:"VerificationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CurrentPage          requests.Integer `position:"Query" name:"CurrentPage"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	To                   string           `position:"Query" name:"To"`
}

// SearchVerificationResponse is the response struct for api SearchVerification
type SearchVerificationResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Message   string                 `json:"Message" xml:"Message"`
	Model     map[string]interface{} `json:"Model" xml:"Model"`
	Code      string                 `json:"Code" xml:"Code"`
	Success   bool                   `json:"Success" xml:"Success"`
}

// CreateSearchVerificationRequest creates a request to invoke SearchVerification API
func CreateSearchVerificationRequest() (request *SearchVerificationRequest) {
	request = &SearchVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi-intl", "2017-07-25", "SearchVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchVerificationResponse creates a response to parse from SearchVerification response
func CreateSearchVerificationResponse() (response *SearchVerificationResponse) {
	response = &SearchVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
