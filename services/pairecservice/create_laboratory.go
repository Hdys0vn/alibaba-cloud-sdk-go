package pairecservice

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

// CreateLaboratory invokes the pairecservice.CreateLaboratory API synchronously
func (client *Client) CreateLaboratory(request *CreateLaboratoryRequest) (response *CreateLaboratoryResponse, err error) {
	response = CreateCreateLaboratoryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLaboratoryWithChan invokes the pairecservice.CreateLaboratory API asynchronously
func (client *Client) CreateLaboratoryWithChan(request *CreateLaboratoryRequest) (<-chan *CreateLaboratoryResponse, <-chan error) {
	responseChan := make(chan *CreateLaboratoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLaboratory(request)
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

// CreateLaboratoryWithCallback invokes the pairecservice.CreateLaboratory API asynchronously
func (client *Client) CreateLaboratoryWithCallback(request *CreateLaboratoryRequest, callback func(response *CreateLaboratoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLaboratoryResponse
		var err error
		defer close(result)
		response, err = client.CreateLaboratory(request)
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

// CreateLaboratoryRequest is the request struct for api CreateLaboratory
type CreateLaboratoryRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateLaboratoryResponse is the response struct for api CreateLaboratory
type CreateLaboratoryResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	LaboratoryId string `json:"LaboratoryId" xml:"LaboratoryId"`
}

// CreateCreateLaboratoryRequest creates a request to invoke CreateLaboratory API
func CreateCreateLaboratoryRequest() (request *CreateLaboratoryRequest) {
	request = &CreateLaboratoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateLaboratory", "/api/v1/laboratories", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateLaboratoryResponse creates a response to parse from CreateLaboratory response
func CreateCreateLaboratoryResponse() (response *CreateLaboratoryResponse) {
	response = &CreateLaboratoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}