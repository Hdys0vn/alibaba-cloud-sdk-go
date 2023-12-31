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

// ReportCensorJobResult invokes the mts.ReportCensorJobResult API synchronously
func (client *Client) ReportCensorJobResult(request *ReportCensorJobResultRequest) (response *ReportCensorJobResultResponse, err error) {
	response = CreateReportCensorJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportCensorJobResultWithChan invokes the mts.ReportCensorJobResult API asynchronously
func (client *Client) ReportCensorJobResultWithChan(request *ReportCensorJobResultRequest) (<-chan *ReportCensorJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportCensorJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportCensorJobResult(request)
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

// ReportCensorJobResultWithCallback invokes the mts.ReportCensorJobResult API asynchronously
func (client *Client) ReportCensorJobResultWithCallback(request *ReportCensorJobResultRequest, callback func(response *ReportCensorJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportCensorJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportCensorJobResult(request)
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

// ReportCensorJobResultRequest is the request struct for api ReportCensorJobResult
type ReportCensorJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Label                string           `position:"Query" name:"Label"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Detail               string           `position:"Query" name:"Detail"`
}

// ReportCensorJobResultResponse is the response struct for api ReportCensorJobResult
type ReportCensorJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportCensorJobResultRequest creates a request to invoke ReportCensorJobResult API
func CreateReportCensorJobResultRequest() (request *ReportCensorJobResultRequest) {
	request = &ReportCensorJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportCensorJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportCensorJobResultResponse creates a response to parse from ReportCensorJobResult response
func CreateReportCensorJobResultResponse() (response *ReportCensorJobResultResponse) {
	response = &ReportCensorJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
