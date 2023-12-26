package cr_ee

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

// GetRepoTagScanSummary invokes the cr.GetRepoTagScanSummary API synchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscansummary.html
func (client *Client) GetRepoTagScanSummary(request *GetRepoTagScanSummaryRequest) (response *GetRepoTagScanSummaryResponse, err error) {
	response = CreateGetRepoTagScanSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoTagScanSummaryWithChan invokes the cr.GetRepoTagScanSummary API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscansummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagScanSummaryWithChan(request *GetRepoTagScanSummaryRequest) (<-chan *GetRepoTagScanSummaryResponse, <-chan error) {
	responseChan := make(chan *GetRepoTagScanSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoTagScanSummary(request)
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

// GetRepoTagScanSummaryWithCallback invokes the cr.GetRepoTagScanSummary API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscansummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagScanSummaryWithCallback(request *GetRepoTagScanSummaryRequest, callback func(response *GetRepoTagScanSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoTagScanSummaryResponse
		var err error
		defer close(result)
		response, err = client.GetRepoTagScanSummary(request)
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

// GetRepoTagScanSummaryRequest is the request struct for api GetRepoTagScanSummary
type GetRepoTagScanSummaryRequest struct {
	*requests.RpcRequest
	RepoId     string `position:"Query" name:"RepoId"`
	ScanTaskId string `position:"Query" name:"ScanTaskId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Tag        string `position:"Query" name:"Tag"`
}

// GetRepoTagScanSummaryResponse is the response struct for api GetRepoTagScanSummary
type GetRepoTagScanSummaryResponse struct {
	*responses.BaseResponse
	GetRepoTagScanSummaryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string `json:"Code" xml:"Code"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	LowSeverity                    int    `json:"LowSeverity" xml:"LowSeverity"`
	MediumSeverity                 int    `json:"MediumSeverity" xml:"MediumSeverity"`
	HighSeverity                   int    `json:"HighSeverity" xml:"HighSeverity"`
	UnknownSeverity                int    `json:"UnknownSeverity" xml:"UnknownSeverity"`
	TotalSeverity                  int    `json:"TotalSeverity" xml:"TotalSeverity"`
}

// CreateGetRepoTagScanSummaryRequest creates a request to invoke GetRepoTagScanSummary API
func CreateGetRepoTagScanSummaryRequest() (request *GetRepoTagScanSummaryRequest) {
	request = &GetRepoTagScanSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetRepoTagScanSummary", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRepoTagScanSummaryResponse creates a response to parse from GetRepoTagScanSummary response
func CreateGetRepoTagScanSummaryResponse() (response *GetRepoTagScanSummaryResponse) {
	response = &GetRepoTagScanSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
