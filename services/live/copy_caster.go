package live

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

// CopyCaster invokes the live.CopyCaster API synchronously
func (client *Client) CopyCaster(request *CopyCasterRequest) (response *CopyCasterResponse, err error) {
	response = CreateCopyCasterResponse()
	err = client.DoAction(request, response)
	return
}

// CopyCasterWithChan invokes the live.CopyCaster API asynchronously
func (client *Client) CopyCasterWithChan(request *CopyCasterRequest) (<-chan *CopyCasterResponse, <-chan error) {
	responseChan := make(chan *CopyCasterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyCaster(request)
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

// CopyCasterWithCallback invokes the live.CopyCaster API asynchronously
func (client *Client) CopyCasterWithCallback(request *CopyCasterRequest, callback func(response *CopyCasterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyCasterResponse
		var err error
		defer close(result)
		response, err = client.CopyCaster(request)
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

// CopyCasterRequest is the request struct for api CopyCaster
type CopyCasterRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	CasterName  string           `position:"Query" name:"CasterName"`
	SrcCasterId string           `position:"Query" name:"SrcCasterId"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// CopyCasterResponse is the response struct for api CopyCaster
type CopyCasterResponse struct {
	*responses.BaseResponse
	CasterId  string `json:"CasterId" xml:"CasterId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCopyCasterRequest creates a request to invoke CopyCaster API
func CreateCopyCasterRequest() (request *CopyCasterRequest) {
	request = &CopyCasterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CopyCaster", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCopyCasterResponse creates a response to parse from CopyCaster response
func CreateCopyCasterResponse() (response *CopyCasterResponse) {
	response = &CopyCasterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
