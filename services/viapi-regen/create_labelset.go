package viapi_regen

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

// CreateLabelset invokes the viapi_regen.CreateLabelset API synchronously
func (client *Client) CreateLabelset(request *CreateLabelsetRequest) (response *CreateLabelsetResponse, err error) {
	response = CreateCreateLabelsetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLabelsetWithChan invokes the viapi_regen.CreateLabelset API asynchronously
func (client *Client) CreateLabelsetWithChan(request *CreateLabelsetRequest) (<-chan *CreateLabelsetResponse, <-chan error) {
	responseChan := make(chan *CreateLabelsetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLabelset(request)
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

// CreateLabelsetWithCallback invokes the viapi_regen.CreateLabelset API asynchronously
func (client *Client) CreateLabelsetWithCallback(request *CreateLabelsetRequest, callback func(response *CreateLabelsetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLabelsetResponse
		var err error
		defer close(result)
		response, err = client.CreateLabelset(request)
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

// CreateLabelsetRequest is the request struct for api CreateLabelset
type CreateLabelsetRequest struct {
	*requests.RpcRequest
	Description string           `position:"Body" name:"Description"`
	Type        string           `position:"Body" name:"Type"`
	PreLabelId  requests.Integer `position:"Body" name:"PreLabelId"`
	TagUserList string           `position:"Body" name:"TagUserList"`
	UserOssUrl  string           `position:"Body" name:"UserOssUrl"`
	ObjectKey   string           `position:"Body" name:"ObjectKey"`
	Name        string           `position:"Body" name:"Name"`
	DatasetId   requests.Integer `position:"Body" name:"DatasetId"`
	TagSettings string           `position:"Body" name:"TagSettings"`
}

// CreateLabelsetResponse is the response struct for api CreateLabelset
type CreateLabelsetResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateLabelsetRequest creates a request to invoke CreateLabelset API
func CreateCreateLabelsetRequest() (request *CreateLabelsetRequest) {
	request = &CreateLabelsetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "CreateLabelset", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLabelsetResponse creates a response to parse from CreateLabelset response
func CreateCreateLabelsetResponse() (response *CreateLabelsetResponse) {
	response = &CreateLabelsetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
