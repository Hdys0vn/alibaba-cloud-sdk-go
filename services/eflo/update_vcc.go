package eflo

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

// UpdateVcc invokes the eflo.UpdateVcc API synchronously
func (client *Client) UpdateVcc(request *UpdateVccRequest) (response *UpdateVccResponse, err error) {
	response = CreateUpdateVccResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVccWithChan invokes the eflo.UpdateVcc API asynchronously
func (client *Client) UpdateVccWithChan(request *UpdateVccRequest) (<-chan *UpdateVccResponse, <-chan error) {
	responseChan := make(chan *UpdateVccResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVcc(request)
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

// UpdateVccWithCallback invokes the eflo.UpdateVcc API asynchronously
func (client *Client) UpdateVccWithCallback(request *UpdateVccRequest, callback func(response *UpdateVccResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVccResponse
		var err error
		defer close(result)
		response, err = client.UpdateVcc(request)
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

// UpdateVccRequest is the request struct for api UpdateVcc
type UpdateVccRequest struct {
	*requests.RpcRequest
	Bandwidth requests.Integer `position:"Body" name:"Bandwidth"`
	OrderId   string           `position:"Body" name:"OrderId"`
	VccName   string           `position:"Body" name:"VccName"`
	VccId     string           `position:"Body" name:"VccId"`
}

// UpdateVccResponse is the response struct for api UpdateVcc
type UpdateVccResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateUpdateVccRequest creates a request to invoke UpdateVcc API
func CreateUpdateVccRequest() (request *UpdateVccRequest) {
	request = &UpdateVccRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "UpdateVcc", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateVccResponse creates a response to parse from UpdateVcc response
func CreateUpdateVccResponse() (response *UpdateVccResponse) {
	response = &UpdateVccResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
