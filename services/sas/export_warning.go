package sas

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

// ExportWarning invokes the sas.ExportWarning API synchronously
func (client *Client) ExportWarning(request *ExportWarningRequest) (response *ExportWarningResponse, err error) {
	response = CreateExportWarningResponse()
	err = client.DoAction(request, response)
	return
}

// ExportWarningWithChan invokes the sas.ExportWarning API asynchronously
func (client *Client) ExportWarningWithChan(request *ExportWarningRequest) (<-chan *ExportWarningResponse, <-chan error) {
	responseChan := make(chan *ExportWarningResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportWarning(request)
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

// ExportWarningWithCallback invokes the sas.ExportWarning API asynchronously
func (client *Client) ExportWarningWithCallback(request *ExportWarningRequest, callback func(response *ExportWarningResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportWarningResponse
		var err error
		defer close(result)
		response, err = client.ExportWarning(request)
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

// ExportWarningRequest is the request struct for api ExportWarning
type ExportWarningRequest struct {
	*requests.RpcRequest
	IsCleartextPwd  requests.Integer `position:"Query" name:"IsCleartextPwd"`
	StatusList      string           `position:"Query" name:"StatusList"`
	RiskLevels      string           `position:"Query" name:"RiskLevels"`
	RiskName        string           `position:"Query" name:"RiskName"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	ExportType      string           `position:"Query" name:"ExportType"`
	Dealed          string           `position:"Query" name:"Dealed"`
	TypeNames       string           `position:"Query" name:"TypeNames"`
	IsSummaryExport requests.Integer `position:"Query" name:"IsSummaryExport"`
	RiskIds         string           `position:"Query" name:"RiskIds"`
	StrategyId      requests.Integer `position:"Query" name:"StrategyId"`
	TypeName        string           `position:"Query" name:"TypeName"`
	SubTypeNames    string           `position:"Query" name:"SubTypeNames"`
	Uuids           string           `position:"Query" name:"Uuids"`
}

// ExportWarningResponse is the response struct for api ExportWarning
type ExportWarningResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FileName  string `json:"FileName" xml:"FileName"`
	Id        int64  `json:"Id" xml:"Id"`
}

// CreateExportWarningRequest creates a request to invoke ExportWarning API
func CreateExportWarningRequest() (request *ExportWarningRequest) {
	request = &ExportWarningRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ExportWarning", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportWarningResponse creates a response to parse from ExportWarning response
func CreateExportWarningResponse() (response *ExportWarningResponse) {
	response = &ExportWarningResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}