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

// UpdateTableModelInfo invokes the dataworks_public.UpdateTableModelInfo API synchronously
func (client *Client) UpdateTableModelInfo(request *UpdateTableModelInfoRequest) (response *UpdateTableModelInfoResponse, err error) {
	response = CreateUpdateTableModelInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTableModelInfoWithChan invokes the dataworks_public.UpdateTableModelInfo API asynchronously
func (client *Client) UpdateTableModelInfoWithChan(request *UpdateTableModelInfoRequest) (<-chan *UpdateTableModelInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateTableModelInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTableModelInfo(request)
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

// UpdateTableModelInfoWithCallback invokes the dataworks_public.UpdateTableModelInfo API asynchronously
func (client *Client) UpdateTableModelInfoWithCallback(request *UpdateTableModelInfoRequest, callback func(response *UpdateTableModelInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTableModelInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateTableModelInfo(request)
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

// UpdateTableModelInfoRequest is the request struct for api UpdateTableModelInfo
type UpdateTableModelInfoRequest struct {
	*requests.RpcRequest
	LevelType          requests.Integer `position:"Query" name:"LevelType"`
	SecondLevelThemeId requests.Integer `position:"Query" name:"SecondLevelThemeId"`
	TableGuid          string           `position:"Query" name:"TableGuid"`
	LevelId            requests.Integer `position:"Query" name:"LevelId"`
	FirstLevelThemeId  requests.Integer `position:"Query" name:"FirstLevelThemeId"`
}

// UpdateTableModelInfoResponse is the response struct for api UpdateTableModelInfo
type UpdateTableModelInfoResponse struct {
	*responses.BaseResponse
	UpdateResult bool   `json:"UpdateResult" xml:"UpdateResult"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateTableModelInfoRequest creates a request to invoke UpdateTableModelInfo API
func CreateUpdateTableModelInfoRequest() (request *UpdateTableModelInfoRequest) {
	request = &UpdateTableModelInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateTableModelInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTableModelInfoResponse creates a response to parse from UpdateTableModelInfo response
func CreateUpdateTableModelInfoResponse() (response *UpdateTableModelInfoResponse) {
	response = &UpdateTableModelInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
