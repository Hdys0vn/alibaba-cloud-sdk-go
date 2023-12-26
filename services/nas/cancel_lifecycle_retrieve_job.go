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

// CancelLifecycleRetrieveJob invokes the nas.CancelLifecycleRetrieveJob API synchronously
func (client *Client) CancelLifecycleRetrieveJob(request *CancelLifecycleRetrieveJobRequest) (response *CancelLifecycleRetrieveJobResponse, err error) {
	response = CreateCancelLifecycleRetrieveJobResponse()
	err = client.DoAction(request, response)
	return
}

// CancelLifecycleRetrieveJobWithChan invokes the nas.CancelLifecycleRetrieveJob API asynchronously
func (client *Client) CancelLifecycleRetrieveJobWithChan(request *CancelLifecycleRetrieveJobRequest) (<-chan *CancelLifecycleRetrieveJobResponse, <-chan error) {
	responseChan := make(chan *CancelLifecycleRetrieveJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelLifecycleRetrieveJob(request)
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

// CancelLifecycleRetrieveJobWithCallback invokes the nas.CancelLifecycleRetrieveJob API asynchronously
func (client *Client) CancelLifecycleRetrieveJobWithCallback(request *CancelLifecycleRetrieveJobRequest, callback func(response *CancelLifecycleRetrieveJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelLifecycleRetrieveJobResponse
		var err error
		defer close(result)
		response, err = client.CancelLifecycleRetrieveJob(request)
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

// CancelLifecycleRetrieveJobRequest is the request struct for api CancelLifecycleRetrieveJob
type CancelLifecycleRetrieveJobRequest struct {
	*requests.RpcRequest
	JobId string `position:"Query" name:"JobId"`
}

// CancelLifecycleRetrieveJobResponse is the response struct for api CancelLifecycleRetrieveJob
type CancelLifecycleRetrieveJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelLifecycleRetrieveJobRequest creates a request to invoke CancelLifecycleRetrieveJob API
func CreateCancelLifecycleRetrieveJobRequest() (request *CancelLifecycleRetrieveJobRequest) {
	request = &CancelLifecycleRetrieveJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CancelLifecycleRetrieveJob", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelLifecycleRetrieveJobResponse creates a response to parse from CancelLifecycleRetrieveJob response
func CreateCancelLifecycleRetrieveJobResponse() (response *CancelLifecycleRetrieveJobResponse) {
	response = &CancelLifecycleRetrieveJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
