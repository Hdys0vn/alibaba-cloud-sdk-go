package drds

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

// DescribeDrdsInstanceMonitor invokes the drds.DescribeDrdsInstanceMonitor API synchronously
func (client *Client) DescribeDrdsInstanceMonitor(request *DescribeDrdsInstanceMonitorRequest) (response *DescribeDrdsInstanceMonitorResponse, err error) {
	response = CreateDescribeDrdsInstanceMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsInstanceMonitorWithChan invokes the drds.DescribeDrdsInstanceMonitor API asynchronously
func (client *Client) DescribeDrdsInstanceMonitorWithChan(request *DescribeDrdsInstanceMonitorRequest) (<-chan *DescribeDrdsInstanceMonitorResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsInstanceMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsInstanceMonitor(request)
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

// DescribeDrdsInstanceMonitorWithCallback invokes the drds.DescribeDrdsInstanceMonitor API asynchronously
func (client *Client) DescribeDrdsInstanceMonitorWithCallback(request *DescribeDrdsInstanceMonitorRequest, callback func(response *DescribeDrdsInstanceMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsInstanceMonitorResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsInstanceMonitor(request)
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

// DescribeDrdsInstanceMonitorRequest is the request struct for api DescribeDrdsInstanceMonitor
type DescribeDrdsInstanceMonitorRequest struct {
	*requests.RpcRequest
	EndTime        requests.Integer `position:"Query" name:"EndTime"`
	StartTime      requests.Integer `position:"Query" name:"StartTime"`
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	Key            string           `position:"Query" name:"Key"`
	PeriodMultiple requests.Integer `position:"Query" name:"PeriodMultiple"`
}

// DescribeDrdsInstanceMonitorResponse is the response struct for api DescribeDrdsInstanceMonitor
type DescribeDrdsInstanceMonitorResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Data      []PartialPerformanceData `json:"Data" xml:"Data"`
}

// CreateDescribeDrdsInstanceMonitorRequest creates a request to invoke DescribeDrdsInstanceMonitor API
func CreateDescribeDrdsInstanceMonitorRequest() (request *DescribeDrdsInstanceMonitorRequest) {
	request = &DescribeDrdsInstanceMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsInstanceMonitor", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsInstanceMonitorResponse creates a response to parse from DescribeDrdsInstanceMonitor response
func CreateDescribeDrdsInstanceMonitorResponse() (response *DescribeDrdsInstanceMonitorResponse) {
	response = &DescribeDrdsInstanceMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
