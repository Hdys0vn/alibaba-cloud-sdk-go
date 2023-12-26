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

// QueryDeviceVodUrl invokes the linkvisual.QueryDeviceVodUrl API synchronously
func (client *Client) QueryDeviceVodUrl(request *QueryDeviceVodUrlRequest) (response *QueryDeviceVodUrlResponse, err error) {
	response = CreateQueryDeviceVodUrlResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceVodUrlWithChan invokes the linkvisual.QueryDeviceVodUrl API asynchronously
func (client *Client) QueryDeviceVodUrlWithChan(request *QueryDeviceVodUrlRequest) (<-chan *QueryDeviceVodUrlResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceVodUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceVodUrl(request)
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

// QueryDeviceVodUrlWithCallback invokes the linkvisual.QueryDeviceVodUrl API asynchronously
func (client *Client) QueryDeviceVodUrlWithCallback(request *QueryDeviceVodUrlRequest, callback func(response *QueryDeviceVodUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceVodUrlResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceVodUrl(request)
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

// QueryDeviceVodUrlRequest is the request struct for api QueryDeviceVodUrl
type QueryDeviceVodUrlRequest struct {
	*requests.RpcRequest
	Scheme           string           `position:"Query" name:"Scheme"`
	PlayUnLimited    requests.Boolean `position:"Query" name:"PlayUnLimited"`
	EncryptType      requests.Integer `position:"Query" name:"EncryptType"`
	IotId            string           `position:"Query" name:"IotId"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	ShouldEncrypt    requests.Boolean `position:"Query" name:"ShouldEncrypt"`
	EnableStun       requests.Boolean `position:"Query" name:"EnableStun"`
	ProductKey       string           `position:"Query" name:"ProductKey"`
	FileName         string           `position:"Query" name:"FileName"`
	SeekTime         requests.Integer `position:"Query" name:"SeekTime"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	DeviceName       string           `position:"Query" name:"DeviceName"`
	UrlValidDuration requests.Integer `position:"Query" name:"UrlValidDuration"`
}

// QueryDeviceVodUrlResponse is the response struct for api QueryDeviceVodUrl
type QueryDeviceVodUrlResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryDeviceVodUrlRequest creates a request to invoke QueryDeviceVodUrl API
func CreateQueryDeviceVodUrlRequest() (request *QueryDeviceVodUrlRequest) {
	request = &QueryDeviceVodUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryDeviceVodUrl", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceVodUrlResponse creates a response to parse from QueryDeviceVodUrl response
func CreateQueryDeviceVodUrlResponse() (response *QueryDeviceVodUrlResponse) {
	response = &QueryDeviceVodUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}