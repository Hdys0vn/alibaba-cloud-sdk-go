package dataworks_public

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

// UpdateTableAddColumn invokes the dataworks_public.UpdateTableAddColumn API synchronously
func (client *Client) UpdateTableAddColumn(request *UpdateTableAddColumnRequest) (response *UpdateTableAddColumnResponse, err error) {
	response = CreateUpdateTableAddColumnResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTableAddColumnWithChan invokes the dataworks_public.UpdateTableAddColumn API asynchronously
func (client *Client) UpdateTableAddColumnWithChan(request *UpdateTableAddColumnRequest) (<-chan *UpdateTableAddColumnResponse, <-chan error) {
	responseChan := make(chan *UpdateTableAddColumnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTableAddColumn(request)
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

// UpdateTableAddColumnWithCallback invokes the dataworks_public.UpdateTableAddColumn API asynchronously
func (client *Client) UpdateTableAddColumnWithCallback(request *UpdateTableAddColumnRequest, callback func(response *UpdateTableAddColumnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTableAddColumnResponse
		var err error
		defer close(result)
		response, err = client.UpdateTableAddColumn(request)
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

// UpdateTableAddColumnRequest is the request struct for api UpdateTableAddColumn
type UpdateTableAddColumnRequest struct {
	*requests.RpcRequest
	Column    *[]UpdateTableAddColumnColumn `position:"Body" name:"Column"  type:"Repeated"`
	TableGuid string                        `position:"Query" name:"TableGuid"`
}

// UpdateTableAddColumnColumn is a repeated param struct in UpdateTableAddColumnRequest
type UpdateTableAddColumnColumn struct {
	ColumnNameCn string `name:"ColumnNameCn"`
	Comment      string `name:"Comment"`
	ColumnName   string `name:"ColumnName"`
	ColumnType   string `name:"ColumnType"`
}

// UpdateTableAddColumnResponse is the response struct for api UpdateTableAddColumn
type UpdateTableAddColumnResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	TaskInfo  TaskInfo `json:"TaskInfo" xml:"TaskInfo"`
}

// CreateUpdateTableAddColumnRequest creates a request to invoke UpdateTableAddColumn API
func CreateUpdateTableAddColumnRequest() (request *UpdateTableAddColumnRequest) {
	request = &UpdateTableAddColumnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateTableAddColumn", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTableAddColumnResponse creates a response to parse from UpdateTableAddColumn response
func CreateUpdateTableAddColumnResponse() (response *UpdateTableAddColumnResponse) {
	response = &UpdateTableAddColumnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}