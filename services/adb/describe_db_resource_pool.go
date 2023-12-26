package adb

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

// DescribeDBResourcePool invokes the adb.DescribeDBResourcePool API synchronously
func (client *Client) DescribeDBResourcePool(request *DescribeDBResourcePoolRequest) (response *DescribeDBResourcePoolResponse, err error) {
	response = CreateDescribeDBResourcePoolResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBResourcePoolWithChan invokes the adb.DescribeDBResourcePool API asynchronously
func (client *Client) DescribeDBResourcePoolWithChan(request *DescribeDBResourcePoolRequest) (<-chan *DescribeDBResourcePoolResponse, <-chan error) {
	responseChan := make(chan *DescribeDBResourcePoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBResourcePool(request)
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

// DescribeDBResourcePoolWithCallback invokes the adb.DescribeDBResourcePool API asynchronously
func (client *Client) DescribeDBResourcePoolWithCallback(request *DescribeDBResourcePoolRequest, callback func(response *DescribeDBResourcePoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBResourcePoolResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBResourcePool(request)
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

// DescribeDBResourcePoolRequest is the request struct for api DescribeDBResourcePool
type DescribeDBResourcePoolRequest struct {
	*requests.RpcRequest
	PoolName             string           `position:"Query" name:"PoolName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBResourcePoolResponse is the response struct for api DescribeDBResourcePool
type DescribeDBResourcePoolResponse struct {
	*responses.BaseResponse
	RequestId   string     `json:"RequestId" xml:"RequestId"`
	DBClusterId string     `json:"DBClusterId" xml:"DBClusterId"`
	PoolsInfo   []PoolInfo `json:"PoolsInfo" xml:"PoolsInfo"`
}

// CreateDescribeDBResourcePoolRequest creates a request to invoke DescribeDBResourcePool API
func CreateDescribeDBResourcePoolRequest() (request *DescribeDBResourcePoolRequest) {
	request = &DescribeDBResourcePoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeDBResourcePool", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBResourcePoolResponse creates a response to parse from DescribeDBResourcePool response
func CreateDescribeDBResourcePoolResponse() (response *DescribeDBResourcePoolResponse) {
	response = &DescribeDBResourcePoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
