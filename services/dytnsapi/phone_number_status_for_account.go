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

// PhoneNumberStatusForAccount invokes the dytnsapi.PhoneNumberStatusForAccount API synchronously
func (client *Client) PhoneNumberStatusForAccount(request *PhoneNumberStatusForAccountRequest) (response *PhoneNumberStatusForAccountResponse, err error) {
	response = CreatePhoneNumberStatusForAccountResponse()
	err = client.DoAction(request, response)
	return
}

// PhoneNumberStatusForAccountWithChan invokes the dytnsapi.PhoneNumberStatusForAccount API asynchronously
func (client *Client) PhoneNumberStatusForAccountWithChan(request *PhoneNumberStatusForAccountRequest) (<-chan *PhoneNumberStatusForAccountResponse, <-chan error) {
	responseChan := make(chan *PhoneNumberStatusForAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PhoneNumberStatusForAccount(request)
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

// PhoneNumberStatusForAccountWithCallback invokes the dytnsapi.PhoneNumberStatusForAccount API asynchronously
func (client *Client) PhoneNumberStatusForAccountWithCallback(request *PhoneNumberStatusForAccountRequest, callback func(response *PhoneNumberStatusForAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PhoneNumberStatusForAccountResponse
		var err error
		defer close(result)
		response, err = client.PhoneNumberStatusForAccount(request)
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

// PhoneNumberStatusForAccountRequest is the request struct for api PhoneNumberStatusForAccount
type PhoneNumberStatusForAccountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	RouteName            string           `position:"Query" name:"RouteName"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResultCount          string           `position:"Query" name:"ResultCount"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AuthCode             string           `position:"Query" name:"AuthCode"`
	InputNumber          string           `position:"Query" name:"InputNumber"`
	FlowName             string           `position:"Query" name:"FlowName"`
	PhoneStatusSceneCode string           `position:"Query" name:"PhoneStatusSceneCode"`
}

// PhoneNumberStatusForAccountResponse is the response struct for api PhoneNumberStatusForAccount
type PhoneNumberStatusForAccountResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePhoneNumberStatusForAccountRequest creates a request to invoke PhoneNumberStatusForAccount API
func CreatePhoneNumberStatusForAccountRequest() (request *PhoneNumberStatusForAccountRequest) {
	request = &PhoneNumberStatusForAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "PhoneNumberStatusForAccount", "", "")
	request.Method = requests.POST
	return
}

// CreatePhoneNumberStatusForAccountResponse creates a response to parse from PhoneNumberStatusForAccount response
func CreatePhoneNumberStatusForAccountResponse() (response *PhoneNumberStatusForAccountResponse) {
	response = &PhoneNumberStatusForAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
