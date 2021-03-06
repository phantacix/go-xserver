package nodemgr

import (
	"context"
	"net"

	"github.com/fananchong/go-xserver/common/config"
	nodecommon "github.com/fananchong/go-xserver/internal/components/node/common"
	"github.com/fananchong/go-xserver/internal/protocol"
	"github.com/fananchong/gotcp"
)

// Session : 网络会话类
type Session struct {
	*nodecommon.SessionBase
}

// Init : 初始化网络会话节点
func (sess *Session) Init(root context.Context, conn net.Conn, derived gotcp.ISession, userdata interface{}) {
	ud := userdata.(*nodecommon.UserData)
	sess.SessionBase = nodecommon.NewSessionBase(ud.Ctx, sess)
	sess.SessionBase.Init(root, conn, derived)
	sess.SessMgr = ud.SessMgr
}

// DoVerify : 验证时保存自己的注册消息
func (sess *Session) DoVerify(msg *protocol.MSG_MGR_REGISTER_SERVER) {
	sess.Info = msg.GetData()
	sess.CacheRegisterMsg = msg
}

// DoRegister : 某节点注册时处理
func (sess *Session) DoRegister(msg *protocol.MSG_MGR_REGISTER_SERVER) {
	if nodecommon.EqualSID(sess.Info.GetId(), msg.GetData().GetId()) == false {
		sess.Close()
		return
	}
	if msg.GetTargetServerType() != uint32(config.Mgr) {
		sess.Close()
		return
	}
	sess.Info = msg.GetData()
	sess.Ctx.Infoln("The service node registers with me, the node ID is ", msg.GetData().GetId().GetID())
	sess.Ctx.Infoln(sess.Info)

	sess.SessMgr.Register(sess.SessionBase)
	sess.SessMgr.ForAll(func(elem *nodecommon.SessionBase) {
		if elem != sess.SessionBase {
			sess.CacheRegisterMsg.TargetServerType = uint32(elem.GetType())
			sess.CacheRegisterMsg.TargetServerID = nodecommon.NodeID2ServerID(elem.GetID())
			elem.SendMsg(uint64(protocol.CMD_MGR_REGISTER_SERVER), sess.CacheRegisterMsg)
		}
	})
	sess.SessMgr.ForAll(func(elem *nodecommon.SessionBase) {
		if elem != sess.SessionBase {
			elem.CacheRegisterMsg.TargetServerType = uint32(sess.GetType())
			elem.CacheRegisterMsg.TargetServerID = nodecommon.NodeID2ServerID(sess.GetID())
			sess.SendMsg(uint64(protocol.CMD_MGR_REGISTER_SERVER), elem.CacheRegisterMsg)
		}
	})
}

// DoLose : 节点丢失时处理
func (sess *Session) DoLose(msg *protocol.MSG_MGR_LOSE_SERVER) {
}

// DoClose : 节点关闭时处理
func (sess *Session) DoClose(sessbase *nodecommon.SessionBase) {
	if sess.SessionBase == sessbase && sessbase.Info != nil {
		sess.SessMgr.Lose1(sessbase)
		msg := &protocol.MSG_MGR_LOSE_SERVER{}
		msg.Id = sess.Info.GetId()
		msg.Type = sess.Info.GetType()
		sess.SessMgr.ForAll(func(elem *nodecommon.SessionBase) {
			elem.SendMsg(uint64(protocol.CMD_MGR_LOSE_SERVER), msg)
		})
		sess.Ctx.Infoln("Service node loses connection, type:", msg.Type, "id:", msg.Id.GetID())
	}
}

// DoRecv : 节点收到消息处理
func (sess *Session) DoRecv(cmd uint64, data []byte, flag byte) (done bool) {
	return
}
