package yundun_dbaudit

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

// EnableInstancePublicAccess invokes the yundun_dbaudit.EnableInstancePublicAccess API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/enableinstancepublicaccess.html
func (client *Client) EnableInstancePublicAccess(request *EnableInstancePublicAccessRequest) (response *EnableInstancePublicAccessResponse, err error) {
	response = CreateEnableInstancePublicAccessResponse()
	err = client.DoAction(request, response)
	return
}

// EnableInstancePublicAccessWithChan invokes the yundun_dbaudit.EnableInstancePublicAccess API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/enableinstancepublicaccess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableInstancePublicAccessWithChan(request *EnableInstancePublicAccessRequest) (<-chan *EnableInstancePublicAccessResponse, <-chan error) {
	responseChan := make(chan *EnableInstancePublicAccessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableInstancePublicAccess(request)
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

// EnableInstancePublicAccessWithCallback invokes the yundun_dbaudit.EnableInstancePublicAccess API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/enableinstancepublicaccess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableInstancePublicAccessWithCallback(request *EnableInstancePublicAccessRequest, callback func(response *EnableInstancePublicAccessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableInstancePublicAccessResponse
		var err error
		defer close(result)
		response, err = client.EnableInstancePublicAccess(request)
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

// EnableInstancePublicAccessRequest is the request struct for api EnableInstancePublicAccess
type EnableInstancePublicAccessRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// EnableInstancePublicAccessResponse is the response struct for api EnableInstancePublicAccess
type EnableInstancePublicAccessResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableInstancePublicAccessRequest creates a request to invoke EnableInstancePublicAccess API
func CreateEnableInstancePublicAccessRequest() (request *EnableInstancePublicAccessRequest) {
	request = &EnableInstancePublicAccessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "EnableInstancePublicAccess", "dbaudit", "openAPI")
	return
}

// CreateEnableInstancePublicAccessResponse creates a response to parse from EnableInstancePublicAccess response
func CreateEnableInstancePublicAccessResponse() (response *EnableInstancePublicAccessResponse) {
	response = &EnableInstancePublicAccessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
