package masternode

type CActiveMasternode struct{

}
//bootup the masternode
//verify vantum collateral and register on network
func (cMn CActiveMasternode) ManageStatus() { 

} 

//get masternode status
func (cMn CActiveMasternode) GetStatus() string { 

} 

//get masternode status
func (cMn CActiveMasternode) SendMasternodePing(errormessage string) bool { 

} 

//register masternode
func (cMn CActiveMasternode) Register(strService string,strKeyMasternode string,strTxHash string,strOutputIndex string,errorMessage string) bool { 

} 

//register masternode
func (cMn CActiveMasternode) Register2(vin CTxIn,service CService,keyCollateralAddress CKey,pubKeyCollateralAddress CPubKey,keyMasternode CKey,pubKeyMasternode CPubKey,errorMessage string) bool { 

} 
 

//get masternode vin
func (cMn CActiveMasternode) GetMasterNodeVin(vin CTxIn,pubkey CPubKey,secretKey CKey) bool { 

} 


//get masternode vin
func (cMn CActiveMasternode) GetMasterNodeVin2(vin CTxIn,pubkey CPubKey,secretKey CKey,strTxHash string,strOutputIndex string) bool { 

} 

//get vin from output
func (cMn CActiveMasternode) GetVinFromOutput(out COutput,vin CTxIn,pubkey CPubKey,secretKey CKey) bool { 

} 

//get all possible outputs for running Masternode
func (cMn CActiveMasternode) SelectCoinsMasternode() void { 

} 

// when starting a Masternode, this can enable to run as a hot wallet with no funds
func (cMn CActiveMasternode) EnableHotColdMasterNode(newVin CTxIn,newService CService) bool { 

} 
	
 