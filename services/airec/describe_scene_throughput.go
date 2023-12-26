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

// DescribeSceneThroughput invokes the airec.DescribeSceneThroughput API synchronously
func (client *Client) DescribeSceneThroughput(request *DescribeSceneThroughputRequest) (response *DescribeSceneThroughputResponse, err error) {
	response = CreateDescribeSceneThroughputResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSceneThroughputWithChan invokes the airec.DescribeSceneThroughput API asynchronously
func (client *Client) DescribeSceneThroughputWithChan(request *DescribeSceneThroughputRequest) (<-chan *DescribeSceneThroughputResponse, <-chan error) {
	responseChan := make(chan *DescribeSceneThroughputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSceneThroughput(request)
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

// DescribeSceneThroughputWithCallback invokes the airec.DescribeSceneThroughput API asynchronously
func (client *Client) DescribeSceneThroughputWithCallback(request *DescribeSceneThroughputRequest, callback func(response *DescribeSceneThroughputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSceneThroughputResponse
		var err error
		defer close(result)
		response, err = client.DescribeSceneThroughput(request)
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

// DescribeSceneThroughputRequest is the request struct for api DescribeSceneThroughput
type DescribeSceneThroughputRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	SceneId    string `position:"Path" name:"sceneId"`
}

// DescribeSceneThroughputResponse is the response struct for api DescribeSceneThroughput
type DescribeSceneThroughputResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeSceneThroughputRequest creates a request to invoke DescribeSceneThroughput API
func CreateDescribeSceneThroughputRequest() (request *DescribeSceneThroughputRequest) {
	request = &DescribeSceneThroughputRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeSceneThroughput", "/v2/openapi/instances/[instanceId]/scenes/[sceneId]/throughput", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSceneThroughputResponse creates a response to parse from DescribeSceneThroughput response
func CreateDescribeSceneThroughputResponse() (response *DescribeSceneThroughputResponse) {
	response = &DescribeSceneThroughputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
