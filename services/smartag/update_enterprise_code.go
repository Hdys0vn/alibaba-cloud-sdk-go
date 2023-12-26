package smartag

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

// UpdateEnterpriseCode invokes the smartag.UpdateEnterpriseCode API synchronously
func (client *Client) UpdateEnterpriseCode(request *UpdateEnterpriseCodeRequest) (response *UpdateEnterpriseCodeResponse, err error) {
	response = CreateUpdateEnterpriseCodeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEnterpriseCodeWithChan invokes the smartag.UpdateEnterpriseCode API asynchronously
func (client *Client) UpdateEnterpriseCodeWithChan(request *UpdateEnterpriseCodeRequest) (<-chan *UpdateEnterpriseCodeResponse, <-chan error) {
	responseChan := make(chan *UpdateEnterpriseCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEnterpriseCode(request)
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

// UpdateEnterpriseCodeWithCallback invokes the smartag.UpdateEnterpriseCode API asynchronously
func (client *Client) UpdateEnterpriseCodeWithCallback(request *UpdateEnterpriseCodeRequest, callback func(response *UpdateEnterpriseCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEnterpriseCodeResponse
		var err error
		defer close(result)
		response, err = client.UpdateEnterpriseCode(request)
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

// UpdateEnterpriseCodeRequest is the request struct for api UpdateEnterpriseCode
type UpdateEnterpriseCodeRequest struct {
	*requests.RpcRequest
	ClientToken    string           `position:"Query" name:"ClientToken"`
	EnterpriseCode string           `position:"Query" name:"EnterpriseCode"`
	IsDefault      requests.Boolean `position:"Query" name:"IsDefault"`
	DryRun         requests.Boolean `position:"Query" name:"DryRun"`
}

// UpdateEnterpriseCodeResponse is the response struct for api UpdateEnterpriseCode
type UpdateEnterpriseCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateEnterpriseCodeRequest creates a request to invoke UpdateEnterpriseCode API
func CreateUpdateEnterpriseCodeRequest() (request *UpdateEnterpriseCodeRequest) {
	request = &UpdateEnterpriseCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UpdateEnterpriseCode", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateEnterpriseCodeResponse creates a response to parse from UpdateEnterpriseCode response
func CreateUpdateEnterpriseCodeResponse() (response *UpdateEnterpriseCodeResponse) {
	response = &UpdateEnterpriseCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
