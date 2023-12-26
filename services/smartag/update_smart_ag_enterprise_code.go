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

// UpdateSmartAGEnterpriseCode invokes the smartag.UpdateSmartAGEnterpriseCode API synchronously
func (client *Client) UpdateSmartAGEnterpriseCode(request *UpdateSmartAGEnterpriseCodeRequest) (response *UpdateSmartAGEnterpriseCodeResponse, err error) {
	response = CreateUpdateSmartAGEnterpriseCodeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSmartAGEnterpriseCodeWithChan invokes the smartag.UpdateSmartAGEnterpriseCode API asynchronously
func (client *Client) UpdateSmartAGEnterpriseCodeWithChan(request *UpdateSmartAGEnterpriseCodeRequest) (<-chan *UpdateSmartAGEnterpriseCodeResponse, <-chan error) {
	responseChan := make(chan *UpdateSmartAGEnterpriseCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSmartAGEnterpriseCode(request)
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

// UpdateSmartAGEnterpriseCodeWithCallback invokes the smartag.UpdateSmartAGEnterpriseCode API asynchronously
func (client *Client) UpdateSmartAGEnterpriseCodeWithCallback(request *UpdateSmartAGEnterpriseCodeRequest, callback func(response *UpdateSmartAGEnterpriseCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSmartAGEnterpriseCodeResponse
		var err error
		defer close(result)
		response, err = client.UpdateSmartAGEnterpriseCode(request)
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

// UpdateSmartAGEnterpriseCodeRequest is the request struct for api UpdateSmartAGEnterpriseCode
type UpdateSmartAGEnterpriseCodeRequest struct {
	*requests.RpcRequest
	ClientToken    string           `position:"Query" name:"ClientToken"`
	EnterpriseCode string           `position:"Query" name:"EnterpriseCode"`
	DryRun         requests.Boolean `position:"Query" name:"DryRun"`
	SmartAGId      string           `position:"Query" name:"SmartAGId"`
}

// UpdateSmartAGEnterpriseCodeResponse is the response struct for api UpdateSmartAGEnterpriseCode
type UpdateSmartAGEnterpriseCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSmartAGEnterpriseCodeRequest creates a request to invoke UpdateSmartAGEnterpriseCode API
func CreateUpdateSmartAGEnterpriseCodeRequest() (request *UpdateSmartAGEnterpriseCodeRequest) {
	request = &UpdateSmartAGEnterpriseCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UpdateSmartAGEnterpriseCode", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSmartAGEnterpriseCodeResponse creates a response to parse from UpdateSmartAGEnterpriseCode response
func CreateUpdateSmartAGEnterpriseCodeResponse() (response *UpdateSmartAGEnterpriseCodeResponse) {
	response = &UpdateSmartAGEnterpriseCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
