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

// GetStructSyncJobAnalyzeResult invokes the dms_enterprise.GetStructSyncJobAnalyzeResult API synchronously
func (client *Client) GetStructSyncJobAnalyzeResult(request *GetStructSyncJobAnalyzeResultRequest) (response *GetStructSyncJobAnalyzeResultResponse, err error) {
	response = CreateGetStructSyncJobAnalyzeResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetStructSyncJobAnalyzeResultWithChan invokes the dms_enterprise.GetStructSyncJobAnalyzeResult API asynchronously
func (client *Client) GetStructSyncJobAnalyzeResultWithChan(request *GetStructSyncJobAnalyzeResultRequest) (<-chan *GetStructSyncJobAnalyzeResultResponse, <-chan error) {
	responseChan := make(chan *GetStructSyncJobAnalyzeResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetStructSyncJobAnalyzeResult(request)
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

// GetStructSyncJobAnalyzeResultWithCallback invokes the dms_enterprise.GetStructSyncJobAnalyzeResult API asynchronously
func (client *Client) GetStructSyncJobAnalyzeResultWithCallback(request *GetStructSyncJobAnalyzeResultRequest, callback func(response *GetStructSyncJobAnalyzeResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetStructSyncJobAnalyzeResultResponse
		var err error
		defer close(result)
		response, err = client.GetStructSyncJobAnalyzeResult(request)
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

// GetStructSyncJobAnalyzeResultRequest is the request struct for api GetStructSyncJobAnalyzeResult
type GetStructSyncJobAnalyzeResultRequest struct {
	*requests.RpcRequest
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	Tid         requests.Integer `position:"Query" name:"Tid"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	OrderId     requests.Integer `position:"Query" name:"OrderId"`
	CompareType string           `position:"Query" name:"CompareType"`
}

// GetStructSyncJobAnalyzeResultResponse is the response struct for api GetStructSyncJobAnalyzeResult
type GetStructSyncJobAnalyzeResultResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	Success                    bool                       `json:"Success" xml:"Success"`
	ErrorMessage               string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode                  string                     `json:"ErrorCode" xml:"ErrorCode"`
	StructSyncJobAnalyzeResult StructSyncJobAnalyzeResult `json:"StructSyncJobAnalyzeResult" xml:"StructSyncJobAnalyzeResult"`
}

// CreateGetStructSyncJobAnalyzeResultRequest creates a request to invoke GetStructSyncJobAnalyzeResult API
func CreateGetStructSyncJobAnalyzeResultRequest() (request *GetStructSyncJobAnalyzeResultRequest) {
	request = &GetStructSyncJobAnalyzeResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetStructSyncJobAnalyzeResult", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetStructSyncJobAnalyzeResultResponse creates a response to parse from GetStructSyncJobAnalyzeResult response
func CreateGetStructSyncJobAnalyzeResultResponse() (response *GetStructSyncJobAnalyzeResultResponse) {
	response = &GetStructSyncJobAnalyzeResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
