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

// GetSQLReviewOptimizeDetail invokes the dms_enterprise.GetSQLReviewOptimizeDetail API synchronously
func (client *Client) GetSQLReviewOptimizeDetail(request *GetSQLReviewOptimizeDetailRequest) (response *GetSQLReviewOptimizeDetailResponse, err error) {
	response = CreateGetSQLReviewOptimizeDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetSQLReviewOptimizeDetailWithChan invokes the dms_enterprise.GetSQLReviewOptimizeDetail API asynchronously
func (client *Client) GetSQLReviewOptimizeDetailWithChan(request *GetSQLReviewOptimizeDetailRequest) (<-chan *GetSQLReviewOptimizeDetailResponse, <-chan error) {
	responseChan := make(chan *GetSQLReviewOptimizeDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSQLReviewOptimizeDetail(request)
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

// GetSQLReviewOptimizeDetailWithCallback invokes the dms_enterprise.GetSQLReviewOptimizeDetail API asynchronously
func (client *Client) GetSQLReviewOptimizeDetailWithCallback(request *GetSQLReviewOptimizeDetailRequest, callback func(response *GetSQLReviewOptimizeDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSQLReviewOptimizeDetailResponse
		var err error
		defer close(result)
		response, err = client.GetSQLReviewOptimizeDetail(request)
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

// GetSQLReviewOptimizeDetailRequest is the request struct for api GetSQLReviewOptimizeDetail
type GetSQLReviewOptimizeDetailRequest struct {
	*requests.RpcRequest
	Tid               requests.Integer `position:"Query" name:"Tid"`
	SQLReviewQueryKey string           `position:"Query" name:"SQLReviewQueryKey"`
}

// GetSQLReviewOptimizeDetailResponse is the response struct for api GetSQLReviewOptimizeDetail
type GetSQLReviewOptimizeDetailResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	OptimizeDetail OptimizeDetail `json:"OptimizeDetail" xml:"OptimizeDetail"`
}

// CreateGetSQLReviewOptimizeDetailRequest creates a request to invoke GetSQLReviewOptimizeDetail API
func CreateGetSQLReviewOptimizeDetailRequest() (request *GetSQLReviewOptimizeDetailRequest) {
	request = &GetSQLReviewOptimizeDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetSQLReviewOptimizeDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSQLReviewOptimizeDetailResponse creates a response to parse from GetSQLReviewOptimizeDetail response
func CreateGetSQLReviewOptimizeDetailResponse() (response *GetSQLReviewOptimizeDetailResponse) {
	response = &GetSQLReviewOptimizeDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
