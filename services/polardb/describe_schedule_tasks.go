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

// DescribeScheduleTasks invokes the polardb.DescribeScheduleTasks API synchronously
func (client *Client) DescribeScheduleTasks(request *DescribeScheduleTasksRequest) (response *DescribeScheduleTasksResponse, err error) {
	response = CreateDescribeScheduleTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScheduleTasksWithChan invokes the polardb.DescribeScheduleTasks API asynchronously
func (client *Client) DescribeScheduleTasksWithChan(request *DescribeScheduleTasksRequest) (<-chan *DescribeScheduleTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeScheduleTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScheduleTasks(request)
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

// DescribeScheduleTasksWithCallback invokes the polardb.DescribeScheduleTasks API asynchronously
func (client *Client) DescribeScheduleTasksWithCallback(request *DescribeScheduleTasksRequest, callback func(response *DescribeScheduleTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScheduleTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeScheduleTasks(request)
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

// DescribeScheduleTasksRequest is the request struct for api DescribeScheduleTasks
type DescribeScheduleTasksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string           `position:"Query" name:"DBClusterDescription"`
	PlannedEndTime       string           `position:"Query" name:"PlannedEndTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OrderId              string           `position:"Query" name:"OrderId"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime     string           `position:"Query" name:"PlannedStartTime"`
	TaskAction           string           `position:"Query" name:"TaskAction"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeScheduleTasksResponse is the response struct for api DescribeScheduleTasks
type DescribeScheduleTasksResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeScheduleTasksRequest creates a request to invoke DescribeScheduleTasks API
func CreateDescribeScheduleTasksRequest() (request *DescribeScheduleTasksRequest) {
	request = &DescribeScheduleTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeScheduleTasks", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeScheduleTasksResponse creates a response to parse from DescribeScheduleTasks response
func CreateDescribeScheduleTasksResponse() (response *DescribeScheduleTasksResponse) {
	response = &DescribeScheduleTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}