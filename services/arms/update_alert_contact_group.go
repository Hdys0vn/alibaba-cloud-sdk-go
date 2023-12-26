package arms

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

// UpdateAlertContactGroup invokes the arms.UpdateAlertContactGroup API synchronously
func (client *Client) UpdateAlertContactGroup(request *UpdateAlertContactGroupRequest) (response *UpdateAlertContactGroupResponse, err error) {
	response = CreateUpdateAlertContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAlertContactGroupWithChan invokes the arms.UpdateAlertContactGroup API asynchronously
func (client *Client) UpdateAlertContactGroupWithChan(request *UpdateAlertContactGroupRequest) (<-chan *UpdateAlertContactGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateAlertContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAlertContactGroup(request)
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

// UpdateAlertContactGroupWithCallback invokes the arms.UpdateAlertContactGroup API asynchronously
func (client *Client) UpdateAlertContactGroupWithCallback(request *UpdateAlertContactGroupRequest, callback func(response *UpdateAlertContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAlertContactGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateAlertContactGroup(request)
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

// UpdateAlertContactGroupRequest is the request struct for api UpdateAlertContactGroup
type UpdateAlertContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupId   requests.Integer `position:"Query" name:"ContactGroupId"`
	ContactGroupName string           `position:"Query" name:"ContactGroupName"`
	ProxyUserId      string           `position:"Query" name:"ProxyUserId"`
	ContactIds       string           `position:"Query" name:"ContactIds"`
}

// UpdateAlertContactGroupResponse is the response struct for api UpdateAlertContactGroup
type UpdateAlertContactGroupResponse struct {
	*responses.BaseResponse
	UpdateAlertContactGroupIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	RequestId                        string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAlertContactGroupRequest creates a request to invoke UpdateAlertContactGroup API
func CreateUpdateAlertContactGroupRequest() (request *UpdateAlertContactGroupRequest) {
	request = &UpdateAlertContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateAlertContactGroup", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAlertContactGroupResponse creates a response to parse from UpdateAlertContactGroup response
func CreateUpdateAlertContactGroupResponse() (response *UpdateAlertContactGroupResponse) {
	response = &UpdateAlertContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}