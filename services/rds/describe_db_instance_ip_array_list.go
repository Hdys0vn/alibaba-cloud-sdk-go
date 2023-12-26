package rds

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

// DescribeDBInstanceIPArrayList invokes the rds.DescribeDBInstanceIPArrayList API synchronously
func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (response *DescribeDBInstanceIPArrayListResponse, err error) {
	response = CreateDescribeDBInstanceIPArrayListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceIPArrayListWithChan invokes the rds.DescribeDBInstanceIPArrayList API asynchronously
func (client *Client) DescribeDBInstanceIPArrayListWithChan(request *DescribeDBInstanceIPArrayListRequest) (<-chan *DescribeDBInstanceIPArrayListResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceIPArrayListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceIPArrayList(request)
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

// DescribeDBInstanceIPArrayListWithCallback invokes the rds.DescribeDBInstanceIPArrayList API asynchronously
func (client *Client) DescribeDBInstanceIPArrayListWithCallback(request *DescribeDBInstanceIPArrayListRequest, callback func(response *DescribeDBInstanceIPArrayListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceIPArrayListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceIPArrayList(request)
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

// DescribeDBInstanceIPArrayListRequest is the request struct for api DescribeDBInstanceIPArrayList
type DescribeDBInstanceIPArrayListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	WhitelistNetworkType string           `position:"Query" name:"WhitelistNetworkType"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceIPArrayListResponse is the response struct for api DescribeDBInstanceIPArrayList
type DescribeDBInstanceIPArrayListResponse struct {
	*responses.BaseResponse
	RequestId string                               `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeDBInstanceIPArrayList `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstanceIPArrayListRequest creates a request to invoke DescribeDBInstanceIPArrayList API
func CreateDescribeDBInstanceIPArrayListRequest() (request *DescribeDBInstanceIPArrayListRequest) {
	request = &DescribeDBInstanceIPArrayListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceIPArrayList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceIPArrayListResponse creates a response to parse from DescribeDBInstanceIPArrayList response
func CreateDescribeDBInstanceIPArrayListResponse() (response *DescribeDBInstanceIPArrayListResponse) {
	response = &DescribeDBInstanceIPArrayListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
