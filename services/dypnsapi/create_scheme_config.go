package dypnsapi

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

// CreateSchemeConfig invokes the dypnsapi.CreateSchemeConfig API synchronously
func (client *Client) CreateSchemeConfig(request *CreateSchemeConfigRequest) (response *CreateSchemeConfigResponse, err error) {
	response = CreateCreateSchemeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSchemeConfigWithChan invokes the dypnsapi.CreateSchemeConfig API asynchronously
func (client *Client) CreateSchemeConfigWithChan(request *CreateSchemeConfigRequest) (<-chan *CreateSchemeConfigResponse, <-chan error) {
	responseChan := make(chan *CreateSchemeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSchemeConfig(request)
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

// CreateSchemeConfigWithCallback invokes the dypnsapi.CreateSchemeConfig API asynchronously
func (client *Client) CreateSchemeConfigWithCallback(request *CreateSchemeConfigRequest, callback func(response *CreateSchemeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSchemeConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateSchemeConfig(request)
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

// CreateSchemeConfigRequest is the request struct for api CreateSchemeConfig
type CreateSchemeConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AndroidPackageSign   string           `position:"Query" name:"AndroidPackageSign"`
	Platform             string           `position:"Query" name:"Platform"`
	H5Url                string           `position:"Query" name:"H5Url"`
	IosBundleId          string           `position:"Query" name:"IosBundleId"`
	AppName              string           `position:"Query" name:"AppName"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	H5Origin             string           `position:"Query" name:"H5Origin"`
	SchemeName           string           `position:"Query" name:"SchemeName"`
	AndroidPackageName   string           `position:"Query" name:"AndroidPackageName"`
}

// CreateSchemeConfigResponse is the response struct for api CreateSchemeConfig
type CreateSchemeConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateCreateSchemeConfigRequest creates a request to invoke CreateSchemeConfig API
func CreateCreateSchemeConfigRequest() (request *CreateSchemeConfigRequest) {
	request = &CreateSchemeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "CreateSchemeConfig", "dypnsapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSchemeConfigResponse creates a response to parse from CreateSchemeConfig response
func CreateCreateSchemeConfigResponse() (response *CreateSchemeConfigResponse) {
	response = &CreateSchemeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}