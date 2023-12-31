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

// ExecuteNodeTool invokes the cassandra.ExecuteNodeTool API synchronously
func (client *Client) ExecuteNodeTool(request *ExecuteNodeToolRequest) (response *ExecuteNodeToolResponse, err error) {
	response = CreateExecuteNodeToolResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteNodeToolWithChan invokes the cassandra.ExecuteNodeTool API asynchronously
func (client *Client) ExecuteNodeToolWithChan(request *ExecuteNodeToolRequest) (<-chan *ExecuteNodeToolResponse, <-chan error) {
	responseChan := make(chan *ExecuteNodeToolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteNodeTool(request)
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

// ExecuteNodeToolWithCallback invokes the cassandra.ExecuteNodeTool API asynchronously
func (client *Client) ExecuteNodeToolWithCallback(request *ExecuteNodeToolRequest, callback func(response *ExecuteNodeToolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteNodeToolResponse
		var err error
		defer close(result)
		response, err = client.ExecuteNodeTool(request)
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

// ExecuteNodeToolRequest is the request struct for api ExecuteNodeTool
type ExecuteNodeToolRequest struct {
	*requests.RpcRequest
	ExecuteNodes string `position:"Query" name:"ExecuteNodes"`
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	Command      string `position:"Query" name:"Command"`
	Arguments    string `position:"Query" name:"Arguments"`
}

// ExecuteNodeToolResponse is the response struct for api ExecuteNodeTool
type ExecuteNodeToolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExecuteNodeToolRequest creates a request to invoke ExecuteNodeTool API
func CreateExecuteNodeToolRequest() (request *ExecuteNodeToolRequest) {
	request = &ExecuteNodeToolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "ExecuteNodeTool", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExecuteNodeToolResponse creates a response to parse from ExecuteNodeTool response
func CreateExecuteNodeToolResponse() (response *ExecuteNodeToolResponse) {
	response = &ExecuteNodeToolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
