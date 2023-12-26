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

// GetEdgeTranscodeJob invokes the live.GetEdgeTranscodeJob API synchronously
func (client *Client) GetEdgeTranscodeJob(request *GetEdgeTranscodeJobRequest) (response *GetEdgeTranscodeJobResponse, err error) {
	response = CreateGetEdgeTranscodeJobResponse()
	err = client.DoAction(request, response)
	return
}

// GetEdgeTranscodeJobWithChan invokes the live.GetEdgeTranscodeJob API asynchronously
func (client *Client) GetEdgeTranscodeJobWithChan(request *GetEdgeTranscodeJobRequest) (<-chan *GetEdgeTranscodeJobResponse, <-chan error) {
	responseChan := make(chan *GetEdgeTranscodeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEdgeTranscodeJob(request)
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

// GetEdgeTranscodeJobWithCallback invokes the live.GetEdgeTranscodeJob API asynchronously
func (client *Client) GetEdgeTranscodeJobWithCallback(request *GetEdgeTranscodeJobRequest, callback func(response *GetEdgeTranscodeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEdgeTranscodeJobResponse
		var err error
		defer close(result)
		response, err = client.GetEdgeTranscodeJob(request)
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

// GetEdgeTranscodeJobRequest is the request struct for api GetEdgeTranscodeJob
type GetEdgeTranscodeJobRequest struct {
	*requests.RpcRequest
	JobId     string           `position:"Query" name:"JobId"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// GetEdgeTranscodeJobResponse is the response struct for api GetEdgeTranscodeJob
type GetEdgeTranscodeJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Job       Job    `json:"Job" xml:"Job"`
}

// CreateGetEdgeTranscodeJobRequest creates a request to invoke GetEdgeTranscodeJob API
func CreateGetEdgeTranscodeJobRequest() (request *GetEdgeTranscodeJobRequest) {
	request = &GetEdgeTranscodeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetEdgeTranscodeJob", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEdgeTranscodeJobResponse creates a response to parse from GetEdgeTranscodeJob response
func CreateGetEdgeTranscodeJobResponse() (response *GetEdgeTranscodeJobResponse) {
	response = &GetEdgeTranscodeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
