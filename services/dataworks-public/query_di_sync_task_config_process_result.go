package dataworks_public

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

// QueryDISyncTaskConfigProcessResult invokes the dataworks_public.QueryDISyncTaskConfigProcessResult API synchronously
func (client *Client) QueryDISyncTaskConfigProcessResult(request *QueryDISyncTaskConfigProcessResultRequest) (response *QueryDISyncTaskConfigProcessResultResponse, err error) {
	response = CreateQueryDISyncTaskConfigProcessResultResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDISyncTaskConfigProcessResultWithChan invokes the dataworks_public.QueryDISyncTaskConfigProcessResult API asynchronously
func (client *Client) QueryDISyncTaskConfigProcessResultWithChan(request *QueryDISyncTaskConfigProcessResultRequest) (<-chan *QueryDISyncTaskConfigProcessResultResponse, <-chan error) {
	responseChan := make(chan *QueryDISyncTaskConfigProcessResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDISyncTaskConfigProcessResult(request)
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

// QueryDISyncTaskConfigProcessResultWithCallback invokes the dataworks_public.QueryDISyncTaskConfigProcessResult API asynchronously
func (client *Client) QueryDISyncTaskConfigProcessResultWithCallback(request *QueryDISyncTaskConfigProcessResultRequest, callback func(response *QueryDISyncTaskConfigProcessResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDISyncTaskConfigProcessResultResponse
		var err error
		defer close(result)
		response, err = client.QueryDISyncTaskConfigProcessResult(request)
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

// QueryDISyncTaskConfigProcessResultRequest is the request struct for api QueryDISyncTaskConfigProcessResult
type QueryDISyncTaskConfigProcessResultRequest struct {
	*requests.RpcRequest
	TaskType       string           `position:"Query" name:"TaskType"`
	AsyncProcessId requests.Integer `position:"Query" name:"AsyncProcessId"`
	ProjectId      requests.Integer `position:"Query" name:"ProjectId"`
}

// QueryDISyncTaskConfigProcessResultResponse is the response struct for api QueryDISyncTaskConfigProcessResult
type QueryDISyncTaskConfigProcessResultResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryDISyncTaskConfigProcessResultRequest creates a request to invoke QueryDISyncTaskConfigProcessResult API
func CreateQueryDISyncTaskConfigProcessResultRequest() (request *QueryDISyncTaskConfigProcessResultRequest) {
	request = &QueryDISyncTaskConfigProcessResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "QueryDISyncTaskConfigProcessResult", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDISyncTaskConfigProcessResultResponse creates a response to parse from QueryDISyncTaskConfigProcessResult response
func CreateQueryDISyncTaskConfigProcessResultResponse() (response *QueryDISyncTaskConfigProcessResultResponse) {
	response = &QueryDISyncTaskConfigProcessResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
