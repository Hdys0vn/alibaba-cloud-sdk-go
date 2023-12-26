package cloudapi

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

// ModifyApiConfiguration invokes the cloudapi.ModifyApiConfiguration API synchronously
func (client *Client) ModifyApiConfiguration(request *ModifyApiConfigurationRequest) (response *ModifyApiConfigurationResponse, err error) {
	response = CreateModifyApiConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiConfigurationWithChan invokes the cloudapi.ModifyApiConfiguration API asynchronously
func (client *Client) ModifyApiConfigurationWithChan(request *ModifyApiConfigurationRequest) (<-chan *ModifyApiConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyApiConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApiConfiguration(request)
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

// ModifyApiConfigurationWithCallback invokes the cloudapi.ModifyApiConfiguration API asynchronously
func (client *Client) ModifyApiConfigurationWithCallback(request *ModifyApiConfigurationRequest, callback func(response *ModifyApiConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyApiConfiguration(request)
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

// ModifyApiConfigurationRequest is the request struct for api ModifyApiConfiguration
type ModifyApiConfigurationRequest struct {
	*requests.RpcRequest
	ContentTypeCategory   string           `position:"Query" name:"ContentTypeCategory"`
	ModelName             string           `position:"Query" name:"ModelName"`
	ErrorCodeSamples      string           `position:"Query" name:"ErrorCodeSamples"`
	AppCodeAuthType       string           `position:"Query" name:"AppCodeAuthType"`
	AuthType              string           `position:"Query" name:"AuthType"`
	HttpConfig            string           `position:"Query" name:"HttpConfig"`
	ServiceParameters     string           `position:"Query" name:"ServiceParameters"`
	FailResultSample      string           `position:"Query" name:"FailResultSample"`
	ContentTypeValue      string           `position:"Query" name:"ContentTypeValue"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	Visibility            string           `position:"Query" name:"Visibility"`
	RequestProtocol       string           `position:"Query" name:"RequestProtocol"`
	RequestMode           string           `position:"Query" name:"RequestMode"`
	BackendName           string           `position:"Query" name:"BackendName"`
	RequestPath           string           `position:"Query" name:"RequestPath"`
	ResultType            string           `position:"Query" name:"ResultType"`
	MockConfig            string           `position:"Query" name:"MockConfig"`
	ResultSample          string           `position:"Query" name:"ResultSample"`
	BodyModel             string           `position:"Query" name:"BodyModel"`
	VpcConfig             string           `position:"Query" name:"VpcConfig"`
	FunctionComputeConfig string           `position:"Query" name:"FunctionComputeConfig"`
	ApiId                 string           `position:"Query" name:"ApiId"`
	UseBackendService     requests.Boolean `position:"Query" name:"UseBackendService"`
	ServiceProtocol       string           `position:"Query" name:"ServiceProtocol"`
	Description           string           `position:"Query" name:"Description"`
	DisableInternet       requests.Boolean `position:"Query" name:"DisableInternet"`
	PostBodyDescription   string           `position:"Query" name:"PostBodyDescription"`
	AllowSignatureMethod  string           `position:"Query" name:"AllowSignatureMethod"`
	RequestHttpMethod     string           `position:"Query" name:"RequestHttpMethod"`
	ServiceParametersMap  string           `position:"Query" name:"ServiceParametersMap"`
	RequestParameters     string           `position:"Query" name:"RequestParameters"`
	BodyFormat            string           `position:"Query" name:"BodyFormat"`
	OssConfig             string           `position:"Query" name:"OssConfig"`
	ApiName               string           `position:"Query" name:"ApiName"`
	ServiceTimeout        requests.Integer `position:"Query" name:"ServiceTimeout"`
	ForceNonceCheck       requests.Boolean `position:"Query" name:"ForceNonceCheck"`
}

// ModifyApiConfigurationResponse is the response struct for api ModifyApiConfiguration
type ModifyApiConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyApiConfigurationRequest creates a request to invoke ModifyApiConfiguration API
func CreateModifyApiConfigurationRequest() (request *ModifyApiConfigurationRequest) {
	request = &ModifyApiConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApiConfiguration", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyApiConfigurationResponse creates a response to parse from ModifyApiConfiguration response
func CreateModifyApiConfigurationResponse() (response *ModifyApiConfigurationResponse) {
	response = &ModifyApiConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
