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

// RetrieveCall invokes the ccc.RetrieveCall API synchronously
func (client *Client) RetrieveCall(request *RetrieveCallRequest) (response *RetrieveCallResponse, err error) {
	response = CreateRetrieveCallResponse()
	err = client.DoAction(request, response)
	return
}

// RetrieveCallWithChan invokes the ccc.RetrieveCall API asynchronously
func (client *Client) RetrieveCallWithChan(request *RetrieveCallRequest) (<-chan *RetrieveCallResponse, <-chan error) {
	responseChan := make(chan *RetrieveCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetrieveCall(request)
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

// RetrieveCallWithCallback invokes the ccc.RetrieveCall API asynchronously
func (client *Client) RetrieveCallWithCallback(request *RetrieveCallRequest, callback func(response *RetrieveCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetrieveCallResponse
		var err error
		defer close(result)
		response, err = client.RetrieveCall(request)
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

// RetrieveCallRequest is the request struct for api RetrieveCall
type RetrieveCallRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	DeviceId   string `position:"Query" name:"DeviceId"`
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	ChannelId  string `position:"Query" name:"ChannelId"`
}

// RetrieveCallResponse is the response struct for api RetrieveCall
type RetrieveCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateRetrieveCallRequest creates a request to invoke RetrieveCall API
func CreateRetrieveCallRequest() (request *RetrieveCallRequest) {
	request = &RetrieveCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "RetrieveCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetrieveCallResponse creates a response to parse from RetrieveCall response
func CreateRetrieveCallResponse() (response *RetrieveCallResponse) {
	response = &RetrieveCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}