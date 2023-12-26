package sddp

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

// ModifyReportTaskStatus invokes the sddp.ModifyReportTaskStatus API synchronously
func (client *Client) ModifyReportTaskStatus(request *ModifyReportTaskStatusRequest) (response *ModifyReportTaskStatusResponse, err error) {
	response = CreateModifyReportTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReportTaskStatusWithChan invokes the sddp.ModifyReportTaskStatus API asynchronously
func (client *Client) ModifyReportTaskStatusWithChan(request *ModifyReportTaskStatusRequest) (<-chan *ModifyReportTaskStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyReportTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReportTaskStatus(request)
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

// ModifyReportTaskStatusWithCallback invokes the sddp.ModifyReportTaskStatus API asynchronously
func (client *Client) ModifyReportTaskStatusWithCallback(request *ModifyReportTaskStatusRequest, callback func(response *ModifyReportTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReportTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyReportTaskStatus(request)
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

// ModifyReportTaskStatusRequest is the request struct for api ModifyReportTaskStatus
type ModifyReportTaskStatusRequest struct {
	*requests.RpcRequest
	SourceIp         string           `position:"Query" name:"SourceIp"`
	ReportTaskStatus requests.Integer `position:"Query" name:"ReportTaskStatus"`
	Lang             string           `position:"Query" name:"Lang"`
}

// ModifyReportTaskStatusResponse is the response struct for api ModifyReportTaskStatus
type ModifyReportTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReportTaskStatusRequest creates a request to invoke ModifyReportTaskStatus API
func CreateModifyReportTaskStatusRequest() (request *ModifyReportTaskStatusRequest) {
	request = &ModifyReportTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "ModifyReportTaskStatus", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyReportTaskStatusResponse creates a response to parse from ModifyReportTaskStatus response
func CreateModifyReportTaskStatusResponse() (response *ModifyReportTaskStatusResponse) {
	response = &ModifyReportTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}