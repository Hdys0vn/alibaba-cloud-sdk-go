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

// UpdateDataset invokes the viapi_regen.UpdateDataset API synchronously
func (client *Client) UpdateDataset(request *UpdateDatasetRequest) (response *UpdateDatasetResponse, err error) {
	response = CreateUpdateDatasetResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDatasetWithChan invokes the viapi_regen.UpdateDataset API asynchronously
func (client *Client) UpdateDatasetWithChan(request *UpdateDatasetRequest) (<-chan *UpdateDatasetResponse, <-chan error) {
	responseChan := make(chan *UpdateDatasetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDataset(request)
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

// UpdateDatasetWithCallback invokes the viapi_regen.UpdateDataset API asynchronously
func (client *Client) UpdateDatasetWithCallback(request *UpdateDatasetRequest, callback func(response *UpdateDatasetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDatasetResponse
		var err error
		defer close(result)
		response, err = client.UpdateDataset(request)
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

// UpdateDatasetRequest is the request struct for api UpdateDataset
type UpdateDatasetRequest struct {
	*requests.RpcRequest
	Description string           `position:"Body" name:"Description"`
	Id          requests.Integer `position:"Body" name:"Id"`
	Name        string           `position:"Body" name:"Name"`
}

// UpdateDatasetResponse is the response struct for api UpdateDataset
type UpdateDatasetResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateDatasetRequest creates a request to invoke UpdateDataset API
func CreateUpdateDatasetRequest() (request *UpdateDatasetRequest) {
	request = &UpdateDatasetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "UpdateDataset", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDatasetResponse creates a response to parse from UpdateDataset response
func CreateUpdateDatasetResponse() (response *UpdateDatasetResponse) {
	response = &UpdateDatasetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}