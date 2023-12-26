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

// PhoneNumberStatusForVirtual invokes the dytnsapi.PhoneNumberStatusForVirtual API synchronously
func (client *Client) PhoneNumberStatusForVirtual(request *PhoneNumberStatusForVirtualRequest) (response *PhoneNumberStatusForVirtualResponse, err error) {
	response = CreatePhoneNumberStatusForVirtualResponse()
	err = client.DoAction(request, response)
	return
}

// PhoneNumberStatusForVirtualWithChan invokes the dytnsapi.PhoneNumberStatusForVirtual API asynchronously
func (client *Client) PhoneNumberStatusForVirtualWithChan(request *PhoneNumberStatusForVirtualRequest) (<-chan *PhoneNumberStatusForVirtualResponse, <-chan error) {
	responseChan := make(chan *PhoneNumberStatusForVirtualResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PhoneNumberStatusForVirtual(request)
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

// PhoneNumberStatusForVirtualWithCallback invokes the dytnsapi.PhoneNumberStatusForVirtual API asynchronously
func (client *Client) PhoneNumberStatusForVirtualWithCallback(request *PhoneNumberStatusForVirtualRequest, callback func(response *PhoneNumberStatusForVirtualResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PhoneNumberStatusForVirtualResponse
		var err error
		defer close(result)
		response, err = client.PhoneNumberStatusForVirtual(request)
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

// PhoneNumberStatusForVirtualRequest is the request struct for api PhoneNumberStatusForVirtual
type PhoneNumberStatusForVirtualRequest struct {
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

// PhoneNumberStatusForVirtualResponse is the response struct for api PhoneNumberStatusForVirtual
type PhoneNumberStatusForVirtualResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePhoneNumberStatusForVirtualRequest creates a request to invoke PhoneNumberStatusForVirtual API
func CreatePhoneNumberStatusForVirtualRequest() (request *PhoneNumberStatusForVirtualRequest) {
	request = &PhoneNumberStatusForVirtualRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "PhoneNumberStatusForVirtual", "", "")
	request.Method = requests.POST
	return
}

// CreatePhoneNumberStatusForVirtualResponse creates a response to parse from PhoneNumberStatusForVirtual response
func CreatePhoneNumberStatusForVirtualResponse() (response *PhoneNumberStatusForVirtualResponse) {
	response = &PhoneNumberStatusForVirtualResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
