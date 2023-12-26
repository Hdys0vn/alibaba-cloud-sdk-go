package dytnsapi

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

// PhoneNumberStatusForPublic invokes the dytnsapi.PhoneNumberStatusForPublic API synchronously
func (client *Client) PhoneNumberStatusForPublic(request *PhoneNumberStatusForPublicRequest) (response *PhoneNumberStatusForPublicResponse, err error) {
	response = CreatePhoneNumberStatusForPublicResponse()
	err = client.DoAction(request, response)
	return
}

// PhoneNumberStatusForPublicWithChan invokes the dytnsapi.PhoneNumberStatusForPublic API asynchronously
func (client *Client) PhoneNumberStatusForPublicWithChan(request *PhoneNumberStatusForPublicRequest) (<-chan *PhoneNumberStatusForPublicResponse, <-chan error) {
	responseChan := make(chan *PhoneNumberStatusForPublicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PhoneNumberStatusForPublic(request)
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

// PhoneNumberStatusForPublicWithCallback invokes the dytnsapi.PhoneNumberStatusForPublic API asynchronously
func (client *Client) PhoneNumberStatusForPublicWithCallback(request *PhoneNumberStatusForPublicRequest, callback func(response *PhoneNumberStatusForPublicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PhoneNumberStatusForPublicResponse
		var err error
		defer close(result)
		response, err = client.PhoneNumberStatusForPublic(request)
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

// PhoneNumberStatusForPublicRequest is the request struct for api PhoneNumberStatusForPublic
type PhoneNumberStatusForPublicRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteName            string           `position:"Query" name:"RouteName"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AuthCode             string           `position:"Query" name:"AuthCode"`
	InputNumber          string           `position:"Query" name:"InputNumber"`
}

// PhoneNumberStatusForPublicResponse is the response struct for api PhoneNumberStatusForPublic
type PhoneNumberStatusForPublicResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePhoneNumberStatusForPublicRequest creates a request to invoke PhoneNumberStatusForPublic API
func CreatePhoneNumberStatusForPublicRequest() (request *PhoneNumberStatusForPublicRequest) {
	request = &PhoneNumberStatusForPublicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "PhoneNumberStatusForPublic", "", "")
	request.Method = requests.POST
	return
}

// CreatePhoneNumberStatusForPublicResponse creates a response to parse from PhoneNumberStatusForPublic response
func CreatePhoneNumberStatusForPublicResponse() (response *PhoneNumberStatusForPublicResponse) {
	response = &PhoneNumberStatusForPublicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}