package edas

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

// GetJvmConfiguration invokes the edas.GetJvmConfiguration API synchronously
func (client *Client) GetJvmConfiguration(request *GetJvmConfigurationRequest) (response *GetJvmConfigurationResponse, err error) {
	response = CreateGetJvmConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// GetJvmConfigurationWithChan invokes the edas.GetJvmConfiguration API asynchronously
func (client *Client) GetJvmConfigurationWithChan(request *GetJvmConfigurationRequest) (<-chan *GetJvmConfigurationResponse, <-chan error) {
	responseChan := make(chan *GetJvmConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJvmConfiguration(request)
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

// GetJvmConfigurationWithCallback invokes the edas.GetJvmConfiguration API asynchronously
func (client *Client) GetJvmConfigurationWithCallback(request *GetJvmConfigurationRequest, callback func(response *GetJvmConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJvmConfigurationResponse
		var err error
		defer close(result)
		response, err = client.GetJvmConfiguration(request)
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

// GetJvmConfigurationRequest is the request struct for api GetJvmConfiguration
type GetJvmConfigurationRequest struct {
	*requests.RoaRequest
	AppId   string `position:"Query" name:"AppId"`
	GroupId string `position:"Query" name:"GroupId"`
}

// GetJvmConfigurationResponse is the response struct for api GetJvmConfiguration
type GetJvmConfigurationResponse struct {
	*responses.BaseResponse
	Code             int              `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	JvmConfiguration JvmConfiguration `json:"JvmConfiguration" xml:"JvmConfiguration"`
}

// CreateGetJvmConfigurationRequest creates a request to invoke GetJvmConfiguration API
func CreateGetJvmConfigurationRequest() (request *GetJvmConfigurationRequest) {
	request = &GetJvmConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetJvmConfiguration", "/pop/v5/app/app_jvm_config", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetJvmConfigurationResponse creates a response to parse from GetJvmConfiguration response
func CreateGetJvmConfigurationResponse() (response *GetJvmConfigurationResponse) {
	response = &GetJvmConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
