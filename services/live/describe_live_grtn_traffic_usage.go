package live

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

// DescribeLiveGrtnTrafficUsage invokes the live.DescribeLiveGrtnTrafficUsage API synchronously
func (client *Client) DescribeLiveGrtnTrafficUsage(request *DescribeLiveGrtnTrafficUsageRequest) (response *DescribeLiveGrtnTrafficUsageResponse, err error) {
	response = CreateDescribeLiveGrtnTrafficUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveGrtnTrafficUsageWithChan invokes the live.DescribeLiveGrtnTrafficUsage API asynchronously
func (client *Client) DescribeLiveGrtnTrafficUsageWithChan(request *DescribeLiveGrtnTrafficUsageRequest) (<-chan *DescribeLiveGrtnTrafficUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveGrtnTrafficUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveGrtnTrafficUsage(request)
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

// DescribeLiveGrtnTrafficUsageWithCallback invokes the live.DescribeLiveGrtnTrafficUsage API asynchronously
func (client *Client) DescribeLiveGrtnTrafficUsageWithCallback(request *DescribeLiveGrtnTrafficUsageRequest, callback func(response *DescribeLiveGrtnTrafficUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveGrtnTrafficUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveGrtnTrafficUsage(request)
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

// DescribeLiveGrtnTrafficUsageRequest is the request struct for api DescribeLiveGrtnTrafficUsage
type DescribeLiveGrtnTrafficUsageRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Area      string           `position:"Query" name:"Area"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Field     string           `position:"Query" name:"Field"`
	AppId     string           `position:"Query" name:"AppId"`
	Interval  string           `position:"Query" name:"Interval"`
}

// DescribeLiveGrtnTrafficUsageResponse is the response struct for api DescribeLiveGrtnTrafficUsage
type DescribeLiveGrtnTrafficUsageResponse struct {
	*responses.BaseResponse
	EndTime              string                                             `json:"EndTime" xml:"EndTime"`
	StartTime            string                                             `json:"StartTime" xml:"StartTime"`
	RequestId            string                                             `json:"RequestId" xml:"RequestId"`
	Filed                string                                             `json:"Filed" xml:"Filed"`
	AppId                string                                             `json:"AppId" xml:"AppId"`
	Area                 string                                             `json:"Area" xml:"Area"`
	UsageDataPerInterval UsageDataPerIntervalInDescribeLiveGrtnTrafficUsage `json:"UsageDataPerInterval" xml:"UsageDataPerInterval"`
}

// CreateDescribeLiveGrtnTrafficUsageRequest creates a request to invoke DescribeLiveGrtnTrafficUsage API
func CreateDescribeLiveGrtnTrafficUsageRequest() (request *DescribeLiveGrtnTrafficUsageRequest) {
	request = &DescribeLiveGrtnTrafficUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveGrtnTrafficUsage", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveGrtnTrafficUsageResponse creates a response to parse from DescribeLiveGrtnTrafficUsage response
func CreateDescribeLiveGrtnTrafficUsageResponse() (response *DescribeLiveGrtnTrafficUsageResponse) {
	response = &DescribeLiveGrtnTrafficUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}