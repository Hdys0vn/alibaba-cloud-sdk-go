package smartag

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

// DescribeUserFlowStatistics invokes the smartag.DescribeUserFlowStatistics API synchronously
func (client *Client) DescribeUserFlowStatistics(request *DescribeUserFlowStatisticsRequest) (response *DescribeUserFlowStatisticsResponse, err error) {
	response = CreateDescribeUserFlowStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserFlowStatisticsWithChan invokes the smartag.DescribeUserFlowStatistics API asynchronously
func (client *Client) DescribeUserFlowStatisticsWithChan(request *DescribeUserFlowStatisticsRequest) (<-chan *DescribeUserFlowStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserFlowStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserFlowStatistics(request)
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

// DescribeUserFlowStatisticsWithCallback invokes the smartag.DescribeUserFlowStatistics API asynchronously
func (client *Client) DescribeUserFlowStatisticsWithCallback(request *DescribeUserFlowStatisticsRequest, callback func(response *DescribeUserFlowStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserFlowStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserFlowStatistics(request)
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

// DescribeUserFlowStatisticsRequest is the request struct for api DescribeUserFlowStatistics
type DescribeUserFlowStatisticsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UserNames            *[]string        `position:"Query" name:"UserNames"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	StatisticsDate       string           `position:"Query" name:"StatisticsDate"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// DescribeUserFlowStatisticsResponse is the response struct for api DescribeUserFlowStatistics
type DescribeUserFlowStatisticsResponse struct {
	*responses.BaseResponse
	RequestId     string                                    `json:"RequestId" xml:"RequestId"`
	SagStatistics SagStatisticsInDescribeUserFlowStatistics `json:"SagStatistics" xml:"SagStatistics"`
}

// CreateDescribeUserFlowStatisticsRequest creates a request to invoke DescribeUserFlowStatistics API
func CreateDescribeUserFlowStatisticsRequest() (request *DescribeUserFlowStatisticsRequest) {
	request = &DescribeUserFlowStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeUserFlowStatistics", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserFlowStatisticsResponse creates a response to parse from DescribeUserFlowStatistics response
func CreateDescribeUserFlowStatisticsResponse() (response *DescribeUserFlowStatisticsResponse) {
	response = &DescribeUserFlowStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
