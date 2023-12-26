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

// DescribeSagRouteProtocolBgp invokes the smartag.DescribeSagRouteProtocolBgp API synchronously
func (client *Client) DescribeSagRouteProtocolBgp(request *DescribeSagRouteProtocolBgpRequest) (response *DescribeSagRouteProtocolBgpResponse, err error) {
	response = CreateDescribeSagRouteProtocolBgpResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagRouteProtocolBgpWithChan invokes the smartag.DescribeSagRouteProtocolBgp API asynchronously
func (client *Client) DescribeSagRouteProtocolBgpWithChan(request *DescribeSagRouteProtocolBgpRequest) (<-chan *DescribeSagRouteProtocolBgpResponse, <-chan error) {
	responseChan := make(chan *DescribeSagRouteProtocolBgpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagRouteProtocolBgp(request)
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

// DescribeSagRouteProtocolBgpWithCallback invokes the smartag.DescribeSagRouteProtocolBgp API asynchronously
func (client *Client) DescribeSagRouteProtocolBgpWithCallback(request *DescribeSagRouteProtocolBgpRequest, callback func(response *DescribeSagRouteProtocolBgpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagRouteProtocolBgpResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagRouteProtocolBgp(request)
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

// DescribeSagRouteProtocolBgpRequest is the request struct for api DescribeSagRouteProtocolBgp
type DescribeSagRouteProtocolBgpRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagRouteProtocolBgpResponse is the response struct for api DescribeSagRouteProtocolBgp
type DescribeSagRouteProtocolBgpResponse struct {
	*responses.BaseResponse
	HoldTime     int         `json:"HoldTime" xml:"HoldTime"`
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	KeepAlive    int         `json:"KeepAlive" xml:"KeepAlive"`
	LocalAs      int         `json:"LocalAs" xml:"LocalAs"`
	RouterId     string      `json:"RouterId" xml:"RouterId"`
	AdvertiseIps string      `json:"AdvertiseIps" xml:"AdvertiseIps"`
	TaskStates   []TaskState `json:"TaskStates" xml:"TaskStates"`
}

// CreateDescribeSagRouteProtocolBgpRequest creates a request to invoke DescribeSagRouteProtocolBgp API
func CreateDescribeSagRouteProtocolBgpRequest() (request *DescribeSagRouteProtocolBgpRequest) {
	request = &DescribeSagRouteProtocolBgpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagRouteProtocolBgp", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagRouteProtocolBgpResponse creates a response to parse from DescribeSagRouteProtocolBgp response
func CreateDescribeSagRouteProtocolBgpResponse() (response *DescribeSagRouteProtocolBgpResponse) {
	response = &DescribeSagRouteProtocolBgpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
