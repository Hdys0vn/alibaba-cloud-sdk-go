package mse

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

// QuerySwimmingLaneById invokes the mse.QuerySwimmingLaneById API synchronously
func (client *Client) QuerySwimmingLaneById(request *QuerySwimmingLaneByIdRequest) (response *QuerySwimmingLaneByIdResponse, err error) {
	response = CreateQuerySwimmingLaneByIdResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySwimmingLaneByIdWithChan invokes the mse.QuerySwimmingLaneById API asynchronously
func (client *Client) QuerySwimmingLaneByIdWithChan(request *QuerySwimmingLaneByIdRequest) (<-chan *QuerySwimmingLaneByIdResponse, <-chan error) {
	responseChan := make(chan *QuerySwimmingLaneByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySwimmingLaneById(request)
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

// QuerySwimmingLaneByIdWithCallback invokes the mse.QuerySwimmingLaneById API asynchronously
func (client *Client) QuerySwimmingLaneByIdWithCallback(request *QuerySwimmingLaneByIdRequest, callback func(response *QuerySwimmingLaneByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySwimmingLaneByIdResponse
		var err error
		defer close(result)
		response, err = client.QuerySwimmingLaneById(request)
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

// QuerySwimmingLaneByIdRequest is the request struct for api QuerySwimmingLaneById
type QuerySwimmingLaneByIdRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	LaneId         requests.Integer `position:"Query" name:"LaneId"`
	Namespace      string           `position:"Query" name:"Namespace"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// QuerySwimmingLaneByIdResponse is the response struct for api QuerySwimmingLaneById
type QuerySwimmingLaneByIdResponse struct {
	*responses.BaseResponse
	RequestId      string                      `json:"RequestId" xml:"RequestId"`
	Success        bool                        `json:"Success" xml:"Success"`
	Code           int                         `json:"Code" xml:"Code"`
	ErrorCode      string                      `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int                         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                      `json:"Message" xml:"Message"`
	Data           DataInQuerySwimmingLaneById `json:"Data" xml:"Data"`
}

// CreateQuerySwimmingLaneByIdRequest creates a request to invoke QuerySwimmingLaneById API
func CreateQuerySwimmingLaneByIdRequest() (request *QuerySwimmingLaneByIdRequest) {
	request = &QuerySwimmingLaneByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QuerySwimmingLaneById", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySwimmingLaneByIdResponse creates a response to parse from QuerySwimmingLaneById response
func CreateQuerySwimmingLaneByIdResponse() (response *QuerySwimmingLaneByIdResponse) {
	response = &QuerySwimmingLaneByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
