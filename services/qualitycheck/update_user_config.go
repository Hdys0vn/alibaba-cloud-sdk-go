package qualitycheck

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

// UpdateUserConfig invokes the qualitycheck.UpdateUserConfig API synchronously
func (client *Client) UpdateUserConfig(request *UpdateUserConfigRequest) (response *UpdateUserConfigResponse, err error) {
	response = CreateUpdateUserConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserConfigWithChan invokes the qualitycheck.UpdateUserConfig API asynchronously
func (client *Client) UpdateUserConfigWithChan(request *UpdateUserConfigRequest) (<-chan *UpdateUserConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateUserConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserConfig(request)
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

// UpdateUserConfigWithCallback invokes the qualitycheck.UpdateUserConfig API asynchronously
func (client *Client) UpdateUserConfigWithCallback(request *UpdateUserConfigRequest, callback func(response *UpdateUserConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserConfig(request)
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

// UpdateUserConfigRequest is the request struct for api UpdateUserConfig
type UpdateUserConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UpdateUserConfigResponse is the response struct for api UpdateUserConfig
type UpdateUserConfigResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateUserConfigRequest creates a request to invoke UpdateUserConfig API
func CreateUpdateUserConfigRequest() (request *UpdateUserConfigRequest) {
	request = &UpdateUserConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateUserConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateUserConfigResponse creates a response to parse from UpdateUserConfig response
func CreateUpdateUserConfigResponse() (response *UpdateUserConfigResponse) {
	response = &UpdateUserConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
