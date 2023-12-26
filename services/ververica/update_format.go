package ververica

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

// UpdateFormat invokes the ververica.UpdateFormat API synchronously
func (client *Client) UpdateFormat(request *UpdateFormatRequest) (response *UpdateFormatResponse, err error) {
	response = CreateUpdateFormatResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFormatWithChan invokes the ververica.UpdateFormat API asynchronously
func (client *Client) UpdateFormatWithChan(request *UpdateFormatRequest) (<-chan *UpdateFormatResponse, <-chan error) {
	responseChan := make(chan *UpdateFormatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFormat(request)
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

// UpdateFormatWithCallback invokes the ververica.UpdateFormat API asynchronously
func (client *Client) UpdateFormatWithCallback(request *UpdateFormatRequest, callback func(response *UpdateFormatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFormatResponse
		var err error
		defer close(result)
		response, err = client.UpdateFormat(request)
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

// UpdateFormatRequest is the request struct for api UpdateFormat
type UpdateFormatRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Name       string `position:"Path" name:"name"`
	Namespace  string `position:"Path" name:"namespace"`
}

// UpdateFormatResponse is the response struct for api UpdateFormat
type UpdateFormatResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateUpdateFormatRequest creates a request to invoke UpdateFormat API
func CreateUpdateFormatRequest() (request *UpdateFormatRequest) {
	request = &UpdateFormatRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "UpdateFormat", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/formats/[name]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateFormatResponse creates a response to parse from UpdateFormat response
func CreateUpdateFormatResponse() (response *UpdateFormatResponse) {
	response = &UpdateFormatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}