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

// DescribeHybridMonitorDataList invokes the cms.DescribeHybridMonitorDataList API synchronously
func (client *Client) DescribeHybridMonitorDataList(request *DescribeHybridMonitorDataListRequest) (response *DescribeHybridMonitorDataListResponse, err error) {
	response = CreateDescribeHybridMonitorDataListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHybridMonitorDataListWithChan invokes the cms.DescribeHybridMonitorDataList API asynchronously
func (client *Client) DescribeHybridMonitorDataListWithChan(request *DescribeHybridMonitorDataListRequest) (<-chan *DescribeHybridMonitorDataListResponse, <-chan error) {
	responseChan := make(chan *DescribeHybridMonitorDataListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHybridMonitorDataList(request)
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

// DescribeHybridMonitorDataListWithCallback invokes the cms.DescribeHybridMonitorDataList API asynchronously
func (client *Client) DescribeHybridMonitorDataListWithCallback(request *DescribeHybridMonitorDataListRequest, callback func(response *DescribeHybridMonitorDataListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHybridMonitorDataListResponse
		var err error
		defer close(result)
		response, err = client.DescribeHybridMonitorDataList(request)
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

// DescribeHybridMonitorDataListRequest is the request struct for api DescribeHybridMonitorDataList
type DescribeHybridMonitorDataListRequest struct {
	*requests.RpcRequest
	Period    string           `position:"Query" name:"Period"`
	Start     requests.Integer `position:"Query" name:"Start"`
	Namespace string           `position:"Query" name:"Namespace"`
	End       requests.Integer `position:"Query" name:"End"`
	PromSQL   string           `position:"Query" name:"PromSQL"`
}

// DescribeHybridMonitorDataListResponse is the response struct for api DescribeHybridMonitorDataList
type DescribeHybridMonitorDataListResponse struct {
	*responses.BaseResponse
	Code       string           `json:"Code" xml:"Code"`
	Message    string           `json:"Message" xml:"Message"`
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Success    string           `json:"Success" xml:"Success"`
	TimeSeries []TimeSeriesItem `json:"TimeSeries" xml:"TimeSeries"`
}

// CreateDescribeHybridMonitorDataListRequest creates a request to invoke DescribeHybridMonitorDataList API
func CreateDescribeHybridMonitorDataListRequest() (request *DescribeHybridMonitorDataListRequest) {
	request = &DescribeHybridMonitorDataListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeHybridMonitorDataList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHybridMonitorDataListResponse creates a response to parse from DescribeHybridMonitorDataList response
func CreateDescribeHybridMonitorDataListResponse() (response *DescribeHybridMonitorDataListResponse) {
	response = &DescribeHybridMonitorDataListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
