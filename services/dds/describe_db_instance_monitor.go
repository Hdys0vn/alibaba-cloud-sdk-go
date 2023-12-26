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

// DescribeDBInstanceMonitor invokes the dds.DescribeDBInstanceMonitor API synchronously
func (client *Client) DescribeDBInstanceMonitor(request *DescribeDBInstanceMonitorRequest) (response *DescribeDBInstanceMonitorResponse, err error) {
	response = CreateDescribeDBInstanceMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceMonitorWithChan invokes the dds.DescribeDBInstanceMonitor API asynchronously
func (client *Client) DescribeDBInstanceMonitorWithChan(request *DescribeDBInstanceMonitorRequest) (<-chan *DescribeDBInstanceMonitorResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceMonitor(request)
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

// DescribeDBInstanceMonitorWithCallback invokes the dds.DescribeDBInstanceMonitor API asynchronously
func (client *Client) DescribeDBInstanceMonitorWithCallback(request *DescribeDBInstanceMonitorRequest, callback func(response *DescribeDBInstanceMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceMonitorResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceMonitor(request)
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

// DescribeDBInstanceMonitorRequest is the request struct for api DescribeDBInstanceMonitor
type DescribeDBInstanceMonitorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceMonitorResponse is the response struct for api DescribeDBInstanceMonitor
type DescribeDBInstanceMonitorResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Granularity string `json:"Granularity" xml:"Granularity"`
}

// CreateDescribeDBInstanceMonitorRequest creates a request to invoke DescribeDBInstanceMonitor API
func CreateDescribeDBInstanceMonitorRequest() (request *DescribeDBInstanceMonitorRequest) {
	request = &DescribeDBInstanceMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstanceMonitor", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceMonitorResponse creates a response to parse from DescribeDBInstanceMonitor response
func CreateDescribeDBInstanceMonitorResponse() (response *DescribeDBInstanceMonitorResponse) {
	response = &DescribeDBInstanceMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}