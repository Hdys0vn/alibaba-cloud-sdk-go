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

// QueryRecordDownloadUrl invokes the linkvisual.QueryRecordDownloadUrl API synchronously
func (client *Client) QueryRecordDownloadUrl(request *QueryRecordDownloadUrlRequest) (response *QueryRecordDownloadUrlResponse, err error) {
	response = CreateQueryRecordDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecordDownloadUrlWithChan invokes the linkvisual.QueryRecordDownloadUrl API asynchronously
func (client *Client) QueryRecordDownloadUrlWithChan(request *QueryRecordDownloadUrlRequest) (<-chan *QueryRecordDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *QueryRecordDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecordDownloadUrl(request)
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

// QueryRecordDownloadUrlWithCallback invokes the linkvisual.QueryRecordDownloadUrl API asynchronously
func (client *Client) QueryRecordDownloadUrlWithCallback(request *QueryRecordDownloadUrlRequest, callback func(response *QueryRecordDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecordDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.QueryRecordDownloadUrl(request)
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

// QueryRecordDownloadUrlRequest is the request struct for api QueryRecordDownloadUrl
type QueryRecordDownloadUrlRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	FileName      string `position:"Query" name:"FileName"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// QueryRecordDownloadUrlResponse is the response struct for api QueryRecordDownloadUrl
type QueryRecordDownloadUrlResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryRecordDownloadUrlRequest creates a request to invoke QueryRecordDownloadUrl API
func CreateQueryRecordDownloadUrlRequest() (request *QueryRecordDownloadUrlRequest) {
	request = &QueryRecordDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryRecordDownloadUrl", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRecordDownloadUrlResponse creates a response to parse from QueryRecordDownloadUrl response
func CreateQueryRecordDownloadUrlResponse() (response *QueryRecordDownloadUrlResponse) {
	response = &QueryRecordDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}