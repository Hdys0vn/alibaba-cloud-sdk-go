package schedulerx2

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

// BatchDeleteRouteStrategy invokes the schedulerx2.BatchDeleteRouteStrategy API synchronously
func (client *Client) BatchDeleteRouteStrategy(request *BatchDeleteRouteStrategyRequest) (response *BatchDeleteRouteStrategyResponse, err error) {
	response = CreateBatchDeleteRouteStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteRouteStrategyWithChan invokes the schedulerx2.BatchDeleteRouteStrategy API asynchronously
func (client *Client) BatchDeleteRouteStrategyWithChan(request *BatchDeleteRouteStrategyRequest) (<-chan *BatchDeleteRouteStrategyResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteRouteStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteRouteStrategy(request)
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

// BatchDeleteRouteStrategyWithCallback invokes the schedulerx2.BatchDeleteRouteStrategy API asynchronously
func (client *Client) BatchDeleteRouteStrategyWithCallback(request *BatchDeleteRouteStrategyRequest, callback func(response *BatchDeleteRouteStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteRouteStrategyResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteRouteStrategy(request)
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

// BatchDeleteRouteStrategyRequest is the request struct for api BatchDeleteRouteStrategy
type BatchDeleteRouteStrategyRequest struct {
	*requests.RpcRequest
	GroupId        string    `position:"Query" name:"GroupId"`
	JobIdList      *[]string `position:"Body" name:"JobIdList"  type:"Repeated"`
	Namespace      string    `position:"Query" name:"Namespace"`
	StrategyIdList *[]string `position:"Body" name:"StrategyIdList"  type:"Repeated"`
}

// BatchDeleteRouteStrategyResponse is the response struct for api BatchDeleteRouteStrategy
type BatchDeleteRouteStrategyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateBatchDeleteRouteStrategyRequest creates a request to invoke BatchDeleteRouteStrategy API
func CreateBatchDeleteRouteStrategyRequest() (request *BatchDeleteRouteStrategyRequest) {
	request = &BatchDeleteRouteStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "BatchDeleteRouteStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchDeleteRouteStrategyResponse creates a response to parse from BatchDeleteRouteStrategy response
func CreateBatchDeleteRouteStrategyResponse() (response *BatchDeleteRouteStrategyResponse) {
	response = &BatchDeleteRouteStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}