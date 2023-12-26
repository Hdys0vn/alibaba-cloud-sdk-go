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

// GetDataTrackJobTableMeta invokes the dms_enterprise.GetDataTrackJobTableMeta API synchronously
func (client *Client) GetDataTrackJobTableMeta(request *GetDataTrackJobTableMetaRequest) (response *GetDataTrackJobTableMetaResponse, err error) {
	response = CreateGetDataTrackJobTableMetaResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataTrackJobTableMetaWithChan invokes the dms_enterprise.GetDataTrackJobTableMeta API asynchronously
func (client *Client) GetDataTrackJobTableMetaWithChan(request *GetDataTrackJobTableMetaRequest) (<-chan *GetDataTrackJobTableMetaResponse, <-chan error) {
	responseChan := make(chan *GetDataTrackJobTableMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataTrackJobTableMeta(request)
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

// GetDataTrackJobTableMetaWithCallback invokes the dms_enterprise.GetDataTrackJobTableMeta API asynchronously
func (client *Client) GetDataTrackJobTableMetaWithCallback(request *GetDataTrackJobTableMetaRequest, callback func(response *GetDataTrackJobTableMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataTrackJobTableMetaResponse
		var err error
		defer close(result)
		response, err = client.GetDataTrackJobTableMeta(request)
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

// GetDataTrackJobTableMetaRequest is the request struct for api GetDataTrackJobTableMeta
type GetDataTrackJobTableMetaRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// GetDataTrackJobTableMetaResponse is the response struct for api GetDataTrackJobTableMeta
type GetDataTrackJobTableMetaResponse struct {
	*responses.BaseResponse
	RequestId     string      `json:"RequestId" xml:"RequestId"`
	Success       bool        `json:"Success" xml:"Success"`
	ErrorMessage  string      `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string      `json:"ErrorCode" xml:"ErrorCode"`
	TableMetaList []TableMeta `json:"TableMetaList" xml:"TableMetaList"`
}

// CreateGetDataTrackJobTableMetaRequest creates a request to invoke GetDataTrackJobTableMeta API
func CreateGetDataTrackJobTableMetaRequest() (request *GetDataTrackJobTableMetaRequest) {
	request = &GetDataTrackJobTableMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataTrackJobTableMeta", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDataTrackJobTableMetaResponse creates a response to parse from GetDataTrackJobTableMeta response
func CreateGetDataTrackJobTableMetaResponse() (response *GetDataTrackJobTableMetaResponse) {
	response = &GetDataTrackJobTableMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}