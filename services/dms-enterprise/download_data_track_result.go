package dms_enterprise

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

// DownloadDataTrackResult invokes the dms_enterprise.DownloadDataTrackResult API synchronously
func (client *Client) DownloadDataTrackResult(request *DownloadDataTrackResultRequest) (response *DownloadDataTrackResultResponse, err error) {
	response = CreateDownloadDataTrackResultResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadDataTrackResultWithChan invokes the dms_enterprise.DownloadDataTrackResult API asynchronously
func (client *Client) DownloadDataTrackResultWithChan(request *DownloadDataTrackResultRequest) (<-chan *DownloadDataTrackResultResponse, <-chan error) {
	responseChan := make(chan *DownloadDataTrackResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadDataTrackResult(request)
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

// DownloadDataTrackResultWithCallback invokes the dms_enterprise.DownloadDataTrackResult API asynchronously
func (client *Client) DownloadDataTrackResultWithCallback(request *DownloadDataTrackResultRequest, callback func(response *DownloadDataTrackResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadDataTrackResultResponse
		var err error
		defer close(result)
		response, err = client.DownloadDataTrackResult(request)
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

// DownloadDataTrackResultRequest is the request struct for api DownloadDataTrackResult
type DownloadDataTrackResultRequest struct {
	*requests.RpcRequest
	FilterStartTime string                              `position:"Query" name:"FilterStartTime"`
	FilterTypeList  *[]string                           `position:"Query" name:"FilterTypeList"  type:"Json"`
	Tid             requests.Integer                    `position:"Query" name:"Tid"`
	RollbackSQLType string                              `position:"Query" name:"RollbackSQLType"`
	EventIdList     *[]string                           `position:"Query" name:"EventIdList"  type:"Json"`
	OrderId         requests.Integer                    `position:"Query" name:"OrderId"`
	FilterTableList *[]string                           `position:"Query" name:"FilterTableList"  type:"Json"`
	FilterEndTime   string                              `position:"Query" name:"FilterEndTime"`
	ColumnFilter    DownloadDataTrackResultColumnFilter `position:"Query" name:"ColumnFilter"  type:"Struct"`
}

// DownloadDataTrackResultColumnFilter is a repeated param struct in DownloadDataTrackResultRequest
type DownloadDataTrackResultColumnFilter struct {
	BetweenStart string    `name:"BetweenStart"`
	BetweenEnd   string    `name:"BetweenEnd"`
	ColumnName   string    `name:"ColumnName"`
	Value        string    `name:"Value"`
	Operator     string    `name:"Operator"`
	InList       *[]string `name:"InList" type:"Repeated"`
}

// DownloadDataTrackResultResponse is the response struct for api DownloadDataTrackResult
type DownloadDataTrackResultResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	ErrorMessage  string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
	DownloadKeyId string `json:"DownloadKeyId" xml:"DownloadKeyId"`
}

// CreateDownloadDataTrackResultRequest creates a request to invoke DownloadDataTrackResult API
func CreateDownloadDataTrackResultRequest() (request *DownloadDataTrackResultRequest) {
	request = &DownloadDataTrackResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "DownloadDataTrackResult", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDownloadDataTrackResultResponse creates a response to parse from DownloadDataTrackResult response
func CreateDownloadDataTrackResultResponse() (response *DownloadDataTrackResultResponse) {
	response = &DownloadDataTrackResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}