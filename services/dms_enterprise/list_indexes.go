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

// ListIndexes invokes the dms_enterprise.ListIndexes API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listindexes.html
func (client *Client) ListIndexes(request *ListIndexesRequest) (response *ListIndexesResponse, err error) {
	response = CreateListIndexesResponse()
	err = client.DoAction(request, response)
	return
}

// ListIndexesWithChan invokes the dms_enterprise.ListIndexes API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listindexes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListIndexesWithChan(request *ListIndexesRequest) (<-chan *ListIndexesResponse, <-chan error) {
	responseChan := make(chan *ListIndexesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIndexes(request)
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

// ListIndexesWithCallback invokes the dms_enterprise.ListIndexes API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listindexes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListIndexesWithCallback(request *ListIndexesRequest, callback func(response *ListIndexesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIndexesResponse
		var err error
		defer close(result)
		response, err = client.ListIndexes(request)
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

// ListIndexesRequest is the request struct for api ListIndexes
type ListIndexesRequest struct {
	*requests.RpcRequest
	TableId string           `position:"Query" name:"TableId"`
	Logic   requests.Boolean `position:"Query" name:"Logic"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
}

// ListIndexesResponse is the response struct for api ListIndexes
type ListIndexesResponse struct {
	*responses.BaseResponse
	RequestId    string    `json:"RequestId" xml:"RequestId"`
	Success      bool      `json:"Success" xml:"Success"`
	ErrorMessage string    `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string    `json:"ErrorCode" xml:"ErrorCode"`
	IndexList    IndexList `json:"IndexList" xml:"IndexList"`
}

// CreateListIndexesRequest creates a request to invoke ListIndexes API
func CreateListIndexesRequest() (request *ListIndexesRequest) {
	request = &ListIndexesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListIndexes", "dmsenterprise", "openAPI")
	return
}

// CreateListIndexesResponse creates a response to parse from ListIndexes response
func CreateListIndexesResponse() (response *ListIndexesResponse) {
	response = &ListIndexesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
