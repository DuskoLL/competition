// SPDX-License-Identifier: MIT
pragma solidity 0.4.25;

contract Insurances {

    event NewInsurance(uint256 indexed _insuranceId, address indexed _beneficiary);
    event CensorInsurance(uint256 indexed _insuranceId, bool result);
    event ExecuteInsurance(uint256 indexed _insuranceId, uint value);
    event InsuranceExpired(uint256 indexed _insuranceId);
    event InsuranceNotExpired(uint256 indexed _insuranceId);
    event Withdraw(address indexed _adminAddress, uint value);

    function receive() external payable{}
    function fallback() external payable{}

    struct Insurance {
        uint256 insuranceId;
        address beneficiaryAddress;
        uint compensation;
        bool insuranceStatus;
        bool censorStatues;
        uint endTime;
        bool isExpired;
    }
    struct WeatherInfo {
        string province;
        string city;
        string weather;
        int8 temperature;
        string winddirection;
        string windpower;
        int8 humidity;
    }


    Insurance[] public insurances;
    WeatherInfo[] public weatherInfos;
    address private adminAddress;
    bool private lock;
    uint public totalCompensation;
    
    mapping(uint256 => uint256) public index;
    mapping(uint256 => bool) public isExist;
    mapping(address => uint) public insuranceIds;

    modifier isOnlyAdmin {
        require(msg.sender == adminAddress, "not admin");
        _;
    }
    modifier isAdminOrBeneficiary(uint256 _insuranceId) {
        require(isExist[_insuranceId], "insurance not exist");
        Insurance memory insurance = insurances[index[_insuranceId]];
        require(msg.sender == adminAddress || msg.sender == insurance.beneficiaryAddress, "not Admin or beneficiary");
        _;
    }

    constructor() internal payable{
        adminAddress = msg.sender;
    }

    function createInsurance(address _beneficiary, uint _compenstation, uint duration) external isOnlyAdmin  returns(uint insuranceId){
        require(totalCompensation <= address(this).balance, "accountBalance < totalCompensation");
        uint _endTime = block.timestamp + duration;
        uint256 _insuranceId = uint256(keccak256(abi.encode(_beneficiary, _compenstation, _endTime)));
        require(!isExist[_insuranceId], "insurance exist");
        Insurance memory _insurance = Insurance(_insuranceId, _beneficiary, _compenstation, false, false, _endTime, false);
        insurances.push(_insurance);
        isExist[_insuranceId] = true;
        index[_insuranceId] = insurances.length - 1;
        totalCompensation += _compenstation;
        emit NewInsurance(_insuranceId, _beneficiary);
        insuranceIds[_beneficiary] = _insuranceId;
        insuranceId = _insuranceId;
    }

    function censorInsurance(uint256 _insuranceId) external isAdminOrBeneficiary(_insuranceId) {
        if (!checkEndTime(_insuranceId)){
            return;
        }
        Insurance storage _insurance = insurances[index[_insuranceId]];
        require(!_insurance.insuranceStatus, "insurance is excted");
        //预言机部分，result直接获取结果
        bool result = true;
        if (result) {
            _insurance.censorStatues = true;
        }
        emit CensorInsurance(_insuranceId, result);
    }

    function executeInsurance(uint256 _insuranceId) external isAdminOrBeneficiary(_insuranceId) {
        if (!checkEndTime(_insuranceId)){
            return;
        }
        Insurance storage _insurance = insurances[index[_insuranceId]];
        require(!lock, "lock is locked");
        lock = true;
        if (_insurance.insuranceStatus) {
            revert(string(abi.encodePacked(_insuranceId, " is excuted")));
            lock = false;
            return;
        }
        if (!_insurance.censorStatues) {
            revert(string(abi.encodePacked(_insuranceId, " is not censored")));
            lock = false;
            return;
        }
        uint _compensation = _insurance.compensation;
        address _beneficiary = _insurance.beneficiaryAddress;
        bool success = _beneficiary.call.value(_compensation)("");
        require(success, "failed call");
        _insurance.insuranceStatus = true;
        totalCompensation -= _compensation;
        lock = false;
        emit ExecuteInsurance(_insuranceId, _compensation);
    }

    function findInsurance(uint256 _insuranceId) external view returns (uint256 insuranceId, address beneficiaryAddress, uint compensation, bool insuranceStatus, bool censorStatues, uint endTime) {
        require(isExist[_insuranceId], "insurance do not exist");
        Insurance memory _insurance = insurances[index[_insuranceId]];
        insuranceId = _insuranceId;
        beneficiaryAddress = _insurance.beneficiaryAddress;
        compensation = _insurance.compensation;
        insuranceStatus = _insurance.insuranceStatus;
        censorStatues = _insurance.censorStatues;
        endTime = _insurance.endTime;
    }

    function uploadWeathers(
        string _province,
        string _city,
        string _weather,
        int8 _temperature,
        string _winddirection,
        string _windpower,
        int8 _humidity) external {

        WeatherInfo memory _weatherInfo = WeatherInfo(_province,_city, _weather, _temperature, _winddirection, _windpower, _humidity);
        weatherInfos.push(_weatherInfo);
    }

    function withdraw() external isOnlyAdmin{
        uint accountBalance = address(this).balance;
        require(totalCompensation < accountBalance, "can not withdraw");        
        bool success = msg.sender.call.value(accountBalance - totalCompensation)("");
        require(success, "failed call");
        emit Withdraw(address(msg.sender), accountBalance - totalCompensation);
    }

    function checkEndTime(uint256 _insuranceId) internal returns (bool){
        // require(isExist[_insuranceId], "insurance not exist");
        Insurance memory _insurance = insurances[index[_insuranceId]];
        // require(!_insurance.insuranceStatus, "insurance is excuted");
        if (_insurance.isExpired) {
            emit InsuranceExpired(_insuranceId);
            return false;
        }
        if (_insurance.endTime < block.timestamp) {
            _insurance.isExpired = true;
            totalCompensation -= _insurance.compensation;
            emit InsuranceExpired(_insuranceId);
            return false;
        }
        emit InsuranceNotExpired(_insuranceId);
        return true;       
    }

    function getBalance() external view returns(uint) {
        return address(this).balance;
    }

}