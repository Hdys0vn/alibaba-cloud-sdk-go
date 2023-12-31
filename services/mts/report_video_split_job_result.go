package mts

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

// ReportVideoSplitJobResult invokes the mts.ReportVideoSplitJobResult API synchronously
func (client *Client) ReportVideoSplitJobResult(request *ReportVideoSplitJobResultRequest) (response *ReportVideoSplitJobResultResponse, err error) {
	response = CreateReportVideoSplitJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportVideoSplitJobResultWithChan invokes the mts.ReportVideoSplitJobResult API asynchronously
func (client *Client) ReportVideoSplitJobResultWithChan(request *ReportVideoSplitJobResultRequest) (<-chan *ReportVideoSplitJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportVideoSplitJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportVideoSplitJobResult(request)
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

// ReportVideoSplitJobResultWithCallback invokes the mts.ReportVideoSplitJobResult API asynchronously
func (client *Client) ReportVideoSplitJobResultWithCallback(request *ReportVideoSplitJobResultRequest, callback func(response *ReportVideoSplitJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportVideoSplitJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportVideoSplitJobResult(request)
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

// ReportVideoSplitJobResultRequest is the request struct for api ReportVideoSplitJobResult
type ReportVideoSplitJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Result               string           `position:"Query" name:"Result"`
	JobId                string           `position:"Query" name:"JobId"`
	Details              string           `position:"Query" name:"Details"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReportVideoSplitJobResultResponse is the response struct for api ReportVideoSplitJobResult
type ReportVideoSplitJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportVideoSplitJobResultRequest creates a request to invoke ReportVideoSplitJobResult API
func CreateReportVideoSplitJobResultRequest() (request *ReportVideoSplitJobResultRequest) {
	request = &ReportVideoSplitJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportVideoSplitJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportVideoSplitJobResultResponse creates a response to parse from ReportVideoSplitJobResult response
func CreateReportVideoSplitJobResultResponse() (response *ReportVideoSplitJobResultResponse) {
	response = &ReportVideoSplitJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
