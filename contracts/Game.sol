pragma solidity ^0.4.21;

import "./AccessControl.sol";


contract Game is AccessControl {

    address public newContractAddress;

    event GameUpgrade(address newContract);

    function () public payable {}

    /// 游戏更新
    /// 使用新的合约
    function upgrade(address _newAddress) public onlyCEO whenPaused {
        newContractAddress = _newAddress;
        emit GameUpgrade(_newAddress);
    }
    
}