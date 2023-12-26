package imagerecog

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

// TaggingImage invokes the imagerecog.TaggingImage API synchronously
func (client *Client) TaggingImage(request *TaggingImageRequest) (response *TaggingImageResponse, err error) {
	response = CreateTaggingImageResponse()
	err = client.DoAction(request, response)
	return
}

// TaggingImageWithChan invokes the imagerecog.TaggingImage API asynchronously
func (client *Client) TaggingImageWithChan(request *TaggingImageRequest) (<-chan *TaggingImageResponse, <-chan error) {
	responseChan := make(chan *TaggingImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaggingImage(request)
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

// TaggingImageWithCallback invokes the imagerecog.TaggingImage API asynchronously
func (client *Client) TaggingImageWithCallback(request *TaggingImageRequest, callback func(response *TaggingImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaggingImageResponse
		var err error
		defer close(result)
		response, err = client.TaggingImage(request)
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

// TaggingImageRequest is the request struct for api TaggingImage
type TaggingImageRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	Mode               string           `position:"Body" name:"Mode"`
	OssFile            string           `position:"Query" name:"OssFile"`
	ImageType          requests.Integer `position:"Body" name:"ImageType"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	Async              requests.Boolean `position:"Body" name:"Async"`
	ImageURL           string           `position:"Body" name:"ImageURL"`
}

// TaggingImageResponse is the response struct for api TaggingImage
type TaggingImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTaggingImageRequest creates a request to invoke TaggingImage API
func CreateTaggingImageRequest() (request *TaggingImageRequest) {
	request = &TaggingImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imagerecog", "2019-09-30", "TaggingImage", "", "")
	request.Method = requests.POST
	return
}

// CreateTaggingImageResponse creates a response to parse from TaggingImage response
func CreateTaggingImageResponse() (response *TaggingImageResponse) {
	response = &TaggingImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}