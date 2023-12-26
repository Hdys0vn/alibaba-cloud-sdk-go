package fnf

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

// UpdateFlow invokes the fnf.UpdateFlow API synchronously
func (client *Client) UpdateFlow(request *UpdateFlowRequest) (response *UpdateFlowResponse, err error) {
	response = CreateUpdateFlowResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFlowWithChan invokes the fnf.UpdateFlow API asynchronously
func (client *Client) UpdateFlowWithChan(request *UpdateFlowRequest) (<-chan *UpdateFlowResponse, <-chan error) {
	responseChan := make(chan *UpdateFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFlow(request)
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

// UpdateFlowWithCallback invokes the fnf.UpdateFlow API asynchronously
func (client *Client) UpdateFlowWithCallback(request *UpdateFlowRequest, callback func(response *UpdateFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFlowResponse
		var err error
		defer close(result)
		response, err = client.UpdateFlow(request)
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

// UpdateFlowRequest is the request struct for api UpdateFlow
type UpdateFlowRequest struct {
	*requests.RpcRequest
	Description             string `position:"Body" name:"Description"`
	Type                    string `position:"Body" name:"Type"`
	RequestId               string `position:"Query" name:"RequestId"`
	RoleArn                 string `position:"Body" name:"RoleArn"`
	Name                    string `position:"Body" name:"Name"`
	Definition              string `position:"Body" name:"Definition"`
	ExternalStorageLocation string `position:"Body" name:"ExternalStorageLocation"`
}

// UpdateFlowResponse is the response struct for api UpdateFlow
type UpdateFlowResponse struct {
	*responses.BaseResponse
	Type                    string `json:"Type" xml:"Type"`
	Definition              string `json:"Definition" xml:"Definition"`
	RoleArn                 string `json:"RoleArn" xml:"RoleArn"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
	Description             string `json:"Description" xml:"Description"`
	ExternalStorageLocation string `json:"ExternalStorageLocation" xml:"ExternalStorageLocation"`
	Name                    string `json:"Name" xml:"Name"`
	CreatedTime             string `json:"CreatedTime" xml:"CreatedTime"`
	LastModifiedTime        string `json:"LastModifiedTime" xml:"LastModifiedTime"`
	Id                      string `json:"Id" xml:"Id"`
}

// CreateUpdateFlowRequest creates a request to invoke UpdateFlow API
func CreateUpdateFlowRequest() (request *UpdateFlowRequest) {
	request = &UpdateFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "UpdateFlow", "fnf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateFlowResponse creates a response to parse from UpdateFlow response
func CreateUpdateFlowResponse() (response *UpdateFlowResponse) {
	response = &UpdateFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
