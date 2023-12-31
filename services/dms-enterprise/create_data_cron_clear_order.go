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

// CreateDataCronClearOrder invokes the dms_enterprise.CreateDataCronClearOrder API synchronously
func (client *Client) CreateDataCronClearOrder(request *CreateDataCronClearOrderRequest) (response *CreateDataCronClearOrderResponse, err error) {
	response = CreateCreateDataCronClearOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataCronClearOrderWithChan invokes the dms_enterprise.CreateDataCronClearOrder API asynchronously
func (client *Client) CreateDataCronClearOrderWithChan(request *CreateDataCronClearOrderRequest) (<-chan *CreateDataCronClearOrderResponse, <-chan error) {
	responseChan := make(chan *CreateDataCronClearOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataCronClearOrder(request)
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

// CreateDataCronClearOrderWithCallback invokes the dms_enterprise.CreateDataCronClearOrder API asynchronously
func (client *Client) CreateDataCronClearOrderWithCallback(request *CreateDataCronClearOrderRequest, callback func(response *CreateDataCronClearOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataCronClearOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateDataCronClearOrder(request)
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

// CreateDataCronClearOrderRequest is the request struct for api CreateDataCronClearOrder
type CreateDataCronClearOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer              `position:"Query" name:"Tid"`
	Param           CreateDataCronClearOrderParam `position:"Query" name:"Param"  type:"Struct"`
	RelatedUserList *[]string                     `position:"Query" name:"RelatedUserList"  type:"Json"`
	AttachmentKey   string                        `position:"Query" name:"AttachmentKey"`
	Comment         string                        `position:"Query" name:"Comment"`
}

// CreateDataCronClearOrderParam is a repeated param struct in CreateDataCronClearOrderRequest
type CreateDataCronClearOrderParam struct {
	Classify          string                                                `name:"Classify"`
	DbItemList        *[]CreateDataCronClearOrderParamDbItemListItem        `name:"DbItemList" type:"Repeated"`
	CronClearItemList *[]CreateDataCronClearOrderParamCronClearItemListItem `name:"CronClearItemList" type:"Repeated"`
	DurationHour      string                                                `name:"DurationHour"`
	CronFormat        string                                                `name:"CronFormat"`
	SpecifyDuration   string                                                `name:"specifyDuration"`
}

// CreateDataCronClearOrderParamDbItemListItem is a repeated param struct in CreateDataCronClearOrderRequest
type CreateDataCronClearOrderParamDbItemListItem struct {
	DbId  string `name:"DbId"`
	Logic string `name:"Logic"`
}

// CreateDataCronClearOrderParamCronClearItemListItem is a repeated param struct in CreateDataCronClearOrderRequest
type CreateDataCronClearOrderParamCronClearItemListItem struct {
	FilterSQL  string `name:"FilterSQL"`
	RemainDays string `name:"RemainDays"`
	TableName  string `name:"TableName"`
	ColumnName string `name:"ColumnName"`
	TimeUnit   string `name:"TimeUnit"`
}

// CreateDataCronClearOrderResponse is the response struct for api CreateDataCronClearOrder
type CreateDataCronClearOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	Success           bool    `json:"Success" xml:"Success"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateDataCronClearOrderRequest creates a request to invoke CreateDataCronClearOrder API
func CreateCreateDataCronClearOrderRequest() (request *CreateDataCronClearOrderRequest) {
	request = &CreateDataCronClearOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateDataCronClearOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataCronClearOrderResponse creates a response to parse from CreateDataCronClearOrder response
func CreateCreateDataCronClearOrderResponse() (response *CreateDataCronClearOrderResponse) {
	response = &CreateDataCronClearOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
