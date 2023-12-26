package sae

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

// DescribeInstanceLog invokes the sae.DescribeInstanceLog API synchronously
func (client *Client) DescribeInstanceLog(request *DescribeInstanceLogRequest) (response *DescribeInstanceLogResponse, err error) {
	response = CreateDescribeInstanceLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceLogWithChan invokes the sae.DescribeInstanceLog API asynchronously
func (client *Client) DescribeInstanceLogWithChan(request *DescribeInstanceLogRequest) (<-chan *DescribeInstanceLogResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceLog(request)
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

// DescribeInstanceLogWithCallback invokes the sae.DescribeInstanceLog API asynchronously
func (client *Client) DescribeInstanceLogWithCallback(request *DescribeInstanceLogRequest, callback func(response *DescribeInstanceLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceLog(request)
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

// DescribeInstanceLogRequest is the request struct for api DescribeInstanceLog
type DescribeInstanceLogRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeInstanceLogResponse is the response struct for api DescribeInstanceLog
type DescribeInstanceLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDescribeInstanceLogRequest creates a request to invoke DescribeInstanceLog API
func CreateDescribeInstanceLogRequest() (request *DescribeInstanceLogRequest) {
	request = &DescribeInstanceLogRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeInstanceLog", "/pop/v1/sam/instance/describeInstanceLog", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeInstanceLogResponse creates a response to parse from DescribeInstanceLog response
func CreateDescribeInstanceLogResponse() (response *DescribeInstanceLogResponse) {
	response = &DescribeInstanceLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
