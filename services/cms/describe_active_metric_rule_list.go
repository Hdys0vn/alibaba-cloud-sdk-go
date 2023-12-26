package cms

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

// DescribeActiveMetricRuleList invokes the cms.DescribeActiveMetricRuleList API synchronously
func (client *Client) DescribeActiveMetricRuleList(request *DescribeActiveMetricRuleListRequest) (response *DescribeActiveMetricRuleListResponse, err error) {
	response = CreateDescribeActiveMetricRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveMetricRuleListWithChan invokes the cms.DescribeActiveMetricRuleList API asynchronously
func (client *Client) DescribeActiveMetricRuleListWithChan(request *DescribeActiveMetricRuleListRequest) (<-chan *DescribeActiveMetricRuleListResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveMetricRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveMetricRuleList(request)
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

// DescribeActiveMetricRuleListWithCallback invokes the cms.DescribeActiveMetricRuleList API asynchronously
func (client *Client) DescribeActiveMetricRuleListWithCallback(request *DescribeActiveMetricRuleListRequest, callback func(response *DescribeActiveMetricRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveMetricRuleListResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveMetricRuleList(request)
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

// DescribeActiveMetricRuleListRequest is the request struct for api DescribeActiveMetricRuleList
type DescribeActiveMetricRuleListRequest struct {
	*requests.RpcRequest
	Product string `position:"Query" name:"Product"`
}

// DescribeActiveMetricRuleListResponse is the response struct for api DescribeActiveMetricRuleList
type DescribeActiveMetricRuleListResponse struct {
	*responses.BaseResponse
	Code       string                                   `json:"Code" xml:"Code"`
	Message    string                                   `json:"Message" xml:"Message"`
	RequestId  string                                   `json:"RequestId" xml:"RequestId"`
	Success    bool                                     `json:"Success" xml:"Success"`
	Datapoints DatapointsInDescribeActiveMetricRuleList `json:"Datapoints" xml:"Datapoints"`
	AlertList  AlertList                                `json:"AlertList" xml:"AlertList"`
}

// CreateDescribeActiveMetricRuleListRequest creates a request to invoke DescribeActiveMetricRuleList API
func CreateDescribeActiveMetricRuleListRequest() (request *DescribeActiveMetricRuleListRequest) {
	request = &DescribeActiveMetricRuleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeActiveMetricRuleList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeActiveMetricRuleListResponse creates a response to parse from DescribeActiveMetricRuleList response
func CreateDescribeActiveMetricRuleListResponse() (response *DescribeActiveMetricRuleListResponse) {
	response = &DescribeActiveMetricRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
