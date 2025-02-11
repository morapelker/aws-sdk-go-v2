// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Summary of the membership details of an AppInstanceUser.
type AppInstanceUserMembershipSummary struct {

	// The time at which a message was last read.
	ReadMarkerTimestamp *time.Time

	// The type of ChannelMembership.
	Type ChannelMembershipType

	noSmithyDocumentSerde
}

// The membership information, including member ARNs, the channel ARN, and
// membership types.
type BatchChannelMemberships struct {

	// The ARN of the channel to which you're adding users.
	ChannelArn *string

	// The identifier of the member who invited another member.
	InvitedBy *Identity

	// The users successfully added to the request.
	Members []Identity

	// The membership types set for the channel users.
	Type ChannelMembershipType

	noSmithyDocumentSerde
}

// A list of failed member ARNs, error codes, and error messages.
type BatchCreateChannelMembershipError struct {

	// The error code.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	// The ARN of the member that the service couldn't add.
	MemberArn *string

	noSmithyDocumentSerde
}

// The details of a channel.
type Channel struct {

	// The ARN of a channel.
	ChannelArn *string

	// The AppInstanceUser who created the channel.
	CreatedBy *Identity

	// The time at which the AppInstanceUser created the channel.
	CreatedTimestamp *time.Time

	// The time at which a member sent the last message in the channel.
	LastMessageTimestamp *time.Time

	// The time at which a channel was last updated.
	LastUpdatedTimestamp *time.Time

	// The channel's metadata.
	Metadata *string

	// The mode of the channel.
	Mode ChannelMode

	// The name of a channel.
	Name *string

	// The channel's privacy setting.
	Privacy ChannelPrivacy

	noSmithyDocumentSerde
}

// The details of a channel ban.
type ChannelBan struct {

	// The ARN of the channel from which a member is being banned.
	ChannelArn *string

	// The AppInstanceUser who created the ban.
	CreatedBy *Identity

	// The time at which the ban was created.
	CreatedTimestamp *time.Time

	// The member being banned from the channel.
	Member *Identity

	noSmithyDocumentSerde
}

// Summary of the details of a ChannelBan.
type ChannelBanSummary struct {

	// The member being banned from a channel.
	Member *Identity

	noSmithyDocumentSerde
}

// The details of a channel member.
type ChannelMembership struct {

	// The ARN of the member's channel.
	ChannelArn *string

	// The time at which the channel membership was created.
	CreatedTimestamp *time.Time

	// The identifier of the member who invited another member.
	InvitedBy *Identity

	// The time at which a channel membership was last updated.
	LastUpdatedTimestamp *time.Time

	// The data of the channel member.
	Member *Identity

	// The membership type set for the channel member.
	Type ChannelMembershipType

	noSmithyDocumentSerde
}

// Summary of the channel membership details of an AppInstanceUser.
type ChannelMembershipForAppInstanceUserSummary struct {

	// Returns the channel membership data for an AppInstance.
	AppInstanceUserMembershipSummary *AppInstanceUserMembershipSummary

	// Returns the channel data for an AppInstance.
	ChannelSummary *ChannelSummary

	noSmithyDocumentSerde
}

// Summary of the details of a ChannelMembership.
type ChannelMembershipSummary struct {

	// A member's summary data.
	Member *Identity

	noSmithyDocumentSerde
}

// The details of a message in a channel.
type ChannelMessage struct {

	// The ARN of the channel.
	ChannelArn *string

	// The message content.
	Content *string

	// The time at which the message was created.
	CreatedTimestamp *time.Time

	// The time at which a message was edited.
	LastEditedTimestamp *time.Time

	// The time at which a message was updated.
	LastUpdatedTimestamp *time.Time

	// The ID of a message.
	MessageId *string

	// The message metadata.
	Metadata *string

	// The persistence setting for a channel message.
	Persistence ChannelMessagePersistenceType

	// Hides the content of a message.
	Redacted bool

	// The message sender.
	Sender *Identity

	// The message type.
	Type ChannelMessageType

	noSmithyDocumentSerde
}

// Summary of the messages in a Channel.
type ChannelMessageSummary struct {

	// The content of the message.
	Content *string

	// The time at which the message summary was created.
	CreatedTimestamp *time.Time

	// The time at which a message was last edited.
	LastEditedTimestamp *time.Time

	// The time at which a message was last updated.
	LastUpdatedTimestamp *time.Time

	// The ID of the message.
	MessageId *string

	// The metadata of the message.
	Metadata *string

	// Indicates whether a message was redacted.
	Redacted bool

	// The message sender.
	Sender *Identity

	// The type of message.
	Type ChannelMessageType

	noSmithyDocumentSerde
}

// Summary of the details of a moderated channel.
type ChannelModeratedByAppInstanceUserSummary struct {

	// Summary of the details of a Channel.
	ChannelSummary *ChannelSummary

	noSmithyDocumentSerde
}

// The details of a channel moderator.
type ChannelModerator struct {

	// The ARN of the moderator's channel.
	ChannelArn *string

	// The AppInstanceUser who created the moderator.
	CreatedBy *Identity

	// The time at which the moderator was created.
	CreatedTimestamp *time.Time

	// The moderator's data.
	Moderator *Identity

	noSmithyDocumentSerde
}

// Summary of the details of a ChannelModerator.
type ChannelModeratorSummary struct {

	// The data for a moderator.
	Moderator *Identity

	noSmithyDocumentSerde
}

// Summary of the details of a Channel.
type ChannelSummary struct {

	// The ARN of the channel.
	ChannelArn *string

	// The time at which the last message in a channel was sent.
	LastMessageTimestamp *time.Time

	// The metadata of the channel.
	Metadata *string

	// The mode of the channel.
	Mode ChannelMode

	// The name of the channel.
	Name *string

	// The privacy setting of the channel.
	Privacy ChannelPrivacy

	noSmithyDocumentSerde
}

// The details of a user.
type Identity struct {

	// The ARN in an Identity.
	Arn *string

	// The name in an Identity.
	Name *string

	noSmithyDocumentSerde
}

// The websocket endpoint used to connect to Amazon Chime SDK messaging.
type MessagingSessionEndpoint struct {

	// The endpoint to which you establish a websocket connection.
	Url *string

	noSmithyDocumentSerde
}

// Describes a tag applied to a resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
