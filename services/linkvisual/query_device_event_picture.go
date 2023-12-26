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

// QueryDeviceEventPicture invokes the linkvisual.QueryDeviceEventPicture API synchronously
func (client *Client) QueryDeviceEventPicture(request *QueryDeviceEventPictureRequest) (response *QueryDeviceEventPictureResponse, err error) {
	response = CreateQueryDeviceEventPictureResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceEventPictureWithChan invokes the linkvisual.QueryDeviceEventPicture API asynchronously
func (client *Client) QueryDeviceEventPictureWithChan(request *QueryDeviceEventPictureRequest) (<-chan *QueryDeviceEventPictureResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceEventPictureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceEventPicture(request)
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

// QueryDeviceEventPictureWithCallback invokes the linkvisual.QueryDeviceEventPicture API asynchronously
func (client *Client) QueryDeviceEventPictureWithCallback(request *QueryDeviceEventPictureRequest, callback func(response *QueryDeviceEventPictureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceEventPictureResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceEventPicture(request)
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

// QueryDeviceEventPictureRequest is the request struct for api QueryDeviceEventPicture
type QueryDeviceEventPictureRequest struct {
	*requests.RpcRequest
	EventId       string `position:"Query" name:"EventId"`
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// QueryDeviceEventPictureResponse is the response struct for api QueryDeviceEventPicture
type QueryDeviceEventPictureResponse struct {
	*responses.BaseResponse
	Code         int    `json:"Code" xml:"Code"`
	Data         string `json:"Data" xml:"Data"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateQueryDeviceEventPictureRequest creates a request to invoke QueryDeviceEventPicture API
func CreateQueryDeviceEventPictureRequest() (request *QueryDeviceEventPictureRequest) {
	request = &QueryDeviceEventPictureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryDeviceEventPicture", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceEventPictureResponse creates a response to parse from QueryDeviceEventPicture response
func CreateQueryDeviceEventPictureResponse() (response *QueryDeviceEventPictureResponse) {
	response = &QueryDeviceEventPictureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}