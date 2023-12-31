package config

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

// GenerateCompliancePackReport invokes the config.GenerateCompliancePackReport API synchronously
func (client *Client) GenerateCompliancePackReport(request *GenerateCompliancePackReportRequest) (response *GenerateCompliancePackReportResponse, err error) {
	response = CreateGenerateCompliancePackReportResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateCompliancePackReportWithChan invokes the config.GenerateCompliancePackReport API asynchronously
func (client *Client) GenerateCompliancePackReportWithChan(request *GenerateCompliancePackReportRequest) (<-chan *GenerateCompliancePackReportResponse, <-chan error) {
	responseChan := make(chan *GenerateCompliancePackReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateCompliancePackReport(request)
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

// GenerateCompliancePackReportWithCallback invokes the config.GenerateCompliancePackReport API asynchronously
func (client *Client) GenerateCompliancePackReportWithCallback(request *GenerateCompliancePackReportRequest, callback func(response *GenerateCompliancePackReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateCompliancePackReportResponse
		var err error
		defer close(result)
		response, err = client.GenerateCompliancePackReport(request)
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

// GenerateCompliancePackReportRequest is the request struct for api GenerateCompliancePackReport
type GenerateCompliancePackReportRequest struct {
	*requests.RpcRequest
	ClientToken      string `position:"Body" name:"ClientToken"`
	CompliancePackId string `position:"Body" name:"CompliancePackId"`
}

// GenerateCompliancePackReportResponse is the response struct for api GenerateCompliancePackReport
type GenerateCompliancePackReportResponse struct {
	*responses.BaseResponse
	CompliancePackId string `json:"CompliancePackId" xml:"CompliancePackId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateGenerateCompliancePackReportRequest creates a request to invoke GenerateCompliancePackReport API
func CreateGenerateCompliancePackReportRequest() (request *GenerateCompliancePackReportRequest) {
	request = &GenerateCompliancePackReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GenerateCompliancePackReport", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateCompliancePackReportResponse creates a response to parse from GenerateCompliancePackReport response
func CreateGenerateCompliancePackReportResponse() (response *GenerateCompliancePackReportResponse) {
	response = &GenerateCompliancePackReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
