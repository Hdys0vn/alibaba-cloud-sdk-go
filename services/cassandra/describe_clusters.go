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

// DescribeClusters invokes the cassandra.DescribeClusters API synchronously
func (client *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
	response = CreateDescribeClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClustersWithChan invokes the cassandra.DescribeClusters API asynchronously
func (client *Client) DescribeClustersWithChan(request *DescribeClustersRequest) (<-chan *DescribeClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusters(request)
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

// DescribeClustersWithCallback invokes the cassandra.DescribeClusters API asynchronously
func (client *Client) DescribeClustersWithCallback(request *DescribeClustersRequest, callback func(response *DescribeClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusters(request)
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

// DescribeClustersRequest is the request struct for api DescribeClusters
type DescribeClustersRequest struct {
	*requests.RpcRequest
	ClusterName     string                 `position:"Query" name:"ClusterName"`
	PageNumber      requests.Integer       `position:"Query" name:"PageNumber"`
	ResourceGroupId string                 `position:"Query" name:"ResourceGroupId"`
	PageSize        requests.Integer       `position:"Query" name:"PageSize"`
	Tag             *[]DescribeClustersTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DescribeClustersTag is a repeated param struct in DescribeClustersRequest
type DescribeClustersTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeClustersResponse is the response struct for api DescribeClusters
type DescribeClustersResponse struct {
	*responses.BaseResponse
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	TotalCount int64                      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                        `json:"PageSize" xml:"PageSize"`
	Clusters   ClustersInDescribeClusters `json:"Clusters" xml:"Clusters"`
}

// CreateDescribeClustersRequest creates a request to invoke DescribeClusters API
func CreateDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeClusters", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClustersResponse creates a response to parse from DescribeClusters response
func CreateDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
