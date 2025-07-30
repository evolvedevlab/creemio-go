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
