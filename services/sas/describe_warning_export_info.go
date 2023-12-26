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

// DescribeWarningExportInfo invokes the sas.DescribeWarningExportInfo API synchronously
func (client *Client) DescribeWarningExportInfo(request *DescribeWarningExportInfoRequest) (response *DescribeWarningExportInfoResponse, err error) {
	response = CreateDescribeWarningExportInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWarningExportInfoWithChan invokes the sas.DescribeWarningExportInfo API asynchronously
func (client *Client) DescribeWarningExportInfoWithChan(request *DescribeWarningExportInfoRequest) (<-chan *DescribeWarningExportInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeWarningExportInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWarningExportInfo(request)
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

// DescribeWarningExportInfoWithCallback invokes the sas.DescribeWarningExportInfo API asynchronously
func (client *Client) DescribeWarningExportInfoWithCallback(request *DescribeWarningExportInfoRequest, callback func(response *DescribeWarningExportInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWarningExportInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeWarningExportInfo(request)
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

// DescribeWarningExportInfoRequest is the request struct for api DescribeWarningExportInfo
type DescribeWarningExportInfoRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	ExportId requests.Integer `position:"Query" name:"ExportId"`
}

// DescribeWarningExportInfoResponse is the response struct for api DescribeWarningExportInfo
type DescribeWarningExportInfoResponse struct {
	*responses.BaseResponse
	Link         string `json:"Link" xml:"Link"`
	Progress     int    `json:"Progress" xml:"Progress"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	CurrentCount int    `json:"CurrentCount" xml:"CurrentCount"`
	Message      string `json:"Message" xml:"Message"`
	FileName     string `json:"FileName" xml:"FileName"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	ExportStatus string `json:"ExportStatus" xml:"ExportStatus"`
	Id           int64  `json:"Id" xml:"Id"`
}

// CreateDescribeWarningExportInfoRequest creates a request to invoke DescribeWarningExportInfo API
func CreateDescribeWarningExportInfoRequest() (request *DescribeWarningExportInfoRequest) {
	request = &DescribeWarningExportInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeWarningExportInfo", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWarningExportInfoResponse creates a response to parse from DescribeWarningExportInfo response
func CreateDescribeWarningExportInfoResponse() (response *DescribeWarningExportInfoResponse) {
	response = &DescribeWarningExportInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}