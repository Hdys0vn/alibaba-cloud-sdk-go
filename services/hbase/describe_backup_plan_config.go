package hbase

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

// DescribeBackupPlanConfig invokes the hbase.DescribeBackupPlanConfig API synchronously
func (client *Client) DescribeBackupPlanConfig(request *DescribeBackupPlanConfigRequest) (response *DescribeBackupPlanConfigResponse, err error) {
	response = CreateDescribeBackupPlanConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupPlanConfigWithChan invokes the hbase.DescribeBackupPlanConfig API asynchronously
func (client *Client) DescribeBackupPlanConfigWithChan(request *DescribeBackupPlanConfigRequest) (<-chan *DescribeBackupPlanConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupPlanConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupPlanConfig(request)
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

// DescribeBackupPlanConfigWithCallback invokes the hbase.DescribeBackupPlanConfig API asynchronously
func (client *Client) DescribeBackupPlanConfigWithCallback(request *DescribeBackupPlanConfigRequest, callback func(response *DescribeBackupPlanConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupPlanConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupPlanConfig(request)
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

// DescribeBackupPlanConfigRequest is the request struct for api DescribeBackupPlanConfig
type DescribeBackupPlanConfigRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeBackupPlanConfigResponse is the response struct for api DescribeBackupPlanConfig
type DescribeBackupPlanConfigResponse struct {
	*responses.BaseResponse
	RequestId           string                           `json:"RequestId" xml:"RequestId"`
	FullBackupCycle     int                              `json:"FullBackupCycle" xml:"FullBackupCycle"`
	MinHFileBackupCount int                              `json:"MinHFileBackupCount" xml:"MinHFileBackupCount"`
	NextFullBackupDate  string                           `json:"NextFullBackupDate" xml:"NextFullBackupDate"`
	Tables              TablesInDescribeBackupPlanConfig `json:"Tables" xml:"Tables"`
}

// CreateDescribeBackupPlanConfigRequest creates a request to invoke DescribeBackupPlanConfig API
func CreateDescribeBackupPlanConfigRequest() (request *DescribeBackupPlanConfigRequest) {
	request = &DescribeBackupPlanConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeBackupPlanConfig", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupPlanConfigResponse creates a response to parse from DescribeBackupPlanConfig response
func CreateDescribeBackupPlanConfigResponse() (response *DescribeBackupPlanConfigResponse) {
	response = &DescribeBackupPlanConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
