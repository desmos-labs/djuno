package standard

import (
	"fmt"

	"firebase.google.com/go/v4/messaging"
	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

	"github.com/desmos-labs/djuno/v2/types"
	notificationsbuilder "github.com/desmos-labs/djuno/v2/x/notifications/builder"
)

var (
	_ notificationsbuilder.PostsNotificationsBuilder = &DefaultPostsNotificationsBuilder{}
)

type DefaultPostsNotificationsBuilder struct {
	m notificationsbuilder.UtilityModule
}

func NewDefaultPostsNotificationsBuilder(utilityModule notificationsbuilder.UtilityModule) *DefaultPostsNotificationsBuilder {
	return &DefaultPostsNotificationsBuilder{
		m: utilityModule,
	}
}

func (d DefaultPostsNotificationsBuilder) ConversationReply() notificationsbuilder.PostNotificationBuilder {
	return func(originalPost types.Post, reply types.Post) *notificationsbuilder.NotificationData {
		return &notificationsbuilder.NotificationData{
			Notification: &messaging.Notification{
				Title: "Someone replied to your post! 💬",
				Body:  fmt.Sprintf("%s replied to your post", d.m.GetDisplayName(reply.Author)),
			},
			Data: map[string]string{
				notificationsbuilder.NotificationTypeKey:   notificationsbuilder.TypeReply,
				notificationsbuilder.NotificationActionKey: notificationsbuilder.ActionOpenPost,

				notificationsbuilder.SubspaceIDKey: fmt.Sprintf("%d", reply.SubspaceID),
				notificationsbuilder.PostIDKey:     fmt.Sprintf("%d", reply.ID),
				notificationsbuilder.PostAuthorKey: reply.Author,
			},
		}
	}
}

func (d DefaultPostsNotificationsBuilder) Comment() notificationsbuilder.PostNotificationBuilder {
	return func(originalPost types.Post, reference types.Post) *notificationsbuilder.NotificationData {
		return &notificationsbuilder.NotificationData{
			Notification: &messaging.Notification{
				Title: "Someone commented your post! 💬",
				Body:  fmt.Sprintf("%s commented on your post", d.m.GetDisplayName(reference.Author)),
			},
			Data: map[string]string{
				notificationsbuilder.NotificationTypeKey:   notificationsbuilder.TypeReply,
				notificationsbuilder.NotificationActionKey: notificationsbuilder.ActionOpenPost,

				notificationsbuilder.SubspaceIDKey: fmt.Sprintf("%d", originalPost.SubspaceID),
				notificationsbuilder.PostIDKey:     fmt.Sprintf("%d", originalPost.ID),
				notificationsbuilder.PostAuthorKey: reference.Author,
			},
		}
	}
}

func (d DefaultPostsNotificationsBuilder) Repost() notificationsbuilder.PostNotificationBuilder {
	return func(originalPost types.Post, reference types.Post) *notificationsbuilder.NotificationData {
		return &notificationsbuilder.NotificationData{
			Notification: &messaging.Notification{
				Title: "Someone reposted your post! 💬",
				Body:  fmt.Sprintf("%s reposted your post", d.m.GetDisplayName(reference.Author)),
			},
			Data: map[string]string{
				notificationsbuilder.NotificationTypeKey:   notificationsbuilder.TypeRepost,
				notificationsbuilder.NotificationActionKey: notificationsbuilder.ActionOpenPost,

				notificationsbuilder.SubspaceIDKey: fmt.Sprintf("%d", originalPost.SubspaceID),
				notificationsbuilder.PostIDKey:     fmt.Sprintf("%d", originalPost.ID),
				notificationsbuilder.PostAuthorKey: reference.Author,
			},
		}
	}
}

func (d DefaultPostsNotificationsBuilder) Quote() notificationsbuilder.PostNotificationBuilder {
	return func(originalPost types.Post, reference types.Post) *notificationsbuilder.NotificationData {
		return &notificationsbuilder.NotificationData{
			Notification: &messaging.Notification{
				Title: "Someone quoted your post! 💬",
				Body:  fmt.Sprintf("%s quoted your post", d.m.GetDisplayName(reference.Author)),
			},
			Data: map[string]string{
				notificationsbuilder.NotificationTypeKey:   notificationsbuilder.TypeQuote,
				notificationsbuilder.NotificationActionKey: notificationsbuilder.ActionOpenPost,

				notificationsbuilder.SubspaceIDKey: fmt.Sprintf("%d", originalPost.SubspaceID),
				notificationsbuilder.PostIDKey:     fmt.Sprintf("%d", originalPost.ID),
				notificationsbuilder.PostAuthorKey: reference.Author,
			},
		}
	}
}

func (d DefaultPostsNotificationsBuilder) Mention() notificationsbuilder.MentionNotificationBuilder {
	return func(post types.Post, mention poststypes.TextTag) *notificationsbuilder.NotificationData {
		return &notificationsbuilder.NotificationData{
			Notification: &messaging.Notification{
				Title: "Someone mentioned you inside a post! 💬",
				Body:  fmt.Sprintf("%s mentioned you post", d.m.GetDisplayName(post.Author)),
			},
			Data: map[string]string{
				notificationsbuilder.NotificationTypeKey:   notificationsbuilder.TypeMention,
				notificationsbuilder.NotificationActionKey: notificationsbuilder.ActionOpenPost,

				notificationsbuilder.SubspaceIDKey: fmt.Sprintf("%d", post.SubspaceID),
				notificationsbuilder.PostIDKey:     fmt.Sprintf("%d", post.ID),
				notificationsbuilder.PostAuthorKey: post.Author,
			},
		}
	}
}
