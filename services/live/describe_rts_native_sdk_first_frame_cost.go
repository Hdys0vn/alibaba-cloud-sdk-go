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

// DescribeRTSNativeSDKFirstFrameCost invokes the live.DescribeRTSNativeSDKFirstFrameCost API synchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameCost(request *DescribeRTSNativeSDKFirstFrameCostRequest) (response *DescribeRTSNativeSDKFirstFrameCostResponse, err error) {
	response = CreateDescribeRTSNativeSDKFirstFrameCostResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRTSNativeSDKFirstFrameCostWithChan invokes the live.DescribeRTSNativeSDKFirstFrameCost API asynchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameCostWithChan(request *DescribeRTSNativeSDKFirstFrameCostRequest) (<-chan *DescribeRTSNativeSDKFirstFrameCostResponse, <-chan error) {
	responseChan := make(chan *DescribeRTSNativeSDKFirstFrameCostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRTSNativeSDKFirstFrameCost(request)
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

// DescribeRTSNativeSDKFirstFrameCostWithCallback invokes the live.DescribeRTSNativeSDKFirstFrameCost API asynchronously
func (client *Client) DescribeRTSNativeSDKFirstFrameCostWithCallback(request *DescribeRTSNativeSDKFirstFrameCostRequest, callback func(response *DescribeRTSNativeSDKFirstFrameCostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRTSNativeSDKFirstFrameCostResponse
		var err error
		defer close(result)
		response, err = client.DescribeRTSNativeSDKFirstFrameCost(request)
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

// DescribeRTSNativeSDKFirstFrameCostRequest is the request struct for api DescribeRTSNativeSDKFirstFrameCost
type DescribeRTSNativeSDKFirstFrameCostRequest struct {
	*requests.RpcRequest
	EndTime        string    `position:"Query" name:"EndTime"`
	DomainNameList *[]string `position:"Query" name:"DomainNameList"  type:"Json"`
	StartTime      string    `position:"Query" name:"StartTime"`
	DataInterval   string    `position:"Query" name:"DataInterval"`
}

// DescribeRTSNativeSDKFirstFrameCostResponse is the response struct for api DescribeRTSNativeSDKFirstFrameCost
type DescribeRTSNativeSDKFirstFrameCostResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	DataInterval       string `json:"DataInterval" xml:"DataInterval"`
	StartTime          string `json:"StartTime" xml:"StartTime"`
	EndTime            string `json:"EndTime" xml:"EndTime"`
	FirstFrameCostData []Data `json:"FirstFrameCostData" xml:"FirstFrameCostData"`
}

// CreateDescribeRTSNativeSDKFirstFrameCostRequest creates a request to invoke DescribeRTSNativeSDKFirstFrameCost API
func CreateDescribeRTSNativeSDKFirstFrameCostRequest() (request *DescribeRTSNativeSDKFirstFrameCostRequest) {
	request = &DescribeRTSNativeSDKFirstFrameCostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRTSNativeSDKFirstFrameCost", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRTSNativeSDKFirstFrameCostResponse creates a response to parse from DescribeRTSNativeSDKFirstFrameCost response
func CreateDescribeRTSNativeSDKFirstFrameCostResponse() (response *DescribeRTSNativeSDKFirstFrameCostResponse) {
	response = &DescribeRTSNativeSDKFirstFrameCostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}