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

// BatchQuerySessionByClientIds invokes the onsmqtt.BatchQuerySessionByClientIds API synchronously
func (client *Client) BatchQuerySessionByClientIds(request *BatchQuerySessionByClientIdsRequest) (response *BatchQuerySessionByClientIdsResponse, err error) {
	response = CreateBatchQuerySessionByClientIdsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchQuerySessionByClientIdsWithChan invokes the onsmqtt.BatchQuerySessionByClientIds API asynchronously
func (client *Client) BatchQuerySessionByClientIdsWithChan(request *BatchQuerySessionByClientIdsRequest) (<-chan *BatchQuerySessionByClientIdsResponse, <-chan error) {
	responseChan := make(chan *BatchQuerySessionByClientIdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchQuerySessionByClientIds(request)
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

// BatchQuerySessionByClientIdsWithCallback invokes the onsmqtt.BatchQuerySessionByClientIds API asynchronously
func (client *Client) BatchQuerySessionByClientIdsWithCallback(request *BatchQuerySessionByClientIdsRequest, callback func(response *BatchQuerySessionByClientIdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchQuerySessionByClientIdsResponse
		var err error
		defer close(result)
		response, err = client.BatchQuerySessionByClientIds(request)
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

// BatchQuerySessionByClientIdsRequest is the request struct for api BatchQuerySessionByClientIds
type BatchQuerySessionByClientIdsRequest struct {
	*requests.RpcRequest
	ClientIdList *[]string `position:"Query" name:"ClientIdList"  type:"Repeated"`
	InstanceId   string    `position:"Query" name:"InstanceId"`
}

// BatchQuerySessionByClientIdsResponse is the response struct for api BatchQuerySessionByClientIds
type BatchQuerySessionByClientIdsResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	OnlineStatusList []OnlineStatusListItem `json:"OnlineStatusList" xml:"OnlineStatusList"`
}

// CreateBatchQuerySessionByClientIdsRequest creates a request to invoke BatchQuerySessionByClientIds API
func CreateBatchQuerySessionByClientIdsRequest() (request *BatchQuerySessionByClientIdsRequest) {
	request = &BatchQuerySessionByClientIdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OnsMqtt", "2020-04-20", "BatchQuerySessionByClientIds", "onsmqtt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchQuerySessionByClientIdsResponse creates a response to parse from BatchQuerySessionByClientIds response
func CreateBatchQuerySessionByClientIdsResponse() (response *BatchQuerySessionByClientIdsResponse) {
	response = &BatchQuerySessionByClientIdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
