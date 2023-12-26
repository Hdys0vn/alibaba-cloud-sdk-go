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

// UpdateConnector invokes the ververica.UpdateConnector API synchronously
func (client *Client) UpdateConnector(request *UpdateConnectorRequest) (response *UpdateConnectorResponse, err error) {
	response = CreateUpdateConnectorResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConnectorWithChan invokes the ververica.UpdateConnector API asynchronously
func (client *Client) UpdateConnectorWithChan(request *UpdateConnectorRequest) (<-chan *UpdateConnectorResponse, <-chan error) {
	responseChan := make(chan *UpdateConnectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConnector(request)
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

// UpdateConnectorWithCallback invokes the ververica.UpdateConnector API asynchronously
func (client *Client) UpdateConnectorWithCallback(request *UpdateConnectorRequest, callback func(response *UpdateConnectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConnectorResponse
		var err error
		defer close(result)
		response, err = client.UpdateConnector(request)
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

// UpdateConnectorRequest is the request struct for api UpdateConnector
type UpdateConnectorRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Name       string `position:"Path" name:"name"`
	Namespace  string `position:"Path" name:"namespace"`
}

// UpdateConnectorResponse is the response struct for api UpdateConnector
type UpdateConnectorResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateUpdateConnectorRequest creates a request to invoke UpdateConnector API
func CreateUpdateConnectorRequest() (request *UpdateConnectorRequest) {
	request = &UpdateConnectorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "UpdateConnector", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/connectors/[name]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateConnectorResponse creates a response to parse from UpdateConnector response
func CreateUpdateConnectorResponse() (response *UpdateConnectorResponse) {
	response = &UpdateConnectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}