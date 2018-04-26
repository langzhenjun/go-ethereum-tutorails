pragma solidity ^0.4.21;


contract AccessControl {

    address public ceoAddress;
    address public cfoAddress;
    address public cooAddress;

    // 游戏状态变量
    bool public paused = false;

    /// 表示CEO权限
    modifier onlyCEO() {
        require(msg.sender == ceoAddress);
        _;
    }

    /// 表示CFO权限
    modifier onlyCFO() {
        require(msg.sender == cfoAddress);
        _;
    }

    /// 表示COO权限
    modifier onlyCOO() {
        require(msg.sender == cooAddress);
        _;
    }

    // 表示CEO/CFO/COO都有权限
    modifier onlyCLevel() {
        require(
            msg.sender == ceoAddress ||
            msg.sender == cfoAddress ||
            msg.sender == cooAddress
        );
        _;
    }

    // 取现
    // 只有CFO有权限
    function withdrawBalance() external onlyCFO {
        cfoAddress.transfer(this.balance);
    }
    
    // 设置CEO
    // 只有CEO有权限
    function setCEO(address _newCEO) public onlyCEO {
        require(_newCEO != address(0));
        ceoAddress = _newCEO;
    }

    // 设置CFO
    // 只有CEO有权限
    function setCFO(address _newCFO) public onlyCEO {
        require(_newCFO != address(0));
        cfoAddress = _newCFO;
    }

    // 设置COO
    // 只有CEO有权限
    function setCOO(address _newCOO) public onlyCEO {
        require(_newCOO != address(0));
        cooAddress = _newCOO;
    }

    // 表示游戏状态为运行状态
    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    // 表示游戏状态为暂停状态
    modifier whenPaused() {
        require(paused);
        _;
    }

    // 暂停游戏
    // 身份要求：CEO/CFO/COO
    function pause() public onlyCLevel whenNotPaused {
        paused = true;
    }

    // 恢复游戏
    // 只有CEO才可以重新开启游戏
    function resume() public onlyCEO whenPaused {
        paused = false;
    }
}