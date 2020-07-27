pragma solidity >=0.5.0 <0.6.0;

contract ISubredditPointsParent {
    function mint(address account, uint256 amount) external; // solium-disable-line indentation

    function burn(uint256 amount, bytes calldata data) external; // solium-disable-line indentation

    function operatorBurn(
        address account,
        uint256 amount,
        bytes calldata data,
        bytes calldata operatorData
    ) external; // solium-disable-line indentation

    function subreddit() external view returns (string memory);
}
