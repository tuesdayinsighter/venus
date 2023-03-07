// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package v1

import (
	"context"
	"time"

	address "github.com/filecoin-project/go-address"
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/piecestore"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/paych"
	cid "github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/filecoin-project/venus/venus-shared/types/gateway"
	"github.com/filecoin-project/venus/venus-shared/types/market"
)

type IMarketStruct struct {
	Internal struct {
		ActorDelete                             func(context.Context, address.Address) error                                                                                                                                                        `perm:"admin"`
		ActorExist                              func(ctx context.Context, addr address.Address) (bool, error)                                                                                                                                       `perm:"read"`
		ActorList                               func(context.Context) ([]market.User, error)                                                                                                                                                        `perm:"read"`
		ActorSectorSize                         func(context.Context, address.Address) (abi.SectorSize, error)                                                                                                                                      `perm:"read"`
		ActorUpsert                             func(context.Context, market.User) (bool, error)                                                                                                                                                    `perm:"admin"`
		AddFsPieceStorage                       func(ctx context.Context, name string, path string, readonly bool) error                                                                                                                            `perm:"admin"`
		AddS3PieceStorage                       func(ctx context.Context, name, endpoit, bucket, subdir, accessKey, secretKey, token string, readonly bool) error                                                                                   `perm:"admin"`
		AssignUnPackedDeals                     func(ctx context.Context, sid abi.SectorID, ssize abi.SectorSize, spec *market.GetDealSpec) ([]*market.DealInfoIncludePath, error)                                                                  `perm:"write"`
		DagstoreGC                              func(ctx context.Context) ([]market.DagstoreShardResult, error)                                                                                                                                     `perm:"admin"`
		DagstoreInitializeAll                   func(ctx context.Context, params market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error)                                                                              `perm:"admin"`
		DagstoreInitializeShard                 func(ctx context.Context, key string) error                                                                                                                                                         `perm:"admin"`
		DagstoreInitializeStorage               func(context.Context, string, market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error)                                                                                 `perm:"admin"`
		DagstoreListShards                      func(ctx context.Context) ([]market.DagstoreShardInfo, error)                                                                                                                                       `perm:"admin"`
		DagstoreRecoverShard                    func(ctx context.Context, key string) error                                                                                                                                                         `perm:"admin"`
		DealsConsiderOfflineRetrievalDeals      func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsConsiderOfflineStorageDeals        func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsConsiderOnlineRetrievalDeals       func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsConsiderOnlineStorageDeals         func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsConsiderUnverifiedStorageDeals     func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsConsiderVerifiedStorageDeals       func(context.Context, address.Address) (bool, error)                                                                                                                                                `perm:"read"`
		DealsImportData                         func(ctx context.Context, dealPropCid cid.Cid, file string) error                                                                                                                                   `perm:"admin"`
		DealsMaxProviderCollateralMultiplier    func(context.Context, address.Address) (uint64, error)                                                                                                                                              `perm:"read"`
		DealsMaxPublishFee                      func(context.Context, address.Address) (types.FIL, error)                                                                                                                                           `perm:"read"`
		DealsMaxStartDelay                      func(context.Context, address.Address) (time.Duration, error)                                                                                                                                       `perm:"read"`
		DealsPieceCidBlocklist                  func(context.Context, address.Address) ([]cid.Cid, error)                                                                                                                                           `perm:"read"`
		DealsPublishMsgPeriod                   func(context.Context, address.Address) (time.Duration, error)                                                                                                                                       `perm:"read"`
		DealsSetConsiderOfflineRetrievalDeals   func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetConsiderOfflineStorageDeals     func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetConsiderOnlineRetrievalDeals    func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetConsiderOnlineStorageDeals      func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetConsiderUnverifiedStorageDeals  func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetConsiderVerifiedStorageDeals    func(context.Context, address.Address, bool) error                                                                                                                                                  `perm:"write"`
		DealsSetMaxProviderCollateralMultiplier func(context.Context, address.Address, uint64) error                                                                                                                                                `perm:"write"`
		DealsSetMaxPublishFee                   func(context.Context, address.Address, types.FIL) error                                                                                                                                             `perm:"write"`
		DealsSetMaxStartDelay                   func(context.Context, address.Address, time.Duration) error                                                                                                                                         `perm:"write"`
		DealsSetPieceCidBlocklist               func(context.Context, address.Address, []cid.Cid) error                                                                                                                                             `perm:"write"`
		DealsSetPublishMsgPeriod                func(context.Context, address.Address, time.Duration) error                                                                                                                                         `perm:"write"`
		GetDeals                                func(ctx context.Context, miner address.Address, pageIndex, pageSize int) ([]*market.DealInfo, error)                                                                                               `perm:"read"`
		GetRetrievalDealStatistic               func(ctx context.Context, miner address.Address) (*market.RetrievalDealStatistic, error)                                                                                                            `perm:"read"`
		GetStorageDealStatistic                 func(ctx context.Context, miner address.Address) (*market.StorageDealStatistic, error)                                                                                                              `perm:"read"`
		GetUnPackedDeals                        func(ctx context.Context, miner address.Address, spec *market.GetDealSpec) ([]*market.DealInfoIncludePath, error)                                                                                   `perm:"read"`
		ID                                      func(context.Context) (peer.ID, error)                                                                                                                                                              `perm:"read"`
		ListPieceStorageInfos                   func(ctx context.Context) market.PieceStorageInfos                                                                                                                                                  `perm:"read"`
		ListenMarketEvent                       func(ctx context.Context, policy *gateway.MarketRegisterPolicy) (<-chan *gateway.RequestEvent, error)                                                                                               `perm:"read"`
		MarkDealsAsPacking                      func(ctx context.Context, miner address.Address, deals []abi.DealID) error                                                                                                                          `perm:"write"`
		MarketAddBalance                        func(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error)                                                                                                          `perm:"sign"`
		MarketCancelDataTransfer                func(ctx context.Context, transferID datatransfer.TransferID, otherPeer peer.ID, isInitiator bool) error                                                                                            `perm:"admin"`
		MarketDataTransferPath                  func(context.Context, address.Address) (string, error)                                                                                                                                              `perm:"admin"`
		MarketDataTransferUpdates               func(ctx context.Context) (<-chan market.DataTransferChannel, error)                                                                                                                                `perm:"admin"`
		MarketGetAsk                            func(ctx context.Context, mAddr address.Address) (*market.SignedStorageAsk, error)                                                                                                                  `perm:"read"`
		MarketGetDealUpdates                    func(ctx context.Context) (<-chan market.MinerDeal, error)                                                                                                                                          `perm:"admin"`
		MarketGetReserved                       func(ctx context.Context, addr address.Address) (types.BigInt, error)                                                                                                                               `perm:"sign"`
		MarketGetRetrievalAsk                   func(ctx context.Context, mAddr address.Address) (*retrievalmarket.Ask, error)                                                                                                                      `perm:"read"`
		MarketImportDealData                    func(ctx context.Context, propcid cid.Cid, path string) error                                                                                                                                       `perm:"admin"`
		MarketImportPublishedDeal               func(ctx context.Context, deal market.MinerDeal) error                                                                                                                                              `perm:"write"`
		MarketListDataTransfers                 func(ctx context.Context) ([]market.DataTransferChannel, error)                                                                                                                                     `perm:"admin"`
		MarketListDeals                         func(ctx context.Context, addrs []address.Address) ([]*types.MarketDeal, error)                                                                                                                     `perm:"read"`
		MarketListIncompleteDeals               func(ctx context.Context, mAddr address.Address) ([]market.MinerDeal, error)                                                                                                                        `perm:"read"`
		MarketListRetrievalAsk                  func(ctx context.Context) ([]*market.RetrievalAsk, error)                                                                                                                                           `perm:"read"`
		MarketListRetrievalDeals                func(ctx context.Context) ([]market.ProviderDealState, error)                                                                                                                                       `perm:"read"`
		MarketListStorageAsk                    func(ctx context.Context) ([]*market.SignedStorageAsk, error)                                                                                                                                       `perm:"read"`
		MarketMaxBalanceAddFee                  func(context.Context, address.Address) (types.FIL, error)                                                                                                                                           `perm:"read"`
		MarketMaxDealsPerPublishMsg             func(context.Context, address.Address) (uint64, error)                                                                                                                                              `perm:"read"`
		MarketPendingDeals                      func(ctx context.Context) ([]market.PendingDealInfo, error)                                                                                                                                         `perm:"write"`
		MarketPublishPendingDeals               func(ctx context.Context) error                                                                                                                                                                     `perm:"admin"`
		MarketReleaseFunds                      func(ctx context.Context, addr address.Address, amt types.BigInt) error                                                                                                                             `perm:"sign"`
		MarketReserveFunds                      func(ctx context.Context, wallet address.Address, addr address.Address, amt types.BigInt) (cid.Cid, error)                                                                                          `perm:"sign"`
		MarketRestartDataTransfer               func(ctx context.Context, transferID datatransfer.TransferID, otherPeer peer.ID, isInitiator bool) error                                                                                            `perm:"admin"`
		MarketSetAsk                            func(ctx context.Context, mAddr address.Address, price types.BigInt, verifiedPrice types.BigInt, duration abi.ChainEpoch, minPieceSize abi.PaddedPieceSize, maxPieceSize abi.PaddedPieceSize) error `perm:"admin"`
		MarketSetDataTransferPath               func(context.Context, address.Address, string) error                                                                                                                                                `perm:"admin"`
		MarketSetMaxBalanceAddFee               func(context.Context, address.Address, types.FIL) error                                                                                                                                             `perm:"write"`
		MarketSetMaxDealsPerPublishMsg          func(context.Context, address.Address, uint64) error                                                                                                                                                `perm:"write"`
		MarketSetRetrievalAsk                   func(ctx context.Context, mAddr address.Address, rask *retrievalmarket.Ask) error                                                                                                                   `perm:"admin"`
		MarketWithdraw                          func(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error)                                                                                                          `perm:"sign"`
		MessagerGetMessage                      func(ctx context.Context, mid cid.Cid) (*types.Message, error)                                                                                                                                      `perm:"read"`
		MessagerPushMessage                     func(ctx context.Context, msg *types.Message, meta *types.MessageSendSpec) (cid.Cid, error)                                                                                                         `perm:"write"`
		MessagerWaitMessage                     func(ctx context.Context, mid cid.Cid) (*types.MsgLookup, error)                                                                                                                                    `perm:"read"`
		NetAddrsListen                          func(context.Context) (peer.AddrInfo, error)                                                                                                                                                        `perm:"read"`
		OfflineDealImport                       func(ctx context.Context, deal market.MinerDeal) error                                                                                                                                              `perm:"admin"`
		PaychVoucherList                        func(ctx context.Context, pch address.Address) ([]*paych.SignedVoucher, error)                                                                                                                      `perm:"read"`
		PiecesGetCIDInfo                        func(ctx context.Context, payloadCid cid.Cid) (*piecestore.CIDInfo, error)                                                                                                                          `perm:"read"`
		PiecesGetPieceInfo                      func(ctx context.Context, pieceCid cid.Cid) (*piecestore.PieceInfo, error)                                                                                                                          `perm:"read"`
		PiecesListCidInfos                      func(ctx context.Context) ([]cid.Cid, error)                                                                                                                                                        `perm:"read"`
		PiecesListPieces                        func(ctx context.Context) ([]cid.Cid, error)                                                                                                                                                        `perm:"read"`
		RemovePieceStorage                      func(ctx context.Context, name string) error                                                                                                                                                        `perm:"admin"`
		ResponseMarketEvent                     func(ctx context.Context, resp *gateway.ResponseEvent) error                                                                                                                                        `perm:"read"`
		SectorGetExpectedSealDuration           func(context.Context, address.Address) (time.Duration, error)                                                                                                                                       `perm:"read"`
		SectorSetExpectedSealDuration           func(context.Context, address.Address, time.Duration) error                                                                                                                                         `perm:"write"`
		UpdateDealOnPacking                     func(ctx context.Context, miner address.Address, dealID abi.DealID, sectorid abi.SectorNumber, offset abi.PaddedPieceSize) error                                                                    `perm:"write"`
		UpdateDealStatus                        func(ctx context.Context, miner address.Address, dealID abi.DealID, pieceStatus market.PieceStatus, dealStatus storagemarket.StorageDealStatus) error                                               `perm:"write"`
		UpdateStorageDealStatus                 func(ctx context.Context, dealProposalCid cid.Cid, state storagemarket.StorageDealStatus, pieceState market.PieceStatus) error                                                                      `perm:"write"`
		Version                                 func(ctx context.Context) (types.Version, error)                                                                                                                                                    `perm:"read"`
	}
}

func (s *IMarketStruct) ActorDelete(p0 context.Context, p1 address.Address) error {
	return s.Internal.ActorDelete(p0, p1)
}
func (s *IMarketStruct) ActorExist(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.ActorExist(p0, p1)
}
func (s *IMarketStruct) ActorList(p0 context.Context) ([]market.User, error) {
	return s.Internal.ActorList(p0)
}
func (s *IMarketStruct) ActorSectorSize(p0 context.Context, p1 address.Address) (abi.SectorSize, error) {
	return s.Internal.ActorSectorSize(p0, p1)
}
func (s *IMarketStruct) ActorUpsert(p0 context.Context, p1 market.User) (bool, error) {
	return s.Internal.ActorUpsert(p0, p1)
}
func (s *IMarketStruct) AddFsPieceStorage(p0 context.Context, p1 string, p2 string, p3 bool) error {
	return s.Internal.AddFsPieceStorage(p0, p1, p2, p3)
}
func (s *IMarketStruct) AddS3PieceStorage(p0 context.Context, p1, p2, p3, p4, p5, p6, p7 string, p8 bool) error {
	return s.Internal.AddS3PieceStorage(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}
func (s *IMarketStruct) AssignUnPackedDeals(p0 context.Context, p1 abi.SectorID, p2 abi.SectorSize, p3 *market.GetDealSpec) ([]*market.DealInfoIncludePath, error) {
	return s.Internal.AssignUnPackedDeals(p0, p1, p2, p3)
}
func (s *IMarketStruct) DagstoreGC(p0 context.Context) ([]market.DagstoreShardResult, error) {
	return s.Internal.DagstoreGC(p0)
}
func (s *IMarketStruct) DagstoreInitializeAll(p0 context.Context, p1 market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error) {
	return s.Internal.DagstoreInitializeAll(p0, p1)
}
func (s *IMarketStruct) DagstoreInitializeShard(p0 context.Context, p1 string) error {
	return s.Internal.DagstoreInitializeShard(p0, p1)
}
func (s *IMarketStruct) DagstoreInitializeStorage(p0 context.Context, p1 string, p2 market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error) {
	return s.Internal.DagstoreInitializeStorage(p0, p1, p2)
}
func (s *IMarketStruct) DagstoreListShards(p0 context.Context) ([]market.DagstoreShardInfo, error) {
	return s.Internal.DagstoreListShards(p0)
}
func (s *IMarketStruct) DagstoreRecoverShard(p0 context.Context, p1 string) error {
	return s.Internal.DagstoreRecoverShard(p0, p1)
}
func (s *IMarketStruct) DealsConsiderOfflineRetrievalDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderOfflineRetrievalDeals(p0, p1)
}
func (s *IMarketStruct) DealsConsiderOfflineStorageDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderOfflineStorageDeals(p0, p1)
}
func (s *IMarketStruct) DealsConsiderOnlineRetrievalDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderOnlineRetrievalDeals(p0, p1)
}
func (s *IMarketStruct) DealsConsiderOnlineStorageDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderOnlineStorageDeals(p0, p1)
}
func (s *IMarketStruct) DealsConsiderUnverifiedStorageDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderUnverifiedStorageDeals(p0, p1)
}
func (s *IMarketStruct) DealsConsiderVerifiedStorageDeals(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.DealsConsiderVerifiedStorageDeals(p0, p1)
}
func (s *IMarketStruct) DealsImportData(p0 context.Context, p1 cid.Cid, p2 string) error {
	return s.Internal.DealsImportData(p0, p1, p2)
}
func (s *IMarketStruct) DealsMaxProviderCollateralMultiplier(p0 context.Context, p1 address.Address) (uint64, error) {
	return s.Internal.DealsMaxProviderCollateralMultiplier(p0, p1)
}
func (s *IMarketStruct) DealsMaxPublishFee(p0 context.Context, p1 address.Address) (types.FIL, error) {
	return s.Internal.DealsMaxPublishFee(p0, p1)
}
func (s *IMarketStruct) DealsMaxStartDelay(p0 context.Context, p1 address.Address) (time.Duration, error) {
	return s.Internal.DealsMaxStartDelay(p0, p1)
}
func (s *IMarketStruct) DealsPieceCidBlocklist(p0 context.Context, p1 address.Address) ([]cid.Cid, error) {
	return s.Internal.DealsPieceCidBlocklist(p0, p1)
}
func (s *IMarketStruct) DealsPublishMsgPeriod(p0 context.Context, p1 address.Address) (time.Duration, error) {
	return s.Internal.DealsPublishMsgPeriod(p0, p1)
}
func (s *IMarketStruct) DealsSetConsiderOfflineRetrievalDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderOfflineRetrievalDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetConsiderOfflineStorageDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderOfflineStorageDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetConsiderOnlineRetrievalDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderOnlineRetrievalDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetConsiderOnlineStorageDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderOnlineStorageDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetConsiderUnverifiedStorageDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderUnverifiedStorageDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetConsiderVerifiedStorageDeals(p0 context.Context, p1 address.Address, p2 bool) error {
	return s.Internal.DealsSetConsiderVerifiedStorageDeals(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetMaxProviderCollateralMultiplier(p0 context.Context, p1 address.Address, p2 uint64) error {
	return s.Internal.DealsSetMaxProviderCollateralMultiplier(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetMaxPublishFee(p0 context.Context, p1 address.Address, p2 types.FIL) error {
	return s.Internal.DealsSetMaxPublishFee(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetMaxStartDelay(p0 context.Context, p1 address.Address, p2 time.Duration) error {
	return s.Internal.DealsSetMaxStartDelay(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetPieceCidBlocklist(p0 context.Context, p1 address.Address, p2 []cid.Cid) error {
	return s.Internal.DealsSetPieceCidBlocklist(p0, p1, p2)
}
func (s *IMarketStruct) DealsSetPublishMsgPeriod(p0 context.Context, p1 address.Address, p2 time.Duration) error {
	return s.Internal.DealsSetPublishMsgPeriod(p0, p1, p2)
}
func (s *IMarketStruct) GetDeals(p0 context.Context, p1 address.Address, p2, p3 int) ([]*market.DealInfo, error) {
	return s.Internal.GetDeals(p0, p1, p2, p3)
}
func (s *IMarketStruct) GetRetrievalDealStatistic(p0 context.Context, p1 address.Address) (*market.RetrievalDealStatistic, error) {
	return s.Internal.GetRetrievalDealStatistic(p0, p1)
}
func (s *IMarketStruct) GetStorageDealStatistic(p0 context.Context, p1 address.Address) (*market.StorageDealStatistic, error) {
	return s.Internal.GetStorageDealStatistic(p0, p1)
}
func (s *IMarketStruct) GetUnPackedDeals(p0 context.Context, p1 address.Address, p2 *market.GetDealSpec) ([]*market.DealInfoIncludePath, error) {
	return s.Internal.GetUnPackedDeals(p0, p1, p2)
}
func (s *IMarketStruct) ID(p0 context.Context) (peer.ID, error) { return s.Internal.ID(p0) }
func (s *IMarketStruct) ListPieceStorageInfos(p0 context.Context) market.PieceStorageInfos {
	return s.Internal.ListPieceStorageInfos(p0)
}
func (s *IMarketStruct) ListenMarketEvent(p0 context.Context, p1 *gateway.MarketRegisterPolicy) (<-chan *gateway.RequestEvent, error) {
	return s.Internal.ListenMarketEvent(p0, p1)
}
func (s *IMarketStruct) MarkDealsAsPacking(p0 context.Context, p1 address.Address, p2 []abi.DealID) error {
	return s.Internal.MarkDealsAsPacking(p0, p1, p2)
}
func (s *IMarketStruct) MarketAddBalance(p0 context.Context, p1, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketAddBalance(p0, p1, p2, p3)
}
func (s *IMarketStruct) MarketCancelDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return s.Internal.MarketCancelDataTransfer(p0, p1, p2, p3)
}
func (s *IMarketStruct) MarketDataTransferPath(p0 context.Context, p1 address.Address) (string, error) {
	return s.Internal.MarketDataTransferPath(p0, p1)
}
func (s *IMarketStruct) MarketDataTransferUpdates(p0 context.Context) (<-chan market.DataTransferChannel, error) {
	return s.Internal.MarketDataTransferUpdates(p0)
}
func (s *IMarketStruct) MarketGetAsk(p0 context.Context, p1 address.Address) (*market.SignedStorageAsk, error) {
	return s.Internal.MarketGetAsk(p0, p1)
}
func (s *IMarketStruct) MarketGetDealUpdates(p0 context.Context) (<-chan market.MinerDeal, error) {
	return s.Internal.MarketGetDealUpdates(p0)
}
func (s *IMarketStruct) MarketGetReserved(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return s.Internal.MarketGetReserved(p0, p1)
}
func (s *IMarketStruct) MarketGetRetrievalAsk(p0 context.Context, p1 address.Address) (*retrievalmarket.Ask, error) {
	return s.Internal.MarketGetRetrievalAsk(p0, p1)
}
func (s *IMarketStruct) MarketImportDealData(p0 context.Context, p1 cid.Cid, p2 string) error {
	return s.Internal.MarketImportDealData(p0, p1, p2)
}
func (s *IMarketStruct) MarketImportPublishedDeal(p0 context.Context, p1 market.MinerDeal) error {
	return s.Internal.MarketImportPublishedDeal(p0, p1)
}
func (s *IMarketStruct) MarketListDataTransfers(p0 context.Context) ([]market.DataTransferChannel, error) {
	return s.Internal.MarketListDataTransfers(p0)
}
func (s *IMarketStruct) MarketListDeals(p0 context.Context, p1 []address.Address) ([]*types.MarketDeal, error) {
	return s.Internal.MarketListDeals(p0, p1)
}
func (s *IMarketStruct) MarketListIncompleteDeals(p0 context.Context, p1 address.Address) ([]market.MinerDeal, error) {
	return s.Internal.MarketListIncompleteDeals(p0, p1)
}
func (s *IMarketStruct) MarketListRetrievalAsk(p0 context.Context) ([]*market.RetrievalAsk, error) {
	return s.Internal.MarketListRetrievalAsk(p0)
}
func (s *IMarketStruct) MarketListRetrievalDeals(p0 context.Context) ([]market.ProviderDealState, error) {
	return s.Internal.MarketListRetrievalDeals(p0)
}
func (s *IMarketStruct) MarketListStorageAsk(p0 context.Context) ([]*market.SignedStorageAsk, error) {
	return s.Internal.MarketListStorageAsk(p0)
}
func (s *IMarketStruct) MarketMaxBalanceAddFee(p0 context.Context, p1 address.Address) (types.FIL, error) {
	return s.Internal.MarketMaxBalanceAddFee(p0, p1)
}
func (s *IMarketStruct) MarketMaxDealsPerPublishMsg(p0 context.Context, p1 address.Address) (uint64, error) {
	return s.Internal.MarketMaxDealsPerPublishMsg(p0, p1)
}
func (s *IMarketStruct) MarketPendingDeals(p0 context.Context) ([]market.PendingDealInfo, error) {
	return s.Internal.MarketPendingDeals(p0)
}
func (s *IMarketStruct) MarketPublishPendingDeals(p0 context.Context) error {
	return s.Internal.MarketPublishPendingDeals(p0)
}
func (s *IMarketStruct) MarketReleaseFunds(p0 context.Context, p1 address.Address, p2 types.BigInt) error {
	return s.Internal.MarketReleaseFunds(p0, p1, p2)
}
func (s *IMarketStruct) MarketReserveFunds(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketReserveFunds(p0, p1, p2, p3)
}
func (s *IMarketStruct) MarketRestartDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return s.Internal.MarketRestartDataTransfer(p0, p1, p2, p3)
}
func (s *IMarketStruct) MarketSetAsk(p0 context.Context, p1 address.Address, p2 types.BigInt, p3 types.BigInt, p4 abi.ChainEpoch, p5 abi.PaddedPieceSize, p6 abi.PaddedPieceSize) error {
	return s.Internal.MarketSetAsk(p0, p1, p2, p3, p4, p5, p6)
}
func (s *IMarketStruct) MarketSetDataTransferPath(p0 context.Context, p1 address.Address, p2 string) error {
	return s.Internal.MarketSetDataTransferPath(p0, p1, p2)
}
func (s *IMarketStruct) MarketSetMaxBalanceAddFee(p0 context.Context, p1 address.Address, p2 types.FIL) error {
	return s.Internal.MarketSetMaxBalanceAddFee(p0, p1, p2)
}
func (s *IMarketStruct) MarketSetMaxDealsPerPublishMsg(p0 context.Context, p1 address.Address, p2 uint64) error {
	return s.Internal.MarketSetMaxDealsPerPublishMsg(p0, p1, p2)
}
func (s *IMarketStruct) MarketSetRetrievalAsk(p0 context.Context, p1 address.Address, p2 *retrievalmarket.Ask) error {
	return s.Internal.MarketSetRetrievalAsk(p0, p1, p2)
}
func (s *IMarketStruct) MarketWithdraw(p0 context.Context, p1, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketWithdraw(p0, p1, p2, p3)
}
func (s *IMarketStruct) MessagerGetMessage(p0 context.Context, p1 cid.Cid) (*types.Message, error) {
	return s.Internal.MessagerGetMessage(p0, p1)
}
func (s *IMarketStruct) MessagerPushMessage(p0 context.Context, p1 *types.Message, p2 *types.MessageSendSpec) (cid.Cid, error) {
	return s.Internal.MessagerPushMessage(p0, p1, p2)
}
func (s *IMarketStruct) MessagerWaitMessage(p0 context.Context, p1 cid.Cid) (*types.MsgLookup, error) {
	return s.Internal.MessagerWaitMessage(p0, p1)
}
func (s *IMarketStruct) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	return s.Internal.NetAddrsListen(p0)
}
func (s *IMarketStruct) OfflineDealImport(p0 context.Context, p1 market.MinerDeal) error {
	return s.Internal.OfflineDealImport(p0, p1)
}
func (s *IMarketStruct) PaychVoucherList(p0 context.Context, p1 address.Address) ([]*paych.SignedVoucher, error) {
	return s.Internal.PaychVoucherList(p0, p1)
}
func (s *IMarketStruct) PiecesGetCIDInfo(p0 context.Context, p1 cid.Cid) (*piecestore.CIDInfo, error) {
	return s.Internal.PiecesGetCIDInfo(p0, p1)
}
func (s *IMarketStruct) PiecesGetPieceInfo(p0 context.Context, p1 cid.Cid) (*piecestore.PieceInfo, error) {
	return s.Internal.PiecesGetPieceInfo(p0, p1)
}
func (s *IMarketStruct) PiecesListCidInfos(p0 context.Context) ([]cid.Cid, error) {
	return s.Internal.PiecesListCidInfos(p0)
}
func (s *IMarketStruct) PiecesListPieces(p0 context.Context) ([]cid.Cid, error) {
	return s.Internal.PiecesListPieces(p0)
}
func (s *IMarketStruct) RemovePieceStorage(p0 context.Context, p1 string) error {
	return s.Internal.RemovePieceStorage(p0, p1)
}
func (s *IMarketStruct) ResponseMarketEvent(p0 context.Context, p1 *gateway.ResponseEvent) error {
	return s.Internal.ResponseMarketEvent(p0, p1)
}
func (s *IMarketStruct) SectorGetExpectedSealDuration(p0 context.Context, p1 address.Address) (time.Duration, error) {
	return s.Internal.SectorGetExpectedSealDuration(p0, p1)
}
func (s *IMarketStruct) SectorSetExpectedSealDuration(p0 context.Context, p1 address.Address, p2 time.Duration) error {
	return s.Internal.SectorSetExpectedSealDuration(p0, p1, p2)
}
func (s *IMarketStruct) UpdateDealOnPacking(p0 context.Context, p1 address.Address, p2 abi.DealID, p3 abi.SectorNumber, p4 abi.PaddedPieceSize) error {
	return s.Internal.UpdateDealOnPacking(p0, p1, p2, p3, p4)
}
func (s *IMarketStruct) UpdateDealStatus(p0 context.Context, p1 address.Address, p2 abi.DealID, p3 market.PieceStatus, p4 storagemarket.StorageDealStatus) error {
	return s.Internal.UpdateDealStatus(p0, p1, p2, p3, p4)
}
func (s *IMarketStruct) UpdateStorageDealStatus(p0 context.Context, p1 cid.Cid, p2 storagemarket.StorageDealStatus, p3 market.PieceStatus) error {
	return s.Internal.UpdateStorageDealStatus(p0, p1, p2, p3)
}
func (s *IMarketStruct) Version(p0 context.Context) (types.Version, error) {
	return s.Internal.Version(p0)
}
