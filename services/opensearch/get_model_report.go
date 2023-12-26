package opensearch

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

// GetModelReport invokes the opensearch.GetModelReport API synchronously
func (client *Client) GetModelReport(request *GetModelReportRequest) (response *GetModelReportResponse, err error) {
	response = CreateGetModelReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetModelReportWithChan invokes the opensearch.GetModelReport API asynchronously
func (client *Client) GetModelReportWithChan(request *GetModelReportRequest) (<-chan *GetModelReportResponse, <-chan error) {
	responseChan := make(chan *GetModelReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetModelReport(request)
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

// GetModelReportWithCallback invokes the opensearch.GetModelReport API asynchronously
func (client *Client) GetModelReportWithCallback(request *GetModelReportRequest, callback func(response *GetModelReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetModelReportResponse
		var err error
		defer close(result)
		response, err = client.GetModelReport(request)
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

// GetModelReportRequest is the request struct for api GetModelReport
type GetModelReportRequest struct {
	*requests.RoaRequest
	ModelName        string `position:"Path" name:"modelName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetModelReportResponse is the response struct for api GetModelReport
type GetModelReportResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"requestId" xml:"requestId"`
}

// CreateGetModelReportRequest creates a request to invoke GetModelReport API
func CreateGetModelReportRequest() (request *GetModelReportRequest) {
	request = &GetModelReportRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetModelReport", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/models/[modelName]/report", "", "")
	request.Method = requests.GET
	return
}

// CreateGetModelReportResponse creates a response to parse from GetModelReport response
func CreateGetModelReportResponse() (response *GetModelReportResponse) {
	response = &GetModelReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
