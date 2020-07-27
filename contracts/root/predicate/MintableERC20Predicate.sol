pragma solidity >=0.5.0 <0.6.0;

import "../../contracts-package/IERC20.sol";

import "./ERC20Predicate.sol";

contract MintableERC20Predicate is ERC20Predicate {
    constructor(
        address _withdrawManager,
        address _depositManager,
        address _registry
    ) public ERC20Predicate(_withdrawManager, _depositManager, _registry) {}

    function onFinalizeExit(bytes calldata data) external onlyWithdrawManager {
        (
            ,
            address token,
            address exitor,
            uint256 amount
        ) = decodeExitForProcessExit(data);
        uint256 toTransfer = IERC20(token).balanceOf(address(depositManager));
        if (toTransfer > 0) {
            if (toTransfer > amount) {
                toTransfer = amount;
            }
            depositManager.transferAssets(token, exitor, toTransfer);
            amount = amount.sub(toTransfer);
        }
        if (amount > 0) {
            IERC20(token).mint(exitor, amount);
        }
    }
}
