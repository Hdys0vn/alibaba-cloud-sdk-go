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

// ListLogicTableRouteConfig invokes the dms_enterprise.ListLogicTableRouteConfig API synchronously
func (client *Client) ListLogicTableRouteConfig(request *ListLogicTableRouteConfigRequest) (response *ListLogicTableRouteConfigResponse, err error) {
	response = CreateListLogicTableRouteConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogicTableRouteConfigWithChan invokes the dms_enterprise.ListLogicTableRouteConfig API asynchronously
func (client *Client) ListLogicTableRouteConfigWithChan(request *ListLogicTableRouteConfigRequest) (<-chan *ListLogicTableRouteConfigResponse, <-chan error) {
	responseChan := make(chan *ListLogicTableRouteConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogicTableRouteConfig(request)
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

// ListLogicTableRouteConfigWithCallback invokes the dms_enterprise.ListLogicTableRouteConfig API asynchronously
func (client *Client) ListLogicTableRouteConfigWithCallback(request *ListLogicTableRouteConfigRequest, callback func(response *ListLogicTableRouteConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogicTableRouteConfigResponse
		var err error
		defer close(result)
		response, err = client.ListLogicTableRouteConfig(request)
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

// ListLogicTableRouteConfigRequest is the request struct for api ListLogicTableRouteConfig
type ListLogicTableRouteConfigRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	TableId requests.Integer `position:"Query" name:"TableId"`
}

// ListLogicTableRouteConfigResponse is the response struct for api ListLogicTableRouteConfig
type ListLogicTableRouteConfigResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode                 string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage              string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                   bool                      `json:"Success" xml:"Success"`
	LogicTableRouteConfigList LogicTableRouteConfigList `json:"LogicTableRouteConfigList" xml:"LogicTableRouteConfigList"`
}

// CreateListLogicTableRouteConfigRequest creates a request to invoke ListLogicTableRouteConfig API
func CreateListLogicTableRouteConfigRequest() (request *ListLogicTableRouteConfigRequest) {
	request = &ListLogicTableRouteConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListLogicTableRouteConfig", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLogicTableRouteConfigResponse creates a response to parse from ListLogicTableRouteConfig response
func CreateListLogicTableRouteConfigResponse() (response *ListLogicTableRouteConfigResponse) {
	response = &ListLogicTableRouteConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}