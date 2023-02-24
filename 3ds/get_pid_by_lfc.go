package friends_3ds

import (
	"github.com/PretendoNetwork/friends-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	nexproto "github.com/PretendoNetwork/nex-protocols-go"
)

func GetPrincipalIDByLocalFriendCode(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64) {
	// Respond with unimplemented, waiting for gRPC to retrieve PID from account server

	rmcResponse := nex.NewRMCResponse(nexproto.Friends3DSProtocolID, callID)
	rmcResponse.SetError(nex.Errors.Core.NotImplemented)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV0(client, nil)

	responsePacket.SetVersion(0)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.NEXServer.Send(responsePacket)
}
