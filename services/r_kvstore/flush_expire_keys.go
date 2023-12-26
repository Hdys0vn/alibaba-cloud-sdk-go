package r_kvstore

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

// FlushExpireKeys invokes the r_kvstore.FlushExpireKeys API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/flushexpirekeys.html
func (client *Client) FlushExpireKeys(request *FlushExpireKeysRequest) (response *FlushExpireKeysResponse, err error) {
	response = CreateFlushExpireKeysResponse()
	err = client.DoAction(request, response)
	return
}

// FlushExpireKeysWithChan invokes the r_kvstore.FlushExpireKeys API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/flushexpirekeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FlushExpireKeysWithChan(request *FlushExpireKeysRequest) (<-chan *FlushExpireKeysResponse, <-chan error) {
	responseChan := make(chan *FlushExpireKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FlushExpireKeys(request)
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

// FlushExpireKeysWithCallback invokes the r_kvstore.FlushExpireKeys API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/flushexpirekeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FlushExpireKeysWithCallback(request *FlushExpireKeysRequest, callback func(response *FlushExpireKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FlushExpireKeysResponse
		var err error
		defer close(result)
		response, err = client.FlushExpireKeys(request)
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

// FlushExpireKeysRequest is the request struct for api FlushExpireKeys
type FlushExpireKeysRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// FlushExpireKeysResponse is the response struct for api FlushExpireKeys
type FlushExpireKeysResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateFlushExpireKeysRequest creates a request to invoke FlushExpireKeys API
func CreateFlushExpireKeysRequest() (request *FlushExpireKeysRequest) {
	request = &FlushExpireKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "FlushExpireKeys", "redisa", "openAPI")
	return
}

// CreateFlushExpireKeysResponse creates a response to parse from FlushExpireKeys response
func CreateFlushExpireKeysResponse() (response *FlushExpireKeysResponse) {
	response = &FlushExpireKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
