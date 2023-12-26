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

// DescribeCasterStreamUrl invokes the live.DescribeCasterStreamUrl API synchronously
func (client *Client) DescribeCasterStreamUrl(request *DescribeCasterStreamUrlRequest) (response *DescribeCasterStreamUrlResponse, err error) {
	response = CreateDescribeCasterStreamUrlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCasterStreamUrlWithChan invokes the live.DescribeCasterStreamUrl API asynchronously
func (client *Client) DescribeCasterStreamUrlWithChan(request *DescribeCasterStreamUrlRequest) (<-chan *DescribeCasterStreamUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeCasterStreamUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasterStreamUrl(request)
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

// DescribeCasterStreamUrlWithCallback invokes the live.DescribeCasterStreamUrl API asynchronously
func (client *Client) DescribeCasterStreamUrlWithCallback(request *DescribeCasterStreamUrlRequest, callback func(response *DescribeCasterStreamUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCasterStreamUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasterStreamUrl(request)
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

// DescribeCasterStreamUrlRequest is the request struct for api DescribeCasterStreamUrl
type DescribeCasterStreamUrlRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCasterStreamUrlResponse is the response struct for api DescribeCasterStreamUrl
type DescribeCasterStreamUrlResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	CasterId      string        `json:"CasterId" xml:"CasterId"`
	Total         int           `json:"Total" xml:"Total"`
	CasterStreams CasterStreams `json:"CasterStreams" xml:"CasterStreams"`
}

// CreateDescribeCasterStreamUrlRequest creates a request to invoke DescribeCasterStreamUrl API
func CreateDescribeCasterStreamUrlRequest() (request *DescribeCasterStreamUrlRequest) {
	request = &DescribeCasterStreamUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasterStreamUrl", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCasterStreamUrlResponse creates a response to parse from DescribeCasterStreamUrl response
func CreateDescribeCasterStreamUrlResponse() (response *DescribeCasterStreamUrlResponse) {
	response = &DescribeCasterStreamUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
