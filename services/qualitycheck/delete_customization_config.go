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

// DeleteCustomizationConfig invokes the qualitycheck.DeleteCustomizationConfig API synchronously
func (client *Client) DeleteCustomizationConfig(request *DeleteCustomizationConfigRequest) (response *DeleteCustomizationConfigResponse, err error) {
	response = CreateDeleteCustomizationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCustomizationConfigWithChan invokes the qualitycheck.DeleteCustomizationConfig API asynchronously
func (client *Client) DeleteCustomizationConfigWithChan(request *DeleteCustomizationConfigRequest) (<-chan *DeleteCustomizationConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteCustomizationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCustomizationConfig(request)
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

// DeleteCustomizationConfigWithCallback invokes the qualitycheck.DeleteCustomizationConfig API asynchronously
func (client *Client) DeleteCustomizationConfigWithCallback(request *DeleteCustomizationConfigRequest, callback func(response *DeleteCustomizationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCustomizationConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteCustomizationConfig(request)
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

// DeleteCustomizationConfigRequest is the request struct for api DeleteCustomizationConfig
type DeleteCustomizationConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// DeleteCustomizationConfigResponse is the response struct for api DeleteCustomizationConfig
type DeleteCustomizationConfigResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteCustomizationConfigRequest creates a request to invoke DeleteCustomizationConfig API
func CreateDeleteCustomizationConfigRequest() (request *DeleteCustomizationConfigRequest) {
	request = &DeleteCustomizationConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DeleteCustomizationConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCustomizationConfigResponse creates a response to parse from DeleteCustomizationConfig response
func CreateDeleteCustomizationConfigResponse() (response *DeleteCustomizationConfigResponse) {
	response = &DeleteCustomizationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
