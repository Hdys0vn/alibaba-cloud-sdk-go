package imageenhan

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

// RecolorImage invokes the imageenhan.RecolorImage API synchronously
func (client *Client) RecolorImage(request *RecolorImageRequest) (response *RecolorImageResponse, err error) {
	response = CreateRecolorImageResponse()
	err = client.DoAction(request, response)
	return
}

// RecolorImageWithChan invokes the imageenhan.RecolorImage API asynchronously
func (client *Client) RecolorImageWithChan(request *RecolorImageRequest) (<-chan *RecolorImageResponse, <-chan error) {
	responseChan := make(chan *RecolorImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecolorImage(request)
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

// RecolorImageWithCallback invokes the imageenhan.RecolorImage API asynchronously
func (client *Client) RecolorImageWithCallback(request *RecolorImageRequest, callback func(response *RecolorImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecolorImageResponse
		var err error
		defer close(result)
		response, err = client.RecolorImage(request)
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

// RecolorImageRequest is the request struct for api RecolorImage
type RecolorImageRequest struct {
	*requests.RpcRequest
	Mode          string                       `position:"Body" name:"Mode"`
	ColorCount    requests.Integer             `position:"Body" name:"ColorCount"`
	ColorTemplate *[]RecolorImageColorTemplate `position:"Body" name:"ColorTemplate"  type:"Repeated"`
	Url           string                       `position:"Body" name:"Url"`
	RefUrl        string                       `position:"Body" name:"RefUrl"`
}

// RecolorImageColorTemplate is a repeated param struct in RecolorImageRequest
type RecolorImageColorTemplate struct {
	Color string `name:"Color"`
}

// RecolorImageResponse is the response struct for api RecolorImage
type RecolorImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecolorImageRequest creates a request to invoke RecolorImage API
func CreateRecolorImageRequest() (request *RecolorImageRequest) {
	request = &RecolorImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "RecolorImage", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecolorImageResponse creates a response to parse from RecolorImage response
func CreateRecolorImageResponse() (response *RecolorImageResponse) {
	response = &RecolorImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}