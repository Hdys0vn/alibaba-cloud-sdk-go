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

// DescribeDBClusterAttribute invokes the clickhouse.DescribeDBClusterAttribute API synchronously
func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (response *DescribeDBClusterAttributeResponse, err error) {
	response = CreateDescribeDBClusterAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAttributeWithChan invokes the clickhouse.DescribeDBClusterAttribute API asynchronously
func (client *Client) DescribeDBClusterAttributeWithChan(request *DescribeDBClusterAttributeRequest) (<-chan *DescribeDBClusterAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAttribute(request)
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

// DescribeDBClusterAttributeWithCallback invokes the clickhouse.DescribeDBClusterAttribute API asynchronously
func (client *Client) DescribeDBClusterAttributeWithCallback(request *DescribeDBClusterAttributeRequest, callback func(response *DescribeDBClusterAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAttribute(request)
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

// DescribeDBClusterAttributeRequest is the request struct for api DescribeDBClusterAttribute
type DescribeDBClusterAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterAttributeResponse is the response struct for api DescribeDBClusterAttribute
type DescribeDBClusterAttributeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	DBCluster DBCluster `json:"DBCluster" xml:"DBCluster"`
}

// CreateDescribeDBClusterAttributeRequest creates a request to invoke DescribeDBClusterAttribute API
func CreateDescribeDBClusterAttributeRequest() (request *DescribeDBClusterAttributeRequest) {
	request = &DescribeDBClusterAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeDBClusterAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterAttributeResponse creates a response to parse from DescribeDBClusterAttribute response
func CreateDescribeDBClusterAttributeResponse() (response *DescribeDBClusterAttributeResponse) {
	response = &DescribeDBClusterAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
