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

// CreateLayer invokes the pairecservice.CreateLayer API synchronously
func (client *Client) CreateLayer(request *CreateLayerRequest) (response *CreateLayerResponse, err error) {
	response = CreateCreateLayerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLayerWithChan invokes the pairecservice.CreateLayer API asynchronously
func (client *Client) CreateLayerWithChan(request *CreateLayerRequest) (<-chan *CreateLayerResponse, <-chan error) {
	responseChan := make(chan *CreateLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLayer(request)
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

// CreateLayerWithCallback invokes the pairecservice.CreateLayer API asynchronously
func (client *Client) CreateLayerWithCallback(request *CreateLayerRequest, callback func(response *CreateLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLayerResponse
		var err error
		defer close(result)
		response, err = client.CreateLayer(request)
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

// CreateLayerRequest is the request struct for api CreateLayer
type CreateLayerRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateLayerResponse is the response struct for api CreateLayer
type CreateLayerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	LayerId   string `json:"LayerId" xml:"LayerId"`
}

// CreateCreateLayerRequest creates a request to invoke CreateLayer API
func CreateCreateLayerRequest() (request *CreateLayerRequest) {
	request = &CreateLayerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateLayer", "/api/v1/layers", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateLayerResponse creates a response to parse from CreateLayer response
func CreateCreateLayerResponse() (response *CreateLayerResponse) {
	response = &CreateLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
