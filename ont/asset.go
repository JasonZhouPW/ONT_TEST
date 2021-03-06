package ont

import (
	"github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"sync"
)

type OntAsset struct {
	assetIdMap map[string]common.Uint256
	assetMap   map[common.Uint256]*asset.Asset
	lock       sync.RWMutex
}

func NewOntAsset() *OntAsset {
	return &OntAsset{
		assetIdMap: make(map[string]common.Uint256, 0),
		assetMap:   make(map[common.Uint256]*asset.Asset, 0),
	}
}

func (this *OntAsset) RegAsset(assetId common.Uint256, as *asset.Asset) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.assetIdMap[as.Name]; ok {
		return false
	}
	this.assetIdMap[as.Name] = assetId
	this.assetMap[assetId] = as
	return true
}

func (this *OntAsset) GetAssetById(assetId common.Uint256) *asset.Asset {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.assetMap[assetId]
}

func (this *OntAsset) GetAssetByName(name string) *asset.Asset {
	this.lock.RLock()
	defer this.lock.RUnlock()
	assetId, ok := this.assetIdMap[name]
	if !ok {
		return nil
	}
	return this.assetMap[assetId]
}

func (this *OntAsset) GetAssetId(name string) common.Uint256 {
	this.lock.RLock()
	defer this.lock.RUnlock()
	assetId, ok := this.assetIdMap[name]
	if !ok {
		return common.Uint256{}
	}
	return assetId
}
