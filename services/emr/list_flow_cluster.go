package emr

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

// ListFlowCluster invokes the emr.ListFlowCluster API synchronously
func (client *Client) ListFlowCluster(request *ListFlowClusterRequest) (response *ListFlowClusterResponse, err error) {
	response = CreateListFlowClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowClusterWithChan invokes the emr.ListFlowCluster API asynchronously
func (client *Client) ListFlowClusterWithChan(request *ListFlowClusterRequest) (<-chan *ListFlowClusterResponse, <-chan error) {
	responseChan := make(chan *ListFlowClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowCluster(request)
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

// ListFlowClusterWithCallback invokes the emr.ListFlowCluster API asynchronously
func (client *Client) ListFlowClusterWithCallback(request *ListFlowClusterRequest, callback func(response *ListFlowClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowClusterResponse
		var err error
		defer close(result)
		response, err = client.ListFlowCluster(request)
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

// ListFlowClusterRequest is the request struct for api ListFlowCluster
type ListFlowClusterRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ProjectId       string           `position:"Query" name:"ProjectId"`
}

// ListFlowClusterResponse is the response struct for api ListFlowCluster
type ListFlowClusterResponse struct {
	*responses.BaseResponse
	RequestId  string                    `json:"RequestId" xml:"RequestId"`
	TotalCount int                       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                       `json:"PageSize" xml:"PageSize"`
	Clusters   ClustersInListFlowCluster `json:"Clusters" xml:"Clusters"`
}

// CreateListFlowClusterRequest creates a request to invoke ListFlowCluster API
func CreateListFlowClusterRequest() (request *ListFlowClusterRequest) {
	request = &ListFlowClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowCluster", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFlowClusterResponse creates a response to parse from ListFlowCluster response
func CreateListFlowClusterResponse() (response *ListFlowClusterResponse) {
	response = &ListFlowClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
