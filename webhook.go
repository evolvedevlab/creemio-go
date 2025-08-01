package creemio

type WebHookEvent string

const (
	WebHookEventCheckoutCompleted    WebHookEvent = "checkout.completed"
	WebHookEventSubscriptionActive   WebHookEvent = "subscription.active"
	WebHookEventSubscriptionPaid     WebHookEvent = "subscription.paid"
	WebHookEventSubscriptionCanceled WebHookEvent = "subscription.canceled"
	WebHookEventSubscriptionExpired  WebHookEvent = "subscription.expired"
	WebHookEventRefundCreated        WebHookEvent = "refund.created"
	WebHookEventDisputeCreated       WebHookEvent = "dispute.created"
	WebHookEventSubscriptionUpdated  WebHookEvent = "subscription.update"
	WebHookEventSubscriptionTrialing WebHookEvent = "subscription.trialing"
)

type WebHookRequest struct {
	ID        string       `json:"id"`
	EventType WebHookEvent `json:"eventType"`
	CreatedAt int64        `json:"created_at"`
}

type WebHookCheckoutRequest struct {
	ID             string       `json:"id"`
	EventType      WebHookEvent `json:"eventType"`
	CreatedAt      int64        `json:"created_at"`
	CheckoutObject Checkout     `json:"object"`
}

type WebHookSubscriptionRequest struct {
	ID                 string       `json:"id"`
	EventType          WebHookEvent `json:"eventType"`
	CreatedAt          int64        `json:"created_at"`
	SubscriptionObject Subscription `json:"object"`
}

type WebHookRefundRequest struct {
	ID           string       `json:"id"`
	EventType    WebHookEvent `json:"eventType"`
	CreatedAt    int64        `json:"created_at"`
	RefundObject Refund       `json:"object"`
}

type WebHookDisputeRequest struct {
	ID            string       `json:"id"`
	EventType     WebHookEvent `json:"eventType"`
	CreatedAt     int64        `json:"created_at"`
	DisputeObject Dispute      `json:"object"`
}

type Refund struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	Status         string         `json:"status"`
	RefundAmount   int            `json:"refund_amount"`
	RefundCurrency string         `json:"refund_currency"`
	Reason         string         `json:"reason"`
	Transaction    *Transaction   `json:"transaction"`
	Subscription   *Subscription  `json:"subscription"`
	Checkout       *Checkout      `json:"checkout"`
	Order          *CheckoutOrder `json:"order"`
	Customer       *Customer      `json:"customer"`
	CreatedAt      int64          `json:"created_at"`
	Mode           Mode           `json:"mode"`
}

type Dispute struct {
	ID           string         `json:"id"`
	Object       string         `json:"object"`
	Amount       int            `json:"amount"`
	Currency     string         `json:"currency"`
	Transaction  *Transaction   `json:"transaction"`
	Subscription *Subscription  `json:"subscription"`
	Checkout     *Checkout      `json:"checkout"`
	Order        *CheckoutOrder `json:"order"`
	Customer     *Customer      `json:"customer"`
	CreatedAt    int64          `json:"created_at"`
	Mode         Mode           `json:"mode"`
}
