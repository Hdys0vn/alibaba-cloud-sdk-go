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

// CreateStructSyncOrder invokes the dms_enterprise.CreateStructSyncOrder API synchronously
func (client *Client) CreateStructSyncOrder(request *CreateStructSyncOrderRequest) (response *CreateStructSyncOrderResponse, err error) {
	response = CreateCreateStructSyncOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStructSyncOrderWithChan invokes the dms_enterprise.CreateStructSyncOrder API asynchronously
func (client *Client) CreateStructSyncOrderWithChan(request *CreateStructSyncOrderRequest) (<-chan *CreateStructSyncOrderResponse, <-chan error) {
	responseChan := make(chan *CreateStructSyncOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStructSyncOrder(request)
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

// CreateStructSyncOrderWithCallback invokes the dms_enterprise.CreateStructSyncOrder API asynchronously
func (client *Client) CreateStructSyncOrderWithCallback(request *CreateStructSyncOrderRequest, callback func(response *CreateStructSyncOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStructSyncOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateStructSyncOrder(request)
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

// CreateStructSyncOrderRequest is the request struct for api CreateStructSyncOrder
type CreateStructSyncOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer           `position:"Query" name:"Tid"`
	Param           CreateStructSyncOrderParam `position:"Query" name:"Param"  type:"Struct"`
	RelatedUserList *[]string                  `position:"Query" name:"RelatedUserList"  type:"Json"`
	AttachmentKey   string                     `position:"Query" name:"AttachmentKey"`
	Comment         string                     `position:"Query" name:"Comment"`
}

// CreateStructSyncOrderParam is a repeated param struct in CreateStructSyncOrderRequest
type CreateStructSyncOrderParam struct {
	SyncType      string                                         `name:"SyncType"`
	TableInfoList *[]CreateStructSyncOrderParamTableInfoListItem `name:"TableInfoList" type:"Repeated"`
	Source        CreateStructSyncOrderParamSource               `name:"Source" type:"Struct"`
	IgnoreError   string                                         `name:"IgnoreError"`
	Target        CreateStructSyncOrderParamTarget               `name:"Target" type:"Struct"`
}

// CreateStructSyncOrderParamTableInfoListItem is a repeated param struct in CreateStructSyncOrderRequest
type CreateStructSyncOrderParamTableInfoListItem struct {
	SourceTableName string `name:"SourceTableName"`
	TargetTableName string `name:"TargetTableName"`
}

// CreateStructSyncOrderParamSource is a repeated param struct in CreateStructSyncOrderRequest
type CreateStructSyncOrderParamSource struct {
	DbSearchName string `name:"DbSearchName"`
	VersionId    string `name:"VersionId"`
	DbId         string `name:"DbId"`
	Logic        string `name:"Logic"`
}

// CreateStructSyncOrderParamTarget is a repeated param struct in CreateStructSyncOrderRequest
type CreateStructSyncOrderParamTarget struct {
	DbSearchName string `name:"DbSearchName"`
	VersionId    string `name:"VersionId"`
	DbId         string `name:"DbId"`
	Logic        string `name:"Logic"`
}

// CreateStructSyncOrderResponse is the response struct for api CreateStructSyncOrder
type CreateStructSyncOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	Success           bool    `json:"Success" xml:"Success"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateStructSyncOrderRequest creates a request to invoke CreateStructSyncOrder API
func CreateCreateStructSyncOrderRequest() (request *CreateStructSyncOrderRequest) {
	request = &CreateStructSyncOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateStructSyncOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStructSyncOrderResponse creates a response to parse from CreateStructSyncOrder response
func CreateCreateStructSyncOrderResponse() (response *CreateStructSyncOrderResponse) {
	response = &CreateStructSyncOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
