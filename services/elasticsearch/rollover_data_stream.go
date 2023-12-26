package elasticsearch

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

// RolloverDataStream invokes the elasticsearch.RolloverDataStream API synchronously
func (client *Client) RolloverDataStream(request *RolloverDataStreamRequest) (response *RolloverDataStreamResponse, err error) {
	response = CreateRolloverDataStreamResponse()
	err = client.DoAction(request, response)
	return
}

// RolloverDataStreamWithChan invokes the elasticsearch.RolloverDataStream API asynchronously
func (client *Client) RolloverDataStreamWithChan(request *RolloverDataStreamRequest) (<-chan *RolloverDataStreamResponse, <-chan error) {
	responseChan := make(chan *RolloverDataStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RolloverDataStream(request)
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

// RolloverDataStreamWithCallback invokes the elasticsearch.RolloverDataStream API asynchronously
func (client *Client) RolloverDataStreamWithCallback(request *RolloverDataStreamRequest, callback func(response *RolloverDataStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RolloverDataStreamResponse
		var err error
		defer close(result)
		response, err = client.RolloverDataStream(request)
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

// RolloverDataStreamRequest is the request struct for api RolloverDataStream
type RolloverDataStreamRequest struct {
	*requests.RoaRequest
	DataStream  string `position:"Path" name:"DataStream"`
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

// RolloverDataStreamResponse is the response struct for api RolloverDataStream
type RolloverDataStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateRolloverDataStreamRequest creates a request to invoke RolloverDataStream API
func CreateRolloverDataStreamRequest() (request *RolloverDataStreamRequest) {
	request = &RolloverDataStreamRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "RolloverDataStream", "/openapi/instances/[InstanceId]/data-streams/[DataStream]/rollover", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRolloverDataStreamResponse creates a response to parse from RolloverDataStream response
func CreateRolloverDataStreamResponse() (response *RolloverDataStreamResponse) {
	response = &RolloverDataStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}