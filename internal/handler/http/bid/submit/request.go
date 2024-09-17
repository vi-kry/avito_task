package submit

import "github.com/google/uuid"

type SubmitBidRequest struct {
	BidId uuid.UUID `json:"bid_id"`
}
