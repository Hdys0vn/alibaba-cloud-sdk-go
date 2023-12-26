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

// DeleteRouteStrategy invokes the schedulerx2.DeleteRouteStrategy API synchronously
func (client *Client) DeleteRouteStrategy(request *DeleteRouteStrategyRequest) (response *DeleteRouteStrategyResponse, err error) {
	response = CreateDeleteRouteStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRouteStrategyWithChan invokes the schedulerx2.DeleteRouteStrategy API asynchronously
func (client *Client) DeleteRouteStrategyWithChan(request *DeleteRouteStrategyRequest) (<-chan *DeleteRouteStrategyResponse, <-chan error) {
	responseChan := make(chan *DeleteRouteStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRouteStrategy(request)
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

// DeleteRouteStrategyWithCallback invokes the schedulerx2.DeleteRouteStrategy API asynchronously
func (client *Client) DeleteRouteStrategyWithCallback(request *DeleteRouteStrategyRequest, callback func(response *DeleteRouteStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRouteStrategyResponse
		var err error
		defer close(result)
		response, err = client.DeleteRouteStrategy(request)
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

// DeleteRouteStrategyRequest is the request struct for api DeleteRouteStrategy
type DeleteRouteStrategyRequest struct {
	*requests.RpcRequest
	GroupId    string           `position:"Query" name:"GroupId"`
	JobId      requests.Integer `position:"Query" name:"JobId"`
	Namespace  string           `position:"Query" name:"Namespace"`
	StrategyId requests.Integer `position:"Query" name:"StrategyId"`
}

// DeleteRouteStrategyResponse is the response struct for api DeleteRouteStrategy
type DeleteRouteStrategyResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteRouteStrategyRequest creates a request to invoke DeleteRouteStrategy API
func CreateDeleteRouteStrategyRequest() (request *DeleteRouteStrategyRequest) {
	request = &DeleteRouteStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "DeleteRouteStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRouteStrategyResponse creates a response to parse from DeleteRouteStrategy response
func CreateDeleteRouteStrategyResponse() (response *DeleteRouteStrategyResponse) {
	response = &DeleteRouteStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
