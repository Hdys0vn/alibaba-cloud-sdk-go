package r_kvstore

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

// DescribeTasks invokes the r_kvstore.DescribeTasks API synchronously
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	response = CreateDescribeTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTasksWithChan invokes the r_kvstore.DescribeTasks API asynchronously
func (client *Client) DescribeTasksWithChan(request *DescribeTasksRequest) (<-chan *DescribeTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTasks(request)
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

// DescribeTasksWithCallback invokes the r_kvstore.DescribeTasks API asynchronously
func (client *Client) DescribeTasksWithCallback(request *DescribeTasksRequest, callback func(response *DescribeTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeTasks(request)
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

// DescribeTasksRequest is the request struct for api DescribeTasks
type DescribeTasksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeTasksResponse is the response struct for api DescribeTasks
type DescribeTasksResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	PageNumber       int                `json:"PageNumber" xml:"PageNumber"`
	PageSize         int                `json:"PageSize" xml:"PageSize"`
	TotalRecordCount int                `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            []TaskProgressInfo `json:"Items" xml:"Items"`
}

// CreateDescribeTasksRequest creates a request to invoke DescribeTasks API
func CreateDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeTasks", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTasksResponse creates a response to parse from DescribeTasks response
func CreateDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
