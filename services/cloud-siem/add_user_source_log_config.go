package cloud_siem

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

// AddUserSourceLogConfig invokes the cloud_siem.AddUserSourceLogConfig API synchronously
func (client *Client) AddUserSourceLogConfig(request *AddUserSourceLogConfigRequest) (response *AddUserSourceLogConfigResponse, err error) {
	response = CreateAddUserSourceLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserSourceLogConfigWithChan invokes the cloud_siem.AddUserSourceLogConfig API asynchronously
func (client *Client) AddUserSourceLogConfigWithChan(request *AddUserSourceLogConfigRequest) (<-chan *AddUserSourceLogConfigResponse, <-chan error) {
	responseChan := make(chan *AddUserSourceLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUserSourceLogConfig(request)
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

// AddUserSourceLogConfigWithCallback invokes the cloud_siem.AddUserSourceLogConfig API asynchronously
func (client *Client) AddUserSourceLogConfigWithCallback(request *AddUserSourceLogConfigRequest, callback func(response *AddUserSourceLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserSourceLogConfigResponse
		var err error
		defer close(result)
		response, err = client.AddUserSourceLogConfig(request)
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

// AddUserSourceLogConfigRequest is the request struct for api AddUserSourceLogConfig
type AddUserSourceLogConfigRequest struct {
	*requests.RpcRequest
	DisPlayLine    string           `position:"Body" name:"DisPlayLine"`
	SubUserId      requests.Integer `position:"Body" name:"SubUserId"`
	SourceProdCode string           `position:"Body" name:"SourceProdCode"`
	SourceLogInfo  string           `position:"Body" name:"SourceLogInfo"`
	Deleted        requests.Integer `position:"Body" name:"Deleted"`
	SourceLogCode  string           `position:"Body" name:"SourceLogCode"`
}

// AddUserSourceLogConfigResponse is the response struct for api AddUserSourceLogConfig
type AddUserSourceLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddUserSourceLogConfigRequest creates a request to invoke AddUserSourceLogConfig API
func CreateAddUserSourceLogConfigRequest() (request *AddUserSourceLogConfigRequest) {
	request = &AddUserSourceLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "AddUserSourceLogConfig", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserSourceLogConfigResponse creates a response to parse from AddUserSourceLogConfig response
func CreateAddUserSourceLogConfigResponse() (response *AddUserSourceLogConfigResponse) {
	response = &AddUserSourceLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}