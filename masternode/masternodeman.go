package masternode

type CMasternodeMan struct{

}

type CompareLastPaid struct
{

}

type CompareScoreTxIn struct {

	
}

type CompareScoreMN struct {

}

//cmasternode DB
func (cMnDb CMasternodeDB) CMasternodeDB(){ 

} 


//cmasternode write
func (cMnDb CMasternodeDB) Write(mnodemanToSave CMasternodeMan){ 

} 



//cmasternode read
func (cMnDb CMasternodeDB) Read(mnodemanToLoad CMasternodeMan,fDryRun bool) ReadResult{ 

} 

//cmasternode dump masternodes
func (cMnDb CMasternodeDB) DumpMasternodes(){ 

} 

//cmasternode manager
func (cMnNode CMasternodeMan) CMasternodeMan(){ 

} 

//cmasternode dump masternodes
func (cMnNode CMasternodeMan) Add(mn CMasternode){ 

} 


//cmasternode dump masternodes
func (cMnNode CMasternodeMan) AskForMN(pnode CNode,vin CTxIn){ 

} 


//check
func (cMnNode CMasternodeMan) Check(){ 

} 

//check and remove
func (cMnNode CMasternodeMan) CheckAndRemove(forceExpiredRemoval bool){ 

} 


//clear
func (cMnNode CMasternodeMan) Clear(){ 

} 

//count enabled
func (cMnNode CMasternodeMan) CountEnabled(protocolVersion int){ 

} 


//DsegUpdate
func (cMnNode CMasternodeMan) DsegUpdate(pnode CNode){ 

} 


//Find
func (cMnNode CMasternodeMan) Find(payee CScript){ 

} 


//Find
func (cMnNode CMasternodeMan) Find(vin CTxIn){ 

} 

//Find
func (cMnNode CMasternodeMan) Find(pubKeyMasternode CPubKey){ 

} 


//GetNextMasternodeInQueueForPayment
func (cMnNode CMasternodeMan) GetNextMasternodeInQueueForPayment(nBlockHeight int,fFilterSigTime bool,nCount int){ 

} 

//FindRandomNotInVec
func (cMnNode CMasternodeMan) FindRandomNotInVec(vecToExclude CTxIn,protocolVersion int){ 

} 

//FindRandomNotInVec
func (cMnNode CMasternodeMan) GetCurrentMasterNode(mod int,nBlockHeight int64,minProtocol int){ 

} 


//GetMasternodeRank
func (cMnNode CMasternodeMan) GetMasternodeRank(vin CTxIn,nBlockHeight int64,minProtocol int,fOnlyActive bool){ 


} 

//GetMasternodeRanks
func (cMnNode CMasternodeMan) GetMasternodeRanks(nBlockHeight int64,minProtocol int){ 


} 


//GetMasternodeRank
func (cMnNode CMasternodeMan) GetMasternodeByRank(nRank int,nBlockHeight int64,minProtocol int,fOnlyActive bool){ 


} 


//ProcessMasternodeConnections
func (cMnNode CMasternodeMan) ProcessMasternodeConnections(){ 


} 

//ProcessMessage
func (cMnNode CMasternodeMan) ProcessMessage(pfrom CNode,strCommand string,vRecv CDataStream){ 


} 

//ProcessMasternodeConnections
func (cMnNode CMasternodeMan) Remove(vin CTxIn){ 


} 


//UpdateMasternodeList
func (cMnNode CMasternodeMan) UpdateMasternodeList(mnb CMasternodeBroadcast){ 


} 

//ToString
func (cMnNode CMasternodeMan) ToString() string{ 


} 


