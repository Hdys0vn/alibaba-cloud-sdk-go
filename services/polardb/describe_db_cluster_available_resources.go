package polardb

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

// DescribeDBClusterAvailableResources invokes the polardb.DescribeDBClusterAvailableResources API synchronously
func (client *Client) DescribeDBClusterAvailableResources(request *DescribeDBClusterAvailableResourcesRequest) (response *DescribeDBClusterAvailableResourcesResponse, err error) {
	response = CreateDescribeDBClusterAvailableResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAvailableResourcesWithChan invokes the polardb.DescribeDBClusterAvailableResources API asynchronously
func (client *Client) DescribeDBClusterAvailableResourcesWithChan(request *DescribeDBClusterAvailableResourcesRequest) (<-chan *DescribeDBClusterAvailableResourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAvailableResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAvailableResources(request)
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

// DescribeDBClusterAvailableResourcesWithCallback invokes the polardb.DescribeDBClusterAvailableResources API asynchronously
func (client *Client) DescribeDBClusterAvailableResourcesWithCallback(request *DescribeDBClusterAvailableResourcesRequest, callback func(response *DescribeDBClusterAvailableResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAvailableResourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAvailableResources(request)
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

// DescribeDBClusterAvailableResourcesRequest is the request struct for api DescribeDBClusterAvailableResources
type DescribeDBClusterAvailableResourcesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBNodeClass          string           `position:"Query" name:"DBNodeClass"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBType               string           `position:"Query" name:"DBType"`
	DBVersion            string           `position:"Query" name:"DBVersion"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	PayType              string           `position:"Query" name:"PayType"`
}

// DescribeDBClusterAvailableResourcesResponse is the response struct for api DescribeDBClusterAvailableResources
type DescribeDBClusterAvailableResourcesResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	AvailableZones []AvailableZone `json:"AvailableZones" xml:"AvailableZones"`
}

// CreateDescribeDBClusterAvailableResourcesRequest creates a request to invoke DescribeDBClusterAvailableResources API
func CreateDescribeDBClusterAvailableResourcesRequest() (request *DescribeDBClusterAvailableResourcesRequest) {
	request = &DescribeDBClusterAvailableResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterAvailableResources", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterAvailableResourcesResponse creates a response to parse from DescribeDBClusterAvailableResources response
func CreateDescribeDBClusterAvailableResourcesResponse() (response *DescribeDBClusterAvailableResourcesResponse) {
	response = &DescribeDBClusterAvailableResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}