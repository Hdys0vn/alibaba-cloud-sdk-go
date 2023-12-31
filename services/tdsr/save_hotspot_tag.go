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

// SaveHotspotTag invokes the tdsr.SaveHotspotTag API synchronously
// api document: https://help.aliyun.com/api/tdsr/savehotspottag.html
func (client *Client) SaveHotspotTag(request *SaveHotspotTagRequest) (response *SaveHotspotTagResponse, err error) {
	response = CreateSaveHotspotTagResponse()
	err = client.DoAction(request, response)
	return
}

// SaveHotspotTagWithChan invokes the tdsr.SaveHotspotTag API asynchronously
// api document: https://help.aliyun.com/api/tdsr/savehotspottag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveHotspotTagWithChan(request *SaveHotspotTagRequest) (<-chan *SaveHotspotTagResponse, <-chan error) {
	responseChan := make(chan *SaveHotspotTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveHotspotTag(request)
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

// SaveHotspotTagWithCallback invokes the tdsr.SaveHotspotTag API asynchronously
// api document: https://help.aliyun.com/api/tdsr/savehotspottag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveHotspotTagWithCallback(request *SaveHotspotTagRequest, callback func(response *SaveHotspotTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveHotspotTagResponse
		var err error
		defer close(result)
		response, err = client.SaveHotspotTag(request)
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

// SaveHotspotTagRequest is the request struct for api SaveHotspotTag
type SaveHotspotTagRequest struct {
	*requests.RpcRequest
	SubSceneUuid string `position:"Query" name:"SubSceneUuid"`
	ParamTag     string `position:"Query" name:"ParamTag"`
}

// SaveHotspotTagResponse is the response struct for api SaveHotspotTag
type SaveHotspotTagResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateSaveHotspotTagRequest creates a request to invoke SaveHotspotTag API
func CreateSaveHotspotTagRequest() (request *SaveHotspotTagRequest) {
	request = &SaveHotspotTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("tdsr", "2020-01-01", "SaveHotspotTag", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveHotspotTagResponse creates a response to parse from SaveHotspotTag response
func CreateSaveHotspotTagResponse() (response *SaveHotspotTagResponse) {
	response = &SaveHotspotTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
