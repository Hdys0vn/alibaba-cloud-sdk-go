package cloudcallcenter

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

// DownloadOriginalStatisticsReport invokes the cloudcallcenter.DownloadOriginalStatisticsReport API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadoriginalstatisticsreport.html
func (client *Client) DownloadOriginalStatisticsReport(request *DownloadOriginalStatisticsReportRequest) (response *DownloadOriginalStatisticsReportResponse, err error) {
	response = CreateDownloadOriginalStatisticsReportResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadOriginalStatisticsReportWithChan invokes the cloudcallcenter.DownloadOriginalStatisticsReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadoriginalstatisticsreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadOriginalStatisticsReportWithChan(request *DownloadOriginalStatisticsReportRequest) (<-chan *DownloadOriginalStatisticsReportResponse, <-chan error) {
	responseChan := make(chan *DownloadOriginalStatisticsReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadOriginalStatisticsReport(request)
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

// DownloadOriginalStatisticsReportWithCallback invokes the cloudcallcenter.DownloadOriginalStatisticsReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadoriginalstatisticsreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadOriginalStatisticsReportWithCallback(request *DownloadOriginalStatisticsReportRequest, callback func(response *DownloadOriginalStatisticsReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadOriginalStatisticsReportResponse
		var err error
		defer close(result)
		response, err = client.DownloadOriginalStatisticsReport(request)
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

// DownloadOriginalStatisticsReportRequest is the request struct for api DownloadOriginalStatisticsReport
type DownloadOriginalStatisticsReportRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
}

// DownloadOriginalStatisticsReportResponse is the response struct for api DownloadOriginalStatisticsReport
type DownloadOriginalStatisticsReportResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DownloadParams DownloadParams `json:"DownloadParams" xml:"DownloadParams"`
}

// CreateDownloadOriginalStatisticsReportRequest creates a request to invoke DownloadOriginalStatisticsReport API
func CreateDownloadOriginalStatisticsReportRequest() (request *DownloadOriginalStatisticsReportRequest) {
	request = &DownloadOriginalStatisticsReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DownloadOriginalStatisticsReport", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadOriginalStatisticsReportResponse creates a response to parse from DownloadOriginalStatisticsReport response
func CreateDownloadOriginalStatisticsReportResponse() (response *DownloadOriginalStatisticsReportResponse) {
	response = &DownloadOriginalStatisticsReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
