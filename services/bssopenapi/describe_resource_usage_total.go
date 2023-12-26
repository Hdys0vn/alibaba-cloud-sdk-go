package bssopenapi

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

// DescribeResourceUsageTotal invokes the bssopenapi.DescribeResourceUsageTotal API synchronously
func (client *Client) DescribeResourceUsageTotal(request *DescribeResourceUsageTotalRequest) (response *DescribeResourceUsageTotalResponse, err error) {
	response = CreateDescribeResourceUsageTotalResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceUsageTotalWithChan invokes the bssopenapi.DescribeResourceUsageTotal API asynchronously
func (client *Client) DescribeResourceUsageTotalWithChan(request *DescribeResourceUsageTotalRequest) (<-chan *DescribeResourceUsageTotalResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceUsageTotalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceUsageTotal(request)
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

// DescribeResourceUsageTotalWithCallback invokes the bssopenapi.DescribeResourceUsageTotal API asynchronously
func (client *Client) DescribeResourceUsageTotalWithCallback(request *DescribeResourceUsageTotalRequest, callback func(response *DescribeResourceUsageTotalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceUsageTotalResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceUsageTotal(request)
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

// DescribeResourceUsageTotalRequest is the request struct for api DescribeResourceUsageTotal
type DescribeResourceUsageTotalRequest struct {
	*requests.RpcRequest
	PeriodType   string           `position:"Query" name:"PeriodType"`
	BillOwnerId  requests.Integer `position:"Query" name:"BillOwnerId"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	StartPeriod  string           `position:"Query" name:"StartPeriod"`
	EndPeriod    string           `position:"Query" name:"EndPeriod"`
}

// DescribeResourceUsageTotalResponse is the response struct for api DescribeResourceUsageTotal
type DescribeResourceUsageTotalResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeResourceUsageTotalRequest creates a request to invoke DescribeResourceUsageTotal API
func CreateDescribeResourceUsageTotalRequest() (request *DescribeResourceUsageTotalRequest) {
	request = &DescribeResourceUsageTotalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "DescribeResourceUsageTotal", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeResourceUsageTotalResponse creates a response to parse from DescribeResourceUsageTotal response
func CreateDescribeResourceUsageTotalResponse() (response *DescribeResourceUsageTotalResponse) {
	response = &DescribeResourceUsageTotalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
