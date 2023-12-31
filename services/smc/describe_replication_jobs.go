package smc

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

// DescribeReplicationJobs invokes the smc.DescribeReplicationJobs API synchronously
func (client *Client) DescribeReplicationJobs(request *DescribeReplicationJobsRequest) (response *DescribeReplicationJobsResponse, err error) {
	response = CreateDescribeReplicationJobsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReplicationJobsWithChan invokes the smc.DescribeReplicationJobs API asynchronously
func (client *Client) DescribeReplicationJobsWithChan(request *DescribeReplicationJobsRequest) (<-chan *DescribeReplicationJobsResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicationJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicationJobs(request)
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

// DescribeReplicationJobsWithCallback invokes the smc.DescribeReplicationJobs API asynchronously
func (client *Client) DescribeReplicationJobsWithCallback(request *DescribeReplicationJobsRequest, callback func(response *DescribeReplicationJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicationJobsResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicationJobs(request)
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

// DescribeReplicationJobsRequest is the request struct for api DescribeReplicationJobs
type DescribeReplicationJobsRequest struct {
	*requests.RpcRequest
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	JobId                *[]string        `position:"Query" name:"JobId"  type:"Repeated"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	SourceId             *[]string        `position:"Query" name:"SourceId"  type:"Repeated"`
	BusinessStatus       string           `position:"Query" name:"BusinessStatus"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeReplicationJobsResponse is the response struct for api DescribeReplicationJobs
type DescribeReplicationJobsResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	ReplicationJobs ReplicationJobs `json:"ReplicationJobs" xml:"ReplicationJobs"`
}

// CreateDescribeReplicationJobsRequest creates a request to invoke DescribeReplicationJobs API
func CreateDescribeReplicationJobsRequest() (request *DescribeReplicationJobsRequest) {
	request = &DescribeReplicationJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("smc", "2019-06-01", "DescribeReplicationJobs", "smc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeReplicationJobsResponse creates a response to parse from DescribeReplicationJobs response
func CreateDescribeReplicationJobsResponse() (response *DescribeReplicationJobsResponse) {
	response = &DescribeReplicationJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
