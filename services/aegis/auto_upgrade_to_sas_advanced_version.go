package aegis

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

// AutoUpgradeToSasAdvancedVersion invokes the aegis.AutoUpgradeToSasAdvancedVersion API synchronously
// api document: https://help.aliyun.com/api/aegis/autoupgradetosasadvancedversion.html
func (client *Client) AutoUpgradeToSasAdvancedVersion(request *AutoUpgradeToSasAdvancedVersionRequest) (response *AutoUpgradeToSasAdvancedVersionResponse, err error) {
	response = CreateAutoUpgradeToSasAdvancedVersionResponse()
	err = client.DoAction(request, response)
	return
}

// AutoUpgradeToSasAdvancedVersionWithChan invokes the aegis.AutoUpgradeToSasAdvancedVersion API asynchronously
// api document: https://help.aliyun.com/api/aegis/autoupgradetosasadvancedversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AutoUpgradeToSasAdvancedVersionWithChan(request *AutoUpgradeToSasAdvancedVersionRequest) (<-chan *AutoUpgradeToSasAdvancedVersionResponse, <-chan error) {
	responseChan := make(chan *AutoUpgradeToSasAdvancedVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AutoUpgradeToSasAdvancedVersion(request)
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

// AutoUpgradeToSasAdvancedVersionWithCallback invokes the aegis.AutoUpgradeToSasAdvancedVersion API asynchronously
// api document: https://help.aliyun.com/api/aegis/autoupgradetosasadvancedversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AutoUpgradeToSasAdvancedVersionWithCallback(request *AutoUpgradeToSasAdvancedVersionRequest, callback func(response *AutoUpgradeToSasAdvancedVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AutoUpgradeToSasAdvancedVersionResponse
		var err error
		defer close(result)
		response, err = client.AutoUpgradeToSasAdvancedVersion(request)
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

// AutoUpgradeToSasAdvancedVersionRequest is the request struct for api AutoUpgradeToSasAdvancedVersion
type AutoUpgradeToSasAdvancedVersionRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// AutoUpgradeToSasAdvancedVersionResponse is the response struct for api AutoUpgradeToSasAdvancedVersion
type AutoUpgradeToSasAdvancedVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAutoUpgradeToSasAdvancedVersionRequest creates a request to invoke AutoUpgradeToSasAdvancedVersion API
func CreateAutoUpgradeToSasAdvancedVersionRequest() (request *AutoUpgradeToSasAdvancedVersionRequest) {
	request = &AutoUpgradeToSasAdvancedVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "AutoUpgradeToSasAdvancedVersion", "vipaegis", "openAPI")
	return
}

// CreateAutoUpgradeToSasAdvancedVersionResponse creates a response to parse from AutoUpgradeToSasAdvancedVersion response
func CreateAutoUpgradeToSasAdvancedVersionResponse() (response *AutoUpgradeToSasAdvancedVersionResponse) {
	response = &AutoUpgradeToSasAdvancedVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
