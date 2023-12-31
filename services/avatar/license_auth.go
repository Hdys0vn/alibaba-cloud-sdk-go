package avatar

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

// LicenseAuth invokes the avatar.LicenseAuth API synchronously
func (client *Client) LicenseAuth(request *LicenseAuthRequest) (response *LicenseAuthResponse, err error) {
	response = CreateLicenseAuthResponse()
	err = client.DoAction(request, response)
	return
}

// LicenseAuthWithChan invokes the avatar.LicenseAuth API asynchronously
func (client *Client) LicenseAuthWithChan(request *LicenseAuthRequest) (<-chan *LicenseAuthResponse, <-chan error) {
	responseChan := make(chan *LicenseAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LicenseAuth(request)
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

// LicenseAuthWithCallback invokes the avatar.LicenseAuth API asynchronously
func (client *Client) LicenseAuthWithCallback(request *LicenseAuthRequest, callback func(response *LicenseAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LicenseAuthResponse
		var err error
		defer close(result)
		response, err = client.LicenseAuth(request)
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

// LicenseAuthRequest is the request struct for api LicenseAuth
type LicenseAuthRequest struct {
	*requests.RpcRequest
	License  string           `position:"Query" name:"License"`
	AppId    string           `position:"Query" name:"AppId"`
	TenantId requests.Integer `position:"Query" name:"TenantId"`
}

// LicenseAuthResponse is the response struct for api LicenseAuth
type LicenseAuthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateLicenseAuthRequest creates a request to invoke LicenseAuth API
func CreateLicenseAuthRequest() (request *LicenseAuthRequest) {
	request = &LicenseAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "LicenseAuth", "", "")
	request.Method = requests.POST
	return
}

// CreateLicenseAuthResponse creates a response to parse from LicenseAuth response
func CreateLicenseAuthResponse() (response *LicenseAuthResponse) {
	response = &LicenseAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
