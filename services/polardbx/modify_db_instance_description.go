package polardbx

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

// ModifyDBInstanceDescription invokes the polardbx.ModifyDBInstanceDescription API synchronously
func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (response *ModifyDBInstanceDescriptionResponse, err error) {
	response = CreateModifyDBInstanceDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceDescriptionWithChan invokes the polardbx.ModifyDBInstanceDescription API asynchronously
func (client *Client) ModifyDBInstanceDescriptionWithChan(request *ModifyDBInstanceDescriptionRequest) (<-chan *ModifyDBInstanceDescriptionResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceDescription(request)
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

// ModifyDBInstanceDescriptionWithCallback invokes the polardbx.ModifyDBInstanceDescription API asynchronously
func (client *Client) ModifyDBInstanceDescriptionWithCallback(request *ModifyDBInstanceDescriptionRequest, callback func(response *ModifyDBInstanceDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceDescriptionResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceDescription(request)
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

// ModifyDBInstanceDescriptionRequest is the request struct for api ModifyDBInstanceDescription
type ModifyDBInstanceDescriptionRequest struct {
	*requests.RpcRequest
	DBInstanceName        string `position:"Query" name:"DBInstanceName"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
}

// ModifyDBInstanceDescriptionResponse is the response struct for api ModifyDBInstanceDescription
type ModifyDBInstanceDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceDescriptionRequest creates a request to invoke ModifyDBInstanceDescription API
func CreateModifyDBInstanceDescriptionRequest() (request *ModifyDBInstanceDescriptionRequest) {
	request = &ModifyDBInstanceDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "ModifyDBInstanceDescription", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceDescriptionResponse creates a response to parse from ModifyDBInstanceDescription response
func CreateModifyDBInstanceDescriptionResponse() (response *ModifyDBInstanceDescriptionResponse) {
	response = &ModifyDBInstanceDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
