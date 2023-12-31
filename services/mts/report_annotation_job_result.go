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

// ReportAnnotationJobResult invokes the mts.ReportAnnotationJobResult API synchronously
func (client *Client) ReportAnnotationJobResult(request *ReportAnnotationJobResultRequest) (response *ReportAnnotationJobResultResponse, err error) {
	response = CreateReportAnnotationJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportAnnotationJobResultWithChan invokes the mts.ReportAnnotationJobResult API asynchronously
func (client *Client) ReportAnnotationJobResultWithChan(request *ReportAnnotationJobResultRequest) (<-chan *ReportAnnotationJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportAnnotationJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportAnnotationJobResult(request)
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

// ReportAnnotationJobResultWithCallback invokes the mts.ReportAnnotationJobResult API asynchronously
func (client *Client) ReportAnnotationJobResultWithCallback(request *ReportAnnotationJobResultRequest, callback func(response *ReportAnnotationJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportAnnotationJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportAnnotationJobResult(request)
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

// ReportAnnotationJobResultRequest is the request struct for api ReportAnnotationJobResult
type ReportAnnotationJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	Details              string           `position:"Query" name:"Details"`
	Annotation           string           `position:"Query" name:"Annotation"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReportAnnotationJobResultResponse is the response struct for api ReportAnnotationJobResult
type ReportAnnotationJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportAnnotationJobResultRequest creates a request to invoke ReportAnnotationJobResult API
func CreateReportAnnotationJobResultRequest() (request *ReportAnnotationJobResultRequest) {
	request = &ReportAnnotationJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportAnnotationJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportAnnotationJobResultResponse creates a response to parse from ReportAnnotationJobResult response
func CreateReportAnnotationJobResultResponse() (response *ReportAnnotationJobResultResponse) {
	response = &ReportAnnotationJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
