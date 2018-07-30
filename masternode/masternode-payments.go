package masternode

type CMasternodePayments struct{

}

type CMasternodePaymentDB struct{

}

//write to masternode database
func (cMndb CMasternodePaymentDB) Write(objToSave CMasternodePayments) bool { 

} 

//write to masternode database
func (cMndb CMasternodePaymentDB) Read(objToLoad CMasternodePayments, fDryRun bool ) ReadResult { 

} 

//dump masternode payment
func  DumpMasternodePayments(objToSave CMasternodePayments){ 

} 

//isblock value valid
func  (cMnpayments CMasternodePayments) IsBlockValueValid(block CBlock,nExpectedValue int64){ 

} 
 

//isblockpayee valid
func  (cMnpayments CMasternodePayments) IsBlockPayeeValid(block CBlock,nBlockHeight int64){ 

} 
	
//fill block payee
func  (cMnpayments CMasternodePayments) FillBlockPayee(txNew CMutableTransaction,nFees int64,fProofOfStake bool){ 

} 


//get minimum masternode payments proto
func  (cMnpayments CMasternodePayments) GetMinMasternodePaymentsProto(){ 

} 

//process message masternode payment
func  (cMnpayments CMasternodePayments) ProcessMessageMasternodePayments(pfrom CNode,strCommand string,vRecv CDataStream){ 

} 


//process message masternode payment
func  (cMnpayments CMasternodePayments) ProcessMessageMasternodePayments2(pfrom CNode,strCommand string,vRecv CDataStream){ 

} 

//sign message
func  (cMnpayments CMasternodePayments) Sign(keyMasternode CKey,pubKeyMasternode CPubKey){ 

} 


//get block payee
func  (cMnpayments CMasternodePayments) GetBlockPayee(nBlockHeight int,payee CScript){ 

} 

//is scheduled
func  (cMnpayments CMasternodePayments) IsScheduled(mn CMasternode,nNotBlockHeight int) bool{ 

} 


//add winning masternode payment
func  (cMnpayments CMasternodePayments) AddWinningMasternode(winnerIn CMasternodePaymentWinner) bool{ 

} 
	


//is transaction valid
func  (cMnpayments CMasternodePayments) IsTransactionValid(txNew CTransaction) bool{ 

} 

//get required payments string
func  (cMnpayments CMasternodePayments) GetRequiredPaymentsString() string{ 

} 

//get required payments string
func  (cMnpayments CMasternodePayments) GetRequiredPaymentsString2(nBlockHeight int) string{ 

}


//is transaction valid
func  (cMnpayments CMasternodePayments) IsTransactionValid(txNew CTransaction,nBlockHeight int) string{ 

}
 

//clean payment list
func  (cMnpayments CMasternodePayments) CleanPaymentList(){ 

}


//isvalid
func  (cMnpayments CMasternodePayments) IsValid(pnode CNode,strError string) bool{ 

}


//process block
func  (cMnpayments CMasternodePayments) ProcessBlock(nBlockHeight int) bool{ 

}


//relay
func  (cMnpayments CMasternodePayments) Relay(){ 

}


//signature valid
func  (cMnpayments CMasternodePayments) SignatureValid(){ 

}

//sync
func  (cMnpayments CMasternodePayments) Sync(node CNode,nCountNeeded int){ 

}


//to string
func  (cMnpayments CMasternodePayments) ToString(){ 

}


//get oldest block
func  (cMnpayments CMasternodePayments) GetOldestBlock(){ 

}


//get newest block
func  (cMnpayments CMasternodePayments) GetNewestBlock(){ 

}