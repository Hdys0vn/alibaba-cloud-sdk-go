package gpdb

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

// ModifyVectorConfiguration invokes the gpdb.ModifyVectorConfiguration API synchronously
func (client *Client) ModifyVectorConfiguration(request *ModifyVectorConfigurationRequest) (response *ModifyVectorConfigurationResponse, err error) {
	response = CreateModifyVectorConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVectorConfigurationWithChan invokes the gpdb.ModifyVectorConfiguration API asynchronously
func (client *Client) ModifyVectorConfigurationWithChan(request *ModifyVectorConfigurationRequest) (<-chan *ModifyVectorConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyVectorConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVectorConfiguration(request)
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

// ModifyVectorConfigurationWithCallback invokes the gpdb.ModifyVectorConfiguration API asynchronously
func (client *Client) ModifyVectorConfigurationWithCallback(request *ModifyVectorConfigurationRequest, callback func(response *ModifyVectorConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVectorConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyVectorConfiguration(request)
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

// ModifyVectorConfigurationRequest is the request struct for api ModifyVectorConfiguration
type ModifyVectorConfigurationRequest struct {
	*requests.RpcRequest
	DBInstanceId              string           `position:"Query" name:"DBInstanceId"`
	VectorConfigurationStatus string           `position:"Query" name:"VectorConfigurationStatus"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyVectorConfigurationResponse is the response struct for api ModifyVectorConfiguration
type ModifyVectorConfigurationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Status       bool   `json:"Status" xml:"Status"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateModifyVectorConfigurationRequest creates a request to invoke ModifyVectorConfiguration API
func CreateModifyVectorConfigurationRequest() (request *ModifyVectorConfigurationRequest) {
	request = &ModifyVectorConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifyVectorConfiguration", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyVectorConfigurationResponse creates a response to parse from ModifyVectorConfiguration response
func CreateModifyVectorConfigurationResponse() (response *ModifyVectorConfigurationResponse) {
	response = &ModifyVectorConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}