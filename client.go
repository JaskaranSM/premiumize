package premiumize

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"premiumize/types"
	"strconv"
	"strings"
)

const BASE_URL string = "https://www.premiumize.me/api"

type Client struct {
	api_key     string
	http_client http.Client
}

func (c *Client) GetBaseParams() *url.Values {
	params := url.Values{}
	params.Add("apikey", c.api_key)
	return &params
}

func (c *Client) ListTransfers() (*types.ListTransfersResponse, error) {
	req, _ := http.NewRequest("GET", BASE_URL+"/transfer/list", nil)
	params := c.GetBaseParams()
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	listResp := &types.ListTransfersResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = listResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

func (c *Client) AddUrl(_url string) (*types.CreateTransferResponse, error) {
	data := url.Values{
		"src": {_url},
	}
	req, _ := http.NewRequest("POST", BASE_URL+"/transfer/create", strings.NewReader(data.Encode()))
	params := c.GetBaseParams()
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	createTransferResp := &types.CreateTransferResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = createTransferResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return createTransferResp, nil
}

func (c *Client) DeleteTransfer(transfer_id string) (*types.DeleteResponse, error) {
	data := url.Values{
		"id": {transfer_id},
	}
	req, _ := http.NewRequest("POST", BASE_URL+"/transfer/delete", strings.NewReader(data.Encode()))
	params := c.GetBaseParams()
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	deleteResp := &types.DeleteResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = deleteResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return deleteResp, nil
}

// empty string for listing root
func (c *Client) ListFolderByFolderId(folder_id string) (*types.ListFolderResponse, error) {
	req, _ := http.NewRequest("GET", BASE_URL+"/folder/list", nil)
	params := c.GetBaseParams()
	params.Add("id", folder_id)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	listResp := &types.ListFolderResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = listResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

func (c *Client) DownloadFile(_url string, filepath string) (*types.Progress, error) {
	out, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("GET", _url, nil)
	params := c.GetBaseParams()
	req.URL.RawQuery = params.Encode()
	resp, err := c.http_client.Do(req)
	if err != nil {
		return nil, err
	}
	var t_size int64
	t_size_string := resp.Header.Get("Content-Length")
	if n, err := strconv.Atoi(t_size_string); err == nil {
		t_size = int64(n)
	} else {
		t_size = 0
	}
	progress := types.NewProgress(t_size)
	go func() {
		io.Copy(out, io.TeeReader(resp.Body, progress))
		progress.IsComplete = true
		resp.Body.Close()
		out.Close()
	}()
	return progress, nil
}

func (c *Client) GetItemDetails(item_id string) (*types.ItemResponse, error) {
	req, _ := http.NewRequest("GET", BASE_URL+"/item/details", nil)
	params := c.GetBaseParams()
	params.Add("id", item_id)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	itemResp := &types.ItemResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = itemResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return itemResp, nil
}

func (c *Client) Delete(id string, folder bool) (*types.DeleteResponse, error) {
	data := url.Values{
		"id": {id},
	}
	var req_suffix string
	if folder {
		req_suffix = "/folder/delete"
	} else {
		req_suffix = "/item/delete"
	}
	req, _ := http.NewRequest("POST", BASE_URL+req_suffix, strings.NewReader(data.Encode()))
	params := c.GetBaseParams()
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	deleteResp := &types.DeleteResponse{}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = deleteResp.Unmarshall(bytes)
	if err != nil {
		return nil, err
	}
	return deleteResp, nil
}

func (c *Client) DeleteFile(file_id string) (*types.DeleteResponse, error) {
	return c.Delete(file_id, false)
}

func (c *Client) DeleteFolder(folder_id string) (*types.DeleteResponse, error) {
	return c.Delete(folder_id, true)
}

func (c *Client) SetHttpClient(client http.Client) {
	c.http_client = client
}

func NewPremiumizeClient(api_key string) *Client {
	return &Client{api_key: api_key}
}
