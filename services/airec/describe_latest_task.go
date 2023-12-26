package airec

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

// DescribeLatestTask invokes the airec.DescribeLatestTask API synchronously
func (client *Client) DescribeLatestTask(request *DescribeLatestTaskRequest) (response *DescribeLatestTaskResponse, err error) {
	response = CreateDescribeLatestTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLatestTaskWithChan invokes the airec.DescribeLatestTask API asynchronously
func (client *Client) DescribeLatestTaskWithChan(request *DescribeLatestTaskRequest) (<-chan *DescribeLatestTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeLatestTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLatestTask(request)
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

// DescribeLatestTaskWithCallback invokes the airec.DescribeLatestTask API asynchronously
func (client *Client) DescribeLatestTaskWithCallback(request *DescribeLatestTaskRequest, callback func(response *DescribeLatestTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLatestTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeLatestTask(request)
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

// DescribeLatestTaskRequest is the request struct for api DescribeLatestTask
type DescribeLatestTaskRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"instanceId"`
	AlgorithmId string `position:"Path" name:"algorithmId"`
}

// DescribeLatestTaskResponse is the response struct for api DescribeLatestTask
type DescribeLatestTaskResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"requestId" xml:"requestId"`
	Result    []IndexVersion `json:"result" xml:"result"`
}

// CreateDescribeLatestTaskRequest creates a request to invoke DescribeLatestTask API
func CreateDescribeLatestTaskRequest() (request *DescribeLatestTaskRequest) {
	request = &DescribeLatestTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeLatestTask", "/v2/openapi/instances/[instanceId]/filtering-algorithms/[algorithmId]/tasks/latest", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeLatestTaskResponse creates a response to parse from DescribeLatestTask response
func CreateDescribeLatestTaskResponse() (response *DescribeLatestTaskResponse) {
	response = &DescribeLatestTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
