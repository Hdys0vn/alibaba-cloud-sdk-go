package aegis

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

// DescribeSuspEventExportInfo invokes the aegis.DescribeSuspEventExportInfo API synchronously
// api document: https://help.aliyun.com/api/aegis/describesuspeventexportinfo.html
func (client *Client) DescribeSuspEventExportInfo(request *DescribeSuspEventExportInfoRequest) (response *DescribeSuspEventExportInfoResponse, err error) {
	response = CreateDescribeSuspEventExportInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspEventExportInfoWithChan invokes the aegis.DescribeSuspEventExportInfo API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspeventexportinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspEventExportInfoWithChan(request *DescribeSuspEventExportInfoRequest) (<-chan *DescribeSuspEventExportInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspEventExportInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspEventExportInfo(request)
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

// DescribeSuspEventExportInfoWithCallback invokes the aegis.DescribeSuspEventExportInfo API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspeventexportinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspEventExportInfoWithCallback(request *DescribeSuspEventExportInfoRequest, callback func(response *DescribeSuspEventExportInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspEventExportInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspEventExportInfo(request)
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

// DescribeSuspEventExportInfoRequest is the request struct for api DescribeSuspEventExportInfo
type DescribeSuspEventExportInfoRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	From     string           `position:"Query" name:"From"`
	ExportId requests.Integer `position:"Query" name:"ExportId"`
}

// DescribeSuspEventExportInfoResponse is the response struct for api DescribeSuspEventExportInfo
type DescribeSuspEventExportInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	FileName     string `json:"FileName" xml:"FileName"`
	GmtModified  int    `json:"GmtModified" xml:"GmtModified"`
	Progress     int    `json:"Progress" xml:"Progress"`
	Id           int    `json:"Id" xml:"Id"`
	Type         string `json:"Type" xml:"Type"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	GmtCreate    int    `json:"GmtCreate" xml:"GmtCreate"`
	Properties   string `json:"Properties" xml:"Properties"`
	ExportStatus string `json:"ExportStatus" xml:"ExportStatus"`
	Link         string `json:"Link" xml:"Link"`
}

// CreateDescribeSuspEventExportInfoRequest creates a request to invoke DescribeSuspEventExportInfo API
func CreateDescribeSuspEventExportInfoRequest() (request *DescribeSuspEventExportInfoRequest) {
	request = &DescribeSuspEventExportInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSuspEventExportInfo", "vipaegis", "openAPI")
	return
}

// CreateDescribeSuspEventExportInfoResponse creates a response to parse from DescribeSuspEventExportInfo response
func CreateDescribeSuspEventExportInfoResponse() (response *DescribeSuspEventExportInfoResponse) {
	response = &DescribeSuspEventExportInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
