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

// StartEdgeTranscodeJob invokes the live.StartEdgeTranscodeJob API synchronously
func (client *Client) StartEdgeTranscodeJob(request *StartEdgeTranscodeJobRequest) (response *StartEdgeTranscodeJobResponse, err error) {
	response = CreateStartEdgeTranscodeJobResponse()
	err = client.DoAction(request, response)
	return
}

// StartEdgeTranscodeJobWithChan invokes the live.StartEdgeTranscodeJob API asynchronously
func (client *Client) StartEdgeTranscodeJobWithChan(request *StartEdgeTranscodeJobRequest) (<-chan *StartEdgeTranscodeJobResponse, <-chan error) {
	responseChan := make(chan *StartEdgeTranscodeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartEdgeTranscodeJob(request)
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

// StartEdgeTranscodeJobWithCallback invokes the live.StartEdgeTranscodeJob API asynchronously
func (client *Client) StartEdgeTranscodeJobWithCallback(request *StartEdgeTranscodeJobRequest, callback func(response *StartEdgeTranscodeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartEdgeTranscodeJobResponse
		var err error
		defer close(result)
		response, err = client.StartEdgeTranscodeJob(request)
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

// StartEdgeTranscodeJobRequest is the request struct for api StartEdgeTranscodeJob
type StartEdgeTranscodeJobRequest struct {
	*requests.RpcRequest
	JobId     string           `position:"Query" name:"JobId"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// StartEdgeTranscodeJobResponse is the response struct for api StartEdgeTranscodeJob
type StartEdgeTranscodeJobResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartEdgeTranscodeJobRequest creates a request to invoke StartEdgeTranscodeJob API
func CreateStartEdgeTranscodeJobRequest() (request *StartEdgeTranscodeJobRequest) {
	request = &StartEdgeTranscodeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StartEdgeTranscodeJob", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartEdgeTranscodeJobResponse creates a response to parse from StartEdgeTranscodeJob response
func CreateStartEdgeTranscodeJobResponse() (response *StartEdgeTranscodeJobResponse) {
	response = &StartEdgeTranscodeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
