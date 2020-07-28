pragma solidity >=0.5.0 < 0.6.0;

import "./contracts-package/Initializable.sol";
import "./contracts-package/Ownable.sol";

import "./libraries/SafeMath.sol";
import "./libraries/Address.sol";
import "./ISubredditPoints.sol";

import "./contracts-package/IERC20.sol";

import "./GSN/UpdatableGSNRecipientSignature.sol";

contract Distributions_v0 is Initializable, Ownable, UpdatableGSNRecipientSignature {

    struct DistributionRound {
        uint256 availablePoints;
        uint256 totalKarma;
    }

    event AdvanceRound(uint256 round, uint256 totalPoints);
    event ClaimPoints(uint256 round, address indexed user, uint256 points);


    using SafeMath for uint256;
    using Address for address;


    address public subredditPoints;
    string public subreddit;
    string private _subredditLowerCase;

    mapping(uint256 => DistributionRound) private _distributionRounds;

    uint256 public initialSupply;

    function initialize (
        address owner_,
        address subredditPoints_,
        uint256 initialSupply_,
        uint256 initialKarma_,
        address gsnApprover_
    ) external initializer {

        require(owner_ != address(0), "Distributions: owner should not be 0");

        Ownable.initialize(owner_);

        subredditPoints = subredditPoints_;

        subreddit = ISubredditPoints(subredditPoints_).subreddit();
        _subredditLowerCase = _toLower(subreddit);

        UpdatableGSNRecipientSignature.initialize(gsnApprover_);

        initialSupply = initialSupply_;
        _distributionRounds[0] = DistributionRound({
            availablePoints: initialSupply,
            totalKarma: initialKarma_
        });

        emit AdvanceRound(0, initialSupply);
    }

    function claim(uint256 round, address account, uint256 karma) external {
        DistributionRound memory dr = _distributionRounds[round];
        require(dr.availablePoints > 0, "Distributions: no points to claim in this round");
        require(dr.totalKarma > 0, "Distributions: this round has no karma");
        uint256 userPoints = dr.availablePoints
            .mul(karma)
            .div(dr.totalKarma);
        require(userPoints > 0, "Distributions: user karma is too low to claim points");
        emit ClaimPoints(round, account, userPoints);
        ISubredditPoints(subredditPoints).mint(address(this), account, userPoints, "", "");
    }

    function advanceToRound(uint256 round, uint256 _roundPoints,uint256 _totalKarma) external {

        require(_totalKarma > 0, "Distributions: totalKarma should be > 0");

        // @poc: add new distribution round
        _distributionRounds[round] = DistributionRound({
            availablePoints: _roundPoints,
            totalKarma: _totalKarma
        });

        emit AdvanceRound(round, _roundPoints);

    }

    function _toLower(string memory str) internal pure returns (string memory) {
        bytes memory bStr = bytes(str);
        bytes memory bLower = new bytes(bStr.length);
        for (uint i = 0; i < bStr.length; i++) {
            if ((int8(bStr[i]) >= 65) && (int8(bStr[i]) <= 90)) {
                bLower[i] = bytes1(int8(bStr[i]) + 32);
            } else {
                bLower[i] = bStr[i];
            }
        }
        return string(bLower);
    }
}