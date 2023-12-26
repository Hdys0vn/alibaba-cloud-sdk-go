package clickhouse

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

// CreateMonitorDataReport invokes the clickhouse.CreateMonitorDataReport API synchronously
func (client *Client) CreateMonitorDataReport(request *CreateMonitorDataReportRequest) (response *CreateMonitorDataReportResponse, err error) {
	response = CreateCreateMonitorDataReportResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMonitorDataReportWithChan invokes the clickhouse.CreateMonitorDataReport API asynchronously
func (client *Client) CreateMonitorDataReportWithChan(request *CreateMonitorDataReportRequest) (<-chan *CreateMonitorDataReportResponse, <-chan error) {
	responseChan := make(chan *CreateMonitorDataReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMonitorDataReport(request)
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

// CreateMonitorDataReportWithCallback invokes the clickhouse.CreateMonitorDataReport API asynchronously
func (client *Client) CreateMonitorDataReportWithCallback(request *CreateMonitorDataReportRequest, callback func(response *CreateMonitorDataReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMonitorDataReportResponse
		var err error
		defer close(result)
		response, err = client.CreateMonitorDataReport(request)
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

// CreateMonitorDataReportRequest is the request struct for api CreateMonitorDataReport
type CreateMonitorDataReportRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateMonitorDataReportResponse is the response struct for api CreateMonitorDataReport
type CreateMonitorDataReportResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateMonitorDataReportRequest creates a request to invoke CreateMonitorDataReport API
func CreateCreateMonitorDataReportRequest() (request *CreateMonitorDataReportRequest) {
	request = &CreateMonitorDataReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CreateMonitorDataReport", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMonitorDataReportResponse creates a response to parse from CreateMonitorDataReport response
func CreateCreateMonitorDataReportResponse() (response *CreateMonitorDataReportResponse) {
	response = &CreateMonitorDataReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
