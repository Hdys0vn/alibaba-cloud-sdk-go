package nas

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

// RetryLifecycleRetrieveJob invokes the nas.RetryLifecycleRetrieveJob API synchronously
func (client *Client) RetryLifecycleRetrieveJob(request *RetryLifecycleRetrieveJobRequest) (response *RetryLifecycleRetrieveJobResponse, err error) {
	response = CreateRetryLifecycleRetrieveJobResponse()
	err = client.DoAction(request, response)
	return
}

// RetryLifecycleRetrieveJobWithChan invokes the nas.RetryLifecycleRetrieveJob API asynchronously
func (client *Client) RetryLifecycleRetrieveJobWithChan(request *RetryLifecycleRetrieveJobRequest) (<-chan *RetryLifecycleRetrieveJobResponse, <-chan error) {
	responseChan := make(chan *RetryLifecycleRetrieveJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetryLifecycleRetrieveJob(request)
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

// RetryLifecycleRetrieveJobWithCallback invokes the nas.RetryLifecycleRetrieveJob API asynchronously
func (client *Client) RetryLifecycleRetrieveJobWithCallback(request *RetryLifecycleRetrieveJobRequest, callback func(response *RetryLifecycleRetrieveJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetryLifecycleRetrieveJobResponse
		var err error
		defer close(result)
		response, err = client.RetryLifecycleRetrieveJob(request)
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

// RetryLifecycleRetrieveJobRequest is the request struct for api RetryLifecycleRetrieveJob
type RetryLifecycleRetrieveJobRequest struct {
	*requests.RpcRequest
	JobId string `position:"Query" name:"JobId"`
}

// RetryLifecycleRetrieveJobResponse is the response struct for api RetryLifecycleRetrieveJob
type RetryLifecycleRetrieveJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRetryLifecycleRetrieveJobRequest creates a request to invoke RetryLifecycleRetrieveJob API
func CreateRetryLifecycleRetrieveJobRequest() (request *RetryLifecycleRetrieveJobRequest) {
	request = &RetryLifecycleRetrieveJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "RetryLifecycleRetrieveJob", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetryLifecycleRetrieveJobResponse creates a response to parse from RetryLifecycleRetrieveJob response
func CreateRetryLifecycleRetrieveJobResponse() (response *RetryLifecycleRetrieveJobResponse) {
	response = &RetryLifecycleRetrieveJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
