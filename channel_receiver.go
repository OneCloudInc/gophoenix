package gophoenix

// ChannelReceiver receives messages for each message type.
type ChannelReceiver interface {
	// Invoked after the client has successfully joined a topic.
	OnJoin(payload interface{})
	// Invoked if the server has refused a topic join request.
	OnJoinError(payload interface{})
	// Invoked after the server closes a Channel.
	OnChannelClose(payload interface{}, joinRef int64)
	// Invoked when a message from the server arrives.
	OnMessage(ref int64, event string, payload interface{})
}
