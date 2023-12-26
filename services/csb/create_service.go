package csb

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

// CreateService invokes the csb.CreateService API synchronously
// api document: https://help.aliyun.com/api/csb/createservice.html
func (client *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	response = CreateCreateServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceWithChan invokes the csb.CreateService API asynchronously
// api document: https://help.aliyun.com/api/csb/createservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateServiceWithChan(request *CreateServiceRequest) (<-chan *CreateServiceResponse, <-chan error) {
	responseChan := make(chan *CreateServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateService(request)
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

// CreateServiceWithCallback invokes the csb.CreateService API asynchronously
// api document: https://help.aliyun.com/api/csb/createservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateServiceWithCallback(request *CreateServiceRequest, callback func(response *CreateServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateService(request)
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

// CreateServiceRequest is the request struct for api CreateService
type CreateServiceRequest struct {
	*requests.RpcRequest
	Data  string           `position:"Body" name:"Data"`
	CsbId requests.Integer `position:"Query" name:"CsbId"`
}

// CreateServiceResponse is the response struct for api CreateService
type CreateServiceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateServiceRequest creates a request to invoke CreateService API
func CreateCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "CreateService", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateServiceResponse creates a response to parse from CreateService response
func CreateCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}