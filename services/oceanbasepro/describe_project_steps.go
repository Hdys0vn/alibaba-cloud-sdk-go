package oceanbasepro

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

// DescribeProjectSteps invokes the oceanbasepro.DescribeProjectSteps API synchronously
func (client *Client) DescribeProjectSteps(request *DescribeProjectStepsRequest) (response *DescribeProjectStepsResponse, err error) {
	response = CreateDescribeProjectStepsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectStepsWithChan invokes the oceanbasepro.DescribeProjectSteps API asynchronously
func (client *Client) DescribeProjectStepsWithChan(request *DescribeProjectStepsRequest) (<-chan *DescribeProjectStepsResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectStepsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjectSteps(request)
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

// DescribeProjectStepsWithCallback invokes the oceanbasepro.DescribeProjectSteps API asynchronously
func (client *Client) DescribeProjectStepsWithCallback(request *DescribeProjectStepsRequest, callback func(response *DescribeProjectStepsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectStepsResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjectSteps(request)
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

// DescribeProjectStepsRequest is the request struct for api DescribeProjectSteps
type DescribeProjectStepsRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// DescribeProjectStepsResponse is the response struct for api DescribeProjectSteps
type DescribeProjectStepsResponse struct {
	*responses.BaseResponse
}

// CreateDescribeProjectStepsRequest creates a request to invoke DescribeProjectSteps API
func CreateDescribeProjectStepsRequest() (request *DescribeProjectStepsRequest) {
	request = &DescribeProjectStepsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeProjectSteps", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectStepsResponse creates a response to parse from DescribeProjectSteps response
func CreateDescribeProjectStepsResponse() (response *DescribeProjectStepsResponse) {
	response = &DescribeProjectStepsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
