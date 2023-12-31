package cassandra

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

// DeleteNodeToolExecutionHistory invokes the cassandra.DeleteNodeToolExecutionHistory API synchronously
func (client *Client) DeleteNodeToolExecutionHistory(request *DeleteNodeToolExecutionHistoryRequest) (response *DeleteNodeToolExecutionHistoryResponse, err error) {
	response = CreateDeleteNodeToolExecutionHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNodeToolExecutionHistoryWithChan invokes the cassandra.DeleteNodeToolExecutionHistory API asynchronously
func (client *Client) DeleteNodeToolExecutionHistoryWithChan(request *DeleteNodeToolExecutionHistoryRequest) (<-chan *DeleteNodeToolExecutionHistoryResponse, <-chan error) {
	responseChan := make(chan *DeleteNodeToolExecutionHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNodeToolExecutionHistory(request)
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

// DeleteNodeToolExecutionHistoryWithCallback invokes the cassandra.DeleteNodeToolExecutionHistory API asynchronously
func (client *Client) DeleteNodeToolExecutionHistoryWithCallback(request *DeleteNodeToolExecutionHistoryRequest, callback func(response *DeleteNodeToolExecutionHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNodeToolExecutionHistoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteNodeToolExecutionHistory(request)
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

// DeleteNodeToolExecutionHistoryRequest is the request struct for api DeleteNodeToolExecutionHistory
type DeleteNodeToolExecutionHistoryRequest struct {
	*requests.RpcRequest
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	JobId        string `position:"Query" name:"JobId"`
}

// DeleteNodeToolExecutionHistoryResponse is the response struct for api DeleteNodeToolExecutionHistory
type DeleteNodeToolExecutionHistoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNodeToolExecutionHistoryRequest creates a request to invoke DeleteNodeToolExecutionHistory API
func CreateDeleteNodeToolExecutionHistoryRequest() (request *DeleteNodeToolExecutionHistoryRequest) {
	request = &DeleteNodeToolExecutionHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DeleteNodeToolExecutionHistory", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNodeToolExecutionHistoryResponse creates a response to parse from DeleteNodeToolExecutionHistory response
func CreateDeleteNodeToolExecutionHistoryResponse() (response *DeleteNodeToolExecutionHistoryResponse) {
	response = &DeleteNodeToolExecutionHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
