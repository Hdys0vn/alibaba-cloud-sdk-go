package ecd

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

// DescribeImages invokes the ecd.DescribeImages API synchronously
func (client *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = CreateDescribeImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImagesWithChan invokes the ecd.DescribeImages API asynchronously
func (client *Client) DescribeImagesWithChan(request *DescribeImagesRequest) (<-chan *DescribeImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImages(request)
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

// DescribeImagesWithCallback invokes the ecd.DescribeImages API asynchronously
func (client *Client) DescribeImagesWithCallback(request *DescribeImagesRequest, callback func(response *DescribeImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeImages(request)
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

// DescribeImagesRequest is the request struct for api DescribeImages
type DescribeImagesRequest struct {
	*requests.RpcRequest
	ImageId             *[]string        `position:"Query" name:"ImageId"  type:"Repeated"`
	GpuCategory         requests.Boolean `position:"Query" name:"GpuCategory"`
	DesktopInstanceType string           `position:"Query" name:"DesktopInstanceType"`
	LanguageType        string           `position:"Query" name:"LanguageType"`
	NextToken           string           `position:"Query" name:"NextToken"`
	FotaChannel         string           `position:"Query" name:"FotaChannel"`
	ImageType           string           `position:"Query" name:"ImageType"`
	OsType              string           `position:"Query" name:"OsType"`
	ImageStatus         string           `position:"Query" name:"ImageStatus"`
	MaxResults          requests.Integer `position:"Query" name:"MaxResults"`
	ProtocolType        string           `position:"Query" name:"ProtocolType"`
	GpuDriverVersion    string           `position:"Query" name:"GpuDriverVersion"`
}

// DescribeImagesResponse is the response struct for api DescribeImages
type DescribeImagesResponse struct {
	*responses.BaseResponse
	NextToken string  `json:"NextToken" xml:"NextToken"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Images    []Image `json:"Images" xml:"Images"`
}

// CreateDescribeImagesRequest creates a request to invoke DescribeImages API
func CreateDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeImages", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeImagesResponse creates a response to parse from DescribeImages response
func CreateDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}