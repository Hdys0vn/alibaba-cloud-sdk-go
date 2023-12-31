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

// AttachDdosToAccelerator invokes the ga.AttachDdosToAccelerator API synchronously
func (client *Client) AttachDdosToAccelerator(request *AttachDdosToAcceleratorRequest) (response *AttachDdosToAcceleratorResponse, err error) {
	response = CreateAttachDdosToAcceleratorResponse()
	err = client.DoAction(request, response)
	return
}

// AttachDdosToAcceleratorWithChan invokes the ga.AttachDdosToAccelerator API asynchronously
func (client *Client) AttachDdosToAcceleratorWithChan(request *AttachDdosToAcceleratorRequest) (<-chan *AttachDdosToAcceleratorResponse, <-chan error) {
	responseChan := make(chan *AttachDdosToAcceleratorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachDdosToAccelerator(request)
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

// AttachDdosToAcceleratorWithCallback invokes the ga.AttachDdosToAccelerator API asynchronously
func (client *Client) AttachDdosToAcceleratorWithCallback(request *AttachDdosToAcceleratorRequest, callback func(response *AttachDdosToAcceleratorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachDdosToAcceleratorResponse
		var err error
		defer close(result)
		response, err = client.AttachDdosToAccelerator(request)
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

// AttachDdosToAcceleratorRequest is the request struct for api AttachDdosToAccelerator
type AttachDdosToAcceleratorRequest struct {
	*requests.RpcRequest
	DdosId        string `position:"Query" name:"DdosId"`
	DdosRegionId  string `position:"Query" name:"DdosRegionId"`
	AcceleratorId string `position:"Query" name:"AcceleratorId"`
}

// AttachDdosToAcceleratorResponse is the response struct for api AttachDdosToAccelerator
type AttachDdosToAcceleratorResponse struct {
	*responses.BaseResponse
	DdosId    string `json:"DdosId" xml:"DdosId"`
	GaId      string `json:"GaId" xml:"GaId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDdosToAcceleratorRequest creates a request to invoke AttachDdosToAccelerator API
func CreateAttachDdosToAcceleratorRequest() (request *AttachDdosToAcceleratorRequest) {
	request = &AttachDdosToAcceleratorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "AttachDdosToAccelerator", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachDdosToAcceleratorResponse creates a response to parse from AttachDdosToAccelerator response
func CreateAttachDdosToAcceleratorResponse() (response *AttachDdosToAcceleratorResponse) {
	response = &AttachDdosToAcceleratorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
