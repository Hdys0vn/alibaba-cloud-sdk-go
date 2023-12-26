package elasticsearch

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

// UpdateAdminPassword invokes the elasticsearch.UpdateAdminPassword API synchronously
func (client *Client) UpdateAdminPassword(request *UpdateAdminPasswordRequest) (response *UpdateAdminPasswordResponse, err error) {
	response = CreateUpdateAdminPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAdminPasswordWithChan invokes the elasticsearch.UpdateAdminPassword API asynchronously
func (client *Client) UpdateAdminPasswordWithChan(request *UpdateAdminPasswordRequest) (<-chan *UpdateAdminPasswordResponse, <-chan error) {
	responseChan := make(chan *UpdateAdminPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAdminPassword(request)
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

// UpdateAdminPasswordWithCallback invokes the elasticsearch.UpdateAdminPassword API asynchronously
func (client *Client) UpdateAdminPasswordWithCallback(request *UpdateAdminPasswordRequest, callback func(response *UpdateAdminPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAdminPasswordResponse
		var err error
		defer close(result)
		response, err = client.UpdateAdminPassword(request)
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

// UpdateAdminPasswordRequest is the request struct for api UpdateAdminPassword
type UpdateAdminPasswordRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UpdateAdminPasswordResponse is the response struct for api UpdateAdminPassword
type UpdateAdminPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAdminPasswordRequest creates a request to invoke UpdateAdminPassword API
func CreateUpdateAdminPasswordRequest() (request *UpdateAdminPasswordRequest) {
	request = &UpdateAdminPasswordRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateAdminPassword", "/openapi/instances/[InstanceId]/admin-pwd", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAdminPasswordResponse creates a response to parse from UpdateAdminPassword response
func CreateUpdateAdminPasswordResponse() (response *UpdateAdminPasswordResponse) {
	response = &UpdateAdminPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
