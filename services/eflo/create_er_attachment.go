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

// CreateErAttachment invokes the eflo.CreateErAttachment API synchronously
func (client *Client) CreateErAttachment(request *CreateErAttachmentRequest) (response *CreateErAttachmentResponse, err error) {
	response = CreateCreateErAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateErAttachmentWithChan invokes the eflo.CreateErAttachment API asynchronously
func (client *Client) CreateErAttachmentWithChan(request *CreateErAttachmentRequest) (<-chan *CreateErAttachmentResponse, <-chan error) {
	responseChan := make(chan *CreateErAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateErAttachment(request)
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

// CreateErAttachmentWithCallback invokes the eflo.CreateErAttachment API asynchronously
func (client *Client) CreateErAttachmentWithCallback(request *CreateErAttachmentRequest, callback func(response *CreateErAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateErAttachmentResponse
		var err error
		defer close(result)
		response, err = client.CreateErAttachment(request)
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

// CreateErAttachmentRequest is the request struct for api CreateErAttachment
type CreateErAttachmentRequest struct {
	*requests.RpcRequest
	ResourceTenantId    string           `position:"Body" name:"ResourceTenantId"`
	AutoReceiveAllRoute requests.Boolean `position:"Body" name:"AutoReceiveAllRoute"`
	InstanceType        string           `position:"Body" name:"InstanceType"`
	ErAttachmentName    string           `position:"Body" name:"ErAttachmentName"`
	ErId                string           `position:"Body" name:"ErId"`
	InstanceId          string           `position:"Body" name:"InstanceId"`
}

// CreateErAttachmentResponse is the response struct for api CreateErAttachment
type CreateErAttachmentResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateCreateErAttachmentRequest creates a request to invoke CreateErAttachment API
func CreateCreateErAttachmentRequest() (request *CreateErAttachmentRequest) {
	request = &CreateErAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "CreateErAttachment", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateErAttachmentResponse creates a response to parse from CreateErAttachment response
func CreateCreateErAttachmentResponse() (response *CreateErAttachmentResponse) {
	response = &CreateErAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}