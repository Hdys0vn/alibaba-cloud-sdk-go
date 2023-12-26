package nas

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

// DescribeDataFlowTasks invokes the nas.DescribeDataFlowTasks API synchronously
func (client *Client) DescribeDataFlowTasks(request *DescribeDataFlowTasksRequest) (response *DescribeDataFlowTasksResponse, err error) {
	response = CreateDescribeDataFlowTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataFlowTasksWithChan invokes the nas.DescribeDataFlowTasks API asynchronously
func (client *Client) DescribeDataFlowTasksWithChan(request *DescribeDataFlowTasksRequest) (<-chan *DescribeDataFlowTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeDataFlowTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataFlowTasks(request)
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

// DescribeDataFlowTasksWithCallback invokes the nas.DescribeDataFlowTasks API asynchronously
func (client *Client) DescribeDataFlowTasksWithCallback(request *DescribeDataFlowTasksRequest, callback func(response *DescribeDataFlowTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataFlowTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataFlowTasks(request)
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

// DescribeDataFlowTasksRequest is the request struct for api DescribeDataFlowTasks
type DescribeDataFlowTasksRequest struct {
	*requests.RpcRequest
	NextToken    string                          `position:"Query" name:"NextToken"`
	FileSystemId string                          `position:"Query" name:"FileSystemId"`
	Filters      *[]DescribeDataFlowTasksFilters `position:"Query" name:"Filters"  type:"Repeated"`
	MaxResults   requests.Integer                `position:"Query" name:"MaxResults"`
}

// DescribeDataFlowTasksFilters is a repeated param struct in DescribeDataFlowTasksRequest
type DescribeDataFlowTasksFilters struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDataFlowTasksResponse is the response struct for api DescribeDataFlowTasks
type DescribeDataFlowTasksResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	NextToken string   `json:"NextToken" xml:"NextToken"`
	TaskInfo  TaskInfo `json:"TaskInfo" xml:"TaskInfo"`
}

// CreateDescribeDataFlowTasksRequest creates a request to invoke DescribeDataFlowTasks API
func CreateDescribeDataFlowTasksRequest() (request *DescribeDataFlowTasksRequest) {
	request = &DescribeDataFlowTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeDataFlowTasks", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataFlowTasksResponse creates a response to parse from DescribeDataFlowTasks response
func CreateDescribeDataFlowTasksResponse() (response *DescribeDataFlowTasksResponse) {
	response = &DescribeDataFlowTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}