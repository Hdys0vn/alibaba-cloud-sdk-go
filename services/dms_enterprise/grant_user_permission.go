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

// GrantUserPermission invokes the dms_enterprise.GrantUserPermission API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/grantuserpermission.html
func (client *Client) GrantUserPermission(request *GrantUserPermissionRequest) (response *GrantUserPermissionResponse, err error) {
	response = CreateGrantUserPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GrantUserPermissionWithChan invokes the dms_enterprise.GrantUserPermission API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/grantuserpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantUserPermissionWithChan(request *GrantUserPermissionRequest) (<-chan *GrantUserPermissionResponse, <-chan error) {
	responseChan := make(chan *GrantUserPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantUserPermission(request)
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

// GrantUserPermissionWithCallback invokes the dms_enterprise.GrantUserPermission API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/grantuserpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantUserPermissionWithCallback(request *GrantUserPermissionRequest, callback func(response *GrantUserPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantUserPermissionResponse
		var err error
		defer close(result)
		response, err = client.GrantUserPermission(request)
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

// GrantUserPermissionRequest is the request struct for api GrantUserPermission
type GrantUserPermissionRequest struct {
	*requests.RpcRequest
	PermTypes  string           `position:"Query" name:"PermTypes"`
	DsType     string           `position:"Query" name:"DsType"`
	ExpireDate string           `position:"Query" name:"ExpireDate"`
	UserId     string           `position:"Query" name:"UserId"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	DbId       string           `position:"Query" name:"DbId"`
	TableId    string           `position:"Query" name:"TableId"`
	Logic      requests.Boolean `position:"Query" name:"Logic"`
	TableName  string           `position:"Query" name:"TableName"`
}

// GrantUserPermissionResponse is the response struct for api GrantUserPermission
type GrantUserPermissionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateGrantUserPermissionRequest creates a request to invoke GrantUserPermission API
func CreateGrantUserPermissionRequest() (request *GrantUserPermissionRequest) {
	request = &GrantUserPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GrantUserPermission", "dmsenterprise", "openAPI")
	return
}

// CreateGrantUserPermissionResponse creates a response to parse from GrantUserPermission response
func CreateGrantUserPermissionResponse() (response *GrantUserPermissionResponse) {
	response = &GrantUserPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}