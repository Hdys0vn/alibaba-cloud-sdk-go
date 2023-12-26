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

// UpdateTaskFlowConstants invokes the dms_enterprise.UpdateTaskFlowConstants API synchronously
func (client *Client) UpdateTaskFlowConstants(request *UpdateTaskFlowConstantsRequest) (response *UpdateTaskFlowConstantsResponse, err error) {
	response = CreateUpdateTaskFlowConstantsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaskFlowConstantsWithChan invokes the dms_enterprise.UpdateTaskFlowConstants API asynchronously
func (client *Client) UpdateTaskFlowConstantsWithChan(request *UpdateTaskFlowConstantsRequest) (<-chan *UpdateTaskFlowConstantsResponse, <-chan error) {
	responseChan := make(chan *UpdateTaskFlowConstantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaskFlowConstants(request)
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

// UpdateTaskFlowConstantsWithCallback invokes the dms_enterprise.UpdateTaskFlowConstants API asynchronously
func (client *Client) UpdateTaskFlowConstantsWithCallback(request *UpdateTaskFlowConstantsRequest, callback func(response *UpdateTaskFlowConstantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaskFlowConstantsResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaskFlowConstants(request)
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

// UpdateTaskFlowConstantsRequest is the request struct for api UpdateTaskFlowConstants
type UpdateTaskFlowConstantsRequest struct {
	*requests.RpcRequest
	DagId        requests.Integer                       `position:"Query" name:"DagId"`
	Tid          requests.Integer                       `position:"Query" name:"Tid"`
	DagConstants *[]UpdateTaskFlowConstantsDagConstants `position:"Query" name:"DagConstants"  type:"Json"`
}

// UpdateTaskFlowConstantsDagConstants is a repeated param struct in UpdateTaskFlowConstantsRequest
type UpdateTaskFlowConstantsDagConstants struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UpdateTaskFlowConstantsResponse is the response struct for api UpdateTaskFlowConstants
type UpdateTaskFlowConstantsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTaskFlowConstantsRequest creates a request to invoke UpdateTaskFlowConstants API
func CreateUpdateTaskFlowConstantsRequest() (request *UpdateTaskFlowConstantsRequest) {
	request = &UpdateTaskFlowConstantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "UpdateTaskFlowConstants", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTaskFlowConstantsResponse creates a response to parse from UpdateTaskFlowConstants response
func CreateUpdateTaskFlowConstantsResponse() (response *UpdateTaskFlowConstantsResponse) {
	response = &UpdateTaskFlowConstantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
