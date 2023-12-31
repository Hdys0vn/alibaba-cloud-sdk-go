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

// ModifySagPortRouteProtocol invokes the smartag.ModifySagPortRouteProtocol API synchronously
func (client *Client) ModifySagPortRouteProtocol(request *ModifySagPortRouteProtocolRequest) (response *ModifySagPortRouteProtocolResponse, err error) {
	response = CreateModifySagPortRouteProtocolResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagPortRouteProtocolWithChan invokes the smartag.ModifySagPortRouteProtocol API asynchronously
func (client *Client) ModifySagPortRouteProtocolWithChan(request *ModifySagPortRouteProtocolRequest) (<-chan *ModifySagPortRouteProtocolResponse, <-chan error) {
	responseChan := make(chan *ModifySagPortRouteProtocolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagPortRouteProtocol(request)
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

// ModifySagPortRouteProtocolWithCallback invokes the smartag.ModifySagPortRouteProtocol API asynchronously
func (client *Client) ModifySagPortRouteProtocolWithCallback(request *ModifySagPortRouteProtocolRequest, callback func(response *ModifySagPortRouteProtocolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagPortRouteProtocolResponse
		var err error
		defer close(result)
		response, err = client.ModifySagPortRouteProtocol(request)
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

// ModifySagPortRouteProtocolRequest is the request struct for api ModifySagPortRouteProtocol
type ModifySagPortRouteProtocolRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RemoteIp             string           `position:"Query" name:"RemoteIp"`
	Vlan                 string           `position:"Query" name:"Vlan"`
	RemoteAs             string           `position:"Query" name:"RemoteAs"`
	RouteProtocol        string           `position:"Query" name:"RouteProtocol"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
	PortName             string           `position:"Query" name:"PortName"`
}

// ModifySagPortRouteProtocolResponse is the response struct for api ModifySagPortRouteProtocol
type ModifySagPortRouteProtocolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagPortRouteProtocolRequest creates a request to invoke ModifySagPortRouteProtocol API
func CreateModifySagPortRouteProtocolRequest() (request *ModifySagPortRouteProtocolRequest) {
	request = &ModifySagPortRouteProtocolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagPortRouteProtocol", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagPortRouteProtocolResponse creates a response to parse from ModifySagPortRouteProtocol response
func CreateModifySagPortRouteProtocolResponse() (response *ModifySagPortRouteProtocolResponse) {
	response = &ModifySagPortRouteProtocolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
