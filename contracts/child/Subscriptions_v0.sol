pragma solidity >=0.5.0 < 0.6.0;

import "../contracts-package/Initializable.sol";
import "../contracts-package/Ownable.sol";
import "../contracts-package/ERC20.sol";
import "./ISubredditPoints.sol";
import "../libraries/Address.sol";


contract Subscriptions_v0 is Initializable, Ownable {
    using SafeMath for uint256;
    using Address for address;

    event Subscribed(address indexed recipient, address indexed payer, uint256 burnedPoints, uint256 expiresAt, bool renewable);
    event Canceled(address indexed recipient, uint256 expiresAt);
    event DurationUpdated(uint256 duration);
    event PriceUpdated(uint256 price);
    event RenewBeforeUpdated(uint256 renewBefore);

    uint256 public renewBeforeVal;
    uint256 public durationVal;
    uint256 public priceVal;

    // maps address to expiration time
    mapping(address => uint256) public _subscriptions;
    // maps address of recipient to address of payer
    mapping(address => address) public _payers;
    address public _subredditPoints;

    function initializecontract(
        address owner_,
        address subredditPoints,
        uint256 price_,
        uint256 duration_,
        uint256 renewBefore
    ) external initializer {
        require(owner_ != address(0), "Subscriptions: owner should not be 0");

        Ownable.initialize(owner_);
        _subredditPoints = subredditPoints;

        updatePrice(price_);
        updateDuration(duration_);
        updateRenewBefore(renewBefore);
    }

    function cancel(address recipient) external {
        address payer = _payers[recipient];
        require(msg.sender == payer || msg.sender == recipient, "Subscriptions: subscription can be cancelled by payer or recipient only");
        delete _payers[recipient];
        emit Canceled(recipient, _subscriptions[recipient]);
    }

    function renew(address recipient) external {
        address payer = _payers[recipient];
        require(payer != address(0), "Subscriptions: subscription is canceled");
        // solium-disable-next-line security/no-block-members
        require(expiration(recipient) < block.timestamp.add(renewBeforeVal), "Subscriptions: too early to renew");
        _subscribe(payer, recipient, true);
    }

    function subscribe(address recipient, bool renewable) external {
        address payer = msg.sender;
        if (renewable) {
            _payers[recipient] = payer;
        }
        _subscribe(payer, recipient, renewable);
    }

    function _subscribe(address payer, address recipient, bool renewable) internal {
        require(recipient != address(0), "Subscriptions: recipient should not be 0");
        uint256 expirationAt = _subscriptions[recipient];
        // solium-disable-next-line security/no-block-members
        if (expirationAt < block.timestamp) {
            // solium-disable-next-line security/no-block-members
            expirationAt = block.timestamp;
        }
        uint256 newExpiration = expirationAt.add(durationVal);
        _subscriptions[recipient] = newExpiration;
        emit Subscribed(recipient, payer, priceVal, newExpiration, renewable);
        ISubredditPoints(_subredditPoints).operatorBurn(payer, priceVal, "", "");
    }

    function updateDuration(uint256 duration_) public onlyOwner {
        require(duration_ > 0, "Subscriptions: duration should be > 0");
        durationVal = duration_;
        emit DurationUpdated(duration_);
    }

    function updatePrice(uint256 price_) public onlyOwner {
        require(price_ > 0, "Subscriptions: price should be > 0");
        priceVal = price_;
        emit PriceUpdated(price_);
    }

    function updateRenewBefore(uint256 renewBefore) public onlyOwner {
        require(renewBefore > 0, "Subscriptions: renewBefore should be > 0");
        renewBeforeVal = renewBefore;
        emit RenewBeforeUpdated(renewBefore);
    }

    function duration() external view returns (uint256) {
        return durationVal;
    }

    function price() external view returns (uint256) {
        return priceVal;
    }

    function renewBefore() external view returns (uint256) {
        return renewBeforeVal;
    }

    function expiration(address account) public view returns (uint256) {
        return _subscriptions[account];
    }
}