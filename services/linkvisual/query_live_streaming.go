package linkvisual

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

// QueryLiveStreaming invokes the linkvisual.QueryLiveStreaming API synchronously
func (client *Client) QueryLiveStreaming(request *QueryLiveStreamingRequest) (response *QueryLiveStreamingResponse, err error) {
	response = CreateQueryLiveStreamingResponse()
	err = client.DoAction(request, response)
	return
}

// QueryLiveStreamingWithChan invokes the linkvisual.QueryLiveStreaming API asynchronously
func (client *Client) QueryLiveStreamingWithChan(request *QueryLiveStreamingRequest) (<-chan *QueryLiveStreamingResponse, <-chan error) {
	responseChan := make(chan *QueryLiveStreamingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLiveStreaming(request)
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

// QueryLiveStreamingWithCallback invokes the linkvisual.QueryLiveStreaming API asynchronously
func (client *Client) QueryLiveStreamingWithCallback(request *QueryLiveStreamingRequest, callback func(response *QueryLiveStreamingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLiveStreamingResponse
		var err error
		defer close(result)
		response, err = client.QueryLiveStreaming(request)
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

// QueryLiveStreamingRequest is the request struct for api QueryLiveStreaming
type QueryLiveStreamingRequest struct {
	*requests.RpcRequest
	Scheme           string           `position:"Query" name:"Scheme"`
	PlayUnLimited    requests.Boolean `position:"Query" name:"PlayUnLimited"`
	EncryptType      requests.Integer `position:"Query" name:"EncryptType"`
	CacheDuration    requests.Integer `position:"Query" name:"CacheDuration"`
	IotId            string           `position:"Query" name:"IotId"`
	ForceIFrame      requests.Boolean `position:"Query" name:"ForceIFrame"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	ShouldEncrypt    requests.Boolean `position:"Query" name:"ShouldEncrypt"`
	EnableStun       requests.Boolean `position:"Query" name:"EnableStun"`
	StreamType       requests.Integer `position:"Query" name:"StreamType"`
	ProductKey       string           `position:"Query" name:"ProductKey"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	DeviceName       string           `position:"Query" name:"DeviceName"`
	UrlValidDuration requests.Integer `position:"Query" name:"UrlValidDuration"`
}

// QueryLiveStreamingResponse is the response struct for api QueryLiveStreaming
type QueryLiveStreamingResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryLiveStreamingRequest creates a request to invoke QueryLiveStreaming API
func CreateQueryLiveStreamingRequest() (request *QueryLiveStreamingRequest) {
	request = &QueryLiveStreamingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryLiveStreaming", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryLiveStreamingResponse creates a response to parse from QueryLiveStreaming response
func CreateQueryLiveStreamingResponse() (response *QueryLiveStreamingResponse) {
	response = &QueryLiveStreamingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}