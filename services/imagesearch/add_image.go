package imagesearch

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

// AddImage invokes the imagesearch.AddImage API synchronously
// api document: https://help.aliyun.com/api/imagesearch/addimage.html
func (client *Client) AddImage(request *AddImageRequest) (response *AddImageResponse, err error) {
	response = CreateAddImageResponse()
	err = client.DoAction(request, response)
	return
}

// AddImageWithChan invokes the imagesearch.AddImage API asynchronously
// api document: https://help.aliyun.com/api/imagesearch/addimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddImageWithChan(request *AddImageRequest) (<-chan *AddImageResponse, <-chan error) {
	responseChan := make(chan *AddImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddImage(request)
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

// AddImageWithCallback invokes the imagesearch.AddImage API asynchronously
// api document: https://help.aliyun.com/api/imagesearch/addimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddImageWithCallback(request *AddImageRequest, callback func(response *AddImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddImageResponse
		var err error
		defer close(result)
		response, err = client.AddImage(request)
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

// AddImageRequest is the request struct for api AddImage
type AddImageRequest struct {
	*requests.RoaRequest
	PicContent    string           `position:"Body" name:"PicContent"`
	StrAttr       string           `position:"Body" name:"StrAttr"`
	InstanceName  string           `position:"Body" name:"InstanceName"`
	IntAttr       requests.Integer `position:"Body" name:"IntAttr"`
	ProductId     string           `position:"Body" name:"ProductId"`
	PicName       string           `position:"Body" name:"PicName"`
	CustomContent string           `position:"Body" name:"CustomContent"`
	Region        string           `position:"Body" name:"Region"`
	CategoryId    requests.Integer `position:"Body" name:"CategoryId"`
	Crop          requests.Boolean `position:"Body" name:"Crop"`
}

// AddImageResponse is the response struct for api AddImage
type AddImageResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Success   bool    `json:"Success" xml:"Success"`
	Message   string  `json:"Message" xml:"Message"`
	Code      int     `json:"Code" xml:"Code"`
	PicInfo   PicInfo `json:"PicInfo" xml:"PicInfo"`
}

// CreateAddImageRequest creates a request to invoke AddImage API
func CreateAddImageRequest() (request *AddImageRequest) {
	request = &AddImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ImageSearch", "2019-03-25", "AddImage", "/v2/image/add", "imagesearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddImageResponse creates a response to parse from AddImage response
func CreateAddImageResponse() (response *AddImageResponse) {
	response = &AddImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
