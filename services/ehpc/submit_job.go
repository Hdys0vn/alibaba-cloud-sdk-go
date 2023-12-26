package ehpc

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

// SubmitJob invokes the ehpc.SubmitJob API synchronously
func (client *Client) SubmitJob(request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
	response = CreateSubmitJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitJobWithChan invokes the ehpc.SubmitJob API asynchronously
func (client *Client) SubmitJobWithChan(request *SubmitJobRequest) (<-chan *SubmitJobResponse, <-chan error) {
	responseChan := make(chan *SubmitJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitJob(request)
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

// SubmitJobWithCallback invokes the ehpc.SubmitJob API asynchronously
func (client *Client) SubmitJobWithCallback(request *SubmitJobRequest, callback func(response *SubmitJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitJob(request)
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

// SubmitJobRequest is the request struct for api SubmitJob
type SubmitJobRequest struct {
	*requests.RpcRequest
	StderrRedirectPath string           `position:"Query" name:"StderrRedirectPath"`
	RunasUserPassword  string           `position:"Query" name:"RunasUserPassword"`
	ClockTime          string           `position:"Query" name:"ClockTime"`
	JobRetryPriority   requests.Integer `position:"Query" name:"JobRetry.Priority"`
	CommandLine        string           `position:"Query" name:"CommandLine"`
	JobQueue           string           `position:"Query" name:"JobQueue"`
	ArrayRequest       string           `position:"Query" name:"ArrayRequest"`
	UnzipCmd           string           `position:"Query" name:"UnzipCmd"`
	PackagePath        string           `position:"Query" name:"PackagePath"`
	Mem                string           `position:"Query" name:"Mem"`
	JobRetryCount      requests.Integer `position:"Query" name:"JobRetry.Count"`
	StdoutRedirectPath string           `position:"Query" name:"StdoutRedirectPath"`
	Variables          string           `position:"Query" name:"Variables"`
	PostCmdLine        string           `position:"Query" name:"PostCmdLine"`
	RunasUser          string           `position:"Query" name:"RunasUser"`
	Cpu                requests.Integer `position:"Query" name:"Cpu"`
	ClusterId          string           `position:"Query" name:"ClusterId"`
	ReRunable          requests.Boolean `position:"Query" name:"ReRunable"`
	Thread             requests.Integer `position:"Query" name:"Thread"`
	Priority           requests.Integer `position:"Query" name:"Priority"`
	Gpu                requests.Integer `position:"Query" name:"Gpu"`
	JobRetryOnExitCode requests.Integer `position:"Query" name:"JobRetry.OnExitCode"`
	Node               requests.Integer `position:"Query" name:"Node"`
	Async              requests.Boolean `position:"Query" name:"Async"`
	Task               requests.Integer `position:"Query" name:"Task"`
	InputFileUrl       string           `position:"Query" name:"InputFileUrl"`
	Name               string           `position:"Query" name:"Name"`
	ContainerId        string           `position:"Query" name:"ContainerId"`
}

// SubmitJobResponse is the response struct for api SubmitJob
type SubmitJobResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitJobRequest creates a request to invoke SubmitJob API
func CreateSubmitJobRequest() (request *SubmitJobRequest) {
	request = &SubmitJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SubmitJob", "", "")
	request.Method = requests.GET
	return
}

// CreateSubmitJobResponse creates a response to parse from SubmitJob response
func CreateSubmitJobResponse() (response *SubmitJobResponse) {
	response = &SubmitJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
