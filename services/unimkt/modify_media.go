package unimkt

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

// ModifyMedia invokes the unimkt.ModifyMedia API synchronously
func (client *Client) ModifyMedia(request *ModifyMediaRequest) (response *ModifyMediaResponse, err error) {
	response = CreateModifyMediaResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMediaWithChan invokes the unimkt.ModifyMedia API asynchronously
func (client *Client) ModifyMediaWithChan(request *ModifyMediaRequest) (<-chan *ModifyMediaResponse, <-chan error) {
	responseChan := make(chan *ModifyMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMedia(request)
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

// ModifyMediaWithCallback invokes the unimkt.ModifyMedia API asynchronously
func (client *Client) ModifyMediaWithCallback(request *ModifyMediaRequest, callback func(response *ModifyMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMediaResponse
		var err error
		defer close(result)
		response, err = client.ModifyMedia(request)
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

// ModifyMediaRequest is the request struct for api ModifyMedia
type ModifyMediaRequest struct {
	*requests.RpcRequest
	Business         string `position:"Query" name:"Business"`
	Media            string `position:"Body" name:"Media"`
	UserId           string `position:"Query" name:"UserId"`
	OriginSiteUserId string `position:"Query" name:"OriginSiteUserId"`
	Environment      string `position:"Query" name:"Environment"`
	AppName          string `position:"Query" name:"AppName"`
	TenantId         string `position:"Query" name:"TenantId"`
	UserSite         string `position:"Query" name:"UserSite"`
}

// ModifyMediaResponse is the response struct for api ModifyMedia
type ModifyMediaResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateModifyMediaRequest creates a request to invoke ModifyMedia API
func CreateModifyMediaRequest() (request *ModifyMediaRequest) {
	request = &ModifyMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "ModifyMedia", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyMediaResponse creates a response to parse from ModifyMedia response
func CreateModifyMediaResponse() (response *ModifyMediaResponse) {
	response = &ModifyMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
