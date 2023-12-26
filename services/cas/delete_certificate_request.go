package cas

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

// DeleteCertificateRequest invokes the cas.DeleteCertificateRequest API synchronously
func (client *Client) DeleteCertificateRequest(request *DeleteCertificateRequestRequest) (response *DeleteCertificateRequestResponse, err error) {
	response = CreateDeleteCertificateRequestResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCertificateRequestWithChan invokes the cas.DeleteCertificateRequest API asynchronously
func (client *Client) DeleteCertificateRequestWithChan(request *DeleteCertificateRequestRequest) (<-chan *DeleteCertificateRequestResponse, <-chan error) {
	responseChan := make(chan *DeleteCertificateRequestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCertificateRequest(request)
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

// DeleteCertificateRequestWithCallback invokes the cas.DeleteCertificateRequest API asynchronously
func (client *Client) DeleteCertificateRequestWithCallback(request *DeleteCertificateRequestRequest, callback func(response *DeleteCertificateRequestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCertificateRequestResponse
		var err error
		defer close(result)
		response, err = client.DeleteCertificateRequest(request)
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

// DeleteCertificateRequestRequest is the request struct for api DeleteCertificateRequest
type DeleteCertificateRequestRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	OrderId  requests.Integer `position:"Query" name:"OrderId"`
}

// DeleteCertificateRequestResponse is the response struct for api DeleteCertificateRequest
type DeleteCertificateRequestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCertificateRequestRequest creates a request to invoke DeleteCertificateRequest API
func CreateDeleteCertificateRequestRequest() (request *DeleteCertificateRequestRequest) {
	request = &DeleteCertificateRequestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "DeleteCertificateRequest", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCertificateRequestResponse creates a response to parse from DeleteCertificateRequest response
func CreateDeleteCertificateRequestResponse() (response *DeleteCertificateRequestResponse) {
	response = &DeleteCertificateRequestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
