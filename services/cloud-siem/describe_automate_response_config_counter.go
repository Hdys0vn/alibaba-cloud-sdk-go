package cloud_siem

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

// DescribeAutomateResponseConfigCounter invokes the cloud_siem.DescribeAutomateResponseConfigCounter API synchronously
func (client *Client) DescribeAutomateResponseConfigCounter(request *DescribeAutomateResponseConfigCounterRequest) (response *DescribeAutomateResponseConfigCounterResponse, err error) {
	response = CreateDescribeAutomateResponseConfigCounterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutomateResponseConfigCounterWithChan invokes the cloud_siem.DescribeAutomateResponseConfigCounter API asynchronously
func (client *Client) DescribeAutomateResponseConfigCounterWithChan(request *DescribeAutomateResponseConfigCounterRequest) (<-chan *DescribeAutomateResponseConfigCounterResponse, <-chan error) {
	responseChan := make(chan *DescribeAutomateResponseConfigCounterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutomateResponseConfigCounter(request)
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

// DescribeAutomateResponseConfigCounterWithCallback invokes the cloud_siem.DescribeAutomateResponseConfigCounter API asynchronously
func (client *Client) DescribeAutomateResponseConfigCounterWithCallback(request *DescribeAutomateResponseConfigCounterRequest, callback func(response *DescribeAutomateResponseConfigCounterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutomateResponseConfigCounterResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutomateResponseConfigCounter(request)
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

// DescribeAutomateResponseConfigCounterRequest is the request struct for api DescribeAutomateResponseConfigCounter
type DescribeAutomateResponseConfigCounterRequest struct {
	*requests.RpcRequest
}

// DescribeAutomateResponseConfigCounterResponse is the response struct for api DescribeAutomateResponseConfigCounter
type DescribeAutomateResponseConfigCounterResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAutomateResponseConfigCounterRequest creates a request to invoke DescribeAutomateResponseConfigCounter API
func CreateDescribeAutomateResponseConfigCounterRequest() (request *DescribeAutomateResponseConfigCounterRequest) {
	request = &DescribeAutomateResponseConfigCounterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAutomateResponseConfigCounter", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutomateResponseConfigCounterResponse creates a response to parse from DescribeAutomateResponseConfigCounter response
func CreateDescribeAutomateResponseConfigCounterResponse() (response *DescribeAutomateResponseConfigCounterResponse) {
	response = &DescribeAutomateResponseConfigCounterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
