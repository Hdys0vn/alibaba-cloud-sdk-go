package linkwan

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

// ReturnJoinPermission invokes the linkwan.ReturnJoinPermission API synchronously
func (client *Client) ReturnJoinPermission(request *ReturnJoinPermissionRequest) (response *ReturnJoinPermissionResponse, err error) {
	response = CreateReturnJoinPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// ReturnJoinPermissionWithChan invokes the linkwan.ReturnJoinPermission API asynchronously
func (client *Client) ReturnJoinPermissionWithChan(request *ReturnJoinPermissionRequest) (<-chan *ReturnJoinPermissionResponse, <-chan error) {
	responseChan := make(chan *ReturnJoinPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReturnJoinPermission(request)
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

// ReturnJoinPermissionWithCallback invokes the linkwan.ReturnJoinPermission API asynchronously
func (client *Client) ReturnJoinPermissionWithCallback(request *ReturnJoinPermissionRequest, callback func(response *ReturnJoinPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReturnJoinPermissionResponse
		var err error
		defer close(result)
		response, err = client.ReturnJoinPermission(request)
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

// ReturnJoinPermissionRequest is the request struct for api ReturnJoinPermission
type ReturnJoinPermissionRequest struct {
	*requests.RpcRequest
	JoinPermissionId   string `position:"Query" name:"JoinPermissionId"`
	JoinPermissionType string `position:"Query" name:"JoinPermissionType"`
	ApiProduct         string `position:"Body" name:"ApiProduct"`
	ApiRevision        string `position:"Body" name:"ApiRevision"`
}

// ReturnJoinPermissionResponse is the response struct for api ReturnJoinPermission
type ReturnJoinPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateReturnJoinPermissionRequest creates a request to invoke ReturnJoinPermission API
func CreateReturnJoinPermissionRequest() (request *ReturnJoinPermissionRequest) {
	request = &ReturnJoinPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ReturnJoinPermission", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReturnJoinPermissionResponse creates a response to parse from ReturnJoinPermission response
func CreateReturnJoinPermissionResponse() (response *ReturnJoinPermissionResponse) {
	response = &ReturnJoinPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
