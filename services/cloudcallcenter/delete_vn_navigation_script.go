package cloudcallcenter

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

// DeleteVnNavigationScript invokes the cloudcallcenter.DeleteVnNavigationScript API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevnnavigationscript.html
func (client *Client) DeleteVnNavigationScript(request *DeleteVnNavigationScriptRequest) (response *DeleteVnNavigationScriptResponse, err error) {
	response = CreateDeleteVnNavigationScriptResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVnNavigationScriptWithChan invokes the cloudcallcenter.DeleteVnNavigationScript API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevnnavigationscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVnNavigationScriptWithChan(request *DeleteVnNavigationScriptRequest) (<-chan *DeleteVnNavigationScriptResponse, <-chan error) {
	responseChan := make(chan *DeleteVnNavigationScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVnNavigationScript(request)
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

// DeleteVnNavigationScriptWithCallback invokes the cloudcallcenter.DeleteVnNavigationScript API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevnnavigationscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVnNavigationScriptWithCallback(request *DeleteVnNavigationScriptRequest, callback func(response *DeleteVnNavigationScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVnNavigationScriptResponse
		var err error
		defer close(result)
		response, err = client.DeleteVnNavigationScript(request)
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

// DeleteVnNavigationScriptRequest is the request struct for api DeleteVnNavigationScript
type DeleteVnNavigationScriptRequest struct {
	*requests.RpcRequest
	InstanceId         string `position:"Query" name:"InstanceId"`
	NavigationScriptId string `position:"Query" name:"NavigationScriptId"`
}

// DeleteVnNavigationScriptResponse is the response struct for api DeleteVnNavigationScript
type DeleteVnNavigationScriptResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVnNavigationScriptRequest creates a request to invoke DeleteVnNavigationScript API
func CreateDeleteVnNavigationScriptRequest() (request *DeleteVnNavigationScriptRequest) {
	request = &DeleteVnNavigationScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteVnNavigationScript", "", "")
	request.Method = requests.GET
	return
}

// CreateDeleteVnNavigationScriptResponse creates a response to parse from DeleteVnNavigationScript response
func CreateDeleteVnNavigationScriptResponse() (response *DeleteVnNavigationScriptResponse) {
	response = &DeleteVnNavigationScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
