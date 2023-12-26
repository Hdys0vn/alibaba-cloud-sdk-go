package resourcecenter

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

// GetResourceConfiguration invokes the resourcecenter.GetResourceConfiguration API synchronously
func (client *Client) GetResourceConfiguration(request *GetResourceConfigurationRequest) (response *GetResourceConfigurationResponse, err error) {
	response = CreateGetResourceConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceConfigurationWithChan invokes the resourcecenter.GetResourceConfiguration API asynchronously
func (client *Client) GetResourceConfigurationWithChan(request *GetResourceConfigurationRequest) (<-chan *GetResourceConfigurationResponse, <-chan error) {
	responseChan := make(chan *GetResourceConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceConfiguration(request)
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

// GetResourceConfigurationWithCallback invokes the resourcecenter.GetResourceConfiguration API asynchronously
func (client *Client) GetResourceConfigurationWithCallback(request *GetResourceConfigurationRequest, callback func(response *GetResourceConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceConfigurationResponse
		var err error
		defer close(result)
		response, err = client.GetResourceConfiguration(request)
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

// GetResourceConfigurationRequest is the request struct for api GetResourceConfiguration
type GetResourceConfigurationRequest struct {
	*requests.RpcRequest
	ResourceId       string `position:"Query" name:"ResourceId"`
	ResourceType     string `position:"Query" name:"ResourceType"`
	ResourceRegionId string `position:"Query" name:"ResourceRegionId"`
}

// GetResourceConfigurationResponse is the response struct for api GetResourceConfiguration
type GetResourceConfigurationResponse struct {
	*responses.BaseResponse
}

// CreateGetResourceConfigurationRequest creates a request to invoke GetResourceConfiguration API
func CreateGetResourceConfigurationRequest() (request *GetResourceConfigurationRequest) {
	request = &GetResourceConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "GetResourceConfiguration", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResourceConfigurationResponse creates a response to parse from GetResourceConfiguration response
func CreateGetResourceConfigurationResponse() (response *GetResourceConfigurationResponse) {
	response = &GetResourceConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}