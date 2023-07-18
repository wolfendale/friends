package types

import (
	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
)

// FriendUser represents a user connected to the friends server
// Contains data relating only to friends
type FriendUser struct {
	NNID              string
	PID               uint32
	Comment           *friends_wiiu_types.Comment
	FriendRequestsOut []*friends_wiiu_types.FriendRequest
	FriendRequestsIn  []*friends_wiiu_types.FriendRequest
	BlockedUsers      []*friends_wiiu_types.BlacklistedPrincipal
	LastOnline        *nex.DateTime
	ActiveTitle       *friends_wiiu_types.GameKey
	Notifications     []*friends_wiiu_types.PersistentNotification
}

func (friendUser *FriendUser) FromPID(pid uint32) {

}

func NewFriendUser() *FriendUser {
	return &FriendUser{}
}
