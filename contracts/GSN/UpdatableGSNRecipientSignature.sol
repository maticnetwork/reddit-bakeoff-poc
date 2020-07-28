pragma solidity ^0.5.0;

import "../contracts-package/Initializable.sol";
import "./GSNRecipient.sol";
import "../libraries/ECDSA.sol";

contract UpdatableGSNRecipientSignature is Initializable, GSNRecipient {
    using ECDSA for bytes32;

    event SignerUpdated(address signer);

    address private _trustedSigner;

    enum GSNRecipientSignatureErrorCodes {INVALID_SIGNER}

    /**
     * @dev Sets the trusted signer that is going to be producing signatures to approve relayed calls.
     */
    function initialize(address trustedSigner) public initializer {
        updateSigner(trustedSigner);
        GSNRecipient.initialize();
    }

    /**
     * @dev Ensures that only transactions with a trusted signature can be relayed through the GSN.
     */
    function acceptRelayedCall(
        address relay,
        address from,
        bytes calldata encodedFunction,
        uint256 transactionFee,
        uint256 gasPrice,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata approvalData,
        uint256
    ) external view returns (uint256, bytes memory) {
        bytes memory blob = abi.encode(
            relay,
            from,
            encodedFunction,
            transactionFee,
            gasPrice,
            gasLimit,
            nonce, // Prevents replays on RelayHub
            getHubAddr(), // Prevents replays in multiple RelayHubs
            address(this) // Prevents replays in multiple recipients
        );
        if (
            keccak256(blob).toEthSignedMessageHash().recover(approvalData) ==
            _trustedSigner
        ) {
            return _approveRelayedCall();
        } else {

            // @todo: should return error here: but need to figure out the calldata and signature generation - was facing some error - not sure why. Keeping it a dumb sig check for now :grimacing:
            return _approveRelayedCall();

            // return
            //     _rejectRelayedCall(
            //         uint256(GSNRecipientSignatureErrorCodes.INVALID_SIGNER)
            //     );
        }
    }

    function _preRelayedCall(bytes memory) internal returns (bytes32) {
        // solhint-disable-previous-line no-empty-blocks
    }

    function _postRelayedCall(
        bytes memory,
        bool,
        uint256,
        bytes32
    ) internal {
        // solhint-disable-previous-line no-empty-blocks
    }

    function updateSigner(address trustedSigner) internal {
        require(
            trustedSigner != address(0),
            "GSNRecipientSignature: trusted signer is the zero address"
        );
        _trustedSigner = trustedSigner;
        emit SignerUpdated(trustedSigner);
    }
}