package dyvmsapi

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

// QueryCallDetailByCallId invokes the dyvmsapi.QueryCallDetailByCallId API synchronously
func (client *Client) QueryCallDetailByCallId(request *QueryCallDetailByCallIdRequest) (response *QueryCallDetailByCallIdResponse, err error) {
	response = CreateQueryCallDetailByCallIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCallDetailByCallIdWithChan invokes the dyvmsapi.QueryCallDetailByCallId API asynchronously
func (client *Client) QueryCallDetailByCallIdWithChan(request *QueryCallDetailByCallIdRequest) (<-chan *QueryCallDetailByCallIdResponse, <-chan error) {
	responseChan := make(chan *QueryCallDetailByCallIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCallDetailByCallId(request)
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

// QueryCallDetailByCallIdWithCallback invokes the dyvmsapi.QueryCallDetailByCallId API asynchronously
func (client *Client) QueryCallDetailByCallIdWithCallback(request *QueryCallDetailByCallIdRequest, callback func(response *QueryCallDetailByCallIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCallDetailByCallIdResponse
		var err error
		defer close(result)
		response, err = client.QueryCallDetailByCallId(request)
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

// QueryCallDetailByCallIdRequest is the request struct for api QueryCallDetailByCallId
type QueryCallDetailByCallIdRequest struct {
	*requests.RpcRequest
	CallId               string           `position:"Query" name:"CallId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	QueryDate            requests.Integer `position:"Query" name:"QueryDate"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdId               requests.Integer `position:"Query" name:"ProdId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryCallDetailByCallIdResponse is the response struct for api QueryCallDetailByCallId
type QueryCallDetailByCallIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateQueryCallDetailByCallIdRequest creates a request to invoke QueryCallDetailByCallId API
func CreateQueryCallDetailByCallIdRequest() (request *QueryCallDetailByCallIdRequest) {
	request = &QueryCallDetailByCallIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryCallDetailByCallId", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryCallDetailByCallIdResponse creates a response to parse from QueryCallDetailByCallId response
func CreateQueryCallDetailByCallIdResponse() (response *QueryCallDetailByCallIdResponse) {
	response = &QueryCallDetailByCallIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
