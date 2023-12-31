package eas

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

// CreateServiceAutoScaler invokes the eas.CreateServiceAutoScaler API synchronously
func (client *Client) CreateServiceAutoScaler(request *CreateServiceAutoScalerRequest) (response *CreateServiceAutoScalerResponse, err error) {
	response = CreateCreateServiceAutoScalerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceAutoScalerWithChan invokes the eas.CreateServiceAutoScaler API asynchronously
func (client *Client) CreateServiceAutoScalerWithChan(request *CreateServiceAutoScalerRequest) (<-chan *CreateServiceAutoScalerResponse, <-chan error) {
	responseChan := make(chan *CreateServiceAutoScalerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServiceAutoScaler(request)
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

// CreateServiceAutoScalerWithCallback invokes the eas.CreateServiceAutoScaler API asynchronously
func (client *Client) CreateServiceAutoScalerWithCallback(request *CreateServiceAutoScalerRequest, callback func(response *CreateServiceAutoScalerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceAutoScalerResponse
		var err error
		defer close(result)
		response, err = client.CreateServiceAutoScaler(request)
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

// CreateServiceAutoScalerRequest is the request struct for api CreateServiceAutoScaler
type CreateServiceAutoScalerRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
	Body        string `position:"Body" name:"body"`
}

// CreateServiceAutoScalerResponse is the response struct for api CreateServiceAutoScaler
type CreateServiceAutoScalerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateServiceAutoScalerRequest creates a request to invoke CreateServiceAutoScaler API
func CreateCreateServiceAutoScalerRequest() (request *CreateServiceAutoScalerRequest) {
	request = &CreateServiceAutoScalerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "CreateServiceAutoScaler", "/api/v2/services/[ClusterId]/[ServiceName]/autoscaler", "eas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateServiceAutoScalerResponse creates a response to parse from CreateServiceAutoScaler response
func CreateCreateServiceAutoScalerResponse() (response *CreateServiceAutoScalerResponse) {
	response = &CreateServiceAutoScalerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
