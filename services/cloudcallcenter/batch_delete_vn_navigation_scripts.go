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

// BatchDeleteVnNavigationScripts invokes the cloudcallcenter.BatchDeleteVnNavigationScripts API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchdeletevnnavigationscripts.html
func (client *Client) BatchDeleteVnNavigationScripts(request *BatchDeleteVnNavigationScriptsRequest) (response *BatchDeleteVnNavigationScriptsResponse, err error) {
	response = CreateBatchDeleteVnNavigationScriptsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteVnNavigationScriptsWithChan invokes the cloudcallcenter.BatchDeleteVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchdeletevnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteVnNavigationScriptsWithChan(request *BatchDeleteVnNavigationScriptsRequest) (<-chan *BatchDeleteVnNavigationScriptsResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteVnNavigationScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteVnNavigationScripts(request)
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

// BatchDeleteVnNavigationScriptsWithCallback invokes the cloudcallcenter.BatchDeleteVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchdeletevnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteVnNavigationScriptsWithCallback(request *BatchDeleteVnNavigationScriptsRequest, callback func(response *BatchDeleteVnNavigationScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteVnNavigationScriptsResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteVnNavigationScripts(request)
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

// BatchDeleteVnNavigationScriptsRequest is the request struct for api BatchDeleteVnNavigationScripts
type BatchDeleteVnNavigationScriptsRequest struct {
	*requests.RpcRequest
	NavigationScriptIds *[]string `position:"Query" name:"NavigationScriptIds"  type:"Repeated"`
	InstanceId          string    `position:"Query" name:"InstanceId"`
}

// BatchDeleteVnNavigationScriptsResponse is the response struct for api BatchDeleteVnNavigationScripts
type BatchDeleteVnNavigationScriptsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchDeleteVnNavigationScriptsRequest creates a request to invoke BatchDeleteVnNavigationScripts API
func CreateBatchDeleteVnNavigationScriptsRequest() (request *BatchDeleteVnNavigationScriptsRequest) {
	request = &BatchDeleteVnNavigationScriptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "BatchDeleteVnNavigationScripts", "", "")
	request.Method = requests.GET
	return
}

// CreateBatchDeleteVnNavigationScriptsResponse creates a response to parse from BatchDeleteVnNavigationScripts response
func CreateBatchDeleteVnNavigationScriptsResponse() (response *BatchDeleteVnNavigationScriptsResponse) {
	response = &BatchDeleteVnNavigationScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
