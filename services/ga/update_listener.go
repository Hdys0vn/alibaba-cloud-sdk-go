package ga

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

// UpdateListener invokes the ga.UpdateListener API synchronously
func (client *Client) UpdateListener(request *UpdateListenerRequest) (response *UpdateListenerResponse, err error) {
	response = CreateUpdateListenerResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateListenerWithChan invokes the ga.UpdateListener API asynchronously
func (client *Client) UpdateListenerWithChan(request *UpdateListenerRequest) (<-chan *UpdateListenerResponse, <-chan error) {
	responseChan := make(chan *UpdateListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateListener(request)
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

// UpdateListenerWithCallback invokes the ga.UpdateListener API asynchronously
func (client *Client) UpdateListenerWithCallback(request *UpdateListenerRequest, callback func(response *UpdateListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateListenerResponse
		var err error
		defer close(result)
		response, err = client.UpdateListener(request)
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

// UpdateListenerRequest is the request struct for api UpdateListener
type UpdateListenerRequest struct {
	*requests.RpcRequest
	ClientToken    string                        `position:"Query" name:"ClientToken"`
	Description    string                        `position:"Query" name:"Description"`
	BackendPorts   *[]UpdateListenerBackendPorts `position:"Query" name:"BackendPorts"  type:"Repeated"`
	ListenerId     string                        `position:"Query" name:"ListenerId"`
	Protocol       string                        `position:"Query" name:"Protocol"`
	ProxyProtocol  string                        `position:"Query" name:"ProxyProtocol"`
	PortRanges     *[]UpdateListenerPortRanges   `position:"Query" name:"PortRanges"  type:"Repeated"`
	Certificates   *[]UpdateListenerCertificates `position:"Query" name:"Certificates"  type:"Repeated"`
	Name           string                        `position:"Query" name:"Name"`
	ClientAffinity string                        `position:"Query" name:"ClientAffinity"`
}

// UpdateListenerBackendPorts is a repeated param struct in UpdateListenerRequest
type UpdateListenerBackendPorts struct {
	FromPort string `name:"FromPort"`
	ToPort   string `name:"ToPort"`
}

// UpdateListenerPortRanges is a repeated param struct in UpdateListenerRequest
type UpdateListenerPortRanges struct {
	FromPort string `name:"FromPort"`
	ToPort   string `name:"ToPort"`
}

// UpdateListenerCertificates is a repeated param struct in UpdateListenerRequest
type UpdateListenerCertificates struct {
	Id string `name:"Id"`
}

// UpdateListenerResponse is the response struct for api UpdateListener
type UpdateListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateListenerRequest creates a request to invoke UpdateListener API
func CreateUpdateListenerRequest() (request *UpdateListenerRequest) {
	request = &UpdateListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "UpdateListener", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateListenerResponse creates a response to parse from UpdateListener response
func CreateUpdateListenerResponse() (response *UpdateListenerResponse) {
	response = &UpdateListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
