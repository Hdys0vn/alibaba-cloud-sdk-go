package snsuapi

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

// MobileStatusQuery invokes the snsuapi.MobileStatusQuery API synchronously
// api document: https://help.aliyun.com/api/snsuapi/mobilestatusquery.html
func (client *Client) MobileStatusQuery(request *MobileStatusQueryRequest) (response *MobileStatusQueryResponse, err error) {
	response = CreateMobileStatusQueryResponse()
	err = client.DoAction(request, response)
	return
}

// MobileStatusQueryWithChan invokes the snsuapi.MobileStatusQuery API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/mobilestatusquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MobileStatusQueryWithChan(request *MobileStatusQueryRequest) (<-chan *MobileStatusQueryResponse, <-chan error) {
	responseChan := make(chan *MobileStatusQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MobileStatusQuery(request)
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

// MobileStatusQueryWithCallback invokes the snsuapi.MobileStatusQuery API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/mobilestatusquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MobileStatusQueryWithCallback(request *MobileStatusQueryRequest, callback func(response *MobileStatusQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MobileStatusQueryResponse
		var err error
		defer close(result)
		response, err = client.MobileStatusQuery(request)
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

// MobileStatusQueryRequest is the request struct for api MobileStatusQuery
type MobileStatusQueryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CorrelationId        string           `position:"Query" name:"CorrelationId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// MobileStatusQueryResponse is the response struct for api MobileStatusQuery
type MobileStatusQueryResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string `json:"ResultMessage" xml:"ResultMessage"`
	ResultModule  bool   `json:"ResultModule" xml:"ResultModule"`
}

// CreateMobileStatusQueryRequest creates a request to invoke MobileStatusQuery API
func CreateMobileStatusQueryRequest() (request *MobileStatusQueryRequest) {
	request = &MobileStatusQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Snsuapi", "2018-07-09", "MobileStatusQuery", "snsuapi", "openAPI")
	return
}

// CreateMobileStatusQueryResponse creates a response to parse from MobileStatusQuery response
func CreateMobileStatusQueryResponse() (response *MobileStatusQueryResponse) {
	response = &MobileStatusQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}