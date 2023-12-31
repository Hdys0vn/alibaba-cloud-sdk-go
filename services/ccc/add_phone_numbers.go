package ccc

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

// AddPhoneNumbers invokes the ccc.AddPhoneNumbers API synchronously
func (client *Client) AddPhoneNumbers(request *AddPhoneNumbersRequest) (response *AddPhoneNumbersResponse, err error) {
	response = CreateAddPhoneNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// AddPhoneNumbersWithChan invokes the ccc.AddPhoneNumbers API asynchronously
func (client *Client) AddPhoneNumbersWithChan(request *AddPhoneNumbersRequest) (<-chan *AddPhoneNumbersResponse, <-chan error) {
	responseChan := make(chan *AddPhoneNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPhoneNumbers(request)
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

// AddPhoneNumbersWithCallback invokes the ccc.AddPhoneNumbers API asynchronously
func (client *Client) AddPhoneNumbersWithCallback(request *AddPhoneNumbersRequest, callback func(response *AddPhoneNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPhoneNumbersResponse
		var err error
		defer close(result)
		response, err = client.AddPhoneNumbers(request)
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

// AddPhoneNumbersRequest is the request struct for api AddPhoneNumbers
type AddPhoneNumbersRequest struct {
	*requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	Usage         string `position:"Query" name:"Usage"`
	NumberGroupId string `position:"Query" name:"NumberGroupId"`
	NumberList    string `position:"Query" name:"NumberList"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// AddPhoneNumbersResponse is the response struct for api AddPhoneNumbers
type AddPhoneNumbersResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Data           []string `json:"Data" xml:"Data"`
}

// CreateAddPhoneNumbersRequest creates a request to invoke AddPhoneNumbers API
func CreateAddPhoneNumbersRequest() (request *AddPhoneNumbersRequest) {
	request = &AddPhoneNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "AddPhoneNumbers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddPhoneNumbersResponse creates a response to parse from AddPhoneNumbers response
func CreateAddPhoneNumbersResponse() (response *AddPhoneNumbersResponse) {
	response = &AddPhoneNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
