package unimkt

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

// CreateSlot invokes the unimkt.CreateSlot API synchronously
func (client *Client) CreateSlot(request *CreateSlotRequest) (response *CreateSlotResponse, err error) {
	response = CreateCreateSlotResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSlotWithChan invokes the unimkt.CreateSlot API asynchronously
func (client *Client) CreateSlotWithChan(request *CreateSlotRequest) (<-chan *CreateSlotResponse, <-chan error) {
	responseChan := make(chan *CreateSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSlot(request)
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

// CreateSlotWithCallback invokes the unimkt.CreateSlot API asynchronously
func (client *Client) CreateSlotWithCallback(request *CreateSlotRequest, callback func(response *CreateSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSlotResponse
		var err error
		defer close(result)
		response, err = client.CreateSlot(request)
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

// CreateSlotRequest is the request struct for api CreateSlot
type CreateSlotRequest struct {
	*requests.RpcRequest
	Business         string `position:"Query" name:"Business"`
	ClientToken      string `position:"Query" name:"ClientToken"`
	UserId           string `position:"Query" name:"UserId"`
	OriginSiteUserId string `position:"Query" name:"OriginSiteUserId"`
	Environment      string `position:"Query" name:"Environment"`
	AppName          string `position:"Query" name:"AppName"`
	TenantId         string `position:"Query" name:"TenantId"`
	UserSite         string `position:"Query" name:"UserSite"`
	AdSlot           string `position:"Body" name:"AdSlot"`
}

// CreateSlotResponse is the response struct for api CreateSlot
type CreateSlotResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateCreateSlotRequest creates a request to invoke CreateSlot API
func CreateCreateSlotRequest() (request *CreateSlotRequest) {
	request = &CreateSlotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "CreateSlot", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSlotResponse creates a response to parse from CreateSlot response
func CreateCreateSlotResponse() (response *CreateSlotResponse) {
	response = &CreateSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}