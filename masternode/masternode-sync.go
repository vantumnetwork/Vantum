package masternode

type CMasternodeSync struct{

}


//masternode sync constructor
func (cMnsync CMasternodeSync) CMasternodeSync() { 

} 

//is synced
func (cMnsync CMasternodeSync) IsSynced() { 

} 


//is blockchain synced
func (cMnsync CMasternodeSync) IsBlockchainSynced() { 

} 


//reset
func (cMnsync CMasternodeSync) Reset() { 

} 

//add masternode to list
func (cMnsync CMasternodeSync) AddedMasternodeList(hash uint256) { 

} 

 //add masternode winner
func (cMnsync CMasternodeSync) AddedMasternodeWinner(hash uint256) { 

} 

//add budget item
func (cMnsync CMasternodeSync) AddedBudgetItem(hash uint256) { 

} 

//is budget propempty
func (cMnsync CMasternodeSync) IsBudgetPropEmpty() { 

} 

//is budget finempty
func (cMnsync CMasternodeSync) IsBudgetFinEmpty() { 

} 

//get next asset
func (cMnsync CMasternodeSync) GetNextAsset() { 

} 


//get sync status
func (cMnsync CMasternodeSync) GetSyncStatus() { 

} 


//get sync status
func (cMnsync CMasternodeSync) ProcessMessage(pfrom CNode,strCommand string,vRecv CDataStream) { 

} 


//clear fulfilled request
func (cMnsync CMasternodeSync) ClearFulfilledRequest() { 

} 


//process
func (cMnsync CMasternodeSync) Process() { 

} 