package ports

import "hexagonal2/core/entity"

type SubscriptionRepository interface {
	CreateSubscription(subscription entity.Subscription) error
	GetSubscriptionByUserID(userID string) (*entity.Subscription, error)
}