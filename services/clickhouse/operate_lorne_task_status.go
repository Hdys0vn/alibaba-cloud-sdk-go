package clickhouse

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

// OperateLorneTaskStatus invokes the clickhouse.OperateLorneTaskStatus API synchronously
func (client *Client) OperateLorneTaskStatus(request *OperateLorneTaskStatusRequest) (response *OperateLorneTaskStatusResponse, err error) {
	response = CreateOperateLorneTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// OperateLorneTaskStatusWithChan invokes the clickhouse.OperateLorneTaskStatus API asynchronously
func (client *Client) OperateLorneTaskStatusWithChan(request *OperateLorneTaskStatusRequest) (<-chan *OperateLorneTaskStatusResponse, <-chan error) {
	responseChan := make(chan *OperateLorneTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateLorneTaskStatus(request)
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

// OperateLorneTaskStatusWithCallback invokes the clickhouse.OperateLorneTaskStatus API asynchronously
func (client *Client) OperateLorneTaskStatusWithCallback(request *OperateLorneTaskStatusRequest, callback func(response *OperateLorneTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateLorneTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.OperateLorneTaskStatus(request)
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

// OperateLorneTaskStatusRequest is the request struct for api OperateLorneTaskStatus
type OperateLorneTaskStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LorneStatus          string           `position:"Query" name:"LorneStatus"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// OperateLorneTaskStatusResponse is the response struct for api OperateLorneTaskStatus
type OperateLorneTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOperateLorneTaskStatusRequest creates a request to invoke OperateLorneTaskStatus API
func CreateOperateLorneTaskStatusRequest() (request *OperateLorneTaskStatusRequest) {
	request = &OperateLorneTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "OperateLorneTaskStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateOperateLorneTaskStatusResponse creates a response to parse from OperateLorneTaskStatus response
func CreateOperateLorneTaskStatusResponse() (response *OperateLorneTaskStatusResponse) {
	response = &OperateLorneTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
