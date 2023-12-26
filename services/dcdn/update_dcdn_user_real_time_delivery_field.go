package dcdn

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

// UpdateDcdnUserRealTimeDeliveryField invokes the dcdn.UpdateDcdnUserRealTimeDeliveryField API synchronously
func (client *Client) UpdateDcdnUserRealTimeDeliveryField(request *UpdateDcdnUserRealTimeDeliveryFieldRequest) (response *UpdateDcdnUserRealTimeDeliveryFieldResponse, err error) {
	response = CreateUpdateDcdnUserRealTimeDeliveryFieldResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDcdnUserRealTimeDeliveryFieldWithChan invokes the dcdn.UpdateDcdnUserRealTimeDeliveryField API asynchronously
func (client *Client) UpdateDcdnUserRealTimeDeliveryFieldWithChan(request *UpdateDcdnUserRealTimeDeliveryFieldRequest) (<-chan *UpdateDcdnUserRealTimeDeliveryFieldResponse, <-chan error) {
	responseChan := make(chan *UpdateDcdnUserRealTimeDeliveryFieldResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDcdnUserRealTimeDeliveryField(request)
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

// UpdateDcdnUserRealTimeDeliveryFieldWithCallback invokes the dcdn.UpdateDcdnUserRealTimeDeliveryField API asynchronously
func (client *Client) UpdateDcdnUserRealTimeDeliveryFieldWithCallback(request *UpdateDcdnUserRealTimeDeliveryFieldRequest, callback func(response *UpdateDcdnUserRealTimeDeliveryFieldResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDcdnUserRealTimeDeliveryFieldResponse
		var err error
		defer close(result)
		response, err = client.UpdateDcdnUserRealTimeDeliveryField(request)
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

// UpdateDcdnUserRealTimeDeliveryFieldRequest is the request struct for api UpdateDcdnUserRealTimeDeliveryField
type UpdateDcdnUserRealTimeDeliveryFieldRequest struct {
	*requests.RpcRequest
	Fields       string `position:"Query" name:"Fields"`
	BusinessType string `position:"Query" name:"BusinessType"`
}

// UpdateDcdnUserRealTimeDeliveryFieldResponse is the response struct for api UpdateDcdnUserRealTimeDeliveryField
type UpdateDcdnUserRealTimeDeliveryFieldResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDcdnUserRealTimeDeliveryFieldRequest creates a request to invoke UpdateDcdnUserRealTimeDeliveryField API
func CreateUpdateDcdnUserRealTimeDeliveryFieldRequest() (request *UpdateDcdnUserRealTimeDeliveryFieldRequest) {
	request = &UpdateDcdnUserRealTimeDeliveryFieldRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "UpdateDcdnUserRealTimeDeliveryField", "", "")
	request.Method = requests.GET
	return
}

// CreateUpdateDcdnUserRealTimeDeliveryFieldResponse creates a response to parse from UpdateDcdnUserRealTimeDeliveryField response
func CreateUpdateDcdnUserRealTimeDeliveryFieldResponse() (response *UpdateDcdnUserRealTimeDeliveryFieldResponse) {
	response = &UpdateDcdnUserRealTimeDeliveryFieldResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
