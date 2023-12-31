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

// ReportFpShotJobResult invokes the mts.ReportFpShotJobResult API synchronously
func (client *Client) ReportFpShotJobResult(request *ReportFpShotJobResultRequest) (response *ReportFpShotJobResultResponse, err error) {
	response = CreateReportFpShotJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportFpShotJobResultWithChan invokes the mts.ReportFpShotJobResult API asynchronously
func (client *Client) ReportFpShotJobResultWithChan(request *ReportFpShotJobResultRequest) (<-chan *ReportFpShotJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportFpShotJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportFpShotJobResult(request)
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

// ReportFpShotJobResultWithCallback invokes the mts.ReportFpShotJobResult API asynchronously
func (client *Client) ReportFpShotJobResultWithCallback(request *ReportFpShotJobResultRequest, callback func(response *ReportFpShotJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportFpShotJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportFpShotJobResult(request)
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

// ReportFpShotJobResultRequest is the request struct for api ReportFpShotJobResult
type ReportFpShotJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Result               string           `position:"Query" name:"Result"`
	JobId                string           `position:"Query" name:"JobId"`
	Details              string           `position:"Query" name:"Details"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReportFpShotJobResultResponse is the response struct for api ReportFpShotJobResult
type ReportFpShotJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportFpShotJobResultRequest creates a request to invoke ReportFpShotJobResult API
func CreateReportFpShotJobResultRequest() (request *ReportFpShotJobResultRequest) {
	request = &ReportFpShotJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportFpShotJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportFpShotJobResultResponse creates a response to parse from ReportFpShotJobResult response
func CreateReportFpShotJobResultResponse() (response *ReportFpShotJobResultResponse) {
	response = &ReportFpShotJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
