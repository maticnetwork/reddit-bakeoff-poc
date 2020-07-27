pragma solidity >=0.5.0 <0.6.0;

import "./ISubredditPointsParent.sol";
import "./contracts-package/Initializable.sol";
import "./contracts-package/Ownable.sol";
import "./contracts-package/ERC20.sol";
import "./libraries/Address.sol";

// ERC20 and borrows only operators notion from ERC777, accounts can revoke default operator
contract SubredditPointsParent is
    Initializable,
    ISubredditPointsParent,
    Ownable,
    ERC20
{
    using SafeMath for uint256;
    using Address for address;

    event Sent(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data,
        bytes operatorData
    );

    event Minted(
        address indexed operator,
        address indexed to,
        uint256 amount,
        bytes data,
        bytes operatorData
    );

    event Burned(
        address indexed operator,
        address indexed from,
        uint256 amount,
        bytes data,
        bytes operatorData
    );

    event AuthorizedOperator(
        address indexed operator,
        address indexed tokenHolder
    );

    event RevokedOperator(
        address indexed operator,
        address indexed tokenHolder
    );

    event DefaultOperatorAdded(address indexed operator);

    event DefaultOperatorRemoved(address indexed operator);

    // ------------------------------------------------------------------------------------
    // VARIABLES BLOCK, MAKE SURE ONLY ADD TO THE END

    string private _subreddit;
    string private _name;
    string private _symbol;

    // operators notion from ERC777, accounts can revoke default operator
    mapping(address => bool) private _defaultOperators;

    // Maps operators and revoked default operators to state (enabled/disabled)
    mapping(address => mapping(address => bool)) private _operators;
    mapping(address => mapping(address => bool))
        private _revokedDefaultOperators;

    // operators notion from ERC777, accounts can revoke default operator
    // This isn't ever read from - it's only used to respond to the defaultOperators query.
    address[] private _defaultOperatorsArray;

    address public erc20Predicate;

    modifier isERC20Predicate() {
        require(msg.sender == erc20Predicate);
        _;
    }

    // END OF VARS
    // ------------------------------------------------------------------------------------

    function initializecontract(
        address owner_, // owner of the contract
        address erc20Predicate_,
        string calldata subreddit_, // name of the subreddit
        string calldata name_, // name of the token - erc20
        string calldata symbol_, // symbol of the token
        address[] calldata defaultOperators_ // erc777 operators - one is subscriptions contract
    ) external initializer {
        require(
            bytes(subreddit_).length != 0,
            "SubredditPoints: subreddit can't be empty"
        );
        require(
            bytes(name_).length != 0,
            "SubredditPoints: name can't be empty"
        );
        require(
            bytes(symbol_).length != 0,
            "SubredditPoints: symbol can't be empty"
        );
        require(owner_ != address(0), "SubredditPoints: owner should not be 0");

        Ownable.initialize(owner_);
        // UpdatableGSNRecipientSignature.initialize(gsnApprover_);

        erc20Predicate = erc20Predicate_;

        _subreddit = subreddit_;
        _name = name_;
        _symbol = symbol_;

        _defaultOperatorsArray = defaultOperators_;
        for (uint256 i = 0; i < defaultOperators_.length; i++) {
            _defaultOperators[defaultOperators_[i]] = true;
            emit DefaultOperatorAdded(defaultOperators_[i]);
        }
    }

    function mint(address account, uint256 amount) external isERC20Predicate {
        ERC20._mint(account, amount);
    }

    function burn(uint256 amount, bytes calldata userData) external {
        address account = msg.sender;
        _burn(account, account, amount, userData, "");
    }

    function isOperatorFor(address operator, address tokenHolder)
        public
        view
        returns (bool)
    {
        return
            operator == tokenHolder ||
            (_defaultOperators[operator] &&
                !_revokedDefaultOperators[tokenHolder][operator]) ||
            _operators[tokenHolder][operator];
    }

    function authorizeOperator(address operator) external {
        require(
            msg.sender != operator,
            "SubredditPoints: authorizing self as operator"
        );
        require(
            address(0) != operator,
            "SubredditPoints: operator can't have 0 address"
        );

        if (_defaultOperators[operator]) {
            delete _revokedDefaultOperators[msg.sender][operator];
        } else {
            _operators[msg.sender][operator] = true;
        }

        emit AuthorizedOperator(operator, msg.sender);
    }

    function revokeOperator(address operator) external {
        require(
            operator != msg.sender,
            "SubredditPoints: revoking self as operator"
        );
        require(
            address(0) != operator,
            "SubredditPoints: operator can't have 0 address"
        );

        if (_defaultOperators[operator]) {
            _revokedDefaultOperators[msg.sender][operator] = true;
        } else {
            delete _operators[msg.sender][operator];
        }

        emit RevokedOperator(operator, msg.sender);
    }

    function operatorSend(
        address sender,
        address recipient,
        uint256 amount,
        bytes calldata userData,
        bytes calldata operatorData
    ) external {
        address operator = msg.sender;
        require(
            isOperatorFor(operator, sender),
            "SubredditPoints: caller is not an operator for holder"
        );
        _transfer(sender, recipient, amount);
        emit Sent(operator, sender, recipient, amount, userData, operatorData);
    }

    function operatorBurn(
        address account,
        uint256 amount,
        bytes calldata data,
        bytes calldata operatorData
    ) external {
        address operator = msg.sender;
        require(
            isOperatorFor(operator, account),
            "SubredditPoints: caller is not an operator for holder"
        );
        _burn(operator, account, amount, data, operatorData);
    }

    function defaultOperators() external view returns (address[] memory) {
        return _defaultOperatorsArray;
    }

    function addDefaultOperator(address operator) external onlyOwner {
        require(
            operator != address(0),
            "SubredditPoints: operator address shouldn't be 0"
        );
        require(
            !_defaultOperators[operator],
            "SubredditPoints: operator already exists"
        );

        _defaultOperatorsArray.push(operator);
        _defaultOperators[operator] = true;

        emit DefaultOperatorAdded(operator);
    }

    function removeDefaultOperator(address operator) external onlyOwner {
        require(
            operator != address(0),
            "SubredditPoints: operator address shouldn't be 0"
        );
        require(
            _defaultOperators[operator],
            "SubredditPoints: operator doesn't exists"
        );

        for (uint256 i = 0; i < _defaultOperatorsArray.length; i++) {
            if (_defaultOperatorsArray[i] == operator) {
                if (i != (_defaultOperatorsArray.length - 1)) {
                    // if it's not last element, replace it from the tail
                    _defaultOperatorsArray[i] = _defaultOperatorsArray[_defaultOperatorsArray
                        .length - 1];
                }
                _defaultOperatorsArray.length =
                    _defaultOperatorsArray.length -
                    1;
                break;
            }
        }
        delete _defaultOperators[operator];

        emit DefaultOperatorRemoved(operator);
    }

    function _burn(
        address operator,
        address account,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData
    ) internal {
        ERC20._burn(account, amount);
        emit Burned(operator, account, amount, userData, operatorData);
    }

    function subreddit() external view returns (string memory) {
        return _subreddit;
    }

    function name() external view returns (string memory) {
        return _name;
    }

    function symbol() external view returns (string memory) {
        return _symbol;
    }

    // function updateGSNApprover(address gsnApprover) external onlyOwner {
    //     updateSigner(gsnApprover);
    // }

    function updateERC20Predicate(address erc20Predicate_) public onlyOwner {
        require(
            erc20Predicate_ != address(0),
            "SubredditPoints: erc20Predicate should not be 0"
        );
        erc20Predicate = erc20Predicate_;
    }
}
