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

// UpdateLeniPrivateIpAddress invokes the eflo.UpdateLeniPrivateIpAddress API synchronously
func (client *Client) UpdateLeniPrivateIpAddress(request *UpdateLeniPrivateIpAddressRequest) (response *UpdateLeniPrivateIpAddressResponse, err error) {
	response = CreateUpdateLeniPrivateIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLeniPrivateIpAddressWithChan invokes the eflo.UpdateLeniPrivateIpAddress API asynchronously
func (client *Client) UpdateLeniPrivateIpAddressWithChan(request *UpdateLeniPrivateIpAddressRequest) (<-chan *UpdateLeniPrivateIpAddressResponse, <-chan error) {
	responseChan := make(chan *UpdateLeniPrivateIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLeniPrivateIpAddress(request)
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

// UpdateLeniPrivateIpAddressWithCallback invokes the eflo.UpdateLeniPrivateIpAddress API asynchronously
func (client *Client) UpdateLeniPrivateIpAddressWithCallback(request *UpdateLeniPrivateIpAddressRequest, callback func(response *UpdateLeniPrivateIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLeniPrivateIpAddressResponse
		var err error
		defer close(result)
		response, err = client.UpdateLeniPrivateIpAddress(request)
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

// UpdateLeniPrivateIpAddressRequest is the request struct for api UpdateLeniPrivateIpAddress
type UpdateLeniPrivateIpAddressRequest struct {
	*requests.RpcRequest
	Description               string `position:"Body" name:"Description"`
	IpName                    string `position:"Body" name:"IpName"`
	ElasticNetworkInterfaceId string `position:"Body" name:"ElasticNetworkInterfaceId"`
}

// UpdateLeniPrivateIpAddressResponse is the response struct for api UpdateLeniPrivateIpAddress
type UpdateLeniPrivateIpAddressResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateUpdateLeniPrivateIpAddressRequest creates a request to invoke UpdateLeniPrivateIpAddress API
func CreateUpdateLeniPrivateIpAddressRequest() (request *UpdateLeniPrivateIpAddressRequest) {
	request = &UpdateLeniPrivateIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "UpdateLeniPrivateIpAddress", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLeniPrivateIpAddressResponse creates a response to parse from UpdateLeniPrivateIpAddress response
func CreateUpdateLeniPrivateIpAddressResponse() (response *UpdateLeniPrivateIpAddressResponse) {
	response = &UpdateLeniPrivateIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
