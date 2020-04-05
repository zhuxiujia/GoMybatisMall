package enum

type UserRedPacketStatus int

const (
	UserRedPacketStatus_UnUse     UserRedPacketStatus = 0  //未使用
	UserRedPacketStatus_Used      UserRedPacketStatus = 1  //已使用
	UserRedPacketStatus_OutOfTime UserRedPacketStatus = -1 //已过期
)
