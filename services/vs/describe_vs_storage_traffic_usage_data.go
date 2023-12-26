package vs

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

// DescribeVsStorageTrafficUsageData invokes the vs.DescribeVsStorageTrafficUsageData API synchronously
func (client *Client) DescribeVsStorageTrafficUsageData(request *DescribeVsStorageTrafficUsageDataRequest) (response *DescribeVsStorageTrafficUsageDataResponse, err error) {
	response = CreateDescribeVsStorageTrafficUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsStorageTrafficUsageDataWithChan invokes the vs.DescribeVsStorageTrafficUsageData API asynchronously
func (client *Client) DescribeVsStorageTrafficUsageDataWithChan(request *DescribeVsStorageTrafficUsageDataRequest) (<-chan *DescribeVsStorageTrafficUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVsStorageTrafficUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsStorageTrafficUsageData(request)
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

// DescribeVsStorageTrafficUsageDataWithCallback invokes the vs.DescribeVsStorageTrafficUsageData API asynchronously
func (client *Client) DescribeVsStorageTrafficUsageDataWithCallback(request *DescribeVsStorageTrafficUsageDataRequest, callback func(response *DescribeVsStorageTrafficUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsStorageTrafficUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsStorageTrafficUsageData(request)
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

// DescribeVsStorageTrafficUsageDataRequest is the request struct for api DescribeVsStorageTrafficUsageData
type DescribeVsStorageTrafficUsageDataRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	SplitBy   string           `position:"Query" name:"SplitBy"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Bucket    string           `position:"Query" name:"Bucket"`
	Interval  string           `position:"Query" name:"Interval"`
}

// DescribeVsStorageTrafficUsageDataResponse is the response struct for api DescribeVsStorageTrafficUsageData
type DescribeVsStorageTrafficUsageDataResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TrafficUsage TrafficUsage `json:"TrafficUsage" xml:"TrafficUsage"`
}

// CreateDescribeVsStorageTrafficUsageDataRequest creates a request to invoke DescribeVsStorageTrafficUsageData API
func CreateDescribeVsStorageTrafficUsageDataRequest() (request *DescribeVsStorageTrafficUsageDataRequest) {
	request = &DescribeVsStorageTrafficUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsStorageTrafficUsageData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsStorageTrafficUsageDataResponse creates a response to parse from DescribeVsStorageTrafficUsageData response
func CreateDescribeVsStorageTrafficUsageDataResponse() (response *DescribeVsStorageTrafficUsageDataResponse) {
	response = &DescribeVsStorageTrafficUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}