package smartag

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

// CreateSagExpressConnectInterface invokes the smartag.CreateSagExpressConnectInterface API synchronously
func (client *Client) CreateSagExpressConnectInterface(request *CreateSagExpressConnectInterfaceRequest) (response *CreateSagExpressConnectInterfaceResponse, err error) {
	response = CreateCreateSagExpressConnectInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSagExpressConnectInterfaceWithChan invokes the smartag.CreateSagExpressConnectInterface API asynchronously
func (client *Client) CreateSagExpressConnectInterfaceWithChan(request *CreateSagExpressConnectInterfaceRequest) (<-chan *CreateSagExpressConnectInterfaceResponse, <-chan error) {
	responseChan := make(chan *CreateSagExpressConnectInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSagExpressConnectInterface(request)
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

// CreateSagExpressConnectInterfaceWithCallback invokes the smartag.CreateSagExpressConnectInterface API asynchronously
func (client *Client) CreateSagExpressConnectInterfaceWithCallback(request *CreateSagExpressConnectInterfaceRequest, callback func(response *CreateSagExpressConnectInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSagExpressConnectInterfaceResponse
		var err error
		defer close(result)
		response, err = client.CreateSagExpressConnectInterface(request)
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

// CreateSagExpressConnectInterfaceRequest is the request struct for api CreateSagExpressConnectInterface
type CreateSagExpressConnectInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vlan                 string           `position:"Query" name:"Vlan"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	IP                   string           `position:"Query" name:"IP"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
	PortName             string           `position:"Query" name:"PortName"`
}

// CreateSagExpressConnectInterfaceResponse is the response struct for api CreateSagExpressConnectInterface
type CreateSagExpressConnectInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateSagExpressConnectInterfaceRequest creates a request to invoke CreateSagExpressConnectInterface API
func CreateCreateSagExpressConnectInterfaceRequest() (request *CreateSagExpressConnectInterfaceRequest) {
	request = &CreateSagExpressConnectInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "CreateSagExpressConnectInterface", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSagExpressConnectInterfaceResponse creates a response to parse from CreateSagExpressConnectInterface response
func CreateCreateSagExpressConnectInterfaceResponse() (response *CreateSagExpressConnectInterfaceResponse) {
	response = &CreateSagExpressConnectInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
