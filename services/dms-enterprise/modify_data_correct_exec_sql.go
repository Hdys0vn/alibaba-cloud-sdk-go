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

// ModifyDataCorrectExecSQL invokes the dms_enterprise.ModifyDataCorrectExecSQL API synchronously
func (client *Client) ModifyDataCorrectExecSQL(request *ModifyDataCorrectExecSQLRequest) (response *ModifyDataCorrectExecSQLResponse, err error) {
	response = CreateModifyDataCorrectExecSQLResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataCorrectExecSQLWithChan invokes the dms_enterprise.ModifyDataCorrectExecSQL API asynchronously
func (client *Client) ModifyDataCorrectExecSQLWithChan(request *ModifyDataCorrectExecSQLRequest) (<-chan *ModifyDataCorrectExecSQLResponse, <-chan error) {
	responseChan := make(chan *ModifyDataCorrectExecSQLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataCorrectExecSQL(request)
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

// ModifyDataCorrectExecSQLWithCallback invokes the dms_enterprise.ModifyDataCorrectExecSQL API asynchronously
func (client *Client) ModifyDataCorrectExecSQLWithCallback(request *ModifyDataCorrectExecSQLRequest, callback func(response *ModifyDataCorrectExecSQLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataCorrectExecSQLResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataCorrectExecSQL(request)
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

// ModifyDataCorrectExecSQLRequest is the request struct for api ModifyDataCorrectExecSQL
type ModifyDataCorrectExecSQLRequest struct {
	*requests.RpcRequest
	ExecSQL string           `position:"Query" name:"ExecSQL"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// ModifyDataCorrectExecSQLResponse is the response struct for api ModifyDataCorrectExecSQL
type ModifyDataCorrectExecSQLResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateModifyDataCorrectExecSQLRequest creates a request to invoke ModifyDataCorrectExecSQL API
func CreateModifyDataCorrectExecSQLRequest() (request *ModifyDataCorrectExecSQLRequest) {
	request = &ModifyDataCorrectExecSQLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ModifyDataCorrectExecSQL", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDataCorrectExecSQLResponse creates a response to parse from ModifyDataCorrectExecSQL response
func CreateModifyDataCorrectExecSQLResponse() (response *ModifyDataCorrectExecSQLResponse) {
	response = &ModifyDataCorrectExecSQLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
