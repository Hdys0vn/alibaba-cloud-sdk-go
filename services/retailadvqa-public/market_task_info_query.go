package retailadvqa_public

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

// MarketTaskInfoQuery invokes the retailadvqa_public.MarketTaskInfoQuery API synchronously
func (client *Client) MarketTaskInfoQuery(request *MarketTaskInfoQueryRequest) (response *MarketTaskInfoQueryResponse, err error) {
	response = CreateMarketTaskInfoQueryResponse()
	err = client.DoAction(request, response)
	return
}

// MarketTaskInfoQueryWithChan invokes the retailadvqa_public.MarketTaskInfoQuery API asynchronously
func (client *Client) MarketTaskInfoQueryWithChan(request *MarketTaskInfoQueryRequest) (<-chan *MarketTaskInfoQueryResponse, <-chan error) {
	responseChan := make(chan *MarketTaskInfoQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MarketTaskInfoQuery(request)
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

// MarketTaskInfoQueryWithCallback invokes the retailadvqa_public.MarketTaskInfoQuery API asynchronously
func (client *Client) MarketTaskInfoQueryWithCallback(request *MarketTaskInfoQueryRequest, callback func(response *MarketTaskInfoQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MarketTaskInfoQueryResponse
		var err error
		defer close(result)
		response, err = client.MarketTaskInfoQuery(request)
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

// MarketTaskInfoQueryRequest is the request struct for api MarketTaskInfoQuery
type MarketTaskInfoQueryRequest struct {
	*requests.RpcRequest
	MarketType  string           `position:"Query" name:"MarketType"`
	AccessId    string           `position:"Query" name:"AccessId"`
	PrivateKey  string           `position:"Query" name:"PrivateKey"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// MarketTaskInfoQueryResponse is the response struct for api MarketTaskInfoQuery
type MarketTaskInfoQueryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateMarketTaskInfoQueryRequest creates a request to invoke MarketTaskInfoQuery API
func CreateMarketTaskInfoQueryRequest() (request *MarketTaskInfoQueryRequest) {
	request = &MarketTaskInfoQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "MarketTaskInfoQuery", "", "")
	request.Method = requests.POST
	return
}

// CreateMarketTaskInfoQueryResponse creates a response to parse from MarketTaskInfoQuery response
func CreateMarketTaskInfoQueryResponse() (response *MarketTaskInfoQueryResponse) {
	response = &MarketTaskInfoQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
