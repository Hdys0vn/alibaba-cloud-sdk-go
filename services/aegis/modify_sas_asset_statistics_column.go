package aegis

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

// ModifySasAssetStatisticsColumn invokes the aegis.ModifySasAssetStatisticsColumn API synchronously
// api document: https://help.aliyun.com/api/aegis/modifysasassetstatisticscolumn.html
func (client *Client) ModifySasAssetStatisticsColumn(request *ModifySasAssetStatisticsColumnRequest) (response *ModifySasAssetStatisticsColumnResponse, err error) {
	response = CreateModifySasAssetStatisticsColumnResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySasAssetStatisticsColumnWithChan invokes the aegis.ModifySasAssetStatisticsColumn API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysasassetstatisticscolumn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySasAssetStatisticsColumnWithChan(request *ModifySasAssetStatisticsColumnRequest) (<-chan *ModifySasAssetStatisticsColumnResponse, <-chan error) {
	responseChan := make(chan *ModifySasAssetStatisticsColumnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySasAssetStatisticsColumn(request)
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

// ModifySasAssetStatisticsColumnWithCallback invokes the aegis.ModifySasAssetStatisticsColumn API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysasassetstatisticscolumn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySasAssetStatisticsColumnWithCallback(request *ModifySasAssetStatisticsColumnRequest, callback func(response *ModifySasAssetStatisticsColumnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySasAssetStatisticsColumnResponse
		var err error
		defer close(result)
		response, err = client.ModifySasAssetStatisticsColumn(request)
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

// ModifySasAssetStatisticsColumnRequest is the request struct for api ModifySasAssetStatisticsColumn
type ModifySasAssetStatisticsColumnRequest struct {
	*requests.RpcRequest
	SourceIp         string `position:"Query" name:"SourceIp"`
	StatisticsColumn string `position:"Query" name:"StatisticsColumn"`
}

// ModifySasAssetStatisticsColumnResponse is the response struct for api ModifySasAssetStatisticsColumn
type ModifySasAssetStatisticsColumnResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySasAssetStatisticsColumnRequest creates a request to invoke ModifySasAssetStatisticsColumn API
func CreateModifySasAssetStatisticsColumnRequest() (request *ModifySasAssetStatisticsColumnRequest) {
	request = &ModifySasAssetStatisticsColumnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifySasAssetStatisticsColumn", "vipaegis", "openAPI")
	return
}

// CreateModifySasAssetStatisticsColumnResponse creates a response to parse from ModifySasAssetStatisticsColumn response
func CreateModifySasAssetStatisticsColumnResponse() (response *ModifySasAssetStatisticsColumnResponse) {
	response = &ModifySasAssetStatisticsColumnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
