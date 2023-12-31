package das

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

// DescribeCacheAnalysisJob invokes the das.DescribeCacheAnalysisJob API synchronously
func (client *Client) DescribeCacheAnalysisJob(request *DescribeCacheAnalysisJobRequest) (response *DescribeCacheAnalysisJobResponse, err error) {
	response = CreateDescribeCacheAnalysisJobResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCacheAnalysisJobWithChan invokes the das.DescribeCacheAnalysisJob API asynchronously
func (client *Client) DescribeCacheAnalysisJobWithChan(request *DescribeCacheAnalysisJobRequest) (<-chan *DescribeCacheAnalysisJobResponse, <-chan error) {
	responseChan := make(chan *DescribeCacheAnalysisJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCacheAnalysisJob(request)
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

// DescribeCacheAnalysisJobWithCallback invokes the das.DescribeCacheAnalysisJob API asynchronously
func (client *Client) DescribeCacheAnalysisJobWithCallback(request *DescribeCacheAnalysisJobRequest, callback func(response *DescribeCacheAnalysisJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCacheAnalysisJobResponse
		var err error
		defer close(result)
		response, err = client.DescribeCacheAnalysisJob(request)
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

// DescribeCacheAnalysisJobRequest is the request struct for api DescribeCacheAnalysisJob
type DescribeCacheAnalysisJobRequest struct {
	*requests.RpcRequest
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeCacheAnalysisJobResponse is the response struct for api DescribeCacheAnalysisJob
type DescribeCacheAnalysisJobResponse struct {
	*responses.BaseResponse
	Message   string                         `json:"Message" xml:"Message"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Code      string                         `json:"Code" xml:"Code"`
	Success   string                         `json:"Success" xml:"Success"`
	Data      DataInDescribeCacheAnalysisJob `json:"Data" xml:"Data"`
}

// CreateDescribeCacheAnalysisJobRequest creates a request to invoke DescribeCacheAnalysisJob API
func CreateDescribeCacheAnalysisJobRequest() (request *DescribeCacheAnalysisJobRequest) {
	request = &DescribeCacheAnalysisJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DescribeCacheAnalysisJob", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCacheAnalysisJobResponse creates a response to parse from DescribeCacheAnalysisJob response
func CreateDescribeCacheAnalysisJobResponse() (response *DescribeCacheAnalysisJobResponse) {
	response = &DescribeCacheAnalysisJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
