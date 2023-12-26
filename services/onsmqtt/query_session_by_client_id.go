package onsmqtt

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

// QuerySessionByClientId invokes the onsmqtt.QuerySessionByClientId API synchronously
func (client *Client) QuerySessionByClientId(request *QuerySessionByClientIdRequest) (response *QuerySessionByClientIdResponse, err error) {
	response = CreateQuerySessionByClientIdResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySessionByClientIdWithChan invokes the onsmqtt.QuerySessionByClientId API asynchronously
func (client *Client) QuerySessionByClientIdWithChan(request *QuerySessionByClientIdRequest) (<-chan *QuerySessionByClientIdResponse, <-chan error) {
	responseChan := make(chan *QuerySessionByClientIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySessionByClientId(request)
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

// QuerySessionByClientIdWithCallback invokes the onsmqtt.QuerySessionByClientId API asynchronously
func (client *Client) QuerySessionByClientIdWithCallback(request *QuerySessionByClientIdRequest, callback func(response *QuerySessionByClientIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySessionByClientIdResponse
		var err error
		defer close(result)
		response, err = client.QuerySessionByClientId(request)
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

// QuerySessionByClientIdRequest is the request struct for api QuerySessionByClientId
type QuerySessionByClientIdRequest struct {
	*requests.RpcRequest
	ClientId   string `position:"Query" name:"ClientId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// QuerySessionByClientIdResponse is the response struct for api QuerySessionByClientId
type QuerySessionByClientIdResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	OnlineStatus bool   `json:"OnlineStatus" xml:"OnlineStatus"`
}

// CreateQuerySessionByClientIdRequest creates a request to invoke QuerySessionByClientId API
func CreateQuerySessionByClientIdRequest() (request *QuerySessionByClientIdRequest) {
	request = &QuerySessionByClientIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OnsMqtt", "2020-04-20", "QuerySessionByClientId", "onsmqtt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySessionByClientIdResponse creates a response to parse from QuerySessionByClientId response
func CreateQuerySessionByClientIdResponse() (response *QuerySessionByClientIdResponse) {
	response = &QuerySessionByClientIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
