package gpdb

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

// DescribeDBClusterPerformance invokes the gpdb.DescribeDBClusterPerformance API synchronously
func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (response *DescribeDBClusterPerformanceResponse, err error) {
	response = CreateDescribeDBClusterPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterPerformanceWithChan invokes the gpdb.DescribeDBClusterPerformance API asynchronously
func (client *Client) DescribeDBClusterPerformanceWithChan(request *DescribeDBClusterPerformanceRequest) (<-chan *DescribeDBClusterPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterPerformance(request)
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

// DescribeDBClusterPerformanceWithCallback invokes the gpdb.DescribeDBClusterPerformance API asynchronously
func (client *Client) DescribeDBClusterPerformanceWithCallback(request *DescribeDBClusterPerformanceRequest, callback func(response *DescribeDBClusterPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterPerformance(request)
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

// DescribeDBClusterPerformanceRequest is the request struct for api DescribeDBClusterPerformance
type DescribeDBClusterPerformanceRequest struct {
	*requests.RpcRequest
	NodeType     string `position:"Query" name:"NodeType"`
	StartTime    string `position:"Query" name:"StartTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	Key          string `position:"Query" name:"Key"`
	EndTime      string `position:"Query" name:"EndTime"`
	Nodes        string `position:"Query" name:"Nodes"`
}

// DescribeDBClusterPerformanceResponse is the response struct for api DescribeDBClusterPerformance
type DescribeDBClusterPerformanceResponse struct {
	*responses.BaseResponse
	EndTime         string           `json:"EndTime" xml:"EndTime"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	StartTime       string           `json:"StartTime" xml:"StartTime"`
	DBClusterId     string           `json:"DBClusterId" xml:"DBClusterId"`
	PerformanceKeys []PerformanceKey `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

// CreateDescribeDBClusterPerformanceRequest creates a request to invoke DescribeDBClusterPerformance API
func CreateDescribeDBClusterPerformanceRequest() (request *DescribeDBClusterPerformanceRequest) {
	request = &DescribeDBClusterPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBClusterPerformance", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterPerformanceResponse creates a response to parse from DescribeDBClusterPerformance response
func CreateDescribeDBClusterPerformanceResponse() (response *DescribeDBClusterPerformanceResponse) {
	response = &DescribeDBClusterPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
