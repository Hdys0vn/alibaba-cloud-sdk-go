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

// ChangeMediaStatus invokes the unimkt.ChangeMediaStatus API synchronously
func (client *Client) ChangeMediaStatus(request *ChangeMediaStatusRequest) (response *ChangeMediaStatusResponse, err error) {
	response = CreateChangeMediaStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeMediaStatusWithChan invokes the unimkt.ChangeMediaStatus API asynchronously
func (client *Client) ChangeMediaStatusWithChan(request *ChangeMediaStatusRequest) (<-chan *ChangeMediaStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeMediaStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeMediaStatus(request)
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

// ChangeMediaStatusWithCallback invokes the unimkt.ChangeMediaStatus API asynchronously
func (client *Client) ChangeMediaStatusWithCallback(request *ChangeMediaStatusRequest, callback func(response *ChangeMediaStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeMediaStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeMediaStatus(request)
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

// ChangeMediaStatusRequest is the request struct for api ChangeMediaStatus
type ChangeMediaStatusRequest struct {
	*requests.RpcRequest
	MessageKey       string `position:"Body" name:"MessageKey"`
	Business         string `position:"Query" name:"Business"`
	MediaId          string `position:"Body" name:"MediaId"`
	MediaStatus      string `position:"Body" name:"MediaStatus"`
	Message          string `position:"Body" name:"Message"`
	UserId           string `position:"Query" name:"UserId"`
	OriginSiteUserId string `position:"Query" name:"OriginSiteUserId"`
	Environment      string `position:"Query" name:"Environment"`
	AppName          string `position:"Query" name:"AppName"`
	TenantId         string `position:"Query" name:"TenantId"`
	UserSite         string `position:"Query" name:"UserSite"`
	AccessStatus     string `position:"Body" name:"AccessStatus"`
}

// ChangeMediaStatusResponse is the response struct for api ChangeMediaStatus
type ChangeMediaStatusResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateChangeMediaStatusRequest creates a request to invoke ChangeMediaStatus API
func CreateChangeMediaStatusRequest() (request *ChangeMediaStatusRequest) {
	request = &ChangeMediaStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "ChangeMediaStatus", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeMediaStatusResponse creates a response to parse from ChangeMediaStatus response
func CreateChangeMediaStatusResponse() (response *ChangeMediaStatusResponse) {
	response = &ChangeMediaStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
