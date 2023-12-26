package dds

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

// DescribeClusterRecoverTime invokes the dds.DescribeClusterRecoverTime API synchronously
func (client *Client) DescribeClusterRecoverTime(request *DescribeClusterRecoverTimeRequest) (response *DescribeClusterRecoverTimeResponse, err error) {
	response = CreateDescribeClusterRecoverTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterRecoverTimeWithChan invokes the dds.DescribeClusterRecoverTime API asynchronously
func (client *Client) DescribeClusterRecoverTimeWithChan(request *DescribeClusterRecoverTimeRequest) (<-chan *DescribeClusterRecoverTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterRecoverTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterRecoverTime(request)
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

// DescribeClusterRecoverTimeWithCallback invokes the dds.DescribeClusterRecoverTime API asynchronously
func (client *Client) DescribeClusterRecoverTimeWithCallback(request *DescribeClusterRecoverTimeRequest, callback func(response *DescribeClusterRecoverTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterRecoverTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterRecoverTime(request)
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

// DescribeClusterRecoverTimeRequest is the request struct for api DescribeClusterRecoverTime
type DescribeClusterRecoverTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeClusterRecoverTimeResponse is the response struct for api DescribeClusterRecoverTime
type DescribeClusterRecoverTimeResponse struct {
	*responses.BaseResponse
	RequestId     string                 `json:"RequestId" xml:"RequestId"`
	RestoreRanges []AvailableRestoreTime `json:"RestoreRanges" xml:"RestoreRanges"`
}

// CreateDescribeClusterRecoverTimeRequest creates a request to invoke DescribeClusterRecoverTime API
func CreateDescribeClusterRecoverTimeRequest() (request *DescribeClusterRecoverTimeRequest) {
	request = &DescribeClusterRecoverTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeClusterRecoverTime", "dds", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterRecoverTimeResponse creates a response to parse from DescribeClusterRecoverTime response
func CreateDescribeClusterRecoverTimeResponse() (response *DescribeClusterRecoverTimeResponse) {
	response = &DescribeClusterRecoverTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
