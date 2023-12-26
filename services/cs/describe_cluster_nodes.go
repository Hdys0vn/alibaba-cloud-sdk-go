package cs

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

// DescribeClusterNodes invokes the cs.DescribeClusterNodes API synchronously
func (client *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
	response = CreateDescribeClusterNodesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterNodesWithChan invokes the cs.DescribeClusterNodes API asynchronously
func (client *Client) DescribeClusterNodesWithChan(request *DescribeClusterNodesRequest) (<-chan *DescribeClusterNodesResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterNodes(request)
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

// DescribeClusterNodesWithCallback invokes the cs.DescribeClusterNodes API asynchronously
func (client *Client) DescribeClusterNodesWithCallback(request *DescribeClusterNodesRequest, callback func(response *DescribeClusterNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterNodesResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterNodes(request)
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

// DescribeClusterNodesRequest is the request struct for api DescribeClusterNodes
type DescribeClusterNodesRequest struct {
	*requests.RoaRequest
	PageSize   string `position:"Query" name:"pageSize"`
	ClusterId  string `position:"Path" name:"ClusterId"`
	State      string `position:"Query" name:"state"`
	NodepoolId string `position:"Query" name:"nodepool_id"`
	PageNumber string `position:"Query" name:"pageNumber"`
}

// DescribeClusterNodesResponse is the response struct for api DescribeClusterNodes
type DescribeClusterNodesResponse struct {
	*responses.BaseResponse
	Page  Page   `json:"page" xml:"page"`
	Nodes []Node `json:"nodes" xml:"nodes"`
}

// CreateDescribeClusterNodesRequest creates a request to invoke DescribeClusterNodes API
func CreateDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
	request = &DescribeClusterNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodes", "/clusters/[ClusterId]/nodes", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterNodesResponse creates a response to parse from DescribeClusterNodes response
func CreateDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
	response = &DescribeClusterNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
