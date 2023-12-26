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

// ListClusterService invokes the emr.ListClusterService API synchronously
func (client *Client) ListClusterService(request *ListClusterServiceRequest) (response *ListClusterServiceResponse, err error) {
	response = CreateListClusterServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterServiceWithChan invokes the emr.ListClusterService API asynchronously
func (client *Client) ListClusterServiceWithChan(request *ListClusterServiceRequest) (<-chan *ListClusterServiceResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterService(request)
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

// ListClusterServiceWithCallback invokes the emr.ListClusterService API asynchronously
func (client *Client) ListClusterServiceWithCallback(request *ListClusterServiceRequest, callback func(response *ListClusterServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceResponse
		var err error
		defer close(result)
		response, err = client.ListClusterService(request)
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

// ListClusterServiceRequest is the request struct for api ListClusterService
type ListClusterServiceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// ListClusterServiceResponse is the response struct for api ListClusterService
type ListClusterServiceResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	TotalCount         int                `json:"TotalCount" xml:"TotalCount"`
	PageNumber         int                `json:"PageNumber" xml:"PageNumber"`
	PageSize           int                `json:"PageSize" xml:"PageSize"`
	ClusterServiceList ClusterServiceList `json:"ClusterServiceList" xml:"ClusterServiceList"`
}

// CreateListClusterServiceRequest creates a request to invoke ListClusterService API
func CreateListClusterServiceRequest() (request *ListClusterServiceRequest) {
	request = &ListClusterServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterService", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterServiceResponse creates a response to parse from ListClusterService response
func CreateListClusterServiceResponse() (response *ListClusterServiceResponse) {
	response = &ListClusterServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
