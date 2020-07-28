pragma solidity ^0.5.0;

import "./GSN/UpdatableGSNRecipientSignature.sol";
import "./contracts-package/Initializable.sol";

contract Counter is Initializable, UpdatableGSNRecipientSignature {
  //it keeps a count to demonstrate stage changes
  uint private count;
  address private _owner;

  function initialize (uint num, address trustedSigner) public initializer {
    UpdatableGSNRecipientSignature.initialize(trustedSigner);
    _owner = _msgSender();
    count = num;
  }

  function _preRelayedCall(bytes memory context) internal returns (bytes32) {
  }

  function _postRelayedCall(bytes memory context, bool, uint256 actualCharge, bytes32) internal {
  }

  function owner() public view returns (address) {
    return _owner;
  }

  // getter
  function getCounter() public view returns (uint) {
    return count;
  }

  //and it can add to a count
  function increaseCounter(uint256 amount) public {
    count = count + amount;
  }

  //We'll upgrade the contract with this function after deploying it
  //Function to decrease the counter
  function decreaseCounter(uint256 amount) public returns (bool) {
    require(count > amount, "Cannot be lower than 0");
    count = count - amount;
    return true;
  }

  function setRelayHubAddress() public {
    if(getHubAddr() == address(0)) {
      _upgradeRelayHub(0xD216153c06E857cD7f72665E0aF1d7D82172F494);
    }
  }

  function getRecipientBalance() public view returns (uint) {
    return IRelayHub(getHubAddr()).balanceOf(address(this));
  }

}
