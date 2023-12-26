package httpdns

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

// GetResolveCountSummary invokes the httpdns.GetResolveCountSummary API synchronously
func (client *Client) GetResolveCountSummary(request *GetResolveCountSummaryRequest) (response *GetResolveCountSummaryResponse, err error) {
	response = CreateGetResolveCountSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetResolveCountSummaryWithChan invokes the httpdns.GetResolveCountSummary API asynchronously
func (client *Client) GetResolveCountSummaryWithChan(request *GetResolveCountSummaryRequest) (<-chan *GetResolveCountSummaryResponse, <-chan error) {
	responseChan := make(chan *GetResolveCountSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResolveCountSummary(request)
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

// GetResolveCountSummaryWithCallback invokes the httpdns.GetResolveCountSummary API asynchronously
func (client *Client) GetResolveCountSummaryWithCallback(request *GetResolveCountSummaryRequest, callback func(response *GetResolveCountSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResolveCountSummaryResponse
		var err error
		defer close(result)
		response, err = client.GetResolveCountSummary(request)
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

// GetResolveCountSummaryRequest is the request struct for api GetResolveCountSummary
type GetResolveCountSummaryRequest struct {
	*requests.RpcRequest
	TimeSpan    requests.Integer `position:"Query" name:"TimeSpan"`
	Granularity string           `position:"Query" name:"Granularity"`
}

// GetResolveCountSummaryResponse is the response struct for api GetResolveCountSummary
type GetResolveCountSummaryResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ResolveSummary ResolveSummary `json:"ResolveSummary" xml:"ResolveSummary"`
}

// CreateGetResolveCountSummaryRequest creates a request to invoke GetResolveCountSummary API
func CreateGetResolveCountSummaryRequest() (request *GetResolveCountSummaryRequest) {
	request = &GetResolveCountSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Httpdns", "2016-02-01", "GetResolveCountSummary", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResolveCountSummaryResponse creates a response to parse from GetResolveCountSummary response
func CreateGetResolveCountSummaryResponse() (response *GetResolveCountSummaryResponse) {
	response = &GetResolveCountSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
