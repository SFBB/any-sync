package syncacl

import (
	"github.com/anyproto/any-sync/commonspace/sync/objectsync/objectmessages"
	"github.com/anyproto/any-sync/consensus/consensusproto"
)

type InnerRequest struct {
	head string
	root *consensusproto.RawRecordWithId
}

func NewRequest(peerId, objectId, spaceId, head string, root *consensusproto.RawRecordWithId) *objectmessages.Request {
	return objectmessages.NewRequest(peerId, spaceId, objectId, &InnerRequest{
		head: head,
		root: root,
	})
}

func (r *InnerRequest) Marshall() ([]byte, error) {
	req := &consensusproto.LogFullSyncRequest{
		Head: r.head,
	}
	fullSync := consensusproto.WrapFullRequest(req, r.root)
	return fullSync.Marshal()
}
