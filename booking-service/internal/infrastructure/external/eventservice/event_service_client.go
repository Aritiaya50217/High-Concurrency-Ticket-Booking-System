package eventservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SeatServiceClient struct {
	baseURL string
	client  *http.Client
}

func NewSeatServiceClient(baseURL string) *SeatServiceClient {
	return &SeatServiceClient{baseURL: baseURL, client: &http.Client{
		Timeout: 5 * time.Second,
	}}
}

func (c *SeatServiceClient) ReserveSeat(ctx context.Context, token string, eventID, seatID uint) error {
	reqBody := ReserveSeatRequest{EventID: eventID, SeatID: seatID}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/api/seat/reserve", c.baseURL), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("reserve seat failed: status=%d body=%s", resp.StatusCode, string(body))
	}
	return nil

}
