package tdsr

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

// PublishHotspot invokes the tdsr.PublishHotspot API synchronously
// api document: https://help.aliyun.com/api/tdsr/publishhotspot.html
func (client *Client) PublishHotspot(request *PublishHotspotRequest) (response *PublishHotspotResponse, err error) {
	response = CreatePublishHotspotResponse()
	err = client.DoAction(request, response)
	return
}

// PublishHotspotWithChan invokes the tdsr.PublishHotspot API asynchronously
// api document: https://help.aliyun.com/api/tdsr/publishhotspot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishHotspotWithChan(request *PublishHotspotRequest) (<-chan *PublishHotspotResponse, <-chan error) {
	responseChan := make(chan *PublishHotspotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishHotspot(request)
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

// PublishHotspotWithCallback invokes the tdsr.PublishHotspot API asynchronously
// api document: https://help.aliyun.com/api/tdsr/publishhotspot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishHotspotWithCallback(request *PublishHotspotRequest, callback func(response *PublishHotspotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishHotspotResponse
		var err error
		defer close(result)
		response, err = client.PublishHotspot(request)
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

// PublishHotspotRequest is the request struct for api PublishHotspot
type PublishHotspotRequest struct {
	*requests.RpcRequest
	SubSceneUuid string `position:"Query" name:"SubSceneUuid"`
	ParamTag     string `position:"Query" name:"ParamTag"`
}

// PublishHotspotResponse is the response struct for api PublishHotspot
type PublishHotspotResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	Success    bool                   `json:"Success" xml:"Success"`
	ErrMessage string                 `json:"ErrMessage" xml:"ErrMessage"`
	Data       map[string]interface{} `json:"Data" xml:"Data"`
}

// CreatePublishHotspotRequest creates a request to invoke PublishHotspot API
func CreatePublishHotspotRequest() (request *PublishHotspotRequest) {
	request = &PublishHotspotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("tdsr", "2020-01-01", "PublishHotspot", "", "")
	request.Method = requests.POST
	return
}

// CreatePublishHotspotResponse creates a response to parse from PublishHotspot response
func CreatePublishHotspotResponse() (response *PublishHotspotResponse) {
	response = &PublishHotspotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
